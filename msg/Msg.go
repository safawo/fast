package msg

import (
	"fmt"
)

const (
	MSG_LEVEL_INFO  = "level_info"
	MSG_LEVEL_WARN  = "level_warn"
	MSG_LEVEL_ERROR = "level_error"
	MSG_LEVEL_FATAL = "level_fatal"
)

const (
	MSG_SUCCESS = "id_success"
	MSG_FAIL    = "id_fail"
	MSG_ERROR   = "id_error"
	MSG_FATAL   = "id_fatal"
)

const (
	MSG_SUCCESS_PART    = "success_part"
	MSG_UNKNOWN_ERROR   = "unknown_error"
	MSG_UNKNOWN_OPERATE = "unknown_operate"

	MSG_FORBID_NOAUTH = "forbid_noauth"

	MSG_BATCH_FORBID = "batch_forbid"
	MSG_BATCH_OVER   = "batch_over"

	MSG_SESSION_INVALID = "session_invalid"
	MSG_SESSION_TIMEOUT = "session_timeout"

	MSG_PARA_ABSENT = "para_absent"
	MSG_OBJ_EXIST   = "obj_exist"
	MSG_OBJ_NOEXIST = "obj_noexist"

	MSG_LICENSE_INVALID = "license_invalid"
	MSG_LICENSE_VOID    = "license_void"

	MSG_LICENSE_NOENOUGH = "license_noenough"
	MSG_LICENSE_Expired  = "license_expired"
)

type FastMsg struct {
	Level   string `json:"level"`
	MsgId   string `json:"msgId"`
	MsgInfo string `json:"msgInfo"`
}

type MsgBuffInfo struct {
	MapMsg map[string](*FastMsg)
}

var msgMgr = MsgBuffInfo{}

func (this *MsgBuffInfo) init() {

	fmt.Println("  Init Msg Mgr")

	this.MapMsg = map[string](*FastMsg){}

	this.regPublicMsg()
}

func Item(msgId string) (msg *FastMsg, ok bool) {
	msg, ok = msgMgr.MapMsg[msgId]
	return
}

func Reg(level string, msgId string, msgInfo string) {
	regMsg := &FastMsg{level, msgId, msgInfo}
	msgMgr.MapMsg[msgId] = regMsg
}

func RegInfo(msgId string, msgInfo string) {
	Reg(MSG_LEVEL_INFO, msgId, msgInfo)
}

func RegWarn(msgId string, msgInfo string) {
	Reg(MSG_LEVEL_WARN, msgId, msgInfo)
}

func RegError(msgId string, msgInfo string) {
	Reg(MSG_LEVEL_ERROR, msgId, msgInfo)
}

func RegFatal(msgId string, msgInfo string) {
	Reg(MSG_LEVEL_FATAL, msgId, msgInfo)
}

func (this *MsgBuffInfo) regPublicMsg() {
	RegInfo(MSG_SUCCESS, "Operate Success")
	RegWarn(MSG_FAIL, "Operate Fail")
	RegError(MSG_ERROR, "Operate Error")
	RegFatal(MSG_FATAL, "Operate Fatal")

	RegError(MSG_SUCCESS_PART, "success part")
	RegError(MSG_UNKNOWN_ERROR, "unknown error")
	RegError(MSG_UNKNOWN_OPERATE, "unknown operate")

	RegError(MSG_FORBID_NOAUTH, "forbid noauth")
	RegError(MSG_BATCH_FORBID, "batch forbid")
	RegError(MSG_BATCH_OVER, "batch over")

	RegError(MSG_SESSION_INVALID, "Session invalid, Please Re Login")
	RegError(MSG_SESSION_TIMEOUT, "session timeout")

	RegError(MSG_PARA_ABSENT, "para absent")
	RegError(MSG_OBJ_EXIST, "obj exist")
	RegError(MSG_OBJ_NOEXIST, "obj noexist")

	RegError(MSG_LICENSE_INVALID, "license invalid")
	RegError(MSG_LICENSE_VOID, "license void")
	RegError(MSG_LICENSE_NOENOUGH, "License not enough")
	RegError(MSG_LICENSE_Expired, "License expired")
}
