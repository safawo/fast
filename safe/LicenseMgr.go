package safe

import (
	"encoding/json"
	"fmt"
	"strings"
	"think/fast/comm"
	"think/fast/utils"
)

type LicenseInterface interface {
	IsEmpty() bool

	GetDistributor() string
	GetDistributeId() string

	GetLicenseId() string
	GetLicenseType() string

	GetCustomerId() string
	GetCustomerName() string

	GetBindNetCard() string
	GetBindDisk() string
	GetBindCPU() string

	GetExpiryDate() string
	GetRemark() string

	GetAllowStr(allowName string) string
	GetAllowInt(allowName string) int

	GetAllowStrs(allowName string) []string
	GetAllowInts(allowName string) []int
}

type LicenseMgrInterface interface {
	QueryLicense() *FastLicenseInfo
	ImportLicense(newLicense string)
	ExportLicense() string
	ClearLicense()
}

type licenseMgr struct {
	licenseData *FastLicenseInfo
}

var licenseMgrInstance *licenseMgr

func (this *licenseMgr) init() {
	this.licenseData = &FastLicenseInfo{}
	this.licenseData.Reset()

	if !utils.SysParaMgr().ExistPara(LICENSE_PARA_NAME) {
		return
	}

	existLicense := utils.SysParaMgr().GetPara(LICENSE_PARA_NAME)
	existLicense = strings.TrimSpace(existLicense)
	if existLicense == comm.NULL_STR {
		return
	}

	existLicense = utils.DeCodeBase64(existLicense)
	existLicense = utils.DeCodeBase64(existLicense)
	existLicense = utils.DeCodeBase64(existLicense)

	json.Unmarshal([]byte(existLicense), this.licenseData)

}

func LicenseMgr() *licenseMgr {

	fmt.Println("  Init License Mgr")

	if licenseMgrInstance == nil {
		licenseMgrInstance = &licenseMgr{}
		licenseMgrInstance.init()
	}

	return licenseMgrInstance
}

func (this *licenseMgr) QueryLicense() *FastLicenseInfo {
	return this.licenseData
}

func (this *licenseMgr) chenckValid(newLicense string) bool {
	checkLicData := strings.TrimSpace(newLicense)

	checkLicData = utils.DeCodeBase64(checkLicData)
	checkLicData = utils.DeCodeBase64(checkLicData)
	checkLicData = utils.DeCodeBase64(checkLicData)

	checkLic := &FastLicenseInfo{}
	err := json.Unmarshal([]byte(checkLicData), checkLic)

	if err != nil {
		return false
	}

	if this.GetLicenseId() != checkLic.LicenseId {
		return false
	}
	if this.GetLicenseType() != checkLic.LicenseType {
		return false
	}
	if this.GetCustomerId() != checkLic.CustomerId {
		return false
	}
	if this.GetCustomerName() != checkLic.CustomerName {
		return false
	}

	return true
}

func (this *licenseMgr) ImportLicense(newLicense string) bool {
	newLicenseData := strings.TrimSpace(newLicense)

	newLicenseData = utils.DeCodeBase64(newLicenseData)
	newLicenseData = utils.DeCodeBase64(newLicenseData)
	newLicenseData = utils.DeCodeBase64(newLicenseData)

	newLicenseObj := &FastLicenseInfo{}
	err := json.Unmarshal([]byte(newLicenseData), newLicenseObj)

	if err != nil {
		return false
	}

	this.licenseData = newLicenseObj
	utils.SysParaMgr().ChangePara(LICENSE_PARA_NAME, newLicense)

	return true
}

func (this *licenseMgr) ExportLicense() string {
	if !utils.SysParaMgr().ExistPara(LICENSE_PARA_NAME) {
		return comm.NULL_STR
	}

	license := utils.SysParaMgr().GetPara(LICENSE_PARA_NAME)
	license = strings.TrimSpace(license)

	return license
}

func (this *licenseMgr) ClearLicense() {
	if !utils.SysParaMgr().ExistPara(LICENSE_PARA_NAME) {
		return
	}

	emptyLicense := &FastLicenseInfo{}
	emptyLicense.Reset()

	byteEmptyLicense, err := json.MarshalIndent(emptyLicense, "", "  ")
	utils.VerifyErr(err)

	emptyLicenseStr := utils.EnCodeBase64(string(byteEmptyLicense))
	emptyLicenseStr = utils.EnCodeBase64(emptyLicenseStr)
	emptyLicenseStr = utils.EnCodeBase64(emptyLicenseStr)

	utils.SysParaMgr().ChangePara(LICENSE_PARA_NAME, emptyLicenseStr)

	return
}

func (this *licenseMgr) IsEmpty() bool {
	return this.QueryLicense().IsEmpty()
}

func (this *licenseMgr) GetDistributor() string {
	return this.QueryLicense().Distributor
}

func (this *licenseMgr) GetDistributeId() string {
	return this.QueryLicense().DistributeId
}

func (this *licenseMgr) GetLicenseId() string {
	return this.QueryLicense().LicenseId
}

func (this *licenseMgr) GetLicenseType() string {
	return this.QueryLicense().LicenseType
}

func (this *licenseMgr) GetCustomerId() string {
	return this.QueryLicense().CustomerId
}

func (this *licenseMgr) GetCustomerName() string {
	return this.QueryLicense().CustomerName
}

func (this *licenseMgr) GetBindNetCard() string {
	return this.QueryLicense().BindNetCard
}

func (this *licenseMgr) GetBindDisk() string {
	return this.QueryLicense().BindDisk
}

func (this *licenseMgr) GetBindCPU() string {
	return this.QueryLicense().BindCPU
}

func (this *licenseMgr) GetExpiryDate() string {
	return this.QueryLicense().ExpiryDate
}

func (this *licenseMgr) GetRemark() string {
	return this.QueryLicense().Remark
}

func (this *licenseMgr) GetAllowStr(allowName string) string {
	return this.QueryLicense().GetStr(allowName)
}

func (this *licenseMgr) GetAllowInt(allowName string) int {
	return this.QueryLicense().GetInt(allowName)
}

func (this *licenseMgr) GetAllowStrs(allowName string) []string {
	return this.QueryLicense().GetStrs(allowName)
}

func (this *licenseMgr) GetAllowInts(allowName string) []int {
	return this.QueryLicense().GetInts(allowName)
}
