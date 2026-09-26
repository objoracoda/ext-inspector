# ext-inspector
CLI tool to inspect Chrome Web Store extensions by ID

ext-inspector/
├── cmd/
│   ├── cli/                 # CLI-версия (то, что делаем сейчас)
│   │   └── main.go
│   └── server/              # будущий HTTP-сервер (пока пусто, создашь позже)
│       └── .gitkeep
├── internal/
│   ├── app/                 # «сценарии» приложения: то, что объединяет парсер + вывод
│   │   └── inspect.go       # логика «взять ID → получить инфу → вернуть отчёт»
│   ├── chrome/
│   │   └── client.go        # работа с Chrome Web Store (HTTP-запросы, парсинг)
│   ├── model/               # структуры данных (Extension, Report и т.п.)
│   │   └── extension.go
│   └── cli/                 # парсинг аргументов, форматирование вывода для CLI
│       └── run.go
├── web/                     # будущий сайт (пока пусто)
│   ├── static/
│   │   └── .gitkeep
│   └── templates/
│       └── .gitkeep
├── .gitignore
├── go.mod
├── go.sum                   
├── LICENSE
└── README.md