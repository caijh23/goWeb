package service

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/caijh23/goWeb/web/negroni-gbk/gbk"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	//n.Use(gbk.Gbk())
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}

	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.HandleFunc("/js", jsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/login",loginHandler(formatter)).Methods("GET")
	mx.HandleFunc("/user",userHandler(formatter)).Methods("POST")
	mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	mx.NotFoundHandler = NotImplementedHandler()

}

func NotImplementedHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("501 page not implemented")
		http.Error(rw, "501 page not implemented", 501)
	})
}