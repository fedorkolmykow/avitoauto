package service

import (
	"errors"
	m "github.com/fedorkolmykow/avitoauto/pkg/models"
	log "github.com/sirupsen/logrus"
)

const(
	errorCustomKeyExist = "custom key already exist"
)

type Service interface {
	SaveURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error)
	GetURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error)
}

type dbClient interface{
	InsertURL(url *m.URL) (id int, err error)
	UpdateKey(urlId int, key string) (err error)
	SelectURLOnKey(key string) (url *m.URL, err error)
	SelectURL(origUrl string) (url *m.URL, err error)
	Exist(origUrl string) (exist bool, err error)
	ExistCustomKey(key string) (exist bool, err error)
}

type service struct{
	db dbClient
}

func (s *service) SaveURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error){
	var exist bool
	err = Req.Validate()
	if err != nil{
		return
	}
	var url = &m.URL{
		URL:   Req.OriginalURL,
		Key:   Req.CustomKey,
	}
	if Req.CustomKey != ""{
		exist, err = s.db.ExistCustomKey(Req.CustomKey)
		if err != nil {
			return
		}
		if exist{
			err = errors.New(errorCustomKeyExist)
			return
		}
	}
	exist, err = s.db.Exist(url.URL)
	if !exist {
		url.Id, err = s.db.InsertURL(url)
		if err != nil {
			return
		}
	} else {
		url, err = s.db.SelectURL(Req.OriginalURL)
		if err != nil {
			return
		}
	}
	if Req.CustomKey == ""{
		url.Key, err = shorten(url.Id)
	}
	err = s.db.UpdateKey(url.Id, url.Key)
	if err != nil{
		return
	}
	Resp = &m.SaveURLResp{
		KeyID: url.Id,
		Key:   url.Key,
	}
	return
}

func (s *service) GetURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error){
	err = Req.Validate()
	if err != nil{
		return
	}
	url, err := s.db.SelectURLOnKey(Req.Key)
	log.Trace(url)
	if err != nil{
		return
	}
	Resp = &m.RedirectResp{OriginalURL: url.URL}
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