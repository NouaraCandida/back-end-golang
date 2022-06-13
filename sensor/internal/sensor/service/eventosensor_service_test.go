package service_test

import (
	"testing"


)

func TestGet(t *testing.T) {
    // Setup
    gin.SetMode(gin.TestMode)

  // add two cases here
  t.Run("Success", func(t *testing.T) {
	uid, _ := uuid.NewRandom()
  
	mockUserResp := &model.User{
	  UID:   uid,
	  Email: "bob@bob.com",
	  Name:  "Bobby Bobson",
	}
  
	mockUserRepository := new(mocks.MockUserRepository)
	us := NewUserService(&USConfig{
	  UserRepository: mockUserRepository,
	})
	mockUserRepository.On("FindByID", mock.Anything, uid).Return(mockUserResp, nil)
  
	ctx := context.TODO()
	u, err := us.Get(ctx, uid)
  
	assert.NoError(t, err)
	assert.Equal(t, u, mockUserResp)
	mockUserRepository.AssertExpectations(t)
  })
}