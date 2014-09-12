package safe

import (
	"fmt"
	"github.com/safawo/fast/mvc"
)

func init() {
	fmt.Println("Start Fast Safe")

	mvc.ProvideSessionValidater(&SafeSessionValidater{})

	mvc.Router("/github.com/safawo/fast/license/query", &QueryLicenseAction{})
	mvc.Router("/github.com/safawo/fast/license/verify", &VerifyLicenseAction{})
	mvc.Router("/github.com/safawo/fast/license/match", &MatchLicenseAction{})

	mvc.Router("/github.com/safawo/fast/license/import", &ImportLicenseAction{})
	mvc.Router("/github.com/safawo/fast/license/export", &ExportLicenseAction{})
	mvc.Router("/github.com/safawo/fast/license/clear", &ClearLicenseAction{})

	mvc.Router("/github.com/safawo/fast/safe/subSys/query", &QuerySafeSubSysAction{})

	mvc.Router("/github.com/safawo/fast/safe/role/query", &QuerySafeRoleAction{})
	mvc.Router("/github.com/safawo/fast/safe/role/create", &CreateSafeRoleAction{})
	mvc.Router("/github.com/safawo/fast/safe/role/delete", &DeleteSafeRoleAction{})
	mvc.Router("/github.com/safawo/fast/safe/role/change", &ChangeSafeRoleAction{})

	mvc.Router("/github.com/safawo/fast/safe/role/alloc", &AllocSafeRoleAction{})
	mvc.Router("/github.com/safawo/fast/safe/role/queryAlloc", &QueryAllocSafeRoleAction{})

	mvc.Router("/github.com/safawo/fast/safe/object/createObject", &CreateSafeObjectAction{})
	mvc.Router("/github.com/safawo/fast/safe/object/delObject", &DeleteSafeObjectAction{})
	mvc.Router("/github.com/safawo/fast/safe/object/queryObject", &QuerySafeObjectAction{})

	mvc.Router("/github.com/safawo/fast/safe/object/auth", &SafeObjectAuthAction{})
	mvc.Router("/github.com/safawo/fast/safe/object/queryAuth", &QuerySafeObjectAuthAction{})

	mvc.Router("/github.com/safawo/fast/safe/operate/queryOperate", &QuerySafeOperateAction{})
	mvc.Router("/github.com/safawo/fast/safe/operate/auth", &SafeOperateAuthAction{})
	mvc.Router("/github.com/safawo/fast/safe/operate/queryAuth", &QuerySafeOperateAuthAction{})

	mvc.Router("/github.com/safawo/fast/safe/user/query", &QueryUserAction{})
	mvc.Router("/github.com/safawo/fast/safe/user/create", &AddUserAction{})
	mvc.Router("/github.com/safawo/fast/safe/user/delete", &DeleteUserAction{})
	mvc.Router("/github.com/safawo/fast/safe/user/change", &ChangeUserAction{})

	mvc.Router("/github.com/safawo/fast/safe/user/modPassword", &ModPasswordAction{})
	mvc.Router("/github.com/safawo/fast/safe/user/lockUser", &LockUserAction{})
	mvc.Router("/github.com/safawo/fast/safe/user/unlockUser", &UnLockUserAction{})
	mvc.Router("/github.com/safawo/fast/safe/user/offlineUser", &OffLineUserAction{})

	mvc.Router("/github.com/safawo/fast/safe/self/queryMyInfo", &SelfQueryMyInfoAction{})
	mvc.Router("/github.com/safawo/fast/safe/self/modPassword", &SelfModPasswordAction{})
	mvc.Router("/github.com/safawo/fast/safe/self/modUserInfo", &SelfModUserInfoAction{})

	mvc.Router("/github.com/safawo/fast/safe/session/login", &LoginAction{})
	mvc.Router("/github.com/safawo/fast/safe/session/logout", &LogoutAction{})
	mvc.Router("/github.com/safawo/fast/safe/session/shakeHand", &ShakeHandAction{})
	mvc.Router("/github.com/safawo/fast/safe/session/awakeLogin", &AwakeLoginAction{})
	mvc.Router("/github.com/safawo/fast/safe/session/forceOffline", &ForceOfflineAction{})

	mvc.Router("/github.com/safawo/fast/safe/session/consultSession", &ConsultSessionAction{})

	mvc.Router("/github.com/safawo/fast/safe/log/queryOperateLog", &QueryOperateLogAction{})
	mvc.Router("/github.com/safawo/fast/safe/log/importOperateLog", &ImportOperateLogAction{})

	mvc.Router("/github.com/safawo/fast/safe/access/queryAccessObj", &QueryMyAccessObjectAction{})

	regMsg()

	buildLicense()
}
