# S2 Geometry Calculator

A user-friendly web interface for converting between GPS coordinates and [S2 cells](http://s2geometry.io/).

## Installation

Requirements: >= Go 1.11

Clone the repository

```bash
go get -u github.com/gojekfarm/s2-calc
```

Install `goexec`

```bash
go get -u github.com/shurcooL/goexec
```

## Usage

Install dependencies

```bash
make init-deps get-deps
```

Build and run server

```bash
make all
```

Navigate to `localhost:<PORT_NUMBER>` (default 8080)
