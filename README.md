# gomarkets
Библиотека торговли из Go

### Connector
    Универсальный интерфейс к биржам
    Для каждой биржи свое расширение connector
    Используется REST API и WebSocket API
### OrderBook
    Запускается с указанием connector
    Подписывается на обновление книги
### ChartData