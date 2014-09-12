package safe

import (
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/utils"
)

type SafeOperateAgent struct {
}

var (
	DefaultSafeOperateAgent = &SafeOperateAgent{}
)

func GetSafeOperateDatas() []SafeOperate {
	return getSafeOperateAgent().getDatas()
}

func getSafeOperateAgent() *SafeOperateAgent {
	return DefaultSafeOperateAgent
}

func (this *SafeOperateAgent) getDatas() []SafeOperate {
	datas := []SafeOperate{}

	db := ds.DB()
	defer db.Close()

	operateSql := "select * from fast.safeOperate where 1=1"
	operateRows, err := db.Query(operateSql)
	utils.VerifyErr(err)

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

		operate.Departs = []string{}
		operate.Users = []string{}
		operate.Roles = []string{}

		datas = append(datas, operate)
	}

	return datas
}
