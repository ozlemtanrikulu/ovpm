package api

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/asaskevich/govalidator"
	"github.com/cad/ovpm/api/pb"
	"github.com/cad/ovpm/bindata"
	"github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// NewRESTServer returns a new REST server.
func NewRESTServer(grpcPort string) (http.Handler, context.CancelFunc, error) {
	mux := http.NewServeMux()
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	if !govalidator.IsNumeric(grpcPort) {
		return nil, cancel, fmt.Errorf("grpcPort should be numeric")
	}
	endPoint := fmt.Sprintf("localhost:%s", grpcPort)
	ctx = NewOriginTypeContext(ctx, OriginTypeREST)
	gmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterVPNServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	err = pb.RegisterUserServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	err = pb.RegisterNetworkServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	err = pb.RegisterAuthServiceHandlerFromEndpoint(ctx, gmux, endPoint, opts)
	if err != nil {
		return nil, cancel, err
	}

	mux.HandleFunc("/specs/", specsHandler)
	mware := middleware.Redoc(middleware.RedocOpts{
		BasePath: "/docs/",
		SpecURL:  "/specs/user.swagger.json",
		Path:     "user",
	}, gmux)
	mware = middleware.Redoc(middleware.RedocOpts{
		BasePath: "/docs/",
		SpecURL:  "/specs/vpn.swagger.json",
		Path:     "vpn",
	}, mware)
	mware = middleware.Redoc(middleware.RedocOpts{
		BasePath: "/docs/",
		SpecURL:  "/specs/network.swagger.json",
		Path:     "network",
	}, mware)
	mware = middleware.Redoc(middleware.RedocOpts{
		BasePath: "/docs/",
		SpecURL:  "/specs/auth.swagger.json",
		Path:     "auth",
	}, mware)
	mux.Handle("/", mware)
	return mux, cancel, nil
}

func specsHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	switch r.URL.Path {
	case "/specs/user.swagger.json":
		userData, err := bindata.Asset("template/user.swagger.json")
		if err != nil {
			logrus.Warn(err)
		}
		w.Write(userData)

	case "/specs/network.swagger.json":
		networkData, err := bindata.Asset("template/network.swagger.json")
		if err != nil {
			logrus.Warn(err)
		}
		w.Write(networkData)
	case "/specs/vpn.swagger.json":
		vpnData, err := bindata.Asset("template/vpn.swagger.json")
		if err != nil {
			logrus.Warn(err)
		}
		w.Write(vpnData)
	case "/specs/auth.swagger.json":
		vpnData, err := bindata.Asset("template/auth.swagger.json")
		if err != nil {
			logrus.Warn(err)
		}
		w.Write(vpnData)
	}
}
