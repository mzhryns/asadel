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
	baseUrl             string
}

func NewHandler(
	url_shortener_uc url_shortener.UseCase,
	meta_data_reader_uc meta_data_reader.UseCase,
	baseUrl string,
) RestHandler {
	return &restHandler{
		url_shortener_uc:    url_shortener_uc,
		meta_data_reader_uc: meta_data_reader_uc,
		baseUrl:             baseUrl,
	}
}

func (h *restHandler) Shorten(c *gin.Context) {
	longUrl := c.Query(QUERY_URL)
	deeplink := c.Query(QUERY_DEEPLINK)
	storeAndroid := c.Query(QUERY_ANDROID_STORE)
	storeIOS := c.Query(QUERY_IOS_STORE)

	shortUrl, err := h.url_shortener_uc.Shorten(longUrl, deeplink, storeAndroid, storeIOS)
	if err == nil {
		c.JSON(http.StatusOK, SuccessResponse{Data: shortUrl})
	} else {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}

}

func (h *restHandler) Redirect(c *gin.Context) {
	shortenUrl := c.Param("shortenUrl")
	result, err := h.url_shortener_uc.Redirect(shortenUrl)

	longUrl := helpers.FormatUrl(h.baseUrl, result)

	if err == nil {
		c.Redirect(http.StatusFound, longUrl)
	} else {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: err.Error()})
	}
}

func (h *restHandler) Resolve(c *gin.Context) {
	url := c.Query(QUERY_URL)
	deeplink := c.Query(QUERY_DEEPLINK)
	storeAndroid := c.Query(QUERY_ANDROID_STORE)
	storeIOS := c.Query(QUERY_IOS_STORE)

	metaData, _ := h.meta_data_reader_uc.GetOpenGraph(url)

	c.HTML(http.StatusOK, "url/index.tmpl", gin.H{
		"title":       metaData.Title,
		"description": metaData.Description,
		"image":       metaData.Image,
		"type":        metaData.Type,
		"url":         metaData.Url,
		"deeplink":    deeplink,
		"android":     storeAndroid,
		"ios":         storeIOS,
	})
}
