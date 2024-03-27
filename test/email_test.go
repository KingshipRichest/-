package test

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "Get <getcharzhaopan@163.com>"
	e.To = []string{"1433698861@qq.com"}
	e.Subject = "验证码发送测试"
	e.Text = []byte("您的验证码：<b>123456</b>")
	//err := e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "cygouhie@gmail.com", "lx", "smtp.gmail.com"))
	// 返回EOF时，关闭SSL重试
	err := e.SendWithTLS("smtp.gmail.com:465",
		smtp.PlainAuth("", "getcharzhaopan@163.com", "PAGHEICPAEQFQKHO", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		t.Fatal(err)
	}
}
