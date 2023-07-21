package meta_data_repository_v1

import (
	"io"
	"strings"

	"github.com/zhryn/asadel/entity"
	"github.com/zhryn/asadel/pkg/helpers"
	"golang.org/x/net/html"
)

type MetaData struct {
	Title       string
	Description string
	Image       string
	Url         string
	Type        string
}

func (md *MetaData) ToMetaDataEntity() *entity.MetaData {
	return &entity.MetaData{
		Title:       md.Title,
		Description: md.Description,
		Image:       md.Image,
		Url:         md.Url,
		Type:        md.Type,
	}
}

func (md *MetaData) FillMetaData(ogType, content string) {
	switch ogType {
	case "og:title":
		md.Title = content
	case "og:description":
		md.Description = content
	case "og:image":
		md.Image = content
	case "og:url":
		md.Url = content
	case "og:type":
		md.Type = content
	}
}

func (MetaData) FromMetaDataEntity(e *entity.MetaData) *MetaData {
	return &MetaData{
		Title:       e.Title,
		Description: e.Description,
		Image:       e.Image,
		Url:         e.Url,
		Type:        e.Type,
	}
}

func (MetaData) FromResponseBody(body io.Reader) *MetaData {
	z := html.NewTokenizer(body)
	metaData := &MetaData{}

	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken: // End of the document
			return metaData

		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()

			if t.Data == "meta" {
				ogType := helpers.GetAttr(t.Attr, "property")
				if strings.HasPrefix(ogType, "og:") {
					content := helpers.GetAttr(t.Attr, "content")
					metaData.FillMetaData(ogType, content)
				}
			}
		}
	}
}
