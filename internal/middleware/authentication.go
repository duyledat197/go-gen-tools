package middleware

// import (
// 	"context"
// 	"net/http"

// 	"github.com/duyledat197/interview-hao/internal/repositories"
// 	"github.com/duyledat197/interview-hao/utils"
// 	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // Middleware ...
// type Middleware func(next http.Handler) http.Handler

// // Authentication ...
// func Authentication(userRepo repositories.UserRepository) grpc_auth.AuthFunc {
// 	return func(ctx context.Context) (context.Context, error) {
// 		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
// 		if err != nil {
// 			return ctx, err
// 		}
// 		userID, err := utils.VerifyToken(token)
// 		if err != nil {
// 			return ctx, err
// 		}
// 		id, err := primitive.ObjectIDFromHex(userID)
// 		if err != nil {
// 			return ctx, err
// 		}
// 		user, err := userRepo.FindByID(ctx, id)
// 		if err != nil {
// 			return ctx, err
// 		}
// 		// grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))
// 		ctx = context.WithValue(ctx, "user", user)
// 		return ctx, nil
// 	}
// }
