package Controllers

import(
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
    "time"
    "fmt"
)


type User struct{

    Id int `json:"id"`
    Name string `json:"name"`
    Completed bool `json:"completed"`
    Due time.Time `json:"due"`

}

var user User

func Adduser(w http.ResponseWriter, r *http.Request){

    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err !=nil{
        panic(nil)
    }
    
    if err :=json.Unmarshal(body,&user); err!=nil{

        w.Header().Set("Content-Type","application/json;charset=UTF-8")
        w.WriteHeader(422)
        if err:=json.NewEncoder(w).Encode(err); err !=nil{
            panic(err)
        }
    }

    fmt.Println(string(body))

}
