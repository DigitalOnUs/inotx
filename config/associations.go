//go:generate stringer -type=AssociationType
package config

type AssociationType int

const (
	Contains AssociationType = iota
	Egress
	Ingress
)
