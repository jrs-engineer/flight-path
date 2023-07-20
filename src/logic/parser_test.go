package logic_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/C001-developer/flight-path/src/logic"
	"github.com/stretchr/testify/require"
)

func TestGetPathManual(t *testing.T) {
	testCases := []struct {
		name         string
		flights      [][2]string
		expectedPath [2]string
		expectedErr  error
	}{
		{
			"invalid path with self cycle",
			[][2]string{{"A", "B"}, {"C", "C"}},
			[2]string{},
			logic.ErrInvalidPathCycle,
		},
		{
			"invalid path with multi sources",
			[][2]string{{"A", "B"}, {"B", "C"}, {"C", "D"}, {"B", "E"}},
			[2]string{},
			logic.ErrInvalidPathMultiSources,
		},
		{
			"invalid path with multi targets",
			[][2]string{{"A", "B"}, {"B", "C"}, {"D", "C"}},
			[2]string{},
			logic.ErrInvalidPathMultiTargets,
		},
		{
			"invalid path with separated",
			[][2]string{{"A", "B"}, {"B", "C"}, {"E", "F"}, {"D", "E"}},
			[2]string{},
			logic.ErrInvalidPathSeparated,
		},
		{
			"invalid path with cycle",
			[][2]string{{"A", "B"}, {"D", "C"}, {"B", "D"}, {"E", "A"}, {"C", "E"}},
			[2]string{},
			logic.ErrInvalidPathCycle,
		},
		{
			"valid path 1",
			[][2]string{{"A", "B"}, {"D", "E"}, {"C", "D"}, {"B", "C"}, {"E", "F"}},
			[2]string{"A", "F"},
			nil,
		},
		{
			"valid path 2",
			[][2]string{{"A", "B"}, {"D", "E"}, {"B", "D"}, {"E", "F"}, {"F", "C"}},
			[2]string{"A", "C"},
			nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path, err := logic.GetSinglePath(tc.flights)
			require.Equal(t, tc.expectedPath, path)
			require.Equal(t, tc.expectedErr, err)
		})
	}
}

func TestGetPathRandom(t *testing.T) {
	for i := 0; i < 1000; i++ {
		airportCount := rand.Uint32()%(26-2) + 2
		pathCount := rand.Uint32()%airportCount + airportCount - 1
		flights := generateRandomFlights(pathCount, airportCount)
		p, err := logic.GetSinglePath(flights)
		if pathCount+1 == airportCount {
			require.NoError(t, err)
		} else {
			if err == nil {
				fmt.Printf("%d, %d\n", pathCount, airportCount)
				fmt.Printf("%v\n", flights)
				fmt.Printf("%v\n", p)
			}
			require.Error(t, err)
		}
	}
}

func generateRandomFlights(pathCount, airportCount uint32) [][2]string {
	flights := make([][2]string, pathCount)
	airports := make([]string, airportCount)
	for i := range airports {
		airports[i] = fmt.Sprintf("%c", 'A'+i)
	}

	if pathCount+1 == airportCount {
		indexes := rand.Perm(int(pathCount))
		for i := range indexes {
			flights[i] = [2]string{airports[indexes[i]], airports[indexes[i]+1]}
		}
	} else {
		for i := 0; i < int(pathCount); i++ {
			flights[i] = [2]string{airports[rand.Uint32()%airportCount], airports[rand.Uint32()%airportCount]}
		}
	}
	return flights
}
