package nonoslv

import "fmt"

// 渡されたファイルのヒントを元にノノグラムを解く
// どのように解いたかはステップ毎に確認が出来る
// - 可能であれば描画したい
// 最終的な結果が描画される（画像じゃなくて文字でもいいが）

type Stage struct {
	width  int
	height int
	cells  [][]Cell
}

func (s *Stage) GetLineVertical(x int) Line {
	cells := make([]Cell, s.height)
	for y := 0; y < s.height; y++ {
		cells[y] = s.GetCell(x, y)
	}
	return cells
}

func (s *Stage) GetLineHorizontal(y int) Line {
	cells := make([]Cell, s.width)
	for x := 0; x < s.height; x++ {
		cells[x] = s.GetCell(x, y)
	}
	return cells
}

func (s *Stage) GetCell(x int, y int) Cell {
	return s.cells[y][x]
}

func (s *Stage) Diff(one *Stage) []Cell {
	var cells []Cell
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			l := s.GetCell(x, y)
			r := one.GetCell(x, y)
			if l.State == r.State {
				continue
			}
			cells = append(cells, r)
		}
	}
	return cells
}

func (s *Stage) Print() {
	for y := 0; y < s.height; y++ {
		row := ""
		for x := 0; x < s.width; x++ {
			switch s.cells[y][x].State {
			case None:
				row += "-"
			case Cross:
				row += "x"
			case Fill:
				row += "o"
			}
		}
		fmt.Println(row)
	}
}

func NewInitialStage(width int, height int) *Stage {
	cells := make([][]Cell, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]Cell, width)
		for x := 0; x < width; x++ {
			cells[y][x] = Cell{
				X: x,
				Y: y,
				State: None,
			}
		}
	}
	return &Stage{width: width, height: height, cells: cells}
}

func CopyStage(stage *Stage) *Stage {
	var cells [][]Cell
	for y := 0; y < stage.height; y++ {
		line := stage.GetLineHorizontal(y)
		line.Copy()
		cells = append(cells, line)
	}
	return &Stage{
		width:  stage.width,
		height: stage.height,
		cells:  cells,
	}
}

type CellState int

const (
	None = iota
	Cross
	Fill
)

type Cell struct {
	X     int
	Y     int
	State CellState
}

func (c *Cell) Fill() bool {
	if c.State == Cross {
		return false
	}
	c.State = Fill
	return true
}

func (c *Cell) Cross() bool {
	if c.State == Fill {
		return false
	}
	c.State = Cross
	return true
}

func (c *Cell) None() {
	c.State = None
}

type Line []Cell

func (l Line) Copy() Line {
	dst := make([]Cell, len(l))
	for i := 0; i < len(l); i++ {
		dst[i] = Cell{
			State: l[i].State,
		}
	}
	return dst
}

func (l Line) FillRange(begin int, n int) (Line, bool) {
	backup := l.Copy()
	end := begin + n
	if end > len(l) {
		end = len(l)
	}
	for i := begin; i < end; i++ {
		if !l[i].Fill() {
			// 埋められないなら元に戻す
			return backup, false
		}
	}
	return l, true
}

func (l Line) CrossRange(begin int, n int) (Line, bool) {
	// 塗る潰す数が0なら何もせずに終了
	if n <= 0 {
		return l, true
	}

	backup := l.Copy()
	end := begin + n
	for i := begin; i < end; i++ {
		if !l[i].Cross() {
			// 埋められないなら元に戻す
			return backup, false
		}
	}
	return l, true
}

func (l Line) GetStates() []CellState {
	var states []CellState
	for _, cell := range l {
		states = append(states, cell.State)
	}
	return states
}

type History struct {
	stages []Stage
}

func (h *History) Add(stage Stage) {
	h.stages = append(h.stages, stage)
}

func (h *History) Print() {
	for step, stage := range h.stages {
		fmt.Printf("step:%d\n", step)
		stage.Print()
	}
}

func Solve(input *Input) (*History, error) {
	history := new(History)
	initial := NewInitialStage(input.width, input.height)
	history.Add(*initial)

	next := initial
	var tv []int
	var th []int
	for x := 0; x < initial.width; x++ {
		tv = append(tv, x)
	}
	for y := 0; y < initial.height; y++ {
		th = append(th, y)
	}

	for {
		cur := CopyStage(next)
		next = searchFixedCell(input, cur, tv, th)
		diff := cur.Diff(next)
		if len(diff) == 0 {
			break
		}
		history.Add(*next)
		tv, th = getChangedLine(diff)
	}
	return history, nil
}

func searchFixedCell(input *Input, stage *Stage, targetsV []int, targetsH []int) *Stage {
	fv := make([]bool, stage.width)
	fh := make([]bool, stage.height)
	for _, x := range targetsV {
		fv[x] = true
	}
	for _, y := range targetsH {
		fh[y] = true
	}
	// vertical
	var vLines []Line
	for x := 0; x < stage.width; x++ {
		line := stage.GetLineVertical(x)
		if fv[x] {
			line = searchCombination(input.vHintsGroup[x], line)
		}
		vLines = append(vLines, line)
	}
	// horizontal
	var hLines []Line
	for y := 0; y < stage.height; y++ {
		line := stage.GetLineHorizontal(y)
		if fh[y] {
			line = searchCombination(input.hHintsGroup[y], line)
		}
		hLines = append(hLines, line)
	}

	cells := make([][]Cell, stage.height)
	for y := 0; y < stage.height; y++ {
		cells[y] = make([]Cell, stage.width)
		for x := 0; x < stage.width; x++ {
			var state CellState
			switch {
			case vLines[x][y].State == Fill || hLines[y][x].State == Fill:
				state = Fill
			case vLines[x][y].State == Cross || hLines[y][x].State == Cross:
				state = Cross
			default: // TODO: ここはありえないのでerrorを返すように
				state = None
			}
			cells[y][x] = Cell{
				X: x,
				Y: y,
				State: state,
			}
		}
	}
	return &Stage{
		width:  stage.width,
		height: stage.height,
		cells:  cells,
	}
}

func searchCombination(hints Hints, line Line) Line {
	var initialState [][]CellState
	cmb := search(0, 0, hints, line, initialState)
	return mergeFinalStates(cmb, line)
}

func search(step int, cur int, hints Hints, line Line, cmb [][]CellState) [][]CellState {
	if step >= len(hints) {
		i := cur-1
		if i < 0 {
			i = 0 // hintsなしの場合の対応
		}
		// 最後まで行ったら末尾まで☓をつける
		line, ok := line.CrossRange(i, len(line)-i)
		if !ok {
			// ただし、すでに塗りつぶされていれば何もしない
			return cmb
		}
		return append(cmb, line.GetStates())
	}

	hint := hints[step]
	for i := cur; i < len(line); i++ {
		// ヒントの数よりもマスの上限を超えた場合は何もしない
		if i+hint > len(line) {
			return cmb
		}

		l := line.Copy()
		// 確定マスがあって塗りつぶせないなら次のマスを試す
		l, ok := l.CrossRange(cur, i - cur)
		if !ok {
			continue
		}
		l, ok = l.FillRange(i, hint)
		if !ok {
			continue
		}
		if i+hint < len(l) {
			l, ok = l.CrossRange(i+hint, 1)
			if !ok {
				continue
			}
		}
		cmb = search(step+1, i+hint+1, hints, l, cmb)
	}
	return cmb
}

func mergeFinalStates(lineStates [][]CellState, fallbackLine Line) Line {
	if len(lineStates) == 0 {
		return fallbackLine
	}
	size := len(lineStates[0])
	fillBits := make([]bool, size)
	crossBits := make([]bool, size)
	for i := 0; i < size; i++ {
		fillBits[i] = true
		crossBits[i] = true
	}
	for _, states := range lineStates {
		for j, state := range states {
			switch state {
			case Fill:
				crossBits[j] = false
			case Cross:
				fillBits[j] = false
			case None: // Noneのパターンあるのか？
				crossBits[j] = false
				fillBits[j] = false
			}
		}
	}
	line := Line(make([]Cell, size))
	for i, _ := range line {
		if fillBits[i] {
			line.FillRange(i, 1)
			continue
		}
		if crossBits[i] {
			line.CrossRange(i, 1)
		}
	}
	return line
}

func getChangedLine(diff []Cell) ([]int, []int) {
	fx := map[int]bool{}
	fy := map[int]bool{}
	var tv []int
	var th []int
	for _, cell := range diff {
		if !fx[cell.X] {
			fx[cell.X] = true
			tv = append(tv, cell.X)
		}
		if !fy[cell.Y] {
			fy[cell.Y] = true
			th = append(th, cell.Y)
		}
	}
	return tv, th
}
