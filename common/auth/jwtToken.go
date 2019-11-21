package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/chuxinplan/gin-mvc/common/config"
	"github.com/chuxinplan/gin-mvc/app/model"
	"strconv"
	"strings"
	"time"
)

//type Token struct {
//	Header    *HeaderData
//	PayLoad   *PayLoadData
//	Signature string
//}

type HeaderData struct {
	EncodeStyle string //加密方式
	Type        string //Token的类型
}

type PayLoadData struct {
	EndTime string //过期时间
	Name    string //用户名
	Id      string //用户Id
}

func GetPayLoad(token string) *PayLoadData {
	var payLoad *PayLoadData
	strs := strings.Split(token, ".")
	if len(strs) == 3 {
		pay, err := base64.URLEncoding.DecodeString(strs[1])
		if err != nil {
			panic(err)
		}
		pays := strings.Split(string(pay), ".")
		if len(pays) == 3 {
			payLoad = &PayLoadData{pays[0], pays[1], pays[2]}
		}
	}
	return payLoad
}

func GetToken(user *model.User) string {
	cfg := config.Get()
	header := &HeaderData{cfg.Jwt.EncodeMethod, "JWT"}
	endTime := strconv.FormatInt(time.Now().UnixNano()/1000000+cfg.Jwt.MaxEffectiveTime*86400, 10)
	payLoad := &PayLoadData{endTime, user.Username, strconv.FormatInt(user.Id, 10)}

	str := header.EncodeStyle + "." + header.Type
	Header := base64.StdEncoding.EncodeToString([]byte(str))

	str = payLoad.EndTime + "." + payLoad.Name + "." + payLoad.Id
	PayLoad := base64.StdEncoding.EncodeToString([]byte(str))

	Signature := ComputeHmac256(Header+"."+PayLoad, GetSecret(payLoad))

	return Header + "." + PayLoad + "." + Signature
}

func CheckToken(token string) bool {
	strs := strings.Split(token, ".")
	expect := ComputeHmac256(strs[0]+"."+strs[1], GetSecret(GetPayLoad(token)))
	if strs[2] == expect {
		return true
	}
	return false
}

func GetSecret(payLoad *PayLoadData) string {
	return ComputeHmac256(payLoad.EndTime+"."+payLoad.Name, "a1b1c2")
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
