package webpack_sass

import (
	"fmt"
	"os/exec"
)

// Config contains all data
type Config struct {
}

// Initalize is the first function that should be executed
func Initalize() *Config {
	return &Config{}
}
func (c *Config) InitWebpack() error {
	cmd := exec.Command("npm",
		[]string{
			"run",
			"--prefix",
			"D:/Development/GO/admin/public/js/src/",
			"start",
		}...)

	bytes, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}

func (c *Config) InitSass() error {
	cmd := exec.Command("sass",
		[]string{
			"input.scss",
			"output.css",
			"--watch",
		}...)

	bytes, err := cmd.Output()

	if err != nil {
		return err
	}

	fmt.Println(string(bytes))
	return nil
}
