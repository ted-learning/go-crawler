package view

import (
	"go-crawler/frontend/model"
	"html/template"
	"net/http"
)

type TeamResultView struct {
	template *template.Template
}

func CreateTeamResultView(filename string) TeamResultView {
	return TeamResultView{
		template: template.Must(template.ParseFiles(filename)),
	}
}

func (v *TeamResultView) Render(writer http.ResponseWriter, result model.SearchResult) error {
	return v.template.Execute(writer, result)
}
