package tasks

import (
	"context"
	"fmt"
	
	"github.com/jfyne/live"
)

type (
	state struct {
		Tasks []task
		Form  form
	}
	
	task struct {
		ID       string
		Name     string
		Complete bool
	}
	
	form struct {
		Errors map[string]string
	}
)

func newState(s live.Socket) *state {
	m, ok := s.Assigns().(*state)
	if !ok {
		return &state{
			Form: form{
				Errors: map[string]string{},
			},
		}
	}
	
	m.Form.Errors = map[string]string{}
	return m
}

func onMount(ctx context.Context, socket live.Socket) (interface{}, error) {
	return newState(socket), nil
}

func onValidate(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	m := newState(s)
	t := p.String("task")
	vm := validateMessage(t)
	if vm != "" {
		m.Form.Errors["message"] = vm
	}
	return m, nil
}

func onSave(ctx context.Context, s live.Socket, p live.Params) (interface{}, error) {
	m := newState(s)
	ts := p.String("task")
	complete := p.Checkbox("complete")
	vm := validateMessage(ts)
	if vm != "" {
		m.Form.Errors["message"] = vm
	} else {
		t := task{
			ID:       live.NewID(),
			Name:     ts,
			Complete: complete,
		}
		m.Tasks = append(m.Tasks, t)
	}
	return m, nil
}

func onComplete(ctx context.Context, socket live.Socket, p live.Params) (interface{}, error) {
	state := newState(socket)
	ID := p.String("id")
	for idx, t := range state.Tasks {
		if t.ID != ID {
			continue
		}
		state.Tasks[idx].Complete = !state.Tasks[idx].Complete
	}
	return state, nil
}

func validateMessage(msg string) string {
	if len(msg) < 10 {
		return fmt.Sprintf("Length of 10 required, have %d", len(msg))
	}
	if len(msg) > 20 {
		return fmt.Sprintf("Your task name is too long > 20, have %d", len(msg))
	}
	return ""
}
