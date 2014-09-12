package backup

import (
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/safe"
)

type NotifyBackupAction struct {
	mvc.JsonAction
}

type NotifyRestoreAction struct {
	mvc.JsonAction
}

type BackupSystemAction struct {
	mvc.JsonAction
}

type RestoreSystemAction struct {
	mvc.JsonAction
}

type QueryBackupAction struct {
	mvc.JsonAction
}

type DeleteBackupAction struct {
	mvc.JsonAction
}

type NotifyBackupRequest struct {
	mvc.FastRequestWrap
	SystemBackupInfo
}

type NotifyBackupResponse struct {
	mvc.FastResponseWrap
}

type NotifyRestoreRequest struct {
	mvc.FastRequestWrap
	SystemBackupInfo
}

type NotifyRestoreResponse struct {
	mvc.FastResponseWrap
}

type BackupSystemRequest struct {
	mvc.FastRequestWrap
}

type BackupSystemResponse struct {
	mvc.FastResponseWrap
	SystemBackupInfo
}

type RestoreSystemRequest struct {
	mvc.FastRequestWrap
	BackupId string `json:"backupId"`
}

type RestoreSystemResponse struct {
	mvc.FastResponseWrap
}

type QueryBackupRequest struct {
	mvc.FastRequestWrap
}

type QueryBackupResponse struct {
	mvc.FastResponseWrap
	Backups []SystemBackupInfo `json:"backups"`
}

type DeleteBackupRequest struct {
	mvc.FastRequestWrap
	BackupId string `json:"backupId"`
}

type DeleteBackupResponse struct {
	mvc.FastResponseWrap
}

func (this *NotifyBackupAction) Post() {
	reqMsg := &NotifyBackupRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &NotifyBackupResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	safe.WriteLog(rspMsg, reqMsg.BackupId, comm.NULL_STR)

	return
}

func (this *NotifyRestoreAction) Post() {
	reqMsg := &NotifyRestoreRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &NotifyRestoreResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	safe.WriteLog(rspMsg, reqMsg.BackupId, comm.NULL_STR)

	return
}

func (this *BackupSystemAction) Post() {
	reqMsg := &BackupSystemRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &BackupSystemResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	safe.WriteLog(rspMsg, rspMsg.BackupId, comm.NULL_STR)

	return
}

func (this *RestoreSystemAction) Post() {
	reqMsg := &RestoreSystemRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &RestoreSystemResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	safe.WriteLog(rspMsg, reqMsg.BackupId, comm.NULL_STR)

	return
}

func (this *QueryBackupAction) Post() {
	reqMsg := &QueryBackupRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &QueryBackupResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	safe.WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}

func (this *DeleteBackupAction) Post() {
	reqMsg := &DeleteBackupRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &DeleteBackupResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

	safe.WriteLog(rspMsg, reqMsg.BackupId, comm.NULL_STR)

	return
}
