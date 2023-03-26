package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserCreate struct {
	Username    string `form:"username"`
	Password    string `form:"password"`
	Email       string `form:"email"`
	DisplayName string `form:"displayName"`
	Age         string `form:"age"`
	Location    string `form:"location"`
	Description string `form:"description"`
}

type User struct {
	ID          primitive.ObjectID `json:"_id"`
	Email       string             `json:"email"`
	Username    string             `json:"username"`
	Avatar      string             `json:"avatar"`
	DisplayName string             `json:"displayName"`
	Description *string            `json:"description"`
	Age         int                `json:"age"`
	CreatedAt   time.Time          `json:"createdAt"`
	Location    *string            `json:"location"`
	Events      []Event            `json:"events"`
	Posts       []Post             `json:"post"`
	Level       Level              `json:"level"`
	Friends     []Friend           `json:"friends"`
	Credentials Credentials        `json:"credentials"`
}

type UserSensored struct {
	CreatedAt   time.Time `json:"createdAt"`
	Username    string    `json:"username"`
	Avatar      string    `json:"avatar"`
	DisplayName string    `json:"displayName"`
	Description *string   `json:"description"`
	Events      []Event   `json:"events"`
	Posts       []Post    `json:"post"`
	Friends     []Friend  `json:"friends"`
	Level       Level     `json:"level"`
}

type Credentials struct {
	Password      string    `json:"password"`
	Hash          []byte    `json:"hash"`
	LastRefreshed time.Time `json:"lastRefreshed"`
}

type Level struct {
	Exp    int `json:"exp"`
	Level  int `json:"level"`
	Badges int `json:"badges"`
}

type Friend struct {
	ID        primitive.ObjectID `json:"_id"`
	CreatedAt time.Time          `json:"createdAt"`
}

type Reaction struct {
	Author   primitive.ObjectID `json:"author"`
	Reaction int                `json:"reaction"`
}

type Comment struct {
	Author    primitive.ObjectID `json:"author"`
	Content   string             `json:"content"`
	Reactions []Reaction         `json:"reactions"`
	CreatedAt time.Time          `json:"createdAt"`
}

type Invite struct {
	Author    string     `json:"author"`
	Reactions []Reaction `json:"reactions"`
}

type Event struct {
	Id          primitive.ObjectID `json:"_id"`
	CreatedAt   time.Time          `json:"createdAt"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	AgeLimit    int                `json:"ageLimit"`
	Author      primitive.ObjectID `json:"author"`
	Invites     []Invite           `json:"invites"`
	Comments    []Comment          `json:"comments"`
	Reactions   []Reaction         `json:"reactions"`
}

type Post struct {
	Id        primitive.ObjectID `json:"id"`
	Author    primitive.ObjectID `json:"author"`
	Image     string             `json:"image"`
	Reactions []Reaction         `json:"reactions"`
	Comments  []Comment          `json:"comments"`
}
