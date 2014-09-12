package fast

import (
	"think/fast/msg"
	"think/fast/mvc"
)

type LoadMsgListAction struct {
	mvc.JsonAction
}

type LoadMsgListRequest struct {
	mvc.FastRequestWrap
}

type LoadMsgListResponse struct {
	mvc.FastResponseWrap
	msgs []msg.FastMsg `json:"msgs"`
}

func (this *LoadMsgListAction) Post() {
	reqMsg := &LoadMsgListRequest{}
	this.GetReqJson(reqMsg)

	rspMsg := &LoadMsgListResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)
}
