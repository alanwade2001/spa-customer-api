package repositories

import (
	"errors"
	"reflect"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"k8s.io/klog/v2"

	"github.com/alanwade2001/spa-customer-api/models/generated"
	"github.com/alanwade2001/spa-customer-api/types"

	mgo "github.com/alanwade2001/spa-common/mongo"
)

// MongoRepository s
type MongoRepository struct {
	service *mgo.MongoService
}

// NewMongoService s
func NewMongoService() types.RepositoryAPI {

	return &MongoRepository{}
}

func (ms *MongoRepository) GetService() *mgo.MongoService {

	if ms.service != nil {
		return ms.service
	}

	uriTemplate := viper.GetString("MONGODB_URI_TEMPLATE")
	username := viper.GetString("MONGODB_USER")
	password := viper.GetString("MONGODB_PASSWORD")
	connectTimeout := viper.GetDuration("MONGODB_TIMEOUT") * time.Second
	database := viper.GetString("MONGODB_DATABASE")
	collection := viper.GetString("MONGODB_COLLECTION")

	structcodec, _ := bsoncodec.NewStructCodec(bsoncodec.JSONFallbackStructTagParser)
	reg := bson.NewRegistryBuilder().
		RegisterTypeEncoder(reflect.TypeOf(generated.CustomerModel{}), structcodec).
		RegisterTypeDecoder(reflect.TypeOf(generated.CustomerModel{}), structcodec).
		Build()

	service := mgo.NewMongoService(uriTemplate, username, password, database, collection, connectTimeout, reg)

	ms.service = service

	return ms.service
}

// CreateCustomer f
func (ms MongoRepository) CreateCustomer(customer *generated.CustomerModel) (*generated.CustomerModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	customer.Id = primitive.NewObjectID().Hex()

	result, err := ms.GetService().GetCollection(connection).InsertOne(connection.Ctx, customer)

	if err != nil {
		klog.Warningf("Could not create Customer: %v", err)
		return nil, err
	}

	klog.Infof("result:[%+v]", result)
	klog.Infof("cus:[%+v]", customer)

	return customer, nil
}

// GetCustomer f
func (ms MongoRepository) GetCustomer(ID string) (*generated.CustomerModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	customer := new(generated.CustomerModel)
	filter := bson.M{"_id": ID}

	if err := ms.GetService().GetCollection(connection).FindOne(connection.Ctx, filter).Decode(customer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			klog.InfoS("Customer not found:", "ID", ID)
			return nil, nil
		}

		klog.Warningf("Error calling - FindOne: %v", err)
		return nil, err
	}

	klog.Infof("found customer:[%+v]", customer)

	return customer, nil
}

// GetCustomers f
func (ms MongoRepository) GetCustomers() (*[]generated.CustomerModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	var cursor *mongo.Cursor
	var err error
	customers := []generated.CustomerModel{}

	filter := bson.M{}
	if cursor, err = ms.GetService().GetCollection(connection).Find(connection.Ctx, filter); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			klog.InfoS("No customers found")
			return &customers, nil
		}

		klog.Warningf("Error calling - Find : %v", err)
		return nil, err
	}

	if err = cursor.All(connection.Ctx, &customers); err != nil {
		return nil, err
	}

	return &customers, nil
}

// FindCustomerByEmail f
func (ms MongoRepository) FindCustomerByEmail(email string) (*generated.CustomerModel, error) {
	connection := ms.GetService().Connect()
	defer connection.Disconnect()

	customer := new(generated.CustomerModel)
	filter := bson.M{"Users.email": email}

	if err := ms.GetService().GetCollection(connection).FindOne(connection.Ctx, filter).Decode(customer); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			klog.InfoS("Customer with email not found:", "Email", email)
			return nil, nil
		}

		klog.Warningf("Error calling - FindOne with filter : %v", err)
		return nil, err
	}

	klog.InfoS("customer:[%+v]", customer)

	return customer, nil
}
