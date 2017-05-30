package nsxresource

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/IBM-tfproviders/govnsx"
	"github.com/IBM-tfproviders/govnsx/nsxtypes"
)

type Network struct {
	Common
}

func NewNetwork(c *govnsx.Client) *Network {
	return &Network{
		Common: NewCommon(c),
	}
}

//GET Method for a NSX network - VirtualWire
func (n Network) Get(location string) (*nsxtypes.VirtualWire, error) {

	resp, err := n.Nsxc.Rclient.R().Get(n.Nsxc.MgrConfig.Uri + location)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		err := fmt.Errorf("[ERROR] %d : %s", resp.StatusCode(), resp.Status())
		return nil, err
	}

	net := nsxtypes.NewVirtualWire()

	err = xml.Unmarshal(resp.Body(), net)

	if err != nil {
		return nil, err
	}

	return net, nil
}

//POST Method for a NSX network - VirtualWire
func (n Network) Post(vwSpec *nsxtypes.VWCreateSpec, scopeId string) (*nsxtypes.VWPostResp, error) {

	postUri := fmt.Sprintf(nsxtypes.VirtualWireUriFormat, n.Nsxc.MgrConfig.Uri, scopeId)

	outputXML, err := xml.MarshalIndent(vwSpec, "  ", "    ")
	if err != nil {
		return nil, err
	}

	resp, err := n.Nsxc.Rclient.R().SetBody(outputXML).Post(postUri)
	if err != nil {
		return nil, err
	}

	sc := resp.StatusCode()
	if (sc < 200) || (sc > 204) {
		err := fmt.Errorf("[ERROR] %d : %s", resp.StatusCode(), resp.Status())
		return nil, err
	}

	net := &nsxtypes.VWPostResp{
		Location:       resp.RawResponse.Header.Get("Location"),
		VirtualWireOID: string(resp.Body()),
	}
	return net, nil
}

//PUT Method for a NSX network - VirtualWire
func (n Network) Put(vWire *nsxtypes.UpdateVirtualWire, location string) error {

	outputXML, err := xml.MarshalIndent(vWire, "  ", "    ")
	if err != nil {
		return err
	}
	resp, err := n.Nsxc.Rclient.R().SetBody(outputXML).Put(n.Nsxc.MgrConfig.Uri + location)
	if err != nil {
		return err
	}

	sc := resp.StatusCode()

	if (sc < 200) || (sc > 204) {
		err := fmt.Errorf("[ERROR] %d : %s : %s", resp.StatusCode(),
			resp.Status(), string(resp.Body()))
		return err
	}

	return nil
}

//Delete Method for a NSX network - VirtualWire
func (n Network) Delete(location string) error {

	tryies := 2
	waitTime := 15 * time.Second
	sc := 0
	status := ""
	errMsg := "Devices may present on Network"

	for i := 0; i <= tryies; i++ {
		resp, err := n.Nsxc.Rclient.R().Delete(n.Nsxc.MgrConfig.Uri + location)
		if err != nil {
			return err
		}

		sc = resp.StatusCode()
		status = resp.Status()
		if sc == 400 {
			time.Sleep(waitTime)
			continue
		} else {
			break
		}
	}

	//For some reson resource already deleted
	//Lets consider delete action is successful
	if sc == 404 {
		log.Printf("[WARN] Network resource %s already deleted !", location)
		return nil
	} else if sc == 400 {
		//Bad delete request, may be device still on network
		err := fmt.Errorf("[ERROR] %d : %s : %s", sc, status, errMsg)
		return err
	} else if (sc < 200) || (sc > 204) {
		//other than done, accepted,..
		err := fmt.Errorf("[ERROR] %d : %s", sc, status)
		return err
	}

	return nil
}
