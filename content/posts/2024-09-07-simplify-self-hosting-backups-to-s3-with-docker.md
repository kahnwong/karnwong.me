+++
title = "Simplify self-hosting backups to S3 with docker"
date = "2024-09-07"
path = "/posts/2024/09/simplify-self-hosting-backups-to-s3-with-docker"

[taxonomies]
categories = ["homelab"]
tags = [ "aws", "docker","postgres"]

+++

These days there are multiple ways to deploy a workload, be it cloud-based or bare-metal. For cloud, depending on whether you are using PaaS or IaaS, backup options can vary.

Why do we need to backup? Because your workloads can contain a state, this can be stored as local files, inside a database, or as other assets outside the application itself.

Take a database for example, ideally you would need a daily backup so you can revert a database to a state before its corruption without losing as much data. Some workloads might store uploaded images, for simplicity let's say they are being written to disk.

For a backup, it can be as simple as running a database dump command, or compress a data directory and save it somewhere. Following the rule of thumb: you should have off-site backups.

In cloud, a managed database would have daily backups as a built-in feature, offloading users from setting up the backup operation themselves.

But in a bare-metal setup (which also applies for self-hosting), you need to somehow store the backup artifacts in, say, AWS S3. And there's a lot of commands involved:

```bash
current_date=$(date +%Y-%m-%d)

backup_bucket="${BUCKET_NAME:-backup}"
backup_prefix="s3://$backup_bucket/$current_date"

filename="$SERVICE_NAME-sqldump-$current_date.bin"
PGPASSWORD="$POSTGRES_PASSWORD" pg_dump -Fc -c -U "$POSTGRES_USERNAME" --host "$POSTGRES_HOSTNAME" >"$filename"

aws s3 cp "${aws_args[@]}" "$filename" "$backup_prefix/$filename"
```

Creating a bash script can work, but this means you have to install required binaries or else the script won't be able to execute. In which after this you can parameterize the inputs.

The final step would involve setting up a cronjob, which is unavoidable, but doing the path logistics is not fun. Additionally, testing a crontab is very clunky, it would involve editing the crontabs to trigger in a minute from now. And when it's triggered, you can't see logs in real-time.

But crontabs still execute commands, why not package this into a docker image, so a backup job would essentially be a `docker run` command? This way it's easier to debug, all dependencies are accounted for, and you can use it anywhere without setting up the environment (you still need docker, but you are unlikely to mess up a docker installation).

So when tying it all together, a Dockerfile would contain following binaries (so we can backup postgres or a path, and push it to s3):

```Dockerfile
FROM nixos/nix:latest

RUN nix-channel --update
RUN nix-env -iA nixpkgs.bash && \
  nix-env -iA nixpkgs.gnutar && \
  nix-env -iA nixpkgs.gzip && \
  nix-env -iA nixpkgs.curl && \
  nix-env -iA nixpkgs.postgresql_16 && \
  nix-env -iA nixpkgs.awscli2

# set entrypoint
WORKDIR /opt/backup
COPY entrypoint.sh .
RUN chmod +x entrypoint.sh

ENTRYPOINT ["bash", "entrypoint.sh"]
```

Notice the entrypoint. Essentially it would take in environment variables, and run the backup commands, then push it to s3. After a successful upload, send a notification.

```bash
#!/bin/bash

# ----------------- VARS ----------------- #
#MODE= # `ARCHIVE`, `DB_POSTGRES`
#SERVICE_NAME=
#BACKUP_PATH=
#BUCKET_NAME=
#BACKUP_PATH_EXCLUDE= # optional

#POSTGRES_USERNAME=
#POSTGRES_PASSWORD=
#POSTGRES_HOSTNAME=

#S3_ENDPOINT= # optional

# either of this
#NTFY_TOPIC_URL
#DISCORD_WEBHOOK_URL

# set filename
current_date=$(date +%Y-%m-%d)

backup_bucket="${BUCKET_NAME:-backup}"
backup_prefix="s3://$backup_bucket/$current_date"

# backup
if [ "$MODE" = "ARCHIVE" ]; then
 filename="$SERVICE_NAME-$current_date.tar.gz"

 # ref: https://stackoverflow.com/a/42985721
 tar_args=()
 if [ -v BACKUP_PATH_EXCLUDE ]; then
  tar_args+=(--exclude "$BACKUP_PATH_EXCLUDE")
 fi
 tar_args+=(
  -czf "$filename" "$BACKUP_PATH"
 )

 tar "${tar_args[@]}"

elif [ "$MODE" = "DB_POSTGRES" ]; then
 filename="$SERVICE_NAME-sqldump-$current_date.bin"
 PGPASSWORD="$POSTGRES_PASSWORD" pg_dump -Fc -c -U "$POSTGRES_USERNAME" --host "$POSTGRES_HOSTNAME" >"$filename"

else
 echo "$MODE is not supported"
fi

# upload
aws_args=()
if [ -v S3_ENDPOINT ]; then
 aws_args+=(--endpoint-url "$S3_ENDPOINT")
fi
aws s3 cp "${aws_args[@]}" "$filename" "$backup_prefix/$filename"

# notify
notify_message="Successfully backup $SERVICE_NAME"
if [ -v NTFY_TOPIC_URL ]; then
 curl -d "$notify_message" "$NTFY_TOPIC_URL"
elif [ -v DISCORD_WEBHOOK_URL ]; then
 curl -i \
  -H "Accept: application/json" \
  -H "Content-Type:application/json" \
  -X POST --data "{\"content\": \"$notify_message\"}" \
  "$DISCORD_WEBHOOK_URL"
else
 echo "Cannot send a notification since no notification backend has been configured."
fi
```

Personally I use Kubernetes for self-hosting, but I have a few services on a small VPS using docker. This setup means I can use the same method to backup docker or kubernetes workloads.
