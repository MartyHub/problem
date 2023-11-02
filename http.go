package problem

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const ContentTypeApplicationProblemJSON = "application/problem+json"

// Write writes the given problem details to the given HTTP response writer.
func Write(writer http.ResponseWriter, pb *Details) error {
	writer.Header().Set("Content-Type", ContentTypeApplicationProblemJSON)
	writer.WriteHeader(pb.Status)

	if err := json.NewEncoder(writer).Encode(pb); err != nil {
		return fmt.Errorf("%w: %s", ErrJSON, err.Error())
	}

	return nil
}
