package main

import (
	"github.com/syaddadSmiley/SeminarPage/api"
	con "github.com/syaddadSmiley/SeminarPage/database"
	repository "github.com/syaddadSmiley/SeminarPage/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := con.Connect()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)
	userAdmin := repository.NewTaskRepo(db)
	route := api.NewAPI(*userRepo, *userAdmin)
	route.Start()
}
