package nsxtypes

import (
	"encoding/xml"
	"testing"
)

func TestEdgeDlrAddInterfacesSpec(t *testing.T) {

	data := `<interfaces>
			<interface>
			<name>interface1</name>
			<addressGroups>
				<addressGroup>
				<primaryAddress>10.10.10.1</primaryAddress>
				<subnetMask>255.255.255.0</subnetMask>
				</addressGroup>
			</addressGroups>
			<mtu>1500</mtu>
			<type>Internal</type>
			<isConnected>true</isConnected>
			<connectedToId>virtualwire-210</connectedToId>
			</interface>

			<interface>
			<isConnected>true</isConnected>
			<connectedToId>virtualwire-212</connectedToId>
			</interface>
		</interfaces>`

	var r EdgeDLRAddInterfacesSpec
	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		t.Fatalf("xml.Unmarshal()  failed : %v\n", err)
	}

	assertValuecheck(t, "Name", r.EdgeDLRInterfaceList[0].Name, "interface1")
	assertValuecheck(t, "PrimaryAddress", r.EdgeDLRInterfaceList[0].AddressGroups[0].PrimaryAddress, "10.10.10.1")
	assertValuecheck(t, "ConnectedToId", r.EdgeDLRInterfaceList[1].ConnectedToId, "virtualwire-212")
}
