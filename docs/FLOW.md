FLOW:
gostarter /path/to/example --flags

flags:
- name: pick what template to make 
- with_optinal: make project structure with optional configs

Reference on best practice of cli arch
- http://www.catb.org/esr/writings/taoup/html/
- https://aosabook.org/en/index.html
- https://github.com/lirantal/nodejs-cli-apps-best-practices

gostarter/
│
├── cmd/
│   └── gostarter/
│       └── main.go
│
├── internal/
│   │
│   ├── app/
│   │   └── container.go
│   │
│   ├── commands/
│   │   ├── root.go
│   │   ├── init.go
│   │   ├── new.go
│   │   └── git.go
│   │
│   ├── database/
│   │   ├── sqlite.go
│   │   └── migrations.go
│   │
│   ├── github/
│   │   ├── model.go
│   │   ├── repository.go
│   │   ├── service.go
│   │   └── ssh.go
│   │
│   ├── template/
│   │   ├── model.go
│   │   ├── loader.go
│   │   ├── generator.go
│   │   ├── repository.go
│   │   └── service.go
│   │
│   └── filesystem/
│       └── filesystem.go
│
├── templates/
│   └── go/
│       └── basic.yaml
│
├── embed.go
├── go.mod
└── go.sum