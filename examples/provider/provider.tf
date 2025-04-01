terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}

provider "eas" {}
