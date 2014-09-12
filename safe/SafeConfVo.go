package safe

import (
	"github.com/safawo/fast/utils"
)

const (
	SAFE_SESSION_INIT          = "Init"
	SAFE_SESSION_LOGIN         = "Login"
	SAFE_SESSION_LOGOUT        = "Logout"
	SAFE_SESSION_TIMEOUT       = "Timeout"
	SAFE_SESSION_FORCE_OFFLINE = "ForceOffLine"
)

type SafeDepart struct {
	Id         string `json:"id"`
	SerialId   int    `json:"serialId"`
	ParentPath string `json:"parentPath"`

	DepartName   string `json:"departName"`
	DepartText   string `json:"departText"`
	DepartImage  string `json:"departImage"`
	DepartRemark string `json:"departRemark"`

	Users    []string `json:"users"`
	Roles    []string `json:"roles"`
	Operates []string `json:"operates"`
	Objects  []string `json:"objects"`
}

type SafeUser struct {
	Id       string `json:"id"`
	DepartId string `json:"departId"`
	Name     string `json:"name"`
	Password string `json:"password"`

	EmployeeId string `json:"employeeId"`
	NickName   string `json:"nickName"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`

	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
	UserRemark string `json:"userRemark"`

	IsLock         bool   `json:"isLock"`
	LockReason     string `json:"lockReason"`
	IsForever      bool   `json:"isForever"`
	AccountExpired string `json:"accountExpired"`

	Roles    []string `json:"roles"`
	Operates []string `json:"operates"`
	Objects  []string `json:"objects"`
}

type SafeSession struct {
	Id       string `json:"id"`
	RandId   string `json:"randId"`
	ConnTime string `json:"connTime"`

	LoginUserId   string `json:"loginUserId"`
	LoginUserName string `json:"loginUserName"`

	ClientIp   string `json:"clientIp"`
	ClientHost string `json:"clientHost"`
	ServerIp   string `json:"serverIp"`
	ServerHost string `json:"serverHost"`

	KeepAliveTime string `json:"keepAliveTime"`
	SessionStatus string `json:"sessionStatus"`
}

type SafeRole struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	RoleDetail string `json:"roleDetail"`
	RoleRemark string `json:"roleRemark"`

	IsUser            bool   `json:"isUser"`
	IsDepart          bool   `json:"isDepart"`
	DefaultSafeObject string `json:"defaultSafeObject"`

	Departs []string `json:"departs"`
	Users   []string `json:"users"`

	Operates []string `json:"operates"`
	Objects  []string `json:"objects"`
}

type SubSys struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ObjectGroup struct {
	SubSysId string `json:"subSysId"`

	Id   string `json:"id"`
	Name string `json:"name"`
}

type SafeObject struct {
	Id         string `json:"id"`
	SerialId   int    `json:"serialId"`
	ObjectType string `json:"objectType"`
	ParentPath string `json:"parentPath"`

	ObjectName   string `json:"objectName"`
	ObjectText   string `json:"objectText"`
	ObjectImage  string `json:"objectImage"`
	ObjectRemark string `json:"objectRemark"`

	Departs []string `json:"departs"`
	Users   []string `json:"users"`
	Roles   []string `json:"roles"`
}

type OperateGroup struct {
	SubSysId string `json:"subSysId"`

	Id   string `json:"id"`
	Name string `json:"name"`
}

type SafeOperate struct {
	Id          string `json:"id"`
	SerialId    int    `json:"serialId"`
	OperateCode string `json:"operateCode"`

	Subsys        string `json:"subsys"`
	OperateGroup  string `json:"operateGroup"`
	OperateName   string `json:"operateName"`
	OperateDetail string `json:"operateDetail"`
	OperateRemark string `json:"operateRemark"`

	IsAuth bool `json:"isAuth"`
	IsLog  bool `json:"isLog"`

	Departs []string `json:"departs"`
	Users   []string `json:"users"`
	Roles   []string `json:"roles"`
}

type RoleAllocInfo struct {
	RoleId string `json:"roleId"`
	UserId string `json:"userId"`
}

type ObjectAuthInfo struct {
	ObjectId string `json:"objectId"`
	RoleId   string `json:"roleId"`
}

type OperateAuthInfo struct {
	OperateId string `json:"operateId"`
	RoleId    string `json:"roleId"`
}

func (this *SafeUser) GetId() (userId string) {
	return this.Id
}

func (this *SafeSession) IsAdmin() bool {
	if this.GetUserId() == "userAdmin" {
		return true
	}
	if this.GetUserId() == "userDaemon" {
		return true
	}
	return false
}

func (this *SafeSession) GetId() (sessionId string) {
	return this.Id
}

func (this *SafeSession) GetUserId() (userId string) {
	return this.LoginUserId
}

func (this *SafeSession) GetUserName() (userId string) {
	return this.LoginUserName
}

func (this *SafeSession) KeepAlive() {
	this.KeepAliveTime = utils.StrTime()
	return
}

func (this *RoleAllocInfo) GetId() (allocId string) {
	allocId = this.RoleId + ":" + this.UserId
	return
}

func (this *ObjectAuthInfo) GetId() (authId string) {
	authId = this.ObjectId + ":" + this.RoleId
	return
}

func (this *OperateAuthInfo) GetId() (authId string) {
	authId = this.OperateId + ":" + this.RoleId
	return
}
