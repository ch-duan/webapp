package model

//Book 图书结构体
type Book struct {
	ID             int
	Title          string  //书名
	Author         string  //作者
	Price          float64 //单价
	Sales          int     //销售数据
	Stock          int     //库存
	Classification string  //分类
	Publisher      string  //出版商
	ImgPath        string  //图书图片路径
	Ebook          bool    //是否电子书
}
