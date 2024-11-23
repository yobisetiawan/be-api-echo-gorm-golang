package controllers

import (
	"be_api/app/database"
	"be_api/app/models"
	"be_api/app/requests"
	"be_api/app/utils"
	"be_api/app/validator"
	"net/http"
	"strconv"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ProductCategoryController struct {
}

func NewProductCategoryController() *ProductCategoryController {
	return &ProductCategoryController{}
}

func (ctr *ProductCategoryController) Index(c echo.Context) error {
	params := utils.BaseCrudIndexDefaultParams(c)

	var data []models.ProductCategory
	var totalItems int64
	searchAble := []string{"title"}
	query := database.DB.Model(data)

	utils.BaseCrudIndexDefaultQuery(params, query, searchAble)

	utils.BaseCrudIndexDefaultLimitQuery(params, query, &totalItems)

	queryData := query.Find(&data)
	if queryData.Error != nil {
		utils.Error500Log(c, queryData.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
		"meta_data": echo.Map{
			"total":        totalItems,
			"current_page": params.Page,
			"per_page":     params.Limit,
		},
	})
}

func (ctr *ProductCategoryController) Show(c echo.Context) error {
	id := c.Param("id")
	var data models.ProductCategory
	query := database.DB.Model(data)
	queryData := query.Where("id = ?", id).First(&data)

	if queryData.Error != nil {
		return utils.ErrorGeneralLog(c, 404, "Data not found!", queryData.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": data})
}

func (ctr *ProductCategoryController) Store(c echo.Context) error {

	formData := requests.ProductCategoryRequest{}

	if err := c.Bind(&formData); err != nil {
		return utils.ErrorGeneralLog(c, 400, "Invalid Form Data", err)
	}

	if err := c.Validate(formData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validator.NewValidationErrorResponse(err, formData))
	}

	data := models.ProductCategory{}
	if err := copier.Copy(&data, formData); err != nil {
		return utils.Error500Log(c, err)
	}

	result := database.DB.Create(&data)
	if result.Error != nil {
		return utils.Error500Log(c, result.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": data})
}

func (ctr *ProductCategoryController) Update(c echo.Context) error {

	id := c.Param("id")
	var data models.ProductCategory
	query := database.DB.Model(data)
	result := query.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return utils.ErrorGeneralLog(c, 404, "Data not found", result.Error)
	}

	formData := requests.ProductCategoryRequest{}

	if err := c.Bind(&formData); err != nil {
		return utils.ErrorGeneralLog(c, 400, "Invalid Form Data", err)
	}

	formData.ID = id

	if err := c.Validate(formData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, validator.NewValidationErrorResponse(err, formData))
	}

	if err := copier.Copy(&data, formData); err != nil {
		return utils.Error500Log(c, err)
	}

	result = database.DB.Save(&data)
	if result.Error != nil {
		return utils.Error500Log(c, result.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"data": data})
}

func (ctr *ProductCategoryController) Destroy(c echo.Context) error {
	id := c.Param("id")
	var data models.ProductCategory
	query := database.DB.Model(data)
	result := query.Where("id = ?", id).First(&data)
	if result.Error != nil {
		return utils.ErrorGeneralLog(c, 404, "Data not found", result.Error)
	}

	result = database.DB.Where("id = ?", id).Delete(&models.ProductCategory{})
	if result.Error != nil {
		return utils.Error500Log(c, result.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}

func (ctr *ProductCategoryController) BulkDestroy(c echo.Context) error {
	idStr := c.QueryParam("ids")
	if idStr == "" {
		return utils.ErrorGeneralLog(c, 404, "Param IDs is required!")
	}

	idStrs := strings.Split(idStr, ",")
	ids := make([]int64, 0, len(idStrs))

	for _, idStr := range idStrs {
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return utils.Error500Log(c, err)
		}
		ids = append(ids, id)
	}

	result := database.DB.Where("id IN ?", ids).Delete(&models.ProductCategory{})
	if result.Error != nil {
		return utils.Error500Log(c, result.Error)
	}

	return c.JSON(http.StatusOK, echo.Map{"success": true})
}
