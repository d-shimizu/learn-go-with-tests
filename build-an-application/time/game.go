package poker

import "time"

type Game interface {
	Start(numberOfPlayers int)
	Finish(winner string)
}

type GameSpy struct {
	alerter BlindAlerter
	store   PlayerStore
}

func NewGame(alerter BlindAlerter, store PlayerStore) *GameSpy {
	return &GameSpy{
		alerter: alerter,
		store:   store,
	}
}

func (p *GameSpy) Start(numberOfPlayers int) {
	blindIncrement := time.Duration(10+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	//blindTime := time.Duration(5 * time.Minute)
	//blindTime := 1 * time.Minute

	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}

func (p *GameSpy) Finish(winner string) {
	p.store.RecordWin(winner)
}
