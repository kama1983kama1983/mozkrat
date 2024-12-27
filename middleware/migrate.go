package middleware

import (
	"io/ioutil"
	"log"
)

func Migrate() {
	db := InitDb()
	defer db.Close()
	read, err := ioutil.ReadFile("./database.sql")
	if err != nil {
		log.Fatalf("لا يمكن قراءة الملف %v", err)
	}
	_, err = db.Exec(string(read))
	if err != nil {
		log.Fatalf("هناك خطأ فى اتصال ببيانات %v", err)
	}
	log.Fatal("تم ادخال البيانات بنجاح !")
}
