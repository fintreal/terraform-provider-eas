terraform {
  required_providers {
    eas = {
      source = "fintreal/eas"
    }
  }
}

provider "eas" { }

resource "eas_app" "this" {
  name = "Terraform Provider EAS"
  slug = "terraform-provider-eas"
}

resource "eas_app_variable" "this" {
  app_id       = eas_app.this.id
  name         = "API_KEY"
  value        = "my-api-key"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}

resource "eas_ios_app_identifier" "this" {
  identifier = local.bundle_identifier
}

resource "eas_ios_app_provisioning_profile" "this" {
  app_identifier_id = eas_ios_app_identifier.this.id
  base64            = var.PROVISIONING_PROFILE_BASE64
}

resource "eas_ios_app_credentials" "this" {
  app_id               = eas_app.this.id
  app_identifier_id    = eas_ios_app_identifier.this.id
  app_store_api_key_id = data.eas_app_store_api_key.this.id
  push_key_id          = data.eas_ios_push_key.this.id
  app_store {
    provisioning_profile_id = eas_ios_app_provisioning_profile.this.id
    certificate_id          = data.eas_ios_certificate.this.id
  }
}

resource "eas_android_app_credentials" "this" {
  app_id                        = eas_app.this.id
  identifier                    = local.bundle_identifier
  google_service_account_key_id = data.eas_google_service_account_key.this.id
  fcm_key                       = var.FCM_KEY
  build_credentials {
    name        = "Default"
    keystore_id = "67484c57-542f-48fc-a470-fa6703a3a6f5"
  }
}
