package utils

import (
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type BaseIndexDefaultParams struct {
	TypeDt  string
	Page    int
	Limit   int
	Offset  int
	OrderBy string
	SortBy  string
	Search  string
	Where   string
}

func BaseCrudIndexDefaultParams(c echo.Context) BaseIndexDefaultParams {
	typeDt := "colection"
	if c.QueryParam("type") == "pagination" {
		typeDt = "pagination"
	}

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 0
	}

	offset := (page - 1) * limit

	sort := c.QueryParam("sort")
	order := c.QueryParam("order_by")
	where := c.QueryParam("where")

	orderBy := ""
	sortBy := ""

	if sort != "" && order != "" {
		orderBy = order
		if strings.ToUpper(sort) == "DESC" {
			sortBy = "DESC"
		} else {
			sortBy = "ASC"
		}
	}

	search := strings.ToLower(c.QueryParam("search"))

	return BaseIndexDefaultParams{
		TypeDt:  typeDt,
		Page:    page,
		Limit:   limit,
		Offset:  offset,
		OrderBy: orderBy,
		SortBy:  sortBy,
		Search:  search,
		Where:   where,
	}

}

func BaseCrudIndexDefaultQuery(params BaseIndexDefaultParams, query *gorm.DB, searchAble []string) {
	if params.Search != "" {
		for i := 0; i < len(searchAble); i++ {
			query.Where(searchAble[i]+" ILIKE ?", "%"+params.Search+"%")
		}
	}

	if params.Where != "" {
		query.Where("(" + params.Where + ")")
	}
}

func BaseCrudIndexDefaultLimitQuery(params BaseIndexDefaultParams, query *gorm.DB, totalItems *int64) {
	if params.TypeDt != "colection" {
		query.Count(totalItems)
	}

	if params.OrderBy != "" && params.SortBy != "" {
		query.Order(params.OrderBy + " " + params.SortBy)
	}

	if params.Limit > 0 {
		query.Limit(params.Limit)
	}

	if params.Page > 0 && params.Limit > 0 {
		query.Offset((params.Page - 1) * params.Limit)
	}
}
