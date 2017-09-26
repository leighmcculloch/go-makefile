package makefile

import (
	"bufio"
	"io"
	"regexp"
)

var targetRegex = regexp.MustCompile(`(?P<Target>\w+):`)

func Unmarshal(r io.Reader) (Makefile, error) {
	m := Makefile{}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		l := s.Text()
		matches := targetRegex.FindStringSubmatch(l)
		if len(matches) > 0 {
			t := Target{Name: matches[1]}
			m.Targets = append(m.Targets, t)
		}
	}
	return m, nil
}
