terraform {
  required_providers {
    eas = {
      source = "fintreal/eas"
    }
  }
}

provider "eas" {
  token        = "EXPO_TOKEN"
  account_name = "EXPO_ACCOUNT_NAME"
}
