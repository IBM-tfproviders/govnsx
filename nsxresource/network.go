package nsxresource 

import (
	"encoding/xml"
	"fmt"

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
	resp, err := n.Nsxc.Rclient.R().Get(location)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200  {
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

