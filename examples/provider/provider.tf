terraform {
  required_providers {
    gitlabci = {
      source = "registry.terraform.io/rsrchboy/gitlabci"
    }
  }
}

provider "gitlabci" {
  # ...or the API address of your particular instance
  base_url = "https://gitlab.com/api/v4"
}
