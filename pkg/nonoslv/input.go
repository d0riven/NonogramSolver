package nonoslv

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Hints []int

type RawInput struct {
	Size struct {
		Width  int
		Height int
	}
	Hints struct {
		HorizontalHintGroups [][]int `yaml:"horizontal"`
		VerticalHintGroups   [][]int `yaml:"vertical"`
	}
}

type Input struct {
	width       int
	height      int
	hHintsGroup []Hints
	vHintsGroup []Hints
}

func NewInput(raw RawInput) (*Input, error) {
	if err := validationHints(raw.Hints.HorizontalHintGroups, raw.Size.Width, "horizontal"); err != nil {
		return nil, err
	}
	if err := validationHints(raw.Hints.VerticalHintGroups, raw.Size.Height, "vertical"); err != nil {
		return nil, err
	}
	vHintGroups := make([]Hints, raw.Size.Height)
	for _, rawGroup := range raw.Hints.VerticalHintGroups {
		vHintGroups = append(vHintGroups, rawGroup)
	}
	hHintGroups := make([]Hints, raw.Size.Width)
	for _, rawGroup := range raw.Hints.VerticalHintGroups {
		hHintGroups = append(hHintGroups, rawGroup)
	}
	return &Input{
		width:       raw.Size.Width,
		height:      raw.Size.Height,
		hHintsGroup: hHintGroups,
		vHintsGroup: vHintGroups,
	}, nil
}

func validationHints(hintsGroups [][]int, size int, direction string) error {
	for i, hHints := range hintsGroups {
		sum := 0
		for _, hv := range hHints {
			sum += hv
		}
		sum += len(hHints) - 1
		if sum > size {
			return fmt.Errorf("sum hints over size. order:%d, direction:%s", i, direction)
		}
	}
	return nil
}

func ReadFromYaml(path string) (*Input, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	var raw RawInput
	if err := yaml.Unmarshal(b, &raw); err != nil {
		return nil, err
	}
	return NewInput(raw)
}
