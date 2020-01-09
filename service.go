package service

import (
	"fmt"

	"bitbucket.org/butenta/pkg-log"
	srv "github.com/kardianos/service"
)

var logger srv.Logger

// Program structure for service
type Program struct {
	Name        string
	DisplayName string
	Description string
	CallBack    func()
}

// Start using for service start
func (p *Program) Start(s srv.Service) error {
	go p.run()
	return nil
}

// Stop function for stop service
func (p *Program) Stop(s srv.Service) error {
	log.Info("Service Stopping!")
	return nil
}

func (p *Program) run() error {
	log.Info("Service running on %v platform.", srv.Platform())

	p.CallBack()

	return nil
}

// New creating and return Program structure
func New(name, display, desc string) *Program {
	return &Program{
		Name:        name,
		DisplayName: display,
		Description: desc,
	}
}

// SetCb for setting CallBack
func (p *Program) SetCb(cb func()) {
	p.CallBack = cb
}

// Controller function using for controll service
func (p *Program) Controller(param string) error {
	svcConfig := &srv.Config{
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
		Arguments:   []string{"service"},
	}

	s, err := srv.New(p, svcConfig)
	if err != nil {
		return err
	}

	logger, err = s.Logger(nil)
	if err != nil {
		return err
	}

	var executionError error
	switch param {
	case "install":
		executionError = s.Install()
		if executionError != nil {
			return fmt.Errorf("Failed to install: %s", executionError)
		}
		fmt.Printf("Service \"%s\" installed.\n", svcConfig.DisplayName)
	case "uninstall":
		executionError = s.Uninstall()
		if executionError != nil {
			return fmt.Errorf("Failed to remove: %s", executionError)
		}
		fmt.Printf("Service \"%s\" removed.\n", svcConfig.DisplayName)
	case "stop":
		executionError = s.Stop()
		if err != nil {
			return fmt.Errorf("Failed to stop: %s", executionError)
		}
		fmt.Printf("Service \"%s\" stopped.\n", svcConfig.DisplayName)
	case "start":
		executionError = s.Start()
		if executionError != nil {
			return fmt.Errorf("Failed to start \"%s\" : %s", svcConfig.DisplayName, executionError)
		}
		fmt.Printf("Service \"%s\" started.\n", svcConfig.DisplayName)
	default:
		executionError = s.Run()
		if executionError != nil {
			return fmt.Errorf("Failed to run \"%s\" : %s", svcConfig.DisplayName, executionError)
		}
		fmt.Printf("Service \"%s\" running.\n", svcConfig.DisplayName)
	}

	return nil
}
