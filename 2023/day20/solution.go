package day20

import (
	"fmt"
	"github.com/jrabasco/aoc/2023/framework/parse"
	"github.com/jrabasco/aoc/2023/framework/utils"
	"strings"
)

type Pulse int

const (
	HIGH Pulse = iota
	LOW
)

func (p Pulse) String() string {
	switch p {
	case HIGH:
		return "high"
	case LOW:
		return "low"
	default:
		return "wat"
	}
}

type SentPulse struct {
	pulse Pulse
	from  string
	to    string
}

func (sp SentPulse) String() string {
	return fmt.Sprintf("%s -%s-> %s", sp.from, sp.pulse, sp.to)
}

type ModType int

const (
	FLIP ModType = iota
	CONJ
	BROAD
	OUT
)

type Module interface {
	Handle(SentPulse) []SentPulse
	AddInput(string)
	AddOutput(string)
	Type() ModType
}

type FlipFlop struct {
	on      bool
	label   string
	outputs []string
	inputs  []string
}

func NewFlipFlop(label string) *FlipFlop {
	return &FlipFlop{
		on:      false,
		label:   label,
		outputs: []string{},
		inputs:  []string{},
	}
}

func (f *FlipFlop) Type() ModType {
	return FLIP
}

func (f *FlipFlop) AddInput(label string) {
	f.inputs = append(f.inputs, label)
}

func (f *FlipFlop) AddOutput(label string) {
	f.outputs = append(f.outputs, label)
}
func (f *FlipFlop) Handle(p SentPulse) []SentPulse {
	res := []SentPulse{}
	if p.pulse == HIGH {
		return res
	}

	f.on = !f.on

	pOut := LOW
	if f.on {
		pOut = HIGH
	}

	for _, d := range f.outputs {
		res = append(res, SentPulse{pOut, f.label, d})
	}
	return res
}

type Conjunction struct {
	label   string
	inputs  map[string]Pulse
	outputs []string
}

func NewConjunction(label string) *Conjunction {
	return &Conjunction{
		label:   label,
		inputs:  map[string]Pulse{},
		outputs: []string{},
	}
}

func (c *Conjunction) Type() ModType {
	return CONJ
}

func (c *Conjunction) AddInput(label string) {
	c.inputs[label] = LOW
}
func (c *Conjunction) AddOutput(label string) {
	c.outputs = append(c.outputs, label)
}

func (c *Conjunction) Handle(p SentPulse) []SentPulse {
	c.inputs[p.from] = p.pulse
	pOut := LOW
	for _, pMem := range c.inputs {
		if pMem == LOW {
			pOut = HIGH
			break
		}
	}
	res := []SentPulse{}
	for _, d := range c.outputs {
		res = append(res, SentPulse{pOut, c.label, d})
	}
	return res
}

type Output struct {
	label string
	input string
}

func NewOutput(label string) *Output {
	return &Output{label, ""}
}

func (o *Output) Type() ModType {
	return OUT
}

func (o *Output) AddInput(label string) {
	if o.input != "" {
		// it's a bit weird but only one output exists in the problem
		// and it has only one input
		panic("Impossible!")
	}
}
func (o *Output) AddOutput(label string) {
}

func (o *Output) Handle(p SentPulse) []SentPulse {
	return []SentPulse{}
}

type Broadcaster struct {
	outputs []string
}

func NewBroadcaster() *Broadcaster {
	return &Broadcaster{[]string{}}
}

func (b *Broadcaster) Type() ModType {
	return BROAD
}

func (b *Broadcaster) AddInput(label string) {
}

func (b *Broadcaster) AddOutput(label string) {
	b.outputs = append(b.outputs, label)
}

func (b *Broadcaster) Handle(p SentPulse) []SentPulse {
	res := []SentPulse{}
	for _, d := range b.outputs {
		res = append(res, SentPulse{p.pulse, "broadcaster", d})
	}
	return res
}

type Circuit struct {
	modules   map[string]Module
	output    *Output
	preOutput *Conjunction
	// when the inputs of preOutput got turned to HIGH
	preOutputStamps map[string]int
	presses         int
	q               utils.Queue[SentPulse]
}

func NewCircuit() Circuit {
	return Circuit{
		modules:         map[string]Module{},
		output:          nil,
		preOutput:       nil,
		preOutputStamps: map[string]int{},
		presses:         0,
		q:               utils.NewQueue[SentPulse](),
	}
}

func (c *Circuit) Stamped() bool {
	for _, v := range c.preOutputStamps {
		if v == 0 {
			return false
		}
	}
	return true
}

func (c *Circuit) StampValue() int {
	res := 1

	for _, v := range c.preOutputStamps {
		res = utils.LCM(res, v)
	}
	return res
}

func (c *Circuit) PressButton() (int, int) {
	lowCount := 0
	highCount := 0
	c.presses += 1
	c.q.Enqueue(SentPulse{LOW, "button", "broadcaster"})
	for !c.q.Empty() {
		sp, _ := c.q.Dequeue()
		if sp.pulse == HIGH {
			highCount += 1
		} else {
			lowCount += 1
		}
		mod, _ := c.modules[sp.to]
		if sp.to == c.preOutput.label && mod.Type() == CONJ {
			for k, v := range c.preOutputStamps {
				// first time k sends HIGH
				if v == 0 && k == sp.from && sp.pulse == HIGH {
					c.preOutputStamps[k] = c.presses
				}
			}
		}
		c.q.Enqueue(mod.Handle(sp)...)
	}
	return lowCount, highCount
}

func parseCircuit(lines []string) (Circuit, error) {
	circuit := NewCircuit()
	inouts := [][]string{}
	for _, line := range lines {
		inout := strings.Split(line, " -> ")
		if len(inout) != 2 {
			return circuit, fmt.Errorf("invalid module line: %s", line)
		}
		inouts = append(inouts, inout)
		modStr := inout[0]
		if len(modStr) < 2 {
			return circuit, fmt.Errorf("invalid module spec: %s", modStr)
		}

		if modStr == "broadcaster" {
			circuit.modules["broadcaster"] = NewBroadcaster()
			continue
		}

		label := modStr[1:len(modStr)]
		switch modStr[0] {
		case '%':
			circuit.modules[label] = NewFlipFlop(label)
		case '&':
			circuit.modules[label] = NewConjunction(label)
		default:
			return circuit, fmt.Errorf("invalid module id: %s", string(modStr[0]))
		}
	}

	// now we have added all the modules to the circuit, need to add
	// connections

	for _, inout := range inouts {
		in := inout[0]
		outstr := inout[1]
		outs := strings.Split(outstr, ", ")
		if in != "broadcaster" {
			in = in[1:len(in)]
		}
		inM, ok := circuit.modules[in]
		if !ok {
			return circuit, fmt.Errorf("could not find %s in circuit?!?", in)
		}

		for _, out := range outs {
			outM, ok := circuit.modules[out]
			if !ok {
				// output-only module
				output := NewOutput(out)
				circuit.output = output
				outM = output
				circuit.modules[out] = output
				// this is based on our knowledge of the problem: there is only
				// one input to the unique output and it's a conjunction
				circuit.preOutput = inM.(*Conjunction)
			}
			inM.AddOutput(out)
			outM.AddInput(in)
		}
	}
	for k := range circuit.preOutput.inputs {
		circuit.preOutputStamps[k] = 0
	}
	return circuit, nil
}

func Solution() int {
	circuit, err := parse.GetLinesAsOne[Circuit]("day20/input.txt", parseCircuit)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return 1
	}

	lowCount := 0
	highCount := 0
	for i := 0; i < 1000; i++ {
		nlc, nhc := circuit.PressButton()
		lowCount += nlc
		highCount += nhc
	}
	fmt.Printf("Part 1: %d\n", lowCount*highCount)

	// This solution is based on the following observations:
	// 1. The inputs the the conjuction that feeds into the output are also
	//    conjunctions
	// 2. There are cycles which leads to those conjunction turning HIGH on
	//    multiples of their initial activation
	// The answer is therefore the LCM of the initial activations of those
	// Conjunctions
	for !circuit.Stamped() {
		circuit.PressButton()
	}

	fmt.Printf("Part 2: %d\n", circuit.StampValue())
	return 0
}
