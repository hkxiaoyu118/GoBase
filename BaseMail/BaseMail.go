package BaseMail

import (
	"net/smtp"
	"strings"
)

func sendMail(message string, username string, password string, host string) bool {
	auth := smtp.PlainAuth("", username, password, host)
	to := []string{"hkxiaoyu118@qq.com"}
	nickname := "XXXXX"
	user := "XXX@qq.com"
	subject := "XXXXX"
	content_type := "Content-Type:text/plain;charset=UTF-8"
	body := message
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	if err == nil {
		return true
	}
	return false
}

func SendMail(username string, nickname string, password string, host string, receiverList []string, subject string, message string) bool {
	auth := smtp.PlainAuth("", username, password, host)
	to := receiverList
	content_type := "Content-Type:text/plain;charset=UTF-8"
	body := message
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + username + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, username, to, msg)
	if err == nil {
		return true
	}
	return false
}
