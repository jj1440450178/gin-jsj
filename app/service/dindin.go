package service

import (
	"gin-jj/app"

	"github.com/wanghuiyt/ding"
)

func DindinSend(str string) error {
	d := ding.Webhook{
		AccessToken: app.C.Dindin.AccessToken,
		Secret:      app.C.Dindin.Secret,
	}

	err := d.SendMessageText(str)
	if err != nil {
		return err
	}
	return nil
}
