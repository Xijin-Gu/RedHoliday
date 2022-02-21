package dao

import (
	"RedHoliday/model"
	"fmt"
)

//该go文件下存放查询数据库的函数

//通过用户名查询密码，
func Query_up(username string)(pw,salt string){
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	querystr := "select password,salt from users where username=?"
	//查询信息
	_ = db.QueryRow(querystr,username).Scan(&u.Password,&u.Salt)
	return u.Password,u.Salt
}

func Query_uid(username string)int{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var u model.Users_mysql
	querystr := "select uid from users where username=?"
	//查询信息
	_ = db.QueryRow(querystr,username).Scan(&u.Uid)
	return u.Uid
}



//通过username查询个人信息表
func Query_username_introduction(username string) model.Person_mysql {
	//连接数据库
	db := Link_mysql()
	//定义参数
	var user model.Person_mysql
	querystring := "select * from "+username+";"
	//提取信息
	r := db.QueryRow(querystring).Scan(&user.Uid, &user.Username,&user.Avatar,&user.Friends,&user.Follow_business,  &user.Shopping_cart,&user.Order_paid,&user.Order_unpaid,&user.Order_received)
	fmt.Println(r,user)
	return user
}


//通过uid查询商品详细信息
func Query_commmidty(uid int)model.Commidity{
	//连接数据库
	db := Link_mysql()
	//定义相关字段
	var c model.Commidity
	querystr := "select * from commidity where uid=?"
	//查询信息
	r := db.QueryRow(querystr,uid).Scan(&c.Uid,&c.Commidity_name,&c.Volume,&c.Evaluations,&c.Detailed_Introduction)
	fmt.Println(r)
	return c
}