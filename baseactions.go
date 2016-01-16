package cdtest

import (
	"errors"
	"strconv"
	"strings"
)

//Action is an interface
type Action interface {
	Run() string
}

//NewAction create a new Action object
func NewAction(action string, params ...string) (Action, error) {
	var res Action
	switch strings.ToUpper(action) {
	case "SAY":
		res = &SayAction{Params: params}
	case "CALC":
		tmp := make([]int, 0, len(params))
		for _, num := range params {
			val, err := strconv.Atoi(num)
			if err != nil {
				return nil, err
			}
			tmp = append(tmp, val)
		}
		res = &CalcAction{Params: tmp}
	case "GREATING":
		if len(params) > 1 {
			return nil, errors.New("too many names")
		}
		res = &GreatingAction{Param: params[0]}

	}
	return res, nil
}

//SayAction i don't know
type SayAction struct {
	Params []string
}

//Run is stupid
func (a SayAction) Run() string {
	return strings.Join(a.Params, " ")
}

//CalcAction this is fun
type CalcAction struct {
	Params []int
}

//Run is stupid as above
func (a CalcAction) Run() string {
	i := 0
	for _, x := range a.Params {
		i += x
	}
	return strconv.Itoa(i)
}

//GreatingAction this is fun
type GreatingAction struct {
	Param string
}

//Run is stupid as above
func (a GreatingAction) Run() string {
	return strings.Join([]string{"HI", a.Param}, ", ")
}
