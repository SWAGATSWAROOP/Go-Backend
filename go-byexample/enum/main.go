package main

import "fmt"

// Enumerated types (enums) are a special case of sum types. An enum is a type that has a fixed number of possible values, each with a distinct name. Go doesn’t have an enum type as a distinct language feature, but enums are simple to implement using existing language idioms.

type ServerState int

// Our enum type ServerState has an underlying int type.
// The possible values for ServerState are defined as constants. The special keyword iota generates successive constant values automatically; in this case 0, 1, 2 and so on.
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
	// sd := 0
	// transition(sd)
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateError:
		return StateError
	case StateRetrying, StateConnected:
		return StateIdle
	default:
		panic(fmt.Errorf("Unknown state : %s", s))
	}
}
