package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	util "github.com/david-side-protocol-technical/utils"
	"github.com/go-playground/assert/v2"
	"github.com/restuwahyu13/go-supertest/supertest"
	. "github.com/smartystreets/goconvey/convey"
)

var router = SetupRouter()

func TestSearchMovies(t *testing.T) {

	Convey("Test Handler Search Movies By Category and Page", t, func() {

		Convey("Results All Movies", func() {

			test := supertest.NewSuperTest(router, t)

			test.Get("/api/v1/movies/search/")
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				var objects []map[string]interface{}
				encoded := util.Strigify(response.Data)
				_ = json.Unmarshal(encoded, &objects)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "Fetched 10 movies successfully", response.Message)
			})
		})
	})
}


func TestDetailMovieById(t *testing.T) {

	Convey("Test Handler Detail Movie By imdbID", t, func() {

		Convey("Results One Movie", func() {

			ID := "tt0372784"

			test := supertest.NewSuperTest(router, t)

			test.Get("/api/v1/movies/detail/" + ID)
			test.Send(nil)
			test.Set("Content-Type", "application/json")
			test.End(func(req *http.Request, rr *httptest.ResponseRecorder) {

				response := util.Parse(rr.Body.Bytes())
				t.Log(response)

				var object map[string]interface{}
				encoded := util.Strigify(response.Data)
				_ = json.Unmarshal(encoded, &object)

				assert.Equal(t, http.StatusOK, rr.Code)
				assert.Equal(t, http.MethodGet, req.Method)
				assert.Equal(t, "Fetched detail successfully", response.Message)
			})
		})
	})
}