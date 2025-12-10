package mapcli

import (
    "os"
    "fmt"
	"path/filepath"
    "gopkg.in/yaml.v3"
)


func isErr(e error) bool {
	if e != nil { return true }
	return false
}

var home, _ = os.UserHomeDir()
var MapCliDir = filepath.Join(home, ".mapcli")
var MapCliBin = filepath.Join(MapCliDir, "bin")
var MapCliMetadata = filepath.Join(MapCliDir, "metadata")
var MapCliMappings = filepath.Join(MapCliDir, "mappings")


type MappingMissingError struct {
	CommandName string
	ErrorMsg string
}

func (e *MappingMissingError) Error() string {
	return fmt.Sprintf("No mapping found for command %s; Error: %s", e.CommandName, e.ErrorMsg)
}

func MakeMapcliDirs() error {
	err := os.MkdirAll(MapCliDir, 0755)
	if isErr(err) { return err }
	err = os.MkdirAll(MapCliBin, 0777)
	if isErr(err) { return err }
	err = os.MkdirAll(MapCliMappings, 0755)
	return err
}

func GetMetadataPath(command string) string {
	yamlPath := fmt.Sprintf("%s.yaml", command)
    return filepath.Join(MapCliMetadata, yamlPath)
}

func GetMappingPath(command string) string {
	yamlPath := fmt.Sprintf("%s.yaml", command)
    return filepath.Join(MapCliMappings, yamlPath)
}

func GetExecutablePath(command string) string {
    return filepath.Join(MapCliBin, command)
}

func ExecutableExists(command string) bool {
	path := GetExecutablePath(command)
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	} else { panic(err) }
}

func ReadMapping(commandName string) (map[string]string, error) {
	path := GetMappingPath(commandName)

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return nil, &MappingMissingError{CommandName: commandName, ErrorMsg: err.Error()}
	} else if err != nil {
		return nil, err
	}

    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    m := make(map[string]string)

    if err := yaml.Unmarshal(data, &m); err != nil {
        return nil, err
    }

    return m, nil
}


func MapArgs(mapping map[string]string, args []string) []string {
    newArgs := make([]string, len(args), cap(args))
    for i, a := range args {
        val, found := mapping[a]
        if found {
            newArgs[i] = val
        } else {
            newArgs[i] = a
        }
    }
    return newArgs
}



