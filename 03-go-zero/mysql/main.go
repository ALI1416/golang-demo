package main

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func main() {
	// 创建连接
	dsn := "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	conn := sqlx.NewMysql(dsn)
	// 插入
	r, err := conn.ExecCtx(context.Background(), "insert into user (type, name) values (?, ?)", 1, "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RowsAffected())
	// 查询
	var u User
	query1 := "select id, name, type, create_at, update_at from user where id = ?"
	err1 := conn.QueryRowCtx(context.Background(), &u, query1, 1)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println(u)
	// 修改
	query2 := "update user set type = ? where name = ?"
	r2, err2 := conn.ExecCtx(context.Background(), query2, 2, "test")
	if err2 != nil {
		panic(err2)
	}
	fmt.Println(r2.RowsAffected())
	// 删除
	query3 := "delete from user where id = ?"
	r3, err3 := conn.ExecCtx(context.Background(), query3, 2)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println(r3.RowsAffected())
}
