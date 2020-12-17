package service

import srv "github.com/kardianos/service"

func (p *Program) getConfig(path string) *srv.Config {
	return &srv.Config{
		Name:             p.Name,
		DisplayName:      p.DisplayName,
		Description:      p.Description,
		WorkingDirectory: path,
		Arguments:        []string{"service"},
	}
}
