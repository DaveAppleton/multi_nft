package etherdb

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // DB selection
)

// Token holds ERC20 token data

//item := etherdb.UniqueSale{filt.Event.Raw.TxHash.Hex(), filt.Event.Buyer.Hex(), filt.Event.TokenId, etherUtils.EtherToStr(filt.Event.Price)}
type UniqueSale struct {
	ID          uint64
	LookupID    int
	Operation   string
	BlockNumber uint64
	TxIndex     uint
	TxHash      string
	Token       string
	TokenID     uint64
	Seller      string
	Buyer       string
	Quantity    uint64
	Price       string
	Hash        string
	Position    uint64
	Timestamp   uint64
}

func NewUniqueSale(
	LookupID int,
	operation string,
	tokenid uint64,
	blocknumber uint64,
	txIndex uint,
	txHash string,
	buyer string,
	seller string,
	price string,
	hash string,
	timestamp uint64) (mt UniqueSale) {
	mt = UniqueSale{
		LookupID:    LookupID,
		Operation:   operation,
		TokenID:     tokenid,
		BlockNumber: blocknumber,
		TxIndex:     txIndex,
		TxHash:      txHash,
		Seller:      seller,
		Buyer:       buyer,
		Price:       price,
		Hash:        hash,
		Timestamp:   timestamp,
	}
	return
}

func (tt *UniqueSale) Add() (err error) {
	//                1       2        3        4         5      6     7     8     9     10     11
	statement := `insert into unique_sale 
				(tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash,timestamp) 
				values 
				($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) 
				returning id`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	//                         1       2              3             4           5          6          7           8        9         10           11           12
	err = db.QueryRowContext(ctx, statement,
		tt.TokenID, tt.LookupID, tt.Operation, tt.BlockNumber, tt.TxIndex, tt.TxHash, tt.Buyer, tt.Seller, tt.Price, tt.Hash, tt.Timestamp).Scan(&tt.ID)
	return
}

// id,tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash,position, timestamp
func (*UniqueSale) getSales(rows *sql.Rows) (transfers []UniqueSale, err error) {
	var t UniqueSale
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.TokenID, &t.LookupID, &t.Operation, &t.BlockNumber, &t.TxIndex, &t.TxHash, &t.Buyer, &t.Seller, &t.Price, &t.Hash, &t.Timestamp)
		if err != nil {
			return
		}
		transfers = append(transfers, t)
	}
	return
}

func (tt *UniqueSale) ListAll(start int) (transfers []UniqueSale, err error) {
	query := `select id,tokenid,lookup_id,operation, blocknumber,index,txhash,buyer,seller,price,hash,timestamp   from unique_sale where lookup_id=$1 offset $2 limit $3`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	rows, err := db.QueryContext(ctx, query, tt.LookupID, start, 50)
	if err != nil {
		return
	}
	return tt.getSales(rows)

}

// FindByAddress returns transfers of specific token to or from an address
func (tt *UniqueSale) FindByAddress(addr string) (transfers []UniqueSale, err error) {
	query := `select id,tokenid,lookup_id,operation, blocknumber,index,txhash,buyer,seller,price,hash,timestamp   from unique_sale  where lookup_id = $1 and tokenid=$2 and (buyer=$3 or seller=$3)`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TokenID, addr)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

func (tt *UniqueSale) FindAllByAddress(addr string) (transfers []UniqueSale, err error) {
	query := `select id,tokenid,lookup_id,operation, blocknumber,index,txhash,buyer,seller,price,hash,timestamp from unique_sale
		 where (buyer=$1 or seller=$1) and lookup_id = $2
		 order by blocknumber desc, index desc`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	fmt.Println(addr, tt.LookupID)
	rows, err := db.QueryContext(ctx, query, addr, tt.LookupID)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

// MaxBlock for a specified token
func (tt *UniqueSale) SearchCountOfThisToken(addr string) (max int, err error) {
	query := `select count(*) from unique_sale where (buyer=$1 or seller=$1) and lookup_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, addr, tt.LookupID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *UniqueSale) CountOfThisToken() (max int, err error) {
	query := `select count(*) from unique_sale where lookup_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *UniqueSale) CountOfThisTokenID() (max int, err error) {
	query := `select count(*) from unique_sale where lookup_id = $1 and token_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TokenID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *UniqueSale) CountOfThisTxHash() (max int, err error) {
	query := `select count(*) from unique_sale where lookup_id = $1 and txhash = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TxHash).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

// id,tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash, timestamp
func (tt *UniqueSale) FindByTokenId() (transfers []UniqueSale, err error) {
	query := `select id,tokenid,lookup_id,operation, blocknumber,index,txhash,buyer,seller,price,hash,timestamp from unique_sale  where lookup_id=$1 and tokenid=$2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TokenID)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

func (tt *UniqueSale) FindByTxHash() (transfers []UniqueSale, err error) {
	query := `select id,tokenid,lookup_id,operation, blocknumber,index,txhash,buyer,seller,price,hash,timestamp  from unique_sale  where lookup_id=$1 and txhash=$2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TxHash)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

/*
// FindAll tokens that match either on name or address
func (t *UniqueSale) FindAll() (sales []UniqueSale, err error) {
	statement := `select id,tx_hash,buyer,token_id, price from unique_sale`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	rows, err := stmt.Query()
	if err != nil {
		return
	}
	var unq UniqueSale
	for rows.Next() {
		err = rows.Scan(&unq.TxHash, &unq.Buyer, &unq.TokenID, &unq.Price)
		if err != nil {
			return
		}
		sales = append(sales, unq)
	}
	return
}

// GetAllTokensAndMaxFrom max as a starting value, from the database
func GetAllSalesAndMaxFrom(max uint64) (sales []UniqueSale, newMax uint64, err error) {
	newMax = max
	statement := `select id,tx_hash,buyer,token_id, price from unique_sale where id > $1 order by name`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	rows, err := stmt.Query(max)
	if err != nil {
		return
	}
	var unq UniqueSale
	for rows.Next() {
		var tID uint64
		err = rows.Scan(&tID, &unq.TxHash, &unq.Buyer, &unq.TokenID, &unq.Price)
		if err != nil {
			return
		}
		sales = append(sales, unq)
		if tID > newMax {
			newMax = tID
		}
	}
	return
}

// GetAllTokensAndMax from the database
func GetAllSalesAndMax() (sales []UniqueSale, max uint64, err error) {
	return GetAllSalesAndMaxFrom(uint64(0))
}

// GetAllTokens that are in the database
func GetAllSales() (sales []UniqueSale, err error) {
	sales, _, err = GetAllSalesAndMax()
	return
}
*/
