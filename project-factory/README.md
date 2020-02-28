# Shared VPC

To use the shared-vpc-attachment, two projects must be created and exposed as
resources to kubernetes. One should be the host and the other should be the
service. The attachment should be applied to the host's namespace.

# Service Account Grant to Group

In order for service-account-grant-to-group to be applied, the service account
must have the Service Account Admin role at the organization level, otherwise there will be an error:
Error 403: Permission iam.serviceAccounts.getIamPolicy is required to perform this operation on service account
