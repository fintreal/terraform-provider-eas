terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_app_store_api_key" "eas_app_store_api_key" {
  identifier = "349P6U74M8"
}

output "name" {
  value = data.eas_app_store_api_key.eas_app_store_api_key
}
