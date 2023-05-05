package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

type Pagination struct {
	Total int64 `json:"total"`
	PerPage int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage int `json:"last_page"`
	FirstPageUrl string `json:"first_page_url"`
	LastPageUrl string `json:"last_page_url"`
	NextPageUrl *string `json:"next_page_url"`
	PrevPageUrl *string `json:"prev_page_url"`
	Path string `json:"path"`
	From int `json:"from"`
	To int `json:"to"`
	Data interface{} `json:"data"`
}

func NewPagination(c echo.Context) *Pagination {
	page,err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	perPage,err := strconv.Atoi(c.QueryParam("perPage"))
	if err != nil {
		perPage = 30
	}


	return &Pagination{
		PerPage: perPage,
		CurrentPage: page,
		Path: fmt.Sprintf("%s://%s%s",c.Scheme() ,c.Request().Host , c.Request().URL.Path),
		FirstPageUrl: fmt.Sprintf("%s://%s%s?page=1",c.Scheme() ,c.Request().Host ,  c.Request().URL.Path),
	}
}


func (p *Pagination) GetOffset() int {
	return (p.GetCurrentPage() - 1) * p.GetPerPage()
}

func (p *Pagination) GetPerPage() int {
	return p.PerPage
}

func (p *Pagination) GetCurrentPage() int {
	if p.CurrentPage == 0 {
		p.CurrentPage = 1
	}
	return p.CurrentPage
}


