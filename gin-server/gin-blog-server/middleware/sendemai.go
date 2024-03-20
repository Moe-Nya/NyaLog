package middleware

import (
	"NyaLog/gin-blog-server/utils"
	"NyaLog/gin-blog-server/utils/errmsg"
	"net/smtp"
)

func SendEmail(email string, msg []byte) int {
	// 邮件服务器的认证信息
	auth := smtp.PlainAuth("", utils.Emailusername, utils.Emailpassword, utils.Smtphost)

	var users []string
	users = append(users, email)
	// 发送邮件
	err := smtp.SendMail(utils.Smtphost+":"+utils.Smtpport, auth, utils.Emailusername, users, msg)
	if err != nil {
		return errmsg.SendEmailFailed
	}
	return errmsg.SUCCESS
}
