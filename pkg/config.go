package pkg

type Config struct {
	TautologyCheckDistance int // Будем искать тавтологии отрезке слов такого количества
	TautologyCheckOffset int // После поиска сделаем такой отступ и поищем еще раз
}
