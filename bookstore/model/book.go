package model

//Books 图书结构体
type Books struct {
	ID             int
	Title          string  //书名
	Author         string  //作者
	Price          float64 //单价
	Sales          int     //销售数据
	Stock          int     //库存
	classification string  //分类
	publisher      string  //出版商
	ImgPath        string  //图书图片路径
	ebook          bool    //是否电子书
}
