package main

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	gomysql "github.com/go-mysql-org/go-mysql/mysql"
	"github.com/go-mysql-org/go-mysql/replication"
)

type sink struct {
	msg string
}

func (s *sink) Write(p []byte) (n int, err error) {
	s.msg += string(p)
	return len(p), nil
}

func (s *sink) Clear() {
	s.msg = ""
}

func (s *sink) Get(key string) string {
	re := regexp.MustCompile(key + `: (.*)\n`)
	buf := re.FindAllStringSubmatch(s.msg, 1)
	if len(buf) > 0 {
		return buf[0][1]
	}
	return ""
}

func do() error {
	// GO_DEMO_DB_PW = 'my_password'
	host := os.Getenv("BINLOG_DB_HOST")
	port, err := strconv.Atoi(os.Getenv("BINLOG_DB_PORT"))
	if err != nil {
		return err
	}
	user := os.Getenv("BINLOG_DB_USER")
	pw := os.Getenv("BINLOG_DB_PW")
	file := os.Getenv("BINLOG_FILE")
	search := os.Getenv("BINLOG_SEARCH")
	s := replication.NewBinlogSyncer(replication.BinlogSyncerConfig{
		ServerID:   1,
		Flavor:     gomysql.MySQLFlavor,
		Host:       host,
		Port:       uint16(port),
		User:       user,
		Password:   pw,
		TLSConfig:  nil,
		UseDecimal: true,
	})
	t, err := s.StartSync(gomysql.Position{Name: file, Pos: 0})
	if err != nil {
		return err
	}
	ctx := context.TODO()
	var cached sink
	for {
		e, err := t.GetEvent(ctx)
		if err != nil {
			return err
		}
		var buf sink
		e.Dump(&buf)

		if e.Header.EventType == replication.TABLE_MAP_EVENT {
			cached.Clear()
			e.Dump(&cached)
		}
		if (strings.Contains(buf.msg, "WriteRowsEvent") ||
			strings.Contains(buf.msg, "UpdateRowsEvent")) &&
			strings.Contains(buf.msg, `"`+search+`"`) {

			if cached.msg != "" {
				fmt.Println("table: ", cached.Get("Table"))
			}
			fmt.Println(buf.msg)
		}
		//if strings.Contains(buf.msg, search) {
		//	fmt.Println(buf.msg)
		//}
	}
}

func main() {
	if err := do(); err != nil {
		fmt.Println(err)
	}
}
