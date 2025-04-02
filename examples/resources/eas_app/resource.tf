terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}

resource "eas_app" "eas_app" {
    name = "My App Name"
    slug = "my-app-slug"
}

output "name" {
  value = eas_app.eas_app
}
