package etherdb

import (
	"math/big"
)

/*
create table balances (
	id serial not null primary key,
	address  varchar(44) not null,
	tokenid  integer,
	amount   numeric,
	constraint one_balance unique (tokenid,address)
);

grant select, insert, update on balances to token;
grant select,  update on balances_id_seq to token;

create index on balances (amount);

*/

type Balance struct {
	TokenID     uint64
	Address     string
	Balance     *big.Int
	TotalSupply *big.Int
}

// AddIfNotFound only adds if the txhash is not found
func (bb *Balance) AddOrUpdate() (err error) {
	statement := `insert into balances (tokenid,address,balance,totalsupply) 
					values ($1,$2,$3,$4) 
					on conflict (one_balance) 
					do update
					set balance = $3, totalsupply=$4
					where tokenid = $1 and address = $2`

	stmt, err := db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(bb.TokenID, bb.Address, bb.Balance, bb.TotalSupply)
	return
}
