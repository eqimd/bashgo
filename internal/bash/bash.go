package bash

/*
 * Интерфейс Bash, который умеет запускать команды
 */
type Bash interface {
	/*
	 * Метод, запускающий команду. Возвращает вывод, код возврата, и возможную ошибку
	 */
	Execute(command string) (string, int, error)
}
