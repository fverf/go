package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5" //импорт библиотеки
)

func mainlib() {
	// подключаемся к бд через URL, берется из переменной окружения DATABASE_URL
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		// если не получилось подключиться, выводится ошибка и выход
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	// после завершения работы соединение закрывается
	defer conn.Close(context.Background())

	var name string  // переменная для имени (получено из бд)
	var weight int64 // переменная для веса

	// тут выполняется sql запрос для получения имя и веса для записи с id = 42 в таблице widgets
	err = conn.QueryRow(context.Background(), "select name, weight from widgets where id=$1", 42).Scan(&name, &weight)
	if err != nil {
		// если запрос не получился, ошибка и завершение выполнение программы
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	/// выводятся результаты запроса (имя, вес)
	fmt.Println(name, weight)
}
