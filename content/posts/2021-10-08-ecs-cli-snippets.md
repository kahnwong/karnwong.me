+++
title = "ecs-cli snippets"
date = "2021-10-08"
path = "/posts/2021/10/ecs-cli-snippets"

[taxonomies]
categories = [ "infrastructure",]
tags = [ "aws",]

+++

```bash
ecs-cli configure profile \
  --access-key $KEY \
  --secret-key $SECRET \
  --profile-name $PROFILE

### launch mode: fargate
ecs-cli configure \
  --cluster $CLUSTER \
  --default-launch-type FARGATE \
  --config-name $NAME \
  --region ap-southeast-1

ecs-cli up \
  --cluster-config $NAME \
  --vpc $VPCID\
  --subnets $SUBNETID1, $SUBNETID2

### launch mode: ec2
ecs-cli configure \
  --cluster $CLUSTER \
  --region ap-southeast-1 \
  --default-launch-type EC2 \
  --config-name $NAME

ecs-cli up --keypair $KEYPAIR \
  --extra-user-data userData.sh \
  --capability-iam --size 1 \
  --instance-type t2.large \
  --cluster-config $NAME \
  --verbose \
  --force \
  --aws-profile $PROFILE

ecs-cli compose \
  --cluster-config $NAME \
  --file docker-compose.yml up \
  --create-log-groups
```
