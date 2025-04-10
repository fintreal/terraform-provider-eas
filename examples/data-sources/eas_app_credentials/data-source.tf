terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_app_credentials" "this" {
  id = "11f2b3f8-ddad-4626-8984-2b96efb28d3c"
  app_id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
}

output "out" {
  value = data.eas_app_credentials.this
}