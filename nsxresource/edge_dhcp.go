package nsxresource

import (
	"encoding/xml"
	"fmt"

	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
)

type EdgeDHCP struct {
	Common
}

func NewEdgeDHCP(c *govnsx.Client) *EdgeDHCP {
	return &EdgeDHCP{
		Common: NewCommon(c),
	}
}

//PUT Method for configuring an Edge with DHCP Service
func (ed EdgeDHCP) Put(dhcpSpec *nsxtypes.ConfigDHCPServiceSpec,
	edgeId string) error {

	putUri := fmt.Sprintf(nsxtypes.EdgeDHCPUriFormat, ed.Nsxc.MgrConfig.Uri, edgeId)

	outputXML, err := xml.MarshalIndent(dhcpSpec, "  ", "    ")
	if err != nil {
		return err
	}

	resp, err := ed.Nsxc.Rclient.R().SetBody(outputXML).Put(putUri)
	if err != nil {
		return err
	}

	sc := resp.StatusCode()
	if (sc < 200) || (sc > 204) {
        err := fmt.Errorf("[ERROR] %d : %s,\n XML: %s\n URI:%s\n, Body:%s",
			resp.StatusCode(),
			resp.Status(), outputXML, putUri, resp.Body())
		return err
	}

	return nil
}

//GET Method to get DHCP configuration from DHCP Service
func (ed EdgeDHCP) Get(edgeId string) (*nsxtypes.DHCPConfig, error) {

	getUri := fmt.Sprintf(nsxtypes.EdgeDHCPUriFormat,
		ed.Nsxc.MgrConfig.Uri, edgeId)

	resp, err := ed.Nsxc.Rclient.R().Get(getUri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		err := fmt.Errorf("[ERROR] %d : %s,\n URI:%s\n",
			resp.StatusCode(),
			resp.Status(), getUri)
		return nil, err
	}

	edge_dhcp := nsxtypes.NewDHCPConfig()

	err = xml.Unmarshal(resp.Body(), edge_dhcp)
	if err != nil {
		return nil, err
	}
	return edge_dhcp, nil
}

//DELETE Method to delete DHCP configuration from DHCP Service
func (ed EdgeDHCP) Delete(edgeId string) error {

	deleteUri := fmt.Sprintf(nsxtypes.EdgeDHCPUriFormat,
		ed.Nsxc.MgrConfig.Uri, edgeId)

	resp, err := ed.Nsxc.Rclient.R().Delete(deleteUri)
	if err != nil {
		return err
	}

	sc := resp.StatusCode()
	if (sc < 200) || (sc > 204) {
		err := fmt.Errorf("[ERROR] %d : %s,\n URI:%s\n",
			resp.StatusCode(),
			resp.Status(), deleteUri)
		return err
	}

	return nil
}
