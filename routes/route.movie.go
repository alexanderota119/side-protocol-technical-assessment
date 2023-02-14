package route

import (
	movieController "github.com/david-side-protocol-technical/controllers"
	movieHandler "github.com/david-side-protocol-technical/handlers"
	"github.com/gin-gonic/gin"
)

func InitMovieRoutes(route *gin.Engine) {

	/**
	@description All Handler Movie
	*/
	movieController := movieController.NewMovieController()
	movieHandler := movieHandler.NewMovieHandler(movieController)

	/**
	@description All Movie Route
	*/
	groupRoute := route.Group("/api/v1/movies/")
	groupRoute.GET("/search", movieHandler.SearchMovies)
	groupRoute.GET("/detail/:id", movieHandler.DetailMovieById)
}
