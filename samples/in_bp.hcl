datacenter "dc1" {
  description = "Main datacenter"
}

resource "firewall" "db" {
  association {
    id = "resource.service-pool.payments"

    type = "ingress"
  }

  association {
    id = "service.db.db1"

    type = "egress"
  }

  association {
    id = "service.db.db2"

    type = "egress"
  }

  location = "datacenter.dc1"
}

resource "load-balancer" "auth" {
  association {
    id = "resource.service-pool.web"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.auth"

    type = "ingress"
  }

  location = "datacenter.dc1"
}

resource "load-balancer" "log" {
  association {
    id = "resource.service-pool.web"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.store"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.auth"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.log"

    type = "egress"
  }

  location = "datacenter.dc1"
}

resource "load-balancer" "notifications" {
  association {
    id = "resource.service-pool.web"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.store"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.auth"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.log"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.notification"

    type = "egress"
  }

  location = "datacenter.dc1"
}

resource "load-balancer" "payments" {
  association {
    id = "resource.service-pool.store"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.payments"

    type = "egress"
  }

  location = "datacenter.dc1"
}

resource "load-balancer" "store" {
  association {
    id = "resource.service-pool.web"

    type = "ingress"
  }

  association {
    id = "resource.service-pool.store"

    type = "egress"
  }
}

resource "load-balancer" "web" {
  association {
    id = "resource.service-pool.web"

    type = "egress"
  }

  location = "datacenter.dc1"
}

resource "service-pool" "auth" {
  association {
    id = "service.auth.auth1"

    type = "contains"
  }

  association {
    id = "service.auth.auth2"

    type = "contains"
  }

  location = "datacenter.dc1"
}

resource "service-pool" "log" {
  association {
    id = "service.log.log1"

    type = "contains"
  }

  association {
    id = "service.log.log2"

    type = "contains"
  }

  location = "datacenter.dc1"
}

resource "service-pool" "notification" {
  association {
    id = "service.noty.not1"

    type = "contains"
  }

  association {
    id = "service.noty.not2"

    type = "contains"
  }

  location = "datacenter.dc1"
}

resource "service-pool" "payment" {
  association {
    id = "service.payment.pay1"

    type = "contains"
  }

  association {
    id = "service.payment.pay2"

    type = "contains"
  }

  association {
    id = "service.payment.pay3"

    type = "contains"
  }

  association {
    id = "service.payment.pay4"

    type = "contains"
  }

  association {
    id = "service.payment.pay5"

    type = "contains"
  }

  location = "datacenter.dc1"
}

resource "service-pool" "store" {
  association {
    id = "service.store.store1"

    type = "contains"
  }

  association {
    id = "service.store.store2"

    type = "contains"
  }

  association {
    id = "service.store.store3"

    type = "contains"
  }

  location = "datacenter.dc1"
}

resource "service-pool" "web" {
  association {
    id = "service.web.web1"

    type = "contains"
  }

  association {
    id = "service.web.web2"

    type = "contains"
  }

  association {
    id = "service.web.web3"

    type = "contains"
  }

  location = "datacenter.dc1"
}

service "auth" "auth1" {
  address = "localhost"

  meta {
    software = "ldap"

    version = 1
  }

  port = 55
}

service "auth" "auth2" {
  address = "localhost"

  meta {
    software = "nginx"

    version = 3
  }

  port = 55
}

service "db" "db1" {
  meta {
    role = "primary"

    software = "postgres"
  }

  port = 5432
}

service "db" "db2" {
  meta {
    role = "secondary"

    software = "postgres"
  }

  port = 5432
}

service "log" "log1" {
  address = "localhost"

  meta {
    role = "primary"

    software = "custom-logger"
  }

  port = 7001
}

service "log" "log2" {
  address = "localhost"

  meta {
    role = "secondary"

    software = "custom-logger"
  }

  port = 8001
}

service "noty" "not1" {
  address = "localhost"

  port = 8000
}

service "noty" "not2" {
  address = "localhost"

  port = 8000
}

service "payment" "pay1" {
  address = "localhost"

  meta {
    software = "outlook"
  }

  port = 8101
}

service "payment" "pay2" {
  address = "localhost"

  meta {
    software = "sms"
  }

  port = 8101
}

service "payment" "pay3" {
  address = "localhost"

  meta {
    software = "whatsapp"
  }

  port = 8101
}

service "payment" "pay4" {
  address = "localhost"

  meta {
    software = "outlook"
  }

  port = 8101
}

service "payment" "pay5" {
  address = "localhost"

  meta {
    software = "outlook"
  }

  port = 8101
}

service "store" "store1" {
  address = "localhost"

  meta {
    software = "rails"
  }

  port = 3000
}

service "store" "store2" {
  address = "localhost"

  meta {
    software = "rails"
  }

  port = 3000
}

service "store" "store3" {
  address = "localhost"

  meta {
    role = "primary"

    software = "rails"
  }

  port = 3100
}

service "web" "web1" {
  address = "localhost"

  meta {
    software = "nginx"

    version = 1
  }

  port = 80

  protocol = "https"
}

service "web" "web2" {
  address = "localhost"

  meta {
    software = "nginx"

    version = 2
  }

  port = 80

  protocol = "https"
}

service "web" "web3" {
  address = "localhost"

  port = 80

  protocol = "https"
}
