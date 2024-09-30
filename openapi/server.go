package openapi

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type OpenAPI struct {
	conf *Config
}

func NewOpenAPI(conf *Config) *OpenAPI {
	return &OpenAPI{conf: conf}
}

func (s *OpenAPI) Run() error {
	return s.startHttp(fmt.Sprintf("%s:%d", s.conf.Host, s.conf.Port))
}

func (s *OpenAPI) startHttp(address string) error {
	router := gin.Default()
	router.Use(cors())
	router.Use(ginLogrus())
	// 创建v1组
	v1 := router.Group("/v1")
	{
		v1.POST("/create-website", apiHandler{}.CreateWebsite)
		v1.POST("/forward-website", apiHandler{}.ForwardWebsite)
	}

	return router.Run(address)
}

// gin use logrus
func ginLogrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.WithFields(log.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"query":  c.Request.URL.RawQuery,
			"ip":     c.ClientIP(),
		}).Info("request")
		c.Next()
	}
}

// enable cors
func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
