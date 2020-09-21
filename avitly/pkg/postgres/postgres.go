package postgres

import (
	"os"

	m "github.com/fedorkolmykow/avitoauto/pkg/models"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const(
	CheckExistence = `SELECT EXISTS(SELECT user_id FROM Users WHERE user_id=$1) ;`
	SelectUserBalance = `SELECT balance FROM Users WHERE user_id=$1;`
	InsertUser = `INSERT INTO Users (user_id, balance) VALUES ($1, $2) RETURNING user_id;`
	UpdateUserBalance = `UPDATE Users SET balance = balance + $1 WHERE user_id = $2 RETURNING balance;`
	SetIsolationSerializable = `SET TRANSACTION ISOLATION LEVEL SERIALIZABLE;`
	InsertTrans = `INSERT INTO Transactions (user_id, init_balance, change, time, comment, source)  
                     VALUES (:user_id, :init_balance, :change, :time, :comment, :source);`
	SelectTransactions = `SELECT * FROM Transactions WHERE user_id=$1;`
)

type DbClient interface{
	InsertURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error)
	SelectURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error)
	Shutdown() error
}

type dbClient struct{
    db *sqlx.DB
}


func rollAndErr(tx *sqlx.Tx, err error) error{
	log.Trace("Rollback")
	errRoll := tx.Rollback()
	if errRoll != nil{
		return errRoll
	}
	return err
}

func (d *dbClient) InsertURL(Req *m.SaveURLReq) (Resp *m.SaveURLResp, err error){
	return
}
func (d *dbClient) SelectURL(Req *m.RedirectReq) (Resp *m.RedirectResp, err error){
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