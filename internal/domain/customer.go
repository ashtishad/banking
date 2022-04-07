package domain

import (
	"github.com/ashtishad/banking/internal/dto"
	"github.com/ashtishad/banking/pkg/lib"
	"time"
)

// Customer is a domain entity
// Id is int64 matched pg big serial
// DateOfBirth is time.Time and is matched pg timestamp

type Customer struct {
	Id          int64     `json:"id"`
	Name        string    `json:"full_name" binding:"required,min=3,max=50,alpha"`
	City        string    `json:"city" binding:"required,min=3,max=50,alpha"`
	Zipcode     string    `json:"zipcode" binding:"required,min=3,max=50,number"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Status      int8      `json:"status"`
}

// CustomerRepository is a SECONDARY PORT on Hexagonal architecture
type CustomerRepository interface {
	FindById(id int64) (*Customer, lib.RestErr)
	FindByStatus(status int8) ([]Customer, lib.RestErr)
	Create(customer *Customer) (*Customer, lib.RestErr)

	//FindByName(name string) (*Customer, lib.RestErr)
	//Update(customer *Customer) (*Customer, lib.RestErr)
	//Delete(id int64) lib.RestErr
}

// statusAsText returns 1= active, 0= inactive
func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == 0 {
		statusAsText = "inactive"
	}
	return statusAsText
}

// timeAsText returns date in format yyyy-mm-dd
func (c Customer) dateAsText() string {
	return c.DateOfBirth.Format(lib.DbDateLayout)
}

// ToCustomerResponse  converts Customer domain entity to DTO response
// there is some type conversion happening when I show it to user in response
// Like: Customer.DateOfBirth is time.Time, However on user level its string
func (c Customer) ToCustomerResponse() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.dateAsText(),
		Status:      c.statusAsText(),
	}
}

//// ValidateNewCustomerRequest validates Customer domain entity
//func ValidateNewCustomerRequest(c dto.NewCustomerRequest) lib.RestErr {
//	rName := regexp.MustCompile(lib.NameRegex.Pattern)
//	rCity := regexp.MustCompile(lib.CityRegex.Pattern)
//	rZipcode := regexp.MustCompile(lib.ZipRegex.Pattern)
//	rDOB := regexp.MustCompile(lib.DOBRegex.Pattern)
//
//	// validate name
//
//	if !rName.MatchString(c.Name) {
//		return lib.NewBadRequestError(lib.NameRegex.Error)
//	}
//	if !rCity.MatchString(c.City) {
//		return lib.NewBadRequestError("city is invalid")
//	}
//	if !rZipcode.MatchString(c.Zipcode) {
//		return lib.NewBadRequestError("zipcode is invalid")
//	}
//	if !rDOB.MatchString(c.DateOfBirth) {
//		return lib.NewBadRequestError("date of birth is invalid")
//	}
//
//	return nil
//}

func ToNewCustomer(cr dto.NewCustomerRequest) *Customer {
	return &Customer{
		Name:        cr.Name,
		City:        cr.City,
		Zipcode:     cr.Zipcode,
		DateOfBirth: lib.DateAsTime(cr.DateOfBirth),
		Status:      1,
	}
}
