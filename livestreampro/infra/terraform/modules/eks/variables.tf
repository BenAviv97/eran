# /home/${USER}/livestreampro/infra/terraform/modules/eks/variables.tf
# /home/${USER}/livestreampro/infra/terraform/modules/eks/variables.tf
variable "cluster_name" {
  type = string
}

variable "subnet_ids" {
  type = list(string)
}
