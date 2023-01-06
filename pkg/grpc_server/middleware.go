package grpc_server

// AuthenticatorWithBearerToken ...
// func AuthenticatorWithBearerToken(authenticator token.Authenticator) grpc_auth.AuthFunc {
// 	return func(ctx context.Context) (context.Context, error) {

// 		token, err := grpc_auth.AuthFromMD(ctx, requestinfo.Bearer)
// 		if err != nil {
// 			return nil, common_error.ErrInvalidToken
// 		}
// 		payload, err := authenticator.Verify(token)
// 		if err != nil {
// 			return nil, err
// 		}
// 		newCtx := context.WithValue(ctx, requestinfo.Info{}, &payload.Info)
// 		return newCtx, nil
// 	}
// }

// // MappingRequestInfo ...
// func MappingRequestInfo(ctx context.Context) (context.Context, error) {
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, common_error.ErrCanNotMappingMetadata
// 	}
// 	i := md[requestinfo.InfoKey]
// 	if len(i) < 1 {
// 		return nil, common_error.ErrCanNotMappingMetadata
// 	}

// 	inf := i[0]
// 	info := &requestinfo.Info{}
// 	if err := json.Unmarshal([]byte(inf), info); err != nil {
// 		return nil, common_error.ErrCanNotExtractInfo
// 	}

// 	newCtx := context.WithValue(ctx, requestinfo.Info{}, info)
// 	return newCtx, nil
// }

// // ExtractRequestInfo ...
// func ExtractRequestInfo(ctx context.Context) (*requestinfo.Info, error) {
// 	info, ok := ctx.Value(requestinfo.Info{}).(*requestinfo.Info)
// 	if !ok {
// 		return nil, common_error.ErrCanNotExtractInfo
// 	}
// 	return info, nil
// }
