package verification

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// VerifyReflectionWithLine is check reflected param for mining
func VerifyReflectionWithLine(body, payload string) (bool, int) {
	bodyArray := strings.Split(body, "\n")
	count := 0
	for l, v := range bodyArray {
		if strings.Contains(v, payload) {
			count = count + l
		}
	}
	if count != 0 {
		return true, count
	}
	return false, 0
}

// VerifyReflection is check reflected param for xss and mining
func VerifyReflection(body, payload string) bool {
	if strings.Contains(body, payload) {
		matched, _ := regexp.MatchString(fmt.Sprintf(`\"\s*\:\s*\"[^\"\n]*%s`, payload), body)
		return !matched
	}
	return false
}

// VerifyDOM is check success inject on code
func VerifyDOM(s string) bool { //(body io.ReadCloser) bool {

	body := ioutil.NopCloser(strings.NewReader(s)) // r type is io.ReadCloser
	defer body.Close()

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	check := false
	if err != nil {
		fmt.Println(err)
		return false
	}
	// Find the review items
	doc.Find(".dalfox").Each(func(i int, s *goquery.Selection) {
		check = true
	})
	if !check {
		doc.Find("#dalfox").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			check = true
		})
	}
	return check
}
