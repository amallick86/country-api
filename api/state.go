package api

import (
	"database/sql"
	"encoding/json"
	"github.com/amallick86/country-api/db"
	"github.com/amallick86/country-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type getStateList struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	CountryId   int         `json:"country_id"`
	CountryCode string      `json:"country_code"`
	Iso2        string      `json:"iso2"`
	Type        interface{} `json:"type"`
	Latitude    string      `json:"latitude"`
	Longitude   string      `json:"longitude"`
}

// Save state data to database by fetching from third party api
// @Summary  Save state data to database by fetching from third party api
// @Tags State
// @ID getStateByAPI
// @Accept json
// @Produce json
// @Success 201 {object} successResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /state/add [get]
func (server *Server) getStateByAPI(ctx *gin.Context) {

	url := "https://api.countrystatecity.in/v1/states"
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
	var resp []getStateList
	json.NewDecoder(res.Body).Decode(&resp)
	for _, item := range resp {
		_, err := server.store.AddState(ctx, db.AddStateParams{
			ID:        item.Id,
			StateName: item.Name,
			CountryId: item.CountryId,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	ctx.JSON(http.StatusCreated, successResponse{Message: "Data is Successfully saved in database"})
}

type stateListResponse struct {
	Sates []models.State `json:"states"`
}

// fetch states List
// @Summary  get states list
// @Tags State
// @ID getStatesList
// @Accept json
// @Produce json
// @Success 200 {object} stateListResponse
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /state/list [get]
func (server *Server) getStatesList(ctx *gin.Context) {
	response, err := server.store.GetStateList(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	resp := stateListResponse{
		Sates: response,
	}
	ctx.JSON(http.StatusOK, resp)
}

type stateListResponseByCountry struct {
	Sates []models.State `json:"states"`
}

// get country states
// @Summary Get states by the country name
// @Tags State
// @ID StateByCountry
// @Accept json
// @Produce json
// @Param        country   path      string  true  "country name"
// @Success 200 {object} stateListResponseByCountry
// @Failure 400 {object} Err
// @Failure 500 {object} Err
// @Router /state/country-state/{country} [get]
func (server *Server) stateByCountry(ctx *gin.Context) {
	country := ctx.Param("country")
	data, err := server.store.GetStateListByCountry(ctx, db.GetStateListByCountryParams{
		Name: strings.ToLower(country),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	res := stateListResponseByCountry{
		Sates: data,
	}

	ctx.JSON(http.StatusCreated, res)
}
