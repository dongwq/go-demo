/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-18，上午1:00
 * @version 0.1
 */
package main

import "fmt"
import "database/sql"
//import _ "github.com/Go-SQL-Driver/MySQL"

func main() {
	fmt.Printf("query:\n")
	db, e := sql.Open("mysql", "root:admin@/longhua?charset=utf8")
	//说明：第一个参数"mysql"表示打开一个MYSQL链接
	//第二个参数是一个DSN，格式是：
	//[username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]

	if e != nil {
		fmt.Printf("conn error.");
		return;
	}
	rows, e := db.Query("select * from tags")
	if e != nil {
		fmt.Printf("query error.");
		return;
	}

	i := 0
	for rows.Next() {
		i++
		var ss, s2 string
		e = rows.Scan(&ss, &s2)
		if e == nil {
			fmt.Printf("Title: %s,%s \n", ss, s2)
		}else {
			fmt.Printf("error")
		}
	}
	fmt.Printf("Total: %d\n" , i)
	db.Close()

}
