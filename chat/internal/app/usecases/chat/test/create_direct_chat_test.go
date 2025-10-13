package chat_test

import (
	"context"
	"errors"
	"testing"

	"chat/internal/app/models"
	"chat/internal/app/models/types"
	"chat/internal/app/usecases/chat"
	"chat/internal/app/usecases/chat/mocks"

	"github.com/stretchr/testify/suite"
)

type ChatServiceTestSuite struct {
	suite.Suite

	ChatRepository *mocks.ChatRepository
	UserIDProvider *mocks.UserIDProvider

	Usecase chat.Usecase
}

func (s *ChatServiceTestSuite) SetupTest() {
	s.ChatRepository = mocks.NewChatRepository(s.T())
	s.UserIDProvider = mocks.NewUserIDProvider(s.T())

	s.Usecase = chat.NewUsecase(chat.Deps{
		ChatRepository: s.ChatRepository,
		UserIDProvider: s.UserIDProvider,
	})
}

func (s *ChatServiceTestSuite) TearDownTest() {
	s.ChatRepository.AssertExpectations(s.T())
	s.UserIDProvider.AssertExpectations(s.T())
}

func TestChatServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ChatServiceTestSuite))
}

func (s *ChatServiceTestSuite) Test_CreateDirectChat_Positive() {
	ctx := context.Background()
	participantID := types.UserID("user2")

	expectedChat := &models.Chat{
		ID: "chat1",
		Members: []types.UserID{
			"user1", "user2",
		},
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		CreateDirectChat(ctx, types.UserID("user1"), participantID). // <- исправлено
		Return(expectedChat, nil).
		Once()

	got, err := s.Usecase.CreateDirectChat(ctx, participantID)

	s.NoError(err)
	s.NotNil(got)
	got.ID = ""
	expectedChat.ID = ""
	s.Equal(expectedChat, got)
}

func (s *ChatServiceTestSuite) Test_CreateDirectChat_Unauthenticated() {
	ctx := context.Background()
	participantID := types.UserID("user2")

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("", errors.New("no user")).
		Once()

	got, err := s.Usecase.CreateDirectChat(ctx, participantID)

	s.Nil(got)
	s.ErrorIs(err, models.ErrUnauthenticated)
}

func (s *ChatServiceTestSuite) Test_CreateDirectChat_SelfChat() {
	ctx := context.Background()
	participantID := types.UserID("user1") // тот же пользователь

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	got, err := s.Usecase.CreateDirectChat(ctx, participantID)

	s.Nil(got)
	s.ErrorIs(err, models.ErrInvalidArgument)
}
