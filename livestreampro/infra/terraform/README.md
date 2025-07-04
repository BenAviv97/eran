# Terraform Setup

This configuration provisions the core AWS infrastructure such as the EKS cluster.

## Example `terraform.tfvars`

```hcl
region       = "us-west-2"
cluster_name = "livestream-dev"
subnet_ids   = ["subnet-123", "subnet-456"]
```

Initialize and apply:

```bash
terraform init
terraform apply -var-file="terraform.tfvars"
```
