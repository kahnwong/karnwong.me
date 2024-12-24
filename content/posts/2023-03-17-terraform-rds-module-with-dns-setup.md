+++
title = "Terraform RDS module with DNS setup"
date = "2023-03-17"
path = "/posts/2023/03/terraform-rds-module-with-dns-setup"

[taxonomies]
categories = ["infrastructure"]
tags = [ "terraform", "aws",]

+++

I love not having to manage databases. Hosting it on your compute is guaranteed to be cheaper, but I don't want to be constantly worrying about backups and database upgrade / maintenance.

AWS offers managed databases, known as RDS. Mostly I use postgres, which works well for most use cases. (It's also a given that - if you don't really need NoSQL database, then don't use it).

If I spin up RDS postgres, I can reach it assuming I have set up the correct networking pathways. (And you shouldn't expose your database to public, unless it's for playground with no sensitive data.)

RDS itself returns a hostname with following format: `mydb.123456789012.us-east-1.rds.amazonaws.com`. This works well for all intent and purposes, but there are some cases where you might want it to be easy to remember, or that you want to switch database instance without incurring downtime. A simple solution is to setup a DNS record, but it would be nice if you can set it all up during database creation!

Enough intro, below is the terraform code (and yes, this works for a replica as well):

```hcl
resource "aws_db_instance" "db" {
  engine            = var.engine
  allocated_storage = var.allocated_storage
  engine_version    = var.engine_version
  instance_class    = var.instance_class

  identifier = var.identifier
  db_name    = var.db_name
  username   = var.username
  password   = var.password

  replicate_source_db = var.source_db_id # for replica

  parameter_group_name = var.parameter_group_name

  performance_insights_enabled    = true
  performance_insights_kms_key_id = var.performance_insights_kms_key_id

  deletion_protection = var.deletion_protection

  ### networking
  vpc_security_group_ids = var.vpc_security_group_ids
  db_subnet_group_name   = var.db_subnet_group_name

  ### backup (in UTC)
  maintenance_window      = var.maintenance_window
  backup_window           = var.backup_window
  backup_retention_period = var.backup_window != null ? 3 : 0 # Backups are required in order to create a replica
  skip_final_snapshot     = var.skip_final_snapshot

  apply_immediately = var.apply_immediately
}

resource "aws_route53_record" "db" {
  zone_id = var.zone_id
  name    = var.dns_name
  type    = "CNAME"
  ttl     = 5
  records = [aws_db_instance.db.address]
}
```
