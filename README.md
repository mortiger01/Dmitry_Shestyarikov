<h1> Web programming and mobile applications</h1>
<details>
  <summary><h2>🎄 Lab 1 - new-year-counter</h2></summary>
  
  <p>Использовали:</p>
  <ul>
    <li>Go</li>
    <li>Docker</li>
  </ul>
  
  <p>Инструкция</p>
  <ul>
    <li>Скачать</li>
    <li>Собрать проект командой - <b>docker build -t new-year-counter .</b></li>
    <li>Запустить проект - <b>docker run -d -p 3000:3000 --name new-year-app new-year-counter</b></li>
    <li>Запрос на получение результата - <b>curl http://localhost:3000</b></li>
  </ul>
  
</details>

<h1>Pipe-api - Social Media Backend Service</h1>

<details>
  <summary><h2>📱 Lab 2 - Pipe-api (Base API)</h2></summary>
  
  <h3>Используемые технологии:</h3>
  <ul>
    <li><b>Go 1.21+</b> - основной язык разработки</li>
    <li><b>Gin Framework</b> - HTTP веб-фреймворк</li>
    <li><b>PostgreSQL 16</b> - реляционная база данных</li>
    <li><b>Docker & Docker Compose</b> - контейнеризация и оркестрация</li>
    <li><b>lib/pq</b> - драйвер PostgreSQL для Go</li>
    <li><b>godotenv</b> - управление переменными окружения</li>
    <li><b>google/uuid</b> - генерация уникальных идентификаторов</li>
  </ul>
  
  <h3>Архитектура проекта</h3>
  <ul>
    <li><b>Models</b> - данные</li>
    <li><b>Repository</b> - доступ к данным</li>
    <li><b>Service</b> - бизнес-логики</li>
    <li><b>Handlers</b> - HTTP обработчики</li>
    <li><b>DTO</b> - объекты передачи данных</li>
    <li><b>Config</b> - управление конфигурацией через .env</li>
  </ul>
  
  <h3>REST API Endpoints (Базовая версия)</h3>
  <table>
    <tr>
      <th>Method</th>
      <th>Endpoint</th>
      <th>Description</th>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/posts</td>
      <td>Получить все посты (с пагинацией)</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/posts/:id</td>
      <td>Получить пост по ID</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/posts</td>
      <td>Создать новый пост</td>
    </tr>
    <tr>
      <td>PUT</td>
      <td>/api/v1/posts/:id</td>
      <td>Полностью обновить пост</td>
    </tr>
    <tr>
      <td>PATCH</td>
      <td>/api/v1/posts/:id</td>
      <td>Частично обновить пост</td>
    </tr>
    <tr>
      <td>DELETE</td>
      <td>/api/v1/posts/:id</td>
      <td>Мягкое удаление поста</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/posts/user/:userId</td>
      <td>Получить посты пользователя</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/health</td>
      <td>Health check сервиса</td>
    </tr>
   </table>
</details>

<details>
  <summary><h2>Lab 3 - Pipe-api (Extended with Auth & Users)</h2></summary>
  
  <h3>Новый функционал</h3>
  <ul>
    <li><b>Аутентификация</b> - JWT токены (Access + Refresh)</li>
    <li><b>Управление пользователями</b> - регистрация, профиль, обновление данных</li>
    <li><b>OAuth 2.0</b> - Вход через Яндекс и ВКонтакте</li>
    <li><b>Безопасность</b> - bcrypt для паролей, хеширование токенов</li>
    <li><b>Сессии</b> - управление несколькими сессиями, logout-all</li>
    <li><b>Soft Delete</b> - безопасное удаление пользователей</li>
  </ul>
  
  <h4>Аутентификация и пользователи</h4>
  <table>
    <tr>
      <th>Method</th>
      <th>Endpoint</th>
      <th>Description</th>
      <th>Auth Required</th>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/auth/register</td>
      <td>Регистрация нового пользователя</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/auth/login</td>
      <td>Вход в систему (устанавливает cookies)</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/auth/whoami</td>
      <td>Получить информацию о текущем пользователе</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/auth/refresh</td>
      <td>Обновить access token через refresh token</td>
      <td>❌</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/auth/logout</td>
      <td>Выход из текущей сессии</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/auth/logout-all</td>
      <td>Выход из всех сессий пользователя</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/users/:id</td>
      <td>Получить пользователя по ID</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>PUT/PATCH</td>
      <td>/api/v1/users/:id</td>
      <td>Обновить данные пользователя</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>DELETE</td>
      <td>/api/v1/users/:id</td>
      <td>Мягкое удаление аккаунта</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/users</td>
      <td>Список пользователей (с пагинацией)</td>
      <td>✅</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/users/email/:email</td>
      <td>Поиск по email</td>
      <td>✅</td>
    </tr>
  </table>
  
  <h4>OAuth 2.0 провайдеры</h4>
  <table>
    <tr>
      <th>Method</th>
      <th>Endpoint</th>
      <th>Description</th>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/auth/oauth/yandex</td>
      <td>Вход через Яндекс (редирект)</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/auth/oauth/vk</td>
      <td>Вход через ВКонтакте (редирект)</td>
    </tr>
  </table>
  
  <h4>Посты (расширенные)</h4>
  <table>
    <tr>
      <th>Method</th>
      <th>Endpoint</th>
      <th>Description</th>
      <th>Auth Required</th>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/posts</td>
      <td>Все посты с пагинацией</td>
      <td>error</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/posts/:id</td>
      <td>Пост по ID</td>
      <td>error</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/api/v1/posts</td>
      <td>Создать пост (привязан к user_id)</td>
      <td>ok</td>
    </tr>
    <tr>
      <td>PUT/PATCH</td>
      <td>/api/v1/posts/:id</td>
      <td>Обновить свой пост</td>
      <td>ok</td>
    </tr>
    <tr>
      <td>DELETE</td>
      <td>/api/v1/posts/:id</td>
      <td>Удалить свой пост</td>
      <td>ok</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/api/v1/posts/user/:userId</td>
      <td>Посты пользователя</td>
      <td>error</td>
    </tr>
  </table>
  
  <h3>Примеры запросов</h3>
  
  <h4>Регистрация</h4>
  <pre><code>curl -X POST http://localhost:4200/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123",
    "phone": "+79991234567"
  }'</code></pre>
  
  <h4>Вход (устанавливает cookies)</h4>
  <pre><code>curl -X POST http://localhost:4200/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }' \
  -c cookies.txt</code></pre>
  
  <h4>Кто я (текущий пользователь)</h4>
  <pre><code>curl http://localhost:4200/api/v1/auth/whoami \
  -b cookies.txt</code></pre>
  
  <h4>Создать пост (авторизованный)</h4>
  <pre><code>curl -X POST http://localhost:4200/api/v1/posts \
  -H "Content-Type: application/json" \
  -d '{
    "user_id": "your-uuid-here",
    "title": "Мой первый пост!",
    "description": "Создано через API",
    "status": "active"
  }' \
  -b cookies.txt</code></pre>
  
  <h4>Обновить профиль</h4>
  <pre><code>curl -X PATCH http://localhost:4200/api/v1/users/{user-id} \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "phone": "+79876543210"
  }' \
  -b cookies.txt</code></pre>
  
  <h4>Выход из всех сессий</h4>
  <pre><code>curl -X POST http://localhost:4200/api/v1/auth/logout-all \
  -b cookies.txt</code></pre>
  
  <h3>Модели данных</h3>
  
  <h4>Users Table</h4>
  <pre><code>{
  "id": "uuid",
  "email": "user@example.com",
  "phone": "+79991234567",
  "yandex_id": "optional",
  "vk_id": "optional",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "deleted_at": "timestamp (soft delete)"
}</code></pre>
  
  <h4>User Tokens Table</h4>
  <pre><code>{
  "id": "uuid",
  "user_id": "uuid",
  "token_hash": "sha256 hash",
  "token_type": "access | refresh",
  "expires_at": "timestamp",
  "revoked": "boolean"
}</code></pre>

  <h3>🔧 Environment Variables</h3>
  
  <pre><code># Database
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=rugram_user
    DB_PASSWORD=rugram_password
    DB_NAME=rugram_db

    # App
    APP_PORT=4200
    APP_ENV=development

    # Pagination
    DEFAULT_PAGE=1
    DEFAULT_LIMIT=10
    MAX_LIMIT=100

    # JWT Secrets (измените в production!)
    JWT_ACCESS_SECRET=your-super-secret-access-key-here
    JWT_REFRESH_SECRET=your-super-secret-refresh-key-here

    # OAuth Yandex
    YANDEX_CLIENT_ID=your_yandex_client_id
    YANDEX_CLIENT_SECRET=your_yandex_client_secret
    YANDEX_REDIRECT_URI=http://localhost:4200/api/v1/auth/oauth/yandex/callback

    # OAuth VK
    VK_CLIENT_ID=your_vk_client_id
    VK_CLIENT_SECRET=your_vk_client_secret
    VK_REDIRECT_URI=http://localhost:4200/api/v1/auth/oauth/vk/callback

</code></pre>

  <h3>Инструкция по запуску</h3>
  
  <p><b>1. Клонировать и настроить</b></p>
  <pre><code>git clone https://github.com/yourusername/Pipe-api.git
cd Pipe-api
cp .env.example .env
# Отредактируйте .env, добавьте JWT_SECRET</code></pre>
  
  <p><b>2. Запуск с Docker Compose</b></p>
  <pre><code># Собрать и запустить
docker-compose up -d --build

# Проверить логи

docker-compose logs -f api

# Выполнить миграции (автоматически при старте)</code></pre>

  <p><b>3. Проверка работы</b></p>
  <pre><code># Health check
curl http://localhost:4200/health

# Регистрация

curl -X POST http://localhost:4200/api/v1/auth/register \
 -H "Content-Type: application/json" \
 -d '{"email":"test@test.com","password":"test123"}'

# Логин

curl -X POST http://localhost:4200/api/v1/auth/login \
 -H "Content-Type: application/json" \
 -d '{"email":"test@test.com","password":"test123"}' \
 -c cookies.txt

# Проверка whoami

curl http://localhost:4200/api/v1/auth/whoami -b cookies.txt</code></pre>

  <h3>Решение проблем</h3>
  
  <p><b>Ошибка: "relation already exists" при миграциях</b></p>
  <ul>
    <li>Используйте <code>CREATE TABLE IF NOT EXISTS</code> и <code>CREATE INDEX IF NOT EXISTS</code></li>
    <li>Очистите volume: <code>docker-compose down -v</code></li>
    <li>Удалите таблицу <code>schema_migrations</code> если используется</li>
  </ul>
  
  <p><b>Ошибка: "invalid or expired token"</b></p>
  <ul>
    <li>Проверьте системное время на сервере</li>
    <li>Убедитесь что JWT_SECRET одинаковый при создании и проверке</li>
    <li>Access token живет 15 минут, используйте /refresh</li>
  </ul>
  
  <p><b>OAuth не работает</b></p>
  <ul>
    <li>Зарегистрируйте приложение в Яндекс.OAuth и VK API</li>
    <li>Укажите правильные Redirect URIs</li>
    <li>Проверьте переменные окружения YANDEX_CLIENT_ID и др.</li>
  </ul>

</details>

<details>
  <summary><h2>📄 Lab 4 - OpenAPI (Swagger) Documentation</h2></summary>

  <h3>Используемые технологии:</h3>
  <ul>
    <li><b>swaggo/swag</b> - генерация OpenAPI спецификации из комментариев к коду</li>
    <li><b>swaggo/gin-swagger</b> - middleware для отдачи Swagger UI</li>
    <li><b>swaggo/files</b> - встроенные статические файлы Swagger UI</li>
    <li><b>Code-First подход</b> – документация создаётся автоматически на основе аннотаций</li>
  </ul>

  <h3>Что сделано</h3>
  <ul>
    <li>Все контроллеры (auth, users, posts) задокументированы с помощью комментариев <code>// @Summary</code>, <code>// @Description</code>, <code>// @Tags</code>, <code>// @Accept</code>, <code>// @Produce</code>, <code>// @Param</code>, <code>// @Success</code>, <code>// @Failure</code>.</li>
    <li>DTO (RegisterRequest, LoginRequest, PostResponse, UserResponse и др.) аннотированы через <code>// @Schema</code> с указанием типов, обязательности и примеров значений.</li>
    <li>Настроена схема безопасности <code>BearerAuth</code> (JWT) для Swagger UI. Хотя реальное приложение использует HttpOnly cookies, в документации добавлена возможность отправлять токен в заголовке <code>Authorization: Bearer &lt;token&gt;</code> для удобного тестирования защищённых эндпоинтов.</li>
    <li>Документация <strong>автоматически генерируется</strong> командой <code>swag init</code> (интегрирована в процесс сборки). Ручное написание YAML/JSON отсутствует.</li>
    <li>Включение Swagger UI происходит только в режиме разработки (<code>APP_ENV=development</code>). При <code>APP_ENV=production</code> документация недоступна (возвращается 404).</li>
  </ul>

  <h3>Установка зависимостей</h3>
  <pre><code>go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files</code></pre>

  <h3>Примеры аннотаций</h3>

  <h4>Контроллер (Auth)</h4>
  <pre><code>// Login godoc
// @Summary      Вход пользователя
// @Description  Аутентификация по email и паролю, установка HttpOnly cookies
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Данные для входа"
// @Success      200 {object} dto.LoginResponse "Успешный вход, cookies установлены"
// @Failure      400 {object} dto.ErrorResponse "Неверный запрос"
// @Failure      401 {object} dto.ErrorResponse "Неверный email или пароль"
// @Router       /api/v1/auth/login [post]</code></pre>

  <h4>DTO (RegisterRequest)</h4>
  <pre><code>type RegisterRequest struct {
    // Email пользователя
    // @Schema(example="user@example.com", required=true)
    Email string `json:"email" binding:"required,email"`

    // Пароль (мин. 6 символов)
    // @Schema(example="strongP@ssw0rd", required=true, minLength=6)
    Password string `json:"password" binding:"required,min=6"`

    // Телефон (опционально)
    // @Schema(example="+79991234567")
    Phone string `json:"phone"`

}</code></pre>

  <h3>Конфигурация условного запуска (main.go)</h3>
  <pre><code>// В функции main() после создания роутера
if os.Getenv("APP_ENV") != "production" {
    // Настройка Swagger UI
    docs.SwaggerInfo.Title = "Pipe‑API"
    docs.SwaggerInfo.Description = "Социальная сеть — документация API"
    docs.SwaggerInfo.Version = "1.0"
    docs.SwaggerInfo.Host = "localhost:4200"
    docs.SwaggerInfo.BasePath = "/api/v1"
    docs.SwaggerInfo.Schemes = []string{"http", "https"}

    // Добавление middleware для Swagger
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}</code></pre>

  <h3>Настройка безопасности (JWT Bearer)</h3>
  <p>В Swagger UI добавлена возможность авторизации через кнопку <strong>Authorize</strong>. Токен, полученный после входа, нужно вручную вставить в поле Bearer Token. Это удобно для тестирования защищённых эндпоинтов. В реальном приложении авторизация осуществляется через HttpOnly cookies, которые браузер автоматически подставляет при запросах из Swagger UI, если документация открыта на том же домене.</p>

  <pre><code>// В программе безопасность объявляется так:
// @SecurityDefinitions BearerAuth
// @in header
// @name Authorization</code></pre>

  <h3>Результат</h3>
  <ul>
    <li>Документация доступна по адресу: <code>http://localhost:4200/swagger/index.html</code> (только при <code>APP_ENV=development</code>).</li>
    <li>Все эндпоинты (auth, users, posts) сгруппированы по тегам, имеют понятные описания, возможные коды ответов и примеры.</li>
    <li>В Swagger UI отображаются схемы всех DTO, включая поля, типы и примеры значений.</li>
    <li>Чувствительные поля (пароль, хеши токенов) скрыты из ответов с помощью <code>// @Schema(hidden=true)</code>.</li>
    <li>Защищённые маршруты помечены значком замка. При нажатии <strong>Authorize</strong> и вводе валидного JWT-токена можно выполнять любые запросы прямо из интерфейса.</li>
  </ul>

  <h3>Инструкция по проверке</h3>
  <ol>
    <li>Убедиться, что в <code>.env</code> установлено <code>APP_ENV=development</code>.</li>
    <li>Запустить проект: <code>docker-compose up --build</code>.</li>
    <li>Открыть в браузере <code>http://localhost:4200/swagger/index.html</code>.</li>
    <li>Выполнить запрос <code>POST /api/v1/auth/register</code> или <code>POST /api/v1/auth/login</code>.</li>
    <li>Скопировать полученный <code>access_token</code> (если он возвращается в теле ответа) или использовать cookies. Для упрощения документация позволяет работать через Bearer Token: нажать <strong>Authorize</strong> → ввести <code>Bearer &lt;access_token&gt;</code>.</li>
    <li>Проверить любой защищённый маршрут (например, <code>GET /api/v1/auth/whoami</code> или <code>POST /api/v1/posts</code>).</li>
    <li>Убедиться, что при <code>APP_ENV=production</code> документация недоступна (проверить, пересобрав контейнер с переменной окружения).</li>
  </ol>

  <h3>Контрольные вопросы (ответы)</h3>
  <ol>
    <li><strong>Что такое спецификация OpenAPI и чем она отличается от Swagger UI?</strong><br>
    OpenAPI — это стандарт описания REST API (в формате JSON/YAML). Swagger UI — это инструмент, который визуализирует эту спецификацию в виде интерактивной веб-страницы.</li>
    <li><strong>Code-First vs Design-First? Какой использован?</strong><br>
    Использован <strong>Code-First</strong>: спецификация генерируется из аннотаций к коду. Плюсы: документация всегда актуальна, минимум дублирования. Минусы: код немного захламляется метаданными.</li>
    <li><strong>Почему важно скрывать документацию в production?</strong><br>
    Открытая документация раскрывает структуру API, эндпоинты, схемы данных и может содержать тестовые данные. Это увеличивает риск атак (например, поиск недокументированных/уязвимых методов).</li>
    <li><strong>Как документировать HttpOnly Cookies в OpenAPI?</strong><br>
    Можно добавить схему типа <code>apiKey</code> с расположением <code>cookie</code>, либо описать <code>Bearer</code> с пояснением в описании, что реальное приложение использует cookies. В данной работе добавлен <code>BearerAuth</code> для удобства тестирования в Swagger UI.</li>
    <li><strong>Зачем нужны примеры в документации?</strong><br>
    Примеры помогают разработчикам фронтенда понять формат запроса/ответа без чтения схемы. Это снижает количество ошибок интеграции и ускоряет разработку.</li>
    <li><strong>Какие HTTP коды обязательно описывать для CRUD?</strong><br>
    <code>200 OK</code> (GET, PUT, PATCH), <code>201 Created</code> (POST), <code>204 No Content</code> (DELETE), <code>400 Bad Request</code>, <code>401 Unauthorized</code>, <code>403 Forbidden</code>, <code>404 Not Found</code>, <code>500 Internal Server Error</code>.</li>
  </ol>

  <h3>Итог</h3>
  <p>Лабораторная работа выполнена в полном объёме. Реализована автоматическая генерация OpenAPI-спецификации, документация доступна только в режиме разработки, все эндпоинты задокументированы с примерами, настроена безопасность для тестирования защищённых маршрутов. Чувствительные данные из ответов исключены. Приложение запускается через <code>docker-compose up --build</code> без дополнительных ручных действий.</p>
</details>
