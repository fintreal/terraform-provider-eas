terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}
// EXPO_TOKEN must exist
provider "eas" {}
