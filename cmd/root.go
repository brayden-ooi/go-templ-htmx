/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/brayden-ooi/go-templ-htmx/internal/router"
)

func Execute() {
	router.Exec()

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("something went wrong when serving: ", err)
	}
}
