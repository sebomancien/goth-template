# goth-template
This repository contains the template project for a GoTH stack (Go, Templ, HTMX) web application.

# Usage
The GoTH stack uses templ for generating front end source files. These files needs to be generated using the templ cli tool.

Install templ using the `go install` command below.
```
go install github.com/a-h/templ/cmd/templ@latest
```

Generate templ go files using `templ generate`.
```
templ generate
```

Download the project dependencies and run the web application.
```
go mod download
go run ./cmd/web/main.go
```

# License
The code in this repository is released under the MIT license. See [LICENSE](LICENSE).
