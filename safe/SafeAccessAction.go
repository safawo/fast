package safe

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/msg"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/utils"
	"strings"
)

type QueryMyAccessObjectAction struct {
	mvc.JsonAction
}

type UpLoadAction struct {
	beego.Controller
}

type QueryMyAccessObjectRequest struct {
	mvc.FastRequestWrap
	ObjectType string `json:"objectType"`
}

type QueryMyAccessObjectResponse struct {
	mvc.FastResponseWrap
	ObjectIds   []string `json:"objectIds"`
	ObjectNames []string `json:"objectNames"`
}

func (this *QueryMyAccessObjectAction) Post() {
	reqMsg := &QueryMyAccessObjectRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.ObjectType = strings.TrimSpace(reqMsg.ObjectType)

	rspMsg := &QueryMyAccessObjectResponse{}
	rspMsg.Init(reqMsg)

	sessionId := reqMsg.GetReqSessionId()
	session, ok := SessionMgr().GetSession(sessionId)
	if !ok {
		rspMsg.SetRspRetId(msg.MSG_FAIL)
		this.SendJson(rspMsg)
		return
	}

	userId := session.GetUserId()

	db := ds.DB()
	defer db.Close()

	querySql := comm.NULL_STR
	if !session.IsAdmin() {
		querySql += "select distinct obj.id,obj.objectName from fast.safeObject obj, fast.objectAuth auth where "
		querySql += "obj.id = auth.objectId "
		if reqMsg.ObjectType != comm.NULL_STR {
			querySql += "and obj.objectType='" + reqMsg.ObjectType + "' "
		}

		querySql += "and (auth.roleId='" + userId + "' or auth.roleId in(select distinct alloc.roleId "
		querySql += "from fast.roleAlloc alloc where alloc.userId='" + userId + "'))"
	} else {
		querySql += "select distinct obj.id,obj.objectName from fast.safeObject obj where 1=1 "
		if reqMsg.ObjectType != comm.NULL_STR {
			querySql += "and obj.objectType='" + reqMsg.ObjectType + "' "
		}
	}

	fmt.Println("query access object sql:", querySql)

	objRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	objectId := comm.NULL_STR
	objectName := comm.NULL_STR
	for objRows.Next() {
		objRows.Scan(&objectId, &objectName)
		rspMsg.ObjectIds = append(rspMsg.ObjectIds, objectId)
		rspMsg.ObjectNames = append(rspMsg.ObjectNames, objectName)
	}

	this.SendJson(rspMsg)
}
