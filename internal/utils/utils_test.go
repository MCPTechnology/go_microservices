package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type sampleDto struct {
	Name string `json:"name" validate:"required,min=3"`
	Age  int    `json:"age" validate:"required,gte=0"`
}

func TestToJson_SuccessfullyUnmarshal(t *testing.T) {
	tests := []struct {
		name   string
		target sampleDto
	}{
		{
			name: "Successfully unmarshal json",
			target: sampleDto{
				Name: "Name",
				Age:  10,
			},
		},
	}

	rw := httptest.NewRecorder()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ToJson(rw, tc.target)
			assert.Nil(t, err)

			got := &sampleDto{}
			err = FromJson(rw.Body, got)
			assert.Nil(t, err)
			assert.EqualValues(t, *got, tc.target)
		})
	}
}
