package main

import (
	"fmt"
	"strings"
)

// expected
// if(eq(min(promoPriceSort_KZ01_96,promoPriceSort_KZ01_97),0),def(min(priceSort_KZ01_96, priceSort_KZ01_97),priceSort_KZ01_63),min(promoPriceSort_KZ01_96,promoPriceSort_KZ01_97))+asc

const promoPrefix = "promoPriceSort"
const defaultPrefix = "priceSort"

func main() {
	groupID := map[int][]string{4: {"63"}, 2: {"96"}}
	// groupID := map[int][]string{4: {"63"}}
	priority := []int{2, 4}

	s := getPriceSortField("KZ01", groupID, priority, "priceSort")

	fmt.Println(s)

}

func getPriceSortField(branch string, groups map[int][]string, sortedPriorities []int, prefix string) (s string) {
	for idx, x := range sortedPriorities {
		if groupIDs, exist := groups[x]; exist {
			defaultPriceSortArray := make([]string, 0)
			promoPriceSortArray := make([]string, 0)
			for _, groupID := range groupIDs {
				defaultPriceSortArray = append(defaultPriceSortArray, prefix+"_"+branch+`_`+groupID)
				promoPriceSortArray = append(promoPriceSortArray, promoPrefix+"_"+branch+`_`+groupID)
			}

			defaultString := `min(` + strings.Join(defaultPriceSortArray, `,`) + `)`
			promoString := `min(` + strings.Join(promoPriceSortArray, `,`) + `)`

			switch len(sortedPriorities) > 1 {
			case false:
				s = fmt.Sprintf("def(%s,%s)", promoString, defaultString)
				return
			case true:
				if idx == 0 {
					s = `if(eq(` + promoString + `,0),{next},` + promoString + `)`
					break
				}
				if len(sortedPriorities)-idx > 1 {
					s = strings.Replace(s, `{next}`, `if(eq(`+defaultString+`,0),{next},`+defaultString+`)`, 1)
				} else {
					s = strings.Replace(s, `{next}`, defaultString, 1)
				}
				break
			default:
				return
			}
		}
	}
	return
}
