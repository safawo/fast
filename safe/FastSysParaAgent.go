package safe

import (
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/utils"
)

type FastSysParaAgent struct {
}

var (
	DefaultFastSysParaAgent = &FastSysParaAgent{}
)

func GetFastSysParas() []utils.FastSysPara {
	return getFastSysParaAgent().getDatas()
}

func getFastSysParaAgent() *FastSysParaAgent {
	return DefaultFastSysParaAgent
}

func (this *FastSysParaAgent) getDatas() []utils.FastSysPara {

	paras := []utils.FastSysPara{}

	db := ds.DB()
	defer db.Close()

	querySql := "select * from fast.sysParaConf"

	rows, err := db.Query(querySql)
	utils.VerifyErr(err)

	for rows.Next() {
		newPara := utils.FastSysPara{}
		err = rows.Scan(
			&newPara.Catalog,
			&newPara.ParaName,
			&newPara.ParaValue,
			&newPara.ParaType,
			&newPara.ParaRemark)
		paras = append(paras, newPara)
	}

	return paras
}
