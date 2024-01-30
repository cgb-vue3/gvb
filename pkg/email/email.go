package email

import (
	"akita/global"
	"github.com/jordan-wright/email"
	"go.uber.org/zap"
	"net/smtp"
)

// Emailer 邮件对象&邮件信息结构体以及模板body缓存区
type Emailer struct {
	form, to, host, user, pass, subject string
	port                                int
}

// Email 声明全局email对象
var Email *Emailer

func InitEmail(to, subject string, html string) error {
	newEmailer(to, subject)
	if err := Email.configEmail(html); err != nil {
		global.Mlog.Error("邮件发送失败", zap.Error(err))
		return err
	}
	return nil
}

func newEmailer(to, subject string) {
	Email = &Emailer{
		subject: subject,
		to:      to,
		form:    global.MConfig.Email.Form,
		host:    global.MConfig.Email.Host,
		port:    global.MConfig.Email.Port,
		user:    global.MConfig.Email.User,
		pass:    global.MConfig.Email.Pass,
	}
}

// 配置email
func (E Emailer) configEmail(html string) error {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = E.form

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{E.to}

	// 设置主题
	em.Subject = E.subject
	em.HTML = []byte(html)
	auth := smtp.PlainAuth("", "1240092443", "cuhrkdhnnixmjejb", "smtp.qq.com")
	// 连接池
	_, err := email.NewPool("smtp.qq.com:587", 5, auth)
	if err != nil {
		global.Mlog.Error("连接池配置失败", zap.Error(err))
		return err
	}
	//设置服务器相关的配置
	err = em.Send("smtp.qq.com:587", auth)

	if err != nil {
		global.Mlog.Error("发送邮件失败", zap.Error(err))
		return err
	}
	global.Mlog.Info("邮件发送成功")
	return nil
}
