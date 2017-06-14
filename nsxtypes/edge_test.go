package nsxtypes

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestEdgeInstallSpec(t *testing.T) {

	data := `<edge>
    <type>gatewayServices</type>
	<name>dhcp_sg_app</name> 
    <datacenterName>NSX Manage To</datacenterName>
	<description>Description for the edge gateway</description>
	<tenant>virtual wire tenant</tenant> 
	<appliances> 
        <deployAppliances>false</deployAppliances>
        <applianceSize>compact</applianceSize>
		<appliance>
			<resourcePoolId>resgroup-60</resourcePoolId>
			<datastoreId>datastore-16</datastoreId>
		</appliance>
	</appliances>
	<vnics> 
		<vnic>
			<index>0</index>
			<type>internal</type>
			<portgroupId>dvportgroup-962</portgroupId> 
			<addressGroups> 
			<addressGroup> 
			    <primaryAddress>10.10.10.2</primaryAddress> 
			    <subnetMask>255.255.255.0</subnetMask> 
			</addressGroup>
			</addressGroups>
			<isConnected>true</isConnected> 
			<mtu>1500</mtu> 
		</vnic>
	</vnics>
	</edge>`

	var r EdgeInstallSpec

	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		t.Fatalf("xml.Unmarshal()  failed : %v\n", err)
	}

	assertValuecheck(t, "Name", r.Name, "dhcp_sg_app")
	assertValuecheck(t, "Type", r.Type, "gatewayServices")
	assertValuecheck(t, "Description", r.Description, "Description for the edge gateway")
	assertValuecheck(t, "Tenant", r.Tenant, "virtual wire tenant")

	assertValuecheck(t, "ResourcePoolId", r.Appliances.AppliancesList[0].ResourcePoolId, "resgroup-60")
	assertValuecheck(t, "DatastoreId", r.Appliances.AppliancesList[0].DatastoreId, "datastore-16")

	assertValuecheck(t, "Index", r.Vnics[0].Index, "0")
	assertValuecheck(t, "PortgroupId", r.Vnics[0].PortgroupId, "dvportgroup-962")
	assertValuecheck(t, "IsConnected", fmt.Sprintf("%v", r.Vnics[0].IsConnected), "true")

	assertValuecheck(t, "PrimaryAddress", r.Vnics[0].AddressGroups[0].PrimaryAddress, "10.10.10.2")
	assertValuecheck(t, "SubnetMask", r.Vnics[0].AddressGroups[0].SubnetMask, "255.255.255.0")

}
