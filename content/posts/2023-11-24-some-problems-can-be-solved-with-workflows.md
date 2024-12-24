+++
title = "Some problems can be solved with workflows"
date = "2023-11-24"
path = "/posts/2023/11/some-problems-can-be-solved-with-workflows"

[taxonomies]
categories = ["software-engineering",]
tags = [ "workflow", ]

+++

When we face with engineering problems, it's too easy to fall into the trap thinking it should be solved with a technical solution. Seasoned engineers think differently, because they realize that most of the time, it's "people" or "workflow" problems. Let me provide a few examples.

## Management wants analysts to use Jupyter notebook to reduce time required to create a routine report

Background:

- Most analysts are comfortable using Microsoft Excel to work with data, some can also use SQL, but it's rare for analysts to be familiar with Python.
- Jupyter notebook is an interactive development interface for data works, since users can execute a chunk of code at a time, and render data without requiring re-running the full code.

Problem: every month a two-person analyst team would spend two days stitching up multiple CSV files (can be up to 60) via VLookup for a monthly report. This is because analysts have to look up information for each record, in which they use a template query and manually execute 50 queries with changed parameters.

Initial solution: management wants analysts to use Jupyter notebook to reduce time spent on monthly report.

Issues with proposed solution: analysts are comfortable with SQL, but have no experience with programming languages. Using Jupyter notebook can help reduce operations time, but this would require a huge effort for analysts to learn Python, and the only use case for them would be to run this report monthly.

Solution: optimize SQL query to reduce operations time, this means analysts don't have to learn Python. At most, an engineer would have to come up with a helper script to generate SQL from initial input, and analysts only have to run the generated SQL and get a single file, instead of 50 files.

## Engineer wants product owner to upload image assets to blob storage directly to reduce operations overhead

Problem: a service requires image assets, which is provided by the business / marketing team. A product owner would receive new images and have to inform engineers to update respective images in a CDN bucket.

Initial solution: Create a cloud IAM user with upload-only permission to allow a product owner to upload images to CDN bucket directly.

Issues with proposed solution: Creating a cloud IAM account increases attack surface, and the product owner could mistakenly upload assets to the wrong path. Given the upload-only permission, they can't delete images in the wrong path, this could lead to a confusion down the line as to which images are actually used. In addition, the product owner doesn't work with cloud services on a daily basis, this would require some onboarding before the product owner is comfortable enough with cloud web console.

Solution: this is a tech product, which means the product owner is familiar with some enginering tools. Upon further inquiry, turns out this product owner is familiar with git. This means we can set up a git repository to have a content of CDN bucket, the product owner just have to update the state of git repo to reflect desired state, then make a PR. After a push to master, a CI/CD would kick off to sync the git repo with CDN bucket.
