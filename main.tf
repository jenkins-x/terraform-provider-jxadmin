provider "jx-admin" {
  name = "my-jx-cluster"
  endpoint = "https://127.0.0.1"
  certificate = "some generated cert..."
}

resource "jx_admin_operator" "foo" {
  pipelineUser = "bot1"
}