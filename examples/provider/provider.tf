terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/expo-eas"
    }
  }
}

provider "eas" {
  token       = "EXPO_TOKEN" # you can set this via environment variable
  accountName = "EXPO_ACCOUNT_NAME" # you can set this via environment variable
}
