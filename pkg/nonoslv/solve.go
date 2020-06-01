package nonoslv

import "github.com/ulule/deepcopier"

// 渡されたファイルのヒントを元にノノグラムを解く
// どのように解いたかはステップ毎に確認が出来る
// - 可能であれば描画したい
// 最終的な結果が描画される（画像じゃなくて文字でもいいが）

type Stage struct {
	width  int
	height int
	cells  [][]Cell
}

//func (s *Stage) Width() int {
//	return s.width
//}
//
//func (s *Stage) Height() int {
//	return s.height
//}

func (s *Stage) GetCell(x int, y int) Cell {
	return s.cells[y][x]
}

func NewInitialStage(width int, height int) *Stage {
	cells := make([][]Cell, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]Cell, width)
		for x := 0; x < width; x++ {
			cells[y][x] = Cell{
				X: x,
				Y: y,
			}
		}
	}
	return &Stage{width: width, height: height, cells: cells}
}

func CopyStage(stage *Stage) *Stage {
	dst := Stage{}
	deepcopier.Copy(dst).From(*stage)
	return &dst
}

type Cell struct {
	X int
	Y int
}

type History struct {
	stages []Stage
}

func (h *History) Add(stage Stage) {
	h.stages = append(h.stages, stage)
}

func Solve(input *Input) (*History, error) {
	history := new(History)
	initial := NewInitialStage(input.width, input.height)
	history.Add(*initial)

	next := initial
	for {
		cur := CopyStage(next)
		next, err := searchFixedCell(input, cur)
		if err != nil {
			return nil, err
		}
		if next == nil {
			break
		}
		history.Add(*next)
	}
	return history, nil
}

func searchFixedCell(input *Input, stage *Stage) (*Stage, error) {
	return nil, nil
}
