package main

import "fmt"

func main() {
	whiteList := []FCList{
		{
			FCType: "BRANCH",
			FCCode: "KZ01",
		},
	}
	blackList := []FCList{
		{
			FCType: "STORE",
			FCCode: "TZ01",
		},
	}

	currentBranch := "KZ01"
	currentStore := "K446"

	fmt.Println(isFcOK(whiteList, blackList, currentBranch, currentStore))
}

// FCList ---
type FCList struct {
	FCType string `json:"fcType"`
	FCCode string `json:"fcCode"`
}

func isFcOK(wl, bl []FCList, currentBranch, currentStore string) bool {
	wlInit := true
	blInit := false

	if len(wl) > 0 {
		wlInit = false
		for _, w := range wl {
			if !wlInit {
				wlInit = w.FCCode == currentBranch || w.FCCode == currentStore
			}
		}
	}

	if len(bl) > 0 {
		for _, w := range bl {
			blInit = w.FCCode == currentBranch || w.FCCode == currentStore
		}
	}

	return wlInit && !blInit
}
