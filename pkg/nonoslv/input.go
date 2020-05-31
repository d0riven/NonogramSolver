package nonoslv

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type HintsGroup [][]int

type RawInput struct {
	Size struct {
		Width  int
		Height int
	}
	Hints struct {
		HorizontalHintsGroups [][]int `yaml:"horizontal"`
		VerticalHintsGroups   [][]int `yaml:"vertical"`
	}
}

type Input struct {
	width        int
	height       int
	hHintsGroups HintsGroup
	vHintsGroups HintsGroup
}

func NewInput(raw RawInput) (*Input, error) {
	if err := validationHints(raw.Hints.HorizontalHintsGroups, raw.Size.Width, "horizontal"); err != nil {
		return nil, err
	}
	if err := validationHints(raw.Hints.VerticalHintsGroups, raw.Size.Height, "vertical"); err != nil {
		return nil, err
	}
	return &Input{
		width:        raw.Size.Width,
		height:       raw.Size.Height,
		hHintsGroups: raw.Hints.HorizontalHintsGroups,
		vHintsGroups: raw.Hints.VerticalHintsGroups,
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
