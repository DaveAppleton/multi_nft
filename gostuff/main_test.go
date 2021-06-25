package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DaveAppleton/ether_go/ethKeys"
)

func TestAll(t *testing.T) {
	mainnet := ethKeys.NewKey("../.mainnet")
	mainnet.RestoreOrCreate()
	fmt.Println(mainnet.PublicKeyAsHexString())
	t.Fail()
}

func chkTErr(t *testing.T, msg string, err error) {
	if err == nil {
		return
	}
	t.Fatal(msg, err)
}

func TestShares(t *testing.T) {
	sam := ethKeys.NewKey("userKeys/samourai")
	pboy := ethKeys.NewKey("userKeys/pboy")
	minty := ethKeys.NewKey("userKeys/minty")
	patron := ethKeys.NewKey("userKeys/patron")
	kol := ethKeys.NewKey("userKeys/kol")

	chkTErr(t, "samourai", sam.RestoreOrCreate())
	chkTErr(t, "minty", minty.RestoreOrCreate())
	chkTErr(t, "patron", patron.RestoreOrCreate())
	chkTErr(t, "pboy", pboy.RestoreOrCreate())
	chkTErr(t, "kol", kol.RestoreOrCreate())

	typeA := []*ethKeys.AccountKey{patron, sam, pboy, kol}
	blue := []int{5, 8, 2, 0}
	yellow := []int{5, 3, 0, 7}
	pink := []int{5, 6, 2, 2}
	blueArr := makeShares(typeA, blue)
	yellowArr := makeShares(typeA, yellow)
	pinkArr := makeShares(typeA, pink)
	_ = yellowArr
	_ = pinkArr
	fmt.Println("blue")
	for _, data := range blueArr {
		fmt.Println(data.Share)
	}
	fmt.Println("yellow")
	for _, data := range yellowArr {
		fmt.Println(data.Share)
	}
	fmt.Println("pink")
	for _, data := range pinkArr {
		fmt.Println(data.Share)
	}
	t.Fail()
}

func TestDemoDB(t *testing.T) {
	//var max int
	var user string
	connectString := "host=localhost user=postgres password=P@ssword1 port=5432 sslmode=disable dbname='minty-demo' connect_timeout=10"
	dbx, err := sql.Open("postgres", connectString) //
	if err != nil {
		t.Fatal(err)
	}
	err = dbx.Ping()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("connected to ", connectString)
	q1 := "select a.address from users a, artist b where a.id=b.user_id and a.nickname='AmirHussain'"
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	row := dbx.QueryRowContext(ctx, q1)

	err = row.Scan(&user)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(user)
	t.Fail()
}

/*
SELECT grantee, privilege_type
 FROM information_schema.role_table_grants
 WHERE table_name='artist';

*/
