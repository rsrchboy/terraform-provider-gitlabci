terraform {
  required_providers {
    gitlabci = {
      source = "terraform.weyl.io/gitlab/gitlabci"
    }
  }
}

provider "gitlabci" {
  # ...or the API address of your particular instance
  base_url = "https://gitlab.com/api/v4"
}
