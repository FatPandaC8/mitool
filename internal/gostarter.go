package internal

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Node struct {
	Name     string  `yaml:"name"`
	Type     string  `yaml:"type"`
	Children []*Node `yaml:"children,omitempty"`
}

type Template struct {
	Name      string   `yaml:"name"`
	Structure []*Node  `yaml:"structure"`
	Files     []string `yaml:"files"`
}

func LoadTemplate(path string) (*Template, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tmpl Template

	err = yaml.Unmarshal(data, &tmpl)
	if err != nil {
		return nil, err
	}

	return &tmpl, nil
}

func CreateNode(base string, node *Node) error {
	path := filepath.Join(base, node.Name)

	switch node.Type {
	case "directory":
		// mode perm is linux full permission
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}

		for _, child := range node.Children {
			err := CreateNode(path, child)
			if err != nil {
				return err
			}
		}
	case "file":
		file, err := os.Create(path)
		if err != nil {
			return err
		}

		defer file.Close()
	}

	return nil
}

func CreateProject(projectName string, template *Template) error {
	err := os.MkdirAll(projectName, os.ModePerm)
	if err != nil {
		return err
	}

	for _, node := range template.Structure {
		err := CreateNode(projectName, node)
		if err != nil {
			return err
		}
	}

	for _, file := range template.Files {
		path := filepath.Join(projectName, file)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
