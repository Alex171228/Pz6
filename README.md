# Практическое занятие №5  
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
   <img width="484" height="474" alt="изображение" src="https://github.com/user-attachments/assets/15fe1d53-d850-40b9-9ee8-54b70bec5b25" />


<img width="1872" height="670" alt="изображение" src="https://github.com/user-attachments/assets/9034897e-b544-430a-b959-a33fed0b0d7e" /> 

<img width="1705" height="569" alt="изображение" src="https://github.com/user-attachments/assets/040fd554-8fcf-4b87-85f7-e8f1b0ef6b98" /> 
