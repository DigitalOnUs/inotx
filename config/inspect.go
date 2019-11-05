package config

import (
	"github.com/fatih/color"
)

// AddConsul to the topology
func AddConsul(root *Root) error {
	// looking for resources
	var DC Datacenter
	var consul Root

	if root.Datacenters == nil && len(root.Datacenters) < 1 {
		// creating a default dc
		color.Yellow("File without dc creating one for default")
		DC.Name = "privateDC"
		DC.Description = "autogenerated DC"
		DC.Default = true
		consul.Datacenters = []*Datacenter{&DC}
	}

	if root.Datacenters != nil && len(root.Datacenters) > 0 {
		consul.Datacenters = root.Datacenters
		// If only one that's the default
		for i, dc := range root.Datacenters {
			if dc != nil && dc.Default {
				color.Green(
					"Using a dc from list with number: ", i)
				DC = *(root.Datacenters[i])
				break
			}
		}

		// No default just pick one
		if DC.Name == "" {
			color.Yellow(
				"-> No default dc found, selecting one")
			DC = *(root.Datacenters[0])
		}

	}

	// looking for the default DC
	color.Green("Using DC %+v", DC)

	/* what we need to know
	- By default we have the cluster per region
	- The load balancers will help for the intention which are connected with whom
	  in the pool and the flow egress/ingress/contains
	- the services list will tell you the size max will calculate the clients
	*/

	return nil
}
