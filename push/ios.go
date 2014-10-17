package push

import (
	"github.com/gysan/room/config"
	"github.com/anachronistic/apns"
	"github.com/golang/glog"
)

type Message struct {
	Token string
	Text  string
}

func Send(messages []*Message) {
	glog.Info("APNS push send ...")
	client := apns.NewClient(config.ApnsUrl, config.ApnsCertificate, config.ApnsKey)

	message := &Message{}
	for _, message = range messages {
		glog.Infof("Push message: %v", message)
		payload := apns.NewPayload()
		payload.Alert = message.Text
		payload.Badge = 1
		payload.Sound = "default"

		pn := apns.NewPushNotification()
		pn.DeviceToken = message.Token
		pn.AddPayload(payload)

		response := client.Send(pn)
		glog.Infof("APNS response: %v", response)
		alert, _ := pn.PayloadString()
		glog.Infof("APNS alert: %v", alert)
	}
	glog.Info("APNS push send end")
}
