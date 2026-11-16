package deploy

import (
	"errors"
	"fmt"
	"path/filepath"
	"strings"
)

var (
	ErrInvalidApplication = errors.New("application name is required")
	ErrInvalidSource      = errors.New("source directory is required")
	ErrInvalidCommand     = errors.New("start command is required")
)

// Spec describes the minimum input needed to run one local application release.
type Spec struct {
	Application string
	SourceDir   string
	StartCommand string
	Args        []string
	Port        int
}

func (s Spec) Validate() error {
	if strings.TrimSpace(s.Application) == "" {
		return ErrInvalidApplication
	}
	if strings.TrimSpace(s.SourceDir) == "" {
		return ErrInvalidSource
	}
	if filepath.IsAbs(s.SourceDir) == false {
		return errors.New("source directory must be absolute")
	}
	if strings.TrimSpace(s.StartCommand) == "" {
		return ErrInvalidCommand
	}
	if s.Port < 1 || s.Port > 65535 {
		return fmt.Errorf("port must be between 1 and 65535")
	}
	return nil
}

// Command returns the executable and arguments for a validated local release.
func (s Spec) Command() ([]string, error) {
	if err := s.Validate(); err != nil {
		return nil, err
	}
	command := make([]string, 0, len(s.Args)+1)
	command = append(command, s.StartCommand)
	command = append(command, s.Args...)
	return command, nil
}
