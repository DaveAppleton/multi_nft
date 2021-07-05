package etherdb

import (
	"context"

	_ "github.com/lib/pq" // DB selection
)

// Token holds ERC20 token data

//item := etherdb.UniqueSale{filt.Event.Raw.TxHash.Hex(), filt.Event.Buyer.Hex(), filt.Event.TokenId, etherUtils.EtherToStr(filt.Event.Price)}
type UniqueSale struct {
	ID          uint64
	LookupRef   int
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
	lookupRef int,
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
		LookupRef:   lookupRef,
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
	//                                         1       2           3        4     5      6     7     8     9     10       11       12
	statement := `insert into unique_sale 
				(tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash,timestamp) 
				values 
				($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) 
				returning id`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	//                         1       2              3             4           5          6          7           8        9         10           11           12
	err = db.QueryRowContext(ctx, statement,
		tt.TokenID, tt.LookupRef, tt.Operation, tt.BlockNumber, tt.TxIndex, tt.TxHash, tt.Buyer, tt.Seller, tt.Price, tt.Hash, tt.Timestamp).Scan(&tt.ID)
	return
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
