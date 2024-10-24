package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/config"
	"github.com/gophercloud/gophercloud/v2/openstack/config/clouds"
)

func New() (*gophercloud.ProviderClient, *gophercloud.EndpointOpts, error) {
	cloudAuthOpts := []clouds.ParseOption{
		clouds.WithCloudName("OS_CLOUD_NAME"),
		clouds.WithIdentityEndpoint("OS_AUTH_URL"),
		clouds.WithProjectName("OS_PROJECT_NAME"),
		clouds.WithUsername("OS_USERNAME"),
		clouds.WithPassword("OS_PASSWORD"),
		clouds.WithRegion("OS_REGION"),
	}
	authOptions, endpointOptions, tlsConfig, err := clouds.Parse(cloudAuthOpts...)
	if err != nil {
		return nil, nil, err
	}
	ctx := context.Background()
	providerClient, err := config.NewProviderClient(ctx, authOptions, config.WithTLSConfig(tlsConfig))
	if err != nil {
		return nil, nil, err
	}
	return providerClient, &endpointOptions, nil
}
