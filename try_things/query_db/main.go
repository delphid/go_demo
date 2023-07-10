package main

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xyproto/randomstring"
)

type Hub struct {
	sync.Mutex
	dsn    string
	pools  map[*sql.DB]string
	ctx    context.Context
	cancel func()
}

func NewHub() *Hub {
	h := Hub{}
	h.ctx, h.cancel = context.WithCancel(context.Background())
	return &h
}

func (h *Hub) Read(stmt string) (*sql.Rows, error) {
	pool, err := h.newPool(h.dsn)
	h.unregister(pool)
	if err != nil {
		return nil, err
	}
	conn, err := pool.Conn(h.ctx)
	defer conn.Close()
	var sessionID string
	if err := conn.QueryRowContext(h.ctx, `SELECT CONNECTION_ID()`).Scan(&sessionID); err != nil {
		return nil, err
	}
	h.pools[pool] = sessionID
	return conn.QueryContext(h.ctx, stmt)
}

func (h *Hub) Write(stmt string) (sql.Result, error) {

}

func (h *Hub) kill(pool *sql.DB, sessionID string) {
	pool.Exec("KILL QUERY ?", sessionID)
}

func (h *Hub) newPool(dsn string) (*sql.DB, error) {
	h.Lock()
	defer h.Unlock()
	pool, err := sql.Open("mysql", dsn)
	h.pools[pool] = pool
	return pool, err
}

func (h *Hub) unregister(pool *sql.DB) error {
	h.Lock()
	defer h.Unlock()
	delete(h.pools, pool)
	return nil
}

func (h *Hub) teardown() {
	fmt.Println("teardown db connections")
	h.cancel()
	for p, s := range h.pools {
		h.kill(p, s)
	}
}

func query(h *hub, dsn, stmt string) error {
	//fmt.Println(stmt)
	db, err := h.new(dsn)
	if err != nil {
		return err
	}
	defer h.close(db)
	_, err = db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}

func name() string {
	return randomstring.HumanFriendlyEnglishString(10)
}

func id() string {
	return strconv.Itoa(rand.Intn(500000))
}

func longIn(size int) string {
	s := `UPDATE large SET a="` + name() + `" where id in (`
	for i := 0; i < size; i++ {
		s += id() + ","
	}
	s = s[:len(s)-1] + ")"
	return s
}

func doQuery(h *hub) error {
	dsn := os.Getenv("QUERY_DB_DSN")
	//stmt := longIn(10)
	stmt := `UPDATE large SET a="` + name() + `" where a="` + name() + `"`
	return query(h, dsn, stmt)
}

func spawn(max int, f func() error) error {
	workers := make(chan int, max)
	done := make(chan int)
	go func() {
		for {
			workers <- 1
			go func() {
				err := f()
				<-workers
				done <- 1
				if err != nil {
					fmt.Println(err)
				}
			}()
		}
	}()
	go func() {
		for {
			fmt.Println("now workers: ", len(workers))
			time.Sleep(5 * time.Second)
		}
	}()
	go func() {
		n := 0
		ticker := time.NewTicker(5 * time.Second)
		report := func() {
			fmt.Println("workers done: ", n)
		}
		for {
			select {
			case one := <-done:
				n += one
				if n%5 == 0 {
					report()
				}
			case <-ticker.C:
				report()
			}
		}
	}()
	return nil
}

func main() {
	h := newHub()
	if err := spawn(15, func() error {
		return doQuery(h)
	}); err != nil {
		fmt.Println(err)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-sigCh
	h.teardown()
	fmt.Println("done")
}
