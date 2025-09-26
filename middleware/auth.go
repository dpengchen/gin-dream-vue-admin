package middleware

import (
	"dream-vue-admin/global"
	"dream-vue-admin/pkg/jwt"
	"dream-vue-admin/util/constants"
	"dream-vue-admin/util/http/response"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.GetHeader("Token")

		//角色默认为空
		role := ""

		//判断token是否存在
		if token == "" {
			_, err := enforce(c, role)
			if err != nil {
				response.Fail(c, http.StatusUnauthorized, err.Error())
			}
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			global.Log.Warnf("token解析失败: %v", err.Error())
			response.Fail(c, http.StatusUnauthorized, "身份令牌过期！")
			c.Abort()
			return
		}

		//获取subject
		sub, err := claims.GetSubject()
		if err != nil {
			global.Log.Warnf("获取subject失败: %v", err.Error())
			response.Fail(c, http.StatusUnauthorized, "身份令牌过期！")
			c.Abort()
			return
		}

		loginCache := global.RedisClient.Get(c, fmt.Sprintf(constants.RedisLoginUserPrefixKey, sub, token))
		if err = loginCache.Err(); err != nil {
			global.Log.Infof("redis获取用户信息失败: %v", err.Error())
			response.FailWithInternalError(c, "获取用户信息失败"+err.Error())
			c.Abort()
			return
		}

		//设置当前用户id到上下文
		userInfoStr := loginCache.Val()
		userId, _ := strconv.Atoi(sub)
		c.Set(constants.GinContextLoginUserIdKey, userId)
		//TODO 后期这里更改为结构体而不是字符串
		c.Set(constants.GinContextLoginUserInfoKey, userInfoStr)

		//TODO 解析用户信息拿到用户角色数据，然后循环判断是否有角色信息
	}
}

func enforce(c *gin.Context, role string) (flag bool, err error) {
	var enforceBool bool
	enforceBool, err = global.Enforcer.Enforce(role, c.Request.RequestURI, c.Request.Method)

	//验证失败，出现错误
	if err != nil {
		return
	}

	//判断是否通过验证
	if !enforceBool {
		err = errors.New("无权限访问，请联系管理员！")
		return
	}
	//通过
	c.Next()
	flag = true
	return
}
