package main

import (
	"log"
	"net/http"

	"github.com/Nickadimas79/movies/handler"
	"github.com/Nickadimas79/movies/repository"
	"github.com/Nickadimas79/movies/service"
	"github.com/julienschmidt/httprouter"
)

func main() {
	movieInMemoryRepository := repository.NewInMemoryMovieRepository()

	//moviePostgreSQLRepository := repository.NewPostgreSQLMovieRepository()
	//movieService := service.NewDefaultMovieService(moviePostgreSQLRepository)
	movieService := service.NewDefaultMovieService(movieInMemoryRepository)
	movieHandler := handler.NewMovieHandler(movieService)

	router := httprouter.New()

	router.GET("/movies", movieHandler.GetMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)

	router.POST("/movies", movieHandler.CreateMovie)

	router.PATCH("/movies/:id", movieHandler.UpdateMovie)

	router.DELETE("/movies", movieHandler.DeleteAllMovies)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)

	log.Println("http server runs on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
