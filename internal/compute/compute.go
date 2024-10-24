package compute

import (
	"context"
	"log"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/melbahja/goph"
)

type Manager struct {
	computeClient     *gophercloud.ServiceClient
	securityGroupName string
	networkName       string
}

func New(providerClient *gophercloud.ProviderClient, endpointOptions *gophercloud.EndpointOpts) (*Manager, error) {
	computeClient, err := openstack.NewComputeV2(providerClient, *endpointOptions)
	if err != nil {
		return nil, err
	}
	manager := &Manager{
		computeClient: computeClient,
	}
	if err := manager.prepareSecurityGroup(); err != nil {
		return nil, err
	}
	if err := manager.prepareNetwork(); err != nil {
		return nil, err
	}
	return manager, nil
}

func (m *Manager) prepareSecurityGroup() error {
	// TODO
	return nil
}

func (m *Manager) prepareNetwork() error {
	// TODO
	return nil
}

func (m *Manager) CreateServer(ctx context.Context, name string, image string, flavor string) (*servers.Server, error) {
	return servers.Create(ctx, m.computeClient, servers.CreateOpts{
		Name:             name,
		ImageRef:         image,
		FlavorRef:        flavor,
		SecurityGroups:   []string{m.securityGroupName},
		UserData:         []byte{},
		AvailabilityZone: "",
		Networks:         m.networkName,
		Metadata:         map[string]string{},
		Personality:      []*servers.File{},
		ConfigDrive:      new(bool),
		AdminPass:        "",
		AccessIPv4:       "",
		AccessIPv6:       "",
		Min:              1,
		Max:              1,
		Tags:             []string{},
		Hostname:         name,
		BlockDevice:      []servers.BlockDevice{},
		DiskConfig:       "",
	}, servers.SchedulerHintOpts{
		Group:                "",
		DifferentHost:        []string{},
		SameHost:             []string{},
		Query:                []any{},
		TargetCell:           "",
		DifferentCell:        []string{},
		BuildNearHostIP:      "",
		AdditionalProperties: map[string]any{},
	}).Extract()
}

func (m *Manager) ListServers(ctx context.Context, prefix string) ([]servers.Server, error) {
	paginatedServers := servers.List(m.computeClient, servers.ListOpts{
		Name: "<REGEX TO GET ALL SERVERS PREFIXED w/ PREFIX>",
	})
	page, err := paginatedServers.AllPages(ctx)
	if err != nil {
		return nil, err
	}
	return servers.ExtractServers(page)
}

func (m *Manager) DeleteServer(ctx context.Context, id string) error {
	return servers.Delete(ctx, m.computeClient, id).ExtractErr()
}

// SSH cereates SSH client to compute with ID. Caller must call defer client.close()
func (m *Manager) SSH(ID string) (*goph.Client, error) {
	auth, err := goph.Key("/home/ubuntu/.ssh/id_rsa", "")
	if err != nil {
		log.Fatal(err)
	}

	client, err := goph.New("root", "192.1.1.3", auth)
	if err != nil {
		return nil, err
	}

	return client, nil
}
