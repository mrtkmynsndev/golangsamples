package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretKey = "EE1w2FSMC2GIZDkLNKcuKmukCL6PIeja"
)

type Authentication struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

type JwtToken struct {
	Token string `json:"token"`
}

func main() {
	os.Remove("./user.db")

	AddMigration()

	app := fiber.New(fiber.Config{})

	app.Post("/signup", SignUp)

	app.Post("/signin", SignIn)

	app.Use("/users", AuthenticationMiddleware())
	app.Get("/users", UserListHandler)

	app.Use("/user", AuthenticationMiddleware())
	app.Get("/user", GetUserHandler)

	log.Fatal(app.Listen(":3000"))
}

func SignUp(c *fiber.Ctx) error {
	var user User
	err := c.BodyParser(&user)
	if err != nil {
		log.Fatal(err)
	}

	db := GetDatabase()
	defer db.Close()

	var counter int
	row := db.QueryRow("select 1 from user where username = $1 limit 1", user.UserName)
	if err := row.Scan(&counter); err == nil && counter > 0 {
		return c.JSON("Username is already in use!")
	}

	inserUserSQL := `insert into user(firstname, lastname, email, username, password) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(inserUserSQL)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	defer statement.Close()

	hashPassword, err := GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	_, err = statement.Exec(user.FirstName, user.LastName, user.Email, user.UserName, hashPassword)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	return c.JSON(user)
}

func SignIn(c *fiber.Ctx) error {
	var auth Authentication
	err := c.BodyParser(&auth)
	if err != nil {
		log.Fatal(err)
	}

	db := GetDatabase()
	defer db.Close()

	var dbUser User
	row := db.QueryRow("select username, password from user where username=$1", auth.UserName)
	if err := row.Scan(&dbUser.UserName, &dbUser.Password); err != nil {
		if err == sql.ErrNoRows {
			return c.JSON("Username or password is invalid!")
		}
	}

	isPasswordMatched := CheckPasswordHash(auth.Password, dbUser.Password)
	if !isPasswordMatched {
		return c.JSON("Username or password is invalid!")
	}

	token, err := GenerateJWTToken(dbUser.UserName)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.SendString(err.Error())
	}

	var jwtToken JwtToken
	jwtToken.Token = token

	return c.JSON(jwtToken)
}

func UserListHandler(ctx *fiber.Ctx) error {
	db := GetDatabase()
	defer db.Close()

	rows, err := db.Query("select firstname, lastname, email, username, password from user")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.FirstName, &user.LastName, &user.Email, &user.UserName, &user.Password)
		if err != nil {
			return err
		}

		users = append(users, user)
	}

	return ctx.JSON(users)
}

func GetUserHandler(ctx *fiber.Ctx) error {
	username, err := ExtractTokenMetaData(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON("Failed extract token")
	}

	var db = GetDatabase()
	defer db.Close()

	var user User
	row := db.QueryRow("select * from user where username = $1", *username)
	if err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.UserName, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return ctx.Status(fiber.StatusBadRequest).JSON("User can not found")
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func GetDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./user.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func AddMigration() {
	db := GetDatabase()

	defer db.Close()

	createAlbumTable := `CREATE TABLE user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"firstname" TEXT,
		"lastname" TEXT,
		"email" TEXT,			
		"username" TEXT,
		"password" TEXT	
	  );`

	_, err := db.Exec(createAlbumTable)

	if err != nil {
		log.Printf("%q: %s\n", err, createAlbumTable)
		return
	}

	log.Println("Migration finished...")
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWTToken(username string) (string, error) {
	var signinKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["expr"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(signinKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ExtractTokenMetaData(ctx *fiber.Ctx) (*string, error) {
	token, err := VerifyToken(ctx)
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	if claims != nil && token.Valid {
		username := claims["username"].(string)
		return &username, nil
	}

	return nil, err
}

func AuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		token, err := VerifyToken(c)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}

		if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}
}

func VerifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	var authorizationHeader = c.Get("Authorization")
	if authorizationHeader == "" {
		return nil, fmt.Errorf("Unauthorized")
	}

	splittedToken := strings.Split(authorizationHeader, " ")
	if len(splittedToken) < 2 {
		return nil, fmt.Errorf("Unauthorized")
	}

	if splittedToken[0] != "Bearer" {
		return nil, fmt.Errorf("Unauthorized")
	}

	var jwtToken = splittedToken[1]

	var signingKey = []byte(secretKey)

	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return signingKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
