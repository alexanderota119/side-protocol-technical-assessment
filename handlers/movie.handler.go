package movieHandler

import (
	"net/http"

	movieController "github.com/david-side-protocol-technical/controllers"
	util "github.com/david-side-protocol-technical/utils"
	"github.com/gin-gonic/gin"
)

type handler struct {
	movieCtrl movieController.Controller
}

func NewMovieHandler(movieCtrl movieController.Controller) *handler {
	return &handler{movieCtrl: movieCtrl}
}

func handleErrorResponse(ctx *gin.Context, statusCode uint16) {
	switch statusCode {
		case 404:
			util.APIResponse(ctx, "Not found the omdbapi endpoint", http.StatusConflict, http.MethodGet, nil)

		case 503:
			util.APIResponse(ctx, "Not ready to fetch", http.StatusConflict, http.MethodGet, nil)

		case 500:
			util.APIResponse(ctx, "Failed to decode json", http.StatusConflict, http.MethodGet, nil)

		default:
			util.APIResponse(ctx, "Something went wrong", http.StatusConflict, http.MethodGet, nil)
	}
}

func (h *handler) SearchMovies(ctx *gin.Context) {
	result, statusCode := h.movieCtrl.SearchMovies(ctx)

	if(statusCode > 400) {
		handleErrorResponse(ctx, statusCode)
	} else {
		util.APIResponse(ctx, "Fetched 10 movies successfully", http.StatusOK, http.MethodGet, result)
	}
	
}

func (h *handler) DetailMovieById(ctx *gin.Context) {
	result, statusCode := h.movieCtrl.DetailMovieById(ctx)

	if(statusCode > 400) {
		handleErrorResponse(ctx, statusCode)
	} else {
		util.APIResponse(ctx, "Fetched detail successfully", http.StatusOK, http.MethodGet, result)
	}
}
