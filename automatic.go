package main
//whr-helen 2019

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	_"strings"
	"io/ioutil"
	"flag"
	"os"
	"strings"
)
//用户输入配置
var(
	db_host string //数据库地址
	db_port string //数据库端口
	db_name string //数据库名称
	db_account string //数据库账号
	db_pwd string //数据库密码
	path string		//结构体保存路径
	tables string  //表名称
)
//表类型
type Dbtabels struct{
	Name string `json:"name"`
}
//数据字段类型
type Column struct {
	Columname string `json:"columname"`
	Datatype string `json:"datatype"`
	Columncomment string `json:"columncomment"`
	Columnkey	string `json:"columnkey"`
	Extra string	`json:"extra"`
}

func init() {
	flag.StringVar(&db_account, "acc", "root", "# Database account")
	flag.StringVar(&db_pwd, "pwd", "123123", "# Database password")
	flag.StringVar(&db_host, "host", "127.0.0.1", "# Database host")
	flag.StringVar(&db_port, "port", "3306", "# Database port")
	flag.StringVar(&db_name, "d", "", "# Database name")
	flag.StringVar(&path, "path", "./models", "# Structure preservation path")
	flag.StringVar(&tables, "t", "all", "# Table name formats such as - t user, rule, config")

}

func main(){
	flag.Parse()
	if db_name ==""{
		flag.Usage()
		return
	}
	tables = convtables(tables)//转换

	fmt.Println("数据库:",db_name)
	fmt.Println("数据库账号:",db_account)
	fmt.Println("数据库密码:",db_pwd)
	fmt.Println("结构体保存路径:",path)
	fmt.Println("指定数生成据表:",tables)
	fmt.Println("Automatic Struct Start ...")
	if checkpath(path){
		fmt.Println("paPath irregularityth  Example：'./models'")
		return
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",db_account,db_pwd,db_host,db_port,db_name)
	db, err := sql.Open("mysql", dataSourceName)

	if err!=nil{
		fmt.Println("mysql connect err:",err)
		return
	}
	//获取所有表
	var sqls = "select table_name from information_schema.tables where table_schema=?"
	if tables != "all"{
		 sqls = fmt.Sprintf("select table_name from information_schema.tables where table_schema=? and table_name in (%s)",tables)
	}
	rows,err2 := db.Query(sqls,db_name)
	if err2!=nil{
		fmt.Println("mysql query err:",err2)
		return
	}
	defer func() {
		if rows != nil {
			rows.Close() //可以关闭掉未scan连接一直占用
		}
	}()

	table := Dbtabels{}
	for rows.Next(){
		err := rows.Scan(&table.Name)
		fmt.Println("正在生成表结构：",table.Name)
		if err != nil {
			fmt.Printf("Scan failed,err:%v", err)
			return
		}
		//获取单个表所有字段
		cloumn_sql := fmt.Sprintf("select column_name columnName, data_type dataType, column_comment columnComment, column_key columnKey, extra from information_schema.columns where table_name = '%s' and table_schema = (select database()) order by ordinal_position",table.Name)
		cloumns,err3:= db.Query(cloumn_sql)
		if err3!=nil{
			fmt.Println("mysql 获取单个表所有字段 err:",err3)
		}
		defer func() {
			if cloumns != nil {
				cloumns.Close() //可以关闭掉未scan连接一直占用
			}
		}()
		struct_str :=fmt.Sprintf("type %s struct { \n",strFirstToUpper(table.Name))
		column := Column{}
		for cloumns.Next(){
			err := cloumns.Scan(&column.Columname,&column.Datatype,&column.Columncomment,&column.Columnkey,&column.Extra)
			//类型判断
			//开始拼接字符串
			if err != nil {
				fmt.Printf("Scan failed,err:%v", err)
				return
			}
			struct_str += "    "+strFirstToUpper(column.Columname)
			if column.Datatype == "int" || column.Datatype =="tinyint"{
				struct_str +=" int "
			}else if column.Datatype == "decimal"{
				struct_str +=" float64 "
			}else{
				struct_str +=" string "
			}

			if column.Extra != "auto_increment"{
				struct_str += fmt.Sprintf("`xorm:\"comment('%s')\" json:\"%s\"` \n",column.Columncomment,column.Columname)
			}else{
				struct_str += fmt.Sprintf("`xorm:\"not null pk autoincr comment('%s') INT(11)\" json:\"%s\"` \n",column.Columncomment,column.Columname)
			}

		}
		struct_str += "}"
		model_head := "package models \n\n"
		//导出文件
		body := model_head+struct_str
		filename:=fmt.Sprintf("%s/%s.go",path,table.Name)
		//创建文件夹
		error2 := os.MkdirAll(path, os.ModePerm)
		if error2 != nil{
			fmt.Println("midkr path err:",error2)
		}
		err4:=ioutil.WriteFile(filename,[]byte(body),0666)
		if err4 != nil{
			fmt.Println("写入文件错误:",err4)
		}
	}

	fmt.Println("End  SUCCESS")
}

//首字母大写
func strFirstToUpper(str string) string {
	var upperStr string
	vv := []rune(str)   // 后文有介绍
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {  // 后文有介绍
				vv[i] -= 32 // string的码表相差32位
				upperStr += string(vv[i])
			} else {
				fmt.Println("Not begins with lowercase letter,")
				return str
			}
		} else {
			upperStr += string(vv[i])
		}
	}
	return upperStr
}

//判断用户输入
func checkpath(path string)bool{
	a := strings.Split(path, "/")
	if len(a)>=2 && a[1] != ""{
		return false
	}else{
		return true
	}

}

//用户输入转换
func convtables(tab string)string{
	if tab=="all"{
		return tab
	}else{
		str_arr := strings.Split(tab, ",")
		var tabs string
		for _,v := range str_arr {
			if v!=""{
				item := fmt.Sprintf("'%s',",v)
				tabs += item
			}
		}
		//'mall_account','mall_goods'
		return tabs[0:len(tabs)-1]
	}
}
