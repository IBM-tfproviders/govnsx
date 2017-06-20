package nsxresource

import (
	"encoding/xml"
	"fmt"

	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
)

type EdgeDHCPIPPool struct {
	Common
}

func NewEdgeDHCPIPPool(c *govnsx.Client) *EdgeDHCPIPPool {
	return &EdgeDHCPIPPool{
		Common: NewCommon(c),
	}
}

//POST Method to add IP Pool into DHCP configuration
func (ed EdgeDHCPIPPool) Post(ipPool *nsxtypes.IPPool, edgeId string) (
	*nsxtypes.AddIPPoolToDHCPServiceResp, error) {

	postUri := fmt.Sprintf(nsxtypes.EdgeDHCPAddIPPoolUriFormat,
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
		err := fmt.Errorf("[ERROR] %d : %s,\n XML: %s\n URI:%s\n Body:%s",
			resp.StatusCode(),
			resp.Status(), outputXML, postUri, resp.Body())
		return nil, err
	}

	ipPoolResp := &nsxtypes.AddIPPoolToDHCPServiceResp{
		Location: resp.RawResponse.Header.Get("Location"),
	}
	return ipPoolResp, nil
}

//DELETE Method to delete IP Pool from DHCP configuration
func (ed EdgeDHCPIPPool) Delete(edgeId string, ipPoolId string) error {

	deleteUri := fmt.Sprintf(nsxtypes.EdgeDHCPDelIPPoolUriFormat,
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
