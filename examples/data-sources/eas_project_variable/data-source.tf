data "eas_project_variable" "this" {
  project_name = "MyProject"
  name         = "MY_VARIABLE"
  environment  = "development"
}
