package controller

import (
	"github.com/labstack/echo"
	"github.com/yejiansnake/go-yedb"
	"net/http"
	"../model"
	"../helper"
	"../sys"
)

const testPageSize int = 50

type TestController struct {
	Model *yedb.DbModel
}

func StockControllerNew() *TestController {
	res := TestController{}
	res.Model = model.GetTestModel()
	res.setRoute()
	return &res
}

func (ptr *TestController) setRoute() {
	sys.AppInstance.GET("/test", ptr.testGet)
	sys.AppInstance.GET("/test/:id", ptr.testGetOne)
}

func (ptr *TestController) testGet(content echo.Context) error {
	pageIndex, _ := helper.StrToInt(content.QueryParam("page_index"))

	var data []model.Test
	pageData, err := helper.GetPageData(
		ptr.Model.Find(),
		&data,
		pageIndex,
		testPageSize)

	if err != nil {
		return err
	}

	return content.JSON(http.StatusOK, &pageData)
}

func (ptr *TestController) testGetOne(content echo.Context) error {
	id := content.Param("id")

	stock := model.Test{}
	err := ptr.Model.Find().AndWhere(&yedb.DbParams{"id": id}).FillRow(&stock)
echo.New()
	if err != nil {
		return err
	}

	return content.JSON(http.StatusOK, &stock)
}
