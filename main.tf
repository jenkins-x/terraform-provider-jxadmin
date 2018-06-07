provider "jx" {

}

resource "jx_team" "team1" {
  name = "team1"
}

resource "jx_team" "team2" {
  name = "team2"
}

resource "jx_environment" "staging" {
  name = "staging"
  promotion_strategy = "auto"
  order = "100"
  namespace = "jx-staging"
}

resource "jx_environment" "production" {
  name = "production"
  promotion_strategy = "manual"
  order = "200"
  namespace = "jx-production"
}
