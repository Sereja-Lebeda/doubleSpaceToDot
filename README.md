# Double Space to Period (Windows)

Простая утилита, которая превращает **два быстрых нажатия пробела** в **точку + пробел** — точно так же, как это работает на macOS и iOS при наборе текста.

## Что делает программа

- Отслеживает нажатия пробела глобально (во всех приложениях).
- Если два пробела нажаты подряд с интервалом менее 500 мс → удаляет эти два пробела и вставляет `. `.

## Требования

- Windows (протестировано на Windows 10/11)
- Go 1.21+ (рекомендуется 1.22)
- MinGW-w64 / gcc (для cgo в robotgo и gohook)

## Установка зависимостей

```bash
go mod init double-space-to-dot
go get github.com/robotn/gohook
go get github.com/go-vgo/robotgo
go mod tidy
```

## Сборка проекта

Для работы файла в фоне следует применить:
'go build -ldflags "-H windowsgui" -o double-space.exe'
