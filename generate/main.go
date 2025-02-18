package main

import (
	"github.com/dave/jennifer/jen"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm" //TODO 注意与"gorm.io/gorm"的不同,"github.com/jinzhu/gorm"是gorm的老版本
	"log"
	"strconv"
	"strings"
)

// 1.定义用于存储表字段信息的结构体
type Column struct {
	ColumnName             string `gorm:"column:COLUMN_NAME"`              //字段名称
	ColumnKey              string `gorm:"column:COLUMN_KEY"`               //索引类型
	DataType               string `gorm:"column:DATA_TYPE"`                //数据类型
	CharacterMaximumLength int    `gorm:"column:CHARACTER_MAXIMUM_LENGTH"` //设置的最大长度限制
	IsNullable             string `gorm:"column:IS_NULLABLE"`              //是否可以为空
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT"`           //默认值
	ColumnComment          string `gorm:"column:COLUMN_COMMENT"`           //注释
}

// 2.数据库地址
const dsn = "root:pass@tcp(192.168.150.129:3306)/exer?charset=utf8mb4&parseTime=True&loc=Local"

// 3.需要生成的表所在库名
const databaseName = "exer"

// 4.需要生成的表名
const tableName = "t_user"

// 5.表名的驼峰命名并且首字母大写
var TabNameCamelCaseUp = underscoreToCamelCase(tableName)

// 6.表名的驼峰命名,并且首字母小写
var TabNameCamelCaselower = lowerFirst(TabNameCamelCaseUp)

func main() {

	//1.获取数据库连接
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//2.执行sql查询指定表的所有字段信息,包括(字段名,数据类型,索引,设置的最大长度,是否允许为null,默认值,注释)
	var columns []Column
	sql := "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, CHARACTER_MAXIMUM_LENGTH, IS_NULLABLE, COLUMN_DEFAULT, COLUMN_COMMENT  FROM information_schema.columns WHERE table_schema = '" + databaseName + "' AND table_name = '" + tableName + "';"
	err = db.Debug().Raw(sql).Scan(&columns).Error
	if err != nil {
		panic(err)
	}

	//3.使用jen开始创建并生成对应表的代码,指定代码所在包
	f := jen.NewFile("main")

	//4.封装结构体
	f.Type().Id(TabNameCamelCaseUp).StructFunc(func(g *jen.Group) {
		for _, c := range columns {

			//4.1将数据库字段名转换为结构体需要的属性名
			statement := g.Id(underscoreToCamelCase(c.ColumnName))
			//4.2获取属性类型,将mysql数据类型转换为golang中使用的
			//TODO 注意此处只是简单的判断了几个类型,要根据实际需求进行转换
			switch c.DataType {
			case "int", "tinyint", "smallint":
				statement.Int()
			case "mediumint", "bigint":
				statement.Qual("github.com/shopspring/decimal", "Decimal")
			case "float":
				statement.Float32()
			case "double":
				statement.Float64()
			case "decimal":
				statement.Qual("github.com/shopspring/decimal", "Decimal")
			case "char", "varchar", "text", "tinytext", "mediumtext", "longtext", "blob":
				statement.String()
			case "date", "time", "datetime", "timestamp":
				statement.Qual("time", "Time")
			case "bit", "bool":
				statement.Bool()
			case "json":
				statement.Interface()
			default:
				statement.Interface()
			}
			//4.3获取属性后的tag,设置tag
			tagM := getTagMap(c)
			statement.Tag(tagM)
		}
	})

	f.Comment("TODO 代表数据库连接的全局变量,项目启动时先初始化数据库连接,在执行操作数据库方法时就可以直接使用全局变量的连接,不要来回传递了," +
		"此处是为了防止生成的代码报错添加的,项目中如果已经存在数据库连接变量,将生成代码中的这个变量删除,直接使用已经存在的就可以")
	f.Var().Id("db").Op("*").Qual("github.com/jinzhu/gorm", "DB")

	//6.生成TableName方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("TableName").Params().String().Block(
		jen.Return(jen.Lit(tableName)),
	)

	//7.生成Add方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("Add").Params().Error().Block(
		jen.Return(jen.Id("db").Dot("Create").Call(jen.Id("c")).Dot("Error")),
	)

	//8.生成Update方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("Update").Params().Error().Block(
		jen.Return(jen.Id("db").Dot("Model").Call(jen.Id("c")).Dot("Update").Call(jen.Id("c")).Dot("Error")),
	)

	//9.生成UpdateSave方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("UpdateSave").Params().Error().Block(
		jen.Return(jen.Id("db").Dot("Model").Call(jen.Id("c")).Dot("Save").Call(jen.Id("c")).Dot("Error")),
	)

	//10.生成SearchFirst方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("SearchFirst").Params().Error().Block(
		jen.Return(jen.Id("db").Dot("Model").Call(jen.Id("c")).Dot("Where").Call(jen.Id("c")).Dot("First").Call(jen.Id("c")).Dot("Error")),
	)

	//11.生成Search方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("Search").Params().Params(jen.Index().Id(TabNameCamelCaseUp), jen.Error()).Block(
		jen.Var().Id(TabNameCamelCaselower+"List").Index().Id(TabNameCamelCaseUp),
		jen.Err().Op(":=").Id("db").Dot("Model").Call(jen.Id("c")).Dot("Where").Call(jen.Id("c")).Dot("Find").Call(jen.Op("&").Id(TabNameCamelCaselower+"List")).Dot("Error"),
		jen.Return(jen.Id(TabNameCamelCaselower+"List"), jen.Err()),
	)

	//12.生成SelectById方法
	f.Func().Id("SelectById").Params(jen.Id("id").Int()).Params(jen.Op("*").Id(TabNameCamelCaseUp), jen.Error()).Block(
		jen.Var().Id(TabNameCamelCaselower).Id(TabNameCamelCaseUp),
		jen.Err().Op(":=").Id("db").Dot("First").Call(jen.Op("&").Id(TabNameCamelCaselower), jen.Id("id")).Dot("Error"),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Err()),
		),
		jen.Return(jen.Op("&").Id(TabNameCamelCaselower), jen.Nil()),
	)

	//13.生成SelectByPage方法
	f.Func().Params(jen.Id("c").Op("*").Id(TabNameCamelCaseUp)).Id("SelectByPage").Params(jen.Id("pageNo").Int(), jen.Id("pageSize").Int()).Params(jen.Index().Id(TabNameCamelCaseUp), jen.Int(), jen.Error()).Block(
		jen.Var().Id(TabNameCamelCaselower+"List").Index().Id(TabNameCamelCaseUp),
		jen.Var().Id("count").Int(),
		jen.Id("db").Op("=").Id("db").Dot("Model").Call(jen.Op("&").Id(TabNameCamelCaseUp+"{}")).Dot("Where").Call(jen.Id("c")),
		jen.Id("limit").Op(":=").Id("pageSize"),
		jen.Id("offset").Op(":=").Id("pageSize").Op("*").Call(jen.Id("pageNo").Op("-").Lit(1)),
		jen.Comment("注意Limit 方法和 Offset 方法必须在 Find 方法之前调用，否则会出现错误。"),
		jen.Err().Op(":=").Id("db").Dot("Count").Call(jen.Op("&").Id("count")).Dot("Limit").Call(jen.Id("limit")).Dot("Offset").Call(jen.Id("offset")).Dot("Find").Call(jen.Op("&").Id(TabNameCamelCaselower+"List")).Dot("Error"),
		jen.If(jen.Err().Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Lit(0), jen.Err()),
		),
		jen.Return(jen.Id(TabNameCamelCaselower+"List"), jen.Id("count"), jen.Nil()),
	)

	//14.保存文件
	err = f.Save(tableName + ".go")
	if err != nil {
		log.Fatal(err)
	}
}

// 组装属性tag
// TODO 在组装tag时,并不是越多越好,例如默认值, 是否允许为空,数据类型,数据长度等,如果代码中使用不到,就不要组装了,因为在后续迭代过程中,由于遗漏,可能会出现数据库表设置与当前代码tag中的不一致
func getTagMap(column Column) map[string]string {
	tagStr := ""
	//拼接索引
	//TODO 当前只判断了主键索引
	if column.ColumnKey == "PRI" {
		tagStr += "primary_key;"
	}

	if column.IsNullable == "NO" {
		tagStr += "not null;"
	}

	tagStr += column.DataType
	if column.CharacterMaximumLength > 0 {
		tagStr += "(" + strconv.Itoa(column.CharacterMaximumLength) + ")"
	}
	tagStr += ";"

	tagStr += "column:" + column.ColumnName + ";"
	if len(column.ColumnComment) > 0 {
		tagStr += "comment:'" + column.ColumnComment + "';"
	}

	if len(column.ColumnDefault) > 0 {
		tagStr += "default:" + column.ColumnDefault + ";"
	}

	m := make(map[string]string)
	m["gorm"] = tagStr
	m["json"] = lowerFirst(lowerFirst(underscoreToCamelCase(column.ColumnName)))
	return m
}

// 将带下划线的数据库命名转换为驼峰命名,首字母大写
func underscoreToCamelCase(name string) string {
	result := ""
	if !strings.Contains(name, "_") {
		first := strings.Title(name[:1])
		return first + name[1:]
	}
	parts := strings.Split(name, "_")

	// 遍历字符串切片，对每个部分进行转换
	for _, part := range parts {
		// 使用Title函数，将每个部分的首字母转换为大写
		part = strings.Title(part)
		// 使用Replace函数，将每个部分的其他字母转换为小写
		part = strings.Replace(part, part[1:], strings.ToLower(part[1:]), 1)
		// 将转换后的部分拼接到结果字符串中
		result += part
	}
	// 返回结果字符串
	return result
}

// 字符串首字母小写
func lowerFirst(str string) string {
	// 判断字符串是否为空
	if str == "" {
		// 如果为空，直接返回空字符串
		return ""
	}
	// 使用strings包的ToLower方法，将字符串的第一个字符转换为小写
	first := strings.ToLower(str[:1])
	// 将转换后的第一个字符和剩余的字符串拼接起来，返回结果
	return first + str[1:]
}
