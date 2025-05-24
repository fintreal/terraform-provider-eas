resource "eas_android_app_credentials" "eas_android_app_credentials" {
  app_id                        = "953ed82f-4ac7-47be-ab46-d9c7a1169fe6"
  identifier                    = "com.example.myapp"
  google_service_account_key_id = "36b45ce5-1cf3-4e29-a04d-88fb3c4b5683"
  fcm_key                       = "{\"type\":\"service_account\",\"project_id\":\"terraform-test-app-o9e\",\"private_key_id\":\"123\",\"private_key\":\"my-private-key\",\"client_email\":\"firebase-adminsdk-fbsvc@terraform-test-app-o9e.iam.gserviceaccount.com\",\"client_id\":\"116186624576095421809\",\"auth_uri\":\"https://accounts.google.com/o/oauth2/auth\",\"token_uri\":\"https://oauth2.googleapis.com/token\",\"auth_provider_x509_cert_url\":\"https://www.googleapis.com/oauth2/v1/certs\",\"client_x509_cert_url\":\"https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-fbsvc%40terraform-test-app-o9e.iam.gserviceaccount.com\",\"universe_domain\":\"googleapis.com\"}"
  build_credentials {
    name        = "Default"
    keystore_id = "67484c57-542f-48fc-a470-fa6703a3a6f5"
  }
}
