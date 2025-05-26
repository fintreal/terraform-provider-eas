resource "eas_app" "eas_app" {
  name = "My App Name"
  slug = "my-app-slug"
}

data "eas_google_service_account_key" "this" {
  project_identifier = "..."
}

resource "eas_android_app_credentials" "this" {
  app_id                        = eas_app.this.id
  identifier                    = "com.example.myapp"
  google_service_account_key_id = "..."
  fcm_key                       = "..."
  build_credentials {
    name        = "Default"
    keystore_id = "..."
  }
}
