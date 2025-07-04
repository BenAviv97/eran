# /home/${USER}/livestreampro/infra/terraform/modules/eks/main.tf
# /home/${USER}/livestreampro/infra/terraform/modules/eks/main.tf
resource "aws_eks_cluster" "this" {
  name = var.cluster_name
  role_arn = aws_iam_role.eks.arn
  vpc_config {
    subnet_ids = var.subnet_ids
  }
}
