package handler_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/C001-developer/flight-path/src/handler"
	"github.com/C001-developer/flight-path/src/logic"
	"github.com/stretchr/testify/require"
)

func TestFlightPathHandler(t *testing.T) {
	testCases := []struct {
		name               string
		payload            []byte
		expectedStatusCode int
		expectedResponse   []byte
	}{
		{
			"invalid request body",
			[]byte(`{"invalid": "json"`),
			http.StatusBadRequest,
			[]byte(handler.ErrInvalidJSONData),
		},
		{
			"invalid separated path",
			[]byte(`[["A", "B"], ["B", "C"], ["D", "E"]]`),
			http.StatusBadRequest,
			[]byte(logic.ErrInvalidPathSeparated.Error()),
		},
		{
			"invalid path with cycle",
			[]byte(`[["A", "B"], ["B", "C"], ["C", "D"], ["D", "E"], ["E", "A"]]`),
			http.StatusBadRequest,
			[]byte(logic.ErrInvalidPathCycle.Error()),
		},
		{
			"invalid path with multi sources",
			[]byte(`[["A", "B"], ["B", "C"], ["C", "D"], ["B", "E"]]`),
			http.StatusBadRequest,
			[]byte(logic.ErrInvalidPathMultiSources.Error()),
		},
		{
			"invalid path with multi targets",
			[]byte(`[["A", "B"], ["B", "C"], ["D", "C"]]`),
			http.StatusBadRequest,
			[]byte(logic.ErrInvalidPathMultiTargets.Error()),
		},
		{
			"valid path",
			[]byte(`[["A", "B"], ["D", "E"], ["C", "D"], ["B", "C"], ["E", "F"]]`),
			http.StatusOK,
			[]byte(`["A","F"]`),
		},
		{
			"valid path",
			[]byte(`[["SFO", "EWR"]]`),
			http.StatusOK,
			[]byte(`["SFO","EWR"]`),
		},
		{
			"valid path",
			[]byte(`[["ATL", "EWR"], ["SFO", "ATL"]]`),
			http.StatusOK,
			[]byte(`["SFO","EWR"]`),
		},
		{
			"valid path",
			[]byte(`[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]`),
			http.StatusOK,
			[]byte(`["SFO","EWR"]`),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/calculate", bytes.NewBuffer(tc.payload))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			handler.FlightPathHandler(rr, req)

			require.Equal(t, tc.expectedStatusCode, rr.Code)
			require.Equal(t, tc.expectedResponse, bytes.Trim(rr.Body.Bytes(), "\n"))
		})
	}
}
