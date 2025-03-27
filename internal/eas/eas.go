package eas

import (
    "fmt"
    "regexp"
    "strings"
    "bufio"
)


type ProjectInfo struct {
	id    string `json:"id"`
    name  string `json:"name"`
	owner string `json:"owner"`
}

type ProjectVariableProps struct {
    name        string `json:"name"`
    value       string `json:"value"`
    visibility  string `json:"visibility"`
    environment string `json:"environment"`
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
			props.name = value
		case "Value":
			props.value = value
		case "Visibility":
			lowerVisibility := strings.ToLower(value)
			if lowerVisibility == "public" {
				props.visibility = "plaintext"
			} else {
				props.visibility = lowerVisibility
			}
		case "Environments":
			props.environment = strings.ToLower(value)
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
    	owner: matches[1],
    	name:  matches[2],
    	id:    matches[3],
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

func createProject(name string) (*ProjectInfo, error) {
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

func createProjectVariable(projectName string, props ProjectVariableProps) (*ProjectVariableProps, error) {
    linkProject(projectName);
    _, err := RunCommand(projectName, "eas", "env:create", "--scope", "project", "--non-interactive", "--name", props.name, "--value", props.value, "--visibility", props.visibility, "--environment", props.environment)
    if err != nil {
        deleteContext(projectName)
        return nil, err
    }
    deleteContext(projectName)
    return &props, err
}

func deleteProjectVariable(projectName string, variableName string, environment string) (string, error) {
    linkProject(projectName)
    out, err := RunCommand(projectName, "eas", "env:delete", environment, "--variable-name", variableName, "--non-interactive", "--scope", "project")
    deleteContext(projectName)
    return out, err
}

func getProjectVariable(projectName string, variableName string, environment string) (*ProjectVariableProps, error) {
    linkProject(projectName)
    out, err := RunCommand(projectName, "eas", "env:get", environment, "--variable-name", variableName, "--non-interactive", "--format", "long", "--scope", "project")
    if err != nil {
        return nil, err
    }
deleteContext(projectName)
    return parseVariable(out)
}
