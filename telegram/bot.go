package telegram

import (
	"log"
	"log/slog"
	"os"
	"rms-api/services/auth"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	button := tgbotapi.NewKeyboardButton("Отправить контакт")
	button.RequestContact = true

	requestButton := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			button,
		),
	)

	slog.Info("Bot is ready!")
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil { // If we got a message
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я главный по еде. Для того, чтобы продолжить авторизацию, мне необходимо получить Ваш контакт. Не переживайте - мы не сохраняем чувствительную информацию и никому не передаём Ваши данные. Нажмите на кнопку \"Отправить контакт\"")

			if update.Message.Command() == "start" {
				msg.ReplyMarkup = requestButton
				bot.Send(msg)
				continue
			}

			if update.Message.Contact != nil {
				phoneNumber := update.Message.Contact.PhoneNumber

				if err := auth.CheckAuthRequest(phoneNumber); err != nil {
					msg.Text = "Ошибка - вы не заполняли форму авторизации на сайте. Вернитесь на сайт, нажмите на кнопку \"Войти\" и следуйте инструкции"
					bot.Send(msg)
					continue
				}

				if err := auth.HandleSecondAuthorizationStep(phoneNumber); err != nil {
					msg.Text = "Возникла ошибка при авторизации. Попробуйте позже"
					bot.Send(msg)
					continue
				}

				msg.Text = "Номер успешно подтверждён!"
				bot.Send(msg)
				continue
			}
		}
	}
}
