package etherdb

import (
	"context"
	"testing"
	"time"
)

const TOKEN_HOST = "host=localhost user=token password='erc20' port=5432 sslmode=disable dbname=ether_cards connect_timeout=10"

func TestUnique(t *testing.T) {
	InitDB(TOKEN_HOST)

	query := `insert into unique_owners (tokenid,lookup_id,owner) 
	values ($1,$2,$3)
	on conflict on constraint unq
	do update set
	owner=$3`
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	strA := []string{"0x139B522955D54482E7662927653ABb0bFB6F19BA", "0x139B522955D54482E7662927653ABb0bFB6F19BA", "0x90Dbd11d4842aE3b51cD0AB1ecC32bD8cD756307"}
	for _, addr := range strA {
		_, err := db.ExecContext(ctx, query, 3, 1, addr)
		if err != nil {
			t.Fatal(err)
		}
	}

}
