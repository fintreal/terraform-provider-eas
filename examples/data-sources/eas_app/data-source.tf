terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}

data "eas_app" "eas_app" {
    id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
}

output "name" {
  value = data.eas_app.eas_app
}
