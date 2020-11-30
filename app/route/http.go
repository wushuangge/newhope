package route

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	_ "encoding/json"
	_ "errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
	log "github.com/sirupsen/logrus"
	"net/http"
	_struct "newhope/app/struct"
	"newhope/config"
)

func SetupHttp(r *gin.Engine) {
	v1 := r.Group("/rget")
	{
		v1.GET("/task", HandleTask)
		v1.GET("/test", HandleTest)
	}

	v2 := r.Group("/rpost")
	{
		v2.POST("/task", HandleTask)
		v2.POST("/test", HandleTest)
	}

	v3 := r.Group("/ui")
	{
		v3.StaticFS("/", http.Dir("./ui"))
	}
}

func HandleTask(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	auth, _ := checkSession(c)
	if !auth {
		c.String(http.StatusOK, "该用户未登录")
		return
	}
	switch c.Request.FormValue("operation") {
	case "UserLogin":
		result, code := userLogin(c.Request.FormValue("username"), c.Request.FormValue("password"))
		if code == http.StatusOK {
			session := sessions.Default(c)
			session.Set("user", _struct.UserSession{c.Request.FormValue("username")})
			session.Save()
		}
		c.String(code, result)
		break
	case "UserLogout":
		session := sessions.Default(c)
		session.Delete("user")
		session.Save()
		c.String(http.StatusOK, "success")
		break
	default:
		log.Error("default!!! operation is ", c.Request.FormValue("operation"))
	}
}

func HandleTest(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, string("hello world!"))
}

func userLogin(username string, password string) (string, int) {
	return string("success"), http.StatusOK
}

func checkSession(c *gin.Context) (bool, _struct.UserSession) {
	session := sessions.Default(c)
	user := session.Get("user")
	if user == nil {
		if c.Request.FormValue("operation") == "UserLogin" ||
			c.Request.FormValue("operation") == "GetServiceUrl" {
			return true, _struct.UserSession{}
		} else {
			return false, _struct.UserSession{}
		}
	}
	return true, user.(_struct.UserSession)
}

func checkUser(username string, password string) bool {
	var userIsExist = false
	request := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	resp, _, errs := request.Get(config.GetAuthUrl()).SetBasicAuth(username, password).Set("User-Agent", "ftp").End()
	if len(errs) <= 0 && resp.StatusCode == 200 {
		userIsExist = true
	}
	return userIsExist
}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
