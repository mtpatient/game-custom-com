package utility

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(content string, to []string) error {
	mailUserName := "mt-z-weixiang@qq.com" //邮箱账号
	mailPassword := "jkggxaaekgahbceh"     //邮箱授权码
	addr := "smtp.qq.com:465"              //TLS地址
	host := "smtp.qq.com"                  //邮件服务器地址
	Subject := "找回密码"                      //发送的主题

	e := email.NewEmail()
	e.From = "mt-z-weixiang <mt-z-weixiang@qq.com>"
	e.To = to
	e.Subject = Subject
	//e.HTML = []byte("你的验证码为：<h1>" + code + "</h1> <p>有效期为5分钟。<p>")
	e.HTML = []byte(content)
	err := e.SendWithTLS(addr, smtp.PlainAuth("", mailUserName, mailPassword, host),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
	if err != nil {
		return err
	}

	return nil
}
