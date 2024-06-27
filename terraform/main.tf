variable "location" {}
variable "project_id" {
}
variable "machine_type" {
  
}
provider "google" {
  project = var.project_id
  region  = var.location
}

data "google_secret_manager_secret_version" "db_token" {
  secret  = "projects/${var.project_id}/secrets/DB_TOKEN"
  version = "latest"
}

data "google_secret_manager_secret_version" "jwt_secret" {
  secret  = "projects/${var.project_id}/secrets/JWT_SECRET"
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

resource "google_container_cluster" "primary" {
  name     = "primary-cluster"
  location = var.location
  initial_node_count = 1

  deletion_protection = false
  node_config {
    machine_type = var.machine_type
  }
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "primary-node-pool"
  location   = google_container_cluster.primary.location
  cluster    = google_container_cluster.primary.name
  node_count = 1

  node_config {
    disk_size_gb = 10
    preemptible  = true
    machine_type = var.machine_type
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
    ]
  }
}


