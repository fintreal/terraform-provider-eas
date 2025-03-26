resource "eas_apple_app_identifier" "eas_apple_app_identifier" {
  identifier = "my.app.identifier"
}

resource "eas_provisioning_profile" "provisioning_profile" {
  app_identifier_id = eas_apple_app_identifier.eas_apple_app_identifier.id
  base64            = "BASE64_ENCODED_PROVISIONING_PROFILE"
}
