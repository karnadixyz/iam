package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	PaginationLimit = 20
	PaginationPage  = 1
)

type PaginationMetadata struct {
	Page         int   `json:"page"`
	Limit        int   `json:"limit"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
}

type Pagination struct {
	Metadata *PaginationMetadata `json:"metadata"`
	Records  interface{}         `json:"records"`
}

func (m *PaginationMetadata) GetOffset() int {
	return m.Limit * (m.Page - 1)
}

func GetPaginationParams(c *gin.Context) *PaginationMetadata {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = PaginationLimit
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = PaginationPage
	}

	return &PaginationMetadata{
		Limit: limit,
		Page:  page,
	}
}
