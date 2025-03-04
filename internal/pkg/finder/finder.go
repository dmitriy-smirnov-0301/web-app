package finder

import (
	"errors"
	"ice-creams-app/internal/pkg/logger"
	"os"
	"path/filepath"
)

func FindConfigsDir(defaultFolder, configFileName string) (string, error) {

	log := logger.GetLogger()

	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %s", err)
		return "", err
	}
	log.Infof("Working directory: %v", workingDir)

	projectRoot, err := findProjectRoot(workingDir, "cmd", "configs", "internal", "app")
	if err != nil {
		log.Errorf("Error: %s", err)
	}
	log.Infof("Root direcrory found: %s", projectRoot)

	configPath, err := findConfigFile(projectRoot, defaultFolder, configFileName)
	if err != nil {
		log.Errorf("Error: %s", err)
		return "", err
	}
	log.Infof("Config file found: %s", configPath)

	return configPath, nil

}

func isProjectRoot(dir string, rootMarkers ...string) bool {

	for _, marker := range rootMarkers {
		if _, err := os.Stat(filepath.Join(dir, marker)); err == nil {
			return true
		}
	}

	return false

}

func findProjectRoot(startDir string, rootMarkers ...string) (string, error) {

	currentDir := startDir

	for {
		if isProjectRoot(currentDir, rootMarkers...) {
			return currentDir, nil
		}

		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", errors.New("project root not found")
		}
		currentDir = parentDir
	}

}

func findConfigFile(root, defaultFolder, fileName string) (string, error) {

	defaultConfigPath := filepath.Join(root, defaultFolder, fileName)

	if _, err := os.Stat(defaultConfigPath); err == nil {
		return defaultConfigPath, nil
	}

	foundPaths := make([]string, 0, 1)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && info.Name() == fileName {
			foundPaths = append(foundPaths, path)
			if len(foundPaths) > 1 {
				return errors.New("multiple files found")
			}
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if len(foundPaths) == 0 {
		return "", errors.New("file not found")
	}

	return foundPaths[0], nil

}
