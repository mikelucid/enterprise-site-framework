// Package domain models Flow Genius: an AI music writer combined with
// Beatport-marketplace label management.
package domain

import (
	"errors"
	"time"
)

type ReleaseStatus string

const (
	ReleaseDraft     ReleaseStatus = "draft"
	ReleaseSubmitted ReleaseStatus = "submitted"
	ReleaseLive      ReleaseStatus = "live"
	ReleaseRejected  ReleaseStatus = "rejected"
)

var (
	ErrInvalidSplits = errors.New("royalty splits must be positive and sum to 10000 basis points")
	ErrNoTracks      = errors.New("release requires at least one track")
	ErrNotDraft      = errors.New("release is not in draft status")
	ErrNoLabel       = errors.New("release requires a label")
)

// Split is a royalty share in basis points (10000 = 100%).
type Split struct {
	PayeeID string
	BPS     int
}

// CompositionRequest is the prompt for the AI music writer.
type CompositionRequest struct {
	Genre     string // Beatport genre, e.g. "Melodic House & Techno"
	BPM       int
	Key       string
	Mood      string
	Prompt    string
	CreatorID string
}

// Track is an AI-composed (and optionally human co-written) track.
type Track struct {
	ID       string
	Title    string
	Genre    string
	BPM      int
	Key      string
	AudioURI string
	ISRC     string
	AIModel  string
	Splits   []Split
}

type Label struct {
	ID      string
	Name    string
	OwnerID string
}

type Release struct {
	ID         string
	LabelID    string
	Title      string
	Tracks     []Track
	Status     ReleaseStatus
	BeatportID string
	ReleaseAt  time.Time
}

// ValidateSplits checks that splits are positive and total 100%.
func ValidateSplits(s []Split) error {
	total := 0
	for _, x := range s {
		if x.BPS <= 0 || x.PayeeID == "" {
			return ErrInvalidSplits
		}
		total += x.BPS
	}
	if total != 10000 {
		return ErrInvalidSplits
	}
	return nil
}

// Allocate divides amount (minor units) by splits; the remainder goes to the
// first payee so the total is always preserved.
func Allocate(amount int64, s []Split) (map[string]int64, error) {
	if err := ValidateSplits(s); err != nil {
		return nil, err
	}
	out := make(map[string]int64, len(s))
	var used int64
	for _, x := range s {
		share := amount * int64(x.BPS) / 10000
		out[x.PayeeID] += share
		used += share
	}
	out[s[0].PayeeID] += amount - used
	return out, nil
}
