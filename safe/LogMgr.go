package safe

import (
	"encoding/json"
	"fmt"
	"github.com/safawo/fast/ds"
	"github.com/safawo/fast/mvc"
	"github.com/safawo/fast/utils"
	"io/ioutil"
	"os"
	"strings"
)

func WriteLog(rspMsg mvc.FastResponseInterface, object, content string) {
	LogMgr().WriteLog(rspMsg, object, content)
}

func ParseLogs(sessionId string) {
	LogMgr().ParseLogs(sessionId)
}

type LogMgrInterface interface {
	WriteLog(rspMsg mvc.FastResponseInterface, object, content string)
	WriteRawLog(log *OperateLog)

	ParseLogs(sessionId string)

	queryLogs(cond string) (logs []OperateLog)
}

func LogMgr() (logMgr LogMgrInterface) {
	if logMgrInstance == nil {
		logMgrInstance = &logMgrImpl{}
		logMgrInstance.init()
	}

	return logMgrInstance
}

type logMgrImpl struct {
}

var logMgrInstance *logMgrImpl

func (this *logMgrImpl) init() {
}

func (this *logMgrImpl) WriteLog(rspMsg mvc.FastResponseInterface, object, content string) {
	sessMgr := SessionMgr()
	session, ok := sessMgr.GetSession(rspMsg.GetReqSessionId())
	if !ok {
		return
	}

	operateMgr := OperateMgr()
	operateCode := strings.Replace(rspMsg.GetReqActionId(), "/", ".", -1)
	operateCode = utils.Substr(operateCode, 1, len(operateCode))
	operate, ok := operateMgr.GetOperate(operateCode)
	if !ok {
		return
	}

	log := &OperateLog{}

	log.UserName = session.LoginUserName
	log.Subsys = operate.Subsys
	log.OperateGroup = operate.OperateGroup
	log.OperateName = operate.OperateName

	log.OperateRet = rspMsg.GetRspRetId()
	log.OperateRetDetail = rspMsg.GetRspRetMsg()

	log.OperateObj = object
	log.OperateContent = content

	log.UserIpAddress = session.ClientIp
	log.UserHostName = session.ClientHost

	log.OperateTime = rspMsg.GetReqStartTime()

	log.LogType = "User"
	log.SerialNum = utils.RandStr()

	this.WriteRawLog(log)

}

func (this *logMgrImpl) WriteRawLog(log *OperateLog) {

	if log == nil {
		return
	}

	if log.IsEmpty() {
		return
	}

	db := ds.DB()
	defer db.Close()

	logSql := "insert into fast.operateLog values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)"
	logStmt, err := db.Prepare(logSql)
	utils.VerifyErr(err)
	logStmt.Exec(
		log.UserName,
		log.Subsys,
		log.OperateGroup,
		log.OperateName,
		log.OperateRet,
		log.OperateRetDetail,
		log.OperateObj,
		log.OperateContent,
		log.UserIpAddress,
		log.UserHostName,
		log.OperateTime,
		log.LogType,
		log.SerialNum)
	logStmt.Close()
}

func (this *logMgrImpl) queryLogs(cond string) (logs []OperateLog) {
	logs = []OperateLog{}
	return
}

func (this *logMgrImpl) ParseLogs(sessionId string) {
	sessMgr := SessionMgr()
	session, ok := sessMgr.GetSession(sessionId)
	if !ok {
		return
	}

	padinDir, err := os.Open(this.getPadinPath())
	utils.VerifyErr(err)
	defer padinDir.Close()

	logFiles, err := padinDir.Readdir(0)
	utils.VerifyErr(err)
	for _, logFileInfo := range logFiles {

		if logFileInfo.IsDir() {
			continue
		}
		logFileName := logFileInfo.Name()
		if !strings.HasSuffix(logFileName, ".json") {
			continue
		}
		if !strings.HasPrefix(logFileName, "PadLog") {
			continue
		}

		fmt.Println("PadLog file,name=", this.getPadinPath()+"/"+logFileName)
		file, err := os.Open(this.getPadinPath() + "/" + logFileName)
		utils.VerifyErr(err)
		defer func() {
			file.Close()
			os.Remove(this.getPadinPath() + "/" + logFileName)
		}()

		bytes, err := ioutil.ReadAll(file)
		utils.VerifyErr(err)
		if len(bytes) < 20 {
			continue
		}

		if bytes[len(bytes)-1] == ',' {
			bytes[len(bytes)-1] = ']'
		} else if bytes[len(bytes)-2] == ',' {
			bytes[len(bytes)-2] = ']'
		} else if bytes[len(bytes)-3] == ',' {
			bytes[len(bytes)-3] = ']'
		} else if bytes[len(bytes)-4] == ',' {
			bytes[len(bytes)-4] = ']'
		} else if bytes[len(bytes)-5] == ',' {
			bytes[len(bytes)-5] = ']'
		}

		logs := []OperateLog{}
		err = json.Unmarshal(bytes, &logs)
		utils.VerifyErr(err)

		for _, log := range logs {
			log.OperateRetDetail = "Import By " + session.GetUserName() + " at " + utils.StrTime()
			this.WriteRawLog(&log)
		}

	}

}

func (this *logMgrImpl) getPadinPath() string {
	padinDir := utils.CachePath() + "/padin"
	_, err := os.Stat(padinDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(padinDir, os.ModeDir)
	}

	return padinDir
}
