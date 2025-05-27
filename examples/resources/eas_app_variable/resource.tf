resource "eas_app" "this" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "eas_app_variable" "this" {
  app_id       = eas_app.this.id
  name         = "API_URL"
  value        = "http://example.com/api"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}
