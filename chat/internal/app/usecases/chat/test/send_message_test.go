package chat_test

import (
	"context"
	"errors"
	"strings"
	"testing"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	"chat/internal/app/usecases/chat"
	"chat/internal/app/usecases/chat/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SendMessageTestSuite struct {
	suite.Suite

	ChatRepository *mocks.ChatRepository
	UserIDProvider *mocks.UserIDProvider

	Usecase chat.Usecase
}

func (s *SendMessageTestSuite) SetupTest() {
	s.ChatRepository = mocks.NewChatRepository(s.T())
	s.UserIDProvider = mocks.NewUserIDProvider(s.T())

	s.Usecase = chat.NewUsecase(chat.Deps{
		ChatRepository: s.ChatRepository,
		UserIDProvider: s.UserIDProvider,
	})
}

func (s *SendMessageTestSuite) TearDownTest() {
	s.ChatRepository.AssertExpectations(s.T())
	s.UserIDProvider.AssertExpectations(s.T())
}

func TestSendMessageTestSuite(t *testing.T) {
	suite.Run(t, new(SendMessageTestSuite))
}

func (s *SendMessageTestSuite) Test_SendMessage_Positive() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	text := "Hello, world!"

	chatFromRepo := &models.Chat{
		ID:      chatID,
		Members: []types.UserID{"user1", "user2"},
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chatFromRepo, nil).
		Once()

	s.ChatRepository.EXPECT().
		CreateMessage(ctx, mock.MatchedBy(func(msg *models.Message) bool {
			return msg.ChatID == chatID && msg.SenderID == "user1" && msg.Text == text
		})).
		Return(&models.Message{ID: "msg1"}, nil).
		Once()

	got, err := s.Usecase.SendMessage(ctx, chatID, text)

	s.NoError(err)
	s.NotNil(got)
	s.Equal(chatID, got.ChatID)
	s.Equal("user1", string(got.SenderID))
	s.Equal(text, got.Text)
}

func (s *SendMessageTestSuite) Test_SendMessage_Unauthenticated() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	text := "Hello"

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("", errors.New("no user")).
		Once()

	got, err := s.Usecase.SendMessage(ctx, chatID, text)

	s.Nil(got)
	s.ErrorIs(err, models.ErrUnauthenticated)
}

func (s *SendMessageTestSuite) Test_SendMessage_TooLong() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	text := strings.Repeat("a", 5000)

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	got, err := s.Usecase.SendMessage(ctx, chatID, text)

	s.Nil(got)
	s.ErrorIs(err, models.ErrInvalidArgument)
}

func (s *SendMessageTestSuite) Test_SendMessage_PermissionDenied() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")
	text := "Hello"

	chatFromRepo := &models.Chat{
		ID:      chatID,
		Members: []types.UserID{"user2", "user3"},
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetChat(ctx, chatID).
		Return(chatFromRepo, nil).
		Once()

	got, err := s.Usecase.SendMessage(ctx, chatID, text)

	s.Nil(got)
	s.ErrorIs(err, models.ErrPermissionDenied)
}
