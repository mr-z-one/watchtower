package Job

import (
	"context"
	"time"
	"watchtower/Server/Model/db"
	custom_color "watchtower/custom_Color"
	"watchtower/rabbitmq"
	"watchtower/utils"
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
	CreatedAt int64             `bson:"created_at"`
	UpdatedAt int64             `bson:"updated_at"`
}

func NewJob(state string, message *rabbitmq.Message) *Job {
	craetedAt := time.Now().Unix()

	return &Job{State: state, Message: message,
		UUID: message.UUID, CreatedAt: craetedAt, UpdatedAt: craetedAt}
}

func InitializeIndexes() {
	db.CreateIndex(collectionName, "uuid", 1)
}

func Insert(state string, Message *rabbitmq.Message) {

	job := NewJob(state, Message)

	watchtower_db := db.Mongo_connect()

	collecton := watchtower_db.Collection(collectionName)

	ctx := context.Background()
	r, err := collecton.InsertOne(ctx, job)

	e := utils.FailOnError(err, "can't insert Job", nil)

	if !e {
		custom_color.Succeed()("job_id : %s has inserted to DB", r.InsertedID)
	}

}
