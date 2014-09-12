package safe

import (
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/utils"
	"strings"
)

type QueryUserAction struct {
	mvc.JsonAction
}

type AddUserAction struct {
	mvc.JsonAction
}

type DeleteUserAction struct {
	mvc.JsonAction
}

type ChangeUserAction struct {
	mvc.JsonAction
}

type ModPasswordAction struct {
	mvc.JsonAction
}

type LockUserAction struct {
	mvc.JsonAction
}

type UnLockUserAction struct {
	mvc.JsonAction
}

type OffLineUserAction struct {
	mvc.JsonAction
}

type QueryUserRequest struct {
	mvc.FastRequestWrap
	UserId     string `json:"userId"`
	DepartId   string `json:"departId"`
	EmployeeId string `json:"employeeId"`
}

type QueryUserResponse struct {
	mvc.FastResponseWrap
	Users []SafeUser `json:"users"`
}

type AddUserRequest struct {
	mvc.FastRequestWrap
	SafeUser
}

type AddUserResponse struct {
	mvc.FastResponseWrap
	SafeUser
}

type DeleteUserRequest struct {
	mvc.FastRequestWrap
	UserId string `json:"userId"`
}

type DeleteUserResponse struct {
	mvc.FastResponseWrap
}

type ChangeUserRequest struct {
	mvc.FastRequestWrap
	SafeUser
}

type ChangeUserResponse struct {
	mvc.FastResponseWrap
	SafeUser
}

type ModPasswordRequest struct {
	mvc.FastRequestWrap
	UserId      string `json:"userId"`
	NewPassword string `json:"newPassword"`
}

type ModPasswordResponse struct {
	mvc.FastResponseWrap
}

type LockUserRequest struct {
	mvc.FastRequestWrap
	UserId     string `json:"userId"`
	LockReason string `json:"lockReason"`
}

type LockUserResponse struct {
	mvc.FastResponseWrap
}

type UnLockUserRequest struct {
	mvc.FastRequestWrap
	UserId string `json:"userId"`
}

type UnLockUserResponse struct {
	mvc.FastResponseWrap
}

type OffLineUserRequest struct {
	mvc.FastRequestWrap
	UserId string `json:"userId"`
}

type OffLineUserResponse struct {
	mvc.FastResponseWrap
}

func (this *QueryUserAction) Post() {
	reqMsg := &QueryUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	reqMsg.UserId = strings.TrimSpace(reqMsg.UserId)
	reqMsg.DepartId = strings.TrimSpace(reqMsg.DepartId)
	reqMsg.EmployeeId = strings.TrimSpace(reqMsg.EmployeeId)

	queryUserSql := "select distinct * from fast.safeUser "
	if reqMsg.UserId != comm.NULL_STR {
		queryUserSql += "where id='" + reqMsg.UserId + "'"
	} else if reqMsg.DepartId != comm.NULL_STR {
		queryUserSql += "where departId ='" + reqMsg.DepartId + "'"
	} else if reqMsg.EmployeeId != comm.NULL_STR {
		queryUserSql += "where employeeId='" + reqMsg.EmployeeId + "'"
	}
	queryUserSql += " order by name"

	userRows, err := db.Query(queryUserSql)
	utils.VerifyErr(err)

	rspMsg := &QueryUserResponse{}
	for userRows.Next() {
		user := SafeUser{}
		userRows.Scan(
			&user.Id,
			&user.DepartId,
			&user.Name,
			&user.Password,
			&user.EmployeeId,
			&user.NickName,
			&user.FirstName,
			&user.LastName,
			&user.Mobile,
			&user.Email,
			&user.UserRemark,
			&user.IsLock,
			&user.LockReason,
			&user.IsForever,
			&user.AccountExpired)

		if user.Name == "daemon" {
			continue
		}
		rspMsg.Users = append(rspMsg.Users, user)
	}

	roleSql := "select distinct roleAlloc.* from fast.roleAlloc,fast.safeUser "
	roleSql += "where roleAlloc.userId = safeUser.id and roleAlloc.roleId != roleAlloc.userId "
	if reqMsg.UserId != comm.NULL_STR {
		roleSql += "and safeUser.id='" + reqMsg.UserId + "'"
	} else if reqMsg.DepartId != comm.NULL_STR {
		roleSql += "where safeUser.departId ='" + reqMsg.DepartId + "'"
	} else if reqMsg.EmployeeId != comm.NULL_STR {
		roleSql += "where safeUser.employeeId='" + reqMsg.EmployeeId + "'"
	}
	roleSql += " order by roleAlloc.userId"

	roleRows, err := db.Query(roleSql)
	utils.VerifyErr(err)

	oldUser := ""
	newUser := ""

	mapRole := map[string]([]string){}
	newRoles := []string{}

	roleId := comm.NULL_STR
	for roleRows.Next() {
		err = roleRows.Scan(&roleId, &newUser)
		utils.VerifyErr(err)
		if oldUser == comm.NULL_STR {
			oldUser = newUser
		} else if oldUser == newUser {
		} else {
			mapRole[oldUser] = newRoles

			oldUser = newUser
			newRoles = []string{}
		}
		newRoles = append(newRoles, roleId)
	}
	if oldUser != comm.NULL_STR {
		mapRole[oldUser] = newRoles
	}

	operateSql := "select distinct operateAuth.operateId,roleAlloc.userId from fast.roleAlloc roleAlloc,fast.operateAuth operateAuth "
	operateSql += "where roleAlloc.roleId = operateAuth.roleId order by roleAlloc.userId"

	operateRows, err := db.Query(operateSql)
	utils.VerifyErr(err)

	oldUser = ""
	newUser = ""

	mapOperate := map[string]([]string){}
	newOperates := []string{}

	operateId := comm.NULL_STR
	for operateRows.Next() {
		err = operateRows.Scan(&operateId, &newUser)
		utils.VerifyErr(err)
		if oldUser == comm.NULL_STR {
			oldUser = newUser
		} else if oldUser == newUser {
		} else {
			mapOperate[oldUser] = newOperates

			oldUser = newUser
			newOperates = []string{}
		}
		newOperates = append(newOperates, roleId)
	}
	if oldUser != comm.NULL_STR {
		mapOperate[oldUser] = newOperates
	}

	objectSql := "select distinct objAuth.objectId,roleAlloc.userId from fast.roleAlloc roleAlloc,fast.objectAuth objAuth "
	objectSql += "where roleAlloc.roleId = objAuth.roleId order by roleAlloc.userId"

	objectRows, err := db.Query(objectSql)
	utils.VerifyErr(err)

	oldUser = ""
	newUser = ""

	mapObject := map[string]([]string){}
	newObjects := []string{}

	objectId := comm.NULL_STR
	for objectRows.Next() {
		err = objectRows.Scan(&objectId, &newUser)
		utils.VerifyErr(err)
		if oldUser == comm.NULL_STR {
			oldUser = newUser
		} else if oldUser == newUser {
		} else {
			mapObject[oldUser] = newObjects

			oldUser = newUser
			newObjects = []string{}
		}
		newObjects = append(newObjects, objectId)
	}
	if oldUser != comm.NULL_STR {
		mapObject[oldUser] = newObjects
	}

	for i, v := range rspMsg.Users {
		roles, ok := mapRole[v.Id]
		if !ok {
			rspMsg.Users[i].Roles = []string{}
		} else {
			rspMsg.Users[i].Roles = roles
		}

		operates, ok := mapOperate[v.Id]
		if !ok {
			rspMsg.Users[i].Operates = []string{}
		} else {
			rspMsg.Users[i].Operates = operates
		}

		objects, ok := mapObject[v.Id]
		if !ok {
			rspMsg.Users[i].Objects = []string{}
		} else {
			rspMsg.Users[i].Objects = objects
		}
	}

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)

}

func (this *AddUserAction) Post() {
	reqMsg := &AddUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	msgId, ok := UserMgr().createUser(&reqMsg.SafeUser)

	rspMsg := &AddUserResponse{}
	rspMsg.Init(reqMsg)
	if !ok {
		rspMsg.SetRsp(msgId)
	} else {
		rspMsg.SafeUser = reqMsg.SafeUser
	}
	rspMsg.Password = comm.NULL_STR

	WriteLog(rspMsg, reqMsg.Name, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *ChangeUserAction) Post() {
	reqMsg := &ChangeUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	if reqMsg.Password != comm.NULL_STR {
		reqMsg.Password = utils.PasswordBySeed(reqMsg.Password, reqMsg.SafeUser.GetId())
	}

	msgId, ok := UserMgr().changeUser(&reqMsg.SafeUser)

	rspMsg := &ChangeUserResponse{}
	rspMsg.Init(reqMsg)
	rspMsg.Init(reqMsg)
	if !ok {
		rspMsg.SetRsp(msgId)
	} else {
		newUser, _ := UserMgr().getUser(reqMsg.Id)
		if newUser != nil {
			rspMsg.SafeUser = *newUser
		}

	}

	WriteLog(rspMsg, reqMsg.Name, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *DeleteUserAction) Post() {
	reqMsg := &DeleteUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.UserId = strings.TrimSpace(reqMsg.UserId)

	UserMgr().delUser(reqMsg.UserId)

	rspMsg := &DeleteUserResponse{}
	rspMsg.Init(reqMsg)

	WriteLog(rspMsg, reqMsg.UserId, comm.NULL_STR)
	this.SendJson(rspMsg)
}

func (this *ModPasswordAction) Post() {
	reqMsg := &ModPasswordRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.UserId = strings.TrimSpace(reqMsg.UserId)
	reqMsg.NewPassword = strings.TrimSpace(reqMsg.NewPassword)

	rspMsg := &DeleteUserResponse{}
	rspMsg.Init(reqMsg)
	if reqMsg.UserId == comm.NULL_STR || reqMsg.NewPassword == comm.NULL_STR {
		this.SendJson(rspMsg)
		return
	}

	db := ds.DB()
	defer db.Close()

	modPasswordSql := "update fast.safeUser set password=$1 where id=$2"

	modPasswordStmt, err := db.Prepare(modPasswordSql)
	utils.VerifyErr(err)
	modPasswordStmt.Exec(reqMsg.NewPassword, reqMsg.UserId)
	modPasswordStmt.Close()

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)

}

func (this *LockUserAction) Post() {
	reqMsg := &LockUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.UserId = strings.TrimSpace(reqMsg.UserId)
	reqMsg.LockReason = strings.TrimSpace(reqMsg.LockReason)
	rspMsg := &LockUserResponse{}
	rspMsg.Init(reqMsg)
	if reqMsg.UserId == comm.NULL_STR {
		this.SendJson(rspMsg)
		return
	}

	db := ds.DB()
	defer db.Close()

	lockUserSql := "update fast.safeUser set isLock=$1, lockReason=$2 where id=$3"

	lockUserStmt, err := db.Prepare(lockUserSql)
	utils.VerifyErr(err)
	lockUserStmt.Exec(true, reqMsg.LockReason, reqMsg.UserId)
	lockUserStmt.Close()

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)

}

func (this *UnLockUserAction) Post() {
	reqMsg := &UnLockUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.UserId = strings.TrimSpace(reqMsg.UserId)
	rspMsg := &UnLockUserResponse{}
	rspMsg.Init(reqMsg)
	if reqMsg.UserId == comm.NULL_STR {
		this.SendJson(rspMsg)
		return
	}

	db := ds.DB()
	defer db.Close()

	unLockUserSql := "update fast.safeUser set isLock=$1, lockReason=$2 where id=$3"

	unLockUserStmt, err := db.Prepare(unLockUserSql)
	utils.VerifyErr(err)
	unLockUserStmt.Exec(false, comm.NULL_STR, reqMsg.UserId)
	unLockUserStmt.Close()

	rspMsg.Init(reqMsg)
	this.SendJson(rspMsg)

}

func (this *OffLineUserAction) Post() {
	reqMsg := &OffLineUserRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.UserId = strings.TrimSpace(reqMsg.UserId)

	rspMsg := &OffLineUserResponse{}
	rspMsg.Init(reqMsg)

	this.SendJson(rspMsg)
}
