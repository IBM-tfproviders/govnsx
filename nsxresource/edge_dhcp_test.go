package nsxresource

import (
	"fmt"
	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
	"path"
	"testing"
)

func getNsxClient() (*govnsx.Client, error) {
	nsxMgrParams := &govnsx.NsxManagerConfig{
		UserName:      "admin",
		Password:      "passw0rd",
		Uri:           "https://9.5.28.153",
		AllowInsecssl: true,
		UserAgentName: "Test Agent",
	}
	nsxClient, err := govnsx.NewClient(nsxMgrParams)

	if err != nil {
		return nil, err
	}

	return nsxClient, nil
}

func TestConfigDHCPService(t *testing.T) {

	nsxClient, err := getNsxClient()
	if err != nil {
		t.Fatalf("[Error] NewClient() returned error : %v", err)
		return
	}

	edgeDHCPObj := NewEdgeDhcp(nsxClient)

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

	edgeId := "edge-181"

	err = edgeDHCPObj.Put(dhcpSpec, edgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Put() returned error : %v", err)
		return
	}
}

func TestDeleteDHCPService(t *testing.T) {

	nsxClient, err := getNsxClient()
	if err != nil {
		t.Fatalf("[Error] NewClient() returned error : %v", err)
		return
	}

	edgeDHCPObj := NewEdgeDhcp(nsxClient)

	edgeId := "edge-181"

	err = edgeDHCPObj.Delete(edgeId)

	if err != nil {
		t.Fatalf("[Error] dhcp.Put() returned error : %v", err)
		return
	}
}

func TestAddIPPoolToDHCPService(t *testing.T) {

	nsxClient, err := getNsxClient()
	if err != nil {
		t.Fatalf("[Error] NewClient() returned error : %v", err)
		return
	}

	edgeDHCPObj := NewEdgeDhcp(nsxClient)

	ipPoolSpec := &nsxtypes.IPPool{
		IPRange:             "192.168.4.192-192.168.4.220",
		DefaultGw:           "192.168.4.1",
		SubnetMask:          "255.255.255.0",
		DomainName:          "eng.vmware.com",
		PrimaryNameServer:   "192.168.4.1",
		SecondaryNameServer: "4.2.2.4",
		LeaseTime:           3600,
		AutoConfigureDNS:    true}

	edgeId := "edge-181"

	ipPoolPostResp, err := edgeDHCPObj.Post(ipPoolSpec, edgeId)

	if err != nil {
		t.Fatalf("[Error] addIPPool.Post() returned error : %v", err)
		return
	}

	fmt.Println("Added IP Pool to Edge DHCP: ", ipPoolPostResp.Location)

	ipPoolId := path.Base(ipPoolPostResp.Location)
	fmt.Println("Deleting IP Pool from Edge DHCP: ", ipPoolId)

	err = edgeDHCPObj.DeleteIPPool(edgeId, ipPoolId)

	if err != nil {
		t.Fatalf("[Error] deleteIPPool.Delete() returned error : %v", err)
		return
	}
}
