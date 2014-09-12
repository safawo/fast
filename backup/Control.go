package backup

import (
	"fmt"
	"github.com/safawo/fast/mvc"
)

func init() {
	fmt.Println("Start Fast Backup")

	mvc.Router("/github.com/safawo/fast/backup/notifyBackup", &NotifyBackupAction{})
	mvc.Router("/github.com/safawo/fast/backup/notifyRestore", &NotifyRestoreAction{})

	mvc.Router("/github.com/safawo/fast/backup/backupSystem", &BackupSystemAction{})
	mvc.Router("/github.com/safawo/fast/backup/restoreSystem", &RestoreSystemAction{})

	mvc.Router("/github.com/safawo/fast/backup/getBackups", &QueryBackupAction{})
	mvc.Router("/github.com/safawo/fast/backup/delBackup", &DeleteBackupAction{})
}
