package safe

import ()

type SysServerMetaInfo struct {
	ServerIp   string `json:"serverIp"`
	ServerHost string `json:"serverHost"`
	ServerPort int    `json:"serverPort"`
}
