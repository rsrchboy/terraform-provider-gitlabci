The `gitlabci_runner` resource allows the trivial creation of a runner token
using a runner registration token, without requiring authentication to the
GitLab instance itself.  It does this by using the [runner registration
token](https://docs.gitlab.com/runner/register/) of the
project/group/instance, rather than the authentication credentials of any
specific user.

Note that, once created, the registration is immutable: any changes will
result in the resource being destroyed and recreated.

See the [Registering a runner](https://registry.terraform.io/providers/rsrchboy/gitlabci/latest/docs/guides/registering-a-runner)
guide for more information.
