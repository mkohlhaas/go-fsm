package main

import (
	"fmt"
	"math/rand"
)

// State /////////////////////////////////////////

var (
	Locked   State = state{"Locked"}
	Unlocked State = state{"Unlocked"}
)

type state struct {
	sz string
}

type State interface {
	isState()
}

func (state) isState() {}

func (s state) String() string { return s.sz }

// Event /////////////////////////////////////////

var (
	Coin Event = event{"Coin"}
	Push Event = event{"Push"}
)

type event struct {
	sz string
}

type Event interface {
	isEvent()
}

func (event) isEvent() {}

func (s event) String() string { return s.sz }

// FSM Interface //////////////////////////////////

type IFSM interface {
	CurrState() State
	TransitionOn(event Event)
}

/// FSM ///////////////////////////////////////////

type FSM struct {
	currentState State
	stMap        stateMap
}

type stateMap map[State]transitionMap
type transitionMap map[Event]Action

type Action struct {
	// can be enriched with more functionality,
	// e.g a real action depending on use case
	dstState State
}

// IFMS Implementation ////////////////////////////

var _ IFSM = &FSM{}

func (fms *FSM) CurrState() State {
	return fms.currentState
}

func (fsm *FSM) TransitionOn(event Event) {
	action, ok := fsm.stMap[fsm.currentState][event]
	if ok {
		fsm.currentState = action.dstState
	} else {
		fmt.Printf("Transition '%s' is invalid in state '%s'.\n",
			event, fsm.currentState)
	}
}

// Main ///////////////////////////////////////////

func main() {
	fsm := FSM{
		currentState: Locked,
		stMap: stateMap{
			Locked: transitionMap{
				Coin: Action{Unlocked},
				Push: Action{Locked},
			},
			Unlocked: transitionMap{
				Coin: Action{Unlocked},
				Push: Action{Locked},
			},
		},
	}

	events := []Event{Coin, Push}
	for i := 0; i < 10000; i++ {
		rndEvent := events[rand.Intn(len(events))]
		fmt.Println("State before:", fsm.CurrState())
		fmt.Printf("Event #%-5d  %s\n", i+1, rndEvent)
		fsm.TransitionOn(rndEvent)
		fmt.Println("State after: ", fsm.CurrState())
		fmt.Println("-----------------------")
	}
}
