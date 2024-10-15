package app

import (
	"bot-oleg/intermal/config"
	"context"
	"fmt"
	"github.com/tkcrm/mx/logger"
	"gopkg.in/telebot.v4"
	"log"
	"time"
)

type Question struct {
	Category   string
	Difficulty string
	Text       string
	VideoNote  string
	Photos     []string
}

// Массив вопросов
var questions = []Question{
	{Category: "КиКи", Difficulty: "Easy", Text: "Возраст Христа: \n\nТеперь тебе 33 года. Поздравляю!\nТак как 33 это возраст в котором распяли Христа, то вспомни 5 фактов о себе, которые так же можно отнести к Иисусу.\nОпиши эти факты и отправь их в текстовом или кружочном формате в чат ПоштаМишта. 100 баллов твои!", VideoNote: "./assets/kiki_1.MP4"},
	{Category: "КиКи", Difficulty: "Medium", Text: "Исход из Египта:\n\nТеперь тебе 33 года. Поздравляю!\nТак как 33 – это возраст Христа, предлагаем тебе проложить в доме \"\"путь через Красное море\"\". Для этого положи одежду или другие предметы на пол в форме волны и представьте, что проводишь своих друзей между ними. Сними фото или видео и отправь его в чат ПоштаМишта. Ура, 300 баллов заработаны!", VideoNote: "./assets/kiki_2.MP4"},
	{Category: "КиКи", Difficulty: "Hard", Text: "Посети место для медитации:\n\nНайди в своем городе тихое место для молитвы или медитации (церковь, парк, холм) и проведи там 15-20 минут в размышлениях о последем годе своей жизни. Сделай заметки о своих ощущениях. 500 баллов у тебя в кармане!", VideoNote: "./assets/kiki_1.MP4"},
	{Category: "Стики", Difficulty: "Easy", Text: "Теперь мы поиграем с тобой в геогессер, тебе нужно отметить на карте где примерно сделаны эти фото, карту для отметок можно получить у Ани", Photos: []string{"./assets/stick_1.jpg", "./assets/stick_2.jpg", "./assets/stick_3.jpg", "./assets/stick_4.jpg", "./assets/stick_5.jpg"}},
	{Category: "Стики", Difficulty: "Medium", Text: "Геогессер второй тур, чтобы получить свои three hundred bucks тебе необходимо также отметить новые фотографии", Photos: []string{"./assets/stick_6.jpg", "./assets/stick_7.jpg", "./assets/stick_8.jpg", "./assets/stick_9.jpg", "./assets/stick_10.jpg"}},
	{Category: "Стики", Difficulty: "Hard", Text: "Геогессер финал, вообще за этот раунд считаю что должны давать больше 500 очков, но увы. Тебе необходимо отметить на карте где пердположительно ты мог видеть голового Кирилла"},
	{Category: "Тен", Difficulty: "Easy", Text: "Внимание! Внимание! Объявляется конкурс чтецов имени Ататюрка! \n\nЧтобы заработать дополнительные баллы, тебе нужно с выражением прочитать стихотворение авторства неизвестного поэта. Распечатанное стихотворение попроси у Ани. Видеозапись прочтения отправить в чат PandaBanda. Экспертные члены независимого жюри присудят количество баллов, соотвествующее выразительности прочтения"},
	{Category: "Тен", Difficulty: "Medium", Text: "Внимание! Внимание! Объявляется конкурс чтецов имени Ататюрка! \n\nЧтобы заработать дополнительные баллы, тебе нужно с выражением прочитать стихотворение авторства неизвестного поэта. Распечатанное стихотворение попроси у Ани. Видеозапись прочтения отправить в чат PandaBanda. Экспертные члены независимого жюри присудят количество баллов, соотвествующее выразительности прочтения"},
	{Category: "Тен", Difficulty: "Hard", Text: "Внимание! Внимание! Объявляется конкурс чтецов имени Ататюрка! \n\nЧтобы заработать дополнительные баллы, тебе нужно с выражением прочитать стихотворение авторства неизвестного поэта. Распечатанное стихотворение попроси у Ани. Видеозапись прочтения отправить в чат PandaBanda. Экспертные члены независимого жюри присудят количество баллов, соотвествующее выразительности прочтения"},
	{Category: "Расим", Difficulty: "Easy", Text: "Чтобы заработать 100 баллов и получить бонус от Расима, ответь правильно на следующие вопросы:\n\n1) Чем отличается азербайджанская окрошка от русской?\n\n2) Что можно увидеть на горизонте, загорая на закате на пляже Шихова под Баку?\n\nСверь правильность ответов с Аней. Бонус: Расим приготовит любое блюдо из его арсенала по твоему запросу."},
	{Category: "Расим", Difficulty: "Medium", Text: "Чтобы заработать 300 баллов и получить бонус от Расима, ответь правильно на следующие вопросы:\n\n1) Как сейчас называется Телеграфная улица в Баку?\n\n2) В каком музее под небом с наскальными рисунками первобытных людей мы пили пиво в январе 2022?\n\nСверь правильность ответов с Аней. Бонус: Расим приготовит любое блюдо из его арсенала по твоему запросу."},
	{Category: "Расим", Difficulty: "Hard", Text: "Чтобы заработать 500 баллов и получить бонус от Расима, ответь правильно на следующие вопросы:\n\n1) Чем отличается русская водка от азербайджанской?\n\n2) Что делать, если объявили штормовое предупреждение на пляже?\n\nСверь правильность ответов с Аней. Бонус: Расим приготовит любое блюдо из его арсенала по твоему запросу."},
	{Category: "Лёня Динамо", Difficulty: "Easy", Text: "Чтобы получить воображаемые 100 баллов, тебе надо отправить ответ на этот вопрос в видеокружочке Лёне:\n\nВ каком матче Сильвестр Слай Игбун забил дубль за московское Динамо?"},
	{Category: "Лёня Динамо", Difficulty: "Medium", Text: "Чтобы получить воображаемые 100 баллов, тебе надо отправить ответ на этот вопрос в видеокружочке Лёне:\n\nВ каком году Точилин забил легендарный гол Алании"},
	{Category: "Лёня Динамо", Difficulty: "Hard", Text: "", VideoNote: "./assets/leny_hard.mp4"},
	{Category: "Ксюша и Гоша", Difficulty: "Easy", Text: "Чтобы заработать 100 баллов, определи, на какой из трёх фотографий запечатлён белый гриб. Ответ отправь в чат Кубка гран-при по Руммикубу", Photos: []string{"./assets/grib1.jpg", "./assets/grib2.jpg", "./assets/grib3.jpg"}},
	{Category: "Ксюша и Гоша", Difficulty: "Medium", Text: "Чтобы заработать 300 баллов, посчитай количество фишек в игре Руммикуб. Ответ отправь в чат Кубка гран-при по Руммикубу"},
	{Category: "Ксюша и Гоша", Difficulty: "Hard", Text: "Чтобы заработать 500 баллов, нарисуй логотип для людей, которые не хотят идти в горы, но им приходится это делать. Результат отправь в чат Кубка гран-при по Руммикубу"},
	{Category: "Катя и Саша", Difficulty: "Easy", Text: "Выполни задание Кати и Саши, чтобы заработать 100 баллов!", VideoNote: "./assets/katy_sasha.mp4", Photos: []string{"./assets/katy1.jpg", "./assets/katy2.jpg", "./assets/katy3.jpg", "./assets/katy4.jpg", "./assets/katy5.jpg", "./assets/katy6.jpg", "./assets/katy7.jpg"}},
	{Category: "Катя и Саша", Difficulty: "Medium", Text: "Выполни задание Кати и Саши, чтобы заработать 300 баллов!", VideoNote: "./assets/katy_sasha.mp4", Photos: []string{"./assets/katy1.jpg", "./assets/katy2.jpg", "./assets/katy3.jpg", "./assets/katy4.jpg", "./assets/katy5.jpg", "./assets/katy6.jpg", "./assets/katy7.jpg"}},
	{Category: "Катя и Саша", Difficulty: "Hard", Text: "Выполни задание Кати и Саши, чтобы заработать 500 баллов!", VideoNote: "./assets/katy_sasha.mp4", Photos: []string{"./assets/katy1.jpg", "./assets/katy2.jpg", "./assets/katy3.jpg", "./assets/katy4.jpg", "./assets/katy5.jpg", "./assets/katy6.jpg", "./assets/katy7.jpg"}},
	{Category: "Маша и Андрей", Difficulty: "Easy", Text: "Выполни задание от Маши и Андрея и получи 100 баллов! Видеоотчёт отправь в сообщении Маше.", VideoNote: "./assets/masha1.mp4"},
	{Category: "Маша и Андрей", Difficulty: "Medium", Text: "Выполни задание от Маши и Андрея и получи 300 баллов! Видеоотчёт отправь в сообщении Маше.", VideoNote: "./assets/masha2.mp4"},
	{Category: "Маша и Андрей", Difficulty: "Hard", Text: "Выполни задание от Маши и Андрея и получи 500 баллов! Видеоотчёт отправь в сообщении Маше.", VideoNote: "./assets/masha3.mp4"},
	{Category: "Олег Егоров", Difficulty: "Easy", Text: "Чтобы заработать 100 баллов, заполни сочинение \"Мой друг Олег\" недостающими словами. В твоём распоряжении будет выбор слов для вставки. Распечатанное сочинение попроси у Ани."},
	{Category: "Олег Егоров", Difficulty: "Medium", Text: "Чтобы заработать 300 баллов, заполни сочинение \"Мой друг Олег\" недостающими словами. В твоём распоряжении будет выбор слов для вставки. Распечатанное сочинение попроси у Ани."},
	{Category: "Олег Егоров", Difficulty: "Hard", Text: "Чтобы заработать 500 баллов, заполни сочинение \"Мой друг Олег\" недостающими словами. В твоём распоряжении будет выбор слов для вставки. Распечатанное сочинение попроси у Ани."},
	{Category: "Женя и Миша", Difficulty: "Easy", Text: "Для того, чтобы заработать 100 баллов, разгадай кроссворд на закадычную тему от Жени и Миши. Распечатанный кроссворд попроси у Ани."},
	{Category: "Женя и Миша", Difficulty: "Medium", Text: "Для того, чтобы заработать 300 баллов, разгадай кроссворд на закадычную тему от Жени и Миши. Распечатанный кроссворд попроси у Ани."},
	{Category: "Женя и Миша", Difficulty: "Hard", Text: "Для того, чтобы заработать 500 баллов, разгадай кроссворд на закадычную тему от Жени и Миши. Распечатанный кроссворд попроси у Ани."},
	{Category: "Макс Мишин", Difficulty: "Easy", Text: "Ура! 100 баллов достаются тебе в подарок просто так! "},
	{Category: "Макс Мишин", Difficulty: "Medium", Text: "Ура! 300 баллов достаются тебе в подарок просто так! "},
	{Category: "Макс Мишин", Difficulty: "Hard", Text: "Ура! 500 баллов достаются тебе в подарок просто так! "},
	{Category: "Аня", Difficulty: "Easy", Text: "Рубрика Орфей: чтобы заработать 100 баллов, назови композиторов 3 мелодий, которые включит Аня. "},
	{Category: "Аня", Difficulty: "Medium", Text: "Чтобы заработать 300 баллов, обыграй Аню в нарды."},
	{Category: "Аня", Difficulty: "Hard", Text: "Чтобы заработать 500 баллов, вместе с Аней слепи из пластилина фигурку как на картинке. Картинку и пластилин возьми у Ани."},
}

var categories = []struct {
	Name   string
	Unique string
}{
	{Name: "КиКи", Unique: "category_kiki"},
	{Name: "Стики", Unique: "category_stiki"},
	{Name: "Тен", Unique: "category_ten"},
	{Name: "Расим", Unique: "category_rasim"},
	{Name: "Олег Егоров", Unique: "category_egorov"},
	{Name: "Макс Мишин", Unique: "category_mishin"},
	{Name: "Женя и Миша", Unique: "category_evg"},
	{Name: "Лёня Динамо", Unique: "category_dinamo"},
	{Name: "Маша и Андрей", Unique: "category_masha"},
	{Name: "Катя и Саша", Unique: "category_shodina"},
	{Name: "Ксюша и Гоша", Unique: "category_gosha"},
	{Name: "Аня", Unique: "category_anna"},
}

// Структура для отслеживания прогресса пользователей
type UserProgress struct {
	UserID    int64
	Completed map[string]bool // Используем уникальный идентификатор задания как ключ
}

var userProgressMap = make(map[int64]*UserProgress)

func Run(ctx context.Context, conf *config.Config, l logger.Logger) {
	l.Info("starting app")

	// Настройка Telegram бота
	pref := telebot.Settings{
		Token:  conf.Telegram.Token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(telebot.OnText, func(c telebot.Context) error {
		logIncomingMessage(c, l)
		return nil
	})

	b.Handle(&telebot.InlineButton{Unique: "all_categories"}, func(c telebot.Context) error {
		// Создаем кнопки категорий, как при старте
		var buttons [][]telebot.InlineButton
		for i := 0; i < len(categories); i += 3 {
			var row []telebot.InlineButton
			for j := i; j < i+3 && j < len(categories); j++ {
				row = append(row, telebot.InlineButton{
					Text:   categories[j].Name,
					Unique: categories[j].Unique,
				})
			}
			buttons = append(buttons, row)
		}

		return c.Send("Выбери категорию:", &telebot.ReplyMarkup{
			InlineKeyboard: buttons,
		})
	})

	// Обработчик команды /start
	b.Handle("/start", func(c telebot.Context) error {
		// Создаем кнопки категорий
		var buttons [][]telebot.InlineButton
		for i := 0; i < len(categories); i += 3 {
			var row []telebot.InlineButton
			for j := i; j < i+3 && j < len(categories); j++ {
				row = append(row, telebot.InlineButton{
					Text:   categories[j].Name,
					Unique: categories[j].Unique,
				})
			}
			buttons = append(buttons, row)
		}

		return c.Send("*Привет, Олег*!\n\nС днём рождения! \n\nВ этот праздничный день мы приготовили для тебя подарочки, но чтобы их получить, предлагаем сыграть в игру! С каждой заработанной 1000 баллов тебе полагается подарок 😉\n\nЗадания ты можешь выбирать в любом порядке и в любое время дня. Подсчёт баллов будет вести искуственный интеллект Аня, поэтому держи её в курсе выполнения заданий. По всем другим вопросам тебе тоже поможет Аня 🤓\n\nУ тебя есть 3 воображаемые золотые монеты, которыми ты можешь откупиться от задания, если оно тебе не понравилось. Потратив золотую монету, ты откажешься от задания, но должен будешь выбрать любой другой уровень сложности в этой же категории.\n\nУдачи! В конце челленджа предлагаем тебе созвониться, чтобы ты рассказал нам, как сильно тебе понравилось! 💙", &telebot.ReplyMarkup{
			InlineKeyboard: buttons,
		}, telebot.ModeMarkdown)
	})

	// Обработчик для категорий
	for _, category := range categories {
		cat := category
		b.Handle(&telebot.InlineButton{Unique: cat.Unique}, func(c telebot.Context) error {
			return chooseDifficulty(c, cat.Unique, cat.Name)
		})
	}

	// Обработчики сложностей
	for _, category := range categories {
		cat := category
		for _, level := range []string{"Easy", "Medium", "Hard"} {
			lev := level
			difficultyUnique := "difficulty_" + lev + "_" + cat.Unique
			b.Handle(&telebot.InlineButton{Unique: difficultyUnique}, func(c telebot.Context) error {
				return sendQuestion(c, cat.Name, lev)
			})
		}
	}

	// Обработчик для завершения задания
	b.Handle(telebot.OnCallback, func(c telebot.Context) error {
		return c.Send("Какой большой молодец! Задание выполнено. Продолжай!")
	})

	b.Start()
}

// Выбор сложности
func chooseDifficulty(c telebot.Context, category, name string) error {
	difficultyMessage := "Ты выбрал категорию *" + name + "*. Теперь выбери уровень сложности:"
	return c.Send(difficultyMessage, generateDifficultyMarkup(category), telebot.ModeMarkdown)
}

// Генерация кнопок сложностей
func generateDifficultyMarkup(category string) *telebot.ReplyMarkup {
	return &telebot.ReplyMarkup{
		InlineKeyboard: [][]telebot.InlineButton{
			{
				{Text: "100", Unique: "difficulty_Easy_" + category},
				{Text: "300", Unique: "difficulty_Medium_" + category},
				{Text: "500", Unique: "difficulty_Hard_" + category},
			},
		},
	}
}

// Отправка вопроса
func sendQuestion(c telebot.Context, category, difficulty string) error {
	for _, q := range questions {
		if q.Category == category && q.Difficulty == difficulty {
			// Генерация уникального идентификатора вопроса
			questionID := fmt.Sprintf("%s_%s", category, difficulty)
			// Проверяем прогресс пользователя
			userID := c.Sender().ID
			if _, exists := userProgressMap[userID]; !exists {
				userProgressMap[userID] = &UserProgress{
					UserID:    userID,
					Completed: make(map[string]bool), // Инициализируем как пустую карту
				}
			}

			// Отправляем вопрос с кнопкой "Задание выполнено"
			replyMarkup := telebot.ReplyMarkup{
				InlineKeyboard: [][]telebot.InlineButton{
					{
						{Text: "Задание выполнено", Unique: questionID}, // Уникальный идентификатор
						{Text: "К другим заданиям", Unique: "all_categories"},
					},
				},
			}

			if q.VideoNote != "" {
				videoNote := &telebot.VideoNote{File: telebot.FromDisk(q.VideoNote)}
				c.Send(videoNote, &replyMarkup)
			}

			if len(q.Photos) > 0 {
				album := telebot.Album{}
				for _, photo := range q.Photos {
					album = append(album, &telebot.Photo{File: telebot.FromDisk(photo)})
				}
				c.SendAlbum(album, &replyMarkup)
			}
			if q.Text != "" {
				return c.Send("*Задание:* "+q.Text, telebot.ModeMarkdown, &replyMarkup)
			}
		}
	}
	return nil
}

func logIncomingMessage(c telebot.Context, l logger.Logger) {
	user := c.Sender()
	message := c.Text()

	// Логируем ID пользователя, имя и текст сообщения
	l.Info("Получено сообщение от пользователя:", user.ID, "-", user.FirstName, user.LastName, ":", message)
}
