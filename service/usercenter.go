package service

import (
	"github.com/joho/godotenv"
	"net/url"
	"os"
)

// token的respond格式
type TokenData struct {
	AccessToken      string `json:"access_token"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string  `json:"token_type"`
	ExpiresIn        float64  `json:"expires_in"`
}

// 用户登录获取token
func GetToken(username string, password string) (TokenData, error) {

	var tokenRespond TokenData

	// 获取配置文件信息
	_ = godotenv.Load()
	//添加post的body内容
	data := make(url.Values)
	data["client_id"] = []string{os.Getenv("CLIENT_ID")}
	data["client_secret"] = []string{os.Getenv("CLIENT_PASSWORD")}
	data["grant_type"] = []string{"password"}
	data["username"] = []string{username}
	data["password"] = []string{password}
	data["scope"] = []string{"read"}
	// 发送Http
	respondStruct,err := HttpPost("http://192.168.202.71:31441/oauth/token",data,"application/x-www-form-urlencoded")
	if respondStruct["error"] != nil{
		return tokenRespond, err
	}
	// 整合成token响应对象
	tokenRespond.AccessToken = respondStruct["access_token"].(string)
	tokenRespond.RefreshToken = respondStruct["refresh_token"].(string)
	tokenRespond.TokenType = respondStruct["token_type"].(string)
	tokenRespond.ExpiresIn = respondStruct["expires_in"].(float64)

	return tokenRespond, err
}

// 响应格式
type UserInfoResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    UserInfo `json:"data"`
}
// userInfo的respond格式
type UserInfo struct {
	UserId         int         `json:"userId"`
	Uid            string         `json:"uid"`
	Username       string         `json:"username"`
	RealName       string         `json:"realName"`
	Tutor          string         `json:"tutor"`
	Avatar         string         `json:"avatar"`
	Phone          string         `json:"phone"`
	Email          string         `json:"email"`
	Sex            int            `json:"sex"`
	Locked         int            `json:"locked"`
	Ctime          int            `json:"ctime"`
	Description    string         `json:"description"`
	OrganizationId string         `json:"organizationId"`
	InnerAccount   string         `json:"innerAccount"`
}

// 组织信息
type Organization struct {
	OrganizationId string `json:"organizationId"`
	Pid            string `json:"pid"`
	Ids            string `json:"ids"`
	Names          string `json:"names"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Ctime          string `json:"ctime"`
}

func GetUserInfo(token string)  (UserInfo, error)  {

	var UserInfoRespond UserInfo
	var basicResponse UserInfoResponse

	// 发送Http
	respondStruct,err := HttpGet("http://192.168.202.71:31441/api/user/",token)
	if respondStruct["error"] != nil{
		return UserInfoRespond, err
	}
	// 整合成整体响应对象
	basicResponse.Code = int(respondStruct["code"].(float64))
	basicResponse.Message = respondStruct["message"].(string)
	data := respondStruct["data"].(map[string]interface{})
	// 整合成UserInfo响应对象
	UserInfoRespond.UserId = int(data["userId"].(float64))
	UserInfoRespond.Uid = data["uid"].(string)
	UserInfoRespond.Username = data["username"].(string)
	if data["realName"] == nil{
		UserInfoRespond.RealName = ""
	}else {
		UserInfoRespond.RealName = data["realName"].(string)
	}
	if data["tutor"] == nil{
		UserInfoRespond.Tutor = ""
	}else {
		UserInfoRespond.Tutor = data["tutor"].(string)
	}
	if data["avatar"] == nil{
		UserInfoRespond.Avatar = ""
	}else {
		UserInfoRespond.Avatar = data["avatar"].(string)
	}
	if data["phone"] == nil{
		UserInfoRespond.Phone = ""
	}else {
		UserInfoRespond.Phone = data["phone"].(string)
	}
	if data["email"] == nil{
		UserInfoRespond.Email = ""
	}else {
		UserInfoRespond.Email = data["email"].(string)
	}
	UserInfoRespond.Sex = int(data["sex"].(float64))
	UserInfoRespond.Locked = int(data["locked"].(float64))
	UserInfoRespond.Ctime = int(data["ctime"].(float64))
	if data["description"] == nil{
		UserInfoRespond.Description = ""
	}else {
		UserInfoRespond.Description = data["description"].(string)
	}
	if data["innerAccount"] == nil{
		UserInfoRespond.InnerAccount = ""
	}else {
		UserInfoRespond.InnerAccount = data["innerAccount"].(string)
	}
	if data["organizationId"] == nil{
		UserInfoRespond.OrganizationId = ""
	}else {
		UserInfoRespond.OrganizationId = data["organizationId"].(string)
	}
	return UserInfoRespond, err
}