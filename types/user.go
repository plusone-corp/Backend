package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserCreate struct {
	Username    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	Email       string `form:"email" json:"email"`
	DisplayName string `form:"displayName" json:"displayName"`
	Age         string `form:"age" json:"age"`
	Location    string `form:"location" json:"location"`
	Description string `form:"description" json:"description"`
}

type User struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Email       string             `bson:"email" json:"email"`
	Username    string             `bson:"username" json:"username"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	DisplayName string             `bson:"displayName" json:"displayName"`
	Description *string            `bson:"description" json:"description"`
	Age         int                `bson:"age" json:"age"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	Location    *string            `bson:"location" json:"location"`
	Events      []Event            `bson:"events" json:"events"`
	Posts       []Post             `bson:"posts" json:"posts"`
	Level       Level              `bson:"level" json:"level"`
	Friends     []Friend           `bson:"friends" json:"friends"`
	Credentials Credentials        `bson:"credentials" json:"credentials"`
}

type UserSensored struct {
	CreatedAt   time.Time `bson:"createdAt" json:"createdAt"`
	Username    string    `bson:"username" json:"username"`
	Avatar      string    `bson:"avatar" json:"avatar"`
	DisplayName string    `bson:"displayName" json:"displayName"`
	Description *string   `bson:"description" json:"description"`
	Events      []Event   `bson:"events" json:"events"`
	Posts       []Post    `bson:"posts" json:"posts"`
	Friends     []Friend  `bson:"friends" json:"friends"`
	Level       Level     `bson:"level" json:"level"`
}

type Credentials struct {
	Password      string    `bson:"password" json:"password"`
	Hash          []byte    `bson:"hash" json:"hash"`
	LastRefreshed time.Time `bson:"lastRefreshed" json:"lastRefreshed"`
}

type Level struct {
	Exp    int `bson:"exp" json:"exp"`
	Level  int `bson:"level" json:"level"`
	Badges int `bson:"badges" json:"badges"`
}

type Friend struct {
	ID        primitive.ObjectID `bson:"id" json:"ID"`
	CreatedAt time.Time          `bson:"createdAt" json:"CreatedAt"`
}

type Reaction struct {
	Author   primitive.ObjectID `bson:"author" json:"author"`
	Reaction int                `bson:"reaction" json:"reaction"`
}

type Comment struct {
	Author    primitive.ObjectID `bson:"author" json:"author"`
	Content   string             `bson:"content" json:"content"`
	Reactions []Reaction         `bson:"reactions" json:"reactions"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
}

type Invite struct {
	Author    string     `bson:"author" json:"author"`
	Reactions []Reaction `bson:"reactions" json:"reactions"`
}

type Event struct {
	Id          primitive.ObjectID `bson:"id" json:"id"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	AgeLimit    int                `bson:"ageLimit" json:"ageLimit"`
	Author      primitive.ObjectID `bson:"author" json:"author"`
	Invites     []Invite           `bson:"invites" json:"invites"`
	Comments    []Comment          `bson:"comments" json:"comments"`
	Reactions   []Reaction         `bson:"reactions" json:"reactions"`
}

type PostCreate struct {
	Id          primitive.ObjectID `form:"id" json:"id"`
	Image       string             `form:"image" json:"image"`
	Reactions   []Reaction         `form:"reactions" json:"reactions"`
	Comments    []Comment          `form:"comments" json:"comments"`
	Description string             `form:"description" json:"description"`
	Title       string             `form:"title" json:"title"`
	Author      primitive.ObjectID `form:"author" json:"author"`
}

type Post struct {
	Id          primitive.ObjectID `bson:"id" json:"id"`
	Description string             `bson:"description" json:"description"`
	Title       string             `bson:"title" json:"title"`
	Author      primitive.ObjectID `bson:"author" json:"author"`
	Image       string             `bson:"image" json:"image"`
	Reactions   []Reaction         `bson:"reactions" json:"reactions"`
	Comments    []Comment          `bson:"comments" json:"comments"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
