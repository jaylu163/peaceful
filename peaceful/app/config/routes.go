package Config

import(
    "net/http"
    "github.com/gorilla/mux"
    "peaceful/app/Logger"
)


func NewRouter() *mux.Router{

    router := mux.NewRouter().StrictSlash(true)
    for _,route :=range routes{
        var handler http.Handler

        handler =route.HandlerFunc
        handler =Logger.WriteLogger(handler,route.Name)
        router.Methods(route.HttpMethod).Path(route.Pattern).Name(route.Name).Handler(handler)
    }
    return router

}
