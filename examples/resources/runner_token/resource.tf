resource "gitlabci_runner_token" "this" {
  registration_token = var.registration_token
  locked             = true
  tags = [
    "jinx",
    "powder",
    "cupcake",
  ]
}

variable "registration_token" {
  type        = string
  description = "Runner registration token"
  sensitive   = true
}

output "token" {
  sensitive = true
  value     = gitlabci_runner_token.this.token
}
