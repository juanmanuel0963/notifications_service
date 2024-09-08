package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solution(t *testing.T) {
	type args struct {
		user   string
		typeId int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"Test 1",
			args{user: "user1", typeId: 1},
			nil,
		},
		{
			"Test 2",
			args{user: "user1", typeId: 1},
			nil,
		},
		{
			"Test 3",
			args{user: "user1", typeId: 1},
			errors.New("expected error"), // Placeholder for a non-nil error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NotificationsProcessing(tt.args.user, tt.args.typeId)

			if tt.want == nil {
				// Check if we expected no error
				assert.Equal(t, tt.want, got, "expected no error, but got one")
			} else {
				// Check if we expected an error (got should not be nil)
				assert.NotNil(t, got, "expected an error but got none")
			}
		})
	}
}

func Test_solution_1(t *testing.T) {

	got := NotificationsProcessing("user1", 1)
	assert.Equal(t, got, nil)

	got2 := NotificationsProcessing("user1", 1)
	assert.Equal(t, got2, nil)

	got3 := NotificationsProcessing("user1", 1)
	assert.NotNil(t, got3)
}

func Test_solution_2(t *testing.T) {

	got := NotificationsProcessing("user1", 2)
	assert.Equal(t, got, nil)

	got1 := NotificationsProcessing("user1", 2)
	assert.Equal(t, got1, nil)

	got2 := NotificationsProcessing("user1", 2)
	assert.Equal(t, got2, nil)

	got3 := NotificationsProcessing("user1", 2)
	assert.NotNil(t, got3)
}

func Test_solution_3(t *testing.T) {

	got := NotificationsProcessing("user1", 3)
	assert.Equal(t, got, nil)

	got2 := NotificationsProcessing("user1", 3)
	assert.NotNil(t, got2)
}
