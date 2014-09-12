package utils

type FastSysPara struct {
	Catalog    string `json:"catalog"`    //参数目录分组
	ParaName   string `json:"paraName"`   //参数名称
	ParaValue  string `json:"paraValue"`  //参数值
	ParaType   string `json:"paraType"`   //参数类型
	ParaRemark string `json:"paraRemark"` //参数说明
}
