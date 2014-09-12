package safe

type OperateLog struct {
	UserName string `json:"userName"`

	Subsys       string `json:"subsys"`
	OperateGroup string `json:"operateGroup"`
	OperateName  string `json:"operateName"`

	OperateRet       string `json:"operateRet"`
	OperateRetDetail string `json:"operateRetDetail"`

	OperateObj     string `json:"operateObj"`
	OperateContent string `json:"operateContent"`

	UserIpAddress string `json:"userIpAddress"`
	UserHostName  string `json:"userHostName"`
	OperateTime   string `json:"operateTime"`

	LogType   string `json:"logType"`
	SerialNum string `json:"serialNum"`
}

func (this *OperateLog) IsEmpty() bool {
	if this.SerialNum == "" {
		return true
	}

	return false
}
