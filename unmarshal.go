package makefile

import (
	"bufio"
	"io"
	"regexp"
)

var targetRegex = regexp.MustCompile(`^(?P<Target>[\w.]+):`)

var targetExcludes = map[string]bool{
	".PHONY": true,
}

func Unmarshal(r io.Reader) (Makefile, error) {
	m := Makefile{}

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		l := s.Text()
		matches := targetRegex.FindStringSubmatch(l)
		if len(matches) > 0 {
			tn := matches[1]
			if targetExcludes[tn] {
				continue
			}
			t := Target{Name: tn}
			m.Targets = append(m.Targets, t)
		}
	}
	return m, nil
}
