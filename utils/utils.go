package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

const (
	base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var (
	randSeed = 0
)

var coder = base64.NewEncoding(base64Table)

func base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}

func Password(password string) string {
	return PasswordBySeed(password, password)
}

func PasswordBySeed(password string, seed string) string {
	passwordSrc := EnCodeBase64(password)
	seedSrc := EnCodeBase64(passwordSrc) + Md5(seed) + EnCodeBase64("ujmiklop")

	passwordObj := EnCodeBase64(passwordSrc + seedSrc + "qazwsx")
	passwordObj = Md5(passwordObj + passwordSrc + seedSrc + "edcrfv")
	passwordObj = EnCodeBase64(passwordObj + passwordSrc + seedSrc + "tgbyhn")
	passwordObj = Md5(passwordObj + passwordSrc + seedSrc + "ujmiklop")

	return passwordObj
}

func Substr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

func RandStr() string {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//randValue := StrTimeTrim() + ":" + strconv.Itoa(r.Intn(100))
	randSeed++
	if randSeed > 9999 {
		randSeed = 1
	}
	randId := "00000" + strconv.Itoa(randSeed)
	randId = Substr(randId, (len(randId) - 5), 5)
	randId = StrTimeTrim() + randId

	return randId
}

func StrTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func StrTimeTrim() string {
	return time.Now().Format("20060102150405")
}

func VerifyErr(err error) {
	if err != nil {
		panic(err)
	}

}

func EnCodeBase64(srcData string) string {
	objData := string(base64Encode([]byte(srcData)))
	//fmt.Println("EnCodeBase64->srcData:", srcData, ",objData:", objData)
	return objData
}

func DeCodeBase64(src string) string {
	data, _ := base64Decode([]byte(src))
	return string(data)
}

func Md5(src string) string {
	h := md5.New()
	h.Write([]byte(src))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5Two(src string) string {
	return Md5(Md5(src))
}

func Md5Three(src string) string {
	return Md5Two(Md5(src))
}

func RootPath() string {
	file, _ := exec.LookPath(os.Args[0])
	dir, _ := path.Split(file)
	os.Chdir(dir)
	rootDir, _ := os.Getwd()

	return rootDir
}

func CachePath() string {
	cacheDir := RootPath() + "/cache"

	_, err := os.Stat(cacheDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(cacheDir, os.ModeDir)
	}

	return cacheDir
}

func TempPath() string {
	tempDir := CachePath() + "/temp"

	_, err := os.Stat(tempDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(tempDir, os.ModeDir)
	}

	return tempDir
}

func DumpPath() string {
	dumpDir := CachePath() + "/dump"

	_, err := os.Stat(dumpDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(dumpDir, os.ModeDir)
	}

	return dumpDir
}

func DownLoadPath() string {
	downloadDir := CachePath() + "/download"

	_, err := os.Stat(downloadDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(downloadDir, os.ModeDir)
	}

	return downloadDir
}

func UpLoadPath() string {
	uploadDir := CachePath() + "/upload"

	_, err := os.Stat(uploadDir)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(uploadDir, os.ModeDir)
	}

	return uploadDir
}
