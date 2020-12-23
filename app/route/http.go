package route

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	_ "encoding/json"
	_ "errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"newhope/app/store/mongodb"
	_struct "newhope/app/struct"
	"newhope/config"
	"os"
)

func SetupHttp(r *gin.Engine) {
	v1 := r.Group("/server")
	{
		v1.GET("/login", HandleUserLogin)
		v1.GET("/logout", HandleUserLogout)
		v1.GET("/entryform", HandleEntryForm)
		v1.GET("/queryform", HandleQueryForm)
		v1.GET("/search", HandleSearch)
		v1.GET("/deleteform", HandleDeleteForm)
		v1.GET("/entryterms", HandleEntryTerms)
		v1.GET("/queryterms", HandleQueryTerms)
		v1.GET("/deleteterms", HandleDeleteTerms)
		v1.GET("/register", HandleRegister)
		v1.GET("/unregister", HandleUnregister)
		v1.GET("/modifypwd", HandleModifyPassword)
		v1.GET("/upload", HandleUpload)
		v1.GET("/download", HandleDownload)
		v1.GET("/test", HandleTest)
	}

	v2 := r.Group("/server")
	{
		v2.POST("/login", HandleUserLogin)
		v2.POST("/logout", HandleUserLogout)
		v2.POST("/entryform", HandleEntryForm)
		v2.POST("/queryform", HandleQueryForm)
		v2.POST("/search", HandleSearch)
		v2.POST("/deleteform", HandleDeleteForm)
		v2.POST("/entryterms", HandleEntryTerms)
		v2.POST("/queryterms", HandleQueryTerms)
		v2.POST("/deleteterms", HandleDeleteTerms)
		v2.POST("/register", HandleRegister)
		v2.POST("/unregister", HandleUnregister)
		v2.POST("/modifypwd", HandleModifyPassword)
		v2.POST("/upload", HandleUpload)
		v2.POST("/download", HandleDownload)
		v2.POST("/test", HandleTest)
	}

	v3 := r.Group("/ui")
	{
		v3.StaticFS("/", http.Dir("./ui"))
	}
}

func HandleUserLogin(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	account, err := mongodb.QueryConditionAccount(bson.D{{"account", c.Request.FormValue("account")},
		{"password", c.Request.FormValue("password")}})
	if account == (_struct.AccountInfo{}) || err != nil {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "密码不正确"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusNotFound, string(jsons))
		return
	}
	account.Password = "******"
	body, _ := json.Marshal(account)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "登录成功"
	responseMessage.Body = string(body)
	jsons, _ := json.Marshal(responseMessage)
	session := sessions.Default(c)
	session.Set("account", _struct.UserSession{Account: account.Account,
		Group: account.Group,Level: account.Level})
	session.Save()
	c.String(http.StatusOK, string(jsons))
}

func HandleUserLogout(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, _ := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	session := sessions.Default(c)
	session.Delete("account")
	session.Save()

	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "账户登出成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleEntryForm(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, account := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	var id primitive.ObjectID
	var err error
	if len(c.Request.FormValue("id")) == 0 {
		id = primitive.NewObjectID()
	}else {
		id, err = primitive.ObjectIDFromHex(c.Request.FormValue("id"))
		if err != nil {
			id = primitive.NewObjectID()
		}
	}
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"company", c.Request.FormValue("company")},
			{"jobs", c.Request.FormValue("jobs")},
			{"working", c.Request.FormValue("working")},
			{"leader", c.Request.FormValue("leader")},
			{"date", c.Request.FormValue("date")},
			{"problem", c.Request.FormValue("problem")},
			{"filename", c.Request.FormValue("filename")},
			{"score", c.Request.FormValue("score")},
			{"type", c.Request.FormValue("type")},
			{"account", account.Account},
		}},
	}
	mongodb.UpdateRecord(filter, update, true)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "操作成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleQueryForm(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, account := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	var response string
	if account.Level == "0"{
		response,_ = mongodb.QueryAllRecord()
	}else {
		response,_ = mongodb.QueryConditionRecord(bson.D{{"account", account.Account}})
	}
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "查询成功"
	responseMessage.Body = response
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleSearch(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, account := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}

	if account.Level != "0"{
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "该账户没有权限"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	var response string
	response,_ = mongodb.QueryConditionRecord(bson.D{{"leader", c.Request.FormValue("leader")},
		{"date",primitive.Regex{Pattern: c.Request.FormValue("date")}}})
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "查询成功"
	responseMessage.Body = response
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleDeleteForm(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, _ := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	var id primitive.ObjectID
	var err error
	if len(c.Request.FormValue("id")) == 0 {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "删除失败，id不正确"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}else {
		id, err = primitive.ObjectIDFromHex(c.Request.FormValue("id"))
		if err != nil {
			responseMessage.ErrCode = -1
			responseMessage.ErrMessage = "删除失败，id不正确"
			jsons, _ := json.Marshal(responseMessage)
			c.String(http.StatusOK, string(jsons))
			return
		}
	}
	filter := bson.M{"_id": id}
	mongodb.DeleteRecord(filter)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "删除成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleEntryTerms(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, _ := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	var id primitive.ObjectID
	var err error
	if len(c.Request.FormValue("id")) == 0 {
		id = primitive.NewObjectID()
	}else {
		id, err = primitive.ObjectIDFromHex(c.Request.FormValue("id"))
		if err != nil {
			id = primitive.NewObjectID()
		}
	}
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"terms", c.Request.FormValue("terms")},
			{"score", c.Request.FormValue("score")},
		}},
	}
	mongodb.UpdateTerms(filter, update, true)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "操作成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleQueryTerms(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, _ := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	response,_ := mongodb.QueryAllTerms()
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "查询成功"
	responseMessage.Body = response
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleDeleteTerms(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, _ := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	var id primitive.ObjectID
	var err error
	if len(c.Request.FormValue("id")) == 0 {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "删除失败，id不正确"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}else {
		id, err = primitive.ObjectIDFromHex(c.Request.FormValue("id"))
		if err != nil {
			responseMessage.ErrCode = -1
			responseMessage.ErrMessage = "删除失败，id不正确"
			jsons, _ := json.Marshal(responseMessage)
			c.String(http.StatusOK, string(jsons))
			return
		}
	}
	filter := bson.D{{"_id", id}}
	mongodb.DeleteTerms(filter)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "删除成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleRegister(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	id := getMd5String(c.Request.FormValue("account"))
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"account", c.Request.FormValue("account")},
			{"password", c.Request.FormValue("password")},
			{"group", c.Request.FormValue("group")},
			{"level", c.Request.FormValue("level")},
		}},
	}
	mongodb.UpdateAccount(filter, update, true)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "账户注册成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleUnregister(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	auth, _ := checkSession(c)
	if !auth {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户未登录"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	id := c.Request.FormValue("id")
	if len(id) == 0 {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "删除失败，id不正确"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	filter :=  bson.M{"_id": id}
	mongodb.DeleteAccount(filter)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "删除成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleModifyPassword(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	responseMessage := _struct.ResponseMessage{}
	if len(c.Request.FormValue("newpassword")) == 0 {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "新密码为空"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	id := getMd5String(c.Request.FormValue("account"))
	acc,_ := mongodb.QueryConditionAccount(bson.D{{"_id", id},
		{"password",c.Request.FormValue("oldpassword")}})
	if len(acc.Account) == 0 {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "账户或密码不存在"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	filter := bson.M{"_id": id}
	update := bson.D{
		{"$set", bson.D{
			{"_id", id},
			{"account", c.Request.FormValue("account")},
			{"password", c.Request.FormValue("newpassword")},
		}},
	}
	mongodb.UpdateAccount(filter, update, false)
	responseMessage.ErrCode = 0
	responseMessage.ErrMessage = "密码修改成功"
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleUpload(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	form, _ := c.MultipartForm()
	files := form.File["upload"]
	path := config.GetImagesPath()
	for _, file := range files {
		c.SaveUploadedFile(file, path + "/" + file.Filename)
	}
	body,_ := json.Marshal(len(files))
	responseMessage := _struct.ResponseMessage{}
	if (len(files) == 0) {
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = "上传失败"
	}else {
		responseMessage.ErrCode = 0
		responseMessage.ErrMessage = "上传成功"
	}
	responseMessage.Body = string(body)
	jsons, _ := json.Marshal(responseMessage)
	c.String(http.StatusOK, string(jsons))
}

func HandleDownload(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	filename := c.Request.FormValue("filename")
	filepath := config.GetImagesPath() + "/" + filename
	if !fileIsExist(filepath){
		responseMessage := _struct.ResponseMessage{}
		responseMessage.ErrCode = -1
		responseMessage.ErrMessage = filename + "文件不存在"
		jsons, _ := json.Marshal(responseMessage)
		c.String(http.StatusOK, string(jsons))
		return
	}
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(filepath)
}

func HandleTest(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.String(http.StatusOK, string("hello world!"))
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

func fileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}