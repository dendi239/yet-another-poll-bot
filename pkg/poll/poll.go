package poll

import (
	"fmt"
	"strings"

	"github.com/dendi239/yet-another-poll-bot/pkg/grammar"
)

// UserOptions represents all data specific to user
type UserOptions struct {
	username string
	context  grammar.Context
}

// Poll represents all data specific to poll and users voted for
type Poll struct {
	Name         string
	OptionsOrder []int
	Users        []UserOptions
	Options      map[int]string
}

func (p *Poll) String() string {
	res := fmt.Sprintf("%s\n", p.Name)
	for _, option := range p.OptionsOrder {
		us := make([]string, 0)
		for _, u := range p.Users {
			if contains, err := u.context.Variables[option]; err && contains {
				us = append(us, u.username)
			}
		}
		line := p.Options[option] + ": " + strings.Join(us, ", ") + "\n"
		res += line
	}
	return res
}

var (
	daniel = UserOptions{
		"@ostashev",
		grammar.Context{
			Variables: map[int]bool{
				1: false,
				2: true,
				3: true,
				4: true,
				5: true,
				6: false,
			},
		},
	}
	denys = UserOptions{
		"@dendi239",
		grammar.Context{
			Variables: map[int]bool{
				1: true,
				2: true,
				3: true,
				4: true,
				5: false,
				6: false,
			},
		},
	}
	belkka = UserOptions{
		"@belkka",
		grammar.Context{
			Variables: map[int]bool{
				1: false,
				2: true,
				3: false,
				4: false,
				5: true,
				6: true,
			},
		},
	}
	// Sample poll to test interface
	Sample = Poll{
		"DAF Pub",
		[]int{1, 2, 3, 4, 5, 6},
		[]UserOptions{daniel, denys, belkka},
		map[int]string{
			1: "tue 20:00",
			2: "tue 21:00",
			3: "wen 20:00",
			4: "wen 19:00",
			5: "wen 21:00",
			6: "stay home",
		},
	}
)
