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

type ListUserChatsTestSuite struct {
	suite.Suite

	ChatRepository *mocks.ChatRepository
	UserIDProvider *mocks.UserIDProvider

	Usecase chat.Usecase
}

func (s *ListUserChatsTestSuite) SetupTest() {
	s.ChatRepository = mocks.NewChatRepository(s.T())
	s.UserIDProvider = mocks.NewUserIDProvider(s.T())

	s.Usecase = chat.NewUsecase(chat.Deps{
		ChatRepository: s.ChatRepository,
		UserIDProvider: s.UserIDProvider,
	})
}

func (s *ListUserChatsTestSuite) TearDownTest() {
	s.ChatRepository.AssertExpectations(s.T())
	s.UserIDProvider.AssertExpectations(s.T())
}

func TestListUserChatsTestSuite(t *testing.T) {
	suite.Run(t, new(ListUserChatsTestSuite))
}

func (s *ListUserChatsTestSuite) Test_ListUserChats_Positive() {
	ctx := context.Background()
	userID := types.UserID("user1")

	expectedChats := []*models.Chat{
		{ID: "chat1", Members: []types.UserID{"user1", "user2"}},
		{ID: "chat2", Members: []types.UserID{"user1", "user3"}},
	}

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user1", nil).
		Once()

	s.ChatRepository.EXPECT().
		GetUserChats(ctx, userID).
		Return(expectedChats, nil).
		Once()

	got, err := s.Usecase.ListUserChats(ctx, userID)

	s.NoError(err)
	s.Equal(expectedChats, got)
}

func (s *ListUserChatsTestSuite) Test_ListUserChats_Unauthenticated() {
	ctx := context.Background()
	userID := types.UserID("user1")

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("", errors.New("no user")).
		Once()

	got, err := s.Usecase.ListUserChats(ctx, userID)

	s.Nil(got)
	s.ErrorIs(err, models.ErrUnauthenticated)
}

func (s *ListUserChatsTestSuite) Test_ListUserChats_PermissionDenied() {
	ctx := context.Background()
	userID := types.UserID("user1")

	s.UserIDProvider.EXPECT().
		GetUserIDFromIncomingContext(ctx).
		Return("user2", nil).
		Once()

	got, err := s.Usecase.ListUserChats(ctx, userID)

	s.Nil(got)
	s.ErrorIs(err, models.ErrPermissionDenied)
}
