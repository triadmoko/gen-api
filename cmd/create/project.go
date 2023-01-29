/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package create

import (
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/triadmoko/gen-api/helper"
)

// projectCmd represents the project command
var (
	projectName string
	projectCmd  = &cobra.Command{
		Use:   "project",
		Short: "Create First Your Project",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			createProject(projectName)
		},
	}
)

func init() {
	projectCmd.Flags().StringVarP(&projectName, "name", "n", "", "your project name")
	Create.AddCommand(projectCmd)
}

func createProject(projectName string) {
	version := getVersion()
	createDirProject(projectName)
	createModule(projectName, version)
	createMain()
}

func createModule(projectName string, version string) {
	d1 := []byte("module " + projectName + "\n\ngo " + version)
	err := os.WriteFile(projectName+"/go.mod", d1, 0644)
	helper.Check(err)
	log.Println("Debug := Create Module Project")
	time.Sleep(1 * time.Second)
}

func createMain() {
	d1 := []byte(`
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
} `)

	err := os.WriteFile(projectName+"/main.go", d1, 0644)
	helper.Check(err)
	log.Println("Debug := Create Main Project")
	time.Sleep(1 * time.Second)
}

func getVersion() string {
	version := helper.GoVersion()
	log.Println("Debug := Get Version Golang")
	time.Sleep(1 * time.Second)
	return version
}

func createDirProject(projectName string) {
	err := os.MkdirAll("./"+projectName, os.ModePerm)
	helper.Check(err)
	log.Println("Debug := Create Directory Project")
	time.Sleep(1 * time.Second)
}
