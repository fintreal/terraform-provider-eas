data "eas_app_store_api_key" "this" {
  identifier = "..."
}

data "eas_ios_certificate" "this" {
  serial_number = "..."
}

resource "eas_app" "this" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "eas_ios_app_identifier" "this" {
  identifier = "my.app.identifier"
}

resource "eas_provisioning_profile" "this" {
  app_identifier_id = eas_ios_app_identifier.this.id
  base64            = "..."
}

resource "eas_ios_app_credentials" "this" {
  app_id               = eas_app.eas_app.id
  app_identifier_id    = eas_ios_app_identifier.this.id
  app_store_api_key_id = data.eas_app_store_api_key.this.id
  app_store {
    provisioning_profile_id = eas_provisioning_profile.this.id
    certificate_id          = data.eas_ios_certificate.this.id
  }
}
