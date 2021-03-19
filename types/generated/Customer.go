// Code generated by schema-generate. DO NOT EDIT.

package generated

import (
    "bytes"
    "encoding/json"
    "errors"
)

// AccountReference A reference to an account
type AccountReference struct {

  // BIC
  BIC string `json:"BIC"`

  // IBAN
  IBAN string `json:"IBAN"`

  // Name on the account
  Name string `json:"Name"`
}

// Customer SPA Customer
type Customer struct {

  // The unique identifier for a customer
  CustomerId string `json:"CustomerId"`

  // The name of the customer
  CustomerName string `json:"CustomerName"`

  // Initiating Parties a customer has
  InitiatingParties []*InitiatingParty `json:"InitiatingParties"`
}

// InitiatingParty An initating party who can initate a payment
type InitiatingParty struct {

  // unique identifier of an initiating party
  Id string `json:"Id"`

  // name of the initiating party
  Name string `json:"Name"`

  // Accounts that the initiating party can use
  RegisteredAccounts []*AccountReference `json:"RegisteredAccounts"`
}

func (strct *AccountReference) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "BIC" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "BIC" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"BIC\": ")
	if tmp, err := json.Marshal(strct.BIC); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "IBAN" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "IBAN" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"IBAN\": ")
	if tmp, err := json.Marshal(strct.IBAN); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Name" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *AccountReference) UnmarshalJSON(b []byte) error {
    BICReceived := false
    IBANReceived := false
    NameReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "BIC":
            if err := json.Unmarshal([]byte(v), &strct.BIC); err != nil {
                return err
             }
            BICReceived = true
        case "IBAN":
            if err := json.Unmarshal([]byte(v), &strct.IBAN); err != nil {
                return err
             }
            IBANReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        }
    }
    // check if BIC (a required property) was received
    if !BICReceived {
        return errors.New("\"BIC\" is required but was not present")
    }
    // check if IBAN (a required property) was received
    if !IBANReceived {
        return errors.New("\"IBAN\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    return nil
}

func (strct *Customer) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "CustomerId" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "CustomerId" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CustomerId\": ")
	if tmp, err := json.Marshal(strct.CustomerId); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "CustomerName" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "CustomerName" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"CustomerName\": ")
	if tmp, err := json.Marshal(strct.CustomerName); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "InitiatingParties" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "InitiatingParties" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"InitiatingParties\": ")
	if tmp, err := json.Marshal(strct.InitiatingParties); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *Customer) UnmarshalJSON(b []byte) error {
    CustomerIdReceived := false
    CustomerNameReceived := false
    InitiatingPartiesReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "CustomerId":
            if err := json.Unmarshal([]byte(v), &strct.CustomerId); err != nil {
                return err
             }
            CustomerIdReceived = true
        case "CustomerName":
            if err := json.Unmarshal([]byte(v), &strct.CustomerName); err != nil {
                return err
             }
            CustomerNameReceived = true
        case "InitiatingParties":
            if err := json.Unmarshal([]byte(v), &strct.InitiatingParties); err != nil {
                return err
             }
            InitiatingPartiesReceived = true
        }
    }
    // check if CustomerId (a required property) was received
    if !CustomerIdReceived {
        return errors.New("\"CustomerId\" is required but was not present")
    }
    // check if CustomerName (a required property) was received
    if !CustomerNameReceived {
        return errors.New("\"CustomerName\" is required but was not present")
    }
    // check if InitiatingParties (a required property) was received
    if !InitiatingPartiesReceived {
        return errors.New("\"InitiatingParties\" is required but was not present")
    }
    return nil
}

func (strct *InitiatingParty) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")
    comma := false
    // "Id" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Id" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Id\": ")
	if tmp, err := json.Marshal(strct.Id); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "Name" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "Name" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"Name\": ")
	if tmp, err := json.Marshal(strct.Name); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true
    // "RegisteredAccounts" field is required
    // only required object types supported for marshal checking (for now)
    // Marshal the "RegisteredAccounts" field
    if comma { 
        buf.WriteString(",") 
    }
    buf.WriteString("\"RegisteredAccounts\": ")
	if tmp, err := json.Marshal(strct.RegisteredAccounts); err != nil {
		return nil, err
 	} else {
 		buf.Write(tmp)
	}
	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()
	return rv, nil
}

func (strct *InitiatingParty) UnmarshalJSON(b []byte) error {
    IdReceived := false
    NameReceived := false
    RegisteredAccountsReceived := false
    var jsonMap map[string]json.RawMessage
    if err := json.Unmarshal(b, &jsonMap); err != nil {
        return err
    }
    // parse all the defined properties
    for k, v := range jsonMap {
        switch k {
        case "Id":
            if err := json.Unmarshal([]byte(v), &strct.Id); err != nil {
                return err
             }
            IdReceived = true
        case "Name":
            if err := json.Unmarshal([]byte(v), &strct.Name); err != nil {
                return err
             }
            NameReceived = true
        case "RegisteredAccounts":
            if err := json.Unmarshal([]byte(v), &strct.RegisteredAccounts); err != nil {
                return err
             }
            RegisteredAccountsReceived = true
        }
    }
    // check if Id (a required property) was received
    if !IdReceived {
        return errors.New("\"Id\" is required but was not present")
    }
    // check if Name (a required property) was received
    if !NameReceived {
        return errors.New("\"Name\" is required but was not present")
    }
    // check if RegisteredAccounts (a required property) was received
    if !RegisteredAccountsReceived {
        return errors.New("\"RegisteredAccounts\" is required but was not present")
    }
    return nil
}