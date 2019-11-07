package config

import (
	"encoding/json"

	"github.com/hashicorp/hcl/v2"
)

//Root document and elements
type Root struct {
	Datacenters []*Datacenter `hcl:"datacenter,block"`
	Resources   []*Resource   `hcl:"resource,block"`
	Services    []*Service    `hcl:"service,block"`
}

//RootJSON : This should be the format supported
type RootJSON struct {
	Datacenter map[string]*Datacenter          `json:"datacenter"`
	Resources  map[string]map[string]*Resource `json:"resource"`
	Services   map[string]map[string]*Service  `json:"service"`
}

//MarshalJSON to be compatible with the reader
func (r *Root) MarshalJSON() ([]byte, error) {
	// creating the structure to marshal
	hierarchy := RootJSON{
		make(map[string]*Datacenter, len(r.Datacenters)),
		make(map[string]map[string]*Resource, 1),
		make(map[string]map[string]*Service, 1),
	}

	for _, dc := range r.Datacenters {
		hierarchy.Datacenter[dc.Name] = dc
	}

	for _, res := range r.Resources {
		// in case not available
		if _, ok := hierarchy.Resources[res.Type]; !ok {
			hierarchy.Resources[res.Type] = make(map[string]*Resource)
		}
		hierarchy.Resources[res.Type][res.Name] = res
	}

	for _, srv := range r.Services {
		if _, ok := hierarchy.Services[srv.Type]; !ok {
			hierarchy.Services[srv.Type] = make(map[string]*Service)
		}
		hierarchy.Services[srv.Type][srv.Name] = srv
	}

	return json.Marshal(&hierarchy)
}

//Datacenter entity
type Datacenter struct {
	Name        string `hcl:"name,label" json:"-"`
	Description string `hcl:"description" json:"description"`
	Default     bool   `hcl:"default,optional" json:"default"`
}

//FQDN - Name
func (dc *Datacenter) FQDN() string {
	return "datacenter." + dc.Name
}

//Resource (load balancers)
type Resource struct {
	Type         string        `hcl:"type,label" json:"-"`
	Name         string        `hcl:"name,label" json:"-"`
	Associations []Association `hcl:"association,block"`
	Location     string        `hcl:"location,optional" json:"location"`
}

func (r *Resource) FQDN() string {
	return "resource." + r.Type + "." + r.Name
}

//Service (representation of local stuff)
type Service struct {
	Type     string `hcl:"type,label" json:"-"`
	Name     string `hcl:"name,label" json:"-"`
	Port     int    `hcl:"port,optional" json:"port,omitempty"`
	Address  string `hcl:"address,optional" json:"address,omitempty"`
	Protocol string `hcl:"protocol,optional" json:"protocol,omitempty"`
	Meta     []Meta `hcl:"meta,block" json:"meta,omitempty"`
}

type AliasService Service

type ServiceJSON struct {
	Meta interface{} `json"meta,omitempty"`
	*AliasService
}

func (s *Service) MarshalJSON() ([]byte, error) {
	// special case for the meta in case they start adding more and more
	var meta interface{}
	// json object by now
	if len(s.Meta) == 1 {
		meta = s.Meta[0]
	}
	// list of
	if len(s.Meta) > 1 {
		meta = s.Meta
	}

	return json.Marshal(&ServiceJSON{
		meta,
		(*AliasService)(s),
	})
}

//Association for any item
type Association struct {
	ID   string `hcl:"id" json:"id"`
	Type string `hcl:"type" json:"type"`
}

//Meta information
type Meta struct {
	Role     string   `hcl:"role,optional" json:"role,omitempty"`
	Version  string   `hcl:"version,optional" json:"version,omitempty"`
	Software string   `hcl:"software,optional" json:"software,omitempty"`
	Extra    hcl.Body `hcl:",remain" json:"extra,omitempty"`
}
