package main

import (
	"gopkg.in/gomail.v2"
	"log"
	"strings"
)



func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "registrasi@hermales.co.id")
	m.SetHeader("To", "fadlimunandar99@gmail.com")
	m.SetHeader("Subject", "Registrasi Berhasil")
	m.SetBody("text/html", getTemplateRegistrationSuccessful("HM0000001", "12345677"))

	d := gomail.NewDialer("srv80.niagahoster.com", 465, "registrasi@hermales.co.id", "r3615Tr@5i")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	log.Println("email sent")
}

func getTemplateRegistration(linkRegister string) string {
	return strings.Replace(registrasi, "{{link_daftar}}", linkRegister, 1)
}

func getTemplateRegistrationSuccessful(username, password string) string {
	template := strings.Replace(registrasiBerhasil, "{{username}}", username, 2)
	template = strings.Replace(template, "{{password}}", password, 1)
	return template
}