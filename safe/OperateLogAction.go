package safe

import (
	"fmt"
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/utils"
	"strings"
)

type QueryOperateLogAction struct {
	mvc.JsonAction
}

type QueryOperateLogRequest struct {
	mvc.FastRequestWrap
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	Subsys         string `json:"subsys"`
	OperateGroup   string `json:"operateGroup"`
	OperateName    string `json:"operateName"`
	UserName       string `json:"userName"`
	OperateObj     string `json:"operateObj"`
	OperateContent string `json:"operateContent"`
}

type QueryOperateLogResponse struct {
	mvc.FastResponseWrap
	Logs []OperateLog `json:"logs"`
}

type ImportOperateLogAction struct {
	mvc.JsonAction
}

type ImportOperateLogRequest struct {
	mvc.FastRequestWrap
	Logs []OperateLog `json:"logs"`
}

type ImportOperateLogResponse struct {
	mvc.FastResponseWrap
}

func (this *QueryOperateLogAction) Post() {
	reqMsg := &QueryOperateLogRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	reqMsg.StartTime = strings.TrimSpace(reqMsg.StartTime)
	reqMsg.EndTime = strings.TrimSpace(reqMsg.EndTime)
	reqMsg.Subsys = strings.TrimSpace(reqMsg.Subsys)
	reqMsg.OperateGroup = strings.TrimSpace(reqMsg.OperateGroup)
	reqMsg.OperateName = strings.TrimSpace(reqMsg.OperateName)
	reqMsg.UserName = strings.TrimSpace(reqMsg.UserName)
	reqMsg.OperateObj = strings.TrimSpace(reqMsg.OperateObj)
	reqMsg.OperateContent = strings.TrimSpace(reqMsg.OperateContent)

	logSql := "select log.* from fast.operateLog log where 1=1"
	if reqMsg.StartTime != comm.NULL_STR {
		logSql += " and log.operateTime >= '" + reqMsg.StartTime + "'"
	}
	if reqMsg.EndTime != comm.NULL_STR {
		logSql += " and log.operateTime <= '" + reqMsg.EndTime + "'"
	}
	if reqMsg.Subsys != comm.NULL_STR {
		logSql += " and log.subsys = '" + reqMsg.Subsys + "'"
	}
	if reqMsg.OperateGroup != comm.NULL_STR {
		logSql += " and log.operateGroup = '" + reqMsg.OperateGroup + "'"
	}
	if reqMsg.OperateName != comm.NULL_STR {
		logSql += " and log.operateName = '" + reqMsg.OperateName + "'"
	}
	if reqMsg.UserName != comm.NULL_STR {
		logSql += " and log.userName = '" + reqMsg.UserName + "'"
	}
	if reqMsg.OperateObj != comm.NULL_STR {
		logSql += " and log.operateObj like '%" + reqMsg.OperateObj + "%'"
	}
	if reqMsg.OperateContent != comm.NULL_STR {
		logSql += " and log.operateContent <= '%" + reqMsg.OperateContent + "%'"
	}

	logSql += " order by log.serialNum desc"

	db := ds.DB()
	defer db.Close()

	fmt.Println("QueryLog:", logSql)
	logRows, err := db.Query(logSql)
	utils.VerifyErr(err)

	rspMsg := &QueryOperateLogResponse{}
	rspMsg.Init(reqMsg)
	rspMsg.Logs = []OperateLog{}

	for logRows.Next() {
		log := OperateLog{}

		logRows.Scan(
			&log.UserName,
			&log.Subsys,
			&log.OperateGroup,
			&log.OperateName,
			&log.OperateRet,
			&log.OperateRetDetail,
			&log.OperateObj,
			&log.OperateContent,
			&log.UserIpAddress,
			&log.UserHostName,
			&log.OperateTime,
			&log.LogType,
			&log.SerialNum)

		rspMsg.Logs = append(rspMsg.Logs, log)
	}

	this.SendJson(rspMsg)

	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}

func (this *ImportOperateLogAction) Post() {
	reqMsg := &ImportOperateLogRequest{}
	if !this.GetReqJson(reqMsg) {
		return
	}

	rspMsg := &ImportOperateLogResponse{}
	rspMsg.Init(reqMsg)

	for _, log := range reqMsg.Logs {
		LogMgr().WriteRawLog(&log)
	}

	this.SendJson(rspMsg)
	WriteLog(rspMsg, comm.NULL_STR, comm.NULL_STR)

	return
}
