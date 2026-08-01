package get_text

import (
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

func Scrape(site string, class_tag string) [][3]string {
	cacheDir := "data/cache"
	os.MkdirAll(cacheDir, 0755)

	c := colly.NewCollector(
		colly.CacheDir(cacheDir),
	)
	// sllice for tiltle , img_url , url
	var final_texts [][3]string
	configure_onhtml(c, &final_texts, class_tag)

	err := c.Visit(site)
	if err != nil {
		return nil
	} else {

		return final_texts
	}

}
func processPost(e *colly.HTMLElement, final_texts *[][3]string) {
	title := e.Text

	// Link
	link := e.Attr("href")
	if link != "" && !strings.HasPrefix(link, "http") {
		link = e.Request.AbsoluteURL(link)
	}

	// Image - search up to parent containers
	img := ""
	for i := 0; i < 5; i++ {
		parent := e.DOM.Parent()
		if parent.Length() > 0 {
			img = parent.Find("img").AttrOr("src", "")
			if img != "" {
				break
			}
			e.DOM = parent
		}
	}

	if img != "" && !strings.HasPrefix(img, "http") {
		img = e.Request.AbsoluteURL(img)
	}

	*final_texts = append(*final_texts, [3]string{title, img, link})
}

func configure_onhtml(c *colly.Collector, final_texts *[][3]string, class_tag string) {
	c.OnHTML(class_tag, func(e *colly.HTMLElement) {
		processPost(e, final_texts)
	})

}
