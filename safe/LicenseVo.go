package safe

import (
	"github.com/safawo/fast/comm"
	"strings"
)

const (
	LICENSE_TYPE_SUPER   = "Super"
	LICENSE_TYPE_FOREVER = "Forever"
	LICENSE_TYPE_DEVELOP = "Develop"
	LICENSE_TYPE_TEST    = "Test"
	LICENSE_TYPE_ISSUE   = "Business"
	LICENSE_TYPE_TEMP    = "Temp"
)

type FastLicenseInfo struct {
	Distributor  string `json:"distributor"`
	DistributeId string `json:"distributeId"`

	LicenseId   string `json:"licenseId"`
	LicenseType string `json:"licenseType"`

	CustomerId   string `json:"customerId"`
	CustomerName string `json:"customerName"`

	BindNetCard string `json:"bindNetCard"`
	BindDisk    string `json:"bindDisk"`
	BindCPU     string `json:"bindCPU"`

	ExpiryDate string `json:"expiryDate"`
	Remark     string `json:"remark"`

	AllowStr  map[string](string)   `json:"allowStr"`
	AllowInt  map[string](int)      `json:"allowInt"`
	AllowStrs map[string]([]string) `json:"allowStrs"`
	AllowInts map[string]([]int)    `json:"allowInts"`
}

func (this *FastLicenseInfo) Reset() {
	this.Distributor = comm.NULL_STR
	this.DistributeId = comm.NULL_STR

	this.LicenseId = comm.NULL_STR
	this.LicenseType = comm.NULL_STR

	this.CustomerId = comm.NULL_STR
	this.CustomerName = comm.NULL_STR

	this.BindNetCard = comm.NULL_STR
	this.BindDisk = comm.NULL_STR
	this.BindCPU = comm.NULL_STR

	this.ExpiryDate = comm.NULL_STR
	this.Remark = comm.NULL_STR

	this.AllowStr = map[string](string){}
	this.AllowInt = map[string](int){}
	this.AllowStrs = map[string]([]string){}
	this.AllowInts = map[string]([]int){}

}

func (this *FastLicenseInfo) IsEmpty() bool {
	if this.Distributor == comm.NULL_STR {
		return true
	}

	return false
}

func (this *FastLicenseInfo) GetStr(allowName string) string {
	allowValue, ok := this.AllowStr[allowName]
	if !ok {
		return comm.NULL_STR
	}

	allowValue = strings.TrimSpace(allowValue)

	return allowValue
}

func (this *FastLicenseInfo) GetInt(allowName string) int {
	allowValue, ok := this.AllowInt[allowName]
	if !ok {
		return comm.NULL_INT
	}

	return allowValue
}

func (this *FastLicenseInfo) GetStrs(allowName string) []string {
	allowValue, ok := this.AllowStrs[allowName]
	if !ok {
		allowValue := []string{}
		return allowValue
	}

	return allowValue
}

func (this *FastLicenseInfo) GetInts(allowName string) []int {
	allowValue, ok := this.AllowInts[allowName]
	if !ok {
		allowValue := []int{}
		return allowValue
	}

	return allowValue
}
