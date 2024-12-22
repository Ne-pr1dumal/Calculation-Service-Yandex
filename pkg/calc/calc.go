package calc

import (
	"errors"
	"strconv"
	"strings"
)

var operations = map[string]int{
	")": -1,
	"(": 0,
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}

var (
	ErrInvalidExpression = errors.New("expression is not valid")
	ErrDivisionByZero    = errors.New("division by zero")
)

type Stack struct {
	values []string
}

func (st *Stack) Counter(s string) error {
	if st.len() < 2 {
		return ErrInvalidExpression
	}
	value, err := st.Pop()
	if err != nil {
		return err
	}
	b, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrInvalidExpression
	}
	value, err = st.Pop()
	if err != nil {
		return nil
	}
	a, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return ErrInvalidExpression
	}
	switch s {
	case "+":
		st.Push(strconv.FormatFloat(a+b, 'f', -1, 64))
	case "-":
		st.Push(strconv.FormatFloat(a-b, 'f', -1, 64))
	case "/":
		if b == 0 {
			return ErrDivisionByZero
		}
		st.Push(strconv.FormatFloat(a/b, 'f', -1, 64))
	case "*":
		st.Push(strconv.FormatFloat(a*b, 'f', -1, 64))
	}
	return nil
}

func (st *Stack) Push(value string) {
	st.values = append(st.values, value)
}

func (st *Stack) Pop() (string, error) {
	if st.len() == 0 {
		return "", errors.New("inccorrect expression")
	}
	index := len(st.values) - 1
	value := st.values[index]
	st.values = st.values[:index]
	return value, nil
}

func (st *Stack) len() int {
	return len(st.values)
}

// Алгоритм Дейкстра
func operationPriority(str string, priorityStr int, output *[]string, s *Stack) error {
	for (*s).len() > 0 {
		value, err := (*s).Pop()
		if err != nil {
			return err
		}
		priorityStack, exist := operations[value] // Приоритет операции из стека
		if !exist {
			return errors.New("can`t find priority")
		}
		if priorityStack > priorityStr && priorityStr != 0 {
			if priorityStack == 0 {
				return nil
			}
			*output = append(*output, value)
		} else if priorityStr != 3 {
			(*s).Push(value)
			(*s).Push(str)
			break
		}
	}
	if (*s).len() == 0 && priorityStr != 3 {
		(*s).Push(str)
	}
	return nil
}

func JoinSlice(MapOfString, output *[]string) {
	if len(*MapOfString) == 0 {
		return
	}
	Str := strings.Join(*MapOfString, "")
	*output = append(*output, Str)
	*MapOfString = []string{}
}

func ShuntingYard(expression string) ([]string, error) {
	var s Stack
	var output, StrForInt []string
	for _, str := range expression {
		str := string(str)
		value, exist := operations[str]
		if exist {
			JoinSlice(&StrForInt, &output)
			operationPriority(str, value, &output, &s)
		} else {
			StrForInt = append(StrForInt, str)
		}
	}
	if len(StrForInt) > 0 {
		JoinSlice(&StrForInt, &output)
	}
	for s.len() > 0 {
		value, err := s.Pop()
		if err != nil {
			return nil, err
		}
		output = append(output, value)
	}
	return output, nil
}

func GetAnsFromStack(s *Stack) (float64, error) {
	if (*s).len() > 1 {
		return 0, errors.New("incorrect expression. want one elem have more")
	}
	value, err := (*s).Pop()
	if err != nil {
		return 0, err
	}
	ans, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, errors.New("can`t convert to int")
	}
	return ans, nil
}

func Calc(expression string) (float64, error) {
	var s Stack
	PostfixExpression, err := ShuntingYard(expression)
	if err != nil {
		return 0, err
	}
	for _, value := range PostfixExpression {
		if _, exist := operations[value]; exist {
			err := s.Counter(value)
			if err != nil {
				return 0, err
			}
		} else if value != " " {
			s.Push(value)
		}
	}
	ans, err := GetAnsFromStack(&s)
	if err != nil {
		return 0, err
	}
	return float64(ans), nil
}
