terraform {
  backend "pg" {
    workspaces {
      name = "go-apis"
    }
  }
}

provider "kubernetes" {
}

module "go_api_deployment_auth" {
  source          = "github.com/blockchain-abstraction-middleware/deployment/modules/deployment"
  namespace       = "go-apis"
  deployment_name = "auth"
  docker_image    = "bamdockerhub/auth"
  config_file     = "review.yml"
  config_path     = "/config/"
}