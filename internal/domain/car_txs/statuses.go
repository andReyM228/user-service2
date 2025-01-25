package car_txs

const (
	StatusNew    = "new"
	StatusDone   = "done"
	StatusFailed = "failed"
)

type Status string

var Statuses = map[Status]struct{}{
	StatusNew:    {},
	StatusFailed: {},
	StatusDone:   {},
}

func (s Status) String() string {
	return string(s)
}

func (c *CarTx) SetStatusFailed(err error) {
	c.Error = err.Error()
	c.Status = StatusFailed
}

func (c *CarTx) SetStatusDone() {
	c.Status = StatusDone
}
