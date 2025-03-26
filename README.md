## Expo Application Services Terraform Provider

- Manage Expo EAS app and environment variables with `terraform`
- Uses Expo EAS GraphQL API

### Provider
```hcl
terraform {
  required_providers {
    eas = {
      source = "registry.terraform.io/fintreal/eas"
    }
  }
}

provider "eas" {
    token       = "EXPO_TOKEN" # you can set this via environment variable
    accountName = "EXPO_ACCOUNT_NAME" # you can set this via environment variable
}
```

### Resources

##### app
```hcl
resource "eas_app" "eas_app" {
    name = "My App Name"
    slug = "my-app-slug"
}
```

##### app_variable
```hcl
resource "eas_app_variable" "app_variable" {
    app_id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
    name = "API_KEY"
    value  = "my-api-key"
    visibility = "PUBLIC"
    environments = ["DEVELOPMENT"]
}
```

### Data Sources

##### app
```hcl
data "eas_app" "eas_app" {
    id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
}
```

##### app_variable
```hcl
data "eas_app_variable" "eas_app_variable" {
  name = "API_URL"
  app_id = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
}
```
