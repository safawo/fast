package safe

import (
	"fmt"
	"think/fast/mvc"
)

func init() {
	fmt.Println("Start Fast Safe")

	mvc.ProvideSessionValidater(&SafeSessionValidater{})

	mvc.Router("/think/fast/license/query", &QueryLicenseAction{})
	mvc.Router("/think/fast/license/verify", &VerifyLicenseAction{})
	mvc.Router("/think/fast/license/match", &MatchLicenseAction{})

	mvc.Router("/think/fast/license/import", &ImportLicenseAction{})
	mvc.Router("/think/fast/license/export", &ExportLicenseAction{})
	mvc.Router("/think/fast/license/clear", &ClearLicenseAction{})

	mvc.Router("/think/fast/safe/subSys/query", &QuerySafeSubSysAction{})

	mvc.Router("/think/fast/safe/role/query", &QuerySafeRoleAction{})
	mvc.Router("/think/fast/safe/role/create", &CreateSafeRoleAction{})
	mvc.Router("/think/fast/safe/role/delete", &DeleteSafeRoleAction{})
	mvc.Router("/think/fast/safe/role/change", &ChangeSafeRoleAction{})

	mvc.Router("/think/fast/safe/role/alloc", &AllocSafeRoleAction{})
	mvc.Router("/think/fast/safe/role/queryAlloc", &QueryAllocSafeRoleAction{})

	mvc.Router("/think/fast/safe/object/createObject", &CreateSafeObjectAction{})
	mvc.Router("/think/fast/safe/object/delObject", &DeleteSafeObjectAction{})
	mvc.Router("/think/fast/safe/object/queryObject", &QuerySafeObjectAction{})

	mvc.Router("/think/fast/safe/object/auth", &SafeObjectAuthAction{})
	mvc.Router("/think/fast/safe/object/queryAuth", &QuerySafeObjectAuthAction{})

	mvc.Router("/think/fast/safe/operate/queryOperate", &QuerySafeOperateAction{})
	mvc.Router("/think/fast/safe/operate/auth", &SafeOperateAuthAction{})
	mvc.Router("/think/fast/safe/operate/queryAuth", &QuerySafeOperateAuthAction{})

	mvc.Router("/think/fast/safe/user/query", &QueryUserAction{})
	mvc.Router("/think/fast/safe/user/create", &AddUserAction{})
	mvc.Router("/think/fast/safe/user/delete", &DeleteUserAction{})
	mvc.Router("/think/fast/safe/user/change", &ChangeUserAction{})

	mvc.Router("/think/fast/safe/user/modPassword", &ModPasswordAction{})
	mvc.Router("/think/fast/safe/user/lockUser", &LockUserAction{})
	mvc.Router("/think/fast/safe/user/unlockUser", &UnLockUserAction{})
	mvc.Router("/think/fast/safe/user/offlineUser", &OffLineUserAction{})

	mvc.Router("/think/fast/safe/self/queryMyInfo", &SelfQueryMyInfoAction{})
	mvc.Router("/think/fast/safe/self/modPassword", &SelfModPasswordAction{})
	mvc.Router("/think/fast/safe/self/modUserInfo", &SelfModUserInfoAction{})

	mvc.Router("/think/fast/safe/session/login", &LoginAction{})
	mvc.Router("/think/fast/safe/session/logout", &LogoutAction{})
	mvc.Router("/think/fast/safe/session/shakeHand", &ShakeHandAction{})
	mvc.Router("/think/fast/safe/session/awakeLogin", &AwakeLoginAction{})
	mvc.Router("/think/fast/safe/session/forceOffline", &ForceOfflineAction{})

	mvc.Router("/think/fast/safe/session/consultSession", &ConsultSessionAction{})

	mvc.Router("/think/fast/safe/log/queryOperateLog", &QueryOperateLogAction{})
	mvc.Router("/think/fast/safe/log/importOperateLog", &ImportOperateLogAction{})

	mvc.Router("/think/fast/safe/access/queryAccessObj", &QueryMyAccessObjectAction{})

	regMsg()

	buildLicense()
}
