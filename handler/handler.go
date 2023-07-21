package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhryn/asadel/app/usecase/meta_data_reader"
	"github.com/zhryn/asadel/app/usecase/url_shortener"
	"github.com/zhryn/asadel/pkg/helpers"
)

type restHandler struct {
	url_shortener_uc    url_shortener.UseCase
	meta_data_reader_uc meta_data_reader.UseCase
}

func NewHandler(
	url_shortener_uc url_shortener.UseCase,
	meta_data_reader_uc meta_data_reader.UseCase,
) RestHandler {
	return &restHandler{
		url_shortener_uc:    url_shortener_uc,
		meta_data_reader_uc: meta_data_reader_uc,
	}
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

func (h *restHandler) Redirect(c *gin.Context) {
	shortenUrl := c.Param("shortenUrl")
	longUrl, err := h.url_shortener_uc.Redirect(shortenUrl)
	longUrl = helpers.FormatUrl("http://localhost:8080", longUrl)

	if err == nil {
		c.Redirect(http.StatusFound, longUrl)
	} else {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) Resolve(c *gin.Context) {
	url := c.Query("url")
	meta, _ := h.meta_data_reader_uc.GetOpenGraph(url)
	c.HTML(http.StatusOK, "url/index.tmpl", gin.H{
		"title":       meta.Title,
		"description": meta.Description,
		"image":       meta.Image,
		"type":        "Website",
		"url":         meta.Url,
	})
}
