terraform {
  required_providers {
    davinci = {
      version = "0.0.1"
      source  = "samir-gandhi/pingidentity/davinci"
    }
  }
}

provider "davinci" {
  username   = var.dv_username
  password   = var.dv_password
  company_id = "dcf2011c-d0fc-4b59-81bc-518c8eec3dab"
}

data "davinci_connections" "all" {
}

output "connections" {
  value = data.davinci_connections.all

}

# output "connection_one" {
#   value = data.davinci_connection
# }
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
