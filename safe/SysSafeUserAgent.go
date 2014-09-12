package safe

import (
	"think/fast/comm"
	"think/fast/ds"
	"think/fast/utils"
)

type SysSafeUserAgent struct {
}

var (
	DefaultSysSafeUserAgent = &SysSafeUserAgent{}
)

func GetSysSafeUserDatas() []SafeUser {
	return getSysSafeUserAgent().getDatas()
}

func getSysSafeUserAgent() *SysSafeUserAgent {
	return DefaultSysSafeUserAgent
}

func (this *SysSafeUserAgent) getDatas() []SafeUser {
	users := []SafeUser{}

	db := ds.DB()
	defer db.Close()

	userSql := "select distinct * from fast.safeUser order by name"

	userRows, err := db.Query(userSql)
	utils.VerifyErr(err)

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
		users = append(users, user)
	}

	roleSql := "select distinct roleAlloc.* from fast.roleAlloc,fast.safeUser "
	roleSql += "where roleAlloc.userId = safeUser.id and roleAlloc.roleId != roleAlloc.userId "
	roleSql += "order by roleAlloc.userId"

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
		newOperates = append(newOperates, operateId)
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

	for i, v := range users {
		roles, ok := mapRole[v.Id]
		if !ok {
			users[i].Roles = []string{}
		} else {
			users[i].Roles = roles
		}

		operates, ok := mapOperate[v.Id]
		if !ok {
			users[i].Operates = []string{}
		} else {
			users[i].Operates = operates
		}

		objects, ok := mapObject[v.Id]
		if !ok {
			users[i].Objects = []string{}
		} else {
			users[i].Objects = objects
		}
	}

	return users
}
