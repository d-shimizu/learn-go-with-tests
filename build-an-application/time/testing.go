package poker

import (
	"fmt"
	"testing"
	"time"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func AssertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
	t.Helper()

	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}

	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

type ScheduledAlert struct {
	At     time.Duration
	Amount int
}

func (s ScheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.Amount, s.At)
}

// SpyBlindAlerter allows you to spy on ScheduleAlertAt calls.
type SpyBlindAlerter struct {
	Alerts []ScheduledAlert
}

// ScheduleAlertAt records alerts that have been scheduled.
func (s *SpyBlindAlerter) ScheduleAlertAt(at time.Duration, amount int) {
	s.Alerts = append(s.Alerts, ScheduledAlert{at, amount})
}
