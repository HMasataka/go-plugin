package shared

import (
	"errors"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	plugin "github.com/hashicorp/go-plugin"
)

type Info struct {
	ID     string
	Path   string
	Client *plugin.Client
}

func NewManager(glob, dir string, pluginImpl plugin.Plugin) *Manager {
	manager := &Manager{
		Glob:       glob,
		Path:       dir,
		Plugins:    map[string]*Info{},
		pluginImpl: pluginImpl,
	}

	return manager
}

type Manager struct {
	Glob        string
	Path        string
	Plugins     map[string]*Info
	initialized bool
	pluginImpl  plugin.Plugin
}

func (m *Manager) Init() error {
	plugins, err := plugin.Discover(m.Glob, m.Path)
	if err != nil {
		return err
	}

	for _, plugin := range plugins {
		_, file := filepath.Split(plugin)
		globAsterisk := strings.LastIndex(m.Glob, "*")
		trim := m.Glob[0:globAsterisk]
		id := strings.TrimPrefix(file, trim)

		m.Plugins[id] = &Info{
			ID:   id,
			Path: plugin,
		}
	}

	m.initialized = true

	return nil
}

func (m *Manager) Launch() error {
	for id, info := range m.Plugins {
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: HandshakeConfig,
			Plugins:         m.pluginMap(id),
			Cmd:             exec.Command(info.Path),
		})

		if _, ok := m.Plugins[id]; !ok {
			continue
		}

		m.Plugins[id].Client = client
	}

	return nil
}

func (m *Manager) Dispose() {
	var wg sync.WaitGroup

	for _, info := range m.Plugins {
		wg.Add(1)

		go func(client *plugin.Client) {
			client.Kill()
			wg.Done()
		}(info.Client)
	}

	wg.Wait()
}

func (m *Manager) Get(id string) (interface{}, error) {
	if _, ok := m.Plugins[id]; !ok {
		return nil, errors.New("plugin not found")
	}

	client := m.Plugins[id].Client

	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense(id)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

func (m *Manager) pluginMap(id string) map[string]plugin.Plugin {
	p := map[string]plugin.Plugin{}
	p[id] = m.pluginImpl

	return p
}
