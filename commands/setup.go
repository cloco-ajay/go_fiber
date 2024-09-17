package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	model := flag.String("model", "", "model name must start with small letter i.e 'go run commands/setup.go -model=modelname'")
	flag.Parse()
	modelName := string(*model)

	// prepparing mode
	PrepareModel(modelName)

	// preparing handler
	PrepareHandler(modelName)

	// prepare usecase
	PrepareUsecase(modelName)

	// prepare repository
	PrepareRepository(modelName)

}

func WriteNewLines(file *os.File) {
	file.WriteString("\n")
}

func PrepareModel(modelName string) {

	modelPath := filepath.Join("models", fmt.Sprint(modelName, ".go"))

	file, modelErr := os.Create(modelPath)
	if modelErr != nil {
		fmt.Println(modelErr.Error())
		return
	}

	file.WriteString("package models")
	WriteNewLines(file)
	WriteNewLines(file)
	file.WriteString(`import "time"`)
	WriteNewLines(file)
	WriteNewLines(file)
	modelAsTitle := strings.Title(modelName)
	file.WriteString(fmt.Sprint("type ", modelAsTitle, " struct { \n"))
	file.WriteString(fmt.Sprint("ID uint ", "`json:\"id\" gorm:\"primaryKey; autoIncrement\"`"))
	WriteNewLines(file)
	file.WriteString(fmt.Sprint("CreatedAt time.Time ", "`json:\"created_at\"`"))
	WriteNewLines(file)
	file.WriteString(fmt.Sprint("UpdatedAt *time.Time ", "`json:\"updated_at\"`"))
	WriteNewLines(file)
	file.WriteString("}")
}

func PrepareHandler(modelName string) {
	titleModelName := strings.Title(modelName)
	handlerPath := filepath.Join("handlers", fmt.Sprint(titleModelName, "Handler.go"))
	handlerFile, handlerErr := os.Create(handlerPath)
	if handlerErr != nil {
		fmt.Println(handlerErr.Error())
		return
	}

	handlerFile.WriteString("package handlers \n")
	WriteNewLines(handlerFile)
	handlerFile.WriteString("import (\n")
	// handlerFile.WriteString(`"sales-api/models"`)
	// WriteNewLines(handlerFile)
	handlerFile.WriteString(`"sales-api/repository"`)
	WriteNewLines(handlerFile)
	handlerFile.WriteString(`"sales-api/usecase"`)
	WriteNewLines(handlerFile)
	WriteNewLines(handlerFile)
	// handlerFile.WriteString(`"github.com/gofiber/fiber/v2"`)
	// WriteNewLines(handlerFile)
	handlerFile.WriteString(`"gorm.io/gorm"`)
	WriteNewLines(handlerFile)
	handlerFile.WriteString(")\n\n\n")

	handlerFile.WriteString(fmt.Sprint("type ", titleModelName, "Handler struct {\n"))
	handlerFile.WriteString(fmt.Sprint(modelName, "Usecase ", "usecase.", titleModelName, "Usecase\n"))
	handlerFile.WriteString("}")
	WriteNewLines(handlerFile)
	WriteNewLines(handlerFile)
	WriteNewLines(handlerFile)

	handlerFile.WriteString(fmt.Sprint("func New", titleModelName, "Handler(db *gorm.DB) *", titleModelName, "Handler {\n"))
	handlerFile.WriteString(fmt.Sprint(modelName, "Repo := repository.New", titleModelName, "Repository(db)\n"))
	handlerFile.WriteString(fmt.Sprint(modelName, "Usecase := usecase.New", titleModelName, "Usecase(", modelName, "Repo)\n"))
	handlerFile.WriteString(fmt.Sprint("return &", titleModelName, "Handler{", modelName, "Usecase: ", modelName, "Usecase}\n"))
	handlerFile.WriteString("}\n")

}

func PrepareUsecase(modelName string) {
	titleModelName := strings.Title(modelName)
	usecasePath := filepath.Join("usecase", fmt.Sprint(titleModelName, "Usecase.go"))
	usecaseFile, usecaseErr := os.Create(usecasePath)
	if usecaseErr != nil {
		fmt.Println(usecaseErr.Error())
		return
	}
	usecaseFile.WriteString("package usecase")
	WriteNewLines(usecaseFile)

	usecaseFile.WriteString("import (\n")
	// usecaseFile.WriteString(`"sales-api/models"`)
	// WriteNewLines(usecaseFile)
	usecaseFile.WriteString(`"sales-api/repository"`)
	WriteNewLines(usecaseFile)
	usecaseFile.WriteString(")\n")
	WriteNewLines(usecaseFile)

	usecaseFile.WriteString(fmt.Sprint("type ", titleModelName, "Usecase interface {\n"))
	usecaseFile.WriteString("}\n")
	WriteNewLines(usecaseFile)

	usecaseFile.WriteString(fmt.Sprint("type ", modelName, "Usecase struct {\n"))
	usecaseFile.WriteString(fmt.Sprint("repo repository.", titleModelName, "Repository"))
	usecaseFile.WriteString("}\n")
	WriteNewLines(usecaseFile)

	usecaseFile.WriteString(fmt.Sprint("func New", titleModelName, "Usecase (repo repository.", titleModelName, "Repository) ", titleModelName, "Usecase {\n"))
	usecaseFile.WriteString(fmt.Sprint("return &", modelName, "Usecase{repo: repo}\n"))
	usecaseFile.WriteString("}\n")
}

func PrepareRepository(modelName string) {
	titleModelName := strings.Title(modelName)
	repoPath := filepath.Join("repository", fmt.Sprint(titleModelName, "Repository.go"))
	repoFile, repoErr := os.Create(repoPath)
	if repoErr != nil {
		fmt.Println(repoErr.Error())
		return
	}

	repoFile.WriteString("package repository\n\n")

	repoFile.WriteString("import (\n")
	// repoFile.WriteString(`"sales-api/models"`)
	// WriteNewLines(repoFile)
	// repoFile.WriteString(`"sales-api/utils"`)
	// WriteNewLines(repoFile)
	repoFile.WriteString(`"gorm.io/gorm"`)
	WriteNewLines(repoFile)
	repoFile.WriteString(")\n\n\n")

	repoFile.WriteString(fmt.Sprint("type ", titleModelName, "Repository interface {"))
	WriteNewLines(repoFile)
	repoFile.WriteString("}\n\n\n")

	repoFile.WriteString(fmt.Sprint("type ", modelName, "Repository struct {"))
	WriteNewLines(repoFile)
	repoFile.WriteString("db *gorm.DB\n}\n\n\n")

	repoFile.WriteString(fmt.Sprint("func New", titleModelName, "Repository(db *gorm.DB) ", titleModelName, "Repository {"))
	WriteNewLines(repoFile)
	repoFile.WriteString(fmt.Sprint("return &", modelName, "Repository{db: db}"))
	WriteNewLines(repoFile)
	repoFile.WriteString("}\n\n\n")

}
