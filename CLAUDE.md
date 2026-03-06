# terraform-provider-eas

Terraform provider for managing Expo Application Services (EAS) apps, credentials, and environment variables via the Expo GraphQL API.

- Repo: `git@github.com:fintreal/terraform-provider-eas.git`
- Default branch: `main`
- Registry namespace: `fintreal/eas`

## Tech Stack

- Go 1.24
- Terraform Plugin SDK v2
- `github.com/fintreal/eas-sdk-go` (private GraphQL client)
- GoReleaser for builds/releases
- GPG-signed releases to Terraform Registry

## Project Structure

```
main.go                          # Provider entrypoint
provider/provider.go             # Provider definition (schema, resources, data sources)
provider/<resource>/schema.go    # Resource/data source schema
provider/<resource>/operations/  # CRUD operations (create.go, read.go, update.go, delete.go)
internal/client/eas.go           # EAS API client wrapper
examples/                        # Terraform example configs (used by tfplugindocs)
docs/                            # Auto-generated registry docs (do not edit manually)
.github/test/                    # Integration test Terraform configs
```

### Resources

- `eas_app` - Expo app
- `eas_app_variable` - App environment variable
- `eas_android_app_credentials` - Android app credentials
- `eas_ios_app_credentials` - iOS app credentials
- `eas_ios_app_identifier` - iOS app identifier (Bundle ID)
- `eas_ios_app_provisioning_profile` - iOS provisioning profile

### Data Sources

- `eas_app_store_api_key`, `eas_ios_certificate`, `eas_ios_push_key`, `eas_google_service_account_key`

## Key Commands

```bash
go build -o terraform-provider-eas   # Build the provider
go install .                         # Install locally for dev testing
go mod tidy                          # Tidy dependencies
```

Set `GOPRIVATE=github.com/fintreal/*` for private module access.

## CI/CD

- **PR / nightly**: Runs `terraform apply` + `terraform destroy` against `.github/test/` configs (test.yml)
- **Push to main**: Auto-generates docs via `tfplugindocs`, creates semver release, builds with GoReleaser (release.yml)
- Docs in `docs/` are auto-generated -- do not edit manually

## Auth

Provider requires `EXPO_TOKEN` and `EXPO_ACCOUNT_NAME` environment variables (or inline config).

## Conventions

- Each resource/data source lives in its own package under `provider/`
- iOS resources are nested under `provider/ios/`, Android under `provider/android/`
- CRUD operations are split into separate files in an `operations/` sub-package
- Schema definition is in `schema.go` at the resource package root
- Resources return `*schema.Resource` via a `Resource()` function; data sources via `DataSource()`
- Commit messages use conventional commits style (e.g., `feat:`, `fix:`, `chore:`)
