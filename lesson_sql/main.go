package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
func dbConn()(db *sql.DB,e error){
	dbDriver:="mysql"
	dbUser:="root"
	dbPass:="root"
	dbName:="douban"
	db,err:=sql.Open(dbDriver,fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",dbUser,dbPass,dbName))
	if err != nil {
		return nil,err
	}
	return db,nil
}

func dbQuery() (err error) {
	db,err:=dbConn()
	if err != nil {
		fmt.Printf("Conn failed! %s",err.Error())
		return err
	}
	defer db.Close()

	//result,err:=db.Exec("select * from posts")
	stmt,err:=db.Prepare("Select * from posts where id>?")
	rows,err:=stmt.Query(2)
	//rows,err:=db.Query("select * from posts")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	post:=Posts{}
	for rows.Next(){
		rows.Scan(&post.Id,&post.Title,&post.Body)
		fmt.Println(post)
	}
	return nil
}

func dbInsert() (err error) {
	db,err:=dbConn()
	if err != nil {
		fmt.Printf("Conn failed! %s",err.Error())
		return err
	}
	defer db.Close()
	stmt,err:=db.Prepare("INSERT INTO posts(title,body) values (?,?)")
	if err != nil {
		fmt.Printf("%s",err.Error())
		return err
	}
	title:="百年孤独"
	body:="一个家族的兴衰！"
	stmt.Exec(title,body)
	return nil
}
type Posts struct {
	Id int
	Title string
	Body string
}

func main() {
	//dbInsert()
	dbQuery()
}