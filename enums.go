package main

import "fmt"

/*
ServerState is a custom type based on int.
This is how Go simulates enums.
*/
type ServerState int

/*
Enum values using iota.
iota starts from 0 and increments automatically.
*/
const (
	StateIdle      ServerState = iota // 0
	StateConnected                    // 1
	StateError                        // 2
	StateRetrying                     // 3
)

/*
Map each ServerState to a readable string.
This helps when printing states.
*/
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

/*
String method lets ServerState satisfy fmt.Stringer.
So fmt.Println(state) prints text instead of number.
*/
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {

	// Start from Idle state
	ns := transition(StateIdle)
	fmt.Println(ns) // prints "connected"

	// Transition again
	ns2 := transition(ns)
	fmt.Println(ns2) // prints "idle"
}

/*
transition defines rules for moving
from one server state to another.
*/
func transition(s ServerState) ServerState {
	switch s {

	case StateIdle:
		return StateConnected

	case StateConnected, StateRetrying:
		return StateIdle

	case StateError:
		return StateError

	default:
		// Panic if unknown state appears
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
