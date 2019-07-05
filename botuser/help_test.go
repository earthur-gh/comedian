package botuser

import (
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"github.com/maddevsio/comedian/model"
	"golang.org/x/text/language"
)

func TestDisplayHelpText(t *testing.T) {

	testCases := []struct {
		language      string
		command       string
		outputMessage string
	}{
		{"en_US", "", ""},
		{"en_US", "add", "To add members use `add` command. Here is an example: `add @user @user1 / pm` You can add members with _pm, developer, designer, tester_ roles, default is a developer role, if the role is not selected. "},
		{"ru_RU", "add", "Для добавления нового пользователя используйте команду `add`. Вот пример: `add @user @user1 / designer` Вы можете добавлять участников с ролями: _pm, developer, designer, tester_. Если роль не указана, то по умолчанию будет developer."},
		{"en_US", "show", "To view members use `show` command"},
		{"ru_RU", "show", "Чтобы посмотреть список пользователей используйте команду `show`. "},
		{"en_US", "remove", "To remove members use `remove` command and tag members to remove"},
		{"ru_RU", "remove", "Чтобы удалить пользователя используйте команду `remove`. "},
		{"en_US", "add_deadline", "To set standup deadline use `add_deadline` command."},
		{"ru_RU", "add_deadline", "Чтобы установить время сдачи стэндапов в канале используйте команду `add_deadline`."},
		{"en_US", "show_deadline", "To view standup deadline in the channel use `show_deadline` command. "},
		{"ru_RU", "show_deadline", "Чтобы узнать установленное время сдачи стендапов в канале используйте команду `show_deadline`."},
		{"en_US", "remove_deadline", "To remove standup deadline in the channel use `remove_deadline` command. "},
		{"ru_RU", "remove_deadline", "Чтобы удалить время сдачи стендапов в канале используйте команду `remove_deadline`."},
	}

	bundle := &i18n.Bundle{DefaultLanguage: language.English}
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	_, err := bundle.LoadMessageFile("../active.en.toml")
	assert.NoError(t, err)
	_, err = bundle.LoadMessageFile("../active.ru.toml")
	assert.NoError(t, err)

	for _, tt := range testCases {
		properties := model.BotSettings{
			Language: tt.language,
		}

		bot := &Bot{
			bundle:     bundle,
			properties: properties,
		}
		text := bot.DisplayHelpText(tt.command)
		assert.Equal(t, tt.outputMessage, text)
	}

}
