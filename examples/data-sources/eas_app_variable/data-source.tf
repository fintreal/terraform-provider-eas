terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}

data "eas_app_variable" "eas_app_variable" {
    id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
    app_id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
}

output "name" {
  value = data.eas_app_variable.eas_app_variable
}
