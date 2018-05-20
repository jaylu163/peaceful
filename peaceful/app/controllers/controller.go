package  Controllers

import(
    "net/http"
    "fmt"
    "time"
)
const DIR="app/controllers"

var t time.Time


func Index(w http.ResponseWriter, r *http.Request){

    t=time.Now()
    var timeNow string
    timeNow =t.Format("2006-01-02 15:04:05")
    w.Write([]byte(timeNow))
    fmt.Println("aaaaaaaaaaa")

}


func Welcome(w http.ResponseWriter, r *http.Request){

    w.Write([]byte("welcome to golang world.............."))

}

func GetList(w http.ResponseWriter,r *http.Request){

    fmt.Fprintf(w,"getList action ##current time: %q",time.Now)
    fmt.Println("aaaaaaaaaaaaaaaa")
}
