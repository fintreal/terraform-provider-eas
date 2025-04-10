terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_apple_app_identifier" "apple_app_identifier" {
  identifier = "gjexymraod"
}

output "output" {
  value = data.eas_apple_app_identifier.apple_app_identifier
}
