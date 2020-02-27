#!/bin/bash
sed -i "s/organization-id: .*/organization-id: /" project-factory/core/*project.yaml
sed -i "s/external: \".*\"/external: \"\"/" project-factory/core/*project.yaml
