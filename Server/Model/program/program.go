package program

import (
	"context"
	"time"
	"watchtower/Server/Model/db"
	custom_color "watchtower/custom_Color"
	"watchtower/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionName string = "programs"

type Wt_Program struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Program_Name string             `json:"program_name" bson:"program_name"`
	In_Scope     []string           `json:"in_scope" bson:"in_scope"`
	Out_Scope    []string           `json:"out_scope" bson:"out_scope"`
}

func GetProgramByID(ID primitive.ObjectID) (*Wt_Program, error) {
	watchtower_db := db.Mongo_connect()
	var result *Wt_Program
	err := watchtower_db.Collection(collectionName).
		FindOne(context.TODO(), bson.D{{"_id", ID}}).Decode(&result)

	return result, err
}

func GetProgramByName(name string) (*Wt_Program, error) {
	watchtower_db := db.Mongo_connect()
	var result *Wt_Program
	err := watchtower_db.Collection(collectionName).
		FindOne(context.TODO(), bson.D{{"program_name", name}}).Decode(&result)

	return result, err
}

func GetAllProgram() ([]*Wt_Program, error) {
	watchtower_db := db.Mongo_connect()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := watchtower_db.Collection(collectionName).Find(ctx, bson.D{})

	e := utils.FailOnError(err, "", nil)
	if e {

		return []*Wt_Program{}, err
	}

	var results []*Wt_Program
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results, nil
}

func UpsertProgram(programs []*Wt_Program) {

	watchtower_db := db.Mongo_connect()
	models := make([]mongo.WriteModel, 0, len(programs))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, myprogram := range programs {

		program_copy := utils.StructToMap(myprogram, "ID")
		model := mongo.NewUpdateOneModel().SetFilter(bson.M{
			"program_name": myprogram.Program_Name,
		}).SetUpdate(bson.M{
			"$set": program_copy,
		}).SetUpsert(true)

		models = append(models, model)
	}
	if len(models) == 0 {
		custom_color.Error()("no any program provide to update")
	}
	_, err := watchtower_db.Collection(collectionName).BulkWrite(ctx, models)

	utils.FailOnError(err, "can't update program", nil)

	//dbconn := db.Mongo_connect()

}
