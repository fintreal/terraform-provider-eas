resource "eas_app" "eas_app" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "eas_app_variable" "eas_app_variable" {
  app_id       = eas_app.eas_app.id
  name         = "API_KEY"
  value        = "my-api-key"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}
