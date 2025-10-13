package chat

// go install github.com/vektra/mockery/v2@latest

//go:generate mockery --disable-version-string --with-expecter --name UserIDProvider --filename user_id_provider_system_mock.go
//go:generate mockery --disable-version-string --with-expecter --name ChatRepository --filename chat_repository_mock.go
