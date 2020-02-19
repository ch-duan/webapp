package model

//Session session结构体
type Session struct {
	SessionID string
	UserName  string
	UserID    int
	Cart      *Cart
}
