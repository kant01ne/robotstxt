# robotstxt

`robotstxt` is a tool that will build URLs based on the robots.txt file of the host provided.


# Install

```
▶ go get -u github.com/NkxxkN/robotstxt
```

# Basic Usage

```
▶ echo "https://yahoo.com" | robotstxt
https://yahoo.com/p/
https://yahoo.com/r/
https://yahoo.com/bin/
https://yahoo.com/includes/
https://yahoo.com/blank.html
https://yahoo.com/_td_api
https://yahoo.com/_tdpp_api
https://yahoo.com/_remote
https://yahoo.com/_multiremote
https://yahoo.com/_tdhl_api
https://yahoo.com/digest
https://yahoo.com/fpjs
https://yahoo.com/myjs
```

## Concurrency

You can set the concurrency level with the -c flag (default = 20):

```
▶ cat domains.txt | robotstxt -c 50
```

## Timeout

You can change the timeout by using the -t flag and specifying a timeout in milliseconds (default = 10000):

```
▶ cat domains.txt | robotstxt -t 20000
```

# Credits

This tool was 100% inspired from [TomNomNom](https://github.com/tomnomnom) tools <3