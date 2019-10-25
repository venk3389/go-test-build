package main

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/labstack/echo"
	"math"
	"net/http"
)

type Url struct {
	UrlString string `json:"UrlString"`
}

func HomeHandler(c echo.Context) error {
	// Please note the the second parameter "home.html" is the template name and should
	// be equal to one of the keys in the TemplateRegistry array defined in main.go
	return c.Render(http.StatusOK, "index.html", nil)
}

func randomBase64String(l int) string {
	buff := make([]byte, int(math.Round(float64(l)/float64(1.33333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)
	return str[:l] // strip the one extra byte we get from half the results.
}
