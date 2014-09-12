package safe

import (
	"github.com/safawo/fast/msg"
)

const (
	MSG_LICENSE_INVALID = "license_invalid"
	MSG_LICENSE_VOID    = "license_void"

	MSG_LICENSE_NOENOUGH = "license_noenough"
	MSG_LICENSE_Expired  = "license_expired"

	MSG_USER_INVALID   = "user_invalid"
	MSG_PASSWORD_ERROR = "password_error"

	MSG_USERNAME_CLASH   = "userName_clash"
	MSG_EMPLOYEEID_CLASH = "employeeId_clash"
)

func regMsg() {
	msg.RegError(MSG_LICENSE_INVALID, "license invalid")
	msg.RegError(MSG_LICENSE_VOID, "license void")
	msg.RegError(MSG_LICENSE_NOENOUGH, "License not enough")
	msg.RegError(MSG_LICENSE_Expired, "License expired")

	msg.RegError(MSG_USER_INVALID, "user invalid")
	msg.RegError(MSG_PASSWORD_ERROR, "password error")

	msg.RegError(MSG_USERNAME_CLASH, "userName clash")
	msg.RegError(MSG_EMPLOYEEID_CLASH, "employeeId clash")

}
