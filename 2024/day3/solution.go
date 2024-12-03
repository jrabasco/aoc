package day3

import (
	"fmt"
	"github.com/jrabasco/aoc/2024/framework/parse"
	"strconv"
	"strings"
	"unicode"
)

type Spec string

const (
	One Spec = "one"
	Two Spec = "two"
)

type FSM struct {
	spec  Spec
	state State
	on    bool
	n1s   string
	n1    int
	n2s   string
	n2    int
	res   int
}

func NewFSM(spec Spec) *FSM {
	return &FSM{spec, Wait{}, true, "", 0, "", 0, 0}
}

type State interface {
	processEvent(*FSM, rune) State
}

type Wait struct{}

func (s Wait) processEvent(fsm *FSM, c rune) State {
	fsm.n1s = ""
	fsm.n2s = ""
	fsm.n1 = 0
	fsm.n2 = 0
	if c == 'm' && fsm.on {
		return M{}
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}

	return Wait{}
}

type WaitOn struct{}

func (s WaitOn) processEvent(fsm *FSM, c rune) State {
	return Wait{}
}

type WaitOff struct{}

func (s WaitOff) processEvent(fsm *FSM, c rune) State {
	return Wait{}
}

type M struct{}

func (s M) processEvent(fsm *FSM, c rune) State {
	if c == 'u' {
		return U{}
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}
	return Wait{}
}

type U struct{}

func (s U) processEvent(fsm *FSM, c rune) State {
	if c == 'l' {
		return L{}
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}
	return Wait{}
}

type L struct{}

func (s L) processEvent(fsm *FSM, c rune) State {
	if c == '(' {
		return MulLParen{}
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}
	return Wait{}
}

type MulLParen struct{}

func (s MulLParen) processEvent(fsm *FSM, c rune) State {
	if unicode.IsDigit(c) {
		fsm.n1s = string(c)
		return Digit1{}
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}
	return Wait{}
}

type Digit1 struct{}

func (s Digit1) processEvent(fsm *FSM, c rune) State {
	if unicode.IsDigit(c) {
		fsm.n1s += string(c)
		return Digit1{}
	}

	if c == ',' {
		fsm.n1, _ = strconv.Atoi(fsm.n1s)
		return Digit2{}
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}
	return Wait{}
}

type Digit2 struct{}

func (s Digit2) processEvent(fsm *FSM, c rune) State {
	if unicode.IsDigit(c) {
		fsm.n2s += string(c)
		return Digit2{}
	}
	if c == ')' && len(fsm.n2s) > 0 {
		fsm.n2, _ = strconv.Atoi(fsm.n2s)
		fsm.res += fsm.n1 * fsm.n2
	}

	if c == 'd' && fsm.spec == Two {
		return D{}
	}
	return Wait{}
}

type D struct{}

func (s D) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}

	if c == 'o' {
		return O{}
	}
	return Wait{}
}

type O struct{}

func (s O) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}

	if c == '(' {
		return DoLParen{}
	}

	if c == 'n' {
		return N{}
	}
	return Wait{}
}

type DoLParen struct{}

func (s DoLParen) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}

	if c == ')' {
		fsm.on = true
	}
	return Wait{}
}

type N struct{}

func (s N) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}

	if c == '\'' {
		return Apos{}
	}

	return Wait{}
}

type Apos struct{}

func (s Apos) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}

	if c == 't' {
		return T{}
	}
	return Wait{}
}

type T struct{}

func (s T) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}

	if c == '(' {
		return DontLParen{}
	}
	return Wait{}
}

type DontLParen struct{}

func (s DontLParen) processEvent(fsm *FSM, c rune) State {
	if c == 'd' {
		return D{}
	}
	if c == ')' {
		fsm.on = false
	}
	return Wait{}
}

func Solution() int {
	parsed, err := parse.GetLines("day3/input.txt")
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}
	code := strings.Join(parsed, "\n")

	fsmOne := NewFSM(One)
	fsmTwo := NewFSM(Two)
	for _, c := range code {
		fsmOne.state = fsmOne.state.processEvent(fsmOne, c)
		fsmTwo.state = fsmTwo.state.processEvent(fsmTwo, c)
	}
	fmt.Printf("Part 1: %d\n", fsmOne.res)
	fmt.Printf("Part 1: %d\n", fsmTwo.res)
	return 0
}
