package nsxresource

import (
	"encoding/xml"
	"fmt"

	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
)

type EdgeDhcp struct {
	Common
}

func NewEdgeDhcp(c *govnsx.Client) *EdgeDhcp {
	return &EdgeDhcp{
		Common: NewCommon(c),
	}
}

//PUT Method for configuring an Edge with DHCP Service
func (ed EdgeDhcp) Put(dhcpSpec *nsxtypes.ConfigDHCPServiceSpec,
	edgeId string) error {

	putUri := fmt.Sprintf(nsxtypes.EdgeDhcpUriFormat, ed.Nsxc.MgrConfig.Uri, edgeId)

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
		err := fmt.Errorf("[ERROR] %d : %s,\n XML: %s\n URI:%s\n",
			resp.StatusCode(),
			resp.Status(), outputXML, putUri)
		return err
	}

	return nil
}

//DELETE Method to delete DHCP configuration from DHCP Service
func (ed EdgeDhcp) Delete(edgeId string) error {

	deleteUri := fmt.Sprintf(nsxtypes.EdgeDhcpUriFormat,
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

//POST Method to add IP Pool into DHCP configuration
func (ed EdgeDhcp) Post(ipPool *nsxtypes.IPPool, edgeId string) (
	*nsxtypes.AddIPPoolToDHCPServiceResp, error) {

	postUri := fmt.Sprintf(nsxtypes.EdgeDhcpAddIPPoolUriFormat,
		ed.Nsxc.MgrConfig.Uri, edgeId)

	outputXML, err := xml.MarshalIndent(ipPool, "  ", "    ")
	if err != nil {
		return nil, err
	}

	resp, err := ed.Nsxc.Rclient.R().SetBody(outputXML).Post(postUri)
	if err != nil {
		return nil, err
	}

	sc := resp.StatusCode()
	if (sc < 200) || (sc > 204) {
		err := fmt.Errorf("[ERROR] %d : %s,\n XML: %s\n URI:%s\n",
			resp.StatusCode(),
			resp.Status(), outputXML, postUri)
		return nil, err
	}

	ipPoolResp := &nsxtypes.AddIPPoolToDHCPServiceResp{
		Location: resp.RawResponse.Header.Get("Location"),
	}
	return ipPoolResp, nil
}

//DELETE Method to delete IP Pool from DHCP configuration
func (ed EdgeDhcp) DeleteIPPool(edgeId string, ipPoolId string) error {

	deleteUri := fmt.Sprintf(nsxtypes.EdgeDhcpDelIPPoolUriFormat,
		ed.Nsxc.MgrConfig.Uri, edgeId, ipPoolId)

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
