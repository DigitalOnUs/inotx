package config

import (
	"errors"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/fatih/color"
)

// AddConsul to the topology per Region
func AddConsul(root *Root) (*Root, error) {
	var consul Root

	// services remains the same
	consul.Services = root.Services

	// base case no DC
	if len(root.Datacenters) < 1 {
		return nil, errors.New("There is no datacenter specified in the blueprint")
	}

	for _, dc := range root.Datacenters {
		createConsulClients(root, &consul, dc)
	}

	return &consul, nil
}

//clientSize
func createConsulClients(src *Root, dest *Root, dc *Datacenter) {
	counter := make(chan map[string]metadata)
	counterFw := make(chan map[string]X)

	// counting servs
	go func() {
		defer close(counter)

		serviceCount := make(map[string]metadata)
		for _, r := range src.Resources {
			// Service Pool size
			if r != nil && r.Type == ServicePool.String() && r.Location == dc.FQDN() {

				serviceCount[r.Name] = metadata{
					InstancesIDs: make(map[string]X),
					Location:     r.Location,
				}

				// contains
				if len(r.Associations) == 0 {
					continue
				}

				for _, assoc := range r.Associations {
					if assoc.Type == Contains.String() {
						serviceCount[r.Name].InstancesIDs[assoc.ID] = X{}
					}
				}
			}
		}
		counter <- serviceCount
	}()

	// firewall resources
	go func() {
		defer close(counterFw)
		serviceCount := make(map[string]X)

		for _, r := range src.Resources {
			if r != nil && r.Type == Firewall.String() && r.Location == dc.FQDN() {
				if len(r.Associations) == 0 {
					continue
				}

				for _, assoc := range r.Associations {
					// we don't care the type just want to get services
					if strings.HasPrefix(assoc.ID, "service.") {
						serviceCount[assoc.ID] = X{}
					}
				}

			}
		}

		counterFw <- serviceCount
	}()

	serviceCount := <-counter
	serviceInFW := <-counterFw

	// nothing to do
	if serviceCount == nil || len(serviceCount) == 0 {
		color.Yellow("No services in the spec no consul clients required")
		return
	}

	if serviceInFW != nil && len(serviceInFW) > 0 {
		crosscheck(serviceCount, serviceInFW, dc)
	}

	// sort by order everything we got
	serviceOrder := make([]metadata, 0, len(serviceCount))

	for k, v := range serviceCount {
		v.Name = k
		serviceOrder = append(serviceOrder, v)
	}

	sort.Stable(sortableMetadata(serviceOrder))

	// now sorting internals
	for i, serv := range serviceOrder {
		serv.CreateIndex()
		serviceOrder[i] = serv
	}

	//required elements
	clientNumber := len(serviceOrder[len(serviceOrder)-1].InstancesIDs)

	if clientNumber <= 0 {
		color.Yellow("No consul clients required after order because there are no instances in the service")
	}

	//consulClients := make([]*Resource, 0, clientNumber)
	color.Green("Required min clients %d", clientNumber)
	consulClients := make([]*Resource, 0, clientNumber)

	var wg sync.WaitGroup

	for i := 0; i < clientNumber; i++ {
		wg.Add(1)
		r := Resource{
			Type:         ConsulClient.String(),
			Name:         fmt.Sprintf("%s-client%d", dc.Name, i+1),
			Associations: make([]Association, 0, 0),
			Location:     dc.FQDN(),
		}
		consulClients = append(consulClients, &r)

		go func(res *Resource, id int) {
			defer wg.Done()
			for _, serv := range serviceOrder {
				if len(serv.IndexIDs) == 0 {
					continue
				}
				if id <= len(serv.IndexIDs)-1 {
					res.Associations = append(res.Associations, Association{
						ID:   serv.IndexIDs[id],
						Type: Contains.String(),
					})
				}
			}

		}(&r, i)
	}

	wg.Wait()

	// adding the cluster as well
	clusterSize := 3
	neighborConns := (clusterSize * 2) - 2

	consulServers := make([]*Resource, 0, clusterSize)
	uniques := make(map[string]X)
	for i := 0; i < clusterSize; i++ {
		server := Resource{
			Type:         ConsulServer.String(),
			Name:         fmt.Sprintf("%s-server%d", dc.Name, i+1),
			Associations: make([]Association, 0, neighborConns),
			Location:     dc.FQDN(),
		}
		consulServers = append(consulServers, &server)
		uniques[server.FQDN()] = X{}
	}

	// adding paths
	for _, server := range consulServers {
		for neighbor := range uniques {
			if neighbor != server.FQDN() {
				server.Associations = append(
					server.Associations,
					Association{
						ID:   neighbor,
						Type: Egress.String(),
					},
					Association{
						ID:   neighbor,
						Type: Ingress.String(),
					},
				)
			}
		}
	}

	// cluster
	resourceSize := len(consulClients) + len(consulServers)
	cluster := Resource{
		Type:         ConsulCluster.String(),
		Name:         fmt.Sprintf("cluster-%s", dc.Name),
		Location:     dc.FQDN(),
		Associations: make([]Association, 0, clusterSize),
	}

	// associations
	for member := range uniques {
		cluster.Associations = append(cluster.Associations,
			Association{
				ID:   member,
				Type: Contains.String(),
			})
	}

	// connecting cluster
	for _, client := range consulClients {
		client.Associations = append(client.Associations,
			Association{
				ID:   cluster.FQDN(),
				Type: Ingress.String(),
			},
			Association{
				ID:   cluster.FQDN(),
				Type: Egress.String(),
			},
		)
	}

	// counting the cluster
	resourceSize++

	finalResourceList := make([]*Resource, 0, resourceSize)
	finalResourceList = append(finalResourceList, consulClients...)
	finalResourceList = append(finalResourceList, consulServers...)
	finalResourceList = append(finalResourceList, &cluster)

	dest.Resources = finalResourceList
	dest.Datacenters = src.Datacenters
}

func crosscheck(services map[string]metadata, fw map[string]X, dc *Datacenter) {
	// getting all services
	current := map[string]X{}
	for _, meta := range services {
		for m := range meta.InstancesIDs {
			current[m] = X{}
		}
	}

	// doing checking if fw are missing ... crappy
	for f := range fw {
		if _, ok := current[f]; !ok {
			// creating own group like service.db.db1 -> fw.service.db
			uri := strings.Split(f, ".")
			if len(uri) < 3 {
				color.Yellow("error processing ref %s", f)
			}
			// to be added
			fqdn := fmt.Sprintf("fw.service.%s", uri[1])
			if _, ok := services[fqdn]; !ok {
				services[fqdn] = metadata{
					InstancesIDs: make(map[string]X),
					Location:     dc.FQDN(),
				}
			}

			services[fqdn].InstancesIDs[f] = X{}
		}
	}
}
