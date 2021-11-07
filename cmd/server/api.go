package main

import (
	"context"
	"log"
	"os"
	"time"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/HongXiangZuniga/CrudGoExample/pkg/http/rest"
	"github.com/HongXiangZuniga/CrudGoExample/pkg/persistence/mongodb"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mdbName string
	mdbURI  string
	mdb     mongo.Database
)

func main() {
	log.Println("Init Api")
	mdb := initMongo(mdbURI, mdbName)
	usersRepo := mongodb.NewUserRepo(mdb)
	usersServices := users.NewUserServices(usersRepo)
	usersHandler := rest.NewUsersHandler(usersServices)
	r := rest.NewHandler(usersHandler)
	if os.Getenv("PORT") == "" {
		log.Println("Api not run for: PORT Not found")
	} else {
		err := r.Run(":" + os.Getenv("PORT"))
		if err != nil {
			log.Println("Api not run for: " + err.Error())
		} else {
			log.Println("Api ok")
		}
	}

}

func init() {
	godotenv.Load()
	mdbURI = os.Getenv("MONGO_URI")
	mdbName = os.Getenv("MONGO_DBNAME")
}
func initMongo(uri string, dbName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Println("Not mongo connection, error: " + err.Error())
		os.Exit(0)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	err = client.Connect(ctx)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Println("Not mongo connection, error: " + err.Error())
		os.Exit(0)
	}
	return client.Database(dbName)
}
