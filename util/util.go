package util

import "log"

//func CheckErr(err error, msg... string) error {
//	if err != nil {
//		log.Fatal("err occurs", err)
//	}
//	return err
//}

func CheckErr(msg string, err error) error {
	if err != nil {
		log.Fatal(msg , err)
	}
	return err
}

func CheckOK(msg string, ok bool) bool {
	if ok == false {
		log.Println(msg , ok)
	}
	return ok
}
