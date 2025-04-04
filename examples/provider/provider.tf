terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4.0"
    }
  }
}


provider "eas" {
  token       = "EXPO_TOKEN"
  account_name = "EXPO_ACCOUNT_NAME"
}
