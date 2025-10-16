package config

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const DatabaseMigrationTable = "migrations"

type Database struct {
	Host             string
	Port             int
	Username         string
	Password         string `json:"-"`
	Database         string
	AdditionalParams map[string]string `json:",omitempty"`
	MaxOpenConns     int
	MaxIdleConns     int
	ConnMaxLifetime  time.Duration
}

func (c *Database) ConnectionString() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("postgres://%s:%s@%s:%d/%s", c.Username, c.Password, c.Host, c.Port, c.Database))

	if len(c.AdditionalParams) > 0 {
		b.WriteString("?")
		params := make([]string, 0, len(c.AdditionalParams))
		for param := range c.AdditionalParams {
			params = append(params, param)
		}

		sort.Strings(params)

		for _, param := range params {
			fmt.Fprintf(&b, "&%s=%s", param, c.AdditionalParams[param])
		}
	}

	return b.String()
}
