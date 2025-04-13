terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_google_service_account_key" "this" {
  project_identifier = "playstore-release-5nszcyhabt"
}

output "out" {
  value = data.eas_google_service_account_key.this
}