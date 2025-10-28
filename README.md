# Практическое занятие №6
## Шишков Алексей Дмитриевич ЭФМО-02-21
## Тема
Использование ORM (GORM). Модели, миграции и связи между таблицами.
## Цели
Понять, что такое ORM и чем удобен GORM.
Научиться описывать модели Go-структурами и автоматически создавать таблицы (миграции через AutoMigrate).
Освоить базовые связи: 1:N и M:N + выборки с Preload.
Написать короткий REST (2–3 ручки) для проверки результата.
## Структура проекта 

<img width="356" height="563" alt="изображение" src="https://github.com/user-attachments/assets/940aa8bc-13fe-4b5d-a707-d1bf16899ad9" />

### 1.  Установка и настройка БД
На удалённом сервере PostgreSQL:
```sql
CREATE DATABASE pz6_gorm;
CREATE USER myuser WITH PASSWORD 'MyStrongPassword';
GRANT ALL PRIVILEGES ON DATABASE pz6_gorm TO myuser;
ALTER DATABASE pz6_gorm OWNER TO myuser;
GRANT ALL PRIVILEGES ON SCHEMA public TO myuser;
```

2.  Переменная окружения
Windows PowerShell:
```
setx DB_DSN "host=<IP_или_домен> user=myuser password=MyStrongPassword dbname=pz6_gorm port=5432 sslmode=disable"
$env:DB_DSN="host=<IP_или_домен> user=myuser password=MyStrongPassword dbname=pz6_gorm port=5432 sslmode=disable"
```
3. Запуск проекта
   Внутри проекта:
   ```
   go mod tidy
   go run ./cmd/server
   ```
4. Проверка
   ```
   curl http://localhost:8080/health
   ```
   
   <img width="484" height="474" alt="изображение" src="https://github.com/user-attachments/assets/15fe1d53-d850-40b9-9ee8-54b70bec5b25" />

ORM (Object-Relational Mapping) — это технология, которая обеспечивает удобное сопоставление объектов программного кода с таблицами реляционной базы данных. Применение ORM позволяет разработчику работать с данными на уровне структур и методов, не прибегая к ручному написанию SQL-запросов. Использование GORM упростило процесс моделирования сущностей, создания и миграции таблиц, а также реализации связей между ними (1:N и M:N). GORM позволил ускорить разработку, повысить читаемость кода и обеспечить более надёжную работу с данными за счёт встроенной защиты и автоматической генерации запросов. Таким образом, ORM выступает инструментом, повышающим продуктивность и качество разработки серверной логики.

### Результаты работы
1. Создание пользователя
   
   <img width="659" height="567" alt="изображение" src="https://github.com/user-attachments/assets/673527d6-535d-4350-a98a-34accaa28cf5" /> 

2. Создание заметки

   <img width="776" height="765" alt="изображение" src="https://github.com/user-attachments/assets/71f9fe6b-c487-47ec-9fd9-1d70db8ee5a7" />

3. Получение заметки

   <img width="778" height="772" alt="изображение" src="https://github.com/user-attachments/assets/a32fb5d9-0a53-4dd9-a66f-db1f4fe443cc" />

4. Ответы сервера

<img width="1705" height="569" alt="изображение" src="https://github.com/user-attachments/assets/040fd554-8fcf-4b87-85f7-e8f1b0ef6b98" /> 

### Фрагмент схемы БД 

<img width="551" height="343" alt="изображение" src="https://github.com/user-attachments/assets/2ca8a0c6-88a6-4b85-b3fc-33f9ccf05cc8" /> 

### Проблемы и их решения 

| Проблема                             | Решение                                                  |
| ------------------------------------ | -------------------------------------------------------- |
| PowerShell не принимал JSON с `curl` | Использовал Postman                                      |
| Нет доступа к БД                     | Настроены права пользователя `myuser` и открыт порт 5432 |
| Ошибка unique constraint             | Проверял уникальные поля Email и Name тегов              |
