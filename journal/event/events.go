package event

import (
	shared "remind/pkg/shared_kernel"
)

type ImageCreated struct {
	shared.DomainEvent
	Type      string `json:"type"`
	ID        int64  `json:"image_id"`
	JournalID int64  `json:"journal_id"`
	Content   string `json:"content"`
	Username  string `json:"username"`
}

func (d *ImageCreated) Identity() string {
	return "ImageCreated"
}

type JournalCreated struct {
	shared.DomainEvent
	Type     string `json:"type"`
	ID       int64  `json:"journal_id"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

func (d *JournalCreated) Identity() string {
	return "JournalCreated"
}

type JournalEmbedded struct {
	shared.DomainEvent
	Type     string `json:"type"`
	ID       int64  `json:"id"`
}

func (d *JournalEmbedded) Identity() string {
	return "JournalEmbedded"
}

type ImageEmbedded struct {
	shared.DomainEvent
	Type      string `json:"type"`
	ID        int64  `json:"id"`
}

func (d *ImageEmbedded) Identity() string {
	return "ImageEmbedded"
}
