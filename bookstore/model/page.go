package model

//Page 页面结构体
type Page struct {
	Books        Books
	PageNum      int
	PageSize     int
	TotalPageNum int
	TotalRecurd  int
	IsLogin      bool
	Username     string
}
