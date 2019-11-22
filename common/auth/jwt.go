package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/chuxinplan/gin-mvc/common/config"
)

type header struct {
	EncodeStyle string // 加密方式
	Type        string // Token的类型
}

type payLoad struct {
	EndTime  time.Time // 过期时间
	Username string    // 用户名
	UserId   int64     // 用户Id
}

func EncodeToken(username string, userId int64) string {
	cfg := config.Get()
	header := &header{cfg.Jwt.EncodeMethod, "JWT"}
	endTime := time.Now().UnixNano()/1000000 + cfg.Jwt.MaxEffectiveTime*86400 // time转换的优雅一些
	payLoad := &payLoad{
		EndTime:  endTime,
		Username: username,
		UserId:   userId,
	}

	str := header.EncodeStyle + "." + header.Type
	Header := base64.StdEncoding.EncodeToString([]byte(str))

	str = payLoad.EndTime.String() + "." + payLoad.Username + "." + strconv.FormatInt(payLoad.UserId, 10)
	PayLoad := base64.StdEncoding.EncodeToString([]byte(str))

	Signature := computeHmac256(Header+"."+PayLoad, getSecret(payLoad))

	return Header + "." + PayLoad + "." + Signature
}

func DecodeToken(token string) (bool, *payLoad) {
	strs := strings.Split(token, ".")
	if len(strs) != 3 {
		return false, nil
	}
	payStr, err := base64.URLEncoding.DecodeString(strs[1])
	if err != nil {
		return false, nil
	}
	payLoad := &payLoad{} // 将payStr转为pay结构体 json!!!

	expect := computeHmac256(strs[0]+"."+strs[1], getSecret(payLoad))
	if strs[2] == expect && payLoad.EndTime.After(time.Now()) {
		return true, payLoad
	}
	return false, nil
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getSecret(payLoad *payLoad) string {
	payStr := "" // json序列化
	return computeHmac256(payStr, "a1b1c2")
}
