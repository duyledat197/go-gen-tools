package metadata

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/duyledat197/go-gen-tools/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/metadata"
)

const MDUserIDKey = "user_id"

func Authentication(ctx context.Context, request *http.Request) metadata.MD {
	bearerToken := request.Header.Get("Authorization")
	var BEARER string = "Bearer "
	if !strings.HasPrefix(bearerToken, BEARER) {
		return metadata.MD{}
	}
	token := bearerToken[len(BEARER):]
	userID, err := utils.VerifyToken(token)
	if err != nil {
		return metadata.MD{}
	}
	log.Println("userID", userID)
	if _, err := primitive.ObjectIDFromHex(userID); err != nil {
		return metadata.MD{}
	}
	md := metadata.Pairs(MDUserIDKey, userID)
	return md
}

func GetUserID(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}
	values := md.Get(MDUserIDKey)
	if len(values) < 1 {
		return "", false
	}
	return values[0], true
}
