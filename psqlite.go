package psqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "admin"
	Name     = "postgres"
)

//Объект базы данных
var DB *sql.DB

//Название и записи таблицы
var defaultTables = map[string]string{
	"users":    `"id" SERIAL PRIMARY KEY, "login" varchar(64), "sha256" varchar(64), "other" varchar(64)`,
	"sessions": `"id" SERIAL PRIMARY KEY, "sessions" varchar(64), "sha256" varchar(64)`,
}

var TableSQL = map[string]string{}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

//Устанавливаем настройки драйвера ДБ
func SettingDB(host string, port int, user, password, name string) {
	Host = host
	Port = port
	User = user
	Password = password
	Name = name
}

//Подключение БД
func OpenDB() {
	var err error //ошибка выполнения запроов к БД
	var dbInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Name) //Информация о БД
	DB, err = sql.Open("postgres", dbInfo) //Подключаем БД
	chk(err)                               //Возвращаем результат подключения БД
}

//Отключение БД
func CloseDB() {
	defer DB.Close()
}

//Создание таблицы
func CreateTableByName(tableName string) {
	var err error //ошибка выполнения запроов к БД
	for k, v := range TableSQL {
		if k == tableName {
			_, err = DB.Exec("CREATE TABLE IF NOT EXISTS " + k + "(" + v + ")")
			chk(err)
			return
		}
	}
	//Находим описания таблицы по имени в списке по умолчанию
	for k, v := range defaultTables {
		if k == tableName {
			_, err = DB.Exec("CREATE TABLE IF NOT EXISTS " + k + "(" + v + ")")
			chk(err)
			return
		}
	}
}

//Удаление таблицы
func DeleteTableByName(name string) {
	_, err := DB.Exec("DROP TABLE " + name)
	chk(err)
}
