package handler

import (
	"fmt"
	"github.com/labstack/echo"
	"imagedb/database"
	"imagedb/model"
	"net/http"
	"strconv"
)

type ReqGameImage struct {
	GameID  uint   `json:"game_id"`
	Address string `json:"address"`
}

func Upload(c echo.Context) error {
	db := database.GetDB()
	var request ReqGameImage
	err := c.Bind(&request)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	image := &model.GameImage{
		GameID:       request.GameID,
		ImageAddress: request.Address,
	}

	if err := db.Create(image).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, "created")
}

func GetImage(c echo.Context) error {
	gameID, err := strconv.Atoi(c.Param("gameID"))
	fmt.Println(gameID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid path parameter", err)

	}
	db := database.GetDB()
	var images []model.GameImage
	if err := db.Model(&model.GameImage{}).Where("game_id = ?", uint(gameID)).Find(&images).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "cant get image from db", err)
	}
	response := NewimgListResponse(images)
	return c.JSON(http.StatusOK, response)
}

func NewimgListResponse(list []model.GameImage) *prjListResponse {
	prjs := make([]*prjResponse, 0)
	for i := range list {
		prjs = append(prjs, NewPrjResponse(&list[i]))
	}
	return &prjListResponse{
		Projects: prjs,
	}
}

type prjResponse struct {
	Address string `json:"address"`
}

type prjListResponse struct {
	Projects []*prjResponse `json:"images"`
}

func NewPrjResponse(prj *model.GameImage) *prjResponse {
	return &prjResponse{
		Address: prj.ImageAddress,
	}
}
