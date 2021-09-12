package middleware

import (
	"GenshinPlanner/Backend/models"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionStr = "mongodb+srv://<user>:<password>@cluster0.fkwok.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
const dbName = "genshin-plan"
const collName = "character-level"

var collection *mongo.Collection

func init() {
	//Setup client options
	clientOpt := options.Client().ApplyURI(connectionStr)

	//Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOpt)

	if err != nil {
		log.Fatal(err)
	}

	//Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database(dbName).Collection(collName)

	fmt.Println("Collection instance created")
}

func GetAllCharLvPlan() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M
	for curr.Next(context.Background()) {
		var res bson.M
		e := curr.Decode(&res)
		if e != nil {
			log.Fatal(e)
		}

		results = append(results, res)
	}

	curr.Close(context.Background())
	return results
}

func InsertOneCharLvPlan(plan models.CharacterPlan) {
	_, err := collection.InsertOne(context.Background(), plan)

	if err != nil {
		log.Fatal(err)
	}
}

func UpdateOneCharLvPlan(plan models.CharacterPlan) {
	filter := bson.M{"charactername": plan.CharacterName}
	updateField := bson.M{
		"currentlevel": plan.CurrentLevel,
		"targetlevel":  plan.TargetLevel,
		"requiredexp":  plan.RequiredEXP,
	}
	update := bson.M{"$set": updateField}

	res, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Count : %d", res.ModifiedCount)
}

func DeleteOneCharLvPlan(planID string) {
	id, _ := primitive.ObjectIDFromHex(planID)
	filter := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
}
