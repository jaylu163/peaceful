package Config

import(
    "net/http"
    contr "peaceful/app/Controllers"
)

type Route struct{
     Name string
     HttpMethod string
     Pattern string
     HandlerFunc http.HandlerFunc

}


type Routes []Route


var routes  =Routes{
    Route{
        "index",
        "GET",
        "/",
        contr.Index,
    },
    Route{
        "Index",
        "GET",
        "/index",
        contr.Index,
    },
    Route{
        "Welcome",
        "GET",
        "/welcome",
        contr.Welcome,
    },
    Route{
        "getlist",
        "GET",
        "/getlist",
        contr.GetList,
    },
    Route{
        "UserList",
        "GET",
        "/userlist",
        contr.UserList,
    },
    Route{
        "Adduser",
        "POST",
        "/add",
        contr.Adduser,
    },

}
