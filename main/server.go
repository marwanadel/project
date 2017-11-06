package main

import (

    //"fmt"
    "net/http"
    "golazo/routes"

)

func main() {

    routes.Route()
    http.ListenAndServe(":8080", nil)

}
