package main

import (
	"log"
	"net/http"
	"udemy/webserver_postgresSQL/router"
	"udemy/webserver_postgresSQL/services"
	"udemy/webserver_postgresSQL/utils"
)

func main(){
	log.Println("In Main App")

	var dbconn = utils.GetConnection()
	services.SetDB(dbconn)
	var appRouter = router.CreateRouter()

	log.Println("Starting Server")
	log.Fatal(http.ListenAndServe(":8080",appRouter))
}