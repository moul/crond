package crond

import (
	"fmt"
	"strings"
)

type Manager struct {
	rules Rules
}

func New() Manager {
	return Manager{
		rules: make(Rules, 0),
	}
}

func (m *Manager) Start() {
	fmt.Println("START")
}

func (m *Manager) AddLine(line string) error {
	rule, err := parseCronLine(line)
	if err != nil {
		return err
	}
	m.rules = append(m.rules, rule)
	return nil
}

func (m *Manager) ParseCrontab(content string) error {
	lines := strings.Split(content, "\n")

	rules := Rules{}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line[0] == '#' {
			continue
		}

		rule, err := parseCronLine(line)
		if err != nil {
			return err
		}
		rules = append(rules, rule)
	}

	m.rules = rules
	return nil
}

func (m *Manager) Rules() Rules { return m.rules }
