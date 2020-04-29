package mikrotik

import (
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/ddelnano/terraform-provider-mikrotik/client"
)

type MikrotikService struct {
	terraform_utils.Service
}

func (m *MikrotikService) generateClient() client.Mikrotik {
	return client.NewClient(
		m.Args["host"].(string),
		m.Args["user"].(string),
		m.Args["password"].(string),
	)
}
