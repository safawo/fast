package safe

import (
	"encoding/json"
	"fmt"
	"github.com/safawo/fast/comm"
	"github.com/safawo/fast/utils"
	"os"
)

func buildLicense() {
	buildDeveloper()
}

func buildDeveloper() {
	licenseInfo := &FastLicenseInfo{}
	licenseInfo.Reset()

	licenseInfo.Distributor = "liuyong"
	licenseInfo.DistributeId = "liuyong1024"

	licenseInfo.LicenseId = "develop"
	licenseInfo.LicenseType = LICENSE_TYPE_DEVELOP

	licenseInfo.CustomerId = "FFFFFFFFFFFF"
	licenseInfo.CustomerName = "developer"

	licenseInfo.BindNetCard = comm.NULL_STR
	licenseInfo.BindDisk = comm.NULL_STR
	licenseInfo.BindCPU = comm.NULL_STR

	licenseInfo.ExpiryDate = "2015-09-15 23:59:59"
	licenseInfo.Remark = comm.NULL_STR

	licenseInfo.AllowInt["clientOnlineLimit"] = 5
	licenseInfo.AllowInt["lockLimit"] = 10000
	licenseInfo.AllowInt["userGroupLimit"] = 256
	licenseInfo.AllowInt["timeTableLimit"] = 8
	licenseInfo.AllowInt["commonDoorLimit"] = 53

	givePads := []string{}
	givePads = append(givePads, "superpad")
	givePads = append(givePads, "353723050291384")
	givePads = append(givePads, "869274012114451")
	givePads = append(givePads, "861519010479951")
	givePads = append(givePads, "BX9034PWKC")
	givePads = append(givePads, "2F32000200000001")
	givePads = append(givePads, "021YHB2133052646")

	licenseInfo.AllowStrs["givePads"] = givePads

	byteLicense, err := json.MarshalIndent(licenseInfo, "", "  ")
	utils.VerifyErr(err)

	licenseContent := utils.EnCodeBase64(string(byteLicense))
	licenseContent = utils.EnCodeBase64(licenseContent)
	licenseContent = utils.EnCodeBase64(licenseContent)

	fmt.Println(licenseContent)

	licenseFile := "developlicense.lic"
	fout, err := os.Create(licenseFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(licenseFile, err)
		return
	}

	fout.Write([]byte(licenseContent))

}
