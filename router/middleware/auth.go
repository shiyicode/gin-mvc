package middleware

import (
	"github.com/gin-gonic/gin"
)

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

func MustGetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}

// func GetCurrentId(token string) int64 {
// 	userId, err := strconv.ParseInt(auth.getPayLoad(token).Id, 10, 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return userId
// }
//
// func GetUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if strings.HasPrefix(c.Request.URL.Path, "/") {
//
// 		}
// 		c.Next()
// 	}
// }
//
// func MustGetUser(checkPath string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if strings.HasPrefix(c.Request.URL.Path, checkPath) {
// 			token, _ := c.Cookie("token")
// 			userId, err := checkLogin(token)
// 			if err != nil {
// 				log.Warningf("Check User Fail!\n")
// 				c.JSON(http.StatusForbidden, "权限检验失败，请重新登录!")
// 				c.Abort()
// 			}
// 			c.Set("userId", userId)
// 		}
// 		c.Next()
// 	}
// }

// input token，get decode userId
// func checkLogin(token string) (int64, error) {
// 	data, err := base64.URLEncoding.DecodeString(token)
// 	if err != nil {
// 		return 0, err
// 	}
//
// 	token = string(data)
// 	if token == "" {
// 		return 0, errors.New("token串为空!")
// 	}
//
// 	nowTime := time.Now().UnixNano() / 1000000
// 	endTime, err := strconv.ParseInt(auth.getPayLoad(token).EndTime, 10, 64)
// 	if err != nil {
// 		return 0, err
// 	}
// 	if nowTime <= endTime && auth.checkToken(token) {
// 		return GetCurrentId(token), nil
// 	}
//
// 	return 0, errors.New("用户会话有误或已失效！")
// }
