# forgerock-terraform-provider
This project (***in real early stages***) aims to create a custom Terraform provider to manage configurations for the ForgeRock Access Manager - part of the ForgeRock Identity stack.

Feed-backs and inputs are welcome!

Resources would cover managing configurations for the following products:

- Identity Gateway
- Identity Manager
- Access Manager
- Directory Services ? Need to evaluate feasibility of developing it

## ToC
1. [SDK](#sdk)


# SDK

Given each product may have specific ways for authenticate or perform other specific tasks, there is a separate SDK repository where those are implemented. It can be found at [ForgeRock golang SDK](https://github.com/jralmaraz/forgerock-go-sdk)

# Debugging

https://developer.hashicorp.com/terraform/plugin/debugging#visual-studio-code

# Test 

```shell
make install

cd examples && rm .terraform.lock.hcl && terraform init

Add credentials to provider config on examples/main.tf

provider "accessmanager" {
    username = "bla"
    password = "bla"
}

Yet to figure out why environment variables are not working as `schema.EnvDefaultFunc` should take care of it when not set in provider config:

 export XOpenAMUsername="bla"
 export XOpenAMPassword="bla"

terraform plan
 ```
