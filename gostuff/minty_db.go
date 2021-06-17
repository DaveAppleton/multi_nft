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

func getContractMetaData(projectID int) (str string, err error) {
	db, err := initDB()
	if err != nil {
		return "", err
	}
	defer db.Close()
	query := "select metadata from project where id=$1"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = db.QueryRowContext(ctx, query, projectID).Scan(&str)
	return str, err
}

type Artist struct {
	UserID    int
	Address   string
	Nickname  sql.NullString
	Avatar    sql.NullString
	ArtistID  int
	AboutMe   sql.NullString
	Approved  bool
	Facebook  sql.NullString
	Twitter   sql.NullString
	Instagram sql.NullString
	Portfolio sql.NullString
}

func getArtists(start int) (pa []Artist, max int, err error) {
	db, err := initDB()
	if err != nil {
		return []Artist{}, 0, err
	}
	defer db.Close()
	q1 := "select count(*) from users a, artist b where a.id=b.user_id order by 1"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = db.QueryRowContext(ctx, q1).Scan(&max)
	if err != nil {
		fmt.Println("cannot get max")
		return []Artist{}, 0, err
	}
	query := "select a.id, a.address, a.nickname, a.avatar, b.id, b.about_me, b.approved, b.facebook, b.twitter, b.instagram, b.portfolio from users a, artist b where a.id=b.user_id order by a.id limit 50 offset $1"
	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row, err := db.QueryContext(ctx, query, start)
	if err != nil {
		return []Artist{}, 0, err
	}
	for row.Next() {
		var p Artist
		err = row.Scan(
			&p.UserID,
			&p.Address,
			&p.Nickname,
			&p.Avatar,
			&p.ArtistID,
			&p.AboutMe,
			&p.Approved,
			&p.Facebook,
			&p.Twitter,
			&p.Instagram,
			&p.Portfolio)
		if err != nil {
			return []Artist{}, 0, err
		}
		pa = append(pa, p)
	}
	return
}

func findArtist(name string) (p *Artist, err error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := "select a.id, a.address, a.nickname, a.avatar, b.id, b.about_me, b.approved, b.facebook, b.twitter, b.instagram, b.portfolio from users a, artist b where a.id=b.user_id and a.nickname=$1"
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err = db.QueryRowContext(ctx, query, projectID).Scan(
		&p.UserID,
		&p.Address,
		&p.Nickname,
		&p.Avatar,
		&p.ArtistID,
		&p.AboutMe,
		&p.Approved,
		&p.Facebook,
		&p.Twitter,
		&p.Instagram,
		&p.Portfolio)
	return
}
