package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomJournal(t *testing.T) Journal {
	content := pgtype.Text{
		String: "Content",
		Valid: true,
	}

	arg := CreateJournalParams{
		Username: "user",
		Title:   "Title",
		Content: content,
		Status: false,
	}

	journal, err := testStore.CreateJournal(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, journal)

	require.Equal(t, arg.Title, journal.Title)
	require.Equal(t, arg.Content, journal.Content)
	require.Equal(t, arg.Status, journal.Status)
	require.NotEmpty(t, journal.ID)
	require.Equal(t, arg.Username, journal.Username)
	require.NotZero(t, journal.CreatedAt)

	return journal
}

func TestCreateJournal(t *testing.T) {
	createRandomJournal(t)
}

func TestGetJournal(t *testing.T) {
	journal1 := createRandomJournal(t)
	journal2, err := testStore.GetJournal(context.Background(), journal1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, journal2)

	require.Equal(t, journal1.ID, journal2.ID)
	require.Equal(t, journal1.Title, journal2.Title)
	require.Equal(t, journal1.Content, journal2.Content)
	require.Equal(t, journal1.Status, journal2.Status)
	require.Equal(t, journal1.Username, journal2.Username)
	require.WithinDuration(t, journal1.CreatedAt, journal2.CreatedAt, time.Second)
	require.WithinDuration(t, journal1.UpdatedAt, journal2.UpdatedAt, time.Second)
}

func TestUpdateJournalOnlyContent(t *testing.T) {
	oldJournal := createRandomJournal(t)

	newContent := pgtype.Text{
		String: "New Content",
		Valid: true,
	}
	updatedJournal, err := testStore.UpdateJournal(context.Background(), UpdateJournalParams{
		ID: oldJournal.ID,
		Content: newContent,
	})

	require.NoError(t, err)
	require.NotEmpty(t, updatedJournal)

	require.Equal(t, oldJournal.ID, updatedJournal.ID)
	require.Equal(t, oldJournal.Title, updatedJournal.Title)
	require.Equal(t, newContent, updatedJournal.Content)
	require.Equal(t, oldJournal.Status, updatedJournal.Status)
	require.Equal(t, oldJournal.Username, updatedJournal.Username)
	require.WithinDuration(t, oldJournal.CreatedAt, updatedJournal.CreatedAt, time.Second)
	require.WithinDuration(t, oldJournal.UpdatedAt, updatedJournal.UpdatedAt, time.Second)
}