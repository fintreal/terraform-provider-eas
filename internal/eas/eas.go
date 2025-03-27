package eas

import (
    "fmt"
    "regexp"
    "strings"
    "bufio"
)


type ProjectInfo struct {
	Id    string `json:"id"`
    Name  string `json:"name"`
	Owner string `json:"owner"`
}

type ProjectVariableProps struct {
    Name        string `json:"name"`
    Value       string `json:"value"`
    Visibility  string `json:"visibility"`
    Environment string `json:"environment"`
}

func parseVariable(input string) (*ProjectVariableProps, error) {
	var props ProjectVariableProps

	// Create a scanner to read line by line
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "  ", 2) // Split by double space (handles variable spacing)
		if len(parts) < 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Map extracted values to the struct fields
		switch key {
		case "Name":
			props.Name = value
		case "Value":
			props.Value = value
		case "Visibility":
			lowerVisibility := strings.ToLower(value)
			if lowerVisibility == "public" {
				props.Visibility = "plaintext"
			} else {
				props.Visibility = lowerVisibility
			}
		case "Environments":
			props.Environment = strings.ToLower(value)
		}
	}

	return &props, nil
}

func parseProjectInfo(input string) (*ProjectInfo, error) {
    re := regexp.MustCompile(`fullName\s+@([^/]+)/([^ \n]+)\nID\s+([a-f0-9-]+)`)

    matches := re.FindStringSubmatch(input)

    if len(matches) < 4 {
    	return nil, fmt.Errorf("Failed to parse project info")
    }

    return &ProjectInfo{
    	Owner: matches[1],
    	Name:  matches[2],
    	Id:    matches[3],
    }, nil
}

func createContext(name string) {
    createDir(name)
    createFile(name, "package.json", "{}")
    content := fmt.Sprintf(`{ "name": "%s", "slug": "%s" }`, name, name)
    createFile(name, "app.json", content)
}

func deleteContext(name string) {
    RunCommand(".", "rm", "-rf", name)
}

func CreateProject(name string) (*ProjectInfo, error) {
    createContext(name)
    out, err := RunCommand(name, "eas", "project:init", "--force")
    if err != nil {
        deleteContext(name)
        return nil, err
    }
    out, err = RunCommand(name, "eas", "project:info")
    deleteContext(name)

    if err !=nil {
        return nil, err
    }

    return parseProjectInfo(out)
}

func linkProject(name string) {
    createContext(name)
    RunCommand(name, "eas", "project:init", "--force")
}

func CreateProjectVariable(projectName string, props ProjectVariableProps) (*ProjectVariableProps, error) {
    linkProject(projectName);
    _, err := RunCommand(projectName, "eas", "env:create", "--scope", "project", "--non-interactive", "--name", props.Name, "--value", props.Value, "--visibility", props.Visibility, "--environment", props.Environment)
    if err != nil {
        deleteContext(projectName)
        return nil, err
    }
    deleteContext(projectName)
    return &props, err
}

func DeleteProjectVariable(projectName string, variableName string, environment string) (string, error) {
    linkProject(projectName)
    out, err := RunCommand(projectName, "eas", "env:delete", environment, "--variable-name", variableName, "--non-interactive", "--scope", "project")
    deleteContext(projectName)
    return out, err
}

func GetProjectVariable(projectName string, variableName string, environment string) (*ProjectVariableProps, error) {
    linkProject(projectName)
    out, err := RunCommand(projectName, "eas", "env:get", environment, "--variable-name", variableName, "--non-interactive", "--format", "long", "--scope", "project")
    if err != nil {
        return nil, err
    }
    deleteContext(projectName)
    return parseVariable(out)
}
