package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	user struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

var (
	users = map[int]*user{}
	seq   = 1
	lock  = sync.Mutex{}
)

// users = append(users, user{ID: '1', Name: "Albatross"})
// users = append(users, user{ID: '2', Name: "Phonenix"})

// var Users []user

func main() {

	// http.HandleFunc("/", handler)
	// port := ":1323"
	// fmt.Println("Server running at http://localhost" + port)
	// if err := http.ListenAndServe(port, nil); err != nil {
	// 	log.Fatal(err)
	// }
	// User = append(User, us{ID: '1', Name: "Albatross"})
	// User = append(User, us{ID: '2', Name: "Phonenix"})
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello Albatross!\n")
	// })
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)

	fmt.Printf("Starting server at port 1323\n")
	log.Fatal(http.ListenAndServe(":1323", e))

	e.Logger.Fatal(e.Start(":1323"))
}

// func handler(w http.ResponseWriter, r *http.Request) {

// if r.URL.Path == "/" {
// 	// fmt.Println("albatross...!")
// 	http.ServeFile(w, r, "index.html")
// 	return
// }
// // fmt.Println("Albatross.!")
// http.NotFound(w, r)
// }

func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, users)
}

func createUser(c echo.Context) error {

	lock.Lock()
	defer lock.Unlock()
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	users[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, users[id])
}

func updateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := new(user)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	users[id].Name = u.Name
	return c.JSON(http.StatusOK, users[id])
}

func deleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, users)
}
