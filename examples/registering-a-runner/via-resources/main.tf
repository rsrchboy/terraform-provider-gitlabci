resource "gitlab_project" "this" {
  name            = "arcane"
  description     = "No magic, plz!"
  visibility_leve = "public"
}

resource "gitlabci_runner_token" "this" {
  registration_token = gitlab_project.this.runners_token
  locked             = true
  tags = [
    "jinx",
    "powder",
    "cupcake",
  ]
}

output "token" {
  sensitive = true
  value     = gitlabci_runner_token.this.token
}
