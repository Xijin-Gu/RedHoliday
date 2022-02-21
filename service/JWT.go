package service

import (
	"RedHoliday/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

//该go文件存放和jwt相关函数及中间件：jwt生成，jwt解析，jwt认证中间件


//jwt生成
func Generate_Token(username string)(string,error){
	//创建声明
	c := model.Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(model.JWT_Effective_Time).Unix(),
			Issuer: "redholiday-project",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,c)
	return token.SignedString(model.Secret)

}


//JWT解析
func Parse_JWT(tokenstring string)(*model.Claims,error){
	//解析token
	token,err := jwt.ParseWithClaims(tokenstring,&model.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return model.Secret,nil
	})
	//处理错误
	if err != nil {
		return nil, err
	}
	//验证是否token有效
	if claims,ok := token.Claims.(*model.Claims); ok && token.Valid {
		return claims,nil
	}
	return nil, errors.New("token无效")
}


//JWT认证
func Verify_JWT()func(c *gin.Context){
	return func(c *gin.Context) {

		//获取含有token信息的头部Authorazition部分
		authorization := c.Request.Header.Get("Authorization")
		fmt.Println(authorization)
		if authorization == "" {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2003,
				"message" : "Authorazition为空",
			})
			c.Abort()
			return
		}

		//提取token信息段
		JWT_information := strings.SplitN(authorization," ",2)
		//验证auth信息段是否合法
		if !(len(JWT_information) == 2 && JWT_information[0] == "Bearer") {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2004,
				"message" : "Authorazition格式错误",
			})
			c.Abort()
			return
		}
		//验证token是否有效
		claim,err := Parse_JWT(JWT_information[1])
		if err != nil {
			c.JSON(http.StatusOK,gin.H{
				"code" : 2005,
				"message" : "token无效",
			})
			c.Abort()
			return
		}
		//将claim信息保存到上下文
		c.Set("username",claim.Username)
		c.Next()
	}
}