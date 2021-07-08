package etherdb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // DB selection
)

// TokenTransfer = the data for a token transfer :-)
type UniqueTokenTransfer struct {
	ID          uint64
	LookupID    int
	BlockNumber uint64
	Timestamp   uint64
	TxIndex     uint
	TxHash      string
	Source      string
	Dest        string
	TokenID     uint64
}

func NewUniqueToken(
	LookupID int,
	blocknumber uint64,
	txindex uint,
	txhash string,
	source string,
	dest string,
	tokenid uint64,
	timestamp uint64) (mt UniqueTokenTransfer) {
	mt = UniqueTokenTransfer{
		LookupID:    LookupID,
		BlockNumber: blocknumber,
		TxIndex:     txindex,
		TxHash:      txhash,

		TokenID:   tokenid,
		Source:    source,
		Dest:      dest,
		Timestamp: timestamp,
	}
	return
}

func (tt *UniqueTokenTransfer) updateUniqueOwners() (err error) {
	query := `insert into unique_owners (tokenid,lookup_id,owner) 
	values ($1,$2,$3) 
	on conflict on constraint unq
	do update set 
	owner=$3`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	_, err = db.ExecContext(ctx, query, tt.TokenID, tt.LookupID, tt.Dest)
	return
}

// Add this token to the database
func (tt *UniqueTokenTransfer) Add() (err error) {
	query := `insert into unique_tokens (lookup_id,tokenid,blocknumber,index,txhash,source,dest,timestamp) values ($1,$2,$3,$4,$5,$6,$7,$8) returning id`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TokenID, tt.BlockNumber, tt.TxIndex, tt.TxHash, tt.Source, tt.Dest, tt.Timestamp).Scan(&tt.ID)
	if err != nil {
		return
	}
	return tt.updateUniqueOwners()
}

// AddIfNotFound only adds if the txhash is not found
func (tt *UniqueTokenTransfer) AddIfNotFound() (err error) {
	query := `insert into unique_tokens (lookup_id,txhash,tokenid,blocknumber,index,source,dest,timestamp) 
					values ($1,$2,$3,$4,$5,$6,$7,$8) on conflict do nothing`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	_, err = db.ExecContext(ctx, query, tt.LookupID, tt.TxHash, tt.TokenID, tt.BlockNumber, tt.TxIndex, tt.Source, tt.Dest, tt.Timestamp)
	return
}

func (*UniqueTokenTransfer) getTransfers(rows *sql.Rows) (transfers []UniqueTokenTransfer, err error) {
	var t UniqueTokenTransfer
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.LookupID, &t.TokenID, &t.BlockNumber, &t.TxIndex, &t.TxHash, &t.Source, &t.Dest, &t.Timestamp)
		if err != nil {
			return
		}
		transfers = append(transfers, t)
	}
	return
}

func (*UniqueTokenTransfer) getUniqueTransfers(rows *sql.Rows) (transfers []UniqueTokenTransfer, err error) {
	var t UniqueTokenTransfer
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.LookupID, &t.TokenID, &t.BlockNumber, &t.TxIndex, &t.TxHash, &t.Source, &t.Dest, &t.Timestamp)
		if err != nil {
			return
		}
		transfers = append(transfers, t)
	}
	return
}

// Find transfers that match tokenID
func (tt *UniqueTokenTransfer) ListAll(start int) (transfers []UniqueTokenTransfer, err error) {
	query := `select id,lookup_id,tokenid,blocknumber,index,txhash,source,dest,timestamp from unique_tokens where lookup_id=$1 offset $2 limit $3`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	rows, err := db.QueryContext(ctx, query, tt.LookupID, start, 50)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)

}

// FindByAddress returns transfers of specific token to or from an address
func (tt *UniqueTokenTransfer) FindByAddress(addr string) (transfers []UniqueTokenTransfer, err error) {
	query := `select id,lookup_id,tokenid,blocknumber,index,txhash,source,dest,timestamp from unique_tokens  where lookup_id = $1 and tokenid=$2 and (source=$3 or dest=$3)`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TokenID, addr)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)
}

// FindByAddress returns transfers of specific token to or from an address
func (tt *UniqueTokenTransfer) FindByTokenId() (transfers []UniqueTokenTransfer, err error) {
	query := `select id,lookup_id,tokenid,blocknumber,index,txhash,source,dest,timestamp from unique_tokens  where lookup_id=$1 and tokenid=$2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TokenID)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)
}

func (tt *UniqueTokenTransfer) FindByTxHash() (transfers []UniqueTokenTransfer, err error) {
	query := `select id,lookup_id,tokenid,blocknumber,index,txhash,source,dest,timestamp from unique_tokens  where lookup_id=$1 and txhash=$2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TxHash)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)
}

// FindAllByAddress returns transfers of any token to or from an address newest first
func (tt *UniqueTokenTransfer) FindAllByAddress(addr string) (transfers []UniqueTokenTransfer, err error) {
	query := `select id,lookup_id,tokenid,blocknumber,index,txhash,source,dest,timestamp from unique_tokens
		 where (source=$1 or dest=$1) and lookup_id = $2
		 order by blocknumber desc, index desc`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	fmt.Println(addr, tt.LookupID)
	rows, err := db.QueryContext(ctx, query, addr, tt.LookupID)
	if err != nil {
		return
	}
	return tt.getTransfers(rows)
}

// MaxBlock for a specified token
func (tt *UniqueTokenTransfer) MaxBlock() (max int64, err error) {
	query := `select max(blocknumber) from unique_tokens`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.TokenID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

// MaxBlock for a specified token
func (tt *UniqueTokenTransfer) SearchCountOfThisToken(addr string) (max int, err error) {
	query := `select count(*) from unique_tokens where (source=$1 or dest=$1) and lookup_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, addr, tt.LookupID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *UniqueTokenTransfer) CountOfThisToken() (max int, err error) {
	query := `select count(*) from unique_tokens where lookup_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *UniqueTokenTransfer) CountOfThisTokenID() (max int, err error) {
	query := `select count(*) from unique_tokens where lookup_id = $1 and token_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TokenID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *UniqueTokenTransfer) CountOfThisTxHash() (max int, err error) {
	query := `select count(*) from unique_tokens where lookup_id = $1 and txhash = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TxHash).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}
