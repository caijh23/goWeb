package service

import (
	"net/http"
	"time"
    "github.com/unrolled/render"
)

func jsHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
		layout := "2006-01-02 15.04.05 PM"
        formatter.JSON(w, http.StatusOK, struct {
            Now      string `json:"now"`
        }{Now: time.Now().Format(layout)})
    }
}