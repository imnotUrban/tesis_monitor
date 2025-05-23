package main

import (
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

func main() {
	frontendURL := os.Getenv("FRONTEND_URL")
	backendURL := os.Getenv("BACKEND_URL")

	if frontendURL == "" || backendURL == "" {
		log.Fatal("FRONTEND_URL y BACKEND_URL deben estar definidos")
	}

	for {
		checkService(frontendURL, "Frontend")
		checkService(backendURL, "Backend")
		time.Sleep(120 * time.Minute)
	}
}

func checkService(url string, serviceName string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error al verificar %s: %v", serviceName, err)
		sendEmail(serviceName+" caído", "Hola, el servicio "+serviceName+" no está respondiendo. Por favor, revisa el estado del servicio.")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Estado del servicio %s: %s", serviceName, resp.Status)
		sendEmail(serviceName+" con problemas", "Hola, el servicio "+serviceName+" está respondiendo con el estado: "+resp.Status)
	} else {
		log.Printf("Estado del servicio %s: OK", serviceName)
	}
}

func sendEmail(subject string, body string) {
	smtpServer := "smtp.gmail.com"
	smtpPort := 587
	senderMail := os.Getenv("SENDER_MAIL")
	senderPass := os.Getenv("SENDER_PASS")
	receiverMail := os.Getenv("RECEIVER_MAIL")

	if senderMail == "" || senderPass == "" || receiverMail == "" {
		log.Println("Variables de entorno SENDER_MAIL, SENDER_PASS o RECEIVER_MAIL no definidas")
		return
	}

	auth := smtp.PlainAuth("", senderMail, senderPass, smtpServer)

	to := []string{receiverMail}
	msg := []byte("To: " + receiverMail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, senderMail, to, msg)
	if err != nil {
		log.Printf("Error al enviar el correo: %v", err)
	}
}
