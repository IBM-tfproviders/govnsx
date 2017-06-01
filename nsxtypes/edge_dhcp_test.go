package nsxtypes

import (
	"encoding/xml"
	"reflect"
	"testing"
)

func TestConfigDHCPServiceSpec(t *testing.T) {

	data := `<dhcp>
          <ipPools>
              <ipPool>
                  <ipRange>192.168.4.192-192.168.4.220</ipRange>
                  <defaultGateway>192.168.4.1</defaultGateway>
                  <subnetMask>255.255.255.0</subnetMask>
                  <domainName>eng.vmware.com</domainName>
                  <primaryNameServer>192.168.4.1</primaryNameServer>
                  <secondaryNameServer>4.2.2.4</secondaryNameServer>
                  <leaseTime>3600</leaseTime>
                  <autoConfigureDNS>true</autoConfigureDNS>
              </ipPool>
          </ipPools>
          <logging>
              <enable>false</enable>
              <logLevel>info</logLevel>
          </logging>
          </dhcp>`

	r := &ConfigDHCPServiceSpec{}

	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		t.Fatalf("xml.Unmarshal() ConfigDHCPServiceSpec failed : %v\n", err)
	}

	ipPools := []IPPool{IPPool{XMLName: xml.Name{"", "ipPool"}, IPRange: "192.168.4.192-192.168.4.220", DefaultGw: "192.168.4.1",
		SubnetMask: "255.255.255.0", DomainName: "eng.vmware.com", PrimaryNameServer: "192.168.4.1",
		SecondaryNameServer: "4.2.2.4", LeaseTime: 3600, AutoConfigureDNS: true}}
	if !reflect.DeepEqual(r.IPPools, ipPools) {
		t.Fatalf("expected value %s, got %s", ipPools, r.IPPools)
	}

	loggingInfo := LoggingInfo{false, "info"}
	if !reflect.DeepEqual(r.Logging, loggingInfo) {
		t.Fatalf("expected value %s, got %s", loggingInfo, r.Logging)
	}
}

func TestAddIPPoolToDHCPServiceSpec(t *testing.T) {

	data := `<ipPool>
                  <ipRange>192.168.4.192-192.168.4.220</ipRange>
                  <defaultGateway>192.168.4.1</defaultGateway>
                  <subnetMask>255.255.255.0</subnetMask>
                  <domainName>eng.vmware.com</domainName>
                  <primaryNameServer>192.168.4.1</primaryNameServer>
                  <secondaryNameServer>4.2.2.4</secondaryNameServer>
                  <leaseTime>3600</leaseTime>
                  <autoConfigureDNS>true</autoConfigureDNS>
              </ipPool>`

	var r IPPool

	err := xml.Unmarshal([]byte(data), &r)

	if err != nil {
		t.Fatalf("xml.Unmarshal() AddIPPoolToDHCPServiceSpec failed : %v\n", err)
	}

	ipPool := IPPool{XMLName: xml.Name{"", "ipPool"}, IPRange: "192.168.4.192-192.168.4.220", DefaultGw: "192.168.4.1",
		SubnetMask: "255.255.255.0", DomainName: "eng.vmware.com", PrimaryNameServer: "192.168.4.1",
		SecondaryNameServer: "4.2.2.4", LeaseTime: 3600, AutoConfigureDNS: true}
	if !reflect.DeepEqual(r, ipPool) {
		t.Fatalf("expected value %s, got %s", ipPool, r)
	}
}
