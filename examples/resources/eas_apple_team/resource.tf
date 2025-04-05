terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

resource "eas_apple_team" "eas_apple_team" {
  identifier = "TEST_APPLE_TEAM_ID"
  type       = "COMPANY_OR_ORGANIZATION"
  name       = "Test Team"
}
