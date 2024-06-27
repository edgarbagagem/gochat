module "gochat-workload-identity" {
  source              = "terraform-google-modules/kubernetes-engine/google//modules/workload-identity"
  name                = "gochat-service-account"
  namespace           = "default"
  project_id          = var.project_id
  roles               = ["roles/storage.admin", "roles/compute.admin"]
}