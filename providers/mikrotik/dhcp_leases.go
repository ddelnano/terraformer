package mikrotik

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ddelnano/terraform-provider-mikrotik/client"
)

type DhcpLeaseGenerator struct {
	MikrotikService
}

func (g DhcpLeaseGenerator) createResources(leases []client.DhcpLease) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for _, lease := range leases {
		resources = append(resources, terraform_utils.NewSimpleResource(
			lease.Id,
			// # TODO: prefer hostname if set, otherwise use Id
			lease.Id,
			"mikrotik_dhcp_lease",
			"mikrotik",
			[]string{}))
	}
	return resources
}

func (g *DhcpLeaseGenerator) InitResources() error {
	client := g.generateClient()
	leases, err := client.ListDhcpLeases()

	if err != nil {
		return err
	}
	g.Resources = g.createResources(leases)
	return nil
}
