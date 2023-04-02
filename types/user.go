package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserForm struct {
	Username    string `form:"username" json:"username"`
	Password    string `form:"password" json:"password"`
	Email       string `form:"email" json:"email"`
	DisplayName string `form:"displayName" json:"displayName" default:""`
	Age         int    `form:"age" json:"age" default:"0"`
}

type UserResponse struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Email       string             `bson:"email" json:"email"`
	Username    string             `bson:"username" json:"username"`
	Avatar      string             `bson:"avatar" json:"avatar"`
	DisplayName string             `bson:"displayName" json:"displayName"`
	Description string             `bson:"description" json:"description"`
	Age         int                `bson:"age" json:"age"`
	Location    GeoJSON            `bson:"location" json:"location"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	Events      []Event            `bson:"events" json:"events"`
	Level       Level              `bson:"level" json:"level"`
	Friends     []UserFiltered     `bson:"friends" json:"friends"`
	Setting     Setting            `json:"setting" bson:"setting"`
}

type User struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id"`
	Email       string               `bson:"email" json:"email"`
	Username    string               `bson:"username" json:"username"`
	Avatar      string               `bson:"avatar" json:"avatar"`
	DisplayName string               `bson:"displayName" json:"displayName"`
	Description string               `bson:"description" json:"description"`
	Age         int                  `bson:"age" json:"age"`
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt" default:"$now"`
	Events      []primitive.ObjectID `bson:"events" json:"events"`
	Location    GeoJSON              `bson:"location" json:"location"`
	Level       Level                `bson:"level" json:"level"`
	Friends     []primitive.ObjectID `bson:"friends" json:"friends"`
	Credentials Credentials          `bson:"credentials" json:"credentials"`
	Setting     Setting              `json:"setting" bson:"setting"`
}

type UserFiltered struct {
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt"`
	Username    string               `bson:"username" json:"username"`
	Avatar      string               `bson:"avatar" json:"avatar"`
	Events      []primitive.ObjectID `bson:"events" json:"events"`
	Location    GeoJSON              `bson:"location" json:"location"`
	DisplayName string               `bson:"displayName" json:"displayName"`
	Description string               `bson:"description" json:"description"`
	Friends     []primitive.ObjectID `bson:"friends" json:"friends"`
	Level       Level                `bson:"level" json:"level"`
	Setting     Setting              `json:"setting" bson:"setting"`
}

type Setting struct {
	Privacy Privacy `json:"privacy" json:"privacy"`
}

type Privacy struct {
	ShareLocation      bool `json:"shareLocation" bson:"shareLocation" default:"true"`
	PrivateProfile     bool `json:"privateProfile" bson:"privateProfile" default:"false"`
	AllowFriendRequest bool `json:"allowFriendRequest" bson:"allowFriendRequest" default:"true"`
	AllowInvite        bool `json:"allowInvite" bson:"allowInvite" default:"true"`
}

type GeoJSON struct {
	Type        string    `bson:"type" json:"type"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type Credentials struct {
	Password      string    `bson:"password" json:"password" default:""`
	Hash          []byte    `bson:"hash" json:"hash"`
	RefreshToken  string    `bson:"refreshToken" json:"refreshToken" default:""`
	LastRefreshed time.Time `bson:"lastRefreshed" json:"lastRefreshed" default:"$now"`
}

type Level struct {
	Exp    int `bson:"exp" json:"exp"`
	Level  int `bson:"level" json:"level"`
	Badges int `bson:"badges" json:"badges"`
}

type Reaction struct {
	Author   primitive.ObjectID `bson:"author" json:"author"`
	Reaction int                `bson:"reaction" json:"reaction"`
}

type Comment struct {
	Author    primitive.ObjectID `bson:"author" json:"author"`
	Content   string             `bson:"content" json:"content"`
	Reactions []Reaction         `bson:"reactions" json:"reactions"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt" default:"$now"`
}

type EventForm struct {
	Title       string   `form:"title" json:"title"`
	Description string   `form:"description" json:"description"`
	AgeLimit    *int     `form:"ageLimit" json:"ageLimit"`
	Invites     []string `form:"invites" json:"invites"`
	Location    string   `form:"location" json:"location"`
	Image       string   `form:"image" json:"image"`
}

type Event struct {
	Id          primitive.ObjectID   `bson:"_id" json:"_id"`
	CreatedAt   time.Time            `bson:"createdAt" json:"createdAt" default:"$now"`
	Title       string               `bson:"title" json:"title"`
	Image       string               `bson:"image" json:"image"`
	Description string               `bson:"description" json:"description" default:""`
	Location    GeoJSON              `bson:"location" json:"location"`
	AgeLimit    int                  `bson:"ageLimit" json:"ageLimit"`
	Author      primitive.ObjectID   `bson:"author" json:"author"`
	Invites     []primitive.ObjectID `bson:"invites" json:"invites"`
	Comments    []Comment            `bson:"comments" json:"comments"`
	Reactions   []Reaction           `bson:"reactions" json:"reactions"`
	Posts       []primitive.ObjectID `bson:"posts" bson:"posts"`
}

type ResEvent struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	AgeLimit    int                `bson:"ageLimit" json:"ageLimit"`
	Location    []float64          `bson:"location" json:"location"`
	Author      UserFiltered       `bson:"author" json:"author"`
	Invites     []UserFiltered     `bson:"invites" json:"invites"`
	Comments    []Comment          `bson:"comments" json:"comments"`
	Reactions   []Reaction         `bson:"reactions" json:"reactions"`
	Posts       []Post             `bson:"posts" bson:"posts"`
}

type PostCreate struct {
	Id          primitive.ObjectID `form:"_id" json:"_id"`
	Image       string             `form:"image" json:"image"`
	Reactions   []Reaction         `form:"reactions" json:"reactions"`
	Comments    []Comment          `form:"comments" json:"comments"`
	Description string             `form:"description" json:"description"`
	Title       string             `form:"title" json:"title"`
	Event       string             `form:"event" json:"event"`
}

type Post struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Description string             `bson:"description" json:"description"`
	Event       primitive.ObjectID `bson:"event" json:"event"`
	Title       string             `bson:"title" json:"title"`
	Author      primitive.ObjectID `bson:"author" json:"author"`
	Image       string             `bson:"image" json:"image"`
	Reactions   []Reaction         `bson:"reactions" json:"reactions"`
	Comments    []Comment          `bson:"comments" json:"comments"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt" default:"$now"`
}

type ResPost struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Description string             `bson:"description" json:"description"`
	Event       primitive.ObjectID `bson:"event" json:"event"`
	Title       string             `bson:"title" json:"title"`
	Author      UserFiltered       `bson:"author" json:"author"`
	Image       string             `bson:"image" json:"image"`
	Reactions   []Reaction         `bson:"reactions" json:"reactions"`
	Comments    []Comment          `bson:"comments" json:"comments"`
	CreatedAt   time.Time          `bson:"createdAt" json:"createdAt"`
}
