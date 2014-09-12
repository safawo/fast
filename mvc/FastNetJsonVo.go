package mvc

import (
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/msg"
	"github.com/safawo/fast/utils"
)

type FastRequestInterface interface {
	GetReqSessionId() (sessionId string)
	GetReqActionId() (actionId string)
	GetReqStartTime() (startTime string)
}

type FastResponseInterface interface {
	Init(reqest FastRequestInterface)

	FastRequestInterface

	IsSuccess() (success bool)
	GetRspFinishTime() (finishTime string)
	GetRspRetLevel() (retLevel string)

	GetRspRetId() (retId string)
	GetRspRetMsg() (retMsg string)

	GetRspRetMsgDetail() (retMsgDetail string)

	SetRspRetId(retId string)
	SetRsp(retId string)
	SetRspRetMsgDetail(retMsgDetail string)
	SetBeforeSend()

	AppendMsgDetail(appendMsg string)
}

type FastRequest struct {
	ReqSessionId string `json:"reqSessionId"`
	ReqActionId  string `json:"reqActionId"`
	ReqStartTime string `json:"reqStartTime"`
}

type FastRequestWrap struct {
	Header FastRequest `json:"header"`
}

type FastResponse struct {
	FastRequest

	RspFinishTime string `json:"rspFinishTime"`
	RspRetLevel   string `json:"rspRetLevel"`

	RspRetId  string `json:"rspRetId"`
	RspRetMsg string `json:"rspRetMsg"`

	RspRetMsgDetail string `json:"rspRetMsgDetail"`
}

type FastResponseWrap struct {
	Header FastResponse `json:"header"`
}

func (this *FastRequest) GetReqSessionId() (sessionId string) {
	return this.ReqSessionId
}

func (this *FastRequest) GetReqActionId() (actionId string) {
	return this.ReqActionId
}

func (this *FastRequest) GetReqStartTime() (startTime string) {
	return this.ReqStartTime
}

func (this *FastRequestWrap) GetReqSessionId() (sessionId string) {
	return this.Header.GetReqSessionId()
}

func (this *FastRequestWrap) GetReqActionId() (actionId string) {
	return this.Header.GetReqActionId()
}

func (this *FastRequestWrap) GetReqStartTime() (startTime string) {
	return this.Header.GetReqStartTime()
}

func (this *FastResponse) Init(reqest FastRequestInterface) {
	this.ReqSessionId = reqest.GetReqSessionId()
	this.ReqActionId = reqest.GetReqActionId()
	this.ReqStartTime = reqest.GetReqStartTime()

	this.RspFinishTime = utils.StrTime()

	if this.GetRspRetId() == comm.NULL_STR {
		this.RspRetId = msg.MSG_SUCCESS
	}

	if this.GetRspRetLevel() == comm.NULL_STR {
		this.RspRetLevel = msg.MSG_LEVEL_INFO
	}

}

func (this *FastResponse) GetReqSessionId() (sessionId string) {
	return this.ReqSessionId
}

func (this *FastResponse) GetReqActionId() (actionId string) {
	return this.ReqActionId
}

func (this *FastResponse) GetReqStartTime() (startTime string) {
	return this.ReqStartTime
}

func (this *FastResponse) IsSuccess() (success bool) {
	if this.GetRspRetLevel() == msg.MSG_LEVEL_INFO {
		return true
	}

	return false
}

func (this *FastResponse) GetRspFinishTime() (finishTime string) {
	return this.RspFinishTime
}

func (this *FastResponse) GetRspRetLevel() (retLevel string) {
	return this.RspRetLevel
}

func (this *FastResponse) GetRspRetId() (retId string) {
	return this.RspRetId
}

func (this *FastResponse) GetRspRetMsg() (retMsg string) {
	return this.RspRetMsg
}

func (this *FastResponse) GetRspRetMsgDetail() (retMsgDetail string) {
	return this.RspRetMsgDetail
}

func (this *FastResponse) SetRspRetId(retId string) {
	this.RspRetId = retId
	this.RspFinishTime = utils.StrTime()

	msgInfo, ok := msg.Item(retId)

	if !ok {
		this.RspRetLevel = msg.MSG_LEVEL_INFO
		this.RspRetMsg = retId
		return
	}

	this.RspRetLevel = msgInfo.Level
	this.RspRetMsg = msgInfo.MsgInfo

	return
}

func (this *FastResponse) SetRsp(retId string) {
	this.SetRspRetId(retId)
}

func (this *FastResponse) SetRspRetMsgDetail(retMsgDetail string) {
	this.RspRetMsgDetail = retMsgDetail
}

func (this *FastResponse) SetBeforeSend() {
	this.RspFinishTime = utils.StrTime()
}

func (this *FastResponse) AppendMsgDetail(appendMsg string) {
	this.RspRetMsgDetail = this.RspRetMsgDetail + appendMsg
}

func (this *FastResponseWrap) Init(reqest FastRequestInterface) {
	this.Header.Init(reqest)
}

func (this *FastResponseWrap) GetReqSessionId() (sessionId string) {
	return this.Header.GetReqSessionId()
}

func (this *FastResponseWrap) GetReqActionId() (actionId string) {
	return this.Header.GetReqActionId()
}

func (this *FastResponseWrap) GetReqStartTime() (startTime string) {
	return this.Header.GetReqStartTime()
}

func (this *FastResponseWrap) IsSuccess() (success bool) {
	if this.GetRspRetLevel() == msg.MSG_LEVEL_INFO {
		return true
	}

	return false
}

func (this *FastResponseWrap) GetRspFinishTime() (finishTime string) {
	return this.Header.GetRspFinishTime()
}

func (this *FastResponseWrap) GetRspRetLevel() (retLevel string) {
	return this.Header.GetRspRetLevel()
}

func (this *FastResponseWrap) GetRspRetId() (retId string) {
	return this.Header.GetRspRetId()
}

func (this *FastResponseWrap) GetRspRetMsg() (retMsg string) {
	return this.Header.GetRspRetMsg()
}

func (this *FastResponseWrap) GetRspRetMsgDetail() (retMsgDetail string) {
	return this.Header.GetRspRetMsgDetail()
}

func (this *FastResponseWrap) SetRspRetId(retId string) {
	this.Header.SetRspRetId(retId)
}

func (this *FastResponseWrap) SetRspRetMsgDetail(retMsgDetail string) {
	this.Header.SetRspRetMsgDetail(retMsgDetail)
}

func (this *FastResponseWrap) SetRsp(retId string) {
	this.SetRspRetId(retId)
}

func (this *FastResponseWrap) SetBeforeSend() {
	this.Header.SetBeforeSend()
}

func (this *FastResponseWrap) AppendMsgDetail(appendMsg string) {
	this.Header.AppendMsgDetail(appendMsg)
}
