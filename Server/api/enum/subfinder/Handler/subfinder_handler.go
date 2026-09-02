package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"watchtower/Server/Message"
	"watchtower/Server/Model/Job"
	"watchtower/Server/validator"
	custom_color "watchtower/custom_Color"
	"watchtower/rabbitmq"
	"watchtower/utils"
)

func SendSubFinderJob(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(b, &resp)

	if utils.FailOnError(err, "", nil) {
		w.Write(Message.CreateErrorMessageJson("json is invalid", 400))
		return
	}

	domain, err := utils.GetStringMap(resp, "domain")
	if err != nil {

		w.Write(Message.CreateErrorMessageJson("domain is required", 400))
		return

	}

	if domain == "" || !validator.IsUrl().MatchString(domain) {

		w.Write(Message.CreateErrorMessageJson("domain is not valid", 400))
		return
	}

	r_message := rabbitmq.NewMessage(rabbitmq.CMD_SUBFINDER, rabbitmq.TYPE_SUBDOMAIN, domain)

	go rabbitmq.SendMessage(r_message, func(ack bool, message *rabbitmq.Message) {
		if !ack {
			custom_color.Error()("message with {%s} not deliver\n", message.UUID)
		}
	})
	Job.Insert(Job.PENDING, r_message)

	w.Write([]byte(Message.CreateMessageJson("the job schedule for execution...", 200)))

}
