data "google_container_registry_image" "app" {
  name = "agetaro-kun"
}

resource "google_cloud_run_service" "default" {
  name     = "cloudrun-service"
  location = var.gcp_region

  template {
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = google_sql_database_instance.agetarou-mysql.connection_name
      }
      labels = {
        environment = "dev"
      }
    }
    spec {
      containers {
        image = data.google_container_registry_image.app.image_url

        env {
          name  = "PROJECT_ID"
          value = var.gcp_project
        }
        env {
          name  = "DB_USER"
          value = "agetarou-kun"
        }
        env {
          name  = "DB_PASSWORD"
          value = "tuyoi-kimochi"
        }
        env {
          name  = "DB_NAME"
          value = "agetaro"
        }

        env {
          name = "DB_INSTANCE"
          value = google_sql_database_instance.agetarou-mysql.connection_name
        }
      }
      service_account_name = google_service_account.run_invoker.email
      timeout_seconds      = 600
    }
  }

  traffic {
    percent         = 100
    latest_revision = true
  }
}

locals {
  url = google_cloud_run_service.default.status[0].url
}

output "url" {
  value = local.url
}
