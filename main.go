package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

var httpClient = &http.Client{}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var wg sync.WaitGroup

	jobs := make(chan string)

	for i := 0; i < 20; i++ {
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

				req.Header.Add("User-Agent", "User-Agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36")

				resp, err := httpClient.Do(req)
				if err != nil {
					return
				}
				if resp.Body == nil {
					return
				}
				defer resp.Body.Close()

				// always read the full body so we can re-use the tcp connection
				b, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					return
				}

				body := string(b)

				lines := bufio.NewScanner(strings.NewReader(body))
				for lines.Scan() {
					line := lines.Text()
					p := strings.Split(line, "llow: ") // Disallow: && Allow:
					if len(p) == 2 {
						path := p[1]
						fmt.Printf("%s%s\n", host, path)
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
