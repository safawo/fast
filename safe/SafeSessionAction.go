package safe

import (
	"os"
	"strings"
	"think/fast/comm"
	"think/fast/msg"
	"think/fast/mvc"
	"think/fast/utils"
)

type LoginAction struct {
	mvc.JsonAction
}

type LogoutAction struct {
	mvc.JsonAction
}

type ShakeHandAction struct {
	mvc.JsonAction
}

type AwakeLoginAction struct {
	mvc.JsonAction
}

type ForceOfflineAction struct {
	mvc.JsonAction
}

type ConsultSessionAction struct {
	mvc.JsonAction
}

type LoginRequest struct {
	mvc.FastRequestWrap
	UserName string `json:"userName"`
	Password string `json:"password"`
	HostName string `json:"hostName"`
	HostIp   string `json:"hostIp"`
}

type LoginResponse struct {
	mvc.FastResponseWrap
	NewSessionId   string `json:"newSessionId"`
	ServerHostName string `json:"serverHostName"`
}

type LogoutRequest struct {
	mvc.FastRequestWrap
}

type LogoutResponse struct {
	mvc.FastResponseWrap
}

type ShakeHandRequest struct {
	mvc.FastRequestWrap
}

type ShakeHandResponse struct {
	mvc.FastResponseWrap
}

type AwakeLoginRequest struct {
	mvc.FastRequestWrap
	Password string `json:"password"`
}

type AwakeLoginResponse struct {
	mvc.FastResponseWrap
}

type ForceOfflineRequest struct {
	mvc.FastRequestWrap
}

type ForceOfflineResponse struct {
	mvc.FastResponseWrap
}

type ConsultSessionRequest struct {
	mvc.FastRequestWrap
}

type ConsultSessionResponse struct {
	mvc.FastResponseWrap
	SessionSeed string `json:"sessionSeed"`
}

func (this *LoginAction) Post() {
	reqMsg := &LoginRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.UserName = strings.TrimSpace(reqMsg.UserName)
	reqMsg.Password = strings.TrimSpace(reqMsg.Password)
	reqMsg.HostIp = strings.TrimSpace(reqMsg.HostIp)
	reqMsg.HostName = strings.TrimSpace(reqMsg.HostName)

	rspMsg := &LoginResponse{}
	rspMsg.Init(reqMsg)
	rspMsg.ServerHostName, _ = os.Hostname()

	if reqMsg.UserName == comm.NULL_STR ||
		reqMsg.Password == comm.NULL_STR ||
		reqMsg.HostName == comm.NULL_STR ||
		reqMsg.HostIp == comm.NULL_STR {
		rspMsg.SetRspRetId(msg.MSG_PARA_ABSENT)
		this.SendJson(rspMsg)
		return
	}

	_, ok := UserMgr().getUserByName(reqMsg.UserName)
	if !ok {
		rspMsg.SetRspRetId(MSG_USER_INVALID)
		this.SendJson(rspMsg)
		return
	}

	session, msgId, ok := SessionMgr().login(this.Ctx.Request, reqMsg.UserName, reqMsg.Password, reqMsg.HostIp, reqMsg.HostName)
	if ok {
		rspMsg.NewSessionId = session.GetId()
		rspMsg.FastResponseWrap.Header.FastRequest.ReqSessionId = session.GetId()
	}

	rspMsg.SetRsp(msgId)

	WriteLog(rspMsg, reqMsg.UserName, comm.NULL_STR)
	this.SendJson(rspMsg)
	return
}

func (this *LogoutAction) Post() {
	reqMsg := &LogoutRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &LogoutResponse{}
	rspMsg.Init(reqMsg)

	SessionMgr().logout(reqMsg.GetReqSessionId())

	this.SendJson(rspMsg)
}

func (this *ShakeHandAction) Post() {
	reqMsg := &ShakeHandRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &ShakeHandResponse{}
	rspMsg.Init(reqMsg)

	sessionId := reqMsg.GetReqSessionId()
	session, ok := SessionMgr().GetSession(sessionId)
	if !ok {
		rspMsg.SetRspRetId(msg.MSG_FAIL)
		this.SendJson(rspMsg)
		return
	}

	if !SessionMgr().IsValid(sessionId) {
		rspMsg.SetRspRetId(msg.MSG_SESSION_INVALID)
		this.SendJson(rspMsg)
		return
	}

	session.KeepAlive()

	this.SendJson(rspMsg)
}

func (this *AwakeLoginAction) Post() {
	reqMsg := &AwakeLoginRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &AwakeLoginResponse{}
	rspMsg.Init(reqMsg)

	reqMsg.Password = strings.TrimSpace(reqMsg.Password)

	if reqMsg.Password == comm.NULL_STR {
		rspMsg.SetRspRetId(msg.MSG_PARA_ABSENT)
		this.SendJson(rspMsg)
		return
	}

	sessionId := reqMsg.GetReqSessionId()
	session, ok := SessionMgr().GetSession(sessionId)
	if !ok {
		rspMsg.SetRspRetId(msg.MSG_FAIL)
		this.SendJson(rspMsg)
		return
	}

	if !SessionMgr().IsValid(sessionId) {
		rspMsg.SetRspRetId(msg.MSG_SESSION_INVALID)
		this.SendJson(rspMsg)
		return
	}

	reqMsg.Password = utils.PasswordBySeed(reqMsg.Password, session.GetUserId())

	user, ok := UserMgr().getUser(session.GetUserId())
	if !ok {
		rspMsg.SetRspRetId(msg.MSG_FAIL)
		this.SendJson(rspMsg)
		return
	}

	if user.Password != reqMsg.Password {
		rspMsg.SetRspRetId(MSG_PASSWORD_ERROR)
		this.SendJson(rspMsg)
		return
	}
	session.KeepAlive()

	this.SendJson(rspMsg)

}

func (this *ForceOfflineAction) Post() {
	reqMsg := &ForceOfflineRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &ForceOfflineResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

}

func (this *ConsultSessionAction) Post() {
	reqMsg := &ConsultSessionRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &ConsultSessionResponse{}
	rspMsg.Init(reqMsg)

	rspMsg.SessionSeed = utils.RandStr() + "sessionId"
	rspMsg.SessionSeed = utils.EnCodeBase64(rspMsg.SessionSeed)
	rspMsg.SessionSeed = utils.Md5(rspMsg.SessionSeed)

	this.SendJson(rspMsg)
}
