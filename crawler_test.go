package crawler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func TestCrawler(t *testing.T) {
	check.Suite(&CrawlerSuite{})
	check.TestingT(t)
}

type CrawlerSuite struct {
	counter int
	request *Request
	crawler *Crawler
	server  *httptest.Server
}

func (s *CrawlerSuite) SetUpSuite(c *check.C) {
	s.server = httptest.NewServer(http.HandlerFunc(s.handler))
}

func (s *CrawlerSuite) SetUpTest(c *check.C) {
	r, err := NewRequest(s.server.URL)
	c.Assert(err, check.IsNil)

	s.crawler = New()
	s.request = r
	s.counter = 0
}

func (s *CrawlerSuite) TestSingle(c *check.C) {
	s.crawler.Do(s.request)
	s.crawler.Do(s.request)
	s.crawler.Run(1)

	_, ok := <-s.request.Done
	c.Assert(ok, check.Equals, false)

	c.Check(s.counter, check.Equals, 1)
}

func (s *CrawlerSuite) handler(w http.ResponseWriter, r *http.Request) {
	s.counter++

	io.WriteString(w, `
		<!DOCTYPE html>
		<html>
			<head></head>
			<body>
				<h1>Example</h1>
				<a href="http://example.org"></a>
				<a href="https://example.com"></a>
			</body>
		</html>
	`)
}
