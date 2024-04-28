// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
)

// DronePlanResponse defines model for DronePlanResponse.
type DronePlanResponse struct {
	// Distance Total distance traveled by the drone
	Distance *int `json:"distance,omitempty"`
	Rest     *struct {
		// X X-coordinate of the location where the drone will first land
		X *int `json:"x,omitempty"`

		// Y Y-coordinate of the location where the drone will first land
		Y *int `json:"y,omitempty"`
	} `json:"rest,omitempty"`
}

// ErrorResponseApi defines model for ErrorResponseApi.
type ErrorResponseApi struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// StatsResponse defines model for StatsResponse.
type StatsResponse struct {
	Count  *int `json:"count,omitempty"`
	Max    *int `json:"max,omitempty"`
	Median *int `json:"median,omitempty"`
	Min    *int `json:"min,omitempty"`
}

// UuidResponse defines model for UuidResponse.
type UuidResponse struct {
	Id *string `json:"id,omitempty"`
}

// PostEstateJSONBody defines parameters for PostEstate.
type PostEstateJSONBody struct {
	Length int `json:"length"`
	Width  int `json:"width"`
}

// GetEstateIdDronePlanParams defines parameters for GetEstateIdDronePlan.
type GetEstateIdDronePlanParams struct {
	MaxDistance *int `form:"max_distance,omitempty" json:"max_distance,omitempty"`
}

// PostEstateIdTreeJSONBody defines parameters for PostEstateIdTree.
type PostEstateIdTreeJSONBody struct {
	Height int `json:"height"`
	X      int `json:"x"`
	Y      int `json:"y"`
}

// PostEstateJSONRequestBody defines body for PostEstate for application/json ContentType.
type PostEstateJSONRequestBody PostEstateJSONBody

// PostEstateIdTreeJSONRequestBody defines body for PostEstateIdTree for application/json ContentType.
type PostEstateIdTreeJSONRequestBody PostEstateIdTreeJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Insert a new estate
	// (POST /estate)
	PostEstate(ctx echo.Context) error
	// Get drone flight plan for a given estate
	// (GET /estate/{id}/drone-plan)
	GetEstateIdDronePlan(ctx echo.Context, id string, params GetEstateIdDronePlanParams) error
	// Get stats of trees in a given estate
	// (GET /estate/{id}/stats)
	GetEstateIdStats(ctx echo.Context, id string) error
	// Store tree data in a given estate
	// (POST /estate/{id}/tree)
	PostEstateIdTree(ctx echo.Context, id string) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// PostEstate converts echo context to params.
func (w *ServerInterfaceWrapper) PostEstate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEstate(ctx)
	return err
}

// GetEstateIdDronePlan converts echo context to params.
func (w *ServerInterfaceWrapper) GetEstateIdDronePlan(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetEstateIdDronePlanParams
	// ------------- Optional query parameter "max_distance" -------------

	err = runtime.BindQueryParameter("form", true, false, "max_distance", ctx.QueryParams(), &params.MaxDistance)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max_distance: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEstateIdDronePlan(ctx, id, params)
	return err
}

// GetEstateIdStats converts echo context to params.
func (w *ServerInterfaceWrapper) GetEstateIdStats(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.GetEstateIdStats(ctx, id)
	return err
}

// PostEstateIdTree converts echo context to params.
func (w *ServerInterfaceWrapper) PostEstateIdTree(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.PostEstateIdTree(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/estate", wrapper.PostEstate)
	router.GET(baseURL+"/estate/:id/drone-plan", wrapper.GetEstateIdDronePlan)
	router.GET(baseURL+"/estate/:id/stats", wrapper.GetEstateIdStats)
	router.POST(baseURL+"/estate/:id/tree", wrapper.PostEstateIdTree)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+yXUW/bNhDHvwpx26NaKSv2ordtLQpjKFYsKbChCAZGPEtXUKR6PDn2An/3gZQcx5Ey",
	"J5kDbMDeJPLIO9397v72DVS+7bxDJwHKGwhVg61Oj2/ZO/xotfsVQ+ddwLjYse+QhTCZGAqiXZV2DIaK",
	"qRPyDkq48KKt2u0rYb1Ci0ZdbZQ0qEy8GzJoyVHbt1AWGcimQyiBnGCNDNsMGINMna6n3n57VXnPhpwW",
	"VH6ZXFhf6bitrhtk3HtV12StWhIHUVY7czSKzdTf7y/nb3u75K++YCVzKxm8Y/a8q8sPHU2zVHmTyjL9",
	"nhZD0PXdzSBMrk6eGL/2xGig/Dxcsbe/nInjXLSEh/mofO9SBf8+w61eP8IIDWn3CDs6ajSX0U89mYc/",
	"hMx8uu6XapsBuaWPxpYqHO9yuo1WHxYX0ZOQ2Pj6KSCrc+QVVTHNK+Qw0HX2unhdREvfodOxuPAmLWXQ",
	"aWlSQDkG0TJE6ocuifEmAhcGSvjog7wbbIaqYpAfvdkMZXGCQ2F011kawM2/BO/2M2CaBIuuliYlV6+H",
	"5H5fFEVxB+izuYpck3n6sXssjr53l13Opn5/QLjHYYCkgqbwvyuKJ338t4xLKOGbfD8h83E85ge0JN+H",
	"A+InRi1oIG0sdW/lZK4nrT/jvne47rASNAqjeUpn6NtW8wZKWLiALEorh9dqBClajFDlN2S2eRperzo7",
	"9FyNM4y9xxGxhbkViwQp6xYFOUD5+f7oXLzdDczBmRKvopNxWC4t1Y2opWeIrQRlQh6yXQ+Rgft1zu4k",
	"btKhN8MlX3vkzf6WVq//uNWuu+cPY/0w8KpC38agV7EVKm2VdkY1nulP7w5V7nbmV9qNmpcpcmrMxpHJ",
	"f/mCwE7FfAabX37+lwL7HuWQj0TM0rPSqqYVugcpjo/hMQAnJXsOvDWKii8UhKpwMnJfEoZD1f6vgZBK",
	"murAiCG21zEIouFjlHJhLqLlUyHY/+CLjhSFiMUVqiA+lvpkNJxCwxuM/XMgxm+OCvj66Zq/+Yd6v4Z4",
	"R7aL93l6f/a/3qeuOY8gDnAaLXq2ZeIB5NUO+Z4tlNCIdGWex782tonNs73c/hUAAP//KWyNODIOAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
