package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./y.sqlite")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type TogyzGame struct {
	Id        int
	WhiteName string
	BlackName string
	Result    string
	Event     string
	Date      string
	Site      string
	Notation  string
}

func GetGames(count int) []TogyzGame {

	rows, err := DB.Query("SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games LIMIT " + strconv.Itoa(count))

	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer rows.Close()

	games := make([]TogyzGame, 0)

	for rows.Next() {
		game := TogyzGame{}
		err = rows.Scan(&game.Id, &game.WhiteName, &game.BlackName, &game.Result, &game.Event, &game.Date, &game.Site, &game.Notation)

		if err != nil {
			log.Fatal(err)
			return nil
		}

		games = append(games, game)
	}

	err = rows.Err()

	if err != nil {
		log.Fatal(err)
		return nil
	}

	return games
}

func GetGameById(id string) TogyzGame {

	stmt, err := DB.Prepare("SELECT id, _WhiteName, _BlackName, _Result, _Event, _Date, _Site, _Notation FROM games WHERE id = ?")

	if err != nil {
		log.Fatal(err)
		return TogyzGame{}
	}

	game := TogyzGame{}

	sqlErr := stmt.QueryRow(id).Scan(&game.Id, &game.WhiteName, &game.BlackName, &game.Result, &game.Event, &game.Date, &game.Site, &game.Notation)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return TogyzGame{}
		}
		log.Fatal(sqlErr)
		return TogyzGame{}
	}

	//    game.Notation = strings.Replace(game.Notation,"\n","<br>",-1)
	return game
}
