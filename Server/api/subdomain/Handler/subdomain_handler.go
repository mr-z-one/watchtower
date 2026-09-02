package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"watchtower/Server/Message"
	"watchtower/utils"
)

func AddSubdomain(w http.ResponseWriter, r *http.Request) {
	resp := map[string]any{}
	b, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(b, &resp)

	if utils.FailOnError(err, "", nil) {
		w.Write(Message.CreateErrorMessageJson("json is invalid", 400))
		return
	}

	uuid, err := utils.GetStringMap(resp, "uuid")
	if utils.FailOnError(err, "", nil) {
		w.Write(Message.CreateErrorMessageJson("uuid is required", 400))
		return
	}
	reponses, err := utils.GetArrayStringMap(resp, "responses")
	if utils.FailOnError(err, "", nil) {
		w.Write(Message.CreateErrorMessageJson("responses array is required", 400))
		return
	}

	w.Write(Message.CreateErrorMessageJson(uuid+"\n"+strings.Join(reponses, " "), 400))
}
