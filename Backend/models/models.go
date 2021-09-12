package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CharacterPlan struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CharacterName string             `json:"characterName"`
	CurrentLevel  int                `json:"currentLevel"`
	TargetLevel   int                `json:"targetLevel"`
	RequiredEXP   int                `json:"requiredEXP"`
}
