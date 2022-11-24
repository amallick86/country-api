package api

import (
	"database/sql"
	"encoding/json"
	"github.com/amallick86/country-api/db"
	"github.com/amallick86/country-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getCountriesList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Iso2 string `json:"iso2"`
}
type successResponse struct {
	Message string `json:"message"`
}

// Save country data to database by fetching from third party api
// @Summary  Save country data to database by fetching from third party api
// @Tags Country
// @ID getCountryByAPI
// @Accept json
// @Produce json
// @Success 201 {object} successResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /country/add [get]
func (server *Server) getCountryByAPI(ctx *gin.Context) {

	url := "https://api.countrystatecity.in/v1/countries"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("X-CSCAPI-KEY", server.config.CountryStateAPIToken)
	res, _ := client.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		var errResp map[string]interface{}
		json.NewDecoder(res.Body).Decode(&errResp)
		ctx.JSON(res.StatusCode, errResp)
		return
	}
	var resp []getCountriesList
	json.NewDecoder(res.Body).Decode(&resp)
	for _, item := range resp {
		_, err := server.store.AddCountry(ctx, db.AddCountryParams{
			ID:               item.Id,
			Name:             item.Name,
			CountryShortName: item.Iso2,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusCreated, successResponse{Message: "Data is Successfully saved in database"})
}

type countriesListResponse struct {
	Countries []models.Country `json:"countries"`
}

// fetch countries List
// @Summary  get countries list
// @Tags Country
// @ID getCountriesList
// @Accept json
// @Produce json
// @Success 200 {object} countriesListResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /country/list [get]
func (server *Server) getCountriesList(ctx *gin.Context) {
	response, err := server.store.GetCountriesList(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := countriesListResponse{
		Countries: response,
	}
	ctx.JSON(http.StatusOK, resp)
}
