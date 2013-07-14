package demo

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

type RequestInfo struct {
	Form   url.Values
	Method string
	Path   string
	Scheme string
}

var serverPort int

func init() {
	flag.IntVar(&serverPort, "port", 9292, "the server (http listen) port")
}

func hiJack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi, Jack!")
	hj, ok := w.(http.Hijacker)
	if !ok {
		errorMsg := "The Web Server does not support Hijacking! "
		http.Error(w, errorMsg, http.StatusInternalServerError)
		fmt.Println(errorMsg)
		return
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		errorMsg := "Internal error!"
		http.Error(w, errorMsg, http.StatusInternalServerError)
		fmt.Printf(errorMsg+" Hijacking Error: %s\n", err)
		return
	}
	defer conn.Close()
	r.ParseForm()
	requestInfo := RequestInfo{Form: r.Form, Method: r.Method, Path: r.URL.Path, Scheme: r.URL.Scheme}
	b, err := json.Marshal(requestInfo)
	if err != nil {
		fmt.Println("JsonMarshalError:", err)
	} else {
		fmt.Printf("Request Detail: %s\n", string(b))
	}
	var sleepSeconds time.Duration
	paramSs := r.FormValue("ss")
	if m, _ := regexp.MatchString("^[0-9]+$", paramSs); !m {
		sleepSeconds = 1
	} else {
		i, err := strconv.ParseInt(paramSs, 10, 64)
		if err != nil {
			fmt.Printf("ParseIntError: %s\n", err)
			sleepSeconds = 1
		} else {
			sleepSeconds = time.Duration(i)
		}
	}
	time.Sleep(sleepSeconds * time.Second)
	resp := fmt.Sprintf("Hi~! How are you? I slept for %d seconds.\n", sleepSeconds)
	_, err = bufrw.Write([]byte(resp))
	if err == nil {
		err = bufrw.Flush()
	}
	if err != nil {
		fmt.Printf("ResponseError: %s\n", err)
	} else {
		fmt.Println("Bye, Jack!")
	}
}

func DemoHijack() {
	flag.Parse()
	http.HandleFunc("/hi-jack", hiJack)
	fmt.Printf("Starting http server with hijack (port=%d)...\n", serverPort)
	err := http.ListenAndServe(":"+fmt.Sprintf("%d", serverPort), nil)
	if err != nil {
		fmt.Println("ListenAndServeError: ", err)
	} else {
		fmt.Println("The http server is started.")
	}
}
