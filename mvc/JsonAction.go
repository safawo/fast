package mvc

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"strconv"
	"think/fast/msg"
	"think/fast/utils"
)

func Router(path string, action JsonActionInterface) {
	beego.Router(path, action)
}

type JsonActionInterface interface {
	beego.ControllerInterface
	GetReqJson(reqMsg interface{}) (ok bool)
	SendJson(rspMsg interface{}) (ok bool)
}

type SessionSafeValidater interface {
	Validate(reqData FastRequestInterface) (ok bool)
}

var (
	sessionValidaterImpl SessionSafeValidater
)

func getSessionValidater() (sessionValidater SessionSafeValidater) {
	return sessionValidaterImpl
}

func ProvideSessionValidater(sessionValidater SessionSafeValidater) {
	sessionValidaterImpl = sessionValidater
}

type JsonAction struct {
	beego.Controller
	reqInterface FastRequestInterface
}

func (this *JsonAction) GetReqInfo() FastRequestInterface {
	return this.reqInterface
}

func (this *JsonAction) Post() {
	http.Error(this.Ctx.ResponseWriter, "Json Request Not Allowed", 405)
	fmt.Println("Json Request Not Allowed Code: 405")
}

func (this *JsonAction) GetReqJson(reqMsg interface{}) (ok bool) {
	ok = true

	reqByte, err := ioutil.ReadAll(this.Ctx.Request.Body)
	utils.VerifyErr(err)

	fmt.Println("Fast Action Request Msg, client:", this.Ctx.Request.RemoteAddr,
		",server:", this.Ctx.Request.Host)
	fmt.Println(string(reqByte))

	err = json.Unmarshal(reqByte, reqMsg)
	utils.VerifyErr(err)

	reqInterfaceTry, isReq := reqMsg.(FastRequestInterface)
	if !isReq {
		return
	}
	this.reqInterface = reqInterfaceTry

	if this.GetReqInfo().GetReqSessionId() == "init" {
		return
	}

	if getSessionValidater() == nil {
		return
	}

	if getSessionValidater().Validate(this.GetReqInfo()) {
		return
	}

	ok = false

	vaidateRsp := &FastResponseWrap{}
	if this.GetReqInfo() != nil {
		vaidateRsp.Init(this.GetReqInfo())
	}

	vaidateRsp.SetRsp(msg.MSG_SESSION_INVALID)
	this.SendJson(vaidateRsp)

	return
}

func (this *JsonAction) SendJson(rspMsg interface{}) (ok bool) {
	ok = true

	rspInterface, isRsp := rspMsg.(FastResponseInterface)
	if !isRsp {
		return
	}
	rspInterface.Init(this.GetReqInfo())
	rspInterface.SetBeforeSend()

	this.Data["json"] = rspMsg
	this.ServeJson()

	return
}

func (this *JsonAction) SendMsg(msgId string) (ok bool) {
	ok = true

	msgIdRsp := &FastResponseWrap{}
	if this.GetReqInfo() != nil {
		msgIdRsp.Init(this.GetReqInfo())
	}

	msgIdRsp.SetRsp(msgId)
	this.SendJson(msgIdRsp)

	return
}

func (this *JsonAction) ServeJson() {
	content, err := json.MarshalIndent(this.Data["json"], "", "  ")

	if err != nil {
		http.Error(this.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	this.Ctx.SetHeader("Content-Length", strconv.Itoa(len(content)), true)
	this.Ctx.ContentType("json")

	this.Ctx.ResponseWriter.Write(content)

	fmt.Println("Fast Action Response Msg, client:", this.Ctx.Request.RemoteAddr,
		",server:", this.Ctx.Request.Host)
	fmt.Println(string(content))
}
