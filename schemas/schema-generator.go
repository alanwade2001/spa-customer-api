package types

//go:generate mkdir -p ../models/generated
//go:generate schema-generate -i ./customermodel-schema-v1.0.json -o ../models/generated/CustomerModel.go -p generated

//go:generate mkdir -p ../types/generated
//go:generate schema-generate -i ./customer-schema-v1.0.json -o ../types/generated/Customer.go -p generated
