package safe

import (
	"fmt"
	"strings"
	"think/fast/comm"
	"think/fast/ds"
	"think/fast/msg"
	"think/fast/utils"
)

type UserMgrInterface interface {
	init()

	existUser(userId string) (exist bool)
	existUserByName(userName string) (exist bool)
	existUserByEmployee(employeeId string) (exist bool)

	isValid(userId string) (msg string, ok bool)
	getUser(userId string) (user *SafeUser, exist bool)
	getUserByName(userName string) (user *SafeUser, exist bool)
	getUserByEmployee(employeeId string) (user *SafeUser, exist bool)

	getUsers() (users [](*SafeUser))
	getUsersByDepart(departId string) (users [](*SafeUser))

	createUser(user *SafeUser) (msgId string, ok bool)
	delUser(userId string)
	changeUser(user *SafeUser) (msgId string, ok bool)
}

func UserMgr() (userMgr UserMgrInterface) {
	if userMgrInstance == nil {
		userMgrInstance = &userMgrImpl{}
		userMgrInstance.init()
	}

	return userMgrInstance
}

type userMgrImpl struct {
	mapUser map[string](*SafeUser)
}

var userMgrInstance *userMgrImpl

func (this *userMgrImpl) init() {

	fmt.Println("  Init User Mgr")

	this.mapUser = map[string](*SafeUser){}

	db := ds.DB()
	defer db.Close()

	querySql := "select * from fast.safeUser"
	userRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	for userRows.Next() {
		user := &SafeUser{}
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
		this.mapUser[user.Id] = user
	}
}

func (this *userMgrImpl) existUser(userId string) (exist bool) {
	_, exist = this.getUser(userId)
	return exist
}

func (this *userMgrImpl) existUserByName(userName string) (exist bool) {
	_, exist = this.getUserByName(userName)
	return exist
}

func (this *userMgrImpl) existUserByEmployee(employeeId string) (exist bool) {
	_, exist = this.getUserByEmployee(employeeId)
	return exist
}

func (this *userMgrImpl) isValid(userId string) (msg string, ok bool) {
	msg = comm.NULL_STR
	ok = true
	return
}

func (this *userMgrImpl) getUser(userId string) (user *SafeUser, exist bool) {
	user = nil
	exist = false

	user, exist = this.mapUser[userId]

	return user, exist
}

func (this *userMgrImpl) getUserByName(userName string) (user *SafeUser, exist bool) {
	user = nil
	exist = false

	for _, v := range this.mapUser {
		if v.Name != userName {
			continue
		}

		user = v
		exist = true
		break
	}

	return user, exist
}

func (this *userMgrImpl) getUserByEmployee(employeeId string) (user *SafeUser, exist bool) {
	user = nil
	exist = false

	for _, v := range this.mapUser {
		if v.EmployeeId != employeeId {
			continue
		}

		user = v
		exist = true
		break
	}

	return user, exist
}

func (this *userMgrImpl) getUsers() (users [](*SafeUser)) {
	users = [](*SafeUser){}

	for _, v := range this.mapUser {
		users = append(users, v)
	}

	return users
}

func (this *userMgrImpl) getUsersByDepart(departId string) (users [](*SafeUser)) {
	users = [](*SafeUser){}

	for _, v := range this.mapUser {
		if v.DepartId != departId {
			continue
		}
		users = append(users, v)
	}

	return users
}

func (this *userMgrImpl) createUser(user *SafeUser) (msgId string, ok bool) {
	msgId = msg.MSG_FAIL
	ok = false

	if user == nil {
		return
	}

	user.Name = strings.TrimSpace(user.Name)
	user.EmployeeId = strings.TrimSpace(user.EmployeeId)
	user.Password = strings.TrimSpace(user.Password)

	if user.Name == comm.NULL_STR ||
		user.Password == comm.NULL_STR {
		msgId = msg.MSG_PARA_ABSENT
		return
	}

	user.Id = utils.RandStr()

	user.Password = utils.PasswordBySeed(user.Password, user.GetId())

	if this.existUserByName(user.Name) {
		msgId = MSG_USERNAME_CLASH
		return
	}

	if this.existUserByEmployee(user.EmployeeId) {
		msgId = MSG_EMPLOYEEID_CLASH
		return
	}

	db := ds.DB()
	defer db.Close()

	createUserSql := "insert into fast.safeUser values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)"
	createUserStmt, err := db.Prepare(createUserSql)
	utils.VerifyErr(err)
	createUserStmt.Exec(
		user.Id,
		user.DepartId,
		user.Name,
		user.Password,
		user.EmployeeId,
		user.NickName,
		user.FirstName,
		user.LastName,
		user.Mobile,
		user.Email,
		user.UserRemark,
		user.IsLock,
		user.LockReason,
		user.IsForever,
		user.AccountExpired)
	createUserStmt.Close()

	createMapRoleSql := "insert into fast.safeRole values($1,$2,$3,$4,$5,$6,$7)"
	createMapRoleStmt, err := db.Prepare(createMapRoleSql)
	utils.VerifyErr(err)
	createMapRoleStmt.Exec(
		user.Id,
		comm.NULL_STR,
		comm.NULL_STR,
		comm.NULL_STR,
		true,
		false,
		comm.NULL_STR)
	createMapRoleStmt.Close()

	insertRoleAllocSql := "insert into fast.roleAlloc values($1,$2)"
	insertRoleAllocStmt, err := db.Prepare(insertRoleAllocSql)
	utils.VerifyErr(err)
	for _, v := range user.Roles {
		insertRoleAllocStmt.Exec(v, user.Id)
	}
	insertRoleAllocStmt.Exec(user.Id, user.Id)

	insertRoleAllocStmt.Close()

	insertOperateAuthSql := "insert into fast.operateAuth values($1,$2)"
	insertOperateAuthStmt, err := db.Prepare(insertOperateAuthSql)
	utils.VerifyErr(err)
	for _, v := range user.Operates {
		insertOperateAuthStmt.Exec(v, user.Id)
	}
	insertOperateAuthStmt.Close()

	insertObjAuthSql := "insert into fast.objectAuth values($1,$2)"
	insertObjAuthStmt, err := db.Prepare(insertObjAuthSql)
	utils.VerifyErr(err)
	for _, v := range user.Objects {
		insertObjAuthStmt.Exec(v, user.Id)
	}
	insertObjAuthStmt.Close()

	this.mapUser[user.Id] = user

	msgId = msg.MSG_SUCCESS
	ok = true

	return
}

func (this *userMgrImpl) delUser(userId string) {
	userId = strings.TrimSpace(userId)
	delete(this.mapUser, userId)

	db := ds.DB()
	defer db.Close()

	delUserSql := "delete from fast.safeUser where id=$1"
	delUserStmt, err := db.Prepare(delUserSql)
	utils.VerifyErr(err)
	delUserStmt.Exec(userId)
	delUserStmt.Close()

	delRoleMapSql := "delete from fast.safeRole where id=$1 and isUser=true"
	delRoleMapStmt, err := db.Prepare(delRoleMapSql)
	utils.VerifyErr(err)
	delRoleMapStmt.Exec(userId)
	delRoleMapStmt.Close()

	delRoleAllocSql := "delete from fast.roleAlloc where userId=$1"
	delRoleAllocStmt, err := db.Prepare(delRoleAllocSql)
	utils.VerifyErr(err)
	delRoleAllocStmt.Exec(userId)
	delRoleAllocStmt.Close()

	delOperateAuthSql := "delete from fast.operateAuth where roleId=$1"
	delOperateAuthStmt, err := db.Prepare(delOperateAuthSql)
	utils.VerifyErr(err)
	delOperateAuthStmt.Exec(userId)
	delOperateAuthStmt.Close()

	delObjAuthSql := "delete from fast.roleAlloc where roleId=$1"
	delObjAuthStmt, err := db.Prepare(delObjAuthSql)
	utils.VerifyErr(err)
	delObjAuthStmt.Exec(userId)
	delObjAuthStmt.Close()

	return
}

func (this *userMgrImpl) changeUser(user *SafeUser) (msgId string, ok bool) {
	msgId = msg.MSG_FAIL
	ok = false

	if user == nil {
		return
	}

	user.Name = strings.TrimSpace(user.Name)
	user.EmployeeId = strings.TrimSpace(user.EmployeeId)
	user.Password = strings.TrimSpace(user.Password)

	if user.Name == comm.NULL_STR {
		msgId = msg.MSG_PARA_ABSENT
		return
	}

	existUser, ok := this.getUser(user.Id)
	if ok {
		if existUser.Name != user.Name && this.existUserByName(user.Name) {
			msgId = MSG_USERNAME_CLASH
			return
		}
		if existUser.EmployeeId != user.EmployeeId && this.existUserByEmployee(user.EmployeeId) {
			msgId = MSG_EMPLOYEEID_CLASH
			return
		}

	} else {
		msgId = MSG_USER_INVALID
		return
	}

	db := ds.DB()
	defer db.Close()

	changeSql := "update fast.safeUser set departId=$1,name=$2,employeeId=$3,nickName=$4,firstName=$5,lastName=$6,"
	changeSql += "mobile=$7,email=$8,userRemark=$9,isForever=$10,accountExpired=$11"
	if user.Password != comm.NULL_STR {
		//user.Password = utils.PasswordBySeed(user.Password, user.GetId())
		changeSql += ",password='" + user.Password + "'"
	}
	changeSql += " where id=$12"

	changeStmt, err := db.Prepare(changeSql)
	utils.VerifyErr(err)
	changeStmt.Exec(
		user.DepartId,
		user.Name,
		user.EmployeeId,
		user.NickName,
		user.FirstName,
		user.LastName,
		user.Mobile,
		user.Email,
		user.UserRemark,
		user.IsForever,
		user.AccountExpired,
		user.Id)
	changeStmt.Close()

	querySql := "select * from fast.safeUser where id='" + user.Id + "'"
	userRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	delRoleSql := "delete from fast.rolealloc alloc where alloc.roleid in "
	delRoleSql += "(select id from fast.safeRole where isUser=false and isDepart=false) "
	delRoleSql += "and alloc.userId='" + user.Id + "'"

	delRoleStmt, err := db.Prepare(delRoleSql)
	utils.VerifyErr(err)
	delRoleStmt.Exec()
	delRoleStmt.Close()

	allocRoleSql := "insert into fast.roleAlloc values($1,$2)"
	allocRoleStmt, err := db.Prepare(allocRoleSql)
	utils.VerifyErr(err)
	for _, v := range user.Roles {
		allocRoleStmt.Exec(v, user.Id)
	}
	allocRoleStmt.Close()

	newUser := &SafeUser{}
	for userRows.Next() {
		userRows.Scan(
			&newUser.Id,
			&newUser.DepartId,
			&newUser.Name,
			&newUser.Password,
			&newUser.EmployeeId,
			&newUser.NickName,
			&newUser.FirstName,
			&newUser.LastName,
			&newUser.Mobile,
			&newUser.Email,
			&newUser.UserRemark,
			&newUser.IsLock,
			&newUser.LockReason,
			&newUser.IsForever,
			&newUser.AccountExpired)
	}
	this.mapUser[newUser.Id] = newUser

	msgId = msg.MSG_SUCCESS
	ok = true

	return
}
