package nsxresource

import (
	"fmt"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
	"testing"
)

func TestAddInterfaces(t *testing.T) {
	nsxClient, err := getNsxClient()
	if err != nil {
		t.Fatalf("[Error] NewClient() returned error : %v", err)
		return
	}

	netobj := NewNetwork(nsxClient)
	vWspec := nsxtypes.NewVWCreateSpec()
	vWspec.Name = "Test net1"
	vWspec.Description = "Testing net"
	vWspec.TenantId = "virtual wire tenant1"
	vWspec.ControlPlaneMode = "UNICAST_MODE"
	vWspec.GuestVlanAllowed = false

	// scopeId := "vdnscope-3"
	scopeId, err := ReadEnv("NSX_VDN_SCOPE")
	if scopeId == "" {
		t.Fatalf("[Error] NSX_VDN_SCOPE is not set")
		return
	}

	vWpostresp, err := netobj.Post(vWspec, scopeId)

	if err != nil {
		t.Fatalf("[Error] Network.Post()  returned error : %v", err)
		return
	}

	fmt.Println("Created Virtual Wire %s: ", vWpostresp.Location)
	fmt.Println("Created Virtual Wire %s: ", vWpostresp.VirtualWireOID)
	vwId := vWpostresp.VirtualWireOID

	edge := NewEdge(nsxClient)

	appliances := nsxtypes.Appliances{DeployAppliances: false}

	edgeInstallSpec := &nsxtypes.EdgeInstallSpec{
		Name:        "Edge-Dhcp-UT2",
		Type:        "distributedRouter",
		Description: "Edge-Dhcp-UT2",
		Tenant:      "Terraform Provider",
		Appliances:  appliances,
	}

	resp, err := edge.Post(edgeInstallSpec)

	if err != nil {
		t.Fatalf("[Error] edge.Post () returned error : %v", err)
		return
	}

	fmt.Println("Created Edge: ", resp.EdgeId)

	edgeId := resp.EdgeId

	dlrInterfaces := NewEdgeDLRInterfaces(nsxClient)

	addrGroups := []nsxtypes.AddressGroup{nsxtypes.AddressGroup{
		PrimaryAddress: "10.10.10.2",
		SubnetMask:     "255.255.255.0"}}

	interfaces := []nsxtypes.EdgeDLRInterface{nsxtypes.EdgeDLRInterface{
		Name:          "Test-Interface",
		AddressGroups: addrGroups,
		IsConnected:   true,
		ConnectedToId: vwId}}

	addInterfacesSpec := &nsxtypes.EdgeDLRAddInterfacesSpec{
		EdgeDLRInterfaceList: interfaces,
	}

	addresp, err := dlrInterfaces.Post(addInterfacesSpec, edgeId)
	if err != nil {
		t.Fatalf("[Error] dlrInterfaces.Post () returned error : %v", err)
		return
	}
	fmt.Println("%v", addresp)

	var index string
	index = addresp.EdgeDLRInterfaceList[0].Index
	err = dlrInterfaces.Delete(edgeId, index)
	if err != nil {
		t.Fatalf("[Error] dlrInterfaces.DELETE() returned error : %v", err)
		return
	}
	fmt.Println("Deleted Edge Interface: ", resp.EdgeId, "-", index)

	err = edge.Delete(edgeId)

	if err != nil {
		t.Fatalf("[Error] edge.Delete () returned error : %v", err)
		return
	}
	fmt.Println("Deleted Edge: ", resp.EdgeId)
}
