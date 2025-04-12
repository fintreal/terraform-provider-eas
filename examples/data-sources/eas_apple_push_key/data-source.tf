terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_apple_push_key" "this" {
  identifier = "M6SPYT2C2L"
}

output "name" {
  value = data.eas_apple_push_key.this
}
