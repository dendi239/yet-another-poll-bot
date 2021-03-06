package poll

import "github.com/dendi239/yet-another-poll-bot/pkg/grammar"

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
	vika = UserOptions{
		"@viskonsin",
		grammar.Context{
			Variables: map[int]bool{
				6: true,
			},
		},
	}
	// Sample poll to test interface
	Sample = Poll{
		"DAF Pub",
		[]int{1, 2, 3, 4, 5, 6},
		[]UserOptions{daniel, denys, belkka, vika},
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
