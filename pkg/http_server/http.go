package http_server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	swaggerui "github.com/esurdam/go-swagger-ui"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewServerMux(docFileName, docFileRoot string, handleFunc func(ctx context.Context, mux *runtime.ServeMux)) (http.Handler, error) {
	ctx := context.Background()
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{MarshalOptions: protojson.MarshalOptions{UseEnumNumbers: false}, UnmarshalOptions: protojson.UnmarshalOptions{AllowPartial: true}}),
		// runtime.WithMetadata(mapMetaData),
	)
	handleFunc(ctx, mux)
	muxh := swaggerui.NewServeMuxWithRoot(func(s string) ([]byte, error) {
		file, err := ioutil.ReadFile(filepath.Join(fmt.Sprintf("../../%s/%s", docFileRoot, s)))
		if err != nil {
			return nil, err
		}
		return file, nil
	}, docFileName, docFileRoot)
	muxh.Handle("/", AllowCORS(mux))

	return muxh, nil
}
