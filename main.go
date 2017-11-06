package main

import (

    //"fmt"
    "net/http"
    "routes"

)

func main() {

    routes.Route()
    http.ListenAndServe(":8080", nil)

}
