package config

import (
	"bytes"
	"os"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func LoadConfigByZK(zkserver []string, path string, watch bool) (*Config, error) {
	con, _, err := zk.Connect(zkserver, time.Second)
	if err != nil {
		return nil, err
	}

	c := NewConfig()

	data, _, chWatch, err := con.GetW(path)
	if err != nil {
		return nil, err
	}

	r := bytes.NewReader(data)

	err = c.UpdateConfig(r)
	if err != nil {
		return nil, err
	}

	if watch {
		go func() {
			for {
				e := <-chWatch

				//again register
				data, _, chWatch, err = con.GetW(path)
				if e.Type == zk.EventNodeDataChanged {
					r := bytes.NewReader(data)

					err = c.UpdateConfig(r)
					if err != nil {
						panic(err)
					}
				}
			}
		}()
	}
	return c, nil
}

func LoadConfigByFile(file string) (*Config, error) {
	fl, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	c := NewConfig()

	err = c.UpdateConfig(fl)
	fl.Close()
	if err != nil {
		return nil, err
	}

	return c, nil
}
