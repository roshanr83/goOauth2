package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	oauth2 "github.com/gobeam/golang-oauth"
	"github.com/gobeam/golang-oauth/example/common"
	"github.com/gobeam/golang-oauth/example/core/models"
	"github.com/gobeam/golang-oauth/example/routers"
	"github.com/jinzhu/gorm"
	"log"
)

func main() {
	dbUrl := common.GetConfig("mysql", "url").String()
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	store := oauth2.NewDefaultStore(
		oauth2.NewConfig(dbUrl),
	)
	defer store.Close()
	models.InitializeDb(db.Debug())

	// register custom validator
	newValidator := common.NewValidatorRegister(db)
	newValidator.RegisterValidator()

	// router setup
	router := routers.SetupRouter(store)

	serverError := router.Run(fmt.Sprintf(":%s", common.GetConfig("system", "httpport").String()))
	if serverError != nil {
		log.Fatalf("Server failed to start %v ", serverError)
	}
}
