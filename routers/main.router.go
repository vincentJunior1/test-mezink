package routers

import (
	"skeleton-svc/controllers"
	"skeleton-svc/helpers"
	"strings"
	"time"

	"skeleton-svc/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterInterface Abstract Class
type (
	RouterInterface interface {
		StartServer() error
		routerControllers()
	}

	// Router Actual Class implementation
	Router struct {
		address    string
		port       string
		Path       map[string]string
		Gin        *gin.Engine
		Controller controllers.Controller
		Logs       *logs.Logger
	}
)

// InitializeRouter return sturct that will implement the abs. class
func InitializeRouter(ctrl controllers.Controller, l *logs.Logger) RouterInterface {
	gin.SetMode(helpers.GetEnv("ROUTER_SETMODE"))
	return &Router{
		address: helpers.GetEnv("ROUTER_SERVER_ADDRESS"),
		port:    helpers.GetEnv("ROUTER_PORT"),
		Path: map[string]string{
			"PATH_VERSION": helpers.GetEnv("PATH_VERSION"),
			"PATH_MAIN":    helpers.GetEnv("PATH_MAIN"),
		},
		Gin:        gin.New(),
		Controller: ctrl,
		Logs:       l,
	}
}

// StartServer Start the server by initialize end-point & create the server
func (r *Router) StartServer() error {
	r.Logs.Info("Starting Server on ", r.address+r.port)
	r.routerControllers()
	// err := r.Gin.Run(r.port)

	err := helpers.GinServerUp(r.address+r.port, r.Gin)
	if err != nil {
		r.Logs.Error("[GinServerUp]Error: ", err)
		return err
	}

	return nil
}

// routerControllers ...
func (r *Router) routerControllers() {

	docs.SwaggerInfo.Title = helpers.GetEnv("SWAG_TITLE")
	docs.SwaggerInfo.Description = helpers.GetEnv("SWAG_DESC")
	docs.SwaggerInfo.Version = helpers.GetEnv("SWAG_VERSION")
	docs.SwaggerInfo.Host = helpers.GetEnv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = r.Path["PATH_MAIN"]
	docs.SwaggerInfo.Schemes = strings.Split(helpers.GetEnv("SWAG_SCHEMES"), ",")

	r.Gin.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "PATCH", "PUT", "POST", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization", "access-control-allow-origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		//AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge: 12 * time.Hour,
	}))

	url := ginSwagger.URL("http://" + docs.SwaggerInfo.Host + "/swagger/doc.json") // The url pointing to API definition
	// }

	r.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	main := r.Gin.Group(r.Path["PATH_MAIN"])

	// BEGIN __INCLUDE_TEMPLATE__
	{
		// version 1
		v1 := main.Group("/v1")
		v1.POST("/records/", r.Controller.V1().GetRecords)

	}
	// END __INCLUDE_TEMPLATE__

	// Invalid URL
	r.Gin.NoRoute(func(c *gin.Context) {
		logs.WithFields(logs.Fields{"URL": r.address + r.port, "Method": c.Request.Method, "Path": c.Request.URL.Path}).Error("[NoRoute]Invalid")
		c.JSON(404, gin.H{"code": "404", "message": "Page not found"})
	})

	r.Logs.Info("initializing Router done")
}
