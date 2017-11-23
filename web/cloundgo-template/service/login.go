package service

import (
	"fmt"
    "net/http"
    "strings"

    "github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        formatter.HTML(w, http.StatusOK, "login", nil)
    }
}
func userHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        fmt.Println("method: ",req.Method)
		formatter.HTML(w,http.StatusOK, "user", struct {
            Username    string `json:"username"`
            Password    string `json:"password"`
        }{Username: strings.Join(req.Form["username"],""), Password: strings.Join(req.Form["password"],"")})
	}
}