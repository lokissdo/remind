package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomAudio(t *testing.T, randomJournal Journal) Audio {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	path := dir + "/test.mp3"

	audioBytes, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("Failed to read image file: %v", err)
    }

	arg := CreateAudioParams{
		Content: audioBytes,
		JournalID: randomJournal.ID,
	}

	audio, err := testStore.CreateAudio(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, audio)

	require.Equal(t, arg.Content, audio.Content)
	require.NotEmpty(t, audio.ID)

	return audio
}

func TestCreateRandomAudio(t *testing.T) {
	randomJournal := createRandomJournal(t)
	createRandomAudio(t, randomJournal)
}