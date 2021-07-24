package email

import (
	"go-template/pkg/log"
	"go-template/setting"
	"time"

	"gopkg.in/mail.v2"
)

func Init() {
	ch := make(chan *mail.Message)
	go func() {
		d := mail.NewDialer(setting.Conf.SMTP.Host, setting.Conf.SMTP.Port, "user", "123456")
		if setting.Conf.SMTP.Encryption {
			d.StartTLSPolicy = mail.MandatoryStartTLS
		}
		var s mail.SendCloser
		var err error
		open := false
		for {
			select {
			case m, ok := <-ch:
				if !ok {
					log.Debugf("邮件发送队列关闭")
					return
				}
				if !open {
					if s, err = d.Dial(); err != nil {
						panic(err)
					}
					open = true
				}
				if err := mail.Send(s, m); err != nil {
					log.Warnf("邮件发送失败：", err)
				} else {
					log.Debugf("邮件已发送")
				}
			// Close the connection to the SMTP server if no email was sent in
			// the last 30 seconds.
			case <-time.After(30 * time.Second):
				if open {
					if err := s.Close(); err != nil {
						log.Warnf("无法关闭SMTP连接", err)
					}
					open = false
				}
			}
		}
	}()
}
