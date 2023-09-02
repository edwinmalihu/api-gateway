package middleware

import (
	"auth-services/model"
	"auth-services/repository"
	"github.com/gin-gonic/gin"
	"sync"
)

// Authorize determines if current user has been authorized to take an action on an object.
func AddUserPlnDetail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.GetString("username")
		if username != "" {
			//fetch user pln detail
			userRepo := repository.NewUserRepository()
			detail, _, err := userRepo.DetailUserPln(model.RequestGetDetail{Username: username})
			if err == nil && detail.CorporateId != "" {
				var mu sync.Mutex
				mu.Lock()
				ctx.Set("user-pln", &model.DetailUserPln{
					Username:    username,
					UserId:      detail.UserId,
					CorporateId: detail.CorporateId,
					ApiKey:      detail.ApiKey,
				})
				mu.Unlock()
			}
		}

		ctx.Next()
	}
}
