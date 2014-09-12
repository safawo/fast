package safe

import (
	"fmt"
	"think/fast/ds"
	"think/fast/utils"
)

type OperateMgrInterface interface {
	GetOperate(operateCode string) (operate *SafeOperate, ok bool)
}

func OperateMgr() (operateMgr OperateMgrInterface) {
	if operateMgrInstance == nil {
		operateMgrInstance = &operateMgrImpl{}
		operateMgrInstance.init()
	}

	return operateMgrInstance
}

type operateMgrImpl struct {
	mapOperate map[string](*SafeOperate)
}

var operateMgrInstance *operateMgrImpl

func (this *operateMgrImpl) init() {
	fmt.Println("  Init Operate Mgr")

	this.mapOperate = map[string](*SafeOperate){}

	db := ds.DB()
	defer db.Close()

	querySql := "select * from fast.safeOperate"
	operateRows, err := db.Query(querySql)
	utils.VerifyErr(err)

	for operateRows.Next() {
		operate := &SafeOperate{}
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
		this.mapOperate[operate.OperateCode] = operate
	}

}

func (this *operateMgrImpl) GetOperate(operateCode string) (operate *SafeOperate, ok bool) {
	operate, ok = this.mapOperate[operateCode]
	return
}
