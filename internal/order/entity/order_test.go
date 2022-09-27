package entity_test

import (
	"intensivo-go/internal/order/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateNewOrder_ThenShouldReceibeAnError(t *testing.T) {
	order := entity.Order{}

	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateNewOrder_ThenShouldReceibeAnError(t *testing.T) {
	order := entity.Order{ID: "123"}

	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenValidParams_WhenCallNewOrder_ThenShouldReceibeCreateOrderWithParams(t *testing.T) {
	order, error := entity.NewOrder("123", 10, 2)

	assert.NoError(t, error)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenValidParams_WhenCallFinalPrice_ThenShouldCalculateFinalFinalAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
