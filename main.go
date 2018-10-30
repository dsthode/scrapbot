package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type PageItem (
	Id string
	Url string
	Image string
	Description string
)

func main() {
	url := "https://www.elcorteingles.es/videojuegos/reacondicionados/consolas/?itemsPerPage=36&level=6&sorting=newInAsc"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error accediendo a la pagina")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatal("La respuesta no es OK")
	}
	doc, err := goquery.NewDOcumentFromReader(res.Body)
	if err != nil {
		log.Fatal("Error construyendo el arbol del documento")
	}
	doc.Find("div#product-list ul.product-list > li.product").Each(func(i int, s *goquery.Selection) {
		
	})
}
