package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	userAgent = "Mozilla/5.0 (compatible; robotxtxt; +https://github.com/NkxxkN/robotstxt)"
)

var httpClient = &http.Client{}

func main() {

	// concurrency flag
	var (
		httpClient  http.Client
		concurrency int
		to          int
		wildcard    bool
	)

	flag.IntVar(&concurrency, "c", 20, "Set the concurrency level")
	flag.BoolVar(&wildcard, "w", true, "Should include wildcards")
	flag.IntVar(&to, "t", 10000, "timeout (milliseconds)")

	flag.Parse()

	// make an actual time.Duration out of the timeout
	timeout := time.Duration(to * 1000000)

	httpClient = http.Client{
		Timeout: timeout,
	}

	sc := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup

	jobs := make(chan string)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for host := range jobs {
				req, err := http.NewRequest("GET", fmt.Sprintf(
					"%s/robots.txt", host,
				), nil)

				if err != nil {
					return
				}

				req.Header.Add("User-Agent", userAgent)

				resp, err := httpClient.Do(req)
				if err != nil {
					continue
				}
				if resp.Body == nil {
					continue
				}
				defer resp.Body.Close()

				// always read the full body so we can re-use the tcp connection
				b, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					continue
				}

				body := string(b)

				lines := bufio.NewScanner(strings.NewReader(body))
				for lines.Scan() {
					line := lines.Text()
					if !wildcard && strings.Contains(line, "*") {
						continue
					}

					p := strings.Split(line, "llow: ") // Disallow: && Allow:
					if len(p) == 2 {
						path := p[1]
						if strings.HasPrefix(path, "https://") || strings.HasPrefix(path, "http://") {
							fmt.Printf("%s\n", path)
						} else {
							if strings.HasPrefix(path, "/") {
								fmt.Printf("%s%s\n", host, path)
							} else {
								fmt.Printf("%s/%s\n", host, path)
							}
						}
					}
				}
			}

		}()
	}

	for sc.Scan() {
		host := sc.Text()
		jobs <- host
	}

	close(jobs)
	wg.Wait()

}
