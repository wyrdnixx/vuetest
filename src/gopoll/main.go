package main

import (
	"database/sql"
	"fmt"
	"log"

	//gonfig -> config aus json file lesen
	//"github.com/wyrdnixx/mygoproject/cmd/api/modules"

	// "database/sql"
	"github.com/wyrdnixx/votegovue/src/gopoll/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"io/ioutil"
)

//var AppConfig = modules.Configuration{}

func main() {

	db := initDB("storage.db")
	migrate(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	/* Test Routes um Echo zu testen
	e.GET("/polls", func(c echo.Context) error {
		fmt.Printf("INFO: GET polls ")
		return c.JSON(200, "GET Polls")
	})

	e.PUT("/polls", func(c echo.Context) error {
		fmt.Printf("INFO: PUT polls ")
		return c.JSON(200, "PUT Polls")
	})

	e.PUT("/polls/:id", func(c echo.Context) error {
		fmt.Printf("INFO: PUT polls:id ")
		return c.JSON(200, "UPDATE Poll "+c.Param("id"))
	})
	*/

	// lese Dateien im aktuellen Verzeichniss
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}

	e.File("/", "src/gopoll/public/index.test.html")
	e.File("/vue", "src/gopoll/public/index.html")
	e.GET("/polls", handlers.GetPolls(db))
	e.PUT("/poll/:index", handlers.UpdatePoll(db))

	// Test einfaches put
	e.PUT("/put", handlers.PutTask())

	e.Logger.Fatal(e.Start(":9000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic("db nill")
	}
	fmt.Println("Info: DB initialisiert...")
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS polls(
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name VARCHAR NOT NULL,
			topic VARCHAR NOT NULL,
			src VARCHAR NOT NULL,
			upvotes INTEGER NOT NULL,
			downvotes INTEGER NOT NULL,
			UNIQUE(name)
	);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Angular','Awesome Angular', 'https://cdn.colorlib.com/wp/wp-content/uploads/sites/2/angular-logo.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Vue', 'Voguish Vue','https://upload.wikimedia.org/wikipedia/commons/thumb/5/53/Vue.js_Logo.svg/400px-Vue.js_Logo.svg.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('React','Remarkable React','https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/1200px-React-icon.svg.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Ember','Excellent Ember','https://cdn-images-1.medium.com/max/741/1*9oD6P0dEfPYp3Vkk2UTzCg.png', 1, 0);

	INSERT OR IGNORE INTO polls(name, topic, src, upvotes, downvotes) VALUES('Knockout','Knightly Knockout','https://images.g2crowd.com/uploads/product/image/social_landscape/social_landscape_1489710848/knockout-js.png', 1, 0);
`

	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
