---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "eas_ios_app_credentials Resource - terraform-provider-eas"
subcategory: ""
description: |-
  
---

# eas_ios_app_credentials (Resource)



## Example Usage

```terraform
data "eas_app_store_api_key" "this" {
  identifier = "..."
}

data "eas_ios_certificate" "this" {
  serial_number = "..."
}

resource "eas_app" "this" {
  name = "My App Name"
  slug = "my-app-slug"
}

resource "eas_ios_app_identifier" "this" {
  identifier = "my.app.identifier"
}

resource "eas_provisioning_profile" "this" {
  app_identifier_id = eas_ios_app_identifier.this.id
  base64            = "..."
}

resource "eas_ios_app_credentials" "this" {
  app_id               = eas_app.eas_app.id
  app_identifier_id    = eas_ios_app_identifier.this.id
  app_store_api_key_id = data.eas_app_store_api_key.this.id
  app_store {
    provisioning_profile_id = eas_provisioning_profile.this.id
    certificate_id          = data.eas_ios_certificate.this.id
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `app_id` (String) EAS App Id
- `app_identifier_id` (String) EAS App Identifier Id

### Optional

- `app_store` (Block List, Max: 1) EAS App Build Credentials for App Store (see [below for nested schema](#nestedblock--app_store))
- `app_store_api_key_id` (String) EAS App Store Api Key Id
- `push_key_id` (String) EAS Apple Push Key Id

### Read-Only

- `id` (String) EAS App Credentials Id

<a id="nestedblock--app_store"></a>
### Nested Schema for `app_store`

Required:

- `certificate_id` (String) EAS Id of the Apple Distribution Certificate
- `provisioning_profile_id` (String) EAS Provisioning Profile Id

Read-Only:

- `id` (String) EAS iOS Build Credential Id
