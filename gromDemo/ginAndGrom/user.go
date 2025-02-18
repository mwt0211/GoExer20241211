package ginAndGrom

type User struct {
	Id   string `json:"id"`
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type Animal struct {
	ID   int64
	UUID string `gorm:"primaryKey"`
	Name string
	Age  int64
}

/*
*
完成数据库表名和struct的映射
*/
func (User) TableName() string {

	return "t_user"

}
