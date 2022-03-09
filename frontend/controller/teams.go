package controller

import (
	"context"
	"github.com/olivere/elastic/v7"
	"go-crawler/common"
	"go-crawler/frontend/model"
	"go-crawler/frontend/view"
	"go-crawler/parser"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

type SearchTeamResultHandler struct {
	view   view.TeamResultView
	client *elastic.Client
}

func CreateSearchTeamResultHandler(template string, client *elastic.Client) *SearchTeamResultHandler {
	return &SearchTeamResultHandler{
		view:   view.CreateTeamResultView(template),
		client: client,
	}
}

func (s *SearchTeamResultHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	q := strings.TrimSpace(request.FormValue("q"))
	result, err := s.search(q)
	common.HandleServerError(writer, err)

	err = s.view.Render(writer, result)
	common.HandleServerError(writer, err)
}

func (s *SearchTeamResultHandler) search(q string) (model.SearchResult, error) {
	var result model.SearchResult
	search := s.client.Search("team").Size(500)
	if q != "" {
		search = search.Query(elastic.NewQueryStringQuery(rewriteQueryString(q)))
	}
	resp, err := search.
		SortBy(
			elastic.NewFieldSort("AreaId"),
			elastic.NewFieldSort("Rank"),
		).
		Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Query = q
	result.Hits = resp.TotalHits()
	result.Items = resp.Each(reflect.TypeOf(parser.Team{}))
	return result, nil
}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
