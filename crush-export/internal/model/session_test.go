package model

import (
	"testing"
	"time"
)

func TestComputeStatistics_EmptyMessages(t *testing.T) {
	s := Session{
		ID:    "test-1",
		Title: "Empty Session",
	}

	stats := s.ComputeStatistics()

	if stats.MessageCount != 0 {
		t.Errorf("MessageCount = %d, want 0", stats.MessageCount)
	}
	if stats.Duration != 0 {
		t.Errorf("Duration = %v, want 0", stats.Duration)
	}
}

func TestComputeStatistics_WithMessages(t *testing.T) {
	now := time.Now()
	s := Session{
		ID:    "test-2",
		Title: "Test Session",
		Messages: []Message{
			{ID: "m1", Role: "user", CreatedAt: now},
			{ID: "m2", Role: "assistant", CreatedAt: now.Add(30 * time.Second)},
			{ID: "m3", Role: "user", CreatedAt: now.Add(60 * time.Second)},
			{ID: "m4", Role: "assistant", CreatedAt: now.Add(90 * time.Second)},
		},
	}

	stats := s.ComputeStatistics()

	if stats.MessageCount != 4 {
		t.Errorf("MessageCount = %d, want 4", stats.MessageCount)
	}
	if stats.UserMessages != 2 {
		t.Errorf("UserMessages = %d, want 2", stats.UserMessages)
	}
	if stats.AssistantMsgs != 2 {
		t.Errorf("AssistantMsgs = %d, want 2", stats.AssistantMsgs)
	}
	if stats.Duration != 90*time.Second {
		t.Errorf("Duration = %v, want 90s", stats.Duration)
	}
}

func TestComputeStatistics_SingleMessage(t *testing.T) {
	now := time.Now()
	s := Session{
		ID:    "test-3",
		Title: "Single Message",
		Messages: []Message{
			{ID: "m1", Role: "user", CreatedAt: now},
		},
	}

	stats := s.ComputeStatistics()

	if stats.MessageCount != 1 {
		t.Errorf("MessageCount = %d, want 1", stats.MessageCount)
	}
	if stats.Duration != 0 {
		t.Errorf("Duration = %v, want 0 (single message)", stats.Duration)
	}
}
