package Job

import (
	"context"
	"time"

	"watchtower/Server/Model/db"
	custom_color "watchtower/custom_Color"
	"watchtower/rabbitmq"
	"watchtower/utils"

	"go.mongodb.org/mongo-driver/bson"
)

var (
	PENDING   = "pending"
	EXECUTION = "execution"
	FAILED    = "failed"
	FINISH    = "finish"
)
var collectionName string = "Jobs"

type Job struct {
	State     string            `bson:"state"`
	Message   *rabbitmq.Message `bson:"message"`
	UUID      string            `bson:"uuid,omitempty"`
	Type      string            `json:"type" bson:"type"`
	CreatedAt int64             `bson:"created_at"`
	UpdatedAt int64             `bson:"updated_at"`
}

func NewJob(state string, message *rabbitmq.Message) *Job {
	craetedAt := time.Now().Unix()

	return &Job{State: state, Message: message,
		UUID: message.UUID, Type: message.Type, CreatedAt: craetedAt, UpdatedAt: craetedAt}
}

func CreateIndex() {
	db.CreateIndex(collectionName, "uuid", 1)
}

func Update(state string, uuid string) {
	watchtower_db := db.Mongo_connect()

	collecton := watchtower_db.Collection(collectionName)
	ctx := context.Background()

	filter := bson.M{"uuid": uuid}
	update := bson.M{"$set": bson.M{"state": state, "updated_at": time.Now().Unix()}}

	r, err := collecton.UpdateOne(ctx, filter, update)

	utils.FailOnError(err, "can't update Job", nil)
	if r.MatchedCount != 1 {
		custom_color.Warning()("jobs with id \"%s\" not Found \n", uuid)
	}
}
func Insert(state string, Message *rabbitmq.Message) {

	job := NewJob(state, Message)

	watchtower_db := db.Mongo_connect()

	collecton := watchtower_db.Collection(collectionName)

	ctx := context.Background()
	_, err := collecton.InsertOne(ctx, job)

	e := utils.FailOnError(err, "can't insert Job", nil)

	if !e {
		custom_color.Succeed()("job_id : %s has inserted to DB\n", Message.UUID)
	}

}
