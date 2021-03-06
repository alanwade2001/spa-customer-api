package main

// Customer s
type Customer struct {
	ID                string `bson:"_id"`
	Name              string
	InitiatingPartyID string
	Users             Users
}

// Customers a
type Customers []Customer

// User s
type User struct {
	Email string
}

// Users a
type Users []User

// Roles s
type Roles struct {
	Submitters Users
	Approvers  Users
	Admins     Users
}
