package model

//结构体的两个属性首字母为大写，表示声明的是全局作用域可见的(标识符首字母大写public, 首字母小写包作用域可见
//结构体中还使用了标签(Tag)。这些标签在encoding/json和encoding/xml中有特殊应用。
type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	ServedBy string `json:"servedBy"`
	Quote Quote `json:"quote"`
}

// NEW struct
type Quote struct {
	Text string `json:"quote"`
	ServedBy string `json:"ipAddress"`
	Language string `json:"language"`
}