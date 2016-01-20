package config

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
	"sync"
)

type Config struct {
	single_data map[string][]string
	mutil_data  map[string]map[string][]string
	changed     bool
	l           *sync.RWMutex
}

func NewConfig() *Config {
	r := new(Config)
	r.single_data = make(map[string][]string)
	r.mutil_data = make(map[string]map[string][]string)
	r.changed = false
	r.l = new(sync.RWMutex)

	return r
}

func (this *Config) IsChange() bool {
	return this.changed
}

func (this *Config) SetChange(b bool) {
	this.changed = b
}

func (this *Config) ConfigRLock() {
	this.l.RLock()
}

func (this *Config) ConfigRUnlock() {
	this.l.RUnlock()
}

func (this *Config) ConfigLock() {
	this.l.Lock()
}

func (this *Config) ConfigUnlock() {
	this.l.Unlock()
}

func (this *Config) UpdateConfig(r io.Reader) error {
	buf := bufio.NewReader(r)
	var curSection string
	var nullSection string
	var section, key, value string
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}

			if len(line) == 0 {
				break
			}
		}
		strline := (string)(line)
		tlen := len(strline)
		strline = strings.TrimSpace(strline)

		switch {
		case tlen == 0:
			continue
		case strline[0] == '[' && strline[tlen-1] == ']':
			curSection = strline[1 : tlen-1]
			continue
		case strline[0] == '#' || strline[0] == ';':
			//skip
			continue
		default:

			pos := strings.Index(strline, "=")
			if pos <= 0 {
				continue
			}

			if len(curSection) != 0 {
				section = curSection
			} else {
				section = nullSection
			}

			key = strline[:pos]
			value = strline[pos+1:]
			key = strings.TrimSpace(key)
			value = strings.TrimSpace(value)
			this.setValue(section, key, value)
		}
	}

	this.SetChange(true)
	return nil
}

func (this *Config) setValue(s, k, v string) {
	if len(s) == 0 {
		this.single_data[k] = append(this.single_data[k], v)
	} else {
		_, ok := this.mutil_data[s]
		if !ok {
			this.mutil_data[s] = make(map[string][]string)
		}
		this.mutil_data[s][k] = append(this.mutil_data[s][k], v)
	}
}

func (this *Config) getValue(s, k string) ([]string, error) {
	if len(s) == 0 {
		v, ok := this.single_data[k]
		if !ok {
			return nil, errors.New("key is not found!")
		}
		return v, nil
	} else {
		_, ok := this.mutil_data[s]
		if !ok {
			return nil, errors.New("section is not found!")
		}

		v, ok := this.mutil_data[s][k]
		if !ok {
			return nil, errors.New("key is not found!")
		}

		return v, nil
	}
}

// Bool returns bool type value.
func (c *Config) Bool(section, key string) (bool, error) {
	value, err := c.getValue(section, key)
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(value[0])
}

// Float64 returns float64 type value.
func (c *Config) Float64(section, key string) (float64, error) {
	value, err := c.getValue(section, key)
	if err != nil {
		return 0.0, err
	}
	return strconv.ParseFloat(value[0], 64)
}

// Int returns int type value.
func (this *Config) Int(section, key string) (int, error) {
	value, err := this.getValue(section, key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(value[0])
}

// Int64 returns int64 type value.
func (this *Config) Int64(section, key string) (int64, error) {
	value, err := this.getValue(section, key)
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(value[0], 10, 64)
}

func (this *Config) MustInt(section, key string, defaultval int) int {
	value, err := this.Int(section, key)
	if err != nil {
		return defaultval
	}

	return value
}

func (this *Config) MustInt64(section, key string, defaultval int64) int64 {
	value, err := this.Int64(section, key)
	if err != nil {
		return defaultval
	}

	return value
}

func (this *Config) MustBool(section, key string, defaultval bool) bool {
	value, err := this.Bool(section, key)
	if err != nil {
		return defaultval
	}

	return value
}

func (this *Config) MustFloat64(section, key string, defaultval float64) float64 {
	value, err := this.Float64(section, key)
	if err != nil {
		return defaultval
	}

	return value
}

func (this *Config) MustString(section, key string, defaultval string) string {
	value, err := this.getValue(section, key)
	if len(value) == 0 || err != nil {
		return defaultval
	}

	return value[0]
}

func (this *Config) MustStringSlice(section, key string, defaultval []string) []string {
	value, err := this.getValue(section, key)
	if err != nil {
		return defaultval
	}

	return value
}
