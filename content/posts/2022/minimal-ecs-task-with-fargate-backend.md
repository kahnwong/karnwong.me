---
title: Minimal ECS task with fargate backend
date: 2022-08-26T21:10:58+07:00
draft: false
ShowToc: true
images:
tags:
  - devops
  - terraform
  - aws
---

To deploy a web application, there are many ways to go about it. I could spin up a bare VM and set up the environment manually. To make things easier, I could have package the app into docker image. But this still means I have to "update" the app manually if I add changes to it.

Things would be super cool if: after I push the changes to master branch, the app would be deployed automatically. In order to achieve this, I could use AWS ECS task to deploy the app, and add CI/CD to it (because this is 2022 after all).

And things would be even better if I don't have to set up the infra manually every time I want to deploy an app, enters terraform!

Below are minimal ecs task with fargate backend setup 😎. Repo [here](https://github.com/devbaygroup/terraform-aws-ecs-fargate-example).

**Updated 2022-09-02**

Notes: you might need to set up autoscaling on LB connections per target. Also this example contains two target tracking policies for the same service. Race conditions can result in undesirable scaling issues. Thanks John Mille!

## Task definition

This is equivalent to docker-compose.yaml

```hcl
resource "aws_cloudwatch_log_group" "this" {
  retention_in_days = 14
  name              = "/aws/ecs/${var.service_name}"
}


resource "aws_ecs_task_definition" "this" {
  family                   = var.service_name
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = 256
  memory                   = 512
  execution_role_arn       = var.task_role
  task_role_arn            = var.task_role
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
  name                               = var.service_name
  cluster                            = var.ecs_cluster_id
  task_definition                    = aws_ecs_task_definition.this.arn
  desired_count                      = 1
  deployment_minimum_healthy_percent = 50
  deployment_maximum_percent         = 200
  launch_type                        = "FARGATE"
  scheduling_strategy                = "REPLICA"

  network_configuration {
    security_groups  = [var.alb_id]
    subnets          = var.subnet_id
    assign_public_ip = true
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
resource "aws_lb" "this" {
  name               = var.service_name
  internal           = false
  load_balancer_type = "application"
  security_groups    = [var.alb_id]
  subnets            = var.subnet_id
  idle_timeout       = 3600

  enable_deletion_protection = true
}


resource "aws_alb_target_group" "this" {
  name        = var.service_name
  port        = 80
  protocol    = "HTTP"
  vpc_id      = var.vpc_id
  target_type = "ip"

  health_check {
    healthy_threshold   = "3"
    interval            = "30"
    protocol            = "HTTP"
    matcher             = "200"
    timeout             = "3"
    path                = var.health_check_path
    unhealthy_threshold = "2"
  }
}


resource "aws_alb_listener" "http" {
  load_balancer_arn = aws_lb.this.id
  port              = 80
  protocol          = "HTTP"

  default_action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.this.arn
  }
}


resource "aws_alb_listener" "https" {
  load_balancer_arn = aws_lb.this.id
  port              = 443
  protocol          = "HTTPS"

  ssl_policy      = "ELBSecurityPolicy-2016-08"
  certificate_arn = aws_acm_certificate.this.arn

  default_action {
    target_group_arn = aws_alb_target_group.this.id
    type             = "forward"
  }
}
```

## Autoscaling

Because we are using cloud, and I love taking advantage of dynamic resources allocation.

```hcl
resource "aws_appautoscaling_target" "this" {
  max_capacity       = 2
  min_capacity       = 1
  resource_id        = "service/${var.ecs_cluster_name}/${aws_ecs_service.this.name}"
  scalable_dimension = "ecs:service:DesiredCount"
  service_namespace  = "ecs"
}


resource "aws_appautoscaling_policy" "memory" {
  name               = "memory-autoscaling"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.this.resource_id
  scalable_dimension = aws_appautoscaling_target.this.scalable_dimension
  service_namespace  = aws_appautoscaling_target.this.service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageMemoryUtilization"
    }

    target_value = 40
  }
}


resource "aws_appautoscaling_policy" "cpu" {
  name               = "cpu-autoscaling"
  policy_type        = "TargetTrackingScaling"
  resource_id        = aws_appautoscaling_target.this.resource_id
  scalable_dimension = aws_appautoscaling_target.this.scalable_dimension
  service_namespace  = aws_appautoscaling_target.this.service_namespace

  target_tracking_scaling_policy_configuration {
    predefined_metric_specification {
      predefined_metric_type = "ECSServiceAverageCPUUtilization"
    }
    target_value = 60
  }
}
```
