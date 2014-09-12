package safe

import (
	"think/fast/ds"
	"think/fast/mvc"
	"think/fast/utils"
)

type QuerySafeSubSysAction struct {
	mvc.JsonAction
}

type QuerySafeSubSysRequest struct {
	mvc.FastRequestWrap
}

type QuerySafeSubSysResponse struct {
	mvc.FastResponseWrap
	SubSyss []SubSys `json:"subSyss"`
}

func (this *QuerySafeSubSysAction) Post() {
	reqMsg := &QuerySafeSubSysRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	db := ds.DB()
	defer db.Close()

	subSysSql := "select * from fast.subSys"

	subSysRows, err := db.Query(subSysSql)
	utils.VerifyErr(err)

	rspMsg := &QuerySafeSubSysResponse{}
	rspMsg.Init(reqMsg)

	for subSysRows.Next() {
		subSys := SubSys{}

		subSysRows.Scan(
			&subSys.Id,
			&subSys.Name)

		rspMsg.SubSyss = append(rspMsg.SubSyss, subSys)
	}

	this.SendJson(rspMsg)
}
