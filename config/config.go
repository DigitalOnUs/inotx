package config

import "github.com/hashicorp/hcl/v2"

//Root document and elements
type Root struct {
	Datacenters []*Datacenter `hcl:"datacenter,block"`
	Resources   []*Resource   `hcl:"resource,block"`
	Services    []*Service    `hcl:"service,block"`
}

//Datacenter entity
type Datacenter struct {
	Name        string `hcl:"name,label"`
	Description string `hcl:"description"`
	Default     bool   `hcl:"default,optional"`
}

//Resource (load balancers)
type Resource struct {
	Type         string        `hcl:"type,label"`
	Name         string        `hcl:"name,label"`
	Associations []Association `hcl:"association,block"`
	Location     string        `hcl:"location,optional"`
}

//Service (representation of local stuff)
type Service struct {
	Type     string `hcl:"type,label"`
	Name     string `hcl:"name,label"`
	Port     int    `hcl:"port,optional"`
	Address  string `hcl:"address,optional"`
	Protocol string `hcl:"protocol,optional"`
	Meta     []Meta `hcl:"meta,block"`
}

//Association for any item
type Association struct {
	ID   string `hcl:"id"`
	Type string `hcl:"type"`
}

//Meta information
type Meta struct {
	Role     string   `hcl:"role,optional"`
	Version  string   `hcl:"version,optional"`
	Software string   `hcl:"software,optional"`
	Extra    hcl.Body `hcl:",remain"`
}
