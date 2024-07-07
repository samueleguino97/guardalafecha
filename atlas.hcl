variable "token" {
  type    = string
  default = getenv("TURSO_TOKEN")
}

env "turso" {
  url     = "libsql://savethedate-samueleguino97.turso.io?authToken=${var.token}"
  exclude = ["_litestream*"]
}
