package etherdb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

// create table locking_data (
//  id serial not null,
// 	lookup_id int,
// 	operation varchar(20),
// 	blocknumber int,
// 	index int,
// 	owner varchar(44) not null,
//  txhash varchar(70) not null,
// 	"timestamp" integer,
// 	amount varchar(44)
// );

// alter table locking_data add column lookup_id int;

// insert into lookup (type,created,description,enabled,address) values (5,15713626,'locking',false,'0x756fe78b65A400F07b6fcA2F92E0482f6DcBF25B');

type LockingData struct {
	ID          uint64
	LookupID    int
	Operation   string
	BlockNumber uint64
	TxIndex     uint
	TxHash      string
	Owner       string
	Amount      string
	Timestamp   uint64
}

func NewLockedTokenData(
	LookupID int,
	operation string,
	blocknumber uint64,
	txIndex uint,
	txHash string,
	owner string,
	amount string,
	timestamp uint64) (mt LockingData) {
	return LockingData{
		LookupID:    LookupID,
		Operation:   operation,
		BlockNumber: blocknumber,
		TxIndex:     txIndex,
		TxHash:      txHash,
		Owner:       owner,
		Amount:      amount,
		Timestamp:   timestamp,
	}
}

var lock_db *sql.DB

func InitLockingDB(connectString string) {
	var err error
	lock_db, err = sql.Open("postgres", connectString) //
	if err != nil {
		fmt.Println("open token database : ", err)
		log.Fatal("open token database : ", err)
	}
	fmt.Println("database is open")
}

func (tt *LockingData) Add() (err error) {

	//                1       2        3        4         5      6     7     8
	statement := `insert into locking_data 
				(lookup_id,operation,blocknumber,index,txhash,owner,amount,timestamp) 
				values 
				($1,$2,$3,$4,$5,$6,$7,$8) 
				returning id`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	//                         1       2              3             4           5          6          7           8        9         10           11           12
	err = lock_db.QueryRowContext(ctx, statement,
		tt.LookupID, tt.Operation, tt.BlockNumber, tt.TxIndex, tt.TxHash, tt.Owner, tt.Amount, tt.Timestamp).Scan(&tt.ID)
	return
}
