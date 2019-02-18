# tf_ccs_test
Custom terraform provider for Cloud Config Server

### Requirements
- golang (https://golang.org/doc/install#install)
- terraform (https://www.terraform.io/downloads.html)

### Testing
From your command line
```
go build -o terraform-provider-oktaccs
terraform init
terraform apply --auto-approve
```
