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
				name: "hints(None)=(-----)=>(xxxxx)",
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
					Cell{State: Cross},
					Cell{State: Cross},
					Cell{State: Cross},
					Cell{State: Cross},
					Cell{State: Cross},
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

func Test_searchFixedCell(t *testing.T) {
	type args struct {
		input *Input
		stage *Stage
		targetV []int
		targetH []int
	}
	tests := []struct {
		name string
		args args
		want *Stage
	}{
		{
			name: "",
			args: args{
				input: &Input{
					width:  5,
					height: 5,
					hHintsGroup: []Hints{
						{5},
						{1},
						{5},
						{1},
						{5},
					},
					vHintsGroup: []Hints{
						{3, 1},
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
						{1, 3},
					},
				},
				stage:   NewInitialStage(5, 5),
				targetV: []int{0,1,2,3,4},
				targetH: []int{0,1,2,3,4},
			},
			want: &Stage{
				width:  5,
				height: 5,
				// ooooo
				// oxxxx
				// ooooo
				// xxxxo
				// ooooo
				cells: [][]Cell{
					{
						{X: 0, Y: 0, State: Fill},
						{X: 1, Y: 0, State: Fill},
						{X: 2, Y: 0, State: Fill},
						{X: 3, Y: 0, State: Fill},
						{X: 4, Y: 0, State: Fill},
					},
					{
						{X: 0, Y: 1, State: Fill},
						{X: 1, Y: 1, State: Cross},
						{X: 2, Y: 1, State: Cross},
						{X: 3, Y: 1, State: Cross},
						{X: 4, Y: 1, State: Cross},
					},
					{
						{X: 0, Y: 2, State: Fill},
						{X: 1, Y: 2, State: Fill},
						{X: 2, Y: 2, State: Fill},
						{X: 3, Y: 2, State: Fill},
						{X: 4, Y: 2, State: Fill},
					},
					{
						{X: 0, Y: 3, State: Cross},
						{X: 1, Y: 3, State: Cross},
						{X: 2, Y: 3, State: Cross},
						{X: 3, Y: 3, State: Cross},
						{X: 4, Y: 3, State: Fill},
					},
					{
						{X: 0, Y: 4, State: Fill},
						{X: 1, Y: 4, State: Fill},
						{X: 2, Y: 4, State: Fill},
						{X: 3, Y: 4, State: Fill},
						{X: 4, Y: 4, State: Fill},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchFixedCell(tt.args.input, tt.args.stage, tt.args.targetV, tt.args.targetH)
			if !assert.Equal(t, tt.want, got) {
				got.Print()
			}
		})
	}
}
