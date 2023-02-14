package movieController

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	model "github.com/david-side-protocol-technical/models"
	util "github.com/david-side-protocol-technical/utils"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	SearchMovies(ctx *gin.Context) (*[]model.EntityMovie, uint16)
	DetailMovieById(ctx *gin.Context) (*model.EntityMovie, uint16)
}

type controller struct {
	omdbEndPoint string
	omdbApiKey string
}

type OmdbSearchResponse struct {
	Response string                `json:"Response"`
	Search   []model.EntityMovie   `json:"Search"`
}

func NewMovieController() *controller {
	return &controller{omdbEndPoint: util.GodotEnv("OMDB_ENDPOINT"), omdbApiKey: util.GodotEnv("OMDB_API_KEY")}
}

func (ctrl *controller) FetchOmdbData(ctx *gin.Context) ([]model.EntityMovie, uint16) {
	var statusCode uint16 = 200
	title := ctx.DefaultQuery("title", "Batman")
	page := ctx.DefaultQuery("page", "1")

	omdbUrl := fmt.Sprintf("%s?apikey=%s&s=%s&page=%s", ctrl.omdbEndPoint, ctrl.omdbApiKey, title, page)

	omdbRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, omdbUrl, bytes.NewBuffer(nil))
	if err != nil {
		statusCode = 503
	}

	omdbClient := &http.Client{}
	omdbResponse, err := omdbClient.Do(omdbRequest)
	if err != nil {
		statusCode = 404
	}
	defer omdbResponse.Body.Close()

	var decodedRes OmdbSearchResponse
	err = json.NewDecoder(omdbResponse.Body).Decode(&decodedRes)
	if err != nil {
		statusCode = 500
	}
	return decodedRes.Search, statusCode
}

func (ctrl *controller) SearchMovies(ctx *gin.Context) (*[]model.EntityMovie, uint16) {
	data, statusCode := ctrl.FetchOmdbData(ctx)
	return &data, statusCode
}

func (ctrl *controller) DetailMovieById(ctx *gin.Context) (*model.EntityMovie, uint16) {
	id := ctx.Param("id")
	data, statusCode := ctrl.FetchOmdbData(ctx)
	
	var result model.EntityMovie
	for i := range data {
    if data[i].ImdbID == id {
			result = data[i]
    }
	}

	return &result, statusCode
}
