package safe

import (
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/mvc"
	"strings"
)

const (
	LICENSE_PARA_NAME = "licenseData"
)

type QueryLicenseAction struct {
	mvc.JsonAction
}

type VerifyLicenseAction struct {
	mvc.JsonAction
}

type MatchLicenseAction struct {
	mvc.JsonAction
}

type ImportLicenseAction struct {
	mvc.JsonAction
}

type ExportLicenseAction struct {
	mvc.JsonAction
}

type ClearLicenseAction struct {
	mvc.JsonAction
}

type QueryLicenseRequest struct {
	mvc.FastRequestWrap
}

type QueryLicenseResponse struct {
	mvc.FastResponseWrap
	FastLicenseInfo
}

type VerifyLicenseRequest struct {
	mvc.FastRequestWrap
	License string `json:"license"`
}

type VerifyLicenseResponse struct {
	mvc.FastResponseWrap
	IsValid bool `json:"isValid"`
}

type MatchLicenseRequest struct {
	mvc.FastRequestWrap
	License string `json:"license"`
}

type MatchLicenseResponse struct {
	mvc.FastResponseWrap
	IsMatch bool `json:"isMatch"`
}

type ImportLicenseRequest struct {
	mvc.FastRequestWrap
	License string `json:"license"`
}

type ImportLicenseResponse struct {
	mvc.FastResponseWrap
}

type ExportLicenseRequest struct {
	mvc.FastRequestWrap
}

type ExportLicenseResponse struct {
	mvc.FastResponseWrap
	License string `json:"license"`
}

type ClearLicenseRequest struct {
	mvc.FastRequestWrap
}

type ClearLicenseResponse struct {
	mvc.FastResponseWrap
}

func (this *QueryLicenseAction) Post() {
	reqMsg := &QueryLicenseRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &QueryLicenseResponse{}
	rspMsg.FastResponseWrap.Init(reqMsg)
	rspMsg.FastLicenseInfo = *LicenseMgr().QueryLicense()

	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}

func (this *VerifyLicenseAction) Post() {
	reqMsg := &VerifyLicenseRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &VerifyLicenseResponse{}
	rspMsg.Init(reqMsg)

	rspMsg.IsValid = LicenseMgr().chenckValid(reqMsg.License)

	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}

func (this *MatchLicenseAction) Post() {
	reqMsg := &MatchLicenseRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &MatchLicenseResponse{}
	rspMsg.Init(reqMsg)

	licData := strings.TrimSpace(LicenseMgr().ExportLicense())
	reqMsg.License = strings.TrimSpace(reqMsg.License)
	if licData == reqMsg.License {
		rspMsg.IsMatch = true
	} else {
		rspMsg.IsMatch = false
	}

	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}

func (this *ImportLicenseAction) Post() {
	reqMsg := &ImportLicenseRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.License = strings.TrimSpace(reqMsg.License)

	importRet := LicenseMgr().ImportLicense(reqMsg.License)

	rspMsg := &ImportLicenseResponse{}
	rspMsg.Init(reqMsg)
	if !importRet {
		rspMsg.SetRsp(MSG_LICENSE_INVALID)
	}
	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}

func (this *ExportLicenseAction) Post() {
	reqMsg := &ExportLicenseRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &ExportLicenseResponse{}
	rspMsg.Init(reqMsg)
	rspMsg.License = comm.NULL_STR
	rspMsg.License = LicenseMgr().ExportLicense()
	if rspMsg.License == comm.NULL_STR {
		rspMsg.SetRsp(MSG_LICENSE_VOID)
	}

	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return

}
func (this *ClearLicenseAction) Post() {
	reqMsg := &ClearLicenseRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	LicenseMgr().ClearLicense()

	rspMsg := &ClearLicenseResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return

}
