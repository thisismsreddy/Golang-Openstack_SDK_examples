package main

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {
	fmt.Println("Hello world")
	opts, err := openstack.AuthOptionsFromEnv()
	fmt.Println(opts, err)
	provider, err := openstack.AuthenticatedClient(opts)

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	fmt.Println("This line 21")
	server, err := servers.Get(client, "ad4e58d1-bc68-40f6-8c67-27a71caa33b8").Extract()

	fmt.Println("server name is ", server.Name, "server Id", server.ID, server.AccessIPv4, server.Status, server.SecurityGroups)
}
