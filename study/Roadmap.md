# Roadmap

## Блок 1: Основы языка и синтаксис

### 1.1 Базовый синтаксис

- [ ] Структура Go-программы
- [ ] Пакеты и импорты
- [ ] Функция main
- [ ] Комментарии (однострочные и многострочные)
- [ ] Точка с запятой (автоматическая вставка)
- [ ] Именование (camelCase, экспортируемые идентификаторы)

### 1.2 Переменные и константы

- [ ] Объявление переменных (var)
- [ ] Краткое объявление (:=)
- [ ] Множественное присваивание
- [ ] Нулевые значения (zero values)
- [ ] Область видимости переменных
- [ ] Константы (const)
- [ ] Перечисления с iota
- [ ] Типизированные и нетипизированные константы

### 1.3 Базовые типы данных

- [ ] Булевый тип (bool)
- [ ] Целочисленные типы (int8, int16, int32, int64)
- [ ] Беззнаковые типы (uint8, uint16, uint32, uint64)
- [ ] Платформозависимые типы (int, uint, uintptr)
- [ ] Байтовый тип (byte как alias для uint8)
- [ ] Руна (rune как alias для int32)
- [ ] Типы с плавающей точкой (float32, float64)
- [ ] Комплексные числа (complex64, complex128)
- [ ] Операции с комплексными числами

### 1.4 Строки

- [ ] Строковый тип (string)
- [ ] Неизменяемость строк
- [ ] Строковые литералы (интерпретируемые и сырые)
- [ ] Длина строки (len)
- [ ] Индексация строк
- [ ] Срезы строк
- [ ] UTF-8 кодировка
- [ ] Руны и байты в строках
- [ ] Итерация по строкам (range)

### 1.5 Операторы

- [ ] Арифметические операторы (+, -, *, /, %)
- [ ] Операторы сравнения (==, !=, <, >, <=, >=)
- [ ] Логические операторы (&&, ||, !)
- [ ] Битовые операторы (&, |, ^, &^, <<, >>)
- [ ] Операторы присваивания (=, +=, -=, *=, /=, и т.д.)
- [ ] Инкремент и декремент (++, --)
- [ ] Адресные операторы (&, *)

### 1.6 Приведение типов

- [ ] Явное приведение типов
- [ ] Преобразование между числовыми типами
- [ ] Преобразование строк в байты и руны
- [ ] Преобразование чисел в строки и обратно

## Блок 2: Управляющие конструкции

### 2.1 Условные операторы

- [ ] Оператор if
- [ ] if с инициализацией
- [ ] if-else цепочки
- [ ] if-else if-else
- [ ] Отсутствие тернарного оператора

### 2.2 Switch

- [ ] Базовый switch
- [ ] Switch с выражением
- [ ] Switch без выражения (как альтернатива if-else)
- [ ] Множественные значения в case
- [ ] Fallthrough
- [ ] Type switch

### 2.3 Циклы

- [ ] Цикл for (классический)
- [ ] For как while
- [ ] Бесконечный цикл
- [ ] For-range для массивов
- [ ] For-range для слайсов
- [ ] For-range для строк
- [ ] For-range для map
- [ ] For-range для каналов
- [ ] Break и continue
- [ ] Метки (labels) для вложенных циклов

### 2.4 Управление потоком

- [ ] Goto и метки
- [ ] Return в функциях
- [ ] Defer (отложенное выполнение)
- [ ] Panic (аварийное завершение)
- [ ] Recover (восстановление после panic)

## Блок 3: Структуры данных

### 3.1 Массивы

- [ ] Объявление массивов
- [ ] Инициализация массивов
- [ ] Доступ к элементам
- [ ] Длина массива (len)
- [ ] Многомерные массивы
- [ ] Массивы как значения (копирование)
- [ ] Сравнение массивов
- [ ] Итерация по массивам

### 3.2 Слайсы

- [ ] Объявление слайсов
- [ ] Создание слайсов с make
- [ ] Слайсы как срезы массивов
- [ ] Длина (len) и ёмкость (cap)
- [ ] Append для добавления элементов
- [ ] Стратегия роста ёмкости
- [ ] Copy для копирования слайсов
- [ ] Многомерные слайсы
- [ ] Нулевой слайс vs пустой слайс
- [ ] Внутреннее устройство слайсов (указатель, длина, ёмкость)
- [ ] Re-slicing
- [ ] Slice tricks (удаление, вставка, фильтрация)

### 3.3 Map (ассоциативные массивы)

- [ ] Объявление map
- [ ] Создание map с make
- [ ] Добавление элементов
- [ ] Получение элементов
- [ ] Проверка существования ключа
- [ ] Удаление элементов (delete)
- [ ] Длина map (len)
- [ ] Итерация по map
- [ ] Нулевой map vs пустой map
- [ ] Map не является потокобезопасным
- [ ] Внутреннее устройство map (bucket, hash function)
- [ ] Порядок итерации не гарантирован

### 3.4 Структуры

- [ ] Объявление структур
- [ ] Инициализация структур
- [ ] Анонимные структуры
- [ ] Доступ к полям структуры
- [ ] Указатели на структуры
- [ ] Вложенные структуры
- [ ] Встраивание структур (embedding)
- [ ] Теги полей структур (tags)
- [ ] Экспортируемые и неэкспортируемые поля
- [ ] Сравнение структур

## Блок 4: Функции

### 4.1 Основы функций

- [ ] Объявление функций
- [ ] Параметры функций
- [ ] Возвращаемые значения
- [ ] Множественные возвращаемые значения
- [ ] Именованные возвращаемые значения
- [ ] Пустой идентификатор (_) для игнорирования значений
- [ ] Рекурсия

### 4.2 Вариативные функции

- [ ] Вариативные параметры (...)
- [ ] Распаковка слайса в вариативную функцию

### 4.3 Функции высшего порядка

- [ ] Функции как значения
- [ ] Функции как параметры
- [ ] Функции как возвращаемые значения
- [ ] Анонимные функции
- [ ] Замыкания (closures)

### 4.4 Методы

- [ ] Методы на типах
- [ ] Методы на структурах
- [ ] Указательные receivers
- [ ] Значимые receivers
- [ ] Выбор между указательным и значимым receiver
- [ ] Методы на не-структурных типах

### 4.5 Defer, Panic, Recover

- [ ] Порядок выполнения defer
- [ ] Defer с именованными возвращаемыми значениями
- [ ] Defer для очистки ресурсов
- [ ] Panic для критических ошибок
- [ ] Recover для перехвата panic
- [ ] Паттерны использования recover

## Блок 5: Интерфейсы и полиморфизм

### 5.1 Основы интерфейсов

- [ ] Объявление интерфейсов
- [ ] Неявная реализация интерфейсов
- [ ] Пустой интерфейс (interface{})
- [ ] Type assertion
- [ ] Type assertion с проверкой
- [ ] Type switch для интерфейсов

### 5.2 Работа с интерфейсами

- [ ] Интерфейсы как параметры функций
- [ ] Интерфейсы как возвращаемые значения
- [ ] Композиция интерфейсов
- [ ] Встраивание интерфейсов
- [ ] Сравнение интерфейсов

### 5.3 Стандартные интерфейсы

- [ ] io.Reader
- [ ] io.Writer
- [ ] io.Closer
- [ ] fmt.Stringer
- [ ] error интерфейс
- [ ] sort.Interface

### 5.4 Внутреннее устройство

- [ ] Структура интерфейса (type, value)
- [ ] Nil интерфейсы
- [ ] Интерфейсы с nil значениями
- [ ] Динамический vs статический тип

## Блок 6: Указатели и память

### 6.1 Указатели

- [ ] Объявление указателей
- [ ] Оператор взятия адреса (&)
- [ ] Оператор разыменования (*)
- [ ] Nil указатели
- [ ] Указатели на структуры
- [ ] Автоматическое разыменование полей структур
- [ ] Передача по значению vs по ссылке

### 6.2 Выделение памяти

- [ ] Функция new
- [ ] Функция make
- [ ] Разница между new и make
- [ ] Composite literals

### 6.3 Управление памятью

- [ ] Стек vs куча
- [ ] Escape analysis
- [ ] Сборщик мусора (GC)
- [ ] Как работает GC в Go
- [ ] Оптимизация использования памяти

## Блок 7: Пакеты и модули

### 7.1 Пакеты

- [ ] Создание пакетов
- [ ] Импорт пакетов
- [ ] Псевдонимы при импорте
- [ ] Dot импорт
- [ ] Blank import (_)
- [ ] Экспортируемые и неэкспортируемые идентификаторы
- [ ] Функция init
- [ ] Порядок инициализации пакетов
- [ ] Внутренние пакеты (internal)

### 7.2 Go Modules

- [ ] Инициализация модуля (go mod init)
- [ ] go.mod файл
- [ ] go.sum файл
- [ ] Добавление зависимостей (go get)
- [ ] Обновление зависимостей
- [ ] Удаление неиспользуемых зависимостей (go mod tidy)
- [ ] Vendor директория
- [ ] Семантическое версионирование
- [ ] Replace директива
- [ ] Exclude директива

### 7.3 Организация проекта

- [ ] Стандартная структура проекта
- [ ] cmd директория
- [ ] pkg директория
- [ ] internal директория
- [ ] api директория
- [ ] Naming conventions

## Блок 8: Обработка ошибок

### 8.1 Базовая обработка ошибок

- [ ] error интерфейс
- [ ] errors.New
- [ ] fmt.Errorf
- [ ] Возврат ошибок из функций
- [ ] Проверка ошибок
- [ ] Игнорирование ошибок (антипаттерн)

### 8.2 Кастомные ошибки

- [ ] Создание кастомных типов ошибок
- [ ] Реализация error интерфейса
- [ ] Ошибки со структурами

### 8.3 Продвинутая обработка ошибок

- [ ] errors.Is (Go 1.13+)
- [ ] errors.As (Go 1.13+)
- [ ] Wrapping ошибок с %w
- [ ] errors.Unwrap
- [ ] Sentinel errors
- [ ] Паттерны обработки ошибок

## Блок 9: Конкурентность (Concurrency)

### 9.1 Основы многопоточности

- [ ] Процессы и потоки
- [ ] Concurrency vs Parallelism
- [ ] Зачем нужна конкурентность
- [ ] Модель CSP (Communicating Sequential Processes)

### 9.2 Горутины

- [ ] Создание горутин (go keyword)
- [ ] Легковесность горутин
- [ ] Планировщик горутин (scheduler)
- [ ] GOMAXPROCS
- [ ] Runtime.NumCPU
- [ ] Runtime.NumGoroutine
- [ ] Завершение горутин
- [ ] Утечки горутин (goroutine leaks)

### 9.3 Каналы

- [ ] Создание каналов
- [ ] Отправка в канал
- [ ] Получение из канала
- [ ] Буферизованные каналы
- [ ] Небуферизованные каналы
- [ ] Закрытие каналов
- [ ] Проверка закрытия канала
- [ ] Range по каналам
- [ ] Однонаправленные каналы (send-only, receive-only)

### 9.4 Select

- [ ] Оператор select
- [ ] Select с несколькими каналами
- [ ] Default case в select
- [ ] Timeout с select
- [ ] Non-blocking операции с select

### 9.5 Паттерны конкурентности

- [ ] Worker pool
- [ ] Fan-out, Fan-in
- [ ] Pipeline
- [ ] Cancellation через каналы
- [ ] Timeout паттерн
- [ ] Rate limiting
- [ ] Semaphore паттерн
- [ ] Or-channel паттерн
- [ ] Bridge-channel паттерн

### 9.6 Sync пакет

- [ ] sync.Mutex
- [ ] sync.RWMutex
- [ ] sync.WaitGroup
- [ ] sync.Once
- [ ] sync.Cond
- [ ] sync.Pool
- [ ] sync.Map

### 9.7 Atomic операции

- [ ] sync/atomic пакет
- [ ] atomic.AddInt32/64
- [ ] atomic.LoadInt32/64
- [ ] atomic.StoreInt32/64
- [ ] atomic.CompareAndSwapInt32/64
- [ ] atomic.Value

### 9.8 Context

- [ ] context.Context интерфейс
- [ ] context.Background
- [ ] context.TODO
- [ ] context.WithCancel
- [ ] context.WithTimeout
- [ ] context.WithDeadline
- [ ] context.WithValue
- [ ] Передача контекста через функции
- [ ] Best practices для context

### 9.9 Продвинутые темы

- [ ] Race detector
- [ ] Data races
- [ ] Happens-before отношения
- [ ] Memory model Go
- [ ] Graceful shutdown
- [ ] Координация горутин

## Блок 10: Стандартная библиотека

### 10.1 fmt пакет

- [ ] Print, Println, Printf
- [ ] Sprint, Sprintln, Sprintf
- [ ] Fprint, Fprintln, Fprintf
- [ ] Scan, Scanln, Scanf
- [ ] Форматирование строк (verbs)
- [ ] Форматирование кастомных типов (Stringer)

### 10.2 strings пакет

- [ ] Contains, ContainsAny, ContainsRune
- [ ] Count
- [ ] Fields, FieldsFunc
- [ ] HasPrefix, HasSuffix
- [ ] Index, IndexAny, LastIndex
- [ ] Join
- [ ] Repeat
- [ ] Replace, ReplaceAll
- [ ] Split, SplitN, SplitAfter
- [ ] ToLower, ToUpper
- [ ] Trim, TrimSpace, TrimPrefix, TrimSuffix
- [ ] strings.Builder для эффективной конкатенации

### 10.3 strconv пакет

- [ ] Atoi, Itoa
- [ ] ParseInt, ParseUint, ParseFloat, ParseBool
- [ ] FormatInt, FormatUint, FormatFloat, FormatBool
- [ ] Quote, Unquote

### 10.4 math пакет

- [ ] Константы (Pi, E)
- [ ] Abs, Max, Min
- [ ] Pow, Sqrt
- [ ] Ceil, Floor, Round
- [ ] Sin, Cos, Tan
- [ ] Log, Log10, Log2, Exp
- [ ] math/rand для генерации случайных чисел

### 10.5 time пакет

- [ ] time.Time тип
- [ ] time.Now
- [ ] time.Date
- [ ] Форматирование времени (Layout)
- [ ] Парсинг времени (Parse)
- [ ] time.Duration
- [ ] Sleep
- [ ] After, AfterFunc
- [ ] Ticker
- [ ] Timer
- [ ] Операции с временем (Add, Sub, Before, After)
- [ ] Часовые пояса (Location)

### 10.6 sort пакет

- [ ] Сортировка слайсов (Ints, Float64s, Strings)
- [ ] sort.Slice
- [ ] sort.SliceStable
- [ ] sort.Interface
- [ ] IsSorted
- [ ] Reverse
- [ ] Search, SearchInts, SearchFloat64s, SearchStrings

### 10.7 regexp пакет

- [ ] Компиляция регулярных выражений (Compile, MustCompile)
- [ ] Match, MatchString
- [ ] Find, FindAll
- [ ] FindString, FindAllString
- [ ] ReplaceAll, ReplaceAllString
- [ ] Split

## Блок 11: Работа с файлами и I/O

### 11.1 os пакет

- [ ] os.File тип
- [ ] Open, OpenFile
- [ ] Create
- [ ] Read, Write
- [ ] ReadAt, WriteAt
- [ ] Seek
- [ ] Close
- [ ] Stat, Lstat
- [ ] os.FileInfo интерфейс
- [ ] Chmod, Chown
- [ ] Remove, RemoveAll
- [ ] Rename
- [ ] Mkdir, MkdirAll
- [ ] Getwd, Chdir
- [ ] Environment variables (Getenv, Setenv, Environ)

### 11.2 io пакет

- [ ] io.Reader интерфейс
- [ ] io.Writer интерфейс
- [ ] io.Closer интерфейс
- [ ] io.ReadWriter
- [ ] io.ReadCloser, io.WriteCloser
- [ ] Copy
- [ ] CopyN
- [ ] ReadAll
- [ ] ReadFull
- [ ] WriteString
- [ ] MultiReader, MultiWriter
- [ ] TeeReader
- [ ] Pipe

### 11.3 bufio пакет

- [ ] bufio.Reader
- [ ] bufio.Writer
- [ ] bufio.Scanner
- [ ] NewReader, NewWriter
- [ ] ReadString, ReadBytes, ReadLine
- [ ] Scan, ScanLines, ScanWords
- [ ] Flush
- [ ] Буферизация для производительности

### 11.4 ioutil пакет (устарел, но знать нужно)

- [ ] ReadFile
- [ ] WriteFile
- [ ] ReadAll
- [ ] ReadDir
- [ ] TempDir, TempFile

### 11.5 os/exec пакет

- [ ] Command
- [ ] CombinedOutput
- [ ] Output
- [ ] Run
- [ ] Start, Wait
- [ ] StdinPipe, StdoutPipe, StderrPipe
- [ ] Env, Dir

### 11.6 path/filepath пакет

- [ ] Join
- [ ] Split
- [ ] Dir, Base
- [ ] Ext
- [ ] Abs
- [ ] Clean
- [ ] Match, Glob
- [ ] Walk
- [ ] WalkDir (Go 1.16+)

### 11.7 embed пакет (Go 1.16+)

- [ ] //go:embed директива
- [ ] embed.FS
- [ ] Встраивание файлов в бинарник

## Блок 12: Тестирование

### 12.1 testing пакет

- [ ] Написание тестов
- [ ] Функции Test*
- [ ] testing.T
- [ ] t.Error, t.Errorf
- [ ] t.Fatal, t.Fatalf
- [ ] t.Log, t.Logf
- [ ] t.Run для subtests
- [ ] Table-driven tests
- [ ] Запуск тестов (go test)
- [ ] Флаги go test (-v, -run, -count)

### 12.2 Бенчмарки

- [ ] Функции Benchmark*
- [ ] testing.B
- [ ] b.N для итераций
- [ ] b.ResetTimer
- [ ] b.StopTimer, b.StartTimer
- [ ] Запуск бенчмарков (go test -bench)
- [ ] Анализ результатов бенчмарков

### 12.3 Examples

- [ ] Функции Example*
- [ ] Output комментарии
- [ ] Testable examples

### 12.4 Coverage

- [ ] Генерация покрытия (go test -cover)
- [ ] Детальный отчёт (go test -coverprofile)
- [ ] Визуализация покрытия (go tool cover)

### 12.5 Test helpers

- [ ] t.Helper
- [ ] Создание вспомогательных функций для тестов
- [ ] Setup и teardown

### 12.6 Mocking

- [ ] Интерфейсы для моков
- [ ] Ручное создание моков
- [ ] Использование библиотек для моков (testify, gomock)

### 12.7 Integration testing

- [ ] Тестирование с реальными зависимостями
- [ ] Использование Docker для тестов
- [ ] Test containers

## Блок 13: Продвинутые возможности языка

### 13.1 Generics (Go 1.18+)

- [ ] Type parameters
- [ ] Type constraints
- [ ] Создание generic функций
- [ ] Создание generic типов
- [ ] any (alias для interface{})
- [ ] comparable constraint
- [ ] Кастомные constraints
- [ ] Type inference

### 13.2 Reflection

- [ ] reflect пакет
- [ ] reflect.Type
- [ ] reflect.Value
- [ ] TypeOf, ValueOf
- [ ] Kind
- [ ] NumField, Field
- [ ] NumMethod, Method
- [ ] Set методы
- [ ] MakeSlice, MakeMap, MakeFunc
- [ ] Когда использовать reflection
- [ ] Performance implications

### 13.3 Unsafe

- [ ] unsafe пакет
- [ ] unsafe.Pointer
- [ ] uintptr
- [ ] Размер типов (Sizeof, Alignof, Offsetof)
- [ ] Когда использовать unsafe
- [ ] Риски использования unsafe

### 13.4 Build tags

- [ ] //go:build директива
- [ ] Условная компиляция
- [ ] Теги для разных ОС и архитектур
- [ ] Кастомные build tags

### 13.5 Code generation

- [ ] go generate
- [ ] //go:generate директива
- [ ] Использование text/template
- [ ] Инструменты кодогенерации (stringer и др.)

## Блок 14: Сетевое программирование

### 14.1 net пакет

- [ ] Типы адресов (IPAddr, TCPAddr, UDPAddr)
- [ ] ResolveTCPAddr, ResolveUDPAddr
- [ ] Dial, DialTimeout
- [ ] Listen
- [ ] Accept
- [ ] net.Conn интерфейс
- [ ] net.Listener интерфейс

### 14.2 TCP

- [ ] TCP клиент
- [ ] TCP сервер
- [ ] Чтение и запись данных
- [ ] Обработка множественных соединений
- [ ] Graceful shutdown TCP сервера

### 14.3 UDP

- [ ] UDP клиент
- [ ] UDP сервер
- [ ] Отличия от TCP

### 14.4 HTTP клиент

- [ ] net/http пакет
- [ ] http.Client
- [ ] GET запросы (http.Get)
- [ ] POST запросы (http.Post, http.PostForm)
- [ ] Кастомные запросы (http.NewRequest)
- [ ] Добавление заголовков
- [ ] Таймауты
- [ ] Редиректы
- [ ] Cookies
- [ ] Multipart/form-data

### 14.5 HTTP сервер

- [ ] http.Server
- [ ] http.ListenAndServe
- [ ] http.Handle, http.HandleFunc
- [ ] http.Handler интерфейс
- [ ] http.HandlerFunc
- [ ] http.Request
- [ ] http.ResponseWriter
- [ ] Роутинг запросов
- [ ] Middleware паттерн
- [ ] Graceful shutdown HTTP сервера

### 14.6 HTTP продвинутые темы

- [ ] HTTPS (TLS)
- [ ] HTTP/2
- [ ] Keep-alive
- [ ] Chunked transfer encoding
- [ ] Server-Sent Events (SSE)
- [ ] File upload/download

### 14.7 WebSocket

- [ ] Протокол WebSocket
- [ ] Установка соединения
- [ ] Отправка и получение сообщений
- [ ] Библиотеки для WebSocket (gorilla/websocket)

### 14.8 DNS

- [ ] net.LookupHost
- [ ] net.LookupIP
- [ ] net.LookupCNAME
- [ ] net.LookupMX

## Блок 15: REST API разработка

### 15.1 Основы REST

- [ ] Принципы REST
- [ ] HTTP методы (GET, POST, PUT, PATCH, DELETE)
- [ ] Статус коды HTTP
- [ ] Headers
- [ ] Query parameters
- [ ] Path parameters
- [ ] Request body
- [ ] Response body

### 15.2 JSON

- [ ] encoding/json пакет
- [ ] Marshal, Unmarshal
- [ ] MarshalIndent
- [ ] json.Encoder, json.Decoder
- [ ] Struct tags для JSON
- [ ] Опциональные поля (omitempty)
- [ ] Кастомная сериализация (MarshalJSON, UnmarshalJSON)
- [ ] json.RawMessage

### 15.3 HTTP роутеры

- [ ] Стандартный роутер (http.ServeMux)
- [ ] Сторонние роутеры (gorilla/mux, chi, gin, echo, fiber)
- [ ] Параметры в URL
- [ ] Группировка роутов
- [ ] Middleware для роутеров

### 15.4 Gin framework

- [ ] Установка Gin
- [ ] Создание роутера (gin.Default, gin.New)
- [ ] Определение роутов
- [ ] Path и query параметры
- [ ] Binding request body
- [ ] Validation
- [ ] Middleware в Gin
- [ ] Группировка роутов
- [ ] Рендеринг JSON, XML, HTML
- [ ] File upload
- [ ] Статические файлы

### 15.5 go-chi router

- [ ] Установка chi
- [ ] Создание роутера
- [ ] Middleware
- [ ] Subrouters
- [ ] URL параметры
- [ ] Context для передачи данных

### 15.6 API лучшие практики

- [ ] Versioning API
- [ ] Pagination
- [ ] Filtering, Sorting
- [ ] Rate limiting
- [ ] Authentication
- [ ] Authorization
- [ ] CORS
- [ ] API documentation

### 15.7 OpenAPI Specification

- [ ] Что такое OpenAPI (Swagger)
- [ ] Описание API в OpenAPI
- [ ] Генерация документации
- [ ] Инструменты (swaggo/swag)
- [ ] Swagger UI

### 15.8 Validation

- [ ] Валидация входных данных
- [ ] go-playground/validator
- [ ] Кастомные валидаторы
- [ ] Обработка ошибок валидации

## Блок 16: gRPC

### 16.1 Основы gRPC

- [ ] Что такое gRPC
- [ ] Преимущества gRPC
- [ ] gRPC vs REST
- [ ] HTTP/2 и gRPC

### 16.2 Protocol Buffers

- [ ] Что такое Protobuf
- [ ] .proto файлы
- [ ] Типы данных в Protobuf
- [ ] Messages
- [ ] Services
- [ ] Компиляция .proto файлов (protoc)
- [ ] Генерация Go кода

### 16.3 gRPC сервер

- [ ] Создание gRPC сервера
- [ ] Реализация сервисов
- [ ] Unary RPC
- [ ] Server streaming RPC
- [ ] Client streaming RPC
- [ ] Bidirectional streaming RPC

### 16.4 gRPC клиент

- [ ] Создание gRPC клиента
- [ ] Вызов методов сервера
- [ ] Обработка потоковых данных

### 16.5 gRPC продвинутые темы

- [ ] Interceptors (middleware для gRPC)
- [ ] Metadata
- [ ] Error handling в gRPC
- [ ] Deadlines и timeouts
- [ ] Authentication
- [ ] Load balancing
- [ ] Health checking

## Блок 17: Базы данных

### 17.1 SQL основы

- [ ] SELECT, INSERT, UPDATE, DELETE
- [ ] WHERE, ORDER BY, LIMIT, OFFSET
- [ ] JOIN (INNER, LEFT, RIGHT, FULL)
- [ ] GROUP BY, HAVING
- [ ] Агрегатные функции (COUNT, SUM, AVG, MAX, MIN)
- [ ] Подзапросы
- [ ] Индексы
- [ ] Транзакции

### 17.2 database/sql пакет

- [ ] sql.DB
- [ ] sql.Open
- [ ] Ping
- [ ] Connection pooling
- [ ] Query, QueryRow
- [ ] Exec
- [ ] Scan
- [ ] Named parameters
- [ ] sql.Rows
- [ ] sql.Result

### 17.3 PostgreSQL

- [ ] Драйвер для PostgreSQL (lib/pq, pgx)
- [ ] Подключение к PostgreSQL
- [ ] CRUD операции
- [ ] Транзакции (Begin, Commit, Rollback)
- [ ] Prepared statements
- [ ] SQL injection защита
- [ ] Работа с NULL значениями (sql.NullString и др.)

### 17.4 ORM - GORM

- [ ] Установка GORM
- [ ] Подключение к БД
- [ ] Модели
- [ ] Auto Migration
- [ ] CRUD операции через GORM
- [ ] Ассоциации (has one, has many, many to many)
- [ ] Запросы (Where, Or, Not)
- [ ] Preloading (Eager loading)
- [ ] Транзакции
- [ ] Hooks (BeforeCreate, AfterCreate и т.д.)

### 17.5 sqlx

- [ ] Установка sqlx
- [ ] sqlx.DB
- [ ] Get, Select
- [ ] Named queries
- [ ] Struct scanning

### 17.6 Миграции

- [ ] Что такое миграции
- [ ] golang-migrate/migrate
- [ ] Создание миграций
- [ ] Применение миграций
- [ ] Откат миграций
- [ ] Версионирование схемы БД

### 17.7 Redis

- [ ] Что такое Redis
- [ ] Use cases для Redis
- [ ] Типы данных в Redis (String, Hash, List, Set, Sorted Set)
- [ ] go-redis/redis клиент
- [ ] Подключение к Redis
- [ ] Set, Get операции
- [ ] Expiration (TTL)
- [ ] Pub/Sub
- [ ] Transactions
- [ ] Pipelining

### 17.8 MongoDB

- [ ] Что такое MongoDB
- [ ] NoSQL vs SQL
- [ ] mongo-driver/mongo
- [ ] Подключение к MongoDB
- [ ] Collections и Documents
- [ ] CRUD операции
- [ ] Фильтры и проекции
- [ ] Агрегации
- [ ] Индексы

## Блок 18: Message Queues

### 18.1 Apache Kafka

- [ ] Что такое Kafka
- [ ] Основные концепции (Topics, Partitions, Producers, Consumers)
- [ ] Use cases
- [ ] segmentio/kafka-go клиент
- [ ] Подключение к Kafka
- [ ] Producer (отправка сообщений)
- [ ] Consumer (получение сообщений)
- [ ] Consumer Groups
- [ ] Offset management
- [ ] Partition balancing

### 18.2 RabbitMQ

- [ ] Что такое RabbitMQ
- [ ] AMQP протокол
- [ ] Основные концепции (Exchange, Queue, Binding, Routing Key)
- [ ] streadway/amqp клиент
- [ ] Подключение к RabbitMQ
- [ ] Объявление очередей
- [ ] Publishing сообщений
- [ ] Consuming сообщений
- [ ] Exchange types (direct, fanout, topic, headers)
- [ ] Acknowledgments
- [ ] Durable queues

## Блок 19: Архитектура и паттерны проектирования

### 19.1 SOLID принципы

- [ ] Single Responsibility Principle
- [ ] Open/Closed Principle
- [ ] Liskov Substitution Principle
- [ ] Interface Segregation Principle
- [ ] Dependency Inversion Principle

### 19.2 KISS, DRY, YAGNI

- [ ] Keep It Simple, Stupid
- [ ] Don't Repeat Yourself
- [ ] You Aren't Gonna Need It

### 19.3 Design patterns

- [ ] Creational patterns (Singleton, Factory, Builder)
- [ ] Structural patterns (Adapter, Decorator, Facade)
- [ ] Behavioral patterns (Strategy, Observer, Command)

### 19.4 Архитектурные паттерны

- [ ] Layered architecture
- [ ] Clean architecture
- [ ] Hexagonal architecture (Ports and Adapters)
- [ ] Domain-Driven Design (DDD)
- [ ] CQRS (Command Query Responsibility Segregation)
- [ ] Event Sourcing

### 19.5 Dependency Injection

- [ ] Что такое DI
- [ ] Constructor injection
- [ ] DI контейнеры (wire, dig, fx)

### 19.6 Microservices

- [ ] Монолит vs Микросервисы
- [ ] Декомпозиция на микросервисы
- [ ] Service discovery
- [ ] API Gateway
- [ ] Circuit breaker
- [ ] Saga pattern
- [ ] Event-driven architecture

## Блок 20: Логирование и мониторинг

### 20.1 Логирование

- [ ] log пакет (стандартная библиотека)
- [ ] Structured logging
- [ ] zap logger
- [ ] zerolog
- [ ] logrus
- [ ] Уровни логирования (Debug, Info, Warn, Error, Fatal)
- [ ] Контекстное логирование
- [ ] Log rotation

### 20.2 ELK Stack

- [ ] Elasticsearch
- [ ] Logstash
- [ ] Kibana
- [ ] Интеграция с Go приложением

### 20.3 Метрики

- [ ] Что такое метрики
- [ ] expvar пакет
- [ ] Prometheus
- [ ] prometheus/client_golang
- [ ] Типы метрик (Counter, Gauge, Histogram, Summary)
- [ ] Экспорт метрик
- [ ] Grafana для визуализации

### 20.4 Tracing

- [ ] Distributed tracing
- [ ] OpenTelemetry
- [ ] Jaeger
- [ ] Spans и traces
- [ ] Интеграция с Go приложением

### 20.5 Healthchecks

- [ ] Liveness probe
- [ ] Readiness probe
- [ ] Healthcheck endpoint

## Блок 21: Docker и контейнеризация

### 21.1 Docker основы

- [ ] Что такое Docker
- [ ] Образы и контейнеры
- [ ] Dockerfile
- [ ] Инструкции Dockerfile (FROM, RUN, COPY, ADD, CMD, ENTRYPOINT)
- [ ] Сборка образа (docker build)
- [ ] Запуск контейнера (docker run)
- [ ] Docker registry

### 21.2 Dockerfile для Go

- [ ] Multi-stage builds
- [ ] Оптимизация размера образа
- [ ] Использование Alpine образов
- [ ] Кэширование слоёв
- [ ] .dockerignore

### 21.3 Docker Compose

- [ ] docker-compose.yml
- [ ] Определение сервисов
- [ ] Networks
- [ ] Volumes
- [ ] Environment variables
- [ ] Зависимости между сервисами
- [ ] Локальное тестирование (БД + MQ + App)

### 21.4 Docker best practices

- [ ] Безопасность образов
- [ ] Non-root пользователь
- [ ] Сканирование уязвимостей
- [ ] Health checks в Docker

## Блок 22: Kubernetes

### 22.1 Kubernetes основы

- [ ] Что такое Kubernetes
- [ ] Архитектура Kubernetes
- [ ] Pods
- [ ] Nodes
- [ ] Cluster

### 22.2 Kubernetes ресурсы

- [ ] Deployments
- [ ] Services (ClusterIP, NodePort, LoadBalancer)
- [ ] ConfigMaps
- [ ] Secrets
- [ ] PersistentVolumes и PersistentVolumeClaims
- [ ] Namespaces

### 22.3 kubectl

- [ ] Базовые команды kubectl
- [ ] apply, get, describe, delete
- [ ] logs, exec
- [ ] port-forward

### 22.4 YAML манифесты

- [ ] Написание манифестов для Deployment
- [ ] Написание манифестов для Service
- [ ] Написание манифестов для ConfigMap и Secret

### 22.5 Helm

- [ ] Что такое Helm
- [ ] Charts
- [ ] Установка приложений с Helm
- [ ] Создание собственных charts

### 22.6 Kubernetes для Go приложений

- [ ] Деплой Go приложения
- [ ] Liveness и Readiness probes
- [ ] Resource limits и requests
- [ ] Horizontal Pod Autoscaling
- [ ] Rolling updates
- [ ] Blue-green deployments

## Блок 23: CI/CD

### 23.1 Git

- [ ] Основы Git
- [ ] Branching strategies (Git Flow, GitHub Flow)
- [ ] Pull Requests
- [ ] Code review
- [ ] Merge vs Rebase

### 23.2 GitLab CI/CD

- [ ] .gitlab-ci.yml
- [ ] Pipelines
- [ ] Stages
- [ ] Jobs
- [ ] Runners
- [ ] Variables
- [ ] Artifacts
- [ ] Cache

### 23.3 CI/CD для Go

- [ ] Линтинг (golangci-lint)
- [ ] Тестирование
- [ ] Coverage reports
- [ ] Сборка бинарников
- [ ] Сборка Docker образов
- [ ] Deployment

### 23.4 GitHub Actions

- [ ] Workflows
- [ ] Actions
- [ ] Events и triggers
- [ ] Jobs и steps
- [ ] Secrets

## Блок 24: Профилирование и оптимизация

### 24.1 Benchmarking

- [ ] Написание бенчмарков
- [ ] Интерпретация результатов
- [ ] benchstat
- [ ] Сравнение бенчмарков

### 24.2 Profiling

- [ ] runtime/pprof пакет
- [ ] net/http/pprof
- [ ] CPU profiling
- [ ] Memory profiling
- [ ] Goroutine profiling
- [ ] Block profiling
- [ ] Mutex profiling

### 24.3 pprof инструмент

- [ ] go tool pprof
- [ ] Интерактивный режим
- [ ] Web UI
- [ ] Flame graphs
- [ ] Анализ профилей

### 24.4 Трассировка

- [ ] runtime/trace пакет
- [ ] go tool trace
- [ ] Анализ трассировок

### 24.5 Оптимизация

- [ ] Анализ узких мест
- [ ] Оптимизация алгоритмов
- [ ] Оптимизация использования памяти
- [ ] Оптимизация конкурентности
- [ ] Оптимизация I/O
- [ ] Преждевременная оптимизация (антипаттерн)

### 24.6 Escape analysis

- [ ] Что такое escape analysis
- [ ] go build -gcflags="-m"
- [ ] Оптимизация аллокаций

## Блок 25: Продвинутые темы и internals

### 25.1 Go scheduler

- [ ] Модель M:N
- [ ] G (горутины), M (threads), P (processors)
- [ ] Work stealing
- [ ] Preemption

### 25.2 Garbage Collector

- [ ] Как работает GC в Go
- [ ] Tri-color marking
- [ ] Write barrier
- [ ] GC tuning (GOGC)
- [ ] Минимизация нагрузки на GC

### 25.3 Memory model

- [ ] Happens-before отношения
- [ ] Synchronization
- [ ] Channel communication
- [ ] Locks
- [ ] Atomic operations

### 25.4 Compiler

- [ ] Процесс компиляции
- [ ] Лексический анализ
- [ ] Парсинг
- [ ] Type checking
- [ ] Code generation
- [ ] Оптимизации компилятора

### 25.5 Assembly

- [ ] Go assembly синтаксис
- [ ] Когда использовать assembly
- [ ] Анализ сгенерированного assembly (go tool objdump)

### 25.6 Runtime

- [ ] runtime пакет
- [ ] Управление горутинами
- [ ] Управление памятью
- [ ] GOMAXPROCS
- [ ] GOGC
- [ ] GODEBUG

## Блок 26: Безопасность

### 26.1 Криптография

- [ ] crypto пакет
- [ ] Хэширование (MD5, SHA256, SHA512)
- [ ] HMAC
- [ ] AES шифрование
- [ ] RSA
- [ ] TLS/SSL

### 26.2 Authentication

- [ ] Basic Auth
- [ ] Token-based auth
- [ ] JWT (JSON Web Tokens)
- [ ] OAuth2
- [ ] Session management

### 26.3 Authorization

- [ ] Role-Based Access Control (RBAC)
- [ ] Attribute-Based Access Control (ABAC)
- [ ] Middleware для авторизации

### 26.4 Безопасность приложений

- [ ] SQL Injection защита
- [ ] XSS (Cross-Site Scripting) защита
- [ ] CSRF (Cross-Site Request Forgery) защита
- [ ] Валидация и санитизация входных данных
- [ ] Секретные данные (environment variables, vault)
- [ ] Rate limiting
- [ ] CORS

## Блок 27: Linux и системное программирование

### 27.1 Linux основы

- [ ] Файловая система
- [ ] Процессы
- [ ] Сигналы
- [ ] Права доступа
- [ ] Environment variables

### 27.2 Unix инструменты

- [ ] bash scripting
- [ ] ssh/rsync
- [ ] grep, sed, awk
- [ ] find
- [ ] Pipes и redirection

### 27.3 Системное программирование в Go

- [ ] syscall пакет
- [ ] os/signal для обработки сигналов
- [ ] Graceful shutdown на SIGTERM/SIGINT
- [ ] Working with processes

### 27.4 Make

- [ ] Makefile синтаксис
- [ ] Targets и dependencies
- [ ] Переменные
- [ ] Использование Make для Go проектов
- [ ] Общие таргеты (build, test, clean, run)

## Блок 28: Инструменты разработки

### 28.1 Go toolchain

- [ ] go build
- [ ] go run
- [ ] go install
- [ ] go get
- [ ] go mod
- [ ] go test
- [ ] go fmt
- [ ] go vet
- [ ] go doc
- [ ] go list

### 28.2 Линтеры

- [ ] golangci-lint
- [ ] Конфигурация линтеров
- [ ] Популярные линтеры (govet, staticcheck, errcheck, gosimple)
- [ ] Интеграция в CI/CD

### 28.3 Code formatting

- [ ] gofmt
- [ ] goimports
- [ ] Автоформатирование в IDE

### 28.4 Documentation

- [ ] Написание godoc комментариев
- [ ] go doc
- [ ] godoc server
- [ ] pkg.go.dev

### 28.5 IDE и редакторы

- [ ] VSCode с Go расширением
- [ ] GoLand
- [ ] Vim/Neovim с vim-go
- [ ] Настройка автодополнения (gopls)
- [ ] Debugging

## Блок 29: Лучшие практики и идиомы Go

### 29.1 Effective Go

- [ ] Форматирование кода
- [ ] Комментарии
- [ ] Именование
- [ ] Control structures
- [ ] Functions
- [ ] Data structures
- [ ] Initialization
- [ ] Methods
- [ ] Interfaces and other types
- [ ] Embedding
- [ ] Concurrency
- [ ] Errors
- [ ] Web servers

### 29.2 Code Review Comments

- [ ] Gofmt
- [ ] Comment Sentences
- [ ] Contexts
- [ ] Copying
- [ ] Crypto Rand
- [ ] Declaring Empty Slices
- [ ] Doc Comments
- [ ] Don't Panic
- [ ] Error Strings
- [ ] Handle Errors
- [ ] Imports
- [ ] In-Band Errors
- [ ] Indent Error Flow
- [ ] Initialisms
- [ ] Interfaces
- [ ] Line Length
- [ ] Mixed Caps
- [ ] Named Result Parameters
- [ ] Naked Returns
- [ ] Package Comments
- [ ] Package Names
- [ ] Pass Values
- [ ] Receiver Names
- [ ] Receiver Type
- [ ] Synchronous Functions
- [ ] Useful Test Failures
- [ ] Variable Names

### 29.3 Go Proverbs

- [ ] Don't communicate by sharing memory, share memory by communicating
- [ ] Concurrency is not parallelism
- [ ] Channels orchestrate; mutexes serialize
- [ ] The bigger the interface, the weaker the abstraction
- [ ] Make the zero value useful
- [ ] interface{} says nothing
- [ ] Gofmt's style is no one's favorite, yet gofmt is everyone's favorite
- [ ] A little copying is better than a little dependency
- [ ] Syscall must always be guarded with build tags
- [ ] Cgo must always be guarded with build tags
- [ ] Cgo is not Go
- [ ] With the unsafe package there are no guarantees
- [ ] Clear is better than clever
- [ ] Reflection is never clear
- [ ] Errors are values
- [ ] Don't just check errors, handle them gracefully
- [ ] Design the architecture, name the components, document the details
- [ ] Documentation is for users
- [ ] Don't panic

### 29.4 Common mistakes

- [ ] Range loop variable capture
- [ ] Slice and map modifications in loops
- [ ] Defer в циклах
- [ ] Неправильное использование goroutines
- [ ] Забытый mutex unlock
- [ ] Неправильная обработка ошибок
- [ ] Закрытие каналов отправителем, а не получателем

## Блок 30: Подготовка к собеседованиям

### 30.1 Алгоритмы и структуры данных

- [ ] Массивы и строки
- [ ] Linked Lists
- [ ] Stacks и Queues
- [ ] Hash Tables
- [ ] Trees (Binary Trees, BST, AVL, Red-Black)
- [ ] Graphs (представление, обходы)
- [ ] Heaps
- [ ] Tries
- [ ] Sorting algorithms (Quick, Merge, Heap, Bubble)
- [ ] Searching algorithms (Binary search, BFS, DFS)
- [ ] Dynamic Programming
- [ ] Greedy algorithms
- [ ] Backtracking

### 30.2 Системный дизайн

- [ ] Масштабируемость
- [ ] Load balancing
- [ ] Caching strategies
- [ ] Database sharding
- [ ] CAP theorem
- [ ] Consistency patterns
- [ ] Availability patterns
- [ ] CDN
- [ ] Message queues
- [ ] Микросервисы vs монолит
