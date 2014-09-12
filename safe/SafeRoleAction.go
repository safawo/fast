package safe

import (
	"fmt"
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/utils"
	"strings"
)

type QuerySafeRoleAction struct {
	mvc.JsonAction
}

type CreateSafeRoleAction struct {
	mvc.JsonAction
}

type DeleteSafeRoleAction struct {
	mvc.JsonAction
}

type ChangeSafeRoleAction struct {
	mvc.JsonAction
}

type AllocSafeRoleAction struct {
	mvc.JsonAction
}

type QueryAllocSafeRoleAction struct {
	mvc.JsonAction
}

type QuerySafeRoleRequest struct {
	mvc.FastRequestWrap
	RoleId   string `json:"roleId"`
	RoleName string `json:"roleName"`
}

type QuerySafeRoleResponse struct {
	mvc.FastResponseWrap
	Roles []SafeRole `json:"roles"`
}

type CreateSafeRoleRequest struct {
	mvc.FastRequestWrap
	SafeRole
}

type CreateSafeRoleResponse struct {
	mvc.FastResponseWrap
	SafeRole
}

type DeleteSafeRoleRequest struct {
	mvc.FastRequestWrap
	RoleId string `json:"roleId"`
}

type DeleteSafeRoleResponse struct {
	mvc.FastResponseWrap
}

type ChangeSafeRoleRequest struct {
	mvc.FastRequestWrap
	SafeRole
}

type ChangeSafeRoleResponse struct {
	mvc.FastResponseWrap
	SafeRole
}

type AllocSafeRoleRequest struct {
	mvc.FastRequestWrap
	Allocs   []RoleAllocInfo `json:"allocs"`
	UnAllocs []RoleAllocInfo `json:"unAllocs"`
}

type AllocSafeRoleResponse struct {
	mvc.FastResponseWrap
}

type QueryAllocSafeRoleRequest struct {
	mvc.FastRequestWrap
	Allocs []RoleAllocInfo `json:"allocs"`
}

type QueryAllocSafeRoleResponse struct {
	mvc.FastResponseWrap
	Allocs []RoleAllocInfo `json:"allocs"`
}

func (this *QuerySafeRoleAction) Post() {
	reqMsg := &QuerySafeRoleRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	reqMsg.RoleId = strings.TrimSpace(reqMsg.RoleId)
	reqMsg.RoleName = strings.TrimSpace(reqMsg.RoleName)

	querySql := "select distinct * from fast.safeRole"
	if reqMsg.RoleId != comm.NULL_STR {
		querySql += " where id='" + reqMsg.RoleId + "' "
	} else if reqMsg.RoleName != comm.NULL_STR {
		querySql += " where name='" + reqMsg.RoleName + "' "
	} else {
		querySql += " where isUser=false and isDepart=false"
	}

	querySql += " order by name"

	roleRows, roleErr := db.Query(querySql)
	utils.VerifyErr(roleErr)

	rspMsg := &QuerySafeRoleResponse{}
	for roleRows.Next() {
		role := SafeRole{}
		roleRows.Scan(
			&role.Id,
			&role.Name,
			&role.RoleDetail,
			&role.RoleRemark,
			&role.IsUser,
			&role.IsDepart,
			&role.DefaultSafeObject)
		rspMsg.Roles = append(rspMsg.Roles, role)
	}

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)
}

func (this *CreateSafeRoleAction) Post() {
	reqMsg := &CreateSafeRoleRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	addRoleSql := "insert into fast.safeRole values($1,$2,$3,$4,$5,$6,$7)"
	addRoleStmt, err := db.Prepare(addRoleSql)
	utils.VerifyErr(err)

	reqMsg.Id = strings.TrimSpace(reqMsg.Id)

	if reqMsg.Id == comm.NULL_STR {
		reqMsg.Id = utils.RandStr()
	}

	addRoleStmt.Exec(
		reqMsg.Id,
		reqMsg.Name,
		reqMsg.RoleDetail,
		reqMsg.RoleRemark,
		reqMsg.IsUser,
		reqMsg.IsDepart,
		reqMsg.DefaultSafeObject)
	addRoleStmt.Close()

	fmt.Println("addRoleSql:", addRoleSql)

	rspMsg := &CreateSafeRoleResponse{}
	rspMsg.Init(reqMsg)
	rspMsg.SafeRole = reqMsg.SafeRole

	WriteLog(rspMsg, rspMsg.Name, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *DeleteSafeRoleAction) Post() {
	reqMsg := &DeleteSafeRoleRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.RoleId = strings.TrimSpace(reqMsg.RoleId)

	rspMsg := &DeleteSafeRoleResponse{}
	rspMsg.Init(reqMsg)
	if reqMsg.RoleId == comm.NULL_STR {
		this.SendJson(rspMsg)
		return
	}

	db := ds.DB()
	defer db.Close()

	delRoleSql := "delete from fast.safeRole where "
	delRoleSql += "id='" + reqMsg.RoleId + "'"

	delRoleStmt, err := db.Prepare(delRoleSql)
	utils.VerifyErr(err)
	delRoleStmt.Exec()
	delRoleStmt.Close()

	WriteLog(rspMsg, reqMsg.RoleId, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *ChangeSafeRoleAction) Post() {
	reqMsg := &ChangeSafeRoleRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	changeSql := "update fast.safeRole set name=$1,roleDetail=$2,roleRemark=$3," +
		"isUser=$4,isDepart=$5,defaultSafeObject=$6 where id=$7"

	changeStmt, err := db.Prepare(changeSql)
	utils.VerifyErr(err)
	changeStmt.Exec(
		reqMsg.Name,
		reqMsg.RoleDetail,
		reqMsg.RoleRemark,
		reqMsg.IsUser,
		reqMsg.IsDepart,
		reqMsg.DefaultSafeObject,
		reqMsg.Id)
	changeStmt.Close()

	rspMsg := &ChangeSafeRoleResponse{}
	rspMsg.Init(reqMsg)
	rspMsg.SafeRole = reqMsg.SafeRole

	WriteLog(rspMsg, reqMsg.Name, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *AllocSafeRoleAction) Post() {
	reqMsg := &AllocSafeRoleRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	unAllocSql := "delete from fast.roleAlloc where roleId=$1 and userId=$2"
	allocSql := "insert into fast.roleAlloc values($1,$2)"

	unAllocStmt, err := db.Prepare(unAllocSql)
	utils.VerifyErr(err)
	for _, v := range reqMsg.UnAllocs {
		unAllocStmt.Exec(v.RoleId, v.UserId)
	}
	unAllocStmt.Close()

	allocStmt, err := db.Prepare(allocSql)
	utils.VerifyErr(err)
	for _, v := range reqMsg.Allocs {
		allocStmt.Exec(v.RoleId, v.UserId)
	}
	allocStmt.Close()

	rspMsg := &AllocSafeRoleResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)
}

func (this *QueryAllocSafeRoleAction) Post() {
	reqMsg := &QueryAllocSafeRoleRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	querySql := "select distinct * from fast.roleAlloc"

	allocRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	rspMsg := &QueryAllocSafeRoleResponse{}

	for allocRows.Next() {
		alloc := RoleAllocInfo{}
		allocRows.Scan(&alloc.RoleId, &alloc.UserId)
		rspMsg.Allocs = append(rspMsg.Allocs, alloc)
	}

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)
}
