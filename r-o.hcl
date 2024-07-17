# readonly-policy.hcl
path "kv/*" {
  capabilities = ["read", "list"]
}

