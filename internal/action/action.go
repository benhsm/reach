package action

type status int

const (
	StatusPending status = iota
	StatusDone
	StatusCancelled
)

type Action struct {
	ID     int64
	Status status // pending, done, cancelled

	Desc          string // short description of what the action is
	Difficulty    int    // rating from 1 to 7
	Notes         string // associated thoughts, feelings, etc
	StartStrategy string // how one might start

	Reflection   string // how it was doing the thing
	OutcomeValue int    // rating of how it was to do the thing
}
