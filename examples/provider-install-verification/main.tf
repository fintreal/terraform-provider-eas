terraform {
    required_providers {
      eas = {
        source = "registry.terraform.io/fintreal/expo-eas"
      }
    }
}


data "eas_project_variable" "this" {
  name = "ID3"
  project_name = "TestProject2"
  environment = "development"
}

output "this" {
  value = data.eas_project_variable.this
}
