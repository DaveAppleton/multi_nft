package etherdb

import (
	"context"
)

type Lookup struct {
	ID          int
	Description string
	Address     string
	Type        int // 1 : ERC721, 2 : ERC1155, 3 : MultiSale, 4 : Sale
	Enabled     bool
	LastChecked uint64
	Created     uint64
}

/*
-- DROP TABLE public.lookup;

CREATE TABLE public.lookup
(
	id serial NOT NULL,
    type integer NOT NULL,
    lastchecked bigint,
    description character varying(80) COLLATE pg_catalog."default",
    enabled boolean DEFAULT true,
    address character varying(44) COLLATE pg_catalog."default"

)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.lookup
    OWNER to token;
*/
// insert into lookup (address, description,type, lastchecked) values ('0x0C110D076E7f0CaEF33F1670945f4D8cA7D86793','Samourai - Sistine Chapel',1,15992772);

func GetAllMultiTokenIds() (lups []Lookup, err error) {
	return getAllTokensOfType(2)
}

func GetAllUniqueTokenIds() (lups []Lookup, err error) {
	return getAllTokensOfType(1)
}

func GetAllMultiSaleIds() (lups []Lookup, err error) {
	return getAllTokensOfType(3)
}

func GetAllUniqueSaleIds() (lups []Lookup, err error) {
	return getAllTokensOfType(4)
}

func getAllTokensOfType(tokenType int) (lups []Lookup, err error) {
	query := "select id,description,address,lastchecked from lookup where type=$1 and enabled"
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	rows, err := db.QueryContext(ctx, query, tokenType)
	if err != nil {
		return
	}
	for rows.Next() {
		var L Lookup
		err = rows.Scan(&L.ID, &L.Description, &L.Address, &L.LastChecked)
		if err != nil {
			return
		}
		lups = append(lups, L)
	}
	return
}

func (l *Lookup) UpdateLastChecked(latest uint64) (err error) {
	query := "update lookup set lastchecked=$1 where id = $2"
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	_, err = db.ExecContext(ctx, query, latest, l.ID)
	return
}

func (l *Lookup) add() (err error) {
	query := `insert into lookup (type,lastchecked,description,enabled,address,created) 
	values ($1,$2,$3,$4,$5,$2)`
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	_, err = db.ExecContext(ctx, query, l.Type, l.LastChecked, l.Description, l.Enabled, l.Address)
	return

}

func addToLookup(lastChecked uint64, address string, description string, _type int, enabled bool) (tp *Lookup, err error) {
	t := Lookup{
		LastChecked: lastChecked,
		Address:     address,
		Description: description,
		Type:        _type,
		Enabled:     enabled,
	}
	return &t, t.add()
}

func AddUniqueSaleToLookup(lastChecked uint64, address string, description string) (tp *Lookup, err error) {
	return addToLookup(lastChecked, address, description, 4, true)
}

func AddMultiSaleToLookup(lastChecked uint64, address string, description string) (tp *Lookup, err error) {
	return addToLookup(lastChecked, address, description, 3, true)
}
func AddUniqueTokenToLookup(lastChecked uint64, address string, description string) (tp *Lookup, err error) {
	return addToLookup(lastChecked, address, description, 1, true)
}
func AddMultiTokenToLookup(lastChecked uint64, address string, description string) (tp *Lookup, err error) {
	return addToLookup(lastChecked, address, description, 2, true)
}

func LookupFromID(ID int) (lup Lookup, err error) {
	var L Lookup
	query := "select id,description,address,lastchecked from lookup where ID=$1"
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	err = db.QueryRowContext(ctx, query, ID).Scan(&L.ID, &L.Description, &L.Address, &L.LastChecked)
	if err != nil {
		return
	}
	return L, nil
}
