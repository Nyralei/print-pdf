# PDF Creator

A simple application to create A4 PDFs from images
Простое приложение для создания PDF-файлов в формате A4 из изображений

![alt text](image.png)

<details>
<summary>English</summary>

## Configuration

The configuration file `config.json` should be located in the `~/.config/print-pdf` directory. Edit this file to set default values for the application.

## Localization

Localization files are in the `internal/ui/locales` directory. You can add more languages by creating new JSON files and updating the `config.json` file to use the new language.

## Running the Application

1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Build the application:
   ```sh
   go build -o print-pdf -v ./cmd/print-pdf
   ```
3. Run the application:
   ```sh
   ./print-pdf
   ```
</details>
<details>
<summary>Русский</summary>

## Конфигурация

Файл конфигурации `config.json` должен находиться в каталоге `~/.config/print-pdf`. Отредактируйте этот файл, чтобы установить значения по умолчанию для приложения.

## Локализация

Файлы локализации находятся в каталоге `internal/ui/locales`. Вы можете добавить больше языков, создав новые JSON-файлы и обновив файл `config.json` для использования нового языка.

## Запуск приложения

1. Установите зависимости:
   ```sh
   go mod tidy
   ```
2. Соберите приложение:
   ```sh
   go build -o print-pdf -v ./cmd/print-pdf
   ```
3. Запустите приложение:
   ```sh
   ./print-pdf
   ```
</details>