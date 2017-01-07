package main

import (
	"fmt"
	"os"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

func main() {
	opts, err := openstack.AuthOptionsFromEnv()
	fmt.Println(opts, err)
	provider, err := openstack.AuthenticatedClient(opts)

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	// To create a new server
	server, err := servers.Create(client, servers.CreateOpts{
		Name:      "My new server!",
		FlavorRef: "m1.small",
		ImageRef:  "ubuntu",
	}).Extract()
	fmt.Println("server name is ", server.Name, "server Id", server.ID, server.AccessIPv4, server.Status, server.SecurityGroups)

	//  Get a list of servers
	//  server, err := servers.Get(client, "server.ID").Extract()

}
