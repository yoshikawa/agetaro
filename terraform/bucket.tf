## GCS for terraform.tfstate  ##
resource "google_storage_bucket" "agetaro-tfstate-bucket" {
  name          = "terraform-tfstate-${var.gcp_project}"
  location      = var.gcp_region
  storage_class = "REGIONAL"

  versioning {
    enabled = true
  }

  lifecycle_rule {
    action {
      type = "Delete"
    }
    condition {
      num_newer_versions = 5
    }
  }
}

terraform {
  backend "gcs" {
    bucket = "terraform-tfstate-agetaro-2517d"
  }
}