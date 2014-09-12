package safe

import (
	"strings"
	"think/fast/comm"
	"think/fast/ds"
	"think/fast/msg"
	"think/fast/mvc"
	"think/fast/utils"
)

type SelfQueryMyInfoAction struct {
	mvc.JsonAction
}

type SelfModPasswordAction struct {
	mvc.JsonAction
}

type SelfModUserInfoAction struct {
	mvc.JsonAction
}

type SelfQueryMyInfoRequest struct {
	mvc.FastRequestWrap
	SafeUser
}

type SelfQueryMyInfoResponse struct {
	mvc.FastResponseWrap
	SafeUser
}

type SelfModPasswordRequest struct {
	mvc.FastRequestWrap
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

type SelfModPasswordResponse struct {
	mvc.FastResponseWrap
}

type SelfModUserInfoRequest struct {
	mvc.FastRequestWrap
	SafeUser
}

type SelfModUserInfoResponse struct {
	mvc.FastResponseWrap
	SafeUser
}

func (this *SelfQueryMyInfoAction) Post() {
	reqMsg := &SelfQueryMyInfoRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &SelfQueryMyInfoResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)
}

func (this *SelfModPasswordAction) Post() {
	reqMsg := &SelfModPasswordRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &SelfModPasswordResponse{}
	rspMsg.Init(reqMsg)

	reqMsg.CurrentPassword = strings.TrimSpace(reqMsg.CurrentPassword)
	reqMsg.NewPassword = strings.TrimSpace(reqMsg.NewPassword)

	if reqMsg.CurrentPassword == comm.NULL_STR ||
		reqMsg.NewPassword == comm.NULL_STR {
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

	reqMsg.CurrentPassword = utils.PasswordBySeed(reqMsg.CurrentPassword, session.GetUserId())
	reqMsg.NewPassword = utils.PasswordBySeed(reqMsg.NewPassword, session.GetUserId())

	user, ok := UserMgr().getUser(session.GetUserId())
	if !ok {
		rspMsg.SetRspRetId(msg.MSG_FAIL)
		this.SendJson(rspMsg)
		return
	}

	if user.Password != reqMsg.CurrentPassword {
		rspMsg.SetRspRetId(MSG_PASSWORD_ERROR)
		this.SendJson(rspMsg)
		return
	}

	user.Password = reqMsg.NewPassword
	msgId, ok := UserMgr().changeUser(user)

	rspMsg.SetRspRetId(msgId)
	WriteLog(rspMsg, session.GetUserName(), comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *SelfModUserInfoAction) Post() {
	reqMsg := &SelfModUserInfoRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	selfModSql := "update fast.safeUser set nickName=$1,firstName=$2,lastName=$3,"
	selfModSql += "mobile=$4,email=$5 where id=$6"
	selfModStmt, err := db.Prepare(selfModSql)
	utils.VerifyErr(err)
	selfModStmt.Exec(reqMsg.NickName,
		reqMsg.FirstName,
		reqMsg.LastName,
		reqMsg.Mobile,
		reqMsg.Email,
		reqMsg.Id)
	selfModStmt.Close()

	rspMsg := &SelfModUserInfoResponse{}
	rspMsg.SafeUser = reqMsg.SafeUser
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

}
