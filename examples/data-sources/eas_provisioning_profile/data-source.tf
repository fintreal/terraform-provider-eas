terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_provisioning_profile" "provisioning_profile" {
  id = "8690db1b-c475-43d0-aa3f-67e103c96426"
}

output "eas_provisioning_profile" {
  value = data.eas_provisioning_profile.provisioning_profile
  sensitive = true
}