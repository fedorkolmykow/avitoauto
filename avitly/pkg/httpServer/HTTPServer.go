package httpServer

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	m "github.com/fedorkolmykow/avitoauto/pkg/models"
)

type service interface {
	SaveURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error)
	GetURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error)
}

type server struct {
	svc service
}

func (s *server) HandleSaveURL(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warn(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req := &m.SaveURLReq{}
	err = req.UnmarshalJSON(body)
	if err != nil{
		log.Warn(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Trace("Received data: " + fmt.Sprintf("%+v", req))
	resp, err := s.svc.SaveURL(req)
	if err != nil{
		log.Warn(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	body, err = resp.MarshalJSON()
	if err != nil{
		log.Warn(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Warn(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (s *server) HandleRedirect(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["key"]
	req := &m.RedirectReq{}
	req.Key = key
	log.Trace("Received data: " + fmt.Sprintf("%+v", req))
	resp, err := s.svc.GetURL(req)
	if err != nil {
		log.Warn(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, resp.OriginalURL, http.StatusSeeOther)
}

func NewHTTPServer(svc service) (httpServer *mux.Router) {
	router := mux.NewRouter()
    s := &server{svc: svc}
	router.HandleFunc("/{key}", s.HandleRedirect).
		Methods("GET")
	router.HandleFunc("/url", s.HandleSaveURL).
		Methods("POST")
	return router
}