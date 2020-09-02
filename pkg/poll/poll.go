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
	res := fmt.Sprintf("%s\n\n", p.Name)
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
