package config

type AssociationType int

const (
	Contains AssociationType = iota
	Egress
	Ingress
)
