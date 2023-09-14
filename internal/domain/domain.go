package domain

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Topology struct {
	Name       string      `yaml:"name"` // Topology name (used to differentiate between topologies)
	Components []Component `yaml:"components"`
	Relations  []Relation  `yaml:"relations"`
}

type Component struct {
	Name        string            `yaml:"name"`
	Type        string            `yaml:"type"`
	Labels      map[string]string `yaml:"labels"`
	Domain      string            `yaml:"domain"`
	Environment string            `yaml:"environment"`
	Layer       string            `yaml:"layer"`
}

type Relation struct {
	Source string `yaml:"source"`
	Target string `yaml:"target"`
	Type   string `yaml:"type"`
}

func ReadTopology(file string) (*Topology, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	t := &Topology{}
	if err := yaml.Unmarshal(b, t); err != nil {
		return nil, err
	}

	return t, nil
}
