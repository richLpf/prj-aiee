package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"prj-aiee/config"
	_ "prj-aiee/docs"
	"prj-aiee/service/apis/common"
	"strings"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ApiGroups = []*Group{}

type Router struct {
	Router   string
	Function func(c *gin.Context)
}

type Group struct {
	Router  string
	Routers []*Router
}

func NewGroup(groupRouter string) *Group {
	return &Group{
		Router:  groupRouter,
		Routers: []*Router{},
	}
}

func (g *Group) NewRouter(router string, function func(c *gin.Context)) {
	g.Routers = append(g.Routers, &Router{
		Router:   router,
		Function: function,
	})
}

func (g *Group) Register() {
	ApiGroups = append(ApiGroups, g)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 排除
		if strings.HasPrefix(c.FullPath(), "/swagger") || strings.HasPrefix(c.FullPath(), "/v1/user/login") || strings.HasPrefix(c.FullPath(), "/v1/user/create") {
			c.Next()
			return
		}
		// 从请求头中获取 Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供 Token"})
			c.Abort()
			return
		}
		// 在这里进行 Token 的解析和验证
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证签名的密钥
			return []byte(common.SecretKey), nil
		})
		if err != nil {
			fmt.Println(err)
		}
		// 将解析后的 Token 信息存储到上下文中，方便后续处理
		claims, _ := token.Claims.(jwt.MapClaims)
		exp := claims["exp"].(float64)
		timestamp := float64(exp)
		now := float64(time.Now().Unix())
		if timestamp < now {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token无效",
			})
			c.Abort()
			return
		}
		// 调用下一个处理程序
		c.Next()
	}
}

// 定义CORS中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func Init() {
	// 设置为发布模式（初始化路由之前设置）
	gin.SetMode(gin.ReleaseMode)
	// gin 默认中间件
	r := gin.Default()
	r.Use(Cors())
	r.Use(AuthMiddleware())
	// 访问一个错误路由时，返回404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  404,
			"message": "404, page not exists!",
		})
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 注册路由
	for _, group := range ApiGroups {
		v1 := r.Group("v1")
		gr := v1.Group(group.Router)
		for _, router := range group.Routers {
			gr.POST(router.Router, router.Function)
		}
	}

	// 启动API服务
	if err := r.Run(fmt.Sprintf(":%v", config.Cfg.Service.Port)); err != nil {
		panic(err)
	}
}
