package main

import (
	"errors"
	"testing"

	"github.com/juanmanuel0963/notification_service/m/notifications"
	"github.com/stretchr/testify/assert"
)

// Type 1, only 2 messages per minute per user
func Test_solution(t *testing.T) {
	type args struct {
		user    string
		typeId  int
		message string
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			"Test 1",
			args{user: "user1", typeId: 1, message: "message"},
			nil,
		},
		{
			"Test 2",
			args{user: "user1", typeId: 1, message: "message"},
			nil,
		},
		{
			"Test 3",
			args{user: "user1", typeId: 1, message: "message"},
			errors.New("expected error"), // Placeholder for a non-nil error
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := notifications.Send(tt.args.user, tt.args.typeId, tt.args.message)

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

// Type 1, only 2 messages per minute per user
func Test_solution_1(t *testing.T) {

	got := notifications.Send("user2", 1, "message")
	assert.Equal(t, got, nil)

	got2 := notifications.Send("user2", 1, "message")
	assert.Equal(t, got2, nil)

	got3 := notifications.Send("user2", 1, "message")
	assert.NotNil(t, got3)
}

// Type 2, only 3 messages per hour per user
func Test_solution_2(t *testing.T) {

	got := notifications.Send("user3", 2, "message")
	assert.Equal(t, got, nil)

	got1 := notifications.Send("user3", 2, "message")
	assert.Equal(t, got1, nil)

	got2 := notifications.Send("user3", 2, "message")
	assert.Equal(t, got2, nil)

	got3 := notifications.Send("user3", 2, "message")
	assert.NotNil(t, got3)
}

// Type 3, only 1 messages per day per user
func Test_solution_3(t *testing.T) {

	got := notifications.Send("user4", 3, "message")
	assert.Equal(t, got, nil)

	got2 := notifications.Send("user4", 3, "message")
	assert.NotNil(t, got2)
}
