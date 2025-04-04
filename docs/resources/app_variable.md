---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "eas_app_variable Resource - terraform-provider-eas"
subcategory: ""
description: |-
  
---

# eas_app_variable (Resource)



## Example Usage

```terraform
resource "eas_app_variable" "app_variable" {
  app_id       = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
  name         = "API_KEY"
  value        = "my-api-key"
  visibility   = "PUBLIC"
  environments = ["DEVELOPMENT"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) The id of the app for the environment variable
- `environments` (Set of String) The environments of the app for the environment variable
- `name` (String) The name of the app for the environment variable
- `value` (String) The value of the app for the environment variable
- `visibility` (String) The visibility of the app for the environment variable

### Read-Only

- `id` (String) The id of the environment variable
