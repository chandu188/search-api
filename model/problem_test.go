package model

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem(t *testing.T) {
	p := &Problem{
		Title:  "service unavailable",
		Status: 500,
		Detail: "error while connecting to the database",
	}
	_, err := json.Marshal(p)
	assert.Nil(t, err, "should be able marshal a problem to json")
}

func TestProblemUnMarshal(t *testing.T) {
	probJSON := `{
			"title" : "service unavailable",
			"status": 500
		}`
	var prob Problem
	err := json.Unmarshal([]byte(probJSON), &prob)
	assert.Nil(t, err)
	assert.Equal(t, 500, prob.Status)

	probJSON = `{
		"title" : "service unavailable",
		"status": "500"
	}`
	err = json.Unmarshal([]byte(probJSON), &prob)
	assert.NotNil(t, err, "should fail when invalid problem json is parsed")
}
