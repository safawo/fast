package safe

import (
	"fmt"
	"net/http"
	"strings"
	"think/fast/comm"
	"think/fast/ds"
	"think/fast/msg"
	"think/fast/utils"
)

type SessionMgrInterface interface {
	IsOnLine(userName string) (onLine bool)
	IsValid(sessionId string) (ok bool)

	GetSessions() (sessions [](*SafeSession))
	GetSession(sessionId string) (session *SafeSession, ok bool)

	GetSessionByName(userName string) (session *SafeSession, ok bool)
	Offline(userName string) (result bool)

	login(req *http.Request, userName, pasword, hostIp, hostName string) (session *SafeSession, msgId string, ok bool)
	logout(sessionId string)
}

func SessionMgr() (sessionMgr SessionMgrInterface) {
	if sessionMgrInstance == nil {
		sessionMgrInstance = &sessionMgrImpl{}
		sessionMgrInstance.init()
	}

	return sessionMgrInstance
}

type sessionMgrImpl struct {
	mapSessions map[string](*SafeSession)
	badSessions map[string](*SafeSession)
}

var sessionMgrInstance *sessionMgrImpl

func (this *sessionMgrImpl) init() {

	fmt.Println("  Init User Session")

	this.mapSessions = map[string](*SafeSession){}
	this.badSessions = map[string](*SafeSession){}

	db := ds.DB()
	defer db.Close()

	initSql := "truncate table fast.safeSession"

	initStmt, err := db.Prepare(initSql)
	utils.VerifyErr(err)

	initStmt.Exec()
	initStmt.Close()
}

func (this *sessionMgrImpl) IsOnLine(userName string) (onLine bool) {
	onLine = false

	for _, v := range this.mapSessions {
		if v.LoginUserName == userName {
			onLine = true
			break
		}
	}

	return
}

func (this *sessionMgrImpl) IsValid(sessionId string) (ok bool) {
	_, ok = this.mapSessions[sessionId]
	return
}

func (this *sessionMgrImpl) GetSessions() (sessions [](*SafeSession)) {
	sessions = [](*SafeSession){}

	for _, v := range this.mapSessions {
		sessions = append(sessions, v)
	}

	return
}

func (this *sessionMgrImpl) GetSession(sessionId string) (session *SafeSession, ok bool) {
	session, ok = this.mapSessions[sessionId]
	return
}

func (this *sessionMgrImpl) GetSessionByName(userName string) (session *SafeSession, ok bool) {
	session = nil
	ok = false
	for _, v := range this.mapSessions {
		if v.LoginUserName == userName {
			ok = true
			session = v
			break
		}
	}

	return
}

func (this *sessionMgrImpl) Offline(userName string) (result bool) {
	result = false

	session, ok := this.GetSessionByName(userName)
	if !ok {
		return true
	}

	delete(this.mapSessions, session.GetId())
	return true
}

func (this *sessionMgrImpl) login(req *http.Request, userName, pasword, hostIp, hostName string) (session *SafeSession, msgId string, ok bool) {
	session = nil
	msgId = comm.NULL_STR
	ok = false

	userName = strings.TrimSpace(userName)
	pasword = strings.TrimSpace(pasword)

	if userName == comm.NULL_STR || pasword == comm.NULL_STR {
		msgId = msg.MSG_PARA_ABSENT
		return
	}

	user, ok := UserMgr().getUserByName(userName)
	if !ok {
		msgId = MSG_USER_INVALID
		return
	}

	msgId, ok = UserMgr().isValid(user.GetId())
	if !ok {
		return
	}

	pasword = utils.PasswordBySeed(pasword, user.GetId())

	if user.Password != pasword {
		msgId = MSG_PASSWORD_ERROR
		ok = false
		return
	}

	if this.IsOnLine(userName) {
		this.Offline(userName)
	}

	clientOnlineLimit := LicenseMgr().GetAllowInt("clientOnlineLimit")
	activateClient := len(this.mapSessions)
	if activateClient >= clientOnlineLimit && !LicenseMgr().IsEmpty() {
		msgId = msg.MSG_LICENSE_NOENOUGH
		ok = false
		return
	}

	session = &SafeSession{}

	session.LoginUserId = user.Id
	session.LoginUserName = user.Name

	session.RandId = utils.RandStr()
	session.ConnTime = utils.StrTime()

	session.ClientIp = hostIp
	session.ClientHost = hostName
	session.ServerIp = req.Host
	session.ServerHost = req.Host

	session.KeepAliveTime = session.ConnTime
	session.SessionStatus = SAFE_SESSION_LOGIN

	session.Id = session.LoginUserName + "&" + session.ConnTime + "&" + utils.RandStr()
	session.Id = utils.EnCodeBase64(session.GetId())
	session.Id = utils.Md5(session.GetId())

	this.mapSessions[session.GetId()] = session

	db := ds.DB()
	defer db.Close()

	insertSql := "insert into fast.safeSession values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)"
	insertStmt, err := db.Prepare(insertSql)
	utils.VerifyErr(err)
	insertStmt.Exec(
		session.Id,
		session.RandId,
		session.ConnTime,
		session.LoginUserId,
		session.LoginUserName,
		session.ClientIp,
		session.ClientHost,
		session.ServerIp,
		session.ServerHost,
		session.KeepAliveTime,
		session.SessionStatus)

	insertStmt.Close()

	msgId = msg.MSG_SUCCESS
	ok = true

	return
}

func (this *sessionMgrImpl) logout(sessionId string) {
	_, ok := this.mapSessions[sessionId]
	if !ok {
		return
	}

	delete(this.mapSessions, sessionId)

	db := ds.DB()
	defer db.Close()

	delSql := "delete from fast.safeSession where id=$1"
	delStmt, err := db.Prepare(delSql)
	utils.VerifyErr(err)
	delStmt.Exec(sessionId)
	delStmt.Close()

	return
}
