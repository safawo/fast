package backup

import (
	"fmt"
	"think/fast/mvc"
)

func init() {
	fmt.Println("Start Fast Backup")

	mvc.Router("/think/fast/backup/notifyBackup", &NotifyBackupAction{})
	mvc.Router("/think/fast/backup/notifyRestore", &NotifyRestoreAction{})

	mvc.Router("/think/fast/backup/backupSystem", &BackupSystemAction{})
	mvc.Router("/think/fast/backup/restoreSystem", &RestoreSystemAction{})

	mvc.Router("/think/fast/backup/getBackups", &QueryBackupAction{})
	mvc.Router("/think/fast/backup/delBackup", &DeleteBackupAction{})
}
