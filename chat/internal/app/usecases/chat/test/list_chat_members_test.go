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

type ListChatMembersTestSuite struct {
	suite.Suite

	ChatRepository *mocks.ChatRepository
	UserIDProvider *mocks.UserIDProvider

	Usecase chat.Usecase
}

func (s *ListChatMembersTestSuite) SetupTest() {
	s.ChatRepository = mocks.NewChatRepository(s.T())
	s.UserIDProvider = mocks.NewUserIDProvider(s.T())

	s.Usecase = chat.NewUsecase(chat.Deps{
		ChatRepository: s.ChatRepository,
		UserIDProvider: s.UserIDProvider,
	})
}

func (s *ListChatMembersTestSuite) TearDownTest() {
	s.ChatRepository.AssertExpectations(s.T())
	s.UserIDProvider.AssertExpectations(s.T())
}

func TestListChatMembersTestSuite(t *testing.T) {
	suite.Run(t, new(ListChatMembersTestSuite))
}

func (s *ListChatMembersTestSuite) Test_ListChatMembers_Positive() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")

	chatFromRepo := &models.Chat{
		ID:      "chat1",
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

	got, err := s.Usecase.ListChatMembers(ctx, chatID)

	s.NoError(err)
	s.Equal(chatFromRepo.Members, got)
}

func (s *ListChatMembersTestSuite) Test_ListChatMembers_Unauthenticated() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("", errors.New("no user")).
		Once()

	got, err := s.Usecase.ListChatMembers(ctx, chatID)

	s.Nil(got)
	s.ErrorIs(err, models.ErrUnauthenticated)
}

func (s *ListChatMembersTestSuite) Test_ListChatMembers_PermissionDenied() {
	ctx := context.Background()
	chatID := types.ChatID("chat1")

	chatFromRepo := &models.Chat{
		ID:      "chat1",
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

	got, err := s.Usecase.ListChatMembers(ctx, chatID)

	s.Nil(got)
	s.ErrorIs(err, models.ErrPermissionDenied)
}
