resource "eas_ios_app_identifier" "this" {
  identifier = "my.app.identifier"
}

resource "eas_ios_app_provisioning_profile" "this" {
  app_identifier_id = eas_ios_app_identifier.this.id
  base64            = "BASE64_ENCODED_PROVISIONING_PROFILE"
}
