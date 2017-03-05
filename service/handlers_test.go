package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFullfillmentStatusReturns200ForExistingSKU(test *testing.T) {
	var (
		request *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()
	targetSKU := "THINGAMAJIG12"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/skus/"+targetSKU, nil)
	server.ServeHTTP(recorder, request)

	var detail fulfillmentStatus

	if recorder.Code != http.StatusOK {
		test.Errorf("Expected %v; received %v", http.StatusOK, recorder.Code)
	}

	payload, err := ioutil.ReadAll(recorder.Body)
	if err != nil {
		test.Errorf("Error parsing response body: %v", err)
	}

	err = json.Unmarshal(payload, &detail)
	if err != nil {
		test.Errorf("Error unmarshaling response to fullfillment status: %v", err)
	}

	if detail.QuantityInStock != 100 {
		test.Errorf("Expected 100 qty in stock, got %d", detail.QuantityInStock)
	}
	if detail.ShipsWithin != 14 {
		test.Errorf("Expected shipswithin 14 days, got %d", detail.ShipsWithin)
	}
	if detail.SKU != "THINGAMAJIG12" {
		test.Errorf("Expected SKU THINGAMAJIG12, got %s", detail.SKU)
	}
}