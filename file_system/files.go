package file_system

import "os"

func WriteToFile(path string, content []byte) error {
	return os.WriteFile(path, content, 0644)
}

func WriteStringToFile(path string, content string) error {
	return WriteToFile(path, []byte(content))
}

func ReadStringInFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	return string(content), nil
}
func DeleteFile(path string) error {
	if _, err := os.Stat(path); err != nil {
		return nil
	}

	return os.Remove(path)
}
