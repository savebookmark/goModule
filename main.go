package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/francoispqt/gojay"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/labstack/echo/v4"
)

type user struct {
	id    int
	name  string
	email string
}

// implement gojay.UnmarshalerJSONObject
func (u *user) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "id":
		return dec.Int(&u.id)
	case "name":
		return dec.String(&u.name)
	case "email":
		return dec.String(&u.email)
	}
	return nil
}
func (u *user) NKeys() int {
	return 3
}

// Game ..
type Game struct{}

// Update ..
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw ..
func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout ..
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	// Echo instance
	// e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	// e.GET("/", hello)
	// Start server
	// e.Logger.Fatal(e.Start(":1323"))
	/////////////////////////////////////////////
	//gods
	// list := arraylist.New()
	// list.Add("list")
	//workiva
	// set := set.New()
	// set.Add("set")
	////////////////////////////////////////////////////
	// ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowTitle("Hello, World!")
	// if err := ebiten.RunGame(&Game{}); err != nil {
	// 	log.Fatal(err)
	// }
	///////////////////////////////////////////
	// sbf := boom.NewDefaultStableBloomFilter(10000, 0.01)
	// fmt.Println("stable point", sbf.StablePoint())

	// sbf.Add([]byte(`a`))
	// if sbf.Test([]byte(`a`)) {
	// 	fmt.Println("contains a")
	// }

	// if !sbf.TestAndAdd([]byte(`b`)) {
	// 	fmt.Println("doesn't contain b")
	// }

	// if sbf.Test([]byte(`b`)) {
	// 	fmt.Println("now it contains b!")
	// }

	// Restore to initial state.
	// sbf.Reset()
	////////////////////////////////////////////////
	u := &user{}
	d := []byte(`{"id":1,"name":"gojay","email":"gojay@email.com"}`)
	err := gojay.UnmarshalJSONObject(d, u)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
