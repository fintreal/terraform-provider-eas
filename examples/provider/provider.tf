terraform {
  required_providers {
    eas = {
      source = "fintreal/eas"
    }
  }
}

provider "eas" {
  token        = "..."
  account_name = "..."
}
