package nonoslv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchCombination(t *testing.T) {
	type args struct {
		hints Hints
		line  Line
	}
	t.Run("initial state (-----)", func(t *testing.T) {
		tests := []struct {
			name string
			args args
			want Line
		}{
			{
				name: "hints(None)=(-----)=>(-----)",
				args: args{
					hints: Hints{},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
				},
			},
			{
				name: "hints(1)=(-----)=>(-----)",
				args: args{
					hints: Hints{1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
				},
			},
			{
				name: "hints(2)=(-----)=>(-----)",
				args: args{
					hints: Hints{2},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
				},
			},
			{
				name: "hints(3)=(-----)=>(--o--)",
				args: args{
					hints: Hints{3},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: Fill},
					Cell{State: None},
					Cell{State: None},
				},
			},
			{
				name: "hints(4)=(-----)=>(-ooo-)",
				args: args{
					hints: Hints{4},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: None},
				},
			},
			{
				name: "hints(5)=(-----)=>(ooooo)",
				args: args{
					hints: Hints{5},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
				},
			},
			{
				name: "hints(1,1)=(-----)=>(-----)",
				args: args{
					hints: Hints{1, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
				},
			},
			{
				name: "hints(1,2)=(-----)=>(---o-)",
				args: args{
					hints: Hints{1, 2},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: Fill},
					Cell{State: None},
				},
			},
			{
				name: "hints(1,3)=(-----)=>(oxooo)",
				args: args{
					hints: Hints{1, 3},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
				},
			},
			{
				name: "hints(2,1)=(-----)=>(-o---)",
				args: args{
					hints: Hints{2, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: Fill},
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
				},
			},
			{
				name: "hints(2,2)=(-----)=>(ooxoo)",
				args: args{
					hints: Hints{2, 2},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
					Cell{State: Fill},
				},
			},
			{
				name: "hints(3,1)=(-----)=>(oooxo)",
				args: args{
					hints: Hints{3, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
				},
			},
			{
				name: "hints(1,1,1)=(-----)=>(oxoxo)",
				args: args{
					hints: Hints{1, 1, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := searchCombination(tt.args.hints, tt.args.line)
				assert.Equal(t, tt.want, got)
			})
		}
	})
	t.Run("initial state (---x-)", func(t *testing.T) {
		tests := []struct {
			name string
			args args
			want Line
		}{
			{
				name: "hints(None)=(---x-)=>(---x-)",
				args: args{
					hints: Hints{1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: Cross},
					Cell{State: None},
				},
			},
			{
				name: "hints(1)=(---x-)=>(---x-)",
				args: args{
					hints: Hints{1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: Cross},
					Cell{State: None},
				},
			},
			{
				name: "hints(2)=(---x-)=>(-o-xx)",
				args: args{
					hints: Hints{2},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: Fill},
					Cell{State: None},
					Cell{State: Cross},
					Cell{State: Cross},
				},
			},
			{
				name: "hints(3)=(---x-)=>(oooxx)",
				args: args{
					hints: Hints{3},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Cross},
				},
			},
			{
				name: "hints(1,1)=(---x-)=>(---x-)",
				args: args{
					hints: Hints{1, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: None},
					Cell{State: None},
					Cell{State: Cross},
					Cell{State: None},
				},
			},
			{
				name: "hints(2,1)=(---x-)=>(-o-xo)",
				args: args{
					hints: Hints{2, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: None},
					Cell{State: Fill},
					Cell{State: None},
					Cell{State: Cross},
					Cell{State: Fill},
				},
			},
			{
				name: "hints(3,1)=(---x-)=>(oooxo)",
				args: args{
					hints: Hints{3, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
				},
			},
			{
				name: "hints(1,1,1)=(---x-)=>(oxoxo)",
				args: args{
					hints: Hints{1, 1, 1},
					line: Line{
						Cell{State: None},
						Cell{State: None},
						Cell{State: None},
						Cell{State: Cross},
						Cell{State: None},
					},
				},
				want: Line{
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
					Cell{State: Cross},
					Cell{State: Fill},
				},
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := searchCombination(tt.args.hints, tt.args.line)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
