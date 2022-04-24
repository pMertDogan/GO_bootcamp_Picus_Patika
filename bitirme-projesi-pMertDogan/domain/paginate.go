package domain

import (
	"strconv"

	"gorm.io/gorm"
)

//Global paginate calculation
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) (*gorm.DB) {

		
		// page ,_:= strconv.Atoi(page)
		// pageSize, _ := strconv.Atoi(pageSize)
		page, pageSize = CalcPageAndSize(page, pageSize)

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func CalcPageAndSize(page int, pageSize int) (int, int) {
	if page == 0 {
		page = 1
	}

	if page == 0 {
		page = 1
	}
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return page, pageSize
}


func CalcPageAndSizeReturnString(page int, pageSize int) (string, string) {
	
	page , size := CalcPageAndSize(page, pageSize)

	return strconv.Itoa(page), strconv.Itoa(size)
}