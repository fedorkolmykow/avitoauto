package postgres

import (
	m "github.com/fedorkolmykow/avitoauto/pkg/models"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const(
	SelectURLOnKey = `SELECT * FROM URLs WHERE key=$1;`
	SelectURL = `SELECT * FROM URLs WHERE url=$1;`
	InsertURL = `INSERT INTO URLs (url, key) VALUES ($1, $2) RETURNING url_id;`
	UpdateURL = `UPDATE URLs SET key=$1 WHERE url_id=$2;`
	SelectExist = `SELECT EXISTS(SELECT url_id FROM URLs WHERE url=$1) ;`
	SelectExistCustomKey = `SELECT EXISTS(SELECT url_id FROM URLs WHERE key=$1) ;`
)

type DbClient interface{
	InsertURL(url *m.URL) (id int, err error)
	UpdateKey(urlId int, key string) (err error)
	SelectURLOnKey(key string) (url *m.URL, err error)
	SelectURL(origUrl string) (url *m.URL, err error)
	Exist(origUrl string) (exist bool, err error)
	ExistCustomKey(key string) (exist bool, err error)
	Shutdown() error
}

type dbClient struct{
    db *sqlx.DB
}

func (d *dbClient) InsertURL(url *m.URL) (id int, err error) {
	err = d.db.QueryRowx(InsertURL, url.URL, url.Key).Scan(&id)
	return
}

func (d *dbClient) UpdateKey(urlId int, key string) (err error){
	_, err = d.db.Exec(UpdateURL, key, urlId)
	return
}

func (d *dbClient) SelectURLOnKey(key string) (url *m.URL, err error){
	url = &m.URL{}
	err = d.db.Get(url,SelectURLOnKey,key)
	return
}

func (d *dbClient) SelectURL(origUrl string) (url *m.URL, err error){
	url = &m.URL{}
	err = d.db.Get(url,SelectURL,origUrl)
	return
}

func (d *dbClient) Exist(origUrl string) (exist bool, err error){
	err = d.db.QueryRowx(SelectExist, origUrl).Scan(&exist)
	return
}

func (d *dbClient) ExistCustomKey(key string) (exist bool, err error){
	err = d.db.QueryRowx(SelectExistCustomKey, key).Scan(&exist)
	return
}

func (d *dbClient) Shutdown() error{
	return d.db.Close()
}

func NewDbClient() DbClient{
	db, err := sqlx.Connect("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	//db.SetMaxIdleConns(n int)
	//db.SetMaxOpenConns(n int)
	return &dbClient{db: db}
}