package controller

import (
	"context"
	"github.com/olivere/elastic/v7"
	"net/http"
	"portal/common"
	"portal/model"
	"portal/view"
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

type ViewTeam struct {
	common.Team
	IsEast bool
	IsWest bool
}

func (s *SearchTeamResultHandler) search(q string) (model.SearchResult, error) {
	var result model.SearchResult
	search := s.client.Search("team").Size(500)
	if q != "" {
		search = search.Query(elastic.NewQueryStringQuery(rewriteQueryString(q)))
	}
	result.Query = q

	resp, err := search.
		SortBy(
			elastic.NewFieldSort("AreaId"),
			elastic.NewFieldSort("Rank"),
		).
		Do(context.Background())

	if err != nil {
		result.Hits = 0
		result.Items = make([]interface{}, 0)
	} else {
		teams := resp.Each(reflect.TypeOf(common.Team{}))
		result.Hits = resp.TotalHits()
		result.Items = make([]interface{}, 0)

		for _, team := range teams {
			t := team.(common.Team)
			result.Items = append(result.Items, ViewTeam{
				t,
				t.Area == "东部",
				t.Area == "西部",
			})
		}
	}
	return result, nil
}

// Rewrites query string. Replaces field names
// like "Age" to "Payload.Age"
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
