+++
title = "pglogical setup"
date = "2023-07-20"
path = "/posts/2023/07/pglogical-setup"

[taxonomies]
categories = ["devops",]
tags = [  "postgres", "database",]

+++

In certain cases, you can't do a full postgres replication to another instance, or you prefer a fine-grained control on what to replicate, pglogical is one way to achieve partial replication, albeit this requires more manual setup.

Below are steps I used to do a pglogical replication from AWS RDS to on-premise database.

Note: If a subscriber (from the above example, the on-premise database) is offline, postgres WAL would balloon up. You'll have to remove all traces of pglogical extension, including uninstalling pglogical extension, then reinitialize everything again to resolve the problem.

## 1. ON PROVIDER INSTANCE: Create a role for subscriber to fetch data from provider instance

```sql
CREATE USER $SUBSCRIBER_ROLE_NAME WITH ENCRYPTED PASSWORD $SUBSCRIBER_ROLE_PASSWORD;
GRANT rds_superuser TO $SUBSCRIBER_ROLE_NAME; -- RDS specific

GRANT CONNECT ON DATABASE $DATABASE_TO_BE_REPLICATED TO $SUBSCRIBER_ROLE_NAME;
GRANT USAGE ON SCHEMA public TO $SUBSCRIBER_ROLE_NAME;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO $SUBSCRIBER_ROLE_NAME;
GRANT rds_replication TO $SUBSCRIBER_ROLE_NAME; -- RDS specific
```

## 2. ON PROVIDER INSTANCE: Init pglogical extension

```sql
CREATE EXTENSION pglogical;
GRANT ALL ON SCHEMA pglogical TO $SUBSCRIBER_ROLE_NAME;
```

## 3. ON PROVIDER INSTANCE: Define provider node

```sql
SELECT pglogical.create_node(
  node_name := 'provider',
  dsn := 'host=$PROVIDER_INSTANCE_HOSTNAME port=5432 user=$SUBSCRIBER_ROLE_NAME dbname=$DATABASE_TO_BE_REPLICATED password=$SUBSCRIBER_ROLE_PASSWORD'
);
```

## 4. ON PROVIDER INSTANCE: Define tables to be replicated

```sql
select pglogical.create_replication_set('replication_set');
select pglogical.replication_set_add_table(
  set_name := 'replication_set',
  relation := '$TABLE_NAME',
  synchronize_data := true
);
```

## 5. ON SUBSCRIBER INSTANCE: Init pglogical extension

```sql
CREATE EXTENSION pglogical;
```

## 6. ON SUBSCRIBER INSTANCE: Define table schemas for replicated tables

pglogical does not transmit schema definition. Basically you need to do the equivalent of schema migration.

```sql
create table public.$TABLE_NAME
(
    subject_id                   varchar not null,
    foo                          integer,
    bar                          varchar,
    primary key (id)
);

-- also define indexes as well,
```

## 7. ON SUBSCRIBER INSTANCE: Define subscriber node

`$SUBSCRIBER_INSTANCE_HOSTNAME` can be `localhost`

```sql
SELECT pglogical.create_node(
  node_name := 'subscriber',
  dsn := 'host=$SUBSCRIBER_INSTANCE_HOSTNAME port=5432 user=$SUBSCRIBER_INSTANCE_ROLE_NAME dbname=$DATABASE_TO_BE_REPLICATED password=$SUBSCRIBER_INSTANCE_ROLE_PASSWORD'
);
```

## 8. ON SUBSCRIBER INSTANCE: Define subscription

```sql
SELECT pglogical.create_subscription(
    subscription_name := 'aws_sub',
    provider_dsn := 'host=$PROVIDER_INSTANCE_HOSTNAME port=5432 dbname=$DATABASE_TO_BE_REPLICATED user=$SUBSCRIBER_ROLE_NAME password=$SUBSCRIBER_ROLE_PASSWORD',
    replication_sets := ARRAY['replication_set']
);
```

## 9. ON PROVIDER INSTANCE: Verify that pglogical is working

```sql
SELECT
  subscription_name,
  status
FROM
  pglogical.show_subscription_status ();
```
