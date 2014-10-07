package goom

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type container struct {
	Lists   []map[string][]map[string]string `json:"lists"`
	allCmds map[string]string
}

func New() (*container, error) {
	user, err := user.Current()
	if err != nil {
		return nil, err
	}

	conf, err := ioutil.ReadFile(filepath.Join(user.HomeDir, ".boom"))
	if err != nil {
		return nil, err
	}
	var c container
	if err := json.Unmarshal(conf, &c); err != nil {
		return nil, err
	}

	c.allCmds = make(map[string]string)
	for _, l := range c.Lists {
		for _, cmds := range l {
			for _, cmd := range cmds {
				for key, val := range cmd {
					c.allCmds[key] = val
				}
			}
		}
	}
	return &c, nil
}

func (c *container) Get(key string) (string, error) {
	if val, ok := c.allCmds[key]; ok {
		err := clipboard.WriteAll(val)
		return val, err
	}
	return "", fmt.Errorf("Unable to find command for key %s", key)
}
