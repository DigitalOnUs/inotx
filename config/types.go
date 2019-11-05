package config

type ResourceType int

const (
	Firewall ResourceType = iota
	LoadBalancer
	ServicePool
	ConsulCluster
	ConsulServer
	ConsulClient
)
