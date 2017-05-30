package nsxresource

import (
	"fmt"
	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
	"testing"
)

func TestInstallEdge(t *testing.T) {

	nsxMgrParams := &govnsx.NsxManagerConfig{
		UserName:      "admin",
		Password:      "passw0rd",
		Uri:           "https://9.5.28.153",
		AllowInsecssl: true,
		UserAgentName: "Test Agent",
	}

	nxclient, err := govnsx.NewClient(nsxMgrParams)

	if err != nil {
		t.Fatalf("[Error] NewClient() returned error : %v", err)
		return
	}

	edge := NewEdge(nxclient)

	var adgps = []nsxtypes.AddressGroup{nsxtypes.AddressGroup{
		PrimaryAddress: "10.10.10.2",
		SubnetMask:     "255.255.255.0"}}

	var vnics = []nsxtypes.Vnic{nsxtypes.Vnic{
		Index:         "0",
		PortgroupId:   "dvportgroup-962",
		AddressGroups: adgps,
		IsConnected:   true}}

	var applns = []nsxtypes.Appliance{nsxtypes.Appliance{
		ResourcePoolId: "resgroup-60",
		DatastoreId:    "datastore-16",
	}}

	edgeInstallSpec := &nsxtypes.EdgeSGWInstallSpec{
		Name:           "Edge-Dhcp-UT1",
		Description:    "Edge-Dhcp-UT1",
		Tenant:         "virtual wire tenant",
		AppliancesList: applns,
		Vnics:          vnics,
	}

	resp, err := edge.Post(edgeInstallSpec)

	if err != nil {
		t.Fatalf("[Error] dhcp.Put() returned error : %v", err)
		return
	}

	fmt.Println("Created Edge: ", resp.EdgeId)

	// /api/4.0/edges/edge-183 ==> edge-183

	fmt.Println("Deleting Edge ID: ", resp.EdgeId)

	err = edge.Delete(resp.EdgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Delete() returned error : %v", err)
		return
	}
}
