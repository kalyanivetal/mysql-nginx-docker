package app

import (
	"time"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/gorilla/mux"
	_"github.com/go-sql-driver/mysql"
)

type App struct {
	Router   *mux.Router
	Database *sql.DB
}

type Post struct {
	ID int `json:"id"`
	Date time.Time `json:"date"`
	Name string `json:"name"`
}

func (app *App) SetupRouter() {
	app.Router.
		Methods("GET").
		Path("/endpoint/{id}").
		HandlerFunc(app.getFunction)

	app.Router.
		Methods("GET").
		Path("/endpoint").
		HandlerFunc(app.getAll)


	app.Router.
		Methods("POST").
		Path("/endpoint").
		HandlerFunc(app.postFunction)
}

func (app *App) getAll(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  var posts []Post
  result, err := app.Database.Query("SELECT id, date, name from test")
  if err != nil {
    panic(err.Error())
  }
  defer result.Close()
  for result.Next() {
    var post Post
    err := result.Scan(&post.ID, &post.Date, &post.Name)
    if err != nil {
      panic(err.Error())
    }
    posts = append(posts, post)
  }
  json.NewEncoder(w).Encode(posts)
}

func (app *App) getFunction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Fatal("No ID in the path")
	}

	dbdata := &DbData{}
	err := app.Database.QueryRow("SELECT id, date, name FROM `test` WHERE id = ?", id).Scan(&dbdata.ID, &dbdata.Date, &dbdata.Name)
	if err != nil {
		log.Fatal("Database SELECT failed")
	}

	log.Println("You fetched a thing!")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(dbdata); err != nil {
		panic(err)
	}
}

/*func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
	_, err := app.Database.Exec("INSERT INTO `test` (name) VALUES ('myname')")
	if err != nil {
		log.Fatal("Database INSERT failed")
	}

	log.Println("You called a thing!")
	w.WriteHeader(http.StatusOK)
}
*/
func (app *App) postFunction(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  stmt, err := app.Database.Prepare("INSERT INTO test(id,date,name) VALUES(?,?,?)")
  if err != nil {
    panic(err.Error())
	log.Fatal("Database INSERT failed")
  }
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    panic(err.Error())
  }
  keyVal := make(map[string]string)
  json.Unmarshal(body, &keyVal)

  id := keyVal["id"]
  date := keyVal["date"]
  name := keyVal["name"]
  _, err = stmt.Exec(id,date,name)
  if err != nil {
    panic(err.Error())
	log.Fatal("Exec statement fails")
  }
  fmt.Fprintf(w, "New post was created")
}
