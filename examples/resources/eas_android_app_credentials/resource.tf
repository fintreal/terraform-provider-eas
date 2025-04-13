terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
    }
  }
}

provider "eas" {}

resource "eas_android_app_credentials" "this" {
  app_id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
  identifier = "com.example.myapp"
  google_service_account_key_id = "36b45ce5-1cf3-4e29-a04d-88fb3c4b5683"
}