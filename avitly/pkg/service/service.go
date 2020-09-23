package service

import (
	m "github.com/fedorkolmykow/avitoauto/pkg/models"
)



type Service interface {
	SaveURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error)
	GetURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error)
}

type dbClient interface{
	InsertURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error)
	SelectURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error)
}

type service struct{
	db dbClient
}

func (s *service) SaveURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error){

	return
}

func (s *service) GetURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error){
	return
}

func shorten(id int) (key string, err error){
	var keysRunes []rune
	for id > 0{
		remain := id % len(m.KeyRunes)
		keysRunes = append(keysRunes, m.KeyRunes[remain])
		id = id / len(m.KeyRunes)
	}
	key = string(keysRunes)
	return
}

func NewService(db dbClient) Service{
    svc := &service{
    	db: db,
	}
    return svc
}