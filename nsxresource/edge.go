package nsxresource

import (
	"encoding/xml"
	"fmt"
	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
	"path"
)

type Edge struct {
	Common
}

func NewEdge(c *govnsx.Client) *Edge {
	return &Edge{
		Common: NewCommon(c),
	}
}

// GET  Method for a NSX Edge
func (e Edge) Get(edgeId string) (*nsxtypes.Edge, error) {

	getUri := fmt.Sprintf(nsxtypes.EdgeUriLocFormat, e.Nsxc.MgrConfig.Uri, edgeId)

	resp, err := e.Nsxc.Rclient.R().Get(getUri)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		err := fmt.Errorf("[ERROR] Get Edge Id %s received error: %s",
			 edgeId, resp.Status())
		return nil, err
	}

	edge := nsxtypes.NewEdge()

	err = xml.Unmarshal(resp.Body(), edge)
	if err != nil {
		return nil, err
	}
	return edge, nil
}

// POST Method for a NSX Edge
func (e Edge) Post(edgeSpec *nsxtypes.EdgeInstallSpec) (*nsxtypes.EdgePostResp, error) {

	postUri := fmt.Sprintf(nsxtypes.EdgeUriFormat, e.Nsxc.MgrConfig.Uri)

	outputXML, err := xml.MarshalIndent(edgeSpec, "  ", "    ")
	if err != nil {
		return nil, err
	}

	resp, err := e.Nsxc.Rclient.R().SetBody(outputXML).Post(postUri)
	if err != nil {
		return nil, err
	}

	sc := resp.StatusCode()
	if (sc < 200) || (sc > 204) {
		err := fmt.Errorf("[ERROR] %d : %s,\n XML: %s\n URI:%s\n", resp.StatusCode(),
			resp.Status(), outputXML, postUri)
		return nil, err
	}

	edgeId := path.Base(resp.RawResponse.Header.Get("Location"))

	edge := &nsxtypes.EdgePostResp{
		EdgeId:   edgeId,
		Location: resp.RawResponse.Header.Get("Location"),
	}

	return edge, nil
}

// PUT Method for a NSX Edge
func (e Edge) Put(edgeSpec *nsxtypes.EdgeInstallSpec, edgeId string) error {

	putUri := fmt.Sprintf(nsxtypes.EdgeUriLocFormat, e.Nsxc.MgrConfig.Uri, edgeId)

	outputXML, err := xml.MarshalIndent(edgeSpec, "  ", "    ")
	if err != nil {
		return err
	}

	resp, err := e.Nsxc.Rclient.R().SetBody(outputXML).Put(putUri)
	if err != nil {
		return err
	}

	sc := resp.StatusCode()
	if (sc < 200) || (sc > 204) {
		err := fmt.Errorf("[ERROR] %d : %s,\n XML: %s\n URI:%s\n", resp.StatusCode(),
			resp.Status(), outputXML, putUri)
		return err
	}

	return nil
}

// DELETE Method for NSX Edge
func (e Edge) Delete(edgeId string) error {

	deleteUri := fmt.Sprintf(nsxtypes.EdgeUriLocFormat, e.Nsxc.MgrConfig.Uri, edgeId)

	resp, err := e.Nsxc.Rclient.R().Delete(deleteUri)
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
