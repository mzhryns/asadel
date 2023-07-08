package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhryn/asadel/app/usecase/url_shortener"
)

type restHandler struct {
	url_shortener_uc url_shortener.UseCase
}

func NewHandler(url_shortener_uc url_shortener.UseCase) RestHandler {
	return &restHandler{url_shortener_uc: url_shortener_uc}
}

func (h *restHandler) Shorten(c *gin.Context) {
	longUrl := c.Query("url")
	shortUrl, err := h.url_shortener_uc.Shorten(longUrl)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: shortUrl})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

}

func (h *restHandler) Resolve(c *gin.Context) {
	shortenUrl := c.Param("shortenUrl")
	longUrl, err := h.url_shortener_uc.Resolve(shortenUrl)

	if err == nil {
		c.Redirect(http.StatusFound, longUrl)
	} else {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	}
}
