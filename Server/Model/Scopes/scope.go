package Scopes

import (
	"watchtower/Server/Model/db"
	"watchtower/Server/Model/program"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collectionName string = "scopes"

type Wt_Scope struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	asset     string             `bson:"asset"`
	ProgramID primitive.ObjectID `bson:program_id,omitempty`
	Type      string             `bson:"type"`
	InScope   bool               `bson:"in_scope"`
	CreatedAt int64              `bson:"created_at"`
}

func Upsert(asset string, ProgramID primitive.ObjectID, Type string, InScope bool) (*Wt_Scope, error) {

	watchtower_db := db.Mongo_connect()
	_, err := program.GetProgramByID(ProgramID)
	if err != nil {
		return nil, err
	}

}
