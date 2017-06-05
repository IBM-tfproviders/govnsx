package nsxresource

import (
	"fmt"
	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
	"path"
	"testing"
)

func getNsxClient() (*govnsx.Client, error) {

	nsxMgrParams, err := GetNsxManagerConfig()
	if err != nil {
		fmt.Errorf("[Error] Test setup for NSX failed: %v", err)
		return nil, err
	}

	nsxClient, err := govnsx.NewClient(nsxMgrParams)

	if err != nil {
		return nil, err
	}

	return nsxClient, nil
}

func TestInstallEdge(t *testing.T) {

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

	vwire, err := netobj.Get(vWpostresp.Location)
	if err != nil {
		t.Fatalf("[Error] Network.Get()  returned error : %v", err)
		return
	}

	fmt.Println("Created Virtual Wire %s: ", vWpostresp.Location)

	edge := NewEdge(nsxClient)

	var addrGroups = []nsxtypes.AddressGroup{nsxtypes.AddressGroup{
		PrimaryAddress: "10.10.10.2",
		SubnetMask:     "255.255.255.0"}}

	var vnics = []nsxtypes.Vnic{nsxtypes.Vnic{
		Index:         "0",
		PortgroupId:   vwire.ObjectId,
		AddressGroups: addrGroups,
		IsConnected:   true}}

	resPoolId, err := ReadEnv("NSX_RESOURCE_POOL_ID")
	if resPoolId == "" {
		t.Fatalf("[Error] NSX_RESOURCE_POOL_ID is not set")
		return
	}

	dataStoreId, err := ReadEnv("NSX_DATASTORE_ID")
	if dataStoreId == "" {
		t.Fatalf("[Error] NSX_DATASTORE_ID is not set")
		return
	}

	dataCenterId, err := ReadEnv("NSX_DATACENTER")
	if dataCenterId == "" {
		t.Fatalf("[Error] NSX_DATACENTER is not set")
		return
	}

	appliances := nsxtypes.Appliances{ApplianceSize: "compact",
		DeployAppliances: true, AppliancesList: []nsxtypes.Appliance{
			nsxtypes.Appliance{ResourcePoolId: resPoolId,
				DatastoreId: dataStoreId}}}

	edgeInstallSpec := &nsxtypes.EdgeInstallSpec{
		Name:        "Edge-Dhcp-UT1",
		Type:        "gatewayServices",
		Description: "Edge-Dhcp-UT1",
		Datacenter:  dataCenterId,
		Tenant:      "virtual wire tenant",
		Appliances:  appliances,
		Vnics:       vnics,
	}

	resp, err := edge.Post(edgeInstallSpec)

	if err != nil {
		t.Fatalf("[Error] edge.Post () returned error : %v", err)
		return
	}

	fmt.Println("Created Edge: ", resp.EdgeId)

	// /api/4.0/edges/edge-183 ==> edge-183

	edgeDHCPObj := NewEdgeDHCP(nsxClient)

	var ipPools = []nsxtypes.IPPool{nsxtypes.IPPool{
		IPRange:             "192.168.4.192-192.168.4.220",
		DefaultGw:           "192.168.4.1",
		SubnetMask:          "255.255.255.0",
		DomainName:          "eng.vmware.com",
		PrimaryNameServer:   "192.168.4.1",
		SecondaryNameServer: "4.2.2.4",
		LeaseTime:           3600,
		AutoConfigureDNS:    true}}

	var logInfo = nsxtypes.LoggingInfo{Enable: true, LogLevel: "info"}
	dhcpSpec := &nsxtypes.ConfigDHCPServiceSpec{
		IPPools: ipPools,
		Logging: logInfo,
	}

	edgeId := resp.EdgeId

	err = edgeDHCPObj.Put(dhcpSpec, edgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Put() returned error : %v", err)
		return
	}

	edgeDHCPIPPoolObj := NewEdgeDHCPIPPool(nsxClient)
	ipPoolSpec := &nsxtypes.IPPool{
		IPRange:             "192.168.5.192-192.168.5.220",
		DefaultGw:           "192.168.5.1",
		SubnetMask:          "255.255.255.0",
		DomainName:          "eng.vmware.com",
		PrimaryNameServer:   "192.168.5.1",
		SecondaryNameServer: "4.2.2.4",
		LeaseTime:           3600,
		AutoConfigureDNS:    true}

	ipPoolPostResp, err := edgeDHCPIPPoolObj.Post(ipPoolSpec, edgeId)

	if err != nil {
		t.Fatalf("[Error] addIPPool.Post() returned error : %v", err)
		return
	}

	fmt.Println("Added IP Pool to Edge DHCP: ", ipPoolPostResp.Location)

	_, err = edgeDHCPObj.Get(edgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Get() returned error : %v", err)
		return
	}

	ipPoolId := path.Base(ipPoolPostResp.Location)
	fmt.Println("Deleting IP Pool from Edge DHCP: ", ipPoolId)

	err = edgeDHCPIPPoolObj.Delete(edgeId, ipPoolId)

	if err != nil {
		t.Fatalf("[Error] deleteIPPool.Delete() returned error : %v", err)
		return
	}

	err = edgeDHCPObj.Delete(edgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Put() returned error : %v", err)
		return
	}

	fmt.Println("Deleting Edge ID: ", resp.EdgeId)

	err = edge.Delete(resp.EdgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Delete() returned error : %v", err)
		return
	}

	fmt.Println("Deleting Virtual Wire: ", vWpostresp.Location)
	err = netobj.Delete(vWpostresp.Location)

	if err != nil {
		t.Fatalf("[Error] Network.Delete()  returned error : %v", err)
		return
	}

}
