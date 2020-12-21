package getdoc

import (
	"fmt"
	"net/url"
	"strconv"
	"unicode"

	"github.com/PuerkitoBio/goquery"
)

func modifyHrefs(selection *goquery.Selection) []string {
	var hrefs []string
	modifyHrefsRecursively(selection, make(map[int]struct{}), &hrefs)
	return hrefs
}

func modifyHrefsRecursively(selection *goquery.Selection, accum map[int]struct{}, hrefs *[]string) {
	if _, processed := accum[selection.Index()]; processed {
		return
	}

	if path, ok := selection.Attr("href"); ok {
		accum[selection.Index()] = struct{}{}

		*hrefs = append(*hrefs, path)

		text, cut := cutRightSpaces(selection.Text())
		text += "[" + strconv.Itoa(len(*hrefs)) + "]"
		text += cut

		selection.SetText(text)
	}

	selection.Find("*").Each(func(i int, s *goquery.Selection) {
		modifyHrefsRecursively(s, accum, hrefs)
	})
}

func cutRightSpaces(input string) (string, string) {
	var (
		r   = []rune(input)
		cut []rune
	)

	for i := len(r) - 1; i >= 0; i-- {
		if unicode.IsSpace(r[i]) {
			cut = append(cut, r[i])
			r = r[:i]
		} else {
			break
		}
	}

	return string(r), string(cut)
}

func createLinks(hrefs []string) string {
	text := "Links:\n"
	for i, href := range hrefs {
		u, err := url.Parse(href)
		if err != nil {
			panic(err)
		}

		if u.Host == "" {
			href = "https://core.telegram.org" + href
		}

		text += fmt.Sprintf("[%d] %s\n", i+1, href)
	}

	return text
}
