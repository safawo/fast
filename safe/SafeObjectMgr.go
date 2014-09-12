package safe

import (
	"fmt"
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/msg"
	"github.com/safawo/fast/utils"
	"strings"
)

type safeObjectMgrImpl struct {
}

var safeObjectMgrInstance *safeObjectMgrImpl

func SafeObjectMgr() (safeObjectMgr *safeObjectMgrImpl) {
	if safeObjectMgrInstance == nil {
		safeObjectMgrInstance = &safeObjectMgrImpl{}
		safeObjectMgrInstance.init()
	}

	return safeObjectMgrInstance
}

func (this *safeObjectMgrImpl) init() {
	fmt.Println("  Init Safe Object Mgr")
}

func (this *safeObjectMgrImpl) CreateObject(safeObj *SafeObject) (msgId string, ok bool) {
	msgId = comm.NULL_STR
	ok = false

	if safeObj == nil {
		return
	}

	db := ds.DB()
	defer db.Close()

	createSql := "insert into fast.safeObject values($1,$2,$3,$4,$5,$6,$7,$8)"
	createStmt, err := db.Prepare(createSql)
	utils.VerifyErr(err)

	createStmt.Exec(
		safeObj.Id,
		safeObj.SerialId,
		safeObj.ObjectType,
		safeObj.ParentPath,
		safeObj.ObjectName,
		safeObj.ObjectText,
		safeObj.ObjectImage,
		safeObj.ObjectRemark)
	createStmt.Close()

	msgId = msg.MSG_SUCCESS
	ok = true

	return
}

func (this *safeObjectMgrImpl) DelObject(objectId string) {
	objectId = strings.TrimSpace(objectId)

	db := ds.DB()
	defer db.Close()

	delSql := "delete from fast.safeObject where id=$1"

	delStmt, err := db.Prepare(delSql)
	utils.VerifyErr(err)
	delStmt.Exec(objectId)
	delStmt.Close()

	return
}

func (this *safeObjectMgrImpl) DelObjectByName(objectName string) {
	objectName = strings.TrimSpace(objectName)

	db := ds.DB()
	defer db.Close()

	delSql := "delete from fast.safeObject where objectName=$1"

	delStmt, err := db.Prepare(delSql)
	utils.VerifyErr(err)
	delStmt.Exec(objectName)
	delStmt.Close()

	return
}

func (this *safeObjectMgrImpl) DelObjectByObjNameAndType(objectName, objectType string) {
	objectName = strings.TrimSpace(objectName)
	objectType = strings.TrimSpace(objectType)

	db := ds.DB()
	defer db.Close()

	delSql := "delete from fast.safeObject where objectName=$1 and objectType=$2"

	delStmt, err := db.Prepare(delSql)
	utils.VerifyErr(err)
	delStmt.Exec(objectName, objectType)
	delStmt.Close()

	this.unAuthByObj(objectName, objectType)

	return
}

func (this *safeObjectMgrImpl) QueryObjects() (objects []SafeObject) {
	return this.QueryObjectsByType(comm.NULL_STR)
}

func (this *safeObjectMgrImpl) QueryObjectsByType(objectType string) (objects []SafeObject) {
	objects = [](SafeObject){}
	objectType = strings.TrimSpace(objectType)

	querySql := "select distinct * from fast.safeObject"
	if objectType != comm.NULL_STR {
		querySql += " where objectType='" + objectType + "'"
	}

	db := ds.DB()
	defer db.Close()

	objectRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	for objectRows.Next() {
		object := SafeObject{}

		objectRows.Scan(
			&object.Id,
			&object.SerialId,
			&object.ObjectType,
			&object.ParentPath,
			&object.ObjectName,
			&object.ObjectText,
			&object.ObjectImage,
			&object.ObjectRemark)

		objects = append(objects, object)
	}

	return
}

func (this *safeObjectMgrImpl) Auth(auths []ObjectAuthInfo) {

	db := ds.DB()
	defer db.Close()

	authSql := "insert into fast.objectAuth values($1,$2)"

	authStmt, err := db.Prepare(authSql)
	utils.VerifyErr(err)
	for _, v := range auths {
		authStmt.Exec(v.ObjectId, v.RoleId)
	}
	authStmt.Close()

	return
}

func (this *safeObjectMgrImpl) unAuthByObj(objectName, objectType string) {
	objectName = strings.TrimSpace(objectName)
	objectType = strings.TrimSpace(objectType)

	db := ds.DB()
	defer db.Close()

	unAuthSql := "delete from fast.objectAuth where objectId in (select id from fast.safeObject where"
	unAuthSql += " objectName=$1 and objectType=$2)"

	unAuthStmt, err := db.Prepare(unAuthSql)
	utils.VerifyErr(err)
	unAuthStmt.Exec(objectName, objectType)

	unAuthStmt.Close()

	return
}

func (this *safeObjectMgrImpl) UnAuth(unAuths []ObjectAuthInfo) {

	db := ds.DB()
	defer db.Close()

	unAuthSql := "delete from fast.objectAuth where objectId=$1 and roleId=$2"

	unAuthStmt, err := db.Prepare(unAuthSql)
	utils.VerifyErr(err)
	for _, v := range unAuths {
		unAuthStmt.Exec(v.ObjectId, v.RoleId)
	}
	unAuthStmt.Close()

	return
}

func (this *safeObjectMgrImpl) QueryAuths() (objects []ObjectAuthInfo) {
	return this.QueryAuthsByType(comm.NULL_STR)
}

func (this *safeObjectMgrImpl) QueryAuthsByType(objectType string) (objects []ObjectAuthInfo) {
	objects = []ObjectAuthInfo{}

	db := ds.DB()
	defer db.Close()

	objectType = strings.TrimSpace(objectType)

	querySql := "select distinct objectAuth.* from fast.objectAuth"
	if objectType != comm.NULL_STR {
		querySql += ",fast.safeObject where objectAuth.objectId=safeObject.id "
		querySql += " and safeObject.objectType='" + objectType + "'"
	}

	authRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	for authRows.Next() {
		auth := ObjectAuthInfo{}
		authRows.Scan(&auth.ObjectId, &auth.RoleId)
		objects = append(objects, auth)
	}

	return
}
