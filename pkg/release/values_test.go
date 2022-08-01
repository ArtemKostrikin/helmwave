package release_test

import (
	"testing"

	"github.com/helmwave/helmwave/pkg/release"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type ValuesTestSuite struct {
	suite.Suite
}

func (s *ValuesTestSuite) TestList() {
	type config struct {
		Values []release.ValuesReference
	}

	src := `
values:
- a
- b
`
	c := &config{}

	err := yaml.Unmarshal([]byte(src), c)
	s.Require().NoError(err)

	s.Require().Equal(&config{
		Values: []release.ValuesReference{
			{Src: "a", Render: true},
			{Src: "b", Render: true},
		},
	}, c)
}

func (s *ValuesTestSuite) TestMap() {
	type config struct {
		Values []release.ValuesReference
	}

	src := `
values:
- src: 1
  render: false
- src: 2
  strict: true
`
	c := &config{}

	err := yaml.Unmarshal([]byte(src), c)
	s.Require().NoError(err)

	s.Require().Equal(&config{
		Values: []release.ValuesReference{
			{Src: "1", Render: false},
			{Src: "2", Strict: true},
		},
	}, c)
}

func TestValuesTestSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ValuesTestSuite))
}
