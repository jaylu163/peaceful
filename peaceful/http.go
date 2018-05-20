package main

import(
    "net/http"
    "log"
    router "peaceful/app/Config"

)


func main(){

    routers:=router.NewRouter()

    err:=http.ListenAndServe("localhost:8899",routers)
    if(err !=nil){
        log.Fatal(err)
    }
}





