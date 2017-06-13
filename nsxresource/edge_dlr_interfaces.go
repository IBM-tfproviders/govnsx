package nsxresource

import (
	"encoding/xml"
	"fmt"

	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
)

type EdgeDLRInterfaces struct {
	Common
}

func NewEdgeDLRInterfaces(c *govnsx.Client) *EdgeDLRInterfaces {
	return &EdgeDLRInterfaces{
		Common: NewCommon(c),
	}
}

// Router modular APIs
// POST Method to add Interfaces
func (edlr EdgeDLRInterfaces) Post(interfaces *nsxtypes.EdgeDLRAddInterfacesSpec, edgeId string) (*nsxtypes.EdgeDLRAddInterfacesResp, error) {

	postUri := fmt.Sprintf(nsxtypes.EdgeDLRAddInterfacesUriFormat,
		edlr.Nsxc.MgrConfig.Uri, edgeId)

	outputXML, err := xml.MarshalIndent(interfaces, "  ", "    ")
	if err != nil {
		return nil, err
	}

	resp, err := edlr.Nsxc.Rclient.R().SetBody(outputXML).Post(postUri)
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

	addInterfaceResp := nsxtypes.NewEdgeDLRAddInterfacesResp()
	err = xml.Unmarshal(resp.Body(), &addInterfaceResp)
	if err != nil {
		return nil, err
	}
	return addInterfaceResp, nil
}

// DELETE method to remove Interface
func (edlr EdgeDLRInterfaces) Delete(edgeId string, index string) error {

	var deleteUri string
	fmt.Errorf("[ERROR] index #%s# \n", index)
	if index == "" { //Delete all intefaces
		deleteUri = fmt.Sprintf(
			nsxtypes.EdgeDLRDelAllInterfacesUriFormat,
			edlr.Nsxc.MgrConfig.Uri, edgeId)
	} else {
		deleteUri = fmt.Sprintf(
			nsxtypes.EdgeDLRDelbyIndexInterfacesUriFormat,
			edlr.Nsxc.MgrConfig.Uri, edgeId, index)
	}

	fmt.Errorf("[ERROR] deleteUri %s \n", deleteUri)
	resp, err := edlr.Nsxc.Rclient.R().Delete(deleteUri)
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
