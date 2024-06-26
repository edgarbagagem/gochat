provider "kubernetes" {
  config_path = "~/.kube/config"
  # host                   = google_container_cluster.primary.endpoint
  # client_certificate     = base64decode(google_container_cluster.primary.master_auth[0].client_certificate)
  # client_key             = base64decode(google_container_cluster.primary.master_auth[0].client_key)
  # cluster_ca_certificate = base64decode(google_container_cluster.primary.master_auth[0].cluster_ca_certificate)
}

resource "kubernetes_deployment" "gochat-frontend" {
  metadata {
    name = "gochat-frontend-deployment"
    namespace = "default"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "gochat-frontend"
      }
    }
    template {
      metadata {
        labels = {
          app = "gochat-frontend"
        }
      }
      spec {
        container {
          name  = "frontend"
          image = "eddygarr/gochat-frontend:latest"
          port {
            container_port = 80
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "gochat-frontend" {
  metadata {
    name = "gochat-frontend-service"
    namespace = "default"
  }
  spec {
    selector = {
      app = "gochat-frontend"
    }
    port {
      port        = 80
      target_port = 80
    }
    type = "ClusterIP"
  }
}

resource "kubernetes_deployment" "gochat-api" {
  metadata {
    name = "gochat-api-deployment"
    namespace = "default"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "gochat-api"
      }
    }
    template {
      metadata {
        labels = {
          app = "gochat-api"
        }
      }
      spec {
        container {
          name  = "gochat-api"
          image = "eddygarr/gochat-api:latest"
          port {
            container_port = 80
          }
          env {
            name = "DB_NAME"
            value = "gochat-edgarbagagem"
          }
                    env {
            name = "DB_TOKEN"
            value_from {
              secret_key_ref {
                name = "api-secrets"
                key  = "DB_TOKEN"
              }
            }
          }
          env {
            name = "JWT_SECRET"
            value_from {
              secret_key_ref {
                name = "api-secrets"
                key  = "JWT_SECRET"
              }
            }
          }
          env {
            name  = "FRONTEND_URL"
            value = "http://gochat-frontend-service"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "gochat-api" {
  metadata {
    name = "gochat-api-service"
    namespace = "default"
  }
  spec {
    selector = {
      app = "gochat-api"
    }
    port {
      port        = 80
      target_port = 8080
    }
    type = "ClusterIP"
  }
}

resource "kubernetes_deployment" "gochat-ws" {
  metadata {
    name = "gochat-ws-deployment"
    namespace = "default"
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "gochat-ws"
      }
    }
    template {
      metadata {
        labels = {
          app = "gochat-ws"
        }
      }
      spec {
        container {
          name  = "gochat-ws"
          image = "eddygarr/gochat-websockets:latest"
          port {
            container_port = 3000
          }
          env {
            name  = "CORS_ORIGIN"
            value = "http://gochat-frontend-service"
          }
        }
      }
    }
  }
}

resource "kubernetes_service" "gochat-ws" {
  metadata {
    name = "gochat-ws-service"
    namespace = "default"
  }
  spec {
    selector = {
      app = "gochat-ws"
    }
    port {
      port        = 3000
      target_port = 3000
    }
    type = "ClusterIP"
  }
}

resource "kubernetes_cron_job_v1" "joke-cronjob" {
  metadata {
    name = "joke-cronjob"
    namespace = "default"
  }
  spec {
    schedule = "* * * * *" # Every minute
    job_template {
        metadata {
          name = "joke-cronjob"
        }
      spec {
        template {
          metadata {
            name = "joke-cronjob"
          }
          spec {
            container {
              name  = "joke-cronjob"
              image = "eddygarr/joke-cronjob:latest"
              env {
                name  = "WEBSOCKET_SERVER"
                value = "http://gochat-ws-service:3000"
              }
            }
            restart_policy = "OnFailure"
          }
        }
      }
    }
  }
}

resource "kubernetes_ingress_v1" "gochat-ingress" {
  metadata {
    name = "gochat-ingress"
    namespace = "default"
  }
  spec {
    rule {
      http {
        path {
          path = "/"
          backend {
            service {
            name = kubernetes_service.gochat-frontend.metadata[0].name
            port { 
                number =  "80"
              }
          }
        }
        }
        path {
          path = "/api"
          backend {
            service {
            name = kubernetes_service.gochat-api.metadata[0].name
            port { 
                number =  "80"
              }
            }
          }
        }
        path {
          path = "/socket.io"
          backend {
            service {
              name = kubernetes_service.gochat-ws.metadata[0].name
              port { 
                number =  "3000"
                }
            }
          }
        }
      }
    }
  }
  }


