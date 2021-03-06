package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/paulsevere/envy/env"
	"github.com/paulsevere/envy/fio"
)

type Action struct {
	Type    string         `json:"type" query:"type"`
	Payload []env.Variable `json:"payload" query:"payload"`
}

func Act(typ string) Action {
	return Action{Type: typ}
}

func (a Action) Load(vs []env.Variable) Action {
	a.Payload = vs
	return a
}
func (a Action) Push(v env.Variable) {
	load := []env.Variable{v}
	a.Payload = load
}

// Run asdf
func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORS())
	e.Static("/", "assets")
	e.File("/", "assets/index.html")
	e.GET("/all", getAll)
	e.POST("/update", handlePost)
	e.DELETE("/variable", handleDelete)
	e.GET("/variable", handleNew)
	e.POST("/variable", handlePost)
	e.Logger.Fatal(e.Start(":8080"))
}

func getAll(c echo.Context) (err error) {
	println("GETTING CURRENT ENV")
	return c.JSON(http.StatusOK, Act("CURRENT_ENV").Load(env.CurrentEnv()))

}

func newVar(c echo.Context) (err error) {
	a := new(Action)
	if err = c.Bind(a); err != nil {
		return
	}
	return c.JSON(http.StatusOK, a)
}

func handlePost(c echo.Context) (err error) {
	println("Post Recieved")
	v := new(env.Variable)
	if err = c.Bind(v); err != nil {
		return
	}
	err = fio.AppendToConfig(*v)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, Act("RADICAL"))

}

func handleDelete(c echo.Context) (err error) {
	println("DELETE RECEIVED")

	v := new(env.Variable)
	if err = c.Bind(v); err != nil {
		return
	}
	action := Act("DELETE_VARIABLE")
	load := []env.Variable{*v}
	action.Payload = load

	return c.JSON(http.StatusOK, action)
}

func handleNew(c echo.Context) (err error) {
	println("HANDLE GET")
	v := new(env.Variable)
	if err = c.Bind(v); err != nil {
		return
	}
	env.SetVar(*v)
	return c.JSON(http.StatusOK, v)
}
