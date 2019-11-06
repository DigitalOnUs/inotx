package config

import "sort"

type ResourceType int

const (
	Firewall ResourceType = iota
	LoadBalancer
	ServicePool
	ConsulCluster
	ConsulServer
	ConsulClient
)

type X struct{}
type metadata struct {
	Name         string
	InstancesIDs map[string]X
	Location     string
	IndexIDs     []string
}

// set stuff
func (meta *metadata) CreateIndex() {
	if len(meta.InstancesIDs) > 0 {
		keys := make([]string, 0, len(meta.InstancesIDs))
		for id := range meta.InstancesIDs {
			keys = append(keys, id)
		}
		sort.Strings(keys)
		meta.IndexIDs = keys
	}
}

// order items
type sortableMetadata []metadata

func (list sortableMetadata) Len() int {
	return len(list)
}

func (list sortableMetadata) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list sortableMetadata) Less(i, j int) bool {
	return len(list[i].InstancesIDs) < len(list[j].InstancesIDs)
}
