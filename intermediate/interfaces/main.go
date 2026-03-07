package main

import "fmt"

// collection of Function
type Notifier interface {
	Send(message string) error
}

// individual structs for each
type EmailNotifier struct {
	Email string
}

type SmsNotifier struct {
	Contact string
}

func (e *EmailNotifier) Send(message string) error {
	fmt.Printf("Email %s is sent to %s \n", message, e.Email)
	return nil
}

func (s *SmsNotifier) Send(message string) error {
	fmt.Printf("SMS %s is sent to %s \n", message, s.Contact)
	return nil
}

type NotificationService struct {
	notifier Notifier
}

func (ns *NotificationService) Notify(message string) error {
	return ns.notifier.Send(message)
}

func main() {
	emailService := NotificationService{&EmailNotifier{Email: "aakash19here@gmail.com"}}

	emailService.Notify("Hello From Email")

	smsService := NotificationService{&SmsNotifier{Contact: "+044234324"}}

	smsService.Notify("Hello From Email")
}
