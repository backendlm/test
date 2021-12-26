package main

import (
	"bytes"
	"database/sql"
	"fmt"
	 "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
)

var db, err = sql.Open("mysql", "root:Ll20030916@/user")
var buffer bytes.Buffer
type user struct {
	id       int
	age      int
	name     string
	username string
	password string
	question string
	answer   string
	email    string
}
type comment struct {
	id       int
	reid     int
	sender   string
	receiver string
	message string
}
var message = make([]string,0)
func main() {
	if err != nil {
		log.Fatal(err)
	}
	r := Default()
	r.GET("/board",board)
	r.POST("/login", login)
	r.POST("/regist", regist)
	r.POST("/find", find)
	r.POST("/message",me)
	check1 := func(c *Context) {

		value, err1 := c.Cookie("login_cookie")
		if err1 != nil {
			c.String(403, "认证失败，没有cookie")
			c.Abort()

		}
		c.Set("cookie", value)
	}
	r.GET("/hello", check1, func(c *Context) {
		cookie, _ := c.Get("cookie")
		s := cookie.(string)
		c.String(200, "登录成功,你好"+s)

	})
	r.Run()
}

func login(c *Context) {
	_, err := c.Cookie("login_cookie")
	if err == nil {
		c.String(403, "您已登录，请勿重复登录。")
		c.Abort()
		return
	}
	uname := c.PostForm("username")
	pword := c.PostForm("password")
	che := "select username,password from users where id > ?"
	stmt, err := db.Prepare(che)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.password)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if uname == u.username && pword == u.password && uname != "" {
			c.SetCookie("login_cookie", uname, 1800, "/", "", false, true)
			c.String(http.StatusOK, "登录成功")
			return
		}
	}
	c.String(403, "用户名或密码错误，请重新登录。")

	return


}

func regist(c *Context) {
	_, err := c.Cookie("login_cookie")
	if err == nil {
		c.String(403, "您已登录，无法进行注册。")
		return
	}
	name := c.PostForm("name")
	age := c.PostForm("age")
	uname := c.PostForm("username")
	pword := c.PostForm("password")
	question := c.PostForm("question")
	answer := c.PostForm("answer")
	email := c.PostForm("email")

	if len(uname) < 6 {
		c.String(403, "用户名不得小于6位，请重新注册。")
		c.Abort()
		return
	}
	if question == "" || answer == "" {
		c.String(403, "请设置密保")
		c.Abort()
		return
	}
	//在这里限定了各项参数，防止和数据库中的参数限制冲突。
	if uname == "" || pword == "" || name == "" || age == "" || email == "" {
		c.String(403, "value不能为空，请重新注册。")
		c.Abort()
		return
	}
	if check(db, uname, email) {
		c.String(403, "出现错误，用户名或邮箱已被注册，请重新注册。")
		c.Abort()
		return
	}
	sqlStr := "insert into users(name,age,username,password,question,answer,email) values (?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, age, uname, pword, question, answer, email)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}

	c.String(403, "注册成功，请重新登录。")
	return
}

func check(db *sql.DB, reguser, email string) bool {
	check := "select username,email from users where id > ?"
	c, err := db.Prepare(check)
	if err != nil {
		fmt.Println("prepare failed")
		return true
	}
	defer c.Close()
	rows, err := c.Query(0)
	if err != nil {
		fmt.Println("query failed")
		return true
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.email)
		if err != nil {
			fmt.Println("scan filed")
			return true
		}
		if reguser == u.username || email == u.email {
			return true
		}
	}
	return false
}

func find(c *Context) {
	q := c.PostForm("question")
	a := c.PostForm("answer")
	e := c.PostForm("email")
	if q == "" || a == "" || e == "" {
		c.String(403, "请输入密保,注册邮箱,找回账号")
	}
	sqlStr := "select username,question,answer,email from users where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.username, &u.question, &u.answer, &u.email)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		if u.question == q && u.answer == a {
			c.String(200, "密保与邮箱输入正确，您的账号是"+u.username)
			return
		}
	}
}