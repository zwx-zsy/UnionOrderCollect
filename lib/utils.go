package lib

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
)

func FloatRound(f float64, n int) float64 {
	format := "%." + strconv.Itoa(n) + "f"
	res, _ := strconv.ParseFloat(fmt.Sprintf(format, f), 64)
	return res
}

//手机号码隐藏
func HidePhoneNumber(PhoneNumber string) (res string) {
	res = PhoneNumber[0:4] + "****" + PhoneNumber[7:]
	return res
}

//生成 MD5
func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//struct to string
func StructToString(s interface{}) string {
	bytes, _ := json.Marshal(s)
	return string(bytes)

}
