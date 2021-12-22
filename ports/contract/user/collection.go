package user

import "time"

// CollectionName : this will be the name of collection
const CollectionName = "users"

// Collection : default collection structure
type Collection struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Role      string    `bson:"role"`
	Address   string    `bson:"address"`
	CreatedAt time.Time `bson:"created_at"`
}
