terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}

resource "eas_app_variable" "app_variable" {
    app_id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
    name = "API_KEY"
    value  = "my-api-key"
    visibility = "PUBLIC"
    environments = ["DEVELOPMENT"]
}

