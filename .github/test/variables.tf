variable "APP_STORE_API_KEY_IDENTIFIER" {
  type = string
}

variable "IOS_CERTIFICATE_SERIAL_NUMBER" {
  type = string
}

variable "IOS_PUSH_KEY_IDENTIFIER" {
  type = string
}

variable "GOOGLE_SERVICE_ACCOUNT_KEY_PROJECT_IDENTIFIER" {
  type = string
}

variable "PROVISIONING_PROFILE_BASE64" {
  type = string
}

variable "FCM_KEY" {
  type = string
}

locals {
  bundle_identifier = "com.example.app.test"
}
