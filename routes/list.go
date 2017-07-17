package routes

import (
	"html/template"
	"net/http"
	"awsplus/aws/securitygroups"
	"awsplus/aws/vpcs"
	"strings"
)

func List(writer http.ResponseWriter, request *http.Request) {

	urlPart := strings.Split(request.URL.Path, "/")
	vpcId := urlPart[len(urlPart)-1]

	// lookup some names
	funcMap := template.FuncMap{
		"sgname":  securitygroups.GetName,
		"vpcname": vpcs.GetName,
	}

	listTemplate, err := template.New("listGroups").Funcs(funcMap).ParseFiles("templates/list.html")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

	err = listTemplate.ExecuteTemplate(writer, "list.html", securitygroups.List(vpcId))
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

}
