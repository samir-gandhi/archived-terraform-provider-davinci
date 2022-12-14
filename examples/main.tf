terraform {
  required_providers {
    // named pingonedv for now until merged to actual pingone provider
    dv = {
      version = "0.0.1"
      source  = "samir-gandhi/pingidentity/davinci"
    }
  }
}

provider "dv" {
  username = var.dv_username
  password = var.dv_password
  # company_id = "dcf2011c-d0fc-4b59-81bc-518c8eec3dab"
  base_url = var.dv_url
}

data "dv_connections" "all" {
}

output "dv_connections" {
  value = data.dv_connections.all.connections

}

resource "dv_connection" "annotation" {
  name         = "myAnnotationConnector"
  connector_id = "annotationConnector"
}

resource "dv_connection" "crowd_strike" {
  connector_id = "crowdStrikeConnector"
  name         = "CrowdStrike2"
  properties {
    name  = "clientId"
    value = "1234"
  }
  properties {
    name  = "clientSecret"
    value = "1234567"
  }
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
