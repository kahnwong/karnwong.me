+++
title = "Spark on Kubernetes"
date = "2023-09-12"
path = "/posts/2023/09/spark-on-kubernetes"

[taxonomies]
categories = ["data-engineering",]
tags = [  "spark",  "finops", "devx", "gcp", "aws", "azure", "gke",]

+++

## Background

For data processing tasks, there are different ways you can go about it:

- using SQL to leverage a database engine to perform data transformation
- dataframe-based frameworks such as pandas, ray, dask, polars
- big data processing frameworks such as spark

Check out [this article](/posts/2023/04/duckdb-vs-polars-vs-spark) for more info on polars vs spark benchmark.

## The problem

At larger data scale, other solutions (except spark) can work, but with a lot of vertical scaling, and this can get very expensive. For a comparison, our team had to scale a database to 4/16 GB and it still took the whole night, whereas spark on a single node can process the data in 2 minutes flat.

Using spark would solve a lot of scaling problems, the only problem is that spark in cluster mode is very hard to set up. That's why most if not all cloud providers provide managed spark runtime (AWS EMR, GCP Dataproc, Azure Databricks). And since these are managed spark, this means you can't test it locally (and you'll have to wait at least 5 minutes for the cluster to start and finish bootstrapping.)

## The solution

As of Apache Spark 3.1 release in March 2021, spark on kubernetes is generally available. This is great in so many ways, including but not limited to:

- you have full control of your infra
- takes less than 10 seconds to start a spark cluster
- can store logs in a central location, to be viewed later via spark history server
- can use minio as local storage backend (better throughput compared to calling S3 via home/work internet)
- cheaper than all managed solutions, even serverless variants (more on this later)

## Demo

### Start a local kubernetes cluster

Minikube, Kind or K3d should work.

```bash
k3d cluster create spark-cluster --api-port 0.0.0.0:6443

# verify that a cluster is created
kubectl get nodes
```

### Create a namespace for spark jobs

```bash
kubectl create namespace spark
```

### Create a service account for spark

```bash
kubectl create serviceaccount spark --namespace spark
kubectl create clusterrolebinding spark-role --clusterrole=edit --serviceaccount=spark:spark --namespace=spark
```

### Install local spark (so you can run `spark-submit`)

```bash
brew install temurin
brew install apache-spark
```

### Run a demo spark job

If encounter `To use support for EC Keys` error: <https://stackoverflow.com/questions/75796747/spark-submit-error-to-use-support-for-ec-keys-you-must-explicitly-add-this-depe>

```bash
# get kubernetes control plane url via `kubectl cluster-info`

spark-submit \
--master k8s://https://0.0.0.0:6443 \
--deploy-mode cluster \
--name spark-pi \
--conf spark.kubernetes.namespace=spark \
--conf spark.kubernetes.authenticate.driver.serviceAccountName=spark \
--conf spark.executor.instances=3 \
--conf spark.kubernetes.container.image=spark:3.4.1 \
local:///opt/spark/examples/src/main/python/pi.py
```

### Tearing down k3d

```bash
k3d cluster delete spark-cluster
```

## Full Demo

![](images/demo.gif)

## Regarding pricing

Since spark-submit on kubernetes is akin to serverless spark, we will be comparing costs with serverless spark runtime.

**Setup**:

- CPU: 8
- Memory: 16
- Runtime: 15 Hours
- Ephemeral storage: 20 GB
- Region: Singapore

Remarks: spark on kubernetes in this article assumes you're using GKE Autopilot, which has a management fee of $60/month (unless it's your only cluster). Compute is spot.

| Serverless Runtime          | Reference                                                                   | Total Cost | Break even point if you run at least 33x workloads |
| --------------------------- | --------------------------------------------------------------------------- | ---------- | -------------------------------------------------- |
| AWS EMR Serverless          | <https://aws.amazon.com/emr/pricing/>                                       | 9.61 USD   | 317.22 USD                                         |
| GCP Dataproc Serverless     | <https://cloud.google.com/products/calculator>                              | 4.26 USD   | 140.58 USD                                         |
| Azure Databricks Serverless | <https://www.databricks.com/product/pricing/product-pricing/instance-types> | 5.61 USD   | 185.13 USD                                         |
| (Bonus) spark on kubernetes | <https://cloud.google.com/kubernetes-engine/pricing>                        | 2.417 USD  | 139.76 USD (Including cluster management fee)      |

The script I used to find a break event point:

```python
aws = 9.61272
gcp = 4.26
databricks = 5.61
k8s = 2.417

found_break_even_point = False
multiplier = 0

while not found_break_even_point:
    multiplier += 1

    print(multiplier)

    aws_multiplied = aws * multiplier
    gcp_multiplied = gcp * multiplier
    databricks_multiplied = databricks * multiplier
    k8s_reference = (k8s * multiplier) + 60  # add gke autopilot management fee

    if k8s_reference < min([aws_multiplied, gcp_multiplied, databricks_multiplied]):
        found_break_even_point = True
```

### Why GKE Autopilot

Traditionally, when you're creating a kubernetes cluster, you'll have to attach nodes yourself. This means, if you're running spark jobs where the compute requirement exceeds the available capacity, your jobs won't be able to start. By using GKE Autopilot, GCP automatically provision nodes and attach to your GKE cluster automatically, which means you won't ever face "insufficient cpu / memory" errors.

## Closing

If you're already using kubernetes, and also use spark for data processing, migrating workloads to k8s can reduce a significant amount of cost. For our case, existing spark jobs are run on a 4vCPU/16GB VM (`109.79 USD / Month`), and this cause a lot of dent in our bills. But if we migrate to spark on k8s, it would only cost us around`16.41 USD / month` 😱. So that's a whopping 85% price reduction 🤯.

For more advanced use cases, check out <https://github.com/kahnwong/spark-on-k8s>.
