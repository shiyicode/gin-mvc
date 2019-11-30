package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/chuxinplan/gin-mvc/common/errors"

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

func EncodeToken(username string, userId int64) (string, error) {
	cfg := config.Get()
	header := &header{cfg.Jwt.EncodeMethod, "JWT"}
	validTime, err := time.ParseDuration(cfg.Jwt.MaxEffectiveTime)
	if err != nil {
		return "", err
	}
	endTime := time.Now().Add(validTime)
	payLoad := &payLoad{
		EndTime:  endTime,
		Username: username,
		UserId:   userId,
	}

	headerStr, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	Header := base64.StdEncoding.EncodeToString(headerStr)

	payLoadStr, err := json.Marshal(payLoad)
	if err != nil {
		return "", err
	}
	PayLoad := base64.StdEncoding.EncodeToString(payLoadStr)

	secretStr, err := getSecret(payLoad)
	if err != nil {
		return "", err
	}
	Signature := computeHmac256(Header+"."+PayLoad, secretStr)

	token := base64.StdEncoding.EncodeToString([]byte(Header + "." + PayLoad + "." + Signature))
	return token, nil
}

func DecodeToken(token string) (*payLoad, *errors.Err) {
	data, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, errors.Warp(errors.ErrTokenValidation, err.Error())
	}
	token = string(data)
	if token == "" {
		return nil, errors.Warp(errors.ErrTokenNotFound)
	}

	strs := strings.Split(token, ".")
	if len(strs) != 3 {
		return nil, errors.Warp(errors.ErrTokenValidation)
	}
	payStr, err := base64.URLEncoding.DecodeString(strs[1])
	if err != nil {
		return nil, errors.Warp(errors.ErrTokenValidation, err.Error())
	}
	payLoad := &payLoad{}
	if err := json.Unmarshal(payStr, payLoad); err != nil {
		return nil, errors.Warp(errors.ErrTokenValidation, err.Error())
	}

	secretStr, err := getSecret(payLoad)
	if err != nil {
		return nil, errors.Warp(errors.ErrTokenValidation, err.Error())
	}
	expect := computeHmac256(strs[0]+"."+strs[1], secretStr)
	if strs[2] != expect {
		return nil, errors.Warp(errors.ErrTokenValidation)
	}
	if !payLoad.EndTime.After(time.Now()) {
		return nil, errors.Warp(errors.ErrTokenExpire)
	}

	return payLoad, nil
}

func computeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func getSecret(payLoad *payLoad) (string, error) {
	payStr, err := json.Marshal(payLoad)
	if err != nil {
		return "", err
	}
	return computeHmac256(string(payStr), "a1b1c2"), nil
}
