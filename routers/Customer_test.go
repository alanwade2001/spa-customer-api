package routers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/alanwade2001/spa-customer-api/models/generated"
	"github.com/alanwade2001/spa-customer-api/repositories"
	"github.com/alanwade2001/spa-customer-api/routers"
	"github.com/alanwade2001/spa-customer-api/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCustomerRouter_GetCustomers(t *testing.T) {

	// change the database to be unittest
	os.Setenv("MONGODB_DATABASE", "unittest")

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	repositoryAPI := repositories.NewMongoService()
	serviceAPI := services.NewService(repositoryAPI)
	customerAPI := routers.NewCustomerRouter(serviceAPI)
	services.NewConfigService().Load("..")

	engine.GET("/customers", customerAPI.GetCustomers)

	type data struct {
		customers []interface{}
	}

	tests := []struct {
		name string
		data data
	}{
		{
			name: "Test01",
			data: data{
				customers: []interface{}{},
			},
		},
		{
			name: "Test02",
			data: data{
				customers: []interface{}{
					generated.CustomerModel{
						Id:     "1122334455667788",
						Active: true,
						Name:   "Corporation ABC",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112233",
								Name: "Payments",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXX",
										IBAN: "IE12AIBKIE2D90909012345678",
										Name: "AJW",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan.test.ie",
							},
						},
					},
				},
			},
		},
		{
			name: "Test03",
			data: data{
				customers: []interface{}{
					generated.CustomerModel{
						Id:     "1122334455667788",
						Active: true,
						Name:   "Corporation ABC",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112233",
								Name: "Payments",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXX",
										IBAN: "IE12AIBKIE2D90909012345678",
										Name: "AJW",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan.test.ie",
							},
						},
					},
					generated.CustomerModel{
						Id:     "1122334455667789",
						Active: true,
						Name:   "Corporation ABD",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112234",
								Name: "Payments2",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXY",
										IBAN: "IE12AIBKIE2D90909012345679",
										Name: "AJX",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan2.test.ie",
							},
						},
					},
					generated.CustomerModel{
						Id:     "1122334455667790",
						Active: true,
						Name:   "Corporation ABE",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112235",
								Name: "Payments3",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXZ",
										IBAN: "IE12AIBKIE2D90909012345680",
										Name: "AJY",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan3.test.ie",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// clear down the database
			mongoRepo := repositoryAPI.(*repositories.MongoRepository)
			conn := mongoRepo.GetService().Connect()
			defer conn.Disconnect()

			// clear the database
			filter := bson.M{}
			if _, err := mongoRepo.GetService().GetCollection(conn).DeleteMany(conn.Ctx, filter); err != nil {
				t.Logf("error deleting customers [%s]", err.Error())
			}

			// insert the seed data
			if _, err := mongoRepo.GetService().GetCollection(conn).InsertMany(conn.Ctx, tt.data.customers); err != nil {
				t.Logf("error deleting customers [%s]", err.Error())
			}

			req, err := http.NewRequest("GET", "/customers", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			engine.ServeHTTP(w, req)
			//fmt.Println(w.Body)

			// Check to see if the response was what you expected
			if w.Code == http.StatusOK {
				t.Logf("Expected to get status %d is same ast %d\n", http.StatusOK, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
			}

			var result []generated.CustomerModel
			if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
				t.Fatalf("Expected Json output")
			}

			// check if the size of the array is equal to the data
			if len(result) == len(tt.data.customers) {
				t.Logf("Expected to get number of customers %d is same as %d\n", len(tt.data.customers), len(result))
			} else {
				t.Fatalf("Expected to get number of customers %d but instead got %d\n", len(tt.data.customers), len(result))
			}

			// random check on the first customer of it exists
			if len(result) > 0 {
				first := tt.data.customers[0].(generated.CustomerModel)
				if result[0].Id == first.Id {
					t.Logf("Expected id to match %s is same as %s\n", first.Id, result[0].Id)
				} else {
					t.Fatalf("Expected id to match %s but instead got %s\n", first.Id, result[0].Id)
				}
			} else {
				t.Logf("No customers in the result")
			}
		})
	}
}

func TestCustomerRouter_GetCustomer(t *testing.T) {

	// change the database to be unittest
	os.Setenv("MONGODB_DATABASE", "unittest")

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	repositoryAPI := repositories.NewMongoService()
	serviceAPI := services.NewService(repositoryAPI)
	customerAPI := routers.NewCustomerRouter(serviceAPI)
	services.NewConfigService().Load("..")

	engine.GET("/customers/:id", customerAPI.GetCustomer)

	type data struct {
		customers []interface{}
	}

	tests := []struct {
		name  string
		id    string
		data  data
		code  int
		index int
	}{
		{
			name: "Test01",
			id:   "11223344",
			data: data{
				customers: []interface{}{},
			},
			code:  http.StatusNotFound,
			index: -1,
		},
		{
			name: "Test02",
			id:   "1122334455667788",
			data: data{
				customers: []interface{}{
					generated.CustomerModel{
						Id:     "1122334455667788",
						Active: true,
						Name:   "Corporation ABC",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112233",
								Name: "Payments",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXX",
										IBAN: "IE12AIBKIE2D90909012345678",
										Name: "AJW",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan.test.ie",
							},
						},
					},
				},
			},
			code:  http.StatusOK,
			index: 0,
		},
		{
			name: "Test03",
			id:   "1122334455667789",
			data: data{
				customers: []interface{}{
					generated.CustomerModel{
						Id:     "1122334455667788",
						Active: true,
						Name:   "Corporation ABC",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112233",
								Name: "Payments",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXX",
										IBAN: "IE12AIBKIE2D90909012345678",
										Name: "AJW",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan.test.ie",
							},
						},
					},
					generated.CustomerModel{
						Id:     "1122334455667789",
						Active: true,
						Name:   "Corporation ABD",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112234",
								Name: "Payments2",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXY",
										IBAN: "IE12AIBKIE2D90909012345679",
										Name: "AJX",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan2.test.ie",
							},
						},
					},
					generated.CustomerModel{
						Id:     "1122334455667790",
						Active: true,
						Name:   "Corporation ABE",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112235",
								Name: "Payments3",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXZ",
										IBAN: "IE12AIBKIE2D90909012345680",
										Name: "AJY",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan3.test.ie",
							},
						},
					},
				},
			},
			code:  http.StatusOK,
			index: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// clear down the database
			mongoRepo := repositoryAPI.(*repositories.MongoRepository)
			conn := mongoRepo.GetService().Connect()
			defer conn.Disconnect()

			// clear the database
			filter := bson.M{}
			if _, err := mongoRepo.GetService().GetCollection(conn).DeleteMany(conn.Ctx, filter); err != nil {
				t.Logf("error deleting customers [%s]", err.Error())
			}

			// insert the seed data
			if len(tt.data.customers) > 0 {
				if _, err := mongoRepo.GetService().GetCollection(conn).InsertMany(conn.Ctx, tt.data.customers); err != nil {
					t.Logf("error inserting customers [%s]", err.Error())
				}
			}

			req, err := http.NewRequest("GET", "/customers/"+tt.id, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			engine.ServeHTTP(w, req)
			//fmt.Println(w.Body)

			// Check to see if the response was what you expected
			if w.Code == tt.code {
				t.Logf("Expected to get status %d is same ast %d\n", tt.code, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.code, w.Code)
			}

			// we have a result
			if w.Code == http.StatusOK {
				var actual generated.CustomerModel
				if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
					t.Fatalf("Expected Json output")
				}

				// random check on the first customer of it exists
				expected := tt.data.customers[tt.index].(generated.CustomerModel)
				if actual.Id == expected.Id {
					t.Logf("Expected id to match %s is same as %s\n", expected.Id, actual.Id)
				} else {
					t.Fatalf("Expected id to match %s but instead got %s\n", expected.Id, actual.Id)
				}
			} else {
				t.Logf("no result found\n")
			}
		})
	}
}

func TestCustomerRouter_CreateCustomer(t *testing.T) {

	// change the database to be unittest
	os.Setenv("MONGODB_DATABASE", "unittest")

	gin.SetMode(gin.TestMode)
	engine := gin.Default()
	repositoryAPI := repositories.NewMongoService()
	serviceAPI := services.NewService(repositoryAPI)
	customerAPI := routers.NewCustomerRouter(serviceAPI)
	services.NewConfigService().Load("..")

	engine.POST("/customers", customerAPI.CreateCustomer)

	type data struct {
		customers   []interface{}
		newCustomer generated.CustomerModel
	}

	tests := []struct {
		name     string
		data     data
		code     int
		location string
	}{
		{
			name: "Test01",
			data: data{
				customers: []interface{}{},
				newCustomer: generated.CustomerModel{
					Active: true,
					Name:   "Corporation ABC",
					InitiatingParties: []*generated.InitiatingParty{
						{
							Id:   "112233",
							Name: "Payments",
							RegisteredAccounts: []*generated.AccountReference{
								{
									BIC:  "AIBK2DXXX",
									IBAN: "IE12AIBKIE2D90909012345678",
									Name: "AJW",
								},
							},
						},
					},
					Users: []*generated.UserReference{
						{
							Email: "alan.test.ie",
						},
					},
				},
			},
			code:     http.StatusCreated,
			location: "/customers/.+",
		},
		{
			name: "Test02",
			data: data{
				customers: []interface{}{
					generated.CustomerModel{
						Id:     "1122334455667788",
						Active: true,
						Name:   "Corporation ABC",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112233",
								Name: "Payments",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXX",
										IBAN: "IE12AIBKIE2D90909012345678",
										Name: "AJW",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan.test.ie",
							},
						},
					},
				},
				newCustomer: generated.CustomerModel{
					Active: true,
					Name:   "Corporation ABC",
					InitiatingParties: []*generated.InitiatingParty{
						{
							Id:   "112233",
							Name: "Payments",
							RegisteredAccounts: []*generated.AccountReference{
								{
									BIC:  "AIBK2DXXX",
									IBAN: "IE12AIBKIE2D90909012345678",
									Name: "AJW",
								},
							},
						},
					},
					Users: []*generated.UserReference{
						{
							Email: "alan.test.ie",
						},
					},
				},
			},
			code:     http.StatusCreated,
			location: "/customers/.+",
		},
		{
			name: "Test03",
			data: data{
				customers: []interface{}{
					generated.CustomerModel{
						Id:     "1122334455667788",
						Active: true,
						Name:   "Corporation ABC",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112233",
								Name: "Payments",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXX",
										IBAN: "IE12AIBKIE2D90909012345678",
										Name: "AJW",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan.test.ie",
							},
						},
					},
					generated.CustomerModel{
						Id:     "1122334455667789",
						Active: true,
						Name:   "Corporation ABD",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112234",
								Name: "Payments2",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXY",
										IBAN: "IE12AIBKIE2D90909012345679",
										Name: "AJX",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan2.test.ie",
							},
						},
					},
					generated.CustomerModel{
						Id:     "1122334455667790",
						Active: true,
						Name:   "Corporation ABE",
						InitiatingParties: []*generated.InitiatingParty{
							{
								Id:   "112235",
								Name: "Payments3",
								RegisteredAccounts: []*generated.AccountReference{
									{
										BIC:  "AIBK2DXXZ",
										IBAN: "IE12AIBKIE2D90909012345680",
										Name: "AJY",
									},
								},
							},
						},
						Users: []*generated.UserReference{
							{
								Email: "alan3.test.ie",
							},
						},
					},
				},
			},
			code:     http.StatusCreated,
			location: "/customers/.+",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// clear down the database
			mongoRepo := repositoryAPI.(*repositories.MongoRepository)
			conn := mongoRepo.GetService().Connect()
			defer conn.Disconnect()

			// clear the database
			filter := bson.M{}
			if _, err := mongoRepo.GetService().GetCollection(conn).DeleteMany(conn.Ctx, filter); err != nil {
				t.Logf("error deleting customers [%s]", err.Error())
			}

			if len(tt.data.customers) > 0 {
				// insert the seed data
				if _, err := mongoRepo.GetService().GetCollection(conn).InsertMany(conn.Ctx, tt.data.customers); err != nil {
					t.Logf("error inserting customers [%s]", err.Error())
				}
			}

			body, err := json.Marshal(tt.data.newCustomer)
			if err != nil {
				t.Fatalf("failed to marshal the new customer\n")
			}

			req, err := http.NewRequest(http.MethodPost, "/customers", bytes.NewReader(body))
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder so you can inspect the response
			w := httptest.NewRecorder()

			// Perform the request
			engine.ServeHTTP(w, req)
			//fmt.Println(w.Body)

			// Check to see if the response was what you expected
			if w.Code == tt.code {
				t.Logf("Expected to get status %d is same ast %d\n", tt.code, w.Code)
			} else {
				t.Fatalf("Expected to get status %d but instead got %d\n", tt.code, w.Code)
			}
		})
	}
}
