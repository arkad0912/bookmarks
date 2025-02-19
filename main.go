package main

import "fmt"

/* Создать приложение, которое сначала выдает меню:
-1. Посмотреть закладки
-2. Добавить закладку
-3. Удалить закладку
-4. Выход

При 1 - Выводит закладки
При 2 - 2 поля ввода названия и адреса и после добавление
При 3 - Ввод названия и удаление по нему
При 4 - Завершение
*/

type bookmarkMap = map[string]string //Делаем новый тип, чтобы постоянно не объявлять мапу (Alias)

func main() {

	bookmarks := bookmarkMap{}

	fmt.Println("Приложение для закладок")

Menu: //label, можно отсылаться
	for {

		variant := menu()

		switch variant {
		case 1:
			printBookmark(bookmarks)
		case 2:
			addBookmark(bookmarks)
		case 3:
			deleteBookmark(bookmarks)
		case 4:
			break Menu
		}

	}

}

func menu() int {

	var variant int
	fmt.Println("Меню:")
	fmt.Println("1. Посмотреть закладки:")
	fmt.Println("2. Добавить закладку:")
	fmt.Println("3. Удалить закладку:")
	fmt.Println("4. Выход:")
	fmt.Scan(&variant)
	return variant

}

func printBookmark(bookmarks bookmarkMap) {

	if len(bookmarks) == 0 {
		fmt.Println("Нет закладок")
	}

	for key, value := range bookmarks {
		fmt.Println(key, ":", value)
	}
}

func addBookmark(bookmarks bookmarkMap) {

	var newBookmarkKey string
	var newBookmarkValue string

	fmt.Println("Введите название:")
	fmt.Scan(&newBookmarkKey)

	fmt.Println("Введите ссылку:")
	fmt.Scan(&newBookmarkValue)

	bookmarks[newBookmarkKey] = newBookmarkValue
}
func deleteBookmark(bookmarks bookmarkMap) {
	var BookmarkKeyToDelete string

	fmt.Print("Введите название:")
	fmt.Scan(&BookmarkKeyToDelete)
	delete(bookmarks, BookmarkKeyToDelete)
}
