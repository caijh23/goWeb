package service

import (
    "net/http"

    "github.com/urfave/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/hello/{id}", testHandler(formatter)).Methods("GET")
}

func testHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        vars := mux.Vars(req)
        id := vars["id"]
        formatter.JSON(w, http.StatusOK, struct{ Test string }{"Hello " + id})
    }
}
// package main

// import (
//   "github.com/urfave/negroni"
//   "net/http"
//   "io"
//   "fmt"
// )
// func main() {
// 	n := negroni.New()
// 	n.UseFunc(printAuthorInfo)
// 	router:=http.NewServeMux()
// 	router.Handle("/",handler())
// 	n.UseHandler(router)
// 	n.Run(":1234")
// }
// func printAuthorInfo(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
// 	fmt.Println("123")
// 	//next(rw,r)
// }
// func handler() http.Handler{
// 	return http.HandlerFunc(myHandler)
// }
// func myHandler(rw http.ResponseWriter, r *http.Request) {
// 	rw.Header().Set("Content-Type", "text/plain")
// 	io.WriteString(rw,"Hello World")
// }