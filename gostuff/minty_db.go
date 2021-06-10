package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/lib/pq" // DB selection
	"github.com/spf13/viper"
)

// minty=# create user mintyadmin with password 'nimdaytnim';
// GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public to mintyadmin;

var (
	db *sql.DB
)

func initDB() (dbx *sql.DB, err error) {
	connectString := viper.GetString("DB_HOST")
	dbx, err = sql.Open("postgres", connectString) //
	if err != nil {
		fmt.Println("open  database : ", err)
		log.Println("open  database : ", err)
		log.Println(connectString)
		return
	}
	err = dbx.Ping()
	if err != nil {
		fmt.Println("database : ", err)
		log.Println("database : ", err)
		log.Println(connectString)
	}
	return
}

func updateProjectWithContract(projectID int, contract common.Address) (err error) {
	db, err := initDB()
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println("update ", projectID, "with", contract.Hex())
	query := "update project set token_address=$1, approved=true where id=$2"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	res, err := db.ExecContext(ctx, query, contract.Hex(), projectID)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("%d row(s) changed", count)
	}
	return
}

type ProjectRecord struct {
	UserID          int
	Name            string
	Address         string
	ArtistID        int
	ArtistApproved  bool
	Title           string
	ProjectID       int
	ProjectApproved bool
	ProjectToken    sql.NullString
}

func getProjectsForUser(userName string) (pa []ProjectRecord, err error) {
	db, err := initDB()
	if err != nil {
		return []ProjectRecord{}, err
	}
	defer db.Close()
	query := "select a.id, a.nickname, a.address, b.id, b.approved, c.approved, c.title, c.id as project_id, c.token_address from users a, artist b, project c where a.id=b.user_id and b.id = c.artist_id and a.nickname=$1"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row, err := db.QueryContext(ctx, query, userName)
	if err != nil {
		return []ProjectRecord{}, err
	}
	for row.Next() {
		var p ProjectRecord
		err = row.Scan(&p.UserID, &p.Name, &p.Address, &p.ArtistID, &p.ArtistApproved, &p.ProjectApproved, &p.Title, &p.ProjectID, &p.ProjectToken)
		if err != nil {
			return []ProjectRecord{}, err
		}
		pa = append(pa, p)
	}
	return pa, nil
}
