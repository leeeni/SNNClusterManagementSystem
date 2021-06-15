package service

import (
	"fmt"
	"snns_srv/session"

	"github.com/kataras/iris/v12"
	"gopkg.in/mgo.v2/bson"
)

//SetLoginSession : Login 生成 Session
func SetLoginSession(ctx iris.Context, uid bson.ObjectId) {
	sess := session.Sess.Start(ctx)
	sess.Set("auth", true)
	sess.Set("uid", uid.Hex())
	GetCookieOptions := session.Sess.GetCookieOptions()
	fmt.Println(GetCookieOptions[0])
}

//SetLogoutSession ：Logout 删除 Session
func SetLogoutSession(ctx iris.Context) {
	sess := session.Sess.Start(ctx)
	sess.Destroy()
}
