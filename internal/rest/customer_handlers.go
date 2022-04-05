package rest

import (
	"github.com/ashtishad/banking/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

// GetCustomerByID is a handler function to get customer by id
func (ch *CustomerHandlers) GetCustomerByID(c *gin.Context) {
	log.Println("Handling GET request on ... /customers/{id}")

	id := c.Param("id")

	customer, err := ch.Service.GetCustomerById(id)
	if err != nil {
		c.JSON(err.AsStatus(), err)
		return
	}

	c.JSON(http.StatusOK, customer)
}