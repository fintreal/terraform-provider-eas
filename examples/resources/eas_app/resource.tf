terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
    }
  }
}

resource "eas_app" "app" {
  name = "My App Name"
  slug = "my-app-slug"
}
