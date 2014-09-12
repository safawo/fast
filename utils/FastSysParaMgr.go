package utils

import (
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"strings"
)

type FastSysParaInterface interface {
	ExistPara(paraName string) bool

	ImportPara(para FastSysPara)
	ChangePara(paraName, paraValue string)

	GetParas() (paras [](*FastSysPara))
	GetPara(paraName string) (paraValue string)
}

type fastSysParaMgr struct {
	mapPara map[string](*FastSysPara)
}

var fastSysParaMgrInstance *fastSysParaMgr

func (this *fastSysParaMgr) init() {

	this.mapPara = map[string](*FastSysPara){}

	db := ds.DB()
	defer db.Close()

	querySql := "select * from fast.sysParaConf"

	rows, err := db.Query(querySql)
	VerifyErr(err)

	for rows.Next() {
		newPara := &FastSysPara{}
		err = rows.Scan(
			&newPara.Catalog,
			&newPara.ParaName,
			&newPara.ParaValue,
			&newPara.ParaType,
			&newPara.ParaRemark)
		this.mapPara[newPara.ParaName] = newPara
	}

}

func (this *fastSysParaMgr) ExistPara(paraName string) bool {
	_, ok := this.mapPara[paraName]
	if ok {
		return true
	}

	return false
}

func (this *fastSysParaMgr) ImportPara(para FastSysPara) {
	importSql := "insert into fast.sysParaConf values($1,$2,$3,$4,$5)"

	this.mapPara[para.ParaName] = &para

	db := ds.DB()
	defer db.Close()
	importStmt, err := db.Prepare(importSql)

	VerifyErr(err)
	importStmt.Exec(
		para.Catalog,
		para.ParaName,
		para.ParaValue,
		para.ParaType,
		para.ParaRemark)
	importStmt.Close()

	return
}

func (this *fastSysParaMgr) ChangePara(paraName, paraValue string) {
	paraName = strings.TrimSpace(paraName)
	paraValue = strings.TrimSpace(paraValue)
	if paraName == comm.NULL_STR || paraValue == comm.NULL_STR {
		return
	}

	para, ok := this.mapPara[paraName]
	if !ok {
		return
	}

	para.ParaValue = paraValue

	db := ds.DB()
	defer db.Close()

	changeSql := "update fast.sysParaConf set paraValue=$1 where paraName=$2"

	changeStmt, err := db.Prepare(changeSql)
	VerifyErr(err)
	changeStmt.Exec(paraValue, paraName)
	changeStmt.Close()

	return
}

func SysParaMgr() FastSysParaInterface {
	if fastSysParaMgrInstance == nil {
		fastSysParaMgrInstance = &fastSysParaMgr{}
		fastSysParaMgrInstance.init()
	}

	return fastSysParaMgrInstance
}

func (this *fastSysParaMgr) GetParas() (paras [](*FastSysPara)) {

	paras = [](*FastSysPara){}

	for _, v := range this.mapPara {
		paras = append(paras, v)
	}

	return
}

func (this *fastSysParaMgr) GetPara(paraName string) (paraValue string) {
	paraValue = comm.NULL_STR

	para, ok := this.mapPara[paraName]
	if ok {
		paraValue = para.ParaValue
	}

	return
}
