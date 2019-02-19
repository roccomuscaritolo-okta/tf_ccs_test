provider "oktaccs" {
  base_url = "https://ct0.cloud-config.auw2l.internal"
}

data "oktaccs_config_server_secret" "main" {
  profiles = ["monolith_ct1", "monolith_ct2"]
}

output "db_user" {
  value = "${data.oktaccs_config_server_secret.main.properties.db.user}"
}
