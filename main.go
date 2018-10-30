package main

import (
	"container/list"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type pageItem struct {
	Ruta        string
	Imagen      string
	Descripcion string
	Precio      string
	Fecha       time.Time
}

func scrapePage() *list.List {
	items := list.New()
	url := "https://www.elcorteingles.es/videojuegos/reacondicionados/consolas/?itemsPerPage=36&level=6&sorting=newInAsc"
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error accediendo a la pagina")
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Println("La respuesta no es OK")
		return nil
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println("Error construyendo el arbol del documento")
		return nil
	}
	doc.Find("div#product-list ul.product-list > li.product").Each(func(i int, s *goquery.Selection) {
		tmpItem := new(pageItem)
		tmpItem.Imagen, _ = s.Find(".product-image img").Attr("src")
		tmpItem.Ruta, _ = s.Find(".product-image a.event").Attr("href")
		tmpItem.Descripcion, _ = s.Find(".product-image a.evetn").Attr("title")
		tmpItem.Precio = s.Find(".product-price current").Text()
		tmpItem.Fecha = time.Now()
		items.PushBack(tmpItem)
	})
	return items
}

func mergeLists(scrapedItems *list.List, currentItems *list.List) *list.List {
	newItems := list.New()
	found := false
	for e := scrapedItems.Front(); e != nil; e = e.Next() {
		found = false
		for e2 := currentItems.Front(); e2 != nil; e2 = e2.Next() {
			if e.Value.(pageItem).Ruta == e.Value.(pageItem).Ruta {
				found = true
			}
		}
		if !found {
			// El elemento extraido de la pagina no esta en la lista de items actuales
			// por lo que es un elemento nuevo y hay que insertarlo en la lista de nuevos items
			newItems.PushBack(e)
		}
	}
	for e := currentItems.Front(); e != nil; e = e.Next() {
		found = false
		for e2 := scrapedItems.Front(); e2 != nil; e2 = e2.Next() {
			if e.Value.(pageItem).Ruta == e.Value.(pageItem).Ruta {
				found = true
			}
		}
		if found {
			// El elemento de la lista de items actuales esta en la lista de items
			// extraidos de la pagina por lo que sigue existiendo y hay que insertarlo
			// en la lista de nuevos items
			newItems.PushBack(e)
		}
	}
	return newItems
}

func main() {

}
