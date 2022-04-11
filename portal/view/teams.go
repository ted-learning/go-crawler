package view

import (
	"html/template"
	"net/http"
	"portal/model"
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
