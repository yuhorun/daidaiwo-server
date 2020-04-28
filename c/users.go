package c

import (
	"github.com/VividCortex/mysqlerr"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"image/png"
	"log"
	"net/http"
	"regexp"
	"server/d"
	"server/middleware"
	"server/pkg/common"
	"server/pkg/e"
	"server/pkg/util"
)

//Shouldxxx和bindxxx区别就是bindxxx会在head中添加400的返回信息，而Shouldxxx不会
//用户登录
func GetUserInfo(c *gin.Context) {
	uid := c.Keys["uid"].(int64)
	user, err := d.GetUserInfo(uid)
	if err != nil {
		common.CJSON(c, http.StatusOK, e.FAIL)
		return
	}

	user.Upwd = ""
	user.Pnumber = ""
	c.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": user,
	})
}

func Login(c *gin.Context) {

	var user d.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		common.CJSON(c, http.StatusOK, e.INVALID_PARAMS)
		return
	}

	uid, err := d.VerifyPwd(user.Pnumber, user.Upwd)
	if err != nil {
		common.CJSON(c, http.StatusOK, e.ERROR_USERNAME_OR_PAWWORD)
		return
	}

	jwtstring, err := middleware.GenerateToken(uid)
	if err != nil {
		log.Println("Faield to GenerateJwtString")
	}
	c.SetCookie("jwtstring", jwtstring, 365*24*3600, "/", "localhost", false, false)
	common.CJSON(c, http.StatusOK, e.LOGIN_SUCCESS)

}

func Logout(c *gin.Context) {
	c.SetCookie("jwtstring", "", -1, "/", "localhost", false, false)
	common.CJSON(c, http.StatusOK, e.LOGOUT_SUCCESS)
}

func GetverifyCodeImage(c *gin.Context) {

	image_id := c.Query("image_id")

	if isOk, _ := regexp.MatchString("^[a-zA-Z0-9][a-zA-Z0-9]{10}$", image_id); !isOk {
		common.CJSON(c, http.StatusOK, e.INVALID_PARAMS)
		return
	} else {
		//生存随即的验证码图片
		caps := util.DefalutCaptcha()

		img, str := caps.Create(4, captcha.NUM)

		//发送图片，客户端接受的即为一个图片： flutter: codeImage = Image.network(api)
		err := png.Encode(c.Writer, img)
		if err != nil {
			log.Println(err)
		}

		err = d.SetVerifyCode("image_code_"+image_id, str, 60*60*10)
		if err != nil {
			log.Println(err)
		}
	}
	common.CJSON(c, http.StatusOK, e.GETVERIFYCODE_SUCCESS)
}

func Signup(c *gin.Context) {
	var user d.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		common.CJSON(c, http.StatusOK, e.INVALID_PARAMS)
		return
	}

	//mysql 错误处理
	//返回的错误代code 则为mysql错误
	//参考https://stackoverflow.com/questions/47009068/how-to-get-the-mysql-error-type-in-golang
	uid, _, err := d.AddUser(user.Uname, user.Upwd, user.Pnumber)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == mysqlerr.ER_DUP_ENTRY {
				common.CJSON(c, http.StatusOK, e.ER_DUP_ENTRY_NAME)
				return
			}
		}
		common.CJSON(c, http.StatusOK, e.ERROR_SIGNUP)
		return
	}

	jwtstring, err := middleware.GenerateToken(uid)
	if err != nil {
		log.Println("Faield to GenerateJwtString")
	}
	//
	//name	    string	    cookie名字
	//value	    string	    cookie值
	//maxAge	int	        有效时间，单位是秒，MaxAge=0  过期时间项为session 关闭浏览器后 cookie失效，MaxAge<0 相当于删除cookie, 通常可以设置-1代表删除，MaxAge>0 多少秒后cookie失效
	//path	    string	    cookie路径
	//domain	string	    cookie作用域
	//secure	bool	    Secure=true，那么这个cookie只能用https协议发送给服务器
	//httpOnly	bool	    设置HttpOnly=true的cookie不能被js获取到
	c.SetCookie("jwtstring", jwtstring, 365*24*3600, "/", "localhost", false, false)
	common.CJSON(c, http.StatusOK, e.SIGNUP_SUCCESS)
}
