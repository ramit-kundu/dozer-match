package service_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/kundu-ramit/dozer_match/domain/mock"
	"github.com/kundu-ramit/dozer_match/service"
	"gopkg.in/go-playground/assert.v1"
)

func TestCheckExistingScrapeSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCache := mocks.NewMockCache(ctrl)
	dozerService := service.MakeNewDozerService(nil, mockCache)
	mockCache.EXPECT().Get("catdotcom").Return("", nil)
	err := dozerService.CheckExistingScrape(context.Background())

	assert.Equal(t, err, nil)
}

func TestCheckExistingScrapeFailure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCache := mocks.NewMockCache(ctrl)
	dozerService := service.MakeNewDozerService(nil, mockCache)
	mockCache.EXPECT().Get("catdotcom").Return("in_progress", nil)
	err := dozerService.CheckExistingScrape(context.Background())
	assert.Equal(t, err.Error(), "scrape is in progress")
}
