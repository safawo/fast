package safe

import (
	"fmt"
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/utils"
	"strings"
)

type QuerySafeOperateAction struct {
	mvc.JsonAction
}

type SafeOperateAuthAction struct {
	mvc.JsonAction
}

type QuerySafeOperateAuthAction struct {
	mvc.JsonAction
}

type QuerySafeOperateRequest struct {
	mvc.FastRequestWrap
	OperateCode  string `json:"operateCode"`
	SubSys       string `json:"subSys"`
	OperateGroup string `json:"operateGroup"`
	OperateName  string `json:"operateName"`
}

type QuerySafeOperateResponse struct {
	mvc.FastResponseWrap
	Operates []SafeOperate `json:"operates"`
}

type SafeOperateAuthRequest struct {
	mvc.FastRequestWrap
	Auths   []OperateAuthInfo `json:"auths"`
	UnAuths []OperateAuthInfo `json:"unAuths"`
}

type SafeOperateAuthResponse struct {
	mvc.FastResponseWrap
}

type QuerySafeOperateAuthRequest struct {
	mvc.FastRequestWrap
}

type QuerySafeOperateAuthResponse struct {
	mvc.FastResponseWrap
	Auths []OperateAuthInfo `json:"auths"`
}

func (this *QuerySafeOperateAction) Post() {

	reqMsg := &QuerySafeOperateRequest{}
	this.GetReqJson(reqMsg)

	reqMsg.OperateCode = strings.TrimSpace(reqMsg.OperateCode)
	reqMsg.SubSys = strings.TrimSpace(reqMsg.SubSys)
	reqMsg.OperateGroup = strings.TrimSpace(reqMsg.OperateGroup)
	reqMsg.OperateName = strings.TrimSpace(reqMsg.OperateName)

	operateSql := "select * from fast.safeOperate where 1=1"
	if reqMsg.OperateCode != comm.NULL_STR {
		operateSql += " and operateCode='" + reqMsg.OperateCode + "'"
	} else if reqMsg.SubSys != comm.NULL_STR {
		operateSql += " and subsys='" + reqMsg.SubSys + "'"

		if reqMsg.OperateGroup != comm.NULL_STR {
			operateSql += " and operateGroup='" + reqMsg.OperateGroup + "'"

			if reqMsg.OperateName != comm.NULL_STR {
				operateSql += " and operateName='" + reqMsg.OperateName + "'"
			}
		}
	}

	fmt.Println("querySql:", operateSql)

	db := ds.DB()
	defer db.Close()

	operateRows, err := db.Query(operateSql)
	utils.VerifyErr(err)

	rspMsg := &QuerySafeOperateResponse{}
	rspMsg.Init(reqMsg)

	for operateRows.Next() {
		operate := SafeOperate{}
		operateRows.Scan(
			&operate.Id,
			&operate.SerialId,
			&operate.OperateCode,
			&operate.Subsys,
			&operate.OperateGroup,
			&operate.OperateName,
			&operate.OperateDetail,
			&operate.OperateRemark,
			&operate.IsAuth,
			&operate.IsLog)
		rspMsg.Operates = append(rspMsg.Operates, operate)
	}

	this.SendJson(rspMsg)
}

func (this *SafeOperateAuthAction) Post() {

	reqMsg := &SafeOperateAuthRequest{}
	this.GetReqJson(reqMsg)

	db := ds.DB()
	defer db.Close()

	unAuthSql := "delete from fast.operateAuth where operateId=$1 and roleId=$2"
	authSql := "insert into fast.operateAuth values($1,$2)"

	unAuthStmt, err := db.Prepare(unAuthSql)
	utils.VerifyErr(err)
	for _, v := range reqMsg.UnAuths {
		unAuthStmt.Exec(v.OperateId, v.RoleId)
	}
	unAuthStmt.Close()

	authStmt, err := db.Prepare(authSql)
	utils.VerifyErr(err)
	for _, v := range reqMsg.Auths {
		authStmt.Exec(v.OperateId, v.RoleId)
	}
	authStmt.Close()

	rspMsg := &SafeOperateAuthResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)

}

func (this *QuerySafeOperateAuthAction) Post() {

	reqMsg := &QuerySafeOperateAuthRequest{}
	this.GetReqJson(reqMsg)

	db := ds.DB()
	defer db.Close()

	querySql := "select distinct * from fast.operateAuth"

	authRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	rspMsg := &QuerySafeOperateAuthResponse{}

	for authRows.Next() {
		auth := OperateAuthInfo{}
		authRows.Scan(&auth.OperateId, &auth.RoleId)
		rspMsg.Auths = append(rspMsg.Auths, auth)
	}

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)
}
