# CloudRun exec service account
resource "google_service_account" "run_invoker" {
  project      = var.gcp_project
  account_id   = "terraform-kun"
  display_name = "Cloud Run Invoker Service Account"
}

# CloudRun Policy
data "google_iam_policy" "invoker" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
      "serviceAccount:${google_service_account.run_invoker.email}"
    ]
  }
}

resource "google_project_iam_member" "cloud_sql_connection" {
  role = "roles/cloudsql.client"
  member = "serviceAccount:${google_service_account.run_invoker.email}"
}

# CloudRun Run Policy
resource "google_cloud_run_service_iam_policy" "run_policy" {
  location    = var.gcp_region
  project     = var.gcp_project
  service     = google_cloud_run_service.default.name
  policy_data = data.google_iam_policy.invoker.policy_data
}
