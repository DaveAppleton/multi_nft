package etherdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq" // DB selection
)

/* MultiTokenTransfers
create table erc1155sales (
	id serial not null primary key,
	lookup_id    int,
	operation    varchar(20),
	tokenid int not null,
	blocknumber int,
	blockhash varchar(70)  not null,
	index int,
	txhash varchar(70) not null,

    token varchar(44) not null,
    buyer varchar(44) not null,
    seller varchar(44) not null,

    quantity  integer,
	price     numeric(20,18),
	hash      varchar(46),

    "timestamp" integer,
    Amount int not null
);
QmTVrehVAuuYNF6GCU4WUyu63gcNsjdLb1mPckR3Dfr5fc
12345
*/

type MultiTokenSale struct {
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

func NewMultiSale(
	LookupID int,
	operation string,
	tokenid uint64,
	blocknumber uint64,
	txIndex uint,
	txHash string,
	token string,
	buyer string,
	seller string,
	quantity uint64,
	price string,
	position uint64,
	hash string,
	timestamp uint64) (mt MultiTokenSale) {
	mt = MultiTokenSale{
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
		Position:    position,
		Quantity:    quantity,
		Timestamp:   timestamp,
	}
	return
}

// Add this token to the database
func (tt *MultiTokenSale) Add() (err error) {
	//                                         1       2           3        4       5      6     7     8     9     10       11       12
	statement := `insert into multi_sales (tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash,position,quantity,timestamp) 
					values                ($1,     $2,       $3,       $4,         $5,   $6,    $7,   $8,    $9,   $10,  $11,    $12,     $13) returning id`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	//                                             1       2              3             4               5          6          7           8        9         10           11           12        13
	err = db.QueryRowContext(ctx, statement, tt.TokenID, tt.LookupID, tt.Operation, tt.BlockNumber, tt.TxIndex, tt.TxHash, tt.Buyer, tt.Seller, tt.Price, tt.Hash, tt.Position, tt.Quantity, tt.Timestamp).Scan(&tt.ID)
	return
}

// AddIfNotFound only adds if the txhash is not found
func (tt *MultiTokenSale) AddIfNotFound() (err error) {
	query := `insert into multi_sales (tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash,position,quantity,timestamp) 
					values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13) on conflict do nothing`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	_, err = db.ExecContext(ctx, query, tt.TokenID, tt.LookupID, tt.Operation, tt.BlockNumber, tt.TxIndex, tt.TxHash, tt.Buyer, tt.Seller, tt.Price, tt.Hash, tt.Position, tt.Quantity, tt.Timestamp)
	return
}

// id,tokenid,lookup_id,operation,blocknumber,index,txhash,buyer,seller,price,hash,position,quantity,timestamp
func (*MultiTokenSale) getSales(rows *sql.Rows) (transfers []MultiTokenSale, err error) {
	var t MultiTokenSale
	for rows.Next() {
		err = rows.Scan(&t.ID, &t.TokenID, &t.LookupID, &t.Operation, &t.BlockNumber, &t.TxIndex, &t.TxHash, &t.Buyer, &t.Seller, &t.Price, &t.Hash, &t.Position, &t.Quantity, &t.Timestamp)
		if err != nil {
			return
		}
		transfers = append(transfers, t)
	}
	return
}

// Find transfers that match tokenID
func (tt *MultiTokenSale) Find() (transfers []MultiTokenSale, err error) {
	statement := `select id,tokenid,lookup_id, blocknumber,index,txhash,buyer,seller,price,hash,position, quantity,timestamp from multi_sales where tokenid=$1`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	rows, err := stmt.Query(tt.TokenID)
	if err != nil {
		return
	}
	return tt.getSales(rows)

}

func (tt *MultiTokenSale) ListAll(start int) (transfers []MultiTokenSale, err error) {
	query := `select id,tokenid,lookup_id, operation, blocknumber,index,txhash,buyer,seller,price,hash,position, quantity,timestamp from multi_sales where lookup_id=$1  offset $2 limit $3`

	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, start, 50)
	if err != nil {
		return
	}
	return tt.getSales(rows)

}

// FindByAddress returns transfers of specific token to or from an address
func (tt *MultiTokenSale) FindByAddress(addr string) (transfers []MultiTokenSale, err error) {
	statement := `select id,tokenid,lookup_id, blocknumber,index,txhash,buyer,seller,price,hash,position,quantity,timestamp from multi_sales  where  (source=$1 or dest=$1)`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	rows, err := stmt.Query(addr)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

// FindAllByAddress returns transfers of any token to or from an address newest first
func (tt *MultiTokenSale) FindAllByAddress(addr string) (transfers []MultiTokenSale, err error) {
	statement := `select id,tokenid,lookup_id, operation, blocknumber,index,txhash,buyer,seller,price,hash,position, quantity,timestamp from multi_sales
		 where (source=$1 or dest=$1) and lookup_id=$2
		 order by blocknumber desc, index desc`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	rows, err := stmt.Query(addr, tt.LookupID)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

func (tt *MultiTokenSale) FindByTxHash() (transfers []MultiTokenSale, err error) {
	query := `select id,tokenid,lookup_id, operation, blocknumber,index,txhash,buyer,seller,price,hash,position, quantity,timestamp from multi_sales
		 where  lookup_id = $1 and txhash = $2
		 order by blocknumber desc, index desc`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TxHash)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

func (tt *MultiTokenSale) FindByTokenId() (transfers []MultiTokenSale, err error) {
	query := `select id,tokenid,lookup_id, operation, blocknumber,index,txhash,buyer,seller,price,hash,position,quantity,timestamp from multi_sales
		 where  lookup_id = $1 and tokenid = $2
		 order by blocknumber desc, index desc`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	rows, err := db.QueryContext(ctx, query, tt.LookupID, tt.TokenID)
	if err != nil {
		return
	}
	return tt.getSales(rows)
}

// MaxBlock for a specified token
func (tt *MultiTokenSale) MaxBlock() (max uint64, err error) {
	statement := `select max(blocknumber) from multi_sales`
	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	err = stmt.QueryRow(tt.TokenID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

// MaxBlock for a specified token
func (tt *MultiTokenSale) SearchCountOfThisToken(addr string) (max int, err error) {
	query := `select count(*) from multi_sales where (source=$1 or dest=$1) and lookup_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, addr, tt.LookupID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *MultiTokenSale) CountOfThisToken() (max int, err error) {
	query := `select count(*) from multi_sales where lookup_id = $1`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *MultiTokenSale) CountOfThisTokenID() (max int, err error) {
	query := `select count(*) from multi_sales where lookup_id = $1 and token_id = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TokenID).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}

func (tt *MultiTokenSale) CountOfThisTxHash() (max int, err error) {
	query := `select count(*) from multi_sales where lookup_id = $1 and txhash = $2`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()

	err = db.QueryRowContext(ctx, query, tt.LookupID, tt.TxHash).Scan(&max)
	if err != nil {
		max = 0
		err = nil
	}
	return
}
