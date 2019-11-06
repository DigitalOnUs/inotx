package config

import (
	"strconv"
)

// avoid the tentation to modify this array values
var (
	_resourceType = [6]string{
		"firewall",
		"load-balancer",
		"service-pool",
		"consul-cluster",
		"consul-server",
		"consul-client",
	}
)

func (i ResourceType) String() string {
	if i < 0 || i > ResourceType(len(_resourceType)-1) {
		return "ResourceType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _resourceType[i]
}
