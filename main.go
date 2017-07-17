package main

import (
	"fmt"
	"net/http"
	"awsplus/config"
	"awsplus/logging"
	"awsplus/routes"
	"strconv"
)

func main() {

	cfg := config.Config
	port := cfg.Port

	mux := http.NewServeMux()

	files := http.FileServer(http.Dir(cfg.Static))
	mux.Handle("/frontend/", http.StripPrefix("/frontend/", files))

	// html template pages
	mux.HandleFunc("/", routes.Index)
	mux.HandleFunc("/list/", routes.List)

	// web services returning JSON
	mux.HandleFunc("/vpcs/", routes.ListVpcs)
	mux.HandleFunc("/securitygroups/", routes.ListSecurityGroups)


	fmt.Println("Starting awsplus on", port)
	logging.Info("Starting awsplus on", port)

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: mux,
	}
	server.ListenAndServe()

}
