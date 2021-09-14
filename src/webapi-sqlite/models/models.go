package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // external sqlite kütüphanesi
)

// Album ile ilgili Data Service
type AlbumDataManager interface {
	AddMigrations()
	SeedData()
	GetAlbumByID(id string) (Album, error)
	AddAlbum(album *Album) error
	GetAlbums() ([]Album, error)
	UpdateAlbum(id string, album Album) error
	DeleteAlbum(id string) error
}

// AlbumDataManager için sql.DB'yi dependency injection uygulayacağız
type albumDataManager struct {
	db *sql.DB
}

// Entitiy
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var DataManager AlbumDataManager

func InitDB() error {
	var err error

	db, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}

	// DataManager'a albumDataManager ile initialize ediyoruz
	DataManager = &albumDataManager{db: db}

	return db.Ping() // sqlite db ping atıyoruz.
}

// AlbumDataManager interface implementasyonları
func (manager *albumDataManager) AddMigrations() {
	createAlbumTable := `CREATE TABLE album (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"title" TEXT,
		"artist" TEXT,
		"price" REAL		
	  );` // SQL Statement for Create Table

	_, err := manager.db.Exec(createAlbumTable)

	if err != nil {
		log.Printf("%q: %s\n", err, createAlbumTable)
		return
	}

	log.Println("Initial Database created...")
}

func (manager *albumDataManager) SeedData() {
	for _, v := range albums {
		insertIntoAlbum(manager, v.Title, v.Artist, v.Price)
	}
}

func insertIntoAlbum(manager *albumDataManager, title string, artist string, price float64) {
	log.Println("Inserting album record ...")
	insertAlbumSQL := `INSERT INTO album(title, artist, price) VALUES (?, ?, ?)`

	statement, err := manager.db.Prepare(insertAlbumSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer statement.Close()
	_, err = statement.Exec(title, artist, price)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// getAlbums responds with the list of all albums as JSON.
func (manager *albumDataManager) GetAlbums() ([]Album, error) {
	rows, err := manager.db.Query("select * from album")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var albumList []Album

	for rows.Next() {
		var album Album
		err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price)
		if err != nil {

			return nil, err
		}

		albumList = append(albumList, album)
	}
	return albumList, nil
}

func (manager *albumDataManager) AddAlbum(album *Album) error {

	insertAlbumSQL := `INSERT INTO album(title, artist, price) VALUES (?, ?, ?)`

	statement, err := manager.db.Prepare(insertAlbumSQL)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(album.Title, album.Artist, album.Price)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func (manager *albumDataManager) GetAlbumByID(id string) (Album, error) {
	var album Album
	row := manager.db.QueryRow("Select * From album Where id = $1", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			return album, fmt.Errorf("album not found")
		}
		return album, err
	}

	return album, nil
}

func (manager *albumDataManager) UpdateAlbum(id string, album Album) error {
	if id != album.ID {
		panic("Album not matched")
	}

	statement, err := manager.db.Prepare("Update album Set Title =?, Artist=?, Price=? where id=?")
	if err != nil {
		return err
	}
	defer statement.Close()

	statement.Exec(album.Title, album.Artist, album.Price, id)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}

func (manager *albumDataManager) DeleteAlbum(id string) error {

	statement, err := manager.db.Prepare("delete from album where id=?")
	if err != nil {
		return err
	}
	defer statement.Close()

	statement.Exec(id)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return nil
}
