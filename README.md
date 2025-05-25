## Expo Application Services Terraform Provider

- Manage Expo EAS app, credentials and environment variables with `terraform`
- Uses Expo EAS GraphQL API

### Authentication

The provider requires an Expo access token and account name for authentication:

1. **Access Token**:
   - Log into your Expo account at https://expo.dev
   - Go to your account settings
   - Create a new access token under the "Access Tokens" section

2. **Account Name**:
   - This is your Expo account username (not email)
   - You can find it in your Expo account settings or in the URL when logged into expo.dev (e.g., https://expo.dev/your-account-name)
   - For organizations, use the organization's account name
   - Example: If your Expo account URL is `https://expo.dev/acme-corp`, then your account name is `acme-corp`

### Documentation

For detailed documentation about available resources and data sources, please refer to the [examples directory](./examples):

- [Resource Examples](./examples/resources) - Examples for managing EAS apps and variables
- [Data Source Examples](./examples/data-sources) - Examples for querying EAS apps and variables
- [Provider Configuration Examples](./examples/provider) - Additional provider configuration examples
