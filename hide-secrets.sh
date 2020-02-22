#!/bin/bash
yq -iy '.spec.billingAccountRef.external = ""' project-factory/core/*project.yaml
yq -iy '.metadata.annotations."cnrm.cloud.google.com/organization-id" = ""' project-factory/core/*project.yaml
