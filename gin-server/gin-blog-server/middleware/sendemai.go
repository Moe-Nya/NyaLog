package middleware

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendEmail(email string, msg []byte) int {
	mail := gomail.NewMessage()

	// 设置发件人
	mail.SetHeader("From", utils.Emailusername)

	// 设置收件人
	mail.SetHeader("To", email)

	// 设置邮件主题和正文
	mail.SetHeader("Subject", "Your Subject")
	mail.SetBody("text/plain", string(msg))

	smtpPort, _ := strconv.Atoi(utils.Smtpport)

	// SMTP 配置
	smtp := gomail.NewDialer(utils.Smtphost, smtpPort, utils.Emailusername, utils.Emailpassword)

	// 发送邮件
	if err := smtp.DialAndSend(mail); err != nil {
		fmt.Println(err)
		return errmsg.SendEmailFailed
	}

	return errmsg.SUCCESS
}
