package gink_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/userfm99/golang-playground/play-gingkgo/gink"
)

var _ = Describe("Gink", func() {
	var (
		longBook  Gink
		shortBook Gink
	)

	BeforeEach(func() {
		longBook = Gink{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  1488,
		}

		shortBook = Gink{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing book length", func() {
		Context("with more than 300 pages", func() {
			It("Should be a novel", func() {
				Expect(longBook.Author).To(Equal("Victor Hugo"))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("Should be Dr. Seuss", func() {
				Expect(shortBook.Author).To(Equal("Dr. Seuss"))
			})
		})
	})
})
