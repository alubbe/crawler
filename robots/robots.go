package robots

import (
	"io"
	"net/http"

	"github.com/bodokaiser/gerenuk/parser"
)

type Robot struct {
	parsers []parser.Parser
}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) Open(url string) error {
	res, err := http.Get(url)

	if err != nil {
		return err
	}

	for i := 0; i < len(r.parsers); i++ {
		io.Copy(r.parsers[i], res.Body)
	}

	return nil
}

func (r *Robot) RegisterParser(p parser.Parser) {
	r.parsers = append(r.parsers, p)
}