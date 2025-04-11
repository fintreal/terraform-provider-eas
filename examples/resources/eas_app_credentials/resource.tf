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
  app_id = "572945eb-2ac9-4c50-96bb-881fc1cbfc41"
  app_identifier_id="41ea486c-676a-4723-97b8-e0f80c53845f"
  app_store_api_key_id="564e9d75-ff77-4860-92ee-7c0ab2066c82"
  app_store {
    provisioning_profile_id = "72157c17-10db-4851-8633-afd5a08384ce"
    certificate_id = "702635c5-3aa1-477c-83b6-bb66a1644aad"
  }
}
