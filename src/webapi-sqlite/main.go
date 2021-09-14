package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"                     // external kütüphane
	"github.com/mrtkmynsndev/golangsamples/models" // projedeki models package
)

func main() {

	// ilk olarak eski db siliyoruz
	os.Remove("sqlite-database.db")

	log.Println("Creating db...")

	// yeni db file oluşturuyoruz
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("Database created...")
	if err != nil {
		log.Fatal(err)
	}

	// models package db'yi configlerini ayarlıyoruz
	models.InitDB()
	models.DataManager.AddMigrations()
	models.DataManager.SeedData()

	router := gin.Default() // router nesnesi tanımlıyoruz

	// router'a ilgili fonksiyonları bind ediyoruz
	router.GET("/", index)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbum)
	router.PUT("/albums/:id", updateAlbum)
	router.DELETE("/albums/:id", deleteAlbum)

	err = router.Run("localhost:8080") // 8080 portundan yayın yapıyoruz
	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ana sayfa
// gin kütüphanesi ile Data fonksiyonu kullnarak html body dönebiliriz
func index(c *gin.Context) {
	htmlBody := `<html>
	<head>
		<title>Golang Web Api</title>
	</head>
	<body>
		<h1 style='text-align:center'>Golang Web Api with SQLite :) </h1>
	</body>
	</html>`
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlBody))
}

// albümleri dönen endpoint
// gin kütüphanesi ile json response dönüyoruz
func getAlbums(c *gin.Context) {
	albums, err := models.DataManager.GetAlbums()

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	c.IndentedJSON(http.StatusOK, albums)
}

// albüm ekleme endpointi
// gin Context üzerinden gelen datayı JSON bind ediyoruz
func postAlbum(c *gin.Context) {
	var newAlbum *models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := models.DataManager.AddAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	c.IndentedJSON(http.StatusCreated, newAlbum) // serialize the album struc
}

// ilgili albümü dönen endpoint
// gin Context nesnesindeki param'dan query ile gönderilen params yakalıyoruz
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	album, err := models.DataManager.GetAlbumByID(id)
	if err != nil {
		// log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, album)
	}
}

// albüm güncelleme endpointi
// gin Context nesnesi üzerinden gönderilen JSON body bind ediyoruz
func updateAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum models.Album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	err := models.DataManager.UpdateAlbum(id, updatedAlbum)

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, updatedAlbum)
	}
}

// albüm silme endpointi
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	err := models.DataManager.DeleteAlbum(id)

	if err != nil {
		log.Fatal(err)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
	}
}
