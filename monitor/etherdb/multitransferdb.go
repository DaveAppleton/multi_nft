package etherdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // DB selection
)

/* MultiTokenTransfers
create table $1 (
	id serial not null primary key,
	tokenid int not null,
	blocknumber int,
	blockhash varchar(70)  not null,
	index int,
	txhash varchar(70) not null,
    source varchar(44) not null,
    dest varchar(44) not null,
    "timestamp" integer,
    Amount int not null
);
*/

type MultiTokenTransfer struct {
	ID           uint64
	LookupRef    int
	TokenID      uint64
	TokenAddress string
	BlockNumber  uint64
	TxIndex      uint
	TxHash       string
	Operator     string
	Source       string
	Dest         string
	Quantity     uint64
	Timestamp    uint64
}

func NewMultiToken(
	lookupRef int,
	tokenid uint64,
	blocknumber uint64,
	txIndex uint,
	txHash string,
	operator string,
	source string,
	dest string,
	quantity uint64,
	timestamp uint64) (mt MultiTokenTransfer) {
	mt = MultiTokenTransfer{
		LookupRef:   lookupRef,
		TokenID:     tokenid,
		BlockNumber: blocknumber,
		TxIndex:     txIndex,
		TxHash:      txHash,
		Operator:    operator,
		Source:      source,
		Dest:        dest,
		Quantity:    quantity,
		Timestamp:   timestamp,
	}
	return
}

// Add this token to the database
func (tt *MultiTokenTransfer) Add() (err error) {
	query := `insert into multi_tokens (tokenid,lookup_id,blocknumber,index,txhash,operator,source,dest,amount,timestamp) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) returning id`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.TokenID, tt.LookupRef, tt.BlockNumber, tt.TxIndex, tt.TxHash, &tt.Operator, tt.Source, tt.Dest, tt.Quantity, tt.Timestamp).Scan(&tt.ID)
	return
}

// AddIfNotFound only adds if the txhash is not found
func (tt *MultiTokenTransfer) AddIfNotFound() (err error) {
	query := `insert into multi_tokens (txhash,lookup_id,tokenid,blocknumber,index,operator,source,dest,amount,timestamp) 
					values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) on conflict do nothing`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	_, err = db.ExecContext(ctx, query, tt.TxHash, tt.LookupRef, tt.TokenID, tt.BlockNumber, tt.TxIndex, &tt.Operator, tt.Source, tt.Dest, tt.Quantity, tt.Timestamp)
	return
}

func (*MultiTokenTransfer) getTransfers(rows *sql.Rows) (transfers []MultiTokenTransfer, err error) {
	var t MultiTokenTransfer
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.TokenID, &t.BlockNumber, &t.TxIndex, &t.TxHash, &t.Operator, &t.Source, &t.Dest, &t.Quantity, &t.Timestamp)
		if err != nil {
			return
		}
		transfers = append(transfers, t)
	}
	return
}

func (*MultiTokenTransfer) getUniqueTransfers(rows *sql.Rows) (transfers []MultiTokenTransfer, err error) {
	var t MultiTokenTransfer
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.TokenID, &t.BlockNumber, &t.TxIndex, &t.TxHash, &t.Operator, &t.Source, &t.Dest, &t.Timestamp)
		if err != nil {
			return
		}
		transfers = append(transfers, t)
	}
	return
}

// Find transfers that match tokenID
func (tt *MultiTokenTransfer) Find() (transfers []MultiTokenTransfer, err error) {
	query := `select id,tokenid,blocknumber,index,txhash,operator,source,dest,quantity,timestamp from multi_tokens where tokenid=$1`

	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.TokenID)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)

}

// FindByAddress returns transfers of specific token to or from an address
func (tt *MultiTokenTransfer) FindByAddress(addr string) (transfers []MultiTokenTransfer, err error) {
	query := `select id,tokenid,blocknumber,index,txhash,operator,source,dest,quantity,timestamp from multi_tokens  where  (source=$1 or dest=$1)`

	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, addr)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)
}

// FindAllByAddress returns transfers of any token to or from an address newest first
func (tt *MultiTokenTransfer) FindAllByAddress(addr string) (transfers []MultiTokenTransfer, err error) {
	query := `select id,tokenid,blocknumber,index,txhash,operator,source,dest,quantity,timestamp from multi_tokens
		 where (source=$1 or dest=$1)
		 order by blocknumber desc, index desc`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, addr)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)
}

// MaxBlock for a specified token
func (tt *MultiTokenTransfer) MaxBlock() (max uint64, err error) {
	query := `select max(blocknumber) from multi_tokens`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	err = db.QueryRowContext(ctx, query, tt.TokenID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}
