package etherdb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

var db *sql.DB

// InitDB has to be called before any activity but needs viper
func InitDB(connectString string) {
	var err error
	db, err = sql.Open("postgres", connectString) //
	if err != nil {
		fmt.Println("open token database : ", err)
		log.Fatal("open token database : ", err)
	}
	fmt.Println("database is open")
	//defer db.Close()
}

func timeout() time.Duration {
	return time.Duration(viper.GetInt("DB_TIMEOUT")) * time.Second
}

func CreateMultiTokenTable() (err error) {
	query := `create table multi_tokens (
		id serial not null primary key,
		lookup_id int not null,
		tokenid int not null,
		blocknumber int,
		index int,
		txhash varchar(70) not null,
		operator varchar(44) not null,
		source varchar(44) not null,
		dest varchar(44) not null,
		"timestamp" integer,
		Amount int not null
	)`
	return create(query)
}

func CreateUniqueTokenTable() (err error) {
	query := `create table unique_tokens (
		id serial not null primary key,
		lookup_id int not null,
		tokenid int not null,
		blocknumber int,
		index int,
		txhash varchar(70) not null,
		source varchar(44) not null,
		dest varchar(44) not null,
		"timestamp" integer
	)`
	return create(query)
}

func CreateLookupTable() (err error) {
	query := `CREATE TABLE public.lookup
	(
		id serial NOT NULL primary key,
		type integer NOT NULL,
		lastchecked bigint,
		created bigint,
		description character varying(80) COLLATE pg_catalog."default",
		enabled boolean DEFAULT true,
		address character varying(44) COLLATE pg_catalog."default"
	
	)`
	return create(query)
}

/*
alter table multi_sales alter column price type numeric(22,18);
alter table unique_sale alter column price type numeric(22,18);
ALTER TABLE multi_tokens RENAME COLUMN token_ref TO lookup_id;
ALTER TABLE unique_tokens RENAME COLUMN token_ref TO lookup_id;
*/

func CreateMultiSaleTable() (err error) {
	query := `CREATE TABLE multi_sales
	(
		id serial NOT NULL,
		tokenid integer NOT NULL,
		lookup_id integer,
		operation character varying(20),
		blocknumber integer,
		index integer,
		txhash character varying(70) NOT NULL,
		buyer character varying(44)  NOT NULL,
		seller character varying(44) NOT NULL,
		price numeric(22,18),
		hash character varying(46),
		"position" integer,
		quantity integer,
		"timestamp" integer
	)`
	return create(query)
}

func CreateUniqueSaleTable() (err error) {
	query := `CREATE TABLE unique_sale
	(
		id serial NOT NULL,
		tokenid integer NOT NULL,
		lookup_id integer,
		operation character varying(20),
		blocknumber integer,
		index integer,
		txhash character varying(70) NOT NULL,
		buyer character varying(44)  NOT NULL,
		seller character varying(44) NOT NULL,
		price numeric(22,18),
		hash character varying(46),
		"timestamp" integer
	)`
	return create(query)
}

func CreateUniqueOwnersTable() (err error) {
	query := `CREATE TABLE unique_owners
	(
		id serial NOT NULL,
		tokenid integer NOT NULL,
		lookup_id integer,
		owner character varying(44)  NOT NULL,
		CONSTRAINT unq UNIQUE (tokenid, lookup_id)
	)`
	if err = create(query); err != nil {
		return
	}
	index := `create unique index unique_owners_index on unique_owners (tokenid,lookup_id)`
	return create(index)
}

func create(query string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout())
	defer cancel()
	_, err = db.ExecContext(ctx, query)
	return
}
