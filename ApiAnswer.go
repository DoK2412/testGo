package main

var answers = map[string]string{
	"0": "Запрос успешно выполнен.",

	"10": "Введеный пароль не соответствует требованиям.",
	"11": "Не верно введен повтор пароля.",
	"12": "Введенный номер телефона не сообветствует требованиям.",
	"13": "Пользователь с таким именем уже есть в базе.",
	"14": "Проверочный код не вереню",
	"15": "Неверный логин или пароль.",

	"20": "Пользователь не авторизован",

	"90": "Ошибка на стороне сервера.",
	"91": "Ошибка на строне базы данных.",
}

func APIAnswer(code string) map[string]string {
	var dictAnswer = map[string]string{
		"code":        code,
		"description": answers[code],
	}
	return dictAnswer
}

func APIAnswerData(code string, data string) map[string]string {
	var dictAnswer = map[string]string{
		"code":        code,
		"description": answers[code],
		"data":        data,
	}
	return dictAnswer
}
