package db

import (
	"context"
	"github.com/AveryKing/tasks/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func createTestTask(t *testing.T) (task Task) {
	user := createTestUser(t)

	arg := CreateTaskParams{
		User:     user.ID,
		Title:    util.RandomString(10),
		Priority: 1,
	}

	task, err := testQueries.CreateTask(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.User, task.User)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Priority, task.Priority)
	return
}

func TestCreateTask(t *testing.T) {
	createTestTask(t)
}
