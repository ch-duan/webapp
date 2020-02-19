package model

//Pages 页面结构体
type Pages struct {
	Books        []*Books
	PageNum      int //页号
	PageSize     int //每页显示的条数
	TotalPageNum int //总页数
	TotalRecurd  int //总记录
	IsLogin      bool
	UserName     string
}

//IsHasPrev 判断是否有上一页
func (p *Pages) IsHasPrev() bool {
	return p.PageNum > 1
}

//IsHasNext 判断是否有下一页
func (p *Pages) IsHasNext() bool {
	return p.PageNum < p.TotalPageNum
}

//GetPrevPageNum 获取上一页
func (p *Pages) GetPrevPageNum() int {
	if p.IsHasPrev() {
		return p.PageNum - 1
	}
	return 1
}

//GetNextPageNum 获取下一页
func (p *Pages) GetNextPageNum() int {
	if p.IsHasNext() {
		return p.PageNum + 1
	}
	return p.TotalPageNum

}
