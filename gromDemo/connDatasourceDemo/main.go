package main

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	Id         string    `gorm:"primary_key;not null;varchar(100);column:id;comment:'主键Id';" json:"id"`
	Name       string    `gorm:"varchar(255);column:name;comment:'名称';" json:"name"`
	Age        int       `gorm:"int;column:age;comment:'年龄';" json:"age"`
	Address    string    `gorm:"varchar(255);column:address;comment:'地址';" json:"address"`
	Email      string    `gorm:"varchar(255);column:email;comment:'邮箱';" json:"email"`
	CreateTime time.Time `gorm:"datetime;column:create_time;comment:'创建时间';" json:"createTime"`
}

/*
*
完成数据库表名和struct的映射
*/
func (User) TableName() string {

	return "t_user"

}

func main() {
	//测试连接数据库
	dsn := "root:pass@tcp(192.168.150.129:3306)/exer?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("数据库初始化异常，%s", err.Error())
	}
	//查询id为1的数据
	user := &User{}
	db.Raw("SELECT id, name, age FROM t_user WHERE id = ?", 1).Scan(user)
	//jsonUser, err := json.Marshal(user)
	//fmt.Println(jsonUser)
	fmt.Println(user)
	//生成UUId
	s := uuid.New().String()
	fmt.Println(s)
	idstr := strings.ReplaceAll(s, "-", "")
	createUser1 := User{Id: idstr, Name: "王五" + idstr, Age: 23, Address: "上海市浦东新区", Email: "221@vip.com", CreateTime: time.Now()}
	result := db.Create(&createUser1)
	if result.Error != nil {
		fmt.Sprintf("数据插入失败%s", result.Error.Error())
	}
	affected := result.RowsAffected
	fmt.Println(affected)
	if affected == 1 {
		fmt.Sprintf("数据插入成功，其插入数据的ID为: %s", idstr)
		fmt.Println("数据插入成功，其插入数据的ID为: %s", idstr)
	}

}
