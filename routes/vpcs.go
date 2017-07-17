package routes

/*
 * Run /vpcs to get a JSON array of VPC ids and descriptions
 */

import (
	"encoding/json"
	"net/http"
	"awsplus/aws/vpcs"
)

type Vpc struct {
	Id, Description string
}

func ListVpcs(writer http.ResponseWriter, request *http.Request) {

	var vs = make([]Vpc, len(vpcs.List()))

	for i, v := range vpcs.List() {

		var tagValue = ""
		for _, tag := range v.Tags {
			if *tag.Key == "Name" {
				tagValue = *tag.Value
			}
		}

		vs[i] = Vpc{*v.VpcId, tagValue}
	}

	json.NewEncoder(writer).Encode(vs)

}
