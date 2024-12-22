package calc

import (
	"errors"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		name       string
		expression string
		want       []string
	}{
		{
			name:       "expression with +",
			expression: "4+2",
			want:       []string{"4", "2", "+"},
		},
		{
			name:       "expression with -",
			expression: "10-2",
			want:       []string{"10", "2", "-"},
		},
		{
			name:       "expression with *",
			expression: "8*4",
			want:       []string{"8", "4", "*"},
		},
		{
			name:       "expression with /",
			expression: "6/3",
			want:       []string{"6", "3", "/"},
		},
		{
			name:       "expression with *, +, -, ()",
			expression: "3+4*(2-1)",
			want:       []string{"3", "4", "2", "1", "-", "*", "+"},
		},
	}

	for _, tc := range cases {
		tc := tc // Защита от замыкания
		t.Run(tc.name, func(t *testing.T) {

			got, err := ShuntingYard(tc.expression)
			if err != nil {
				t.Errorf("ShuntingYard(%v) = %v; want %v", tc.expression, err, tc.want)
			}
			if !equal(got, tc.want) {
				t.Errorf("ShuntingYard(%v) = %v; want %v", tc.expression, got, tc.want)
			}
		})
	}
}

// Функция для сравнения двух срезов строк
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "expression with +",
			expression:     "4+2",
			expectedResult: 6,
		},
		{
			name:           "expression with -",
			expression:     "10-2",
			expectedResult: 8,
		},
		{
			name:           "expression with *",
			expression:     "8*4",
			expectedResult: 32,
		},
		{
			name:           "expression with /",
			expression:     "6/3",
			expectedResult: 2,
		},
		{
			name:           "expression with *, +, -, ()",
			expression:     "3+4*(2-1)",
			expectedResult: 7,
		},
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "missing operand at the end",
			expression:  "1+1*",
			expectedErr: errors.New("incorrect expression"),
		},
		{
			name:        "invalid operator sequence",
			expression:  "2+2**2",
			expectedErr: errors.New("incorrect expression"),
		},
		{
			name:        "invalid characters with parenthesis",
			expression:  "((2+2-*(2",
			expectedErr: errors.New("incorrect expression"),
		},
		{
			name:        "err divide by zero",
			expression:  "6/0",
			expectedErr: errors.New("err divide by zero"),
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := Calc(testCase.expression)
			if err == nil {
				t.Fatalf("fail case %s should return error", testCase.expression)
			}
			if err.Error() != testCase.expectedErr.Error() {
				t.Fatalf("expected error %v, got %v", testCase.expectedErr, err)
			}
		})
	}
}
