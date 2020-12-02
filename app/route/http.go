package route

import (
	"crypto/md5"
	"encoding/hex"
	_ "encoding/json"
	_ "errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"newhope/app/store/mongodb"
	_struct "newhope/app/struct"
)

func SetupHttp(r *gin.Engine) {
	v1 := r.Group("/server")
	{
		v1.GET("/login", HandleUserLogin)
		v1.GET("/logout", HandleUserLogout)
		v1.GET("/entry", HandleFormEntry)
		v1.GET("/query", HandleFormQuery)
		v1.GET("/register", HandleRegister)
		v1.GET("/test", HandleTest)
	}

	v2 := r.Group("/server")
	{
		v2.POST("/login", HandleUserLogin)
		v2.POST("/logout", HandleUserLogout)
		v2.POST("/entry", HandleFormEntry)
		v2.POST("/query", HandleFormQuery)
		v2.POST("/register", HandleRegister)
		v2.POST("/test", HandleTest)
	}

	v3 := r.Group("/ui")
	{
		v3.StaticFS("/", http.Dir("./ui"))
	}
}

func HandleUserLogin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	result, code := userLogin(c.Request.FormValue("account"), c.Request.FormValue("password"))
	if code == http.StatusOK {
		session := sessions.Default(c)
		session.Set("account", _struct.UserSession{Account: c.Request.FormValue("account")})
		session.Save()
	}
	c.String(code, result)
}

func HandleUserLogout(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	auth, _ := checkSession(c)
	if !auth {
		c.String(http.StatusOK, "Not logged in")
		return
	}
	session := sessions.Default(c)
	session.Delete("account")
	session.Save()
	c.String(http.StatusOK, "success")
}

func HandleFormEntry(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	auth, account := checkSession(c)
	if !auth {
		c.String(http.StatusOK, "Not logged in")
		return
	}
	id := primitive.NewObjectID()
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"company", c.Request.FormValue("company")},
			{"jobs", c.Request.FormValue("jobs")},
			{"working", c.Request.FormValue("working")},
			{"leader", c.Request.FormValue("leader")},
			{"time_limit", c.Request.FormValue("time_limit")},
			{"exist_problem", c.Request.FormValue("exist_problem")},
			{"problem_type", c.Request.FormValue("problem_type")},
			{"account", account.Account},
		}},
	}
	mongodb.UpdateRecord(filter, update, true)
	c.String(http.StatusOK, "success")
}

func HandleFormQuery(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	auth, _ := checkSession(c)
	if !auth {
		c.String(http.StatusOK, "Not logged in")
		return
	}

	response,_ := mongodb.QueryAllRecord()
	c.String(http.StatusOK, response)
}

func HandleRegister(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id := getMd5String("admin")
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"account", "admin"},
			{"password", "123456"},
			{"group", "admin"},
			{"level", 0},
		}},
	}
	mongodb.UpdateAccount(filter, update, true)
	c.String(http.StatusOK, string("success"))
}

func HandleTest(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, string("hello world!"))
}

func userLogin(account string, password string) (string, int) {
	res, err := mongodb.QueryConditionAccount2json(bson.D{{"account", account},
		{"password", password}})
	if res == "[]" || err != nil {
		return "failure", http.StatusNotFound
	}
	return "success", http.StatusOK
}

func checkSession(c *gin.Context) (bool, _struct.UserSession) {
	session := sessions.Default(c)
	user := session.Get("account")
	if user == nil {
		return false, _struct.UserSession{}
	}
	return true, user.(_struct.UserSession)
}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
