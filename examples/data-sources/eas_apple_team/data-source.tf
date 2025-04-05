terraform {
  required_providers {
    eas = {
      source  = "fintreal/eas"
      version = "~> 1.4"
    }
  }
}

provider "eas" {}

data "eas_apple_team" "eas_apple_team" {
  identifier = "TEST_APPLE_TEAM_ID"
}

output "team" {
  value = data.eas_apple_team.eas_apple_team
}