# robotstxt

`robotstxt` is a tool that will build URLs based on the robots.txt file of the host provided.

# Usage example

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


# Install

```
▶ go get -u github.com/NkxxkN/robotstxt
```

# Credits

This tool was 100% inspired from [TomNomNom](https://github.com/tomnomnom) tools.