package handler

import (
	"encoding/json"
	"net/http"
	"watchtower/Server/Message"
	"watchtower/Server/Model/Job"
	"watchtower/Server/validator"
	custom_color "watchtower/custom_Color"
	"watchtower/rabbitmq"
)

func SendSubFinderJob(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	domain := r.PostForm["domain"]

	if len(domain) == 0 || !validator.IsUrl().MatchString(domain[0]) {
		err := Message.CreateErrorMessage("domain is required", 400)
		d, _ := json.Marshal(err)
		w.Write([]byte(d))
		return
	}

	r_message := rabbitmq.NewMessage(rabbitmq.SUBFINDER, domain[0])

	go rabbitmq.SendMessage(r_message, func(ack bool, message *rabbitmq.Message) {
		if !ack {
			custom_color.Error()("message with {%s} not deliver\n", message.UUID)
		}
	})
	Job.Insert(Job.PENDING, r_message)

	w.Write([]byte("the job schedule for execution.."))

}
