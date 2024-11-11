package poker

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game *Game
}

const PlayerPrompt = "Please enter the number of players: "

func NewCLI(in io.Reader, out io.Writer, game *Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayerInput := cli.readLine()
	numberOfPlayers, _ := strconv.Atoi(strings.Trim(numberOfPlayerInput, "\n"))

	cli.game.Start(numberOfPlayers)

	winnerInput := cli.readLine()
	winner := extractWinner(winnerInput)
	cli.game.Finish(winner)
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
	//blindIncrement := 10 * time.Minute
	blindIncrement := time.Duration(10+numberOfPlayers) * time.Minute

	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	//blindTime := 5 * time.Minute
	for _, blind := range blinds {
		cli.game.alerter.ScheduleAlertAt(blindTime, blind)
		blindTime = blindTime + blindIncrement
	}
}
