# go sdk for NSX

This is the repository that contains helper SDK for the [Terraform NSV Provider]( https://github.com/IBM-tfproviders/terraform-provider-nsxv), 
which one can use with Terraform to work with [VMware NSX-V](https://www.vmware.com/products/nsx.html).

# Build information

In order to reduce the security risk build this SDK using the latest version of golang. 
You may also need to upgrade the version of required modules in go.mod and regenerate the go.sum for any security issues.
## Steps to rebuild this SDK

- export GOPATH=<YOUR_GO_PATH>
- export GO_BIN_LOCATION=<YOUR_GO_BIN_PATH>
- export GOVX_VERSION=1.0.2
- mkdir -p $GOPATH/src/github.com/IBM-tfproviders
- git clone -b v$NSXV_VERSION https://github.com/IBM-tfproviders/govnsx
- cd $GOPATH/src/github.com/IBM-tfproviders/govnsx
- $GO_BIN_LOCATION mod tidy
- $GO_BIN_LOCATION install