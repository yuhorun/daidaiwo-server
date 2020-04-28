package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AddAllowHeaders("Cookie")
	config.AllowOrigins = []string{"http://localhost:3000"}

	return cors.New(config)
}

/*	客户端配置
const handlerLogInBtn = ()=>{
	axios({
		url:"/login",
		baseURL:"http://localhost:8000",
		method:"POST",
		data:JSON.stringify({pnumber:state.logInPhoneNumber,upwd:state.logInPassword}),
	}).then(res => {
		const data = res.data;
		console.log(data)
	})
}
*/

/*	服务端配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:3000"}

	return cors.New(config)
}
*/
