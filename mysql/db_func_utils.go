package mysql

import (
	"log"
	"time"
)

func getNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func getSkip(page, limit int) int {
	log.Println("getSkip", page, limit)
	if page != 0 && limit != 0 {
		return (page - 1) * limit
	}
	return 0
}

func getLimit(limit int) int {
	log.Println("getLimit", limit)
	if limit != 0 {
		return limit
	}
	return 10
}

func getTotalPages(total int, limit int) int {
	log.Println("getTotalPages", total, limit)
	if limit != 0 {
		return (int(total) + limit - 1) / limit
	}
	return 1
}
