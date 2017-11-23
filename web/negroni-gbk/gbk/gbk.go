package gbk

import (
	"net/http"
	"strings"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding/simplifiedchinese"
	"github.com/urfave/negroni"
)
const (
	headerContentType = "Content-type"
	encodingGbk = "gbk"
	encodingUTF8 = "UTF-8"
)
type gbkResponseWriter struct {
	w *transform.Writer
	negroni.ResponseWriter
	wroteHeader bool
}
func (grw *gbkResponseWriter) WriteHeader(code int) {
	headers := grw.ResponseWriter.Header()
	if headers.Get(headerContentType) == "" {
		headers.Set(headerContentType,encodingGbk)
	} else if strings.Contains(headers.Get(headerContentType),encodingUTF8) {
		headers.Set(headerContentType,strings.Replace(headers.Get(headerContentType),encodingUTF8,encodingGbk,-1))
	} else {
		grw.w = nil
	}
	grw.ResponseWriter.WriteHeader(code)
	grw.wroteHeader = true
}

func (grw *gbkResponseWriter) Write(b []byte) (int, error) {
	if !grw.wroteHeader {
		grw.WriteHeader(http.StatusOK)
	}
	if grw.w == nil {
		return grw.ResponseWriter.Write(b)
	}
	if len(grw.Header().Get(headerContentType)) == 0 {
		grw.Header().Set(headerContentType,http.DetectContentType(b))
	}
	return grw.w.Write(b)
}

type handler struct {

}

func Gbk() *handler {
	h := &handler{}
	return h
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if len(r.Header.Get(headerContentType)) == 0 {
		next(w, r)
		return
	}
	var nrw negroni.ResponseWriter
	var gb *transform.Writer
	if strings.Contains(r.Header.Get(headerContentType), encodingUTF8) {
		nrw = negroni.NewResponseWriter(w)
		gb = transform.NewWriter(nrw, nil)
	}
	if strings.Contains(r.Header.Get(headerContentType), encodingGbk) {
		rd := transform.NewReader(r.Body, simplifiedchinese.GBK.NewDecoder())
		r.Body = ioutil.NopCloser(rd)
		nrw = negroni.NewResponseWriter(w)
		gb = transform.NewWriter(nrw, simplifiedchinese.GBK.NewEncoder())
	}
	grw := gbkResponseWriter{gb, nrw, false}

	next(&grw, r)
}