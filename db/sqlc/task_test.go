package db

import (
	"context"
	"github.com/AveryKing/tasks/util"
	"github.com/stretchr/testify/require"
	"testing"
)

// If userId is not nil, task will be created for the specified user
func createTestTask(t *testing.T, userId int64) Task {
	arg := CreateTaskParams{
		User:     userId,
		Title:    util.RandomString(10),
		Priority: 1,
	}

	task, err := testQueries.CreateTask(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, arg.User, task.User)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Priority, task.Priority)
	return task
}

func TestCreateTask(t *testing.T) {
	user := createTestUser(t)
	createTestTask(t, user.ID)
}

func TestGetTasks(t *testing.T) {
	user := createTestUser(t)

	task1 := createTestTask(t, user.ID)
	task2 := createTestTask(t, user.ID)

	require.NotEmpty(t, task1)
	require.NotEmpty(t, task2)
	require.Equal(t, task1.User, user.ID)
	require.Equal(t, task2.User, user.ID)
}
