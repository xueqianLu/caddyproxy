package openapi

import (
	"caddyproxy/caddy"
	"caddyproxy/types"
	"caddyproxy/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

type apiHandler struct {
	conf    *Config
	backend *caddy.CaddyAPI
}

func (api apiHandler) CreateWebsite(c *gin.Context) {
	var req types.CreateWebsite
	err := c.ShouldBindJSON(&req) // 解析req参数
	if err != nil {
		log.WithError(err).Error("CreateWebsite ctx.ShouldBindJSON error")
		api.response(c, http.StatusBadRequest, err, nil)
	}

	uid := uuid.NewString()
	target := filepath.Join(api.conf.TempDir, uid)
	if err := utils.Download(req.Resource, target); err != nil {
		log.WithError(err).Error("Download resource failed")
		api.response(c, http.StatusInternalServerError, err, nil)
	}

	if err := api.backend.CreateWebsite(req.Domain, target); err != nil {
		log.WithError(err).Error("CreateWebsite backend.CreateWebsite error")
		api.response(c, http.StatusInternalServerError, err, nil)
	} else {
		api.response(c, http.StatusOK, nil, nil)
	}
}

func (api apiHandler) ForwardWebsite(c *gin.Context) {
	var req types.ForwardWebsite
	err := c.ShouldBindJSON(&req) // 解析req参数
	if err != nil {
		log.WithError(err).Error("ForwardWebsite ctx.ShouldBindJSON error")
		api.response(c, http.StatusBadRequest, err, nil)
	}

	if err := api.backend.ForwardWebsite(req); err != nil {
		log.WithError(err).Error("ForwardWebsite backend.ForwardWebsite error")
		api.response(c, http.StatusInternalServerError, err, nil)
	} else {
		api.response(c, http.StatusOK, nil, nil)
	}
}

func (api apiHandler) response(c *gin.Context, code int, err error, data interface{}) {
	result := make(map[string]interface{})
	result["code"] = code
	if err != nil {
		result["message"] = err.Error()
	}
	if data != nil {
		result["data"] = data
	}

	c.JSON(code, result)
}
