data "eas_app_store_api_key" "this" {
  identifier = var.APP_STORE_API_KEY_IDENTIFIER
}

data "eas_ios_certificate" "this" {
  serial_number = var.IOS_CERTIFICATE_SERIAL_NUMBER
}

data "eas_ios_push_key" "this" {
  identifier = var.IOS_PUSH_KEY_IDENTIFIER
}

data "eas_google_service_account_key" "this" {
  project_identifier = var.GOOGLE_SERVICE_ACCOUNT_KEY_PROJECT_IDENTIFIER
}
