package dao

import (
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	log.Println("开始测试")
	time1 := time.Now()
	log.Println(time1, time1.Format("2006-01-02 15:04:05"), time.Local)
	timeStr := time.Now().Format("2006-01-02 15:04:05")

	w, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	log.Println(w)
	// utils.DbIDUpdate("users")
	// utils.DbIDUpdate("books")
	// utils.DbIDUpdate("cartitems")
	// utils.DbIDUpdate("orderitems")
	log.Println("结束测试")
}
