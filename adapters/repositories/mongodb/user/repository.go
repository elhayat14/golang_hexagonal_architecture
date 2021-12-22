package user

import (
	"context"
	"errors"
	"github.com/devfeel/mapper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mongodbUtil "golang_hexagonal_architecture/adapters/repositories/mongodb/utils/mongodb"
	portContractUser "golang_hexagonal_architecture/ports/contract/user"
	"time"
)

type MongoDb struct {
	collection *mongo.Collection
}

func NewMongoDBRepository(db *mongo.Database) (*MongoDb, error) {
	repository := MongoDb{db.Collection(portContractUser.CollectionName)}
	//to some indexing
	err := mongodbUtil.NewIndexing(
		repository.collection,
		options.Index().SetUnique(true),
		[]string{"name"})
	if err != nil {
		return nil, err
	}
	return &repository, nil
}
func (repo *MongoDb) Create(userObj portContractUser.Object) (*portContractUser.Object, error) {
	var (
		collection portContractUser.Collection
		result     portContractUser.Object
	)
	//to auto mapper
	mapper.AutoMapper(userObj, &collection)
	collection.Id = primitive.NewObjectID().Hex()
	collection.CreatedAt = time.Now()
	_, err := repo.collection.InsertOne(context.TODO(), collection)
	if err != nil {
		return nil, err
	}
	mapper.AutoMapper(collection, &result)
	return &result, nil

}
func (repo *MongoDb) Update(userId string, userObj portContractUser.Object) (*portContractUser.Object, error) {
	var (
		collection portContractUser.Collection
	)
	mapper.AutoMapper(userObj, &collection)
	updated := bson.M{
		"$set": collection,
	}
	_, err := repo.collection.UpdateByID(context.TODO(), userId, updated)
	if err != nil {
		return nil, err
	}
	updatedData, er := repo.GetById(userId)
	if er != nil {
		return nil, er
	}
	return updatedData, nil
}
func (repo *MongoDb) Delete(userId string) (*bool, error) {
	filter := bson.M{
		"_id": userId,
	}
	deleteCount, err := repo.collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	if deleteCount.DeletedCount == 0 {
		return nil, errors.New("error_delete_from_database")
	}
	result := true
	return &result, nil
}
func (repo *MongoDb) GetById(userId string) (*portContractUser.Object, error) {
	filter := bson.M{
		"_id": userId,
	}
	var (
		resultQuery portContractUser.Collection
		result      portContractUser.Object
	)
	err := repo.collection.FindOne(context.TODO(), filter).Decode(&resultQuery)
	if err != nil {
		return nil, err
	}
	mapper.AutoMapper(resultQuery, &result)
	return &result, nil
}
