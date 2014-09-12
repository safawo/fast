package safe

import ()

type SysServerMetaAgent struct {
}

var (
	DefaultSysServerMetaAgent = &SysServerMetaAgent{}
)

func GetSysServerMetaData() SysServerMetaInfo {
	return getSysServerMetaAgent().getDatas()
}

func getSysServerMetaAgent() *SysServerMetaAgent {
	return DefaultSysServerMetaAgent
}

func (this *SysServerMetaAgent) getDatas() SysServerMetaInfo {
	metaInfo := SysServerMetaInfo{}
	metaInfo.ServerHost = "HotelServer"
	metaInfo.ServerIp = "HotelServer"
	metaInfo.ServerPort = 65533
	return metaInfo
}
