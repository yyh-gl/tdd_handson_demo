package fizzbuzz

import (
	"testing"

	// "github.com/smartystreets/assertions/assert"
	"github.com/stretchr/testify/assert"
)

func Test_FizzBuzz(t *testing.T) {
	testCases := []struct {
		name  string
		input int
		want  string
	}{
		{
			name:  "1を送ったら1を返す",
			input: 1,
			want:  "1",
		},
		{
			name:  "2を送ったら2を返す",
			input: 2,
			want:  "2",
		},
		{
			name:  "3を送ったらFizzを返す",
			input: 3,
			want:  "Fizz",
		},
		{
			name:  "6を送ったらFizzを返す",
			input: 6,
			want:  "Fizz",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, FizzBuzz(test.input))
		})
	}
}
