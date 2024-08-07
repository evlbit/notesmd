package main

import (
	"flag"
	"log"

	"github.com/evlbit/notesmd/internal/db"
	"github.com/evlbit/notesmd/internal/env"
	"github.com/evlbit/notesmd/internal/server"
	"github.com/go-sql-driver/mysql"
)

func main() {
	appEnv := flag.String("env", "", "Used to specify app environment (prod, dev, ...)")

	flag.Parse()

	if *appEnv == "" {
		log.Fatal("Environment was not specified")
	}

	err := env.InitEnv(*appEnv)
	if err != nil {
		log.Fatalf("Could not load environment\n%s", err)
	}

	log.Println("Successfully load environment")

	db, err := db.NewDB(mysql.Config{
		User:                 env.Vars.DBUser,
		Passwd:               env.Vars.DBPassword,
		Addr:                 env.Vars.DBAddress,
		DBName:               env.Vars.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatalf("Could not connect to database\n%s", err)
	}

	log.Println("Successfully connected to database")

	server.StartServer(db)
}
