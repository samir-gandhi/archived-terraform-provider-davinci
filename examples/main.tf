terraform {
  required_providers {
    davinci = {
      version = "0.0.1"
      source  = "samir-gandhi/pingidentity/davinci"
    }
  }
}

provider "davinci" {
  username = var.dv_username
  password = var.dv_password
}

data "davinci_connnnection" "one" {
  connection_id = "ca1b3b9e1e389c8ce534a872678ffc6a"
}
# data "davinci_customers" "customers" {}

# output "customers" {
#   value = data.davinci_customers.customers

# }

# module "tdf" {
#   source = "./customers"
#   customer_name = "tempdvflows"
# }

# output "psl" {
#   value = module.tdf.customer
# }
