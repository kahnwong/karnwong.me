---
title: Terraform with ECS task on EC2 backend
date: 2022-10-04T22:04:34+07:00
draft: false
ShowToc: false
images:
tags:
  - devops
  - terraform
  - aws
---

[Previously]({{< ref "/posts/minimal-ecs-task-with-fargate-backend" >}}) I wrote about setting up ECS task on fargate backend. But we can also use EC2 as backend too, in some cases where the workload is consistent, ie scaling is not required, since EC2 would be cheaper than fargate backend, even more so if you have reserved instance on top. There's a few modifications from the fargate version to make it work with EC2 backend, if you are curious you can try to hunt those down 😎. Repo [here](https://github.com/devbaygroup/terraform-aws-ecs-ec2-example).

## Task definition

```hcl
#tfsec:ignore:aws-cloudwatch-log-group-customer-key
resource "aws_cloudwatch_log_group" "this" {
  retention_in_days = 14
  name              = "/aws/ecs/${var.service_name}"
}


resource "aws_ecs_task_definition" "this" {
  family             = var.service_name
  cpu                = var.cpu
  memory             = var.memory
  execution_role_arn = var.task_role
  task_role_arn      = var.task_role

  container_definitions = jsonencode(
    [
      {
        name        = var.service_name
        image       = var.image_uri
        essential   = true
        environment = []

        portMappings = [
          {
            protocol      = "tcp"
            containerPort = 80
            hostPort      = 80
          }
        ]
        logConfiguration = {
          logDriver = "awslogs"
          options = {
            awslogs-group         = aws_cloudwatch_log_group.this.name
            awslogs-region        = var.aws_region
            awslogs-stream-prefix = "ecs"
          }
        }
      }
    ]
  )
}

resource "aws_ecs_service" "this" {
  name            = var.service_name
  cluster         = var.ecs_cluster_id
  task_definition = aws_ecs_task_definition.this.arn
  desired_count   = 1

  # https://github.com/hashicorp/terraform/issues/26950
  depends_on = [aws_lb.this, aws_alb_target_group.this]

  ordered_placement_strategy {
    type  = "spread"
    field = "instanceId"
  }

  load_balancer {
    target_group_arn = aws_alb_target_group.this.arn
    container_name   = var.service_name
    container_port   = 80
  }
}
```

## SSL certificate

```hcl
resource "aws_acm_certificate" "this" {
  domain_name       = var.domain_name
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
  }

  tags = {
    Name = var.domain_name
  }
}
```

## Load balancer

```hcl
#tfsec:ignore:aws-elb-alb-not-public
#tfsec:ignore:aws-elb-drop-invalid-headers
resource "aws_lb" "this" {
  name               = var.service_name
  internal           = false
  load_balancer_type = "application"
  security_groups    = [var.alb_id]
  subnets            = var.subnet_id
}


resource "aws_alb_target_group" "this" {
  name        = var.service_name
  port        = 80
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "instance"

  health_check {
    port    = "traffic-port"
    path    = var.health_check_path
    matcher = "200-499"
  }
}


#tfsec:ignore:aws-elb-http-not-used
resource "aws_alb_listener" "http" {
  load_balancer_arn = aws_lb.this.id
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}


#tfsec:ignore:aws-elb-use-secure-tls-policy
resource "aws_alb_listener" "https" {
  load_balancer_arn = aws_lb.this.id
  port              = 443
  protocol          = "HTTPS"

  ssl_policy      = "ELBSecurityPolicy-2016-08"
  certificate_arn = aws_acm_certificate.this.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.this.id
  }
}
```
