package logic

import "errors"

var (
	// ErrInvalidPathCycle is an error that is returned when the input path contains a cycle.
	ErrInvalidPathCycle = errors.New("invalid path, cycle")
	// ErrInvalidPathSeparated is an error that is returned when the input path is separated.
	ErrInvalidPathSeparated = errors.New("invalid path, separated")
	// ErrInvalidPathMultiSources is an error that is returned when the input path contains multiple sources.
	ErrInvalidPathMultiSources = errors.New("invalid path, multi sources")
	// ErrInvalidPathMultiTargets is an error that is returned when the input path contains multiple targets.
	ErrInvalidPathMultiTargets = errors.New("invalid path, multi targets")
)

// GetSinglePath is a function that takes a slice of flight paths and returns a single path.
func GetSinglePath(flights [][2]string) ([2]string, error) {
	var source, target string
	flightMap := make(map[string][2]int32)
	// Check if the input is valid
	for _, flight := range flights {
		source, target = flight[0], flight[1]
		if source == target {
			return [2]string{}, ErrInvalidPathCycle
		}
		if d, ok := flightMap[source]; ok {
			// Check if the path contains a multiple sources
			if d[0] == 1 {
				return [2]string{}, ErrInvalidPathMultiSources
			}
			flightMap[source] = [2]int32{d[0] + 1, d[1]}
		} else {
			flightMap[flight[0]] = [2]int32{1, 0}
		}
		if d, ok := flightMap[target]; ok {
			// Check if the path contains a multiple targets
			if d[1] == 1 {
				return [2]string{}, ErrInvalidPathMultiTargets
			}
			flightMap[target] = [2]int32{d[0], d[1] + 1}
		} else {
			flightMap[flight[1]] = [2]int32{0, 1}
		}
	}

	// it ensures that each airport is visited exactly once,
	// in other words, each airport has at most one incoming and one outgoing flight.

	source, target = "", ""
	// Find the source and target
	for k, v := range flightMap {
		if v[1] == 0 {
			if source != "" {
				return [2]string{}, ErrInvalidPathSeparated
			}
			source = k
		}
		if v[0] == 0 {
			if target != "" {
				return [2]string{}, ErrInvalidPathSeparated
			}
			target = k
		}
	}
	if source == "" || target == "" {
		return [2]string{}, ErrInvalidPathCycle
	}

	return [2]string{source, target}, nil
}
