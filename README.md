# Daigou

![](header.jpg)

### structure

```
.
├── Makefile
├── Procfile
├── README.md
├── config
│   ├── config.go
│   ├── development.yaml
│   ├── production.yaml
│   └── test.yaml
├── controllers
│   └── user.go
├── db
│   └── db.go
├── forms
│   └── user.go
├── header.jpg
├── main.go
├── middlewares
│   └── auth.go
├── models
│   └── user.go
└── server
    ├── router.go
    └── server.go
```

## Befor Installation

__Install vgo__

`go get -u golang.org/x/vgo`

then run:

```sh
vgo install
```

__Install Reflex__

`go get github.com/cespare/reflex`

then run:

```sh
reflex -s -r '\.go$' vgo run main.go
```

## License

