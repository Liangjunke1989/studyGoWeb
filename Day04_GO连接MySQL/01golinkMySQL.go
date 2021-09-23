package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 使用驱动，不用的时候不报错
)

func main() {
	// "用户名：密码@[连接方式](主机名：端口号)/数据库名"
	db, _ := sql.Open("mysql", "root:liang521@(127.0.0.1:3306)/MyGameDB") // 设置连接参数
	defer db.Close()                                                      // 关闭数据库
	err := db.Ping()                                                      // 连接数据库
	if err != nil {
		fmt.Println("数据库连接失败！！！", err.Error())
		return
	}
	fmt.Println("数据库连接成功！！！")

	//操作一：执行数据操作语句
	{
		//now := time.Now().Unix()
		//timeLayout := "2006-01-02 15:04:05" // go语言固定日期模版
		//timeStr := time.Unix(now, 0).Format(timeLayout)
		//sqlStr :="insert into users values(null,'tom44','1234567','"+timeStr+"')"
		//result, _ := db.Exec(sqlStr) //执行SQL语句
		//n, _:= result.RowsAffected() //获取受影响的记录数
		//fmt.Println("受影响的记录为：",n)}
	}

	//操作二：执行预处理语句
	{
		//now := time.Now().Unix()
		//timeLayout := "2006-01-02 15:04:05" // go语言固定日期模版
		//timeStr := time.Unix(now, 0).Format(timeLayout)
		//
		//user11:=[2][4] string{{"298","ketty22","123",timeStr},{"299","rose22","123",timeStr}}
		//stmt, err0:= db.Prepare("insert into users values (?,?,?,?)") //获取预处理
		//if err0!=nil {
		//	fmt.Println("文件插入有误",err0.Error())
		//	return
		//}
		//for _, user := range user11 {
		//	result, err02 := stmt.Exec(user[0], user[1],user[2],user[3]) //调用预处理语句
		//	if err02!=nil {
		//		fmt.Println("执行预处理有误！",err02.Error())
		//	}
		//	n, _ := result.RowsAffected()
		//	fmt.Println("影响到的行数为：",n)
		//}
	}

	//操作三：单行查询
	{
		//var id,name,pwd,date string
		//queryRow := db.QueryRow("select * from users where id=198") //QueryRow() 获取一行数据
		////fmt.Printf("%T",queryRow)
		//queryRow.Scan(&id,&name,&pwd,&date) //将查询到的第一行数据放到id,name,pwd,date中
		//fmt.Println(id,name,pwd,date)
	}

	//操作四：多行查询
	var id, name, pwd, date string
	rows, _ := db.Query("select * from users") //Query() 获取所有数据
	fmt.Println("ID", "姓名", "密码", "日期")
	for rows.Next() { //循环显示所有的数据
		rows.Scan(&id, &name, &pwd, &date) //将查询到的第一行数据放到id,name,pwd,date中
		fmt.Println(id, name, pwd, date)
	}
}
