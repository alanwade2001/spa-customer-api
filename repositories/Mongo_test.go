package repositories

import (
	"reflect"
	"testing"

	"github.com/alanwade2001/spa-customer-api/models/generated"
	"github.com/alanwade2001/spa-customer-api/services"
	"github.com/stretchr/testify/assert"
)

func TestMongoService_CreateCustomer(t *testing.T) {

	services.NewConfigService().Load("..")
	ms := NewMongoService()

	type args struct {
		customer *generated.CustomerModel
	}
	tests := []struct {
		name    string
		ms      MongoRepository
		args    args
		want    *generated.CustomerModel
		wantErr bool
	}{
		{
			name: "test001",
			ms:   *ms.(*MongoRepository),
			args: args{
				customer: &generated.CustomerModel{
					Active: true,
					Name:   "CorporationABC",
					Users: []*generated.UserReference{
						{
							Email: "test001@test.ie",
						},
						{
							Email: "test001@test1.ie",
						},
					},
					InitiatingParties: []*generated.InitiatingParty{
						{
							Id:   "1234",
							Name: "Accounts Dept",
							RegisteredAccounts: []*generated.AccountReference{
								{
									BIC:  "AIBKIE2D",
									IBAN: "IE12AIBKIE90909012345678",
									Name: "Accounts Dept",
								},
							},
						},
					},
				},
			},
			want: &generated.CustomerModel{Active: true,
				Name: "CorporationABC",
				Users: []*generated.UserReference{
					{
						Email: "test001@test.ie",
					},
					{
						Email: "test001@test1.ie",
					},
				},
				InitiatingParties: []*generated.InitiatingParty{
					{
						Id:   "1234",
						Name: "Accounts Dept",
						RegisteredAccounts: []*generated.AccountReference{
							{
								BIC:  "AIBKIE2D",
								IBAN: "IE12AIBKIE90909012345678",
								Name: "Accounts Dept",
							},
						},
					},
				}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ms := MongoRepository{}
			got, err := ms.CreateCustomer(tt.args.customer)
			if (err != nil) != tt.wantErr {
				t.Errorf("MongoService.CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.NotNil(t, got.Id, "id is nil")

			// assign the id to the args.customer to pass test
			tt.want.Id = got.Id

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MongoService.CreateCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}
