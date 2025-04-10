terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_apple_certificate" "apple_certificate" {
  serial_number = "3D986E25FF1B48C2417853A07AA15C55"
}

output "name" {
  value = data.eas_apple_certificate.apple_certificate
}
