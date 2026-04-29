package main

import (
	"log"
	"strings"
	"testing"
)

func TestValidity(t *testing.T) {
	str1 := "https://www.alfacart.com/catalog/makanan-minuman-buah-segar-723"

	if isURLValid(&str1) {
		log.Println("url valid")
	} else {
		t.Errorf("url invalid: %s", str1)
	}
}

func TestGetCategoryID(t *testing.T) {
	str1 := "https://www.alfacart.com/catalog/makanan-minuman-buah-segar-723"
	splits := strings.Split(str1, "/")
	lastPath := splits[len(splits)-1]

	log.Println("Lastpath ", lastPath)
	categoryID := getCategoryIdFromPath(&lastPath)
	if categoryID != "723" {
		t.Errorf("Error getting category id: Expected: %s. Actual: %s\n", "723", categoryID)
	}
	log.Printf("Category ID is: %s\n", categoryID)
}

func TestBuildDeepLink(t *testing.T) {
	str1 := "https://www.alfacart.com/catalog/makanan-minuman-buah-segar-723"

	deepLink := buildDeeplink(&str1)

	if deepLink == "" {
		t.Errorf("error building deeplink. %s\n", deepLink)
	}

	log.Println("The deeplink", deepLink)
}
