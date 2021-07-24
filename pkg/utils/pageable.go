package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Pagination 自定义分页
type Pagination struct {
	CurrentPageNumber int
	TotalSize         int
	TotalPageNumber   int
	HasNext           bool
	HasLast           bool
	NextPageNumber    int
	LastPageNumber    int

	// 判断当前页与首页的偏移
	HasStart   bool
	HasLeftDot bool
	// 当前页页尾页的偏移
	HasEnd      bool
	HasRightDot bool
	PageRange   []int
}

// GetPageParams 获取正确的offset 和 limit
func GetPageParams(pageStr, pageSizeStr string) (int, int) {
	page, _ := strconv.Atoi(pageStr)
	if page <= 0 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	offset := (page - 1) * pageSize
	return offset, pageSize
}

// CheckPage 判断当前页面数 是否正常
func CheckPage(p, limit, total int) (bool, int) {
	var pn int
	if total%limit == 0 {
		pn = (total - total%limit) / limit
	} else {
		pn = (total-total%limit)/limit + 1
	}
	if p >= pn {
		return false, pn
	} else {
		return true, p
	}
}

// GetPagination 获取当前分页
func GetPagination(p, pagesize, count, total int) *Pagination {
	// p 最小为1
	//总页数 pn 最小为0
	var (
		pn, nextNumber, lastNumber int
	)
	var (
		hasNext, hasLast bool
	)
	if total%pagesize == 0 {
		pn = (total - total%pagesize) / pagesize
	} else {
		pn = (total-total%pagesize)/pagesize + 1
	}
	if p >= pn {
		p = pn
		nextNumber = pn
		lastNumber = pn - 1
		hasNext = false
		if p > 1 {
			hasLast = true
		} else {
			hasLast = false
		}
	} else {
		nextNumber = p + 1
		lastNumber = p - 1
		hasNext = true
		if p > 1 {
			hasLast = true
		} else {
			hasLast = false
		}
	}
	hasStart := true
	hasEnd := true
	hasLeftDot := false
	hasRightDot := false
	// 待省略号的分页
	list := []int{}
	// 当前页与首页的差值
	leftOffset := p - 1
	if leftOffset > 3 {
		list = append(list, p-2, p-1)
		hasStart = true
		hasLeftDot = true
	} else if leftOffset == 3 {
		list = append(list, p-2, p-1)
		hasStart = true
		hasLeftDot = false
	} else {
		for i := leftOffset; i > 0; i-- {
			list = append(list, p-i)
		}
		hasStart = false
		hasLeftDot = false
	}
	list = append(list, p)
	// 当前页与末页的差值
	rightOffset := pn - p
	if rightOffset > 3 {
		list = append(list, p+1, p+2)
		hasEnd = true
		hasRightDot = true
	} else if rightOffset == 3 {
		list = append(list, p+1, p+2)
		hasEnd = true
		hasRightDot = false
	} else {
		for i := 1; i <= rightOffset; i++ {
			list = append(list, p+i)
		}
		hasEnd = false
		hasEnd = false
	}

	pagination := &Pagination{
		CurrentPageNumber: p,
		TotalSize:         total,
		TotalPageNumber:   pn,
		HasLast:           hasLast,
		HasNext:           hasNext,
		NextPageNumber:    nextNumber,
		LastPageNumber:    lastNumber,
		PageRange:         list,
		HasStart:          hasStart,
		HasLeftDot:        hasLeftDot,
		HasRightDot:       hasRightDot,
		HasEnd:            hasEnd,
	}
	return pagination
}

// Paginate 数据库分页
func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("pageNo"))
		if page == 0 {
			page = 1
		}
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
