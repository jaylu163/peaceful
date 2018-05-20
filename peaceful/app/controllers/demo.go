package Controllers

import(
    "net/http"
    "encoding/json"
)

type Todo struct{

    Id int `json:"id"`
    Username string `json:"username"`
    Type int `json:"type"`

}

type Todos []Todo

var todos Todos =Todos{
    Todo{
        1,
        "huaan",
        1,
    },
    Todo{
        2,
        "tom",
        0,
    },
    {
        3,
        "pony",
        2,
    },

}


func UserList(w http.ResponseWriter ,r *http.Request){

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    err :=json.NewEncoder(w).Encode(todos)
    if err !=nil{
        panic(err)

    }

}




