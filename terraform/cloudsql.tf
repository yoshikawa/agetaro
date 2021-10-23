resource "google_sql_database_instance" "agetaro-sql" {
  name             = "agetaro-sql"
  database_version = var.database_version
  region           = var.gcp_region

  settings {
    tier              = "db-f1-micro"
    availability_type = "REGIONAL"
    backup_configuration {
      binary_log_enabled = true
      enabled=true
    }
  }
}
