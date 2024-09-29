## Запускать проект с помощью 
```
docker-compose up --build
```
### Работает - ping 
```
curl http://localhost:8080/api/ping
```

### Работает - tender
```
curl -X POST "http://localhost:8080/api/tenders/new" \
-H "Content-Type: application/json" \
-d '{
      "name": "Test Tender",
      "description": "This is a test tender.",
      "serviceType": "Construction",
      "status": "Created",
      "organizationId": "org123",
      "creatorUsername": "user123"
    }'
```
```
curl -X GET "http://localhost:8080/api/tenders"
```

### Пока не работает - tender
```
curl -X GET "http://localhost:8080/api/tenders/{tenderId}"
```
```
curl -X PATCH "http://localhost:8080/api/tenders/{tenderId}/status" \
-H "Content-Type: application/json" \
-d '{
  "status": "Published"
}'
```

### Работает - bids
```
curl -X POST "http://localhost:8080/api/bids/new" \
-H "Content-Type: application/json" \
-d '{
  "name": "New Bid",
  "description": "This is a new bid.",
  "tenderId": "50137766-7f07-45ba-96e0-3ec252554a3e",
  "amount": 1500.00,
  "creatorUsername": "user123",
  "status": "Submitted"
}'
```

## internal/controllers:

### bid.go:
В приведённом коде находится контроллер для управления предложениями (bids) в тендерной системе с использованием фреймворка Echo. Давайте разберём каждую функцию:

### 1. **`GetBids` (Получение списка предложений для тендера)**
   - **Описание**: Эта функция отвечает за получение списка предложений, связанных с конкретным тендером.
   - **Параметры**:
     - `tenderId`: Получен из query-параметра запроса. Используется для фильтрации предложений по тендеру.
     - `limit`, `offset`: Дополнительные параметры для пагинации.
   - **Логика**:
     - Вызывает сервисный слой (`services.GetBids`), который возвращает список предложений для указанного тендера.
     - Возвращает данные клиенту в формате JSON с кодом статуса 200 (`http.StatusOK`).
     - В случае ошибки возвращает ошибку с кодом 500.

### 2. **`CreateBid` (Создание нового предложения)**
   - **Описание**: Создает новое предложение в системе.
   - **Параметры**:
     - Функция принимает данные предложения в теле запроса. Эти данные связываются с моделью `Bid` через метод `c.Bind(&bid)`.
   - **Логика**:
     - Если связывание данных (binding) успешно, вызывается функция `services.CreateBid`, которая сохраняет новое предложение в базе данных.
     - Если предложение успешно создано, оно возвращается клиенту с кодом 200. В случае ошибки возвращается код 500 и сообщение об ошибке.

### 3. **`GetBid` (Получение предложения по его ID)**
   - **Описание**: Получает конкретное предложение по его идентификатору.
   - **Параметры**:
     - `bidId`: Идентификатор предложения передаётся как часть пути URL через `c.Param("bidId")`.
   - **Логика**:
     - Функция вызывает `services.GetBid`, чтобы получить данные предложения по его ID.
     - Если данные успешно получены, они возвращаются клиенту с кодом 200.
     - Если предложение не найдено или возникает ошибка, возвращается ошибка с кодом 500.

### 4. **`UpdateBid` (Обновление предложения)**
   - **Описание**: Обновляет данные существующего предложения.
   - **Параметры**:
     - `bidId`: Идентификатор предложения, который обновляется.
     - Тело запроса содержит данные, которые нужно изменить (как `map[string]interface{}` для гибкости обновления полей).
   - **Логика**:
     - Сначала происходит привязка данных обновления (`c.Bind(&updates)`).
     - Затем вызывается сервисная функция `services.UpdateBid`, которая обновляет соответствующие поля предложения в базе данных.
     - Если обновление успешно, возвращаются обновлённые данные с кодом 200. В случае ошибки — код 500 и сообщение об ошибке.

### 5. **`DeleteBid` (Удаление предложения)**
   - **Описание**: Удаляет предложение по его идентификатору.
   - **Параметры**:
     - `bidId`: Идентификатор предложения передаётся как параметр пути URL.
   - **Логика**:
     - Функция вызывает `services.DeleteBid` для удаления предложения из базы данных.
     - В случае успешного удаления возвращается сообщение с кодом 200. Если возникает ошибка — код 500.

### Общий поток работы:
- Эти функции работают как слой контроллеров для обработки HTTP-запросов.
- Они принимают входящие запросы от клиента, извлекают параметры (например, `bidId`, `tenderId`), вызывают соответствующие функции из сервисного слоя для выполнения бизнес-логики (например, создание или обновление предложения), и затем возвращают ответ клиенту в виде JSON.

### ping.go:
Функция `CheckServer` из приведённого кода выполняет простую задачу — проверку доступности сервера. Она используется для обработки запросов на эндпоинт `/ping`. Давайте подробнее разберём её работу:

### Описание функции `CheckServer`

- **Цель**: Проверить, работает ли сервер. Обычно такие проверки используются для "ping" запросов, которые могут быть отправлены для проверки состояния или здоровья сервера.
  
- **Параметры**:
  - `c echo.Context`: Это объект контекста запроса, предоставляемый фреймворком Echo. Он содержит всю информацию о текущем запросе, такой как параметры, заголовки, тело запроса и т.д. Через этот объект также отправляется ответ клиенту.

- **Логика**:
  - Когда приходит запрос на эндпоинт `/ping`, функция возвращает строку `"ok"` в ответ на запрос.
  - Ответ имеет HTTP статус-код `200` (`http.StatusOK`), что означает успешное выполнение запроса.

- **Пример использования**:
  - Если сервер работает и запрос на `/ping` обрабатывается успешно, клиент получает в ответ сообщение `"ok"` с кодом `200`. Это сигнализирует, что сервер запущен и функционирует корректно.

### Поток работы:
1. При запросе к серверу на адрес `/ping`, вызывается эта функция.
2. Сервер отправляет клиенту текст `"ok"` с кодом успешного выполнения.
  
### Пример curl-запроса:

```bash
curl -X GET "http://localhost:8080/ping"
```

Ответ:

```
ok
```

Это помогает убедиться, что сервер работает и отвечает на запросы.

### tender.go:
Функции в контроллере управляют запросами и передают их в соответствующие сервисы для обработки. Они работают с тендерами и реализуют основные действия (CRUD — создание, чтение, обновление, удаление). Давайте разберём каждую функцию:

### 1. **GetTenders**
   - **Назначение**: Получение списка тендеров с возможностью фильтрации по `service_type`, а также поддержка пагинации через параметры `limit` и `offset`.
   - **Параметры**:
     - `service_type` — фильтр по типу услуги.
     - `limit` — количество записей для пагинации.
     - `offset` — смещение для пагинации.
   - **Логика**: Вызывает функцию из сервиса `GetTenders`, передаёт параметры фильтрации и пагинации. Возвращает JSON-ответ с тендерами или ошибкой.

### 2. **CreateTender**
   - **Назначение**: Создание нового тендера.
   - **Параметры**:
     - Тело запроса, содержащее данные тендера (например, имя, описание).
   - **Логика**: Принимает и связывает данные из тела запроса в структуру `Tender`, вызывает сервис для создания тендера. Возвращает созданный тендер или ошибку.

### 3. **GetUserTenders**
   - **Назначение**: Получение тендеров, созданных пользователем.
   - **Параметры**:
     - `username` — имя пользователя, для которого нужно получить тендеры.
     - `limit`, `offset` — параметры пагинации.
   - **Логика**: Передаёт параметры в сервис `GetUserTenders` и возвращает список тендеров или ошибку.

### 4. **GetTenderStatus**
   - **Назначение**: Получение текущего статуса тендера по его ID.
   - **Параметры**:
     - `tenderId` — идентификатор тендера.
     - `username` — имя пользователя.
   - **Логика**: Получает статус тендера через сервис `GetTenderStatus` и возвращает его клиенту.

### 5. **UpdateTenderStatus**
   - **Назначение**: Обновление статуса тендера.
   - **Параметры**:
     - `tenderId` — идентификатор тендера.
     - `status` — новый статус тендера.
     - `username` — имя пользователя.
   - **Логика**: Обновляет статус тендера через сервис `UpdateTenderStatus` и возвращает обновлённые данные тендера или ошибку.

### 6. **EditTender**
   - **Назначение**: Редактирование полей тендера.
   - **Параметры**:
     - `tenderId` — идентификатор тендера.
     - `username` — имя пользователя.
     - Тело запроса с данными для обновления.
   - **Логика**: Принимает изменения для тендера, передаёт их в сервис `EditTender` и возвращает обновлённый тендер или ошибку.

### 7. **RollbackTender**
   - **Назначение**: Откат к предыдущей версии тендера.
   - **Параметры**:
     - `tenderId` — идентификатор тендера.
     - `version` — версия тендера, к которой нужно откатить.
     - `username` — имя пользователя.
   - **Логика**: Осуществляет откат тендера на указанную версию через сервис `RollbackTender` и возвращает результат или ошибку. Проверяет правильность формата версии перед вызовом сервиса.

### Общая структура
- Эти функции принимают запросы от клиента, обрабатывают параметры запроса, вызывают соответствующие сервисы для бизнес-логики и возвращают результаты в виде JSON-ответов.

## db:

### db.go:
Этот код настраивает подключение к базе данных PostgreSQL с использованием Gorm, выполняет миграции моделей и управляет пулом соединений. Давайте разберём работу каждой части функции.

### Переменные:
- **`DB *gorm.DB`**: Глобальная переменная, которая будет хранить подключение к базе данных, чтобы к нему могли обращаться другие части приложения.
- **`ErrRecordNotFound`**: Переменная для хранения ошибки "запись не найдена", которая часто используется при работе с базой данных.

### Функция `InitDB`:
Эта функция инициализирует подключение к базе данных PostgreSQL и настраивает параметры соединения.

#### Этапы работы:

1. **Получение строки подключения**:
   ```go
   dsn := os.Getenv("POSTGRES_CONN")
   ```
   Строка подключения берётся из переменной окружения `POSTGRES_CONN`. Если переменная не установлена, функция возвращает ошибку.

2. **Проверка строки подключения**:
   ```go
   if dsn == "" {
       return nil, fmt.Errorf("POSTGRES_CONN environment variable is not set")
   }
   ```
   Если строка подключения пустая (не задана), возвращается ошибка с соответствующим сообщением.

3. **Логирование строки подключения**:
   ```go
   fmt.Println("Connection string:", dsn)
   ```
   Это выводит строку подключения в консоль для отладки.

4. **Открытие подключения к базе данных**:
   ```go
   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
       Logger: logger.Default.LogMode(logger.Info),
   })
   ```
   Подключение к базе данных открывается с помощью драйвера PostgreSQL. Также включено логирование запросов к базе данных на уровне `Info`, что позволяет отслеживать выполняемые SQL-запросы.

5. **Автоматическая миграция**:
   ```go
   if err := db.AutoMigrate(&models.Tender{}, &models.Bid{}, &models.BidReview{}); err != nil {
       return nil, err
   }
   ```
   После подключения к базе данных Gorm автоматически создаёт или обновляет таблицы для моделей `Tender`, `Bid` и `BidReview`, если они не существуют или требуют изменений.

6. **Настройка параметров соединения**:
   ```go
   sqlDB, err := db.DB()
   ```
   Из объекта `gorm.DB` извлекается объект `sql.DB`, который управляет соединениями с базой данных.

   Далее настраиваются параметры соединений:
   - **`SetConnMaxIdleTime(30 * time.Second)`**: Максимальное время простоя соединения перед его закрытием.
   - **`SetConnMaxLifetime(5 * time.Minute)`**: Максимальное время жизни соединения перед его повторным использованием.
   - **`SetMaxOpenConns(10)`**: Максимальное количество открытых соединений к базе данных.
   - **`SetMaxIdleConns(5)`**: Максимальное количество "бездействующих" соединений в пуле.

7. **Сохранение подключения в глобальную переменную**:
   ```go
   DB = db
   ```
   Успешно созданное подключение сохраняется в глобальную переменную `DB`, чтобы его можно было использовать в других частях приложения.

8. **Возврат подключения**:
   Функция возвращает подключение к базе данных и возможную ошибку, если что-то пошло не так.

### Итог:
- **`InitDB`** настраивает подключение к PostgreSQL, управляет миграциями моделей и настраивает пул соединений для оптимизации использования ресурсов.
- Логирование запросов помогает отслеживать SQL-запросы во время работы приложения.
- Настройки пула соединений помогают управлять количеством одновременных соединений с базой данных, предотвращая их переполнение или истечение времени.

## models:

### bid.go
В этом коде определяются две модели данных для системы управления тендерами: **`Bid`** (предложение) и **`BidReview`** (отзыв на предложение). Эти структуры используются для представления и работы с данными в базе данных через GORM (ORM для Go). Давайте рассмотрим каждую часть этого кода.

### Перечисления (Enums):

1. **`BidStatus`**:
   ```go
   type BidStatus string
   ```
   Это тип, представляющий статус предложения, который может принимать одно из следующих значений:
   - **`BidCreated`**: Предложение создано.
   - **`BidPublished`**: Предложение опубликовано.
   - **`BidCanceled`**: Предложение отменено.
   - **`BidApproved`**: Предложение утверждено.
   - **`BidRejected`**: Предложение отклонено.

   Этот тип используется для обозначения текущего состояния предложения в системе.

2. **`BidAuthorType`**:
   ```go
   type BidAuthorType string
   ```
   Это тип, определяющий, кто является автором предложения:
   - **`Organization`**: Автором является организация.
   - **`User`**: Автором является пользователь.

   Этот тип используется для указания типа автора предложения.

### Модель **`Bid`** (Предложение):

```go
type Bid struct {
    ID             string       `json:"id" gorm:"primary_key"`
    Name           string       `json:"name"`
    Description    string       `json:"description"`
    TenderID       string       `json:"tenderId"`
    Amount         float64      `json:"amount"`
    CreatorUsername string      `json:"creatorUsername"`
    AuthorType     BidAuthorType `json:"authorType"`
    AuthorID       string       `json:"authorId"`
    Status         BidStatus    `json:"status"`
    CreatedAt      time.Time    `json:"createdAt"`
    Version        int          `json:"version"`
}
```

Эта структура описывает модель предложения (Bid), которая хранится в базе данных. Рассмотрим поля:
- **`ID`**: Уникальный идентификатор предложения. Является первичным ключом в базе данных.
- **`Name`**: Название предложения.
- **`Description`**: Описание предложения.
- **`TenderID`**: Идентификатор тендера, к которому относится предложение.
- **`Amount`**: Сумма предложения, предложенная участником тендера.
- **`CreatorUsername`**: Имя пользователя, создавшего предложение.
- **`AuthorType`**: Тип автора предложения (организация или пользователь).
- **`AuthorID`**: Идентификатор автора предложения (например, пользователя или организации).
- **`Status`**: Текущий статус предложения, использует тип `BidStatus`.
- **`CreatedAt`**: Время создания предложения.
- **`Version`**: Версия предложения, которая может использоваться для отслеживания изменений в предложении.

### Модель **`BidReview`** (Отзыв на предложение):

```go
type BidReview struct {
    ID          string    `json:"id" gorm:"primary_key"`
    BidID       string    `json:"bidId"`
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
}
```

Эта структура описывает отзыв на предложение:
- **`ID`**: Уникальный идентификатор отзыва. Является первичным ключом в базе данных.
- **`BidID`**: Идентификатор предложения, к которому относится этот отзыв. Это поле создаёт связь между отзывом и предложением.
- **`Description`**: Описание отзыва (например, текст обратной связи).
- **`CreatedAt`**: Время создания отзыва.

### Итог:
- **`Bid`**: Основная модель для представления предложений в тендере. Включает информацию о названии, описании, сумме, авторе, статусе и версии предложения.
- **`BidReview`**: Модель для хранения отзывов на предложения. Связана с конкретным предложением через поле `BidID`.

Обе модели могут автоматически мигрировать в базу данных с помощью GORM, и к ним могут применяться CRUD операции (создание, чтение, обновление, удаление).

### tender.go
Этот код описывает модель тендера **`Tender`** в системе управления тендерами. Модель используется для работы с данными тендеров в базе данных через GORM (ORM для Go). Рассмотрим основные части кода и объясним работу каждой из них.

### Перечисления (Enums):

1. **`TenderStatus`**:
   ```go
   type TenderStatus string
   ```
   Это перечисление, которое описывает возможные статусы тендера:
   - **`TenderCreated`**: Тендер создан, но ещё не опубликован.
   - **`TenderPublished`**: Тендер опубликован и доступен для участников.
   - **`TenderClosed`**: Тендер завершён.

   Этот тип используется для обозначения текущего состояния тендера.

2. **`TenderServiceType`**:
   ```go
   type TenderServiceType string
   ```
   Это перечисление, которое определяет тип услуги, связанной с тендером:
   - **`Construction`**: Тендер связан со строительными услугами.
   - **`Delivery`**: Тендер связан с услугами доставки.
   - **`Manufacture`**: Тендер связан с производством.

   Это перечисление помогает классифицировать тендеры по типу услуг.

### Модель **`Tender`** (Тендер):

```go
type Tender struct {
    ID             string             `json:"id" gorm:"primary_key"`
    Name           string             `json:"name"`
    Description    string             `json:"description"`
    ServiceType    TenderServiceType  `json:"serviceType"`
    Status         TenderStatus       `json:"status"`
    OrganizationID string             `json:"organizationId"`
    Version        int                `json:"version"`
    CreatedAt      time.Time          `json:"createdAt"`
}
```

#### Поля модели:

1. **`ID`**: Уникальный идентификатор тендера (строка). Это поле является первичным ключом в базе данных.
2. **`Name`**: Название тендера.
3. **`Description`**: Описание тендера, в котором могут быть указаны детали тендерного проекта.
4. **`ServiceType`**: Тип услуги, связанной с тендером (использует перечисление `TenderServiceType`).
5. **`Status`**: Текущий статус тендера (использует перечисление `TenderStatus`).
6. **`OrganizationID`**: Идентификатор организации, которая создала тендер. Это поле связывает тендер с организацией.
7. **`Version`**: Версия тендера, используется для отслеживания изменений и контроля версий.
8. **`CreatedAt`**: Время создания тендера. Это поле хранит метку времени, когда тендер был создан.

### Итог:

Модель **`Tender`** используется для хранения и работы с тендерами в базе данных. В ней определены основные атрибуты тендера: название, описание, тип услуги, статус и идентификатор организации. Поле версии позволяет отслеживать изменения в тендере, а время создания указывает на момент его регистрации в системе.

Эта структура может использоваться в различных операциях, таких как создание, обновление и получение информации о тендерах.

## routes:

### bid_service.go:
Этот код описывает сервисные функции для работы с предложениями (**`Bid`**) в системе управления тендерами. Каждая функция взаимодействует с базой данных через GORM (ORM для Go). Рассмотрим каждую функцию подробно.

### 1. **`GetBids`** — Получение списка предложений для тендера с пагинацией
```go
func GetBids(tenderID, limit, offset string) ([]models.Bid, error) {
    lim, off := 10, 0 // Значения по умолчанию
    var err error

    if limit != "" {
        lim, err = strconv.Atoi(limit)
        if err != nil {
            return nil, errors.New("invalid limit parameter")
        }
    }

    if offset != "" {
        off, err = strconv.Atoi(offset)
        if err != nil {
            return nil, errors.New("invalid offset parameter")
        }
    }

    var bids []models.Bid
    if err := db.DB.Where("tender_id = ?", tenderID).Offset(off).Limit(lim).Find(&bids).Error; err != nil {
        return nil, err
    }

    return bids, nil
}
```
- Функция получает предложения, связанные с конкретным тендером, используя идентификатор тендера (`tenderID`).
- Параметры **`limit`** и **`offset`** позволяют задавать количество возвращаемых записей и смещение для пагинации.
- Значения **`limit`** и **`offset`** по умолчанию: 10 и 0 соответственно.
- Используется GORM-запрос для выборки предложений из базы данных, связанных с определённым тендером.

### 2. **`CreateBid`** — Создание нового предложения
```go
func CreateBid(bid *models.Bid) (*models.Bid, error) {
    bid.ID = uuid.New().String() // Генерация уникального идентификатора для предложения

    if err := validateBid(bid); err != nil {
        return nil, err
    }

    if err := db.DB.Create(bid).Error; err != nil {
        return nil, err
    }

    return bid, nil
}
```
- Создаёт новое предложение.
- Генерируется уникальный идентификатор для предложения с помощью **UUID**.
- Валидация данных предложения через функцию **`validateBid`**.
- После успешной валидации, предложение сохраняется в базу данных.

### 3. **`GetBid`** — Получение предложения по его ID
```go
func GetBid(bidID string) (*models.Bid, error) {
    var bid models.Bid
    if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("bid not found")
        }
        return nil, err
    }
    return &bid, nil
}
```
- Возвращает предложение по его **ID**.
- Используется GORM для поиска первой записи, соответствующей ID предложения.
- Если предложение не найдено, возвращается ошибка **`bid not found`**.

### 4. **`UpdateBid`** — Обновление предложения
```go
func UpdateBid(bidID string, updates map[string]interface{}) (*models.Bid, error) {
    var bid models.Bid

    if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("bid not found")
        }
        return nil, err
    }

    if err := db.DB.Model(&bid).Updates(updates).Error; err != nil {
        return nil, err
    }

    return &bid, nil
}
```
- Обновляет предложение по его **ID**.
- Сначала выполняется поиск предложения, затем обновляются только те поля, которые указаны в **`updates`** (карта с полями и значениями).
- Если предложение не найдено, возвращается ошибка **`bid not found`**.

### 5. **`DeleteBid`** — Удаление предложения по его ID
```go
func DeleteBid(bidID string) error {
    if err := db.DB.Where("id = ?", bidID).Delete(&models.Bid{}).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return errors.New("bid not found")
        }
        return err
    }
    return nil
}
```
- Удаляет предложение по его **ID**.
- Если предложение не найдено, возвращается ошибка **`bid not found`**.

### 6. **`validateBid`** — Валидация данных предложения
```go
func validateBid(bid *models.Bid) error {
    if bid.Name == "" {
        return errors.New("name cannot be empty")
    }
    if bid.Description == "" {
        return errors.New("description cannot be empty")
    }
    if bid.TenderID == "" {
        return errors.New("tenderId cannot be empty")
    }
    return nil
}
```
- Проверяет корректность данных предложения:
  - Название (**Name**) не должно быть пустым.
  - Описание (**Description**) не должно быть пустым.
  - Идентификатор тендера (**TenderID**) не должен быть пустым.
  
Эта функция вызывается перед сохранением нового предложения для проверки его корректности.

### tender_service.go:
Этот код описывает сервисные функции для управления тендерами (**`Tender`**) в системе управления тендерами. Каждая функция взаимодействует с базой данных через GORM (ORM для Go). Давайте рассмотрим каждую функцию подробнее.

### 1. **`GetTenders`** — Получение списка тендеров с фильтрацией и пагинацией
```go
func GetTenders(serviceType, limit, offset string) ([]models.Tender, error) {
    lim, off := 10, 0 // Значения по умолчанию
    var err error

    if limit != "" {
        lim, err = strconv.Atoi(limit)
        if err != nil {
            return nil, errors.New("invalid limit parameter")
        }
    }

    if offset != "" {
        off, err = strconv.Atoi(offset)
        if err != nil {
            return nil, errors.New("invalid offset parameter")
        }
    }

    // Создание запроса с учетом фильтрации по типу услуг
    query := db.DB.Model(&models.Tender{}).Offset(off).Limit(lim)
    if serviceType != "" {
        query = query.Where("service_type IN (?)", serviceType)
    }

    var tenders []models.Tender
    if err := query.Find(&tenders).Error; err != nil {
        return nil, err
    }

    return tenders, nil
}
```
- Функция получает список тендеров с возможностью фильтрации по типу услуг (`serviceType`) и пагинации через параметры **`limit`** и **`offset`**.
- Параметры по умолчанию: **10** для **`limit`** и **0** для **`offset`**.
- Используется GORM для создания запроса, который фильтрует тендеры по типу услуг и затем применяет пагинацию. Если **`serviceType`** не пустой, добавляется фильтр.

### 2. **`CreateTender`** — Создание нового тендера
```go
func CreateTender(tender *models.Tender) (*models.Tender, error) {
    tender.ID = uuid.New().String() // Генерация уникального идентификатора

    if err := validateTender(tender); err != nil {
        return nil, err
    }

    if err := db.DB.Create(tender).Error; err != nil {
        return nil, err
    }

    return tender, nil
}
```
- Создаёт новый тендер.
- Генерирует уникальный идентификатор для тендера с помощью **UUID**.
- Вызывает функцию **`validateTender`** для проверки корректности данных перед сохранением.
- Сохраняет тендер в базу данных.

### 3. **`GetUserTenders`** — Получение тендеров пользователя с пагинацией
```go
func GetUserTenders(username, limit, offset string) ([]models.Tender, error) {
    lim, off := 10, 0
    var err error

    if limit != "" {
        lim, err = strconv.Atoi(limit)
        if err != nil {
            return nil, errors.New("invalid limit parameter")
        }
    }

    if offset != "" {
        off, err = strconv.Atoi(offset)
        if err != nil {
            return nil, errors.New("invalid offset parameter")
        }
    }

    var tenders []models.Tender
    if err := db.DB.Where("creator_username = ?", username).Offset(off).Limit(lim).Find(&tenders).Error; err != nil {
        return nil, err
    }

    return tenders, nil
}
```
- Получает тендеры, созданные пользователем с указанным **`username`**.
- Параметры **`limit`** и **`offset`** применяются для пагинации.
- Выполняет GORM-запрос, чтобы найти тендеры, которые были созданы указанным пользователем.

### 4. **`GetTenderStatus`** — Получение статуса тендера
```go
func GetTenderStatus(tenderID, username string) (models.TenderStatus, error) {
    var tender models.Tender

    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return "", errors.New("tender not found")
        }
        return "", err
    }

    return tender.Status, nil
}
```
- Возвращает статус тендера по его **ID** и имени создателя (**username**).
- Выполняет поиск тендера с использованием GORM. Если тендер не найден, возвращается ошибка.

### 5. **`UpdateTenderStatus`** — Изменение статуса тендера
```go
func UpdateTenderStatus(tenderID, status, username string) (*models.Tender, error) {
    var tender models.Tender

    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender not found")
        }
        return nil, err
    }

    tender.Status = models.TenderStatus(status)
    if err := db.DB.Save(&tender).Error; err != nil {
        return nil, err
    }

    return &tender, nil
}
```
- Изменяет статус тендера.
- Сначала выполняется поиск тендера, затем статус обновляется с помощью метода **Save** из GORM.

### 6. **`EditTender`** — Редактирование параметров тендера
```go
func EditTender(tenderID, username string, updates map[string]interface{}) (*models.Tender, error) {
    var tender models.Tender

    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender not found")
        }
        return nil, err
    }

    if err := db.DB.Model(&tender).Updates(updates).Error; err != nil {
        return nil, err
    }

    return &tender, nil
}
```
- Позволяет редактировать параметры тендера.
- Ищет тендер по ID и имени создателя. Обновляет поля тендера, используя карту **updates**.

### 7. **`RollbackTender`** — Откат тендера к указанной версии
```go
func RollbackTender(tenderID string, version int, username string) (*models.Tender, error) {
    var tender models.Tender

    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender not found")
        }
        return nil, err
    }

    if version >= tender.Version {
        return nil, errors.New("invalid version number")
    }

    var oldTender models.Tender
    if err := db.DB.Where("id = ? AND version = ?", tenderID, version).First(&oldTender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender version not found")
        }
        return nil, err
    }

    oldTender.Version++ // Обновление версии
    if err := db.DB.Save(&oldTender).Error; err != nil {
        return nil, err
    }

    return &oldTender, nil
}
```
- Откатывает параметры тендера к указанной версии.
- Сначала ищет текущий тендер, затем проверяет, является ли запрашиваемая версия действительной.
- Выполняет поиск старой версии тендера и обновляет текущий тендер до старой версии.

### 8. **`validateTender`** — Валидация данных тендера
```go
func validateTender(tender *models.Tender) error {
    if tender.Name == "" {
        return errors.New("name cannot be empty")
    }
    if tender.Description == "" {
        return errors.New("description cannot be empty")
    }
    return nil
}
```
- Проверяет корректность данных тендера.
- Возвращает ошибку, если название или описание пустые.

Эти функции предоставляют основные CRUD (Create, Read, Update, Delete) операции для работы с тендерами и обеспечивают валидацию и контроль доступа через имя пользователя.