/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-3-18，下午2:53
 * @version 0.1
 */
package dbdemo

import (
	"github.com/ziutek/mymysql/mysql"
	"os"
	"fmt"
)


func DemoMysql(){

	user :="root"
	pass :="admin"
	dbname :="test"

	db := mysql.New("tcp", "", "127.0.0.1:3306", user, pass, dbname)

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	rows, res, err := db.Query("select * from t_user where id > %d", 20)
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		for _, col := range row {
			if col == nil {
				// col has NULL value
			} else {
				// Do something with text in col (type []byte)
			}
		}
		// You can get specific value from a row
		val1 := row[1].([]byte)

		// You can use it directly if conversion isn't needed
		os.Stdout.Write(val1)

		// You can get converted value
		number := row.Int(0)      // Zero value
		str    := row.Str(1)      // First value
		bignum := row.Uint(2) // Second value

		fmt.Println(number)
		fmt.Println(str)
		fmt.Println(bignum)
		//// You may get values by column name
		first := res.Map("FirstColumn")
//		second := res.Map("SecondColumn")
		fmt.Println(first)

//		val1, val2 := row.Int(first), row.Str(second)
	}
}
