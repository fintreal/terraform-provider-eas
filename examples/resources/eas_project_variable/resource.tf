resource "eas_project_variable" "plaintext" {
  project_name = "MyProjectName"
  name         = "API_URL"
  value        = "http://localhost:3000/"
  environment  = "development"
  visibility   = "plaintext"
}

