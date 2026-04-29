package main

import (
	"fmt"
	"strings"
)

type StaticContent struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Img      string `json:"image"`
	DeepLink string `json:"deepLink"`
}

type Contents struct {
	Statics []StaticContent
}

func main() {
	/*link := "https://cfstg.alfacart.com/media/revamp/apps/assets/img/icon/ic_e-service_paket data.png"
	u, _ := url.Parse(link)
	fmt.Println(u.String())

	str1 := `https://www.alfacart.com/catalog/makanan-minuman-buah-segar-723`

	splits := strings.Split(str1, "/")

	for _, s := range splits {
		fmt.Println(s)
	}

	fmt.Println("///")
	fmt.Println(buildDeeplink(&str1))*/

	testUrl := "https://gomapistg.alfacart.com/v5/eservices/numbers/add"
	fmt.Println("is contain", testContain(testUrl, "/v5/eservices"))
}

func buildDeeplink(url *string) (deeplink string) {
	if !isURLValid(url) {
		return ""
	}
	splits := strings.Split(*url, "/")
	term := splits[3]
	lastPath := splits[len(splits)-1]

	prefix := "alfacart://"

	switch term {
	case "promo":
		deeplink = prefix + "promo/"
	case "catalog":
		deeplink = prefix + "catalog/" + getCategoryIdFromPath(&lastPath)
	default:
		deeplink = prefix + "webview/"
	}
	return
}

func isURLValid(url *string) bool {
	s := strings.Split(*url, "/")
	length := len(s)
	if length > 3 && strings.Contains(s[0], "https") && strings.Contains(s[2], "alfacart.com") {
		return true
	}
	return false
}

func getCategoryIdFromPath(path *string) string {
	splits := strings.Split(*path, "-")
	length := len(splits)
	categoryID := splits[length-1]
	return categoryID
}

func testContain(url string, word string) bool {
	return strings.Contains(url, word)
}
