package service

import (
	m "github.com/fedorkolmykow/avitoauto/pkg/models"
)

var runes = [...]rune{'0','1','2','3','4','5','6','7','8','9',
	'A','B','C','D','E','F','G','H','I','J','K','L','M',
	'N', 'O','P','Q','R','S','T','U','V','W','X','Y','Z',
	'a','b','c','d','e','f','g','h','i','j','k','l','m',
	'n', 'o','p','q','r','s','t','u','v','w','x','y','z',
	}

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
		remain := id % len(runes)
		keysRunes = append(keysRunes, runes[remain])
		id = id / len(runes)
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