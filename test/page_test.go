package poppler

import (
	"testing"

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

		for _, a := range annots {
			a.Close()
		}

		page.Close()
	}

}

