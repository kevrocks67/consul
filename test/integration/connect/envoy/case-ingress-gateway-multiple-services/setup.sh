#!/bin/bash

set -euo pipefail

upsert_config_entry primary '
kind = "proxy-defaults"
name = "global"
config {
  protocol = "http"
}
'

upsert_config_entry primary '
kind = "ingress-gateway"
name = "ingress-gateway"
Defaults {
  MaxConnections        = 10
  MaxPendingRequests    = 20
  MaxConcurrentRequests = 30
}
listeners = [
  {
    port     = 9999
    protocol = "http"
    services = [
      {
        name = "*"
      }
    ]
  },
  {
    port     = 9998
    protocol = "http"
    services = [
      {
        name                  = "s1"
        hosts                 = ["test.example.com"]
        MaxConnections        = 100
        MaxPendingRequests    = 200
        MaxConcurrentRequests = 300
      }
    ]
  }
]
'

register_services primary

gen_envoy_bootstrap ingress-gateway 20000 primary true
gen_envoy_bootstrap s1 19000
gen_envoy_bootstrap s2 19001
