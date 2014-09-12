package safe

import (
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/msg"
	"github.com/safawo/fast/mvc"
	"strings"
)

type CreateSafeObjectAction struct {
	mvc.JsonAction
}

type DeleteSafeObjectAction struct {
	mvc.JsonAction
}

type QuerySafeObjectAction struct {
	mvc.JsonAction
}

type SafeObjectAuthAction struct {
	mvc.JsonAction
}

type QuerySafeObjectAuthAction struct {
	mvc.JsonAction
}

type CreateSafeObjectRequest struct {
	mvc.FastRequestWrap
	SafeObject
}

type CreateSafeObjectResponse struct {
	mvc.FastResponseWrap
	SafeObject
}

type DeleteSafeObjectRequest struct {
	mvc.FastRequestWrap
	ObjectId   string `json:"objectId"`
	ObjectName string `json:"objectName"`
}

type DeleteSafeObjectResponse struct {
	mvc.FastResponseWrap
}

type QuerySafeObjectRequest struct {
	mvc.FastRequestWrap
	ObjectType string `json:"objectType"`
}

type QuerySafeObjectResponse struct {
	mvc.FastResponseWrap
	Objects []SafeObject `json:"objects"`
}

type SafeObjectAuthRequest struct {
	mvc.FastRequestWrap
	Auths   []ObjectAuthInfo `json:"auths"`
	UnAuths []ObjectAuthInfo `json:"unAuths"`
}

type SafeObjectAuthResponse struct {
	mvc.FastResponseWrap
}

type QuerySafeObjectAuthRequest struct {
	mvc.FastRequestWrap
	ObjectType string `json:"objectType"`
}

type QuerySafeObjectAuthResponse struct {
	mvc.FastResponseWrap
	Auths []ObjectAuthInfo `json:"auths"`
}

func (this *CreateSafeObjectAction) Post() {
	reqMsg := &CreateSafeObjectRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &CreateSafeObjectResponse{}
	rspMsg.Init(reqMsg)

	msgId, ok := SafeObjectMgr().CreateObject(&reqMsg.SafeObject)
	if !ok {
		rspMsg.SetRsp(msgId)
	} else {
		rspMsg.SafeObject = reqMsg.SafeObject
	}

	this.SendJson(rspMsg)
}

func (this *DeleteSafeObjectAction) Post() {
	reqMsg := &DeleteSafeObjectRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.ObjectId = strings.TrimSpace(reqMsg.ObjectId)
	reqMsg.ObjectName = strings.TrimSpace(reqMsg.ObjectName)

	rspMsg := &DeleteSafeObjectResponse{}
	rspMsg.Init(reqMsg)

	if reqMsg.ObjectId != comm.NULL_STR {
		SafeObjectMgr().DelObject(reqMsg.ObjectId)
	} else if reqMsg.ObjectName != comm.NULL_STR {
		SafeObjectMgr().DelObjectByName(reqMsg.ObjectName)
	} else {
		rspMsg.SetRsp(msg.MSG_PARA_ABSENT)
		this.SendJson(rspMsg)
		return
	}

	this.SendJson(rspMsg)
}

func (this *QuerySafeObjectAction) Post() {
	reqMsg := &QuerySafeObjectRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.ObjectType = strings.TrimSpace(reqMsg.ObjectType)

	rspMsg := &QuerySafeObjectResponse{}
	if reqMsg.ObjectType == comm.NULL_STR {
		rspMsg.Objects = SafeObjectMgr().QueryObjects()
	} else {
		rspMsg.Objects = SafeObjectMgr().QueryObjectsByType(reqMsg.ObjectType)
	}

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)
}

func (this *SafeObjectAuthAction) Post() {
	reqMsg := &SafeObjectAuthRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	SafeObjectMgr().Auth(reqMsg.Auths)
	SafeObjectMgr().UnAuth(reqMsg.UnAuths)

	rspMsg := &SafeObjectAuthResponse{}
	rspMsg.Init(reqMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *QuerySafeObjectAuthAction) Post() {
	reqMsg := &QuerySafeObjectAuthRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.ObjectType = strings.TrimSpace(reqMsg.ObjectType)

	rspMsg := &QuerySafeObjectAuthResponse{}
	if reqMsg.ObjectType == comm.NULL_STR {
		rspMsg.Auths = SafeObjectMgr().QueryAuths()
	} else {
		rspMsg.Auths = SafeObjectMgr().QueryAuthsByType(reqMsg.ObjectType)
	}

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)
}
