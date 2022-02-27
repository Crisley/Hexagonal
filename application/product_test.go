package application_test

import (
	"testing"

	"github.com/crisley/go-hexagonal/application"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}

	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be grater then zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T){
	product := application.Product{
		Name: "Hello",
		Status: application.ENABLED,
		Price: 0,
	}

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T){
	product := application.Product{
		ID: uuid.NewV4().String(),
		Name: "hello",
		Status: application.DISABLED,
		Price: 10,		
	}

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disable", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())

}