package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Username    string             `bson:"username"`
	DisplayName string             `bson:"displayName"`
	CreatedAt   int64              `bson:"createdAt"`
	IconURL     string             `bson:"iconURL,omitempty"`
	ID          primitive.ObjectID `bson:"_id"`
	Email       string             `bson:"email"`
	Posts       []Post             `bson:"posts,omitempty"`
	Events      []Event            `bson:"events"`
	Credentials Credentials        `bson:"credentials"`
	Friends     []string           `bson:"friends"`
}

type SensoredUser struct {
	Username    string             `bson:"username"`
	DisplayName string             `bson:"displayName"`
	CreatedAt   int64              `bson:"createdAt"`
	IconURL     string             `bson:"iconURL,omitempty"`
	ID          primitive.ObjectID `bson:"_id"`
	Email       string             `bson:"email"`
	Posts       []Post             `bson:"posts,omitempty"`
	Events      []Event            `bson:"events"`
	Friends     []string           `bson:"friends"`
}

type Credentials struct {
	Hash          []byte `bson:"hash"`
	Password      string `bson:"password"`
	RefreshToken  string `bson:"refreshToken"`
	LastRefreshed int64  `bson:"lastRefreshed"`
}

type Event struct {
	ID        primitive.ObjectID `bson:"_id"`
	Date      string             `bson:"date"`
	Title     string             `bson:"title"`
	Author    string             `bson:"author"`
	Invites   []Invite           `bson:"invites"`
	Public    bool               `bson:"public"`
	Reactions []Reaction         `bson:"reaction"`
	Comments  []Comment          `bson:"comment"`
	CreatedAt int64              `bson:"createdAt"`
	Post      []Post             `bson:"post"`
	Location  [2]string          `bson:"location"`
}

type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt int64              `bson:"createdAt"`
	Title     string             `bson:"title"`
	Comments  []Comment          `bson:"comments"`
	Reactions []Reaction         `bson:"reactions"`
}

type Reaction struct {
	Author    primitive.ObjectID `bson:"author"`
	CreatedAt int64              `bson:"createdAt"`
	Type      int                `bson:"type"`
}

type Comment struct {
	Author    primitive.ObjectID `bson:"author"`
	CreatedAt int64              `bson:"createdAt"`
	Content   string             `bson:"content"`
	Type      int                `bson:"type"`
}

type Invite struct {
	Author    primitive.ObjectID `bson:"author"`
	Target    primitive.ObjectID `bson:"target"`
	CreatedAt string             `bson:"createdAt"`
	Status    string             `bson:"status"`
}
