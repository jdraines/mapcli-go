package mapcli

import (
    "io"
	"os"
	"fmt"
)


func CreateMappedCli(srcMappingPath string, commandName string, makeCopy bool) {

	MakeMapcliDirs()
	
	if makeCopy {
		err := copyMappingToMapcliMappings(srcMappingPath, commandName)
		if err != nil { panic(err) }
	} else {
		err := linkMappingToMapcliMappings(srcMappingPath, commandName)
		if err != nil { panic(err) }
	}
	writeScript(commandName)
}

func linkMappingToMapcliMappings(srcPath string, commandName string) error {
	linkPath := GetMappingPath(commandName)
	err := os.Symlink(srcPath, linkPath)
	return err
}


func copyMappingToMapcliMappings(srcPath string, commandName string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil { return err }
	defer srcFile.Close()

	dstPath := GetMappingPath(commandName)
	if err != nil { return err }
	
	dstFile, err := os.Open(dstPath)
	if err != nil { return err }
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil { return err }

	err = dstFile.Sync()
	if err != nil { return err }

	return nil
}

func writeScript(commandName string) error {
	scriptText := fmt.Sprintf("mapcli run %s $@\n", commandName)
	scriptPath := GetExecutablePath(commandName)
	
	scriptFile, err := os.Create(scriptPath)
	if err != nil { panic(err) }
	defer scriptFile.Close()

	_, err = scriptFile.WriteString(scriptText)
	if err != nil { panic(err) }
	err = scriptFile.Sync()
	scriptFile.Chmod(0777)
	return err
}
