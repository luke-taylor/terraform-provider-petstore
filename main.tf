terraform {
  required_providers {
    petstore = {
      source = "hashicorp.com/edu/petstore"
    }
  }
}

provider "petstore" {}
