package nsxtypes

import (
	"encoding/xml"
	"testing"
)

func TestVirtualWireCreateSpecType_1(t *testing.T) {

	data := `<virtualWireCreateSpec>
      <name>name_test1</name>
      <description>spec testing</description>
      <tenantId>test tid</tenantId>
      <ControlPlaneMode>UNICAST</ControlPlaneMode>
      <GuestVlanAllowed>true</GuestVlanAllowed>
 </virtualWireCreateSpec>
`
	r := &VirtualWireCreateSpec{
		Name: "",
	}

	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		t.Fatalf("xml.Unmarshal()  failed : %v\n", err)
	}
	assertValuecheck(t, "Name", r.Name, "name_test1")
	assertValuecheck(t, "Description", r.Description, "spec testing")
	assertValuecheck(t, "TenantId", r.TenantId, "test tid")
	assertValuecheck(t, "ControlPlaneMode", r.ControlPlaneMode, "UNICAST")
	assertValuecheck(t, "GuestVlanAllowed", r.GuestVlanAllowed, "true")
}

func TestVirtualWireDecode_1(t *testing.T) {
	data := `
		<?xml version="1.0" encoding="UTF-8"?>
		<virtualWire>
	   <objectId>virtualwire-104</objectId>
	  <objectTypeName>VirtualWire</objectTypeName>
	  <vsmUuid>4230B4F8-24C1-E884-B70C-9EA9DA7E1FFC</vsmUuid>
	  <nodeId>cf9735f6-2d0f-429c-984d-2d0dcf651b00</nodeId>
	  <revision>2</revision>
	  <type>
		<typeName>VirtualWire</typeName>
	  </type>
	  <name>9569bfb0-2433-48d6-87ae-72937c1a2dc5</name>
	  <clientHandle/>
	  <extendedAttributes/>
	  <isUniversal>false</isUniversal>
	  <universalRevision>0</universalRevision>
	  <tenantId>virtual wire tenant</tenantId>
	  <vdnScopeId>vdnscope-3</vdnScopeId>
	  <vdsContextWithBacking>
		<switch>
		  <objectId>dvs-61</objectId>
		  <objectTypeName>VmwareDistributedVirtualSwitch</objectTypeName>
		  <vsmUuid>4230B4F8-24C1-E884-B70C-9EA9DA7E1FFC</vsmUuid>
		  <nodeId>cf9735f6-2d0f-429c-984d-2d0dcf651b00</nodeId>
		  <revision>344</revision>
		  <type>
			<typeName>VmwareDistributedVirtualSwitch</typeName>
		  </type>
		  <name>dvSwitch1</name>
		  <scope>
			<id>datacenter-2</id>
			<objectTypeName>Datacenter</objectTypeName>
			<name>NSX Manage To</name>
		  </scope>
		  <clientHandle/>
		  <extendedAttributes/>
		  <isUniversal>false</isUniversal>
		  <universalRevision>0</universalRevision>
		</switch>
		<mtu>1600</mtu>
		<promiscuousMode>false</promiscuousMode>
		<backingType>portgroup</backingType>
		<backingValue>dvportgroup-693</backingValue>
	  </vdsContextWithBacking>
	  <vdnId>5000</vdnId>
	  <guestVlanAllowed>false</guestVlanAllowed>
	  <controlPlaneMode>UNICAST_MODE</controlPlaneMode>
	  <ctrlLsUuid>2dd5e9bf-87f2-480d-8f3a-09639d64092f</ctrlLsUuid>
	  <macLearningEnabled>false</macLearningEnabled>
	</virtualWire>
	`
	r := VirtualWire{
		ObjectId: "",
		TenantId: "",
	}

	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		t.Fatalf("xml.Unmarshal()  failed : %v\n", err)
	}

	assertValuecheck(t, "ObjectId", r.ObjectId, "virtualwire-104")
	assertValuecheck(t, "TenantId", r.TenantId, "virtual wire tenant")
	assertValuecheck(t, "SwitchOId", r.SwitchOId, "dvs-61")
	assertValuecheck(t, "SwitchName", r.SwitchName, "dvSwitch1")
	assertValuecheck(t, "VdnId", r.VdnId, "5000")
	assertValuecheck(t, "GuestVlanAllowed", r.GuestVlanAllowed, "false")
	assertValuecheck(t, "ControlPlaneMode", r.ControlPlaneMode, "UNICAST_MODE")
}

func assertValuecheck(t *testing.T, name string, value string, expected string) {
	if value != expected {
		t.Fatalf("expected %s value %s, got %v", name, expected, value)
	}
	return
}
