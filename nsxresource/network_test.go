package nsxresource

import (
	"fmt"
	"os"

	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
	"testing"
)

func ReadEnv(evar string) (string, error) {
	if v := os.Getenv(evar); v != "" {
		return v, nil
	}
	return "", fmt.Errorf("[ERROR] Environment variable %s not set", evar)
}

func GetNsxManagerConfig() (*govnsx.NsxManagerConfig, error) {
	nsxMgrParams := &govnsx.NsxManagerConfig{
		AllowInsecssl: true,
		UserAgentName: "Test Agent",
	}

	v, err := ReadEnv("NSX_USER")
	if err != nil {
		return nil, err
	}
	nsxMgrParams.UserName = v

	v, err = ReadEnv("NSX_PASSWORD")
	if err != nil {
		return nil, err
	}
	nsxMgrParams.Password = v

	v, err = ReadEnv("NSX_URI")
	if err != nil {
		return nil, err
	}
	nsxMgrParams.Uri = v

	return nsxMgrParams, nil
}

func TestNetworkResource_test1(t *testing.T) {
	nsxMgrParams, err := GetNsxManagerConfig()
	if err != nil {
		t.Fatalf("[Error] Test setup for NSX failed: %v", err)
		return
	}

	nsxclient, err := govnsx.NewClient(nsxMgrParams)

	if err != nil {
		t.Fatalf("[Error] NewClient() returned error : %v", err)
		return
	}

	netobj := NewNetwork(nsxclient)
	vWspec := nsxtypes.NewVWCreateSpec()
	vWspec.Name = "Test net1"
	vWspec.Description = "Testing net"
	vWspec.TenantId = "virtual wire tenant1"
	vWspec.ControlPlaneMode = "UNICAST_MODE"
	vWspec.GuestVlanAllowed = false

	scopeId := "vdnscope-3"

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

	if vwire.Name != vWspec.Name {
		t.Fatalf("[Error] Network.Post() and Network.Get() Doesn't match the name")
		return
	}

	UpdatedVWire := nsxtypes.NewUpdateVirtualWire()
	UpdatedVWire.Name = "New Name"
	UpdatedVWire.ControlPlaneMode = "UNICAST_MODE"

	err = netobj.Put(UpdatedVWire, vWpostresp.Location)
	if err != nil {
		t.Fatalf("[Error] Network.Put()  returned error : %v", err)
		return
	}

	vwire1, err := netobj.Get(vWpostresp.Location)
	if err != nil {
		t.Fatalf("[Error] Network.Get()  returned error after update : %v", err)
		return
	}

	if vwire1.Name != UpdatedVWire.Name {
		t.Fatalf("[Error] Network.put() name didn't change")
		return
	}

	err = netobj.Delete(vWpostresp.Location)

	if err != nil {
		t.Fatalf("[Error] Network.Delete()  returned error : %v", err)
		return
	}
}
