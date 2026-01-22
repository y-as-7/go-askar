package debugger

import (
	"sync"
	"time"
)

type Query struct {
	SQL      string        `json:"sql"`
	Duration time.Duration `json:"duration"`
	Time     time.Time     `json:"time"`
}

type Log struct {
	Message string    `json:"message"`
	Level   string    `json:"level"`
	Time    time.Time `json:"time"`
}

type DebugData struct {
	StartTime time.Time `json:"start_time"`
	Duration  time.Duration `json:"duration"`
	Queries   []Query   `json:"queries"`
	Logs      []Log     `json:"logs"`
	Memory    uint64    `json:"memory"`
	Views     []string  `json:"views"`
}

type Debugger struct {
	mu   sync.Mutex
	Data *DebugData
}

func New() *Debugger {
	return &Debugger{
		Data: &DebugData{
			StartTime: time.Now(),
			Queries:   []Query{},
			Logs:      []Log{},
			Views:     []string{},
		},
	}
}

func (d *Debugger) AddQuery(sql string, duration time.Duration) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Data.Queries = append(d.Data.Queries, Query{
		SQL:      sql,
		Duration: duration,
		Time:     time.Now(),
	})
}

func (d *Debugger) AddLog(level, message string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Data.Logs = append(d.Data.Logs, Log{
		Level:   level,
		Message: message,
		Time:    time.Now(),
	})
}

func (d *Debugger) AddView(name string) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.Data.Views = append(d.Data.Views, name)
}

func (d *Debugger) Finish() {
	d.Data.Duration = time.Since(d.Data.StartTime)
}
