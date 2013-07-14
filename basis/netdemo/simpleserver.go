/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-1-11，下午4:30
 * @version 0.1
 */
package netdemo

import (
	"net/http"
)

func SimpleServer() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("hello"))
		}))
}
