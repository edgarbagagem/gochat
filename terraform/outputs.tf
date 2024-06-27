output "kubeconfig" {
  value = google_container_cluster.primary.endpoint
}

output "client_certificate" {
  value = base64decode(google_container_cluster.primary.master_auth[0].client_certificate)
}

output "client_key" {
  value = base64decode(google_container_cluster.primary.master_auth[0].client_key)
  sensitive = true
}

output "cluster_ca_certificate" {
  value = base64decode(google_container_cluster.primary.master_auth[0].cluster_ca_certificate)
}