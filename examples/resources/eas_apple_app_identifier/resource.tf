terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

resource "eas_apple_app_identifier" "eas_apple_app_identifier" {
  identifier = "my.app.identifier"
}
