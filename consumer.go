package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"text/template"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formatedMsg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: Test Email\r\n\r\nHi %s, \nThis is a test mail Campain !!!\r\n", "arvind@arvind-sharma.co.in", recipient.Email, recipient.Name)
		// msg := []byte(formatedMsg)
		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("worker: %d Error parsing template for %s", id, recipient.Email)
			continue
		}

		fmt.Println("Worker:", id, "Sending Email to.", recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "arvind@arvind-sharma.co.in", []string{recipient.Email}, []byte(msg))
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Microsecond * 100)
	}

}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer
	err = t.Execute(&tpl, r)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
