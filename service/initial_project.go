package service

import (
	"log"
	"os"
)

func InitProject() {
	if err := os.MkdirAll("configs", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("docs", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("domain", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("helpers", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("injectors", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("middleware", os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll("routers", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}
