
datacenter "dc1" {
  description = "Main datacenter"
  default     = false
}

resource "consul-client" "dc1-client1" {

  association {
    id   = "service.log.log1"
    type = "contains"
  }
  association {
    id   = "service.noty.not1"
    type = "contains"
  }
  association {
    id   = "service.auth.auth1"
    type = "contains"
  }
  association {
    id   = "service.store.store1"
    type = "contains"
  }
  association {
    id   = "service.web.web1"
    type = "contains"
  }
  association {
    id   = "service.payment.pay1"
    type = "contains"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "egress"
  }

  location = "datacenter.dc1"
}
resource "consul-client" "dc1-client2" {

  association {
    id   = "service.log.log2"
    type = "contains"
  }
  association {
    id   = "service.noty.not2"
    type = "contains"
  }
  association {
    id   = "service.auth.auth2"
    type = "contains"
  }
  association {
    id   = "service.store.store2"
    type = "contains"
  }
  association {
    id   = "service.web.web2"
    type = "contains"
  }
  association {
    id   = "service.payment.pay2"
    type = "contains"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "egress"
  }

  location = "datacenter.dc1"
}
resource "consul-client" "dc1-client3" {

  association {
    id   = "service.store.store3"
    type = "contains"
  }
  association {
    id   = "service.web.web3"
    type = "contains"
  }
  association {
    id   = "service.payment.pay3"
    type = "contains"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "egress"
  }

  location = "datacenter.dc1"
}
resource "consul-client" "dc1-client4" {

  association {
    id   = "service.payment.pay4"
    type = "contains"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "egress"
  }

  location = "datacenter.dc1"
}
resource "consul-client" "dc1-client5" {

  association {
    id   = "service.payment.pay5"
    type = "contains"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-cluster.cluster-dc1"
    type = "egress"
  }

  location = "datacenter.dc1"
}
resource "consul-server" "dc1-server1" {

  association {
    id   = "resource.consul-server.dc1-server2"
    type = "egress"
  }
  association {
    id   = "resource.consul-server.dc1-server2"
    type = "ingress"
  }
  association {
    id   = "resource.consul-server.dc1-server3"
    type = "egress"
  }
  association {
    id   = "resource.consul-server.dc1-server3"
    type = "ingress"
  }

  location = "datacenter.dc1"
}
resource "consul-server" "dc1-server2" {

  association {
    id   = "resource.consul-server.dc1-server1"
    type = "egress"
  }
  association {
    id   = "resource.consul-server.dc1-server1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-server.dc1-server3"
    type = "egress"
  }
  association {
    id   = "resource.consul-server.dc1-server3"
    type = "ingress"
  }

  location = "datacenter.dc1"
}
resource "consul-server" "dc1-server3" {

  association {
    id   = "resource.consul-server.dc1-server1"
    type = "egress"
  }
  association {
    id   = "resource.consul-server.dc1-server1"
    type = "ingress"
  }
  association {
    id   = "resource.consul-server.dc1-server2"
    type = "egress"
  }
  association {
    id   = "resource.consul-server.dc1-server2"
    type = "ingress"
  }

  location = "datacenter.dc1"
}
resource "consul-cluster" "cluster-dc1" {

  association {
    id   = "resource.consul-server.dc1-server1"
    type = "contains"
  }
  association {
    id   = "resource.consul-server.dc1-server2"
    type = "contains"
  }
  association {
    id   = "resource.consul-server.dc1-server3"
    type = "contains"
  }

  location = "datacenter.dc1"
}

service "auth" "auth1" {
  port     = 55
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = "1"
    software = "ldap"
  }
}
service "auth" "auth2" {
  port     = 55
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = "3"
    software = "nginx"
  }
}
service "db" "db1" {
  port     = 5432
  address  = ""
  protocol = ""

  meta {
    role     = "primary"
    version  = ""
    software = "postgres"
  }
}
service "db" "db2" {
  port     = 5432
  address  = ""
  protocol = ""

  meta {
    role     = "secondary"
    version  = ""
    software = "postgres"
  }
}
service "log" "log1" {
  port     = 7001
  address  = "localhost"
  protocol = ""

  meta {
    role     = "primary"
    version  = ""
    software = "custom-logger"
  }
}
service "log" "log2" {
  port     = 8001
  address  = "localhost"
  protocol = ""

  meta {
    role     = "secondary"
    version  = ""
    software = "custom-logger"
  }
}
service "noty" "not1" {
  port     = 8000
  address  = "localhost"
  protocol = ""
}
service "noty" "not2" {
  port     = 8000
  address  = "localhost"
  protocol = ""
}
service "payment" "pay1" {
  port     = 8101
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "outlook"
  }
}
service "payment" "pay2" {
  port     = 8101
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "sms"
  }
}
service "payment" "pay3" {
  port     = 8101
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "whatsapp"
  }
}
service "payment" "pay4" {
  port     = 8101
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "outlook"
  }
}
service "payment" "pay5" {
  port     = 8101
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "outlook"
  }
}
service "store" "store1" {
  port     = 3000
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "rails"
  }
}
service "store" "store2" {
  port     = 3000
  address  = "localhost"
  protocol = ""

  meta {
    role     = ""
    version  = ""
    software = "rails"
  }
}
service "store" "store3" {
  port     = 3100
  address  = "localhost"
  protocol = ""

  meta {
    role     = "primary"
    version  = ""
    software = "rails"
  }
}
service "web" "web1" {
  port     = 80
  address  = "localhost"
  protocol = "https"

  meta {
    role     = ""
    version  = "1"
    software = "nginx"
  }
}
service "web" "web2" {
  port     = 80
  address  = "localhost"
  protocol = "https"

  meta {
    role     = ""
    version  = "2"
    software = "nginx"
  }
}
service "web" "web3" {
  port     = 80
  address  = "localhost"
  protocol = "https"
}
