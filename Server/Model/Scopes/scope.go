package Scopes

import (
	"context"
	"errors"
	"strings"
	"time"
	"watchtower/Server/Model/db"
	"watchtower/Server/Model/program"
	"watchtower/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collectionName string = "scopes"
var (
	TYPE_WILD_CART = "wild_cart"
	TYPE_DOMAIN    = "domain"
	TYPE_URL       = "url"
)

type Wt_Scope struct {
	ID        primitive.ObjectID `bson:"_id"`
	Asset     string             `bson:"asset"`
	ProgramID primitive.ObjectID `bson:"program_id,omitempty"`
	Type      string             `bson:"type"`
	InScope   bool               `bson:"in_scope"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"update_at"`
}

func CreateIndex() {
	db.CreateIndex(collectionName, "program_id", 1)
	db.CreateIndex(collectionName, "asset", 1)
}

func GetScope(asset string, ProgramID primitive.ObjectID) (*Wt_Scope, error) {
	watchtower_db := db.Mongo_connect()
	filter := bson.M{"program_id": ProgramID, "asset": asset}
	scope := &Wt_Scope{}
	err := watchtower_db.Collection(collectionName).
		FindOne(context.Background(), filter).Decode(&scope)

	return scope, err

}

func detectType(asset string) string {
	if strings.HasPrefix(asset, "http://") ||
		strings.HasPrefix(asset, "https://") {
		return TYPE_URL
	}

	if strings.HasPrefix(asset, "*.") {
		return TYPE_WILD_CART
	}
	if strings.Contains(asset, "*") {
		return "wildcard"
	}
	return TYPE_DOMAIN
}

func Upsert(asset string, ProgramID primitive.ObjectID, InScope bool) error {

	watchtower_db := db.Mongo_connect()
	_, err := program.GetProgramByID(ProgramID)
	if err != nil {
		return errors.New("program not found")
	}

	asset_type := detectType(asset)

	new_data := map[string]interface{}{

		"asset":      asset,
		"program_id": ProgramID,
		"type":       asset_type,
		"in_scope":   InScope,
	}

	s, e := GetScope(asset, ProgramID)

	if e != nil {

		new_data["created_at"] = time.Now().Unix()
		new_data["update_at"] = time.Now().Unix()

		filter := bson.M{"program_id": ProgramID, "asset": asset}
		update := bson.M{"$set": new_data}
		_, err = watchtower_db.
			Collection(collectionName).
			UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))

		if err != nil {
			return err
		}
	} else {
		current_data := map[string]interface{}{
			"asset":      s.Asset,
			"program_id": s.ProgramID,
			"type":       s.Type,
			"in_scope":   s.InScope,
		}
		d1, _ := utils.HashObject(current_data)
		d2, _ := utils.HashObject(new_data)
		if d1 != d2 {
			new_data["update_at"] = time.Now().Unix()

			filter := bson.M{"program_id": ProgramID, "asset": asset}
			update := bson.M{"$set": new_data}
			_, err = watchtower_db.
				Collection(collectionName).
				UpdateOne(context.Background(), filter, update)

			if err != nil {
				return err
			}
		}

	}

	return nil
}
