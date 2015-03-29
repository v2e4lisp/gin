package gin

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

var (
	watchPatterns    = []string{}
	WatchPatternFile = ".ginwatch"
)

func IsWatched(path string) bool {
	// ignore hidden files
	if filepath.Base(path)[0] == '.' {
		return false
	}

	if filepath.Ext(path) == ".go" {
		return true
	}

	for _, pat := range watchPatterns {
		m, err := filepath.Match(pat, path)
		if err == nil && m {
			return true
		}
	}

	return false
}

func LoadWatchPatterns(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		pat := strings.TrimSpace((sc.Text()))
		if pat == "" || pat[0] == '#' {
			continue
		}

		watchPatterns = append(watchPatterns, pat)
	}

	return nil
}
