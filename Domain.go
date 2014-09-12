package fast

import (
	"fmt"
)

type DomainInfo struct {
	DomainId   string `json:"domainId"`
	DomainName string `json:"domainName"`
}

type DomainInterface interface {
	Register()
}

func (this *DomainInfo) Register() {
	fmt.Println("DomainInfo..., domainId:", this.DomainId, "domainName:,", this.DomainName)
}
