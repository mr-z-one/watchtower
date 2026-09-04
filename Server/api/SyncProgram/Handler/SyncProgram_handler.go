package handler

import (
	"net/http"
	"watchtower/Server/Model/Scopes"
	"watchtower/Server/Model/program"
	"watchtower/config"
)

func GetAllSyncProgram(w http.ResponseWriter, r *http.Request) {
	programConfig := config.ReadeConfig("config.json")

	program.UpsertPrograms(programConfig)

	for _, p := range programConfig {
		prog, _ := program.GetProgramByName(p.Program_Name)

		for _, s := range p.In_Scope {
			Scopes.Upsert(s, prog.ID, true)
		}
	}

	w.Write([]byte("sync is successful!"))
}
