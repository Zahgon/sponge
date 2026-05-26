package mgo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Model embedded structs, add `bson: ",inline"` when defining table structs
type Model struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CreatedAt *time.Time         `bson:"created_at" json:"createdAt"`
	UpdatedAt *time.Time         `bson:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deletedAt,omitempty"`
}

// SetModelValue set model fields
func (p *Model) SetModelValue() { _ = "STUB: not implemented"; return }

// ExcludeDeleted exclude soft deleted records
func ExcludeDeleted(filter bson.M) bson.M { _ = "STUB: not implemented"; return *new(bson.M) }

// EmbedUpdatedAt embed updated_at datetime column
func EmbedUpdatedAt(update bson.M) bson.M { _ = "STUB: not implemented"; return *new(bson.M) }

// EmbedDeletedAt embed deleted_at datetime column
func EmbedDeletedAt(update bson.M) bson.M { _ = "STUB: not implemented"; return *new(bson.M) }

// ConvertToObjectIDs convert ids to objectIDs
func ConvertToObjectIDs(ids []string) []primitive.ObjectID { _ = "STUB: not implemented"; return nil }
