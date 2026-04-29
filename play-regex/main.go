package main

import (
	"fmt"
	"regexp"
)

func main() {
	// str1 := `https://www.alfacart.com/catalog/makanan-minuman-buah-segar-723`

	// re := regexp.MustCompile(`http[s]?:\/\/([^\/]+\/){6}(?<field1>[^\/]+)\/`)
	// fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	// fmt.Println(re.MatchString(str1))        // true

	// for _, n := range re.FindAllString(str1, 10) {
	// 	fmt.Println(n)
	// 	fmt.Println("---")
	// }

	// for _, i := range re.FindAllStringSubmatch(str1, 10) {
	// 	fmt.Println(i[0])
	// }

	fmt.Print("zuul test below\n\n")

	v := "${url.mapi}/oauth/authorize"
	reg, err := regexp.Compile(`\$\{(.*?)\}`)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	match := reg.FindStringSubmatch(v)

	fmt.Println(match)
}
