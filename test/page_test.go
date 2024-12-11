package poppler

import (
	"testing"
	"fmt"

	"github.com/orlandofra/go-poppler"
)

func TestCloseAnnots(t *testing.T) {
	doc, err := poppler.Open("test.pdf")
	defer doc.Close()
	if err != nil {
		return
	}


	n_pages := doc.GetNPages()

	for i := 0; i < n_pages; i++ {
		page := doc.GetPage(i)
		annots := page.GetAnnots()

		fmt.Println("page: ", i)
		for _, a := range annots {
			fmt.Printf("Freeing: %v", a)
			a.Close()
		}

		page.Close()
	}

}

/* get quads for annots that are not text markups */
func __TestGetQuads(t *testing.T) {
	doc, err := poppler.Open("test.pdf")
	defer doc.Close()
	if err != nil {
		return
	}


	n_pages := doc.GetNPages()

	for i := 0; i < n_pages; i++ {
		page := doc.GetPage(i)
		annots := page.GetAnnots()

		for _, a := range annots {
			fmt.Println(a.Quads())
		}

		page.Close()
	}

}

func TestPrintText(t *testing.T) {
	doc, err := poppler.Open("test.pdf")
	defer doc.Close()
	if err != nil {
		return
	}


	n_pages := doc.GetNPages()

	for i := 0; i < n_pages; i++ {
		page := doc.GetPage(i)

		fmt.Println(page.Text())

		page.Close()
	}

}
