terraform {
  required_providers {
    petstore = {
      source = "hashicorp.com/edu/petstore"
    }
  }
}

provider "petstore" {}

resource "petstore_pet" "clifford" {
  id   = 12345
  name = "Clifford the Big Red Dog"
  photo_urls = [
    "https://example.com/pet.jpg"
  ]
  status = "available"
  category = {
    id   = 1
    name = "dog"
  }
  tags = [
    { id = 1, name = "red" },
    { id = 2, name = "vizsla" }
  ]
}

