data "eas_app_store_api_key" "eas_app_store_api_key" {
  identifier = "APP_STORE_API_KEY_IDENTIFIER"
}

data "eas_apple_certificate" "eas_apple_certificate" {
  serial_number = "APPLE_CERTIFICATE_SERIAL_NUMBER"
}

resource "eas_app" "eas_app" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "eas_apple_app_identifier" "eas_apple_app_identifier" {
  identifier = "my.app.identifier"
}

resource "eas_provisioning_profile" "provisioning_profile" {
  app_identifier_id = eas_apple_app_identifier.eas_apple_app_identifier.id
  base64            = "BASE64_ENCODED_PROVISIONING_PROFILE"
}

resource "eas_apple_app_credentials" "eas_apple_app_credentials" {
  app_id               = eas_app.eas_app.id
  app_identifier_id    = eas_apple_app_identifier.eas_apple_app_identifier.id
  app_store_api_key_id = data.eas_app_store_api_key.eas_app_store_api_key.id
  app_store {
    provisioning_profile_id = eas_provisioning_profile.provisioning_profile.id
    certificate_id          = data.eas_apple_certificate.eas_apple_certificate.id
  }
}
