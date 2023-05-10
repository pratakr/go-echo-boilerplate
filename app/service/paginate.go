package service

import (
	"fmt"
	"kancha-api/app/utils"
	"math"

	"gorm.io/gorm"
)

func paginate(pagination *utils.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var total int64
	//db.Model(value).Count(&total)
	db.Count(&total)
	//fmt.Printf("Total=%d\n",total)
	totalPages := int(math.Ceil(float64(total) / float64(pagination.PerPage)))

	pagination.Total = total
	pagination.LastPage = totalPages
	pagination.From = (pagination.CurrentPage-1)*pagination.PerPage + 1

	if int64(pagination.CurrentPage*pagination.PerPage) >= total {
		pagination.To = int(total)
	} else {
		pagination.To = pagination.CurrentPage * pagination.PerPage
	}

	pagination.LastPageUrl = fmt.Sprintf("%s?page=%d", pagination.Path, totalPages)

	if pagination.CurrentPage < totalPages {
		nextPageUrl := fmt.Sprintf("%s?page=%d", pagination.Path, pagination.CurrentPage+1)
		pagination.NextPageUrl = &nextPageUrl
	}

	if pagination.CurrentPage > 1 {
		prevPageUrl := fmt.Sprintf("%s?page=%d", pagination.Path, pagination.CurrentPage-1)
		pagination.PrevPageUrl = &prevPageUrl
	}

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetPerPage()) //.Order(pagination.GetSort())
	}
}
