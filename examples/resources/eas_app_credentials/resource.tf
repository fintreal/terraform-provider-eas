terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

resource "eas_app_credentials" "eas_app_credentials" {
  app_id = "11225d64-bd93-468e-9743-e40f2877a614"
  app_identifier_id="2b29d514-3083-41e0-b19b-da302289d844"
}
