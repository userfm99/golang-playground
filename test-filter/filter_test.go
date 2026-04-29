package main

import (
	"fmt"
	"testing"
)

var (
	mStore  = "K4467"
	mBranch = "K427"
)

// Test1 test empty whitelist and blacklist
func Test1(t *testing.T) {
	wl := make([]FCList, 0)
	bl := make([]FCList, 0)

	if ok := isFcOK(wl, bl, mBranch, mStore); !ok {
		t.Errorf("value: %v", ok)
	}
}

func Test2(t *testing.T) {
	wl := []FCList{{"BRANCH", "K427"}}
	bl := make([]FCList, 0)

	if ok := isFcOK(wl, bl, mBranch, mStore); !ok {
		t.Errorf("value: %v", ok)
	}
}

func Test3(t *testing.T) {
	wl := []FCList{{"STORE", "KZ01"}}
	bl := make([]FCList, 0)

	if ok := isFcOK(wl, bl, mBranch, mStore); !ok {
		t.Errorf("value: %v", ok)
	}
}

func Test4(t *testing.T) {
	wl := []FCList{{"STORE", "KZ01"}, {"BRANCH", "K427"}}
	bl := make([]FCList, 0)

	if ok := isFcOK(wl, bl, mBranch, mStore); !ok {
		t.Errorf("value: %v", ok)
	}
}

func Test5(t *testing.T) {
	wl := []FCList{{"BRANCH", "K427"}}
	bl := []FCList{{"STORE", "KZ01"}}

	if ok := isFcOK(wl, bl, mBranch, mStore); ok {
		t.Errorf("value: %v, expected: %v", ok, !ok)
	}
}

func Test6(t *testing.T) {
	bl := []FCList{{"BRANCH", "K427"}}
	wl := []FCList{{"STORE", "KZ01"}}

	if ok := isFcOK(wl, bl, mBranch, mStore); !ok {
		t.Errorf("value: %v, expected: %v", !ok, ok)
	}
}

func Test7(t *testing.T) {
	bl := []FCList{{"BRANCH", "K427"}}
	wl := make([]FCList, 0)

	if ok := isFcOK(wl, bl, mBranch, mStore); ok {
		t.Errorf("value: %v, expected: %v", ok, !ok)
	}
}

func Test8(t *testing.T) {
	wl := []FCList{{"BRANCH", "K427"}}
	bl := []FCList{{"STORE", "KZ01"}}

	if ok := isFcOK(wl, bl, mBranch, mStore); ok {
		t.Errorf("value: %v, expected: %v", ok, !ok)
	}
}

func Test9(t *testing.T) {
	wl := []FCList{{"STORE", "K446"}}
	bl := []FCList{{"BRANCH", "TZ01"}}

	fmt.Printf("value: %v\n", isFcOK(wl, bl, mBranch, mStore))
}
