package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomImage(t *testing.T, randomJournal Journal) Image {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}

	path := dir + "/test.jpeg"

	imageBytes, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("Failed to read image file: %v", err)
    }

	arg := CreateImageParams{
		Content: imageBytes,
		JournalID: randomJournal.ID,
	}

	image, err := testStore.CreateImage(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, image)

	require.Equal(t, arg.Content, image.Content)
	require.NotEmpty(t, image.ID)

	err = os.WriteFile(dir + "/output.jpeg", []byte(image.Content), 0644)
    if err != nil {
        t.Fatalf("Failed to write image file: %v", err)
    }

	return image
}

func TestCreateRandomImage(t *testing.T) {
	randomJournal := createRandomJournal(t)
	createRandomImage(t, randomJournal)
}

func TestGetImageOfJournal(t *testing.T) {
	randomJournal := createRandomJournal(t)
	randomImage1 := createRandomImage(t, randomJournal)
	randomImage2 := createRandomImage(t, randomJournal)
	randomImage3 := createRandomImage(t, randomJournal)

	images, err := testStore.GetImageOfJournal(context.Background(), randomJournal.ID)

	require.NoError(t, err)
	require.NotEmpty(t, images)

	require.Equal(t, randomImage3.ID, images[0].ID)
	require.Equal(t, randomImage2.ID, images[1].ID)
	require.Equal(t, randomImage1.ID, images[2].ID)
}