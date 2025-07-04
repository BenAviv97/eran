# /home/${USER}/livestreampro/infra/terraform/main.tf
# /home/${USER}/livestreampro/infra/terraform/main.tf
terraform {
  required_version = ">= 1.5.0"
}

provider "aws" {
  region = var.region
}

module "eks" {
  source       = "./modules/eks"
  cluster_name = var.cluster_name
  subnet_ids   = var.subnet_ids
}
