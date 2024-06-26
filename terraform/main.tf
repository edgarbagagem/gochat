provider "google" {
  project = "fastuga"
  region  = "us-central1"
}

data "google_secret_manager_secret_version" "db_token" {
  secret  = "projects/fastuga/secrets/DB_TOKEN"
  version = "latest"
}

data "google_secret_manager_secret_version" "jwt_secret" {
  secret  = "projects/fastuga/secrets/JWT_SECRET"
  version = "latest"
}

resource "kubernetes_secret" "api_secrets" {
  metadata {
    name      = "api-secrets"
    namespace = "default"
  }

  data = {
    DB_TOKEN   = data.google_secret_manager_secret_version.db_token.secret_data
    JWT_SECRET = data.google_secret_manager_secret_version.jwt_secret.secret_data
  }
}

# resource "google_container_cluster" "primary" {
#   name     = "primary-cluster"
#   location = "us-central1"
#   initial_node_count = 1

#   node_config {
#     machine_type = "e2-medium"
#   }
# }

# resource "google_container_node_pool" "primary_nodes" {
#   name       = "primary-node-pool"
#   location   = google_container_cluster.primary.location
#   cluster    = google_container_cluster.primary.name
#   node_count = 1

#   node_config {
#     preemptible  = true
#     machine_type = "e2-medium"
#     oauth_scopes = [
#       "https://www.googleapis.com/auth/cloud-platform",
#     ]
#   }
# }

# output "kubeconfig" {
#   value = google_container_cluster.primary.endpoint
# }

# output "client_certificate" {
#   value = base64decode(google_container_cluster.primary.master_auth[0].client_certificate)
# }

# output "client_key" {
#   value = base64decode(google_container_cluster.primary.master_auth[0].client_key)
# }

# output "cluster_ca_certificate" {
#   value = base64decode(google_container_cluster.primary.master_auth[0].cluster_ca_certificate)
# }
