package cors

import (
	"github.com/gominima/minima"
	"github.com/rs/cors"
	"net/http"
)

type Options = cors.Options

/**
@info The CorsWrapper for minima's instancr
@property {*cors.Cors} [raw] The raw cors instance around which wrapper build
@property {bool} [optionPassthrough] The bool value of options
*/
type corsWrapper struct {
	*cors.Cors
	optionPassthrough bool
}

/**
@info Creates a corsWrapper instance
@return {corsWrapper}
*/
func New() *corsWrapper {
	return &corsWrapper{}
}

/**
@info Builds the raw cors to minima handler
@return {minima.Handler}
*/
func (c *corsWrapper) Build() minima.Handler {
	return func(res *minima.Response, req *minima.Request) {
		c.HandlerFunc(res.Raw(), req.Raw())
		if !c.optionPassthrough && req.Method() == http.MethodOptions && req.GetHeader("Access-Control-Request-Method") != "" {
			res.OK()
			res.CloseConn()
		}

	}
}

/**
@info Sets headers to allow all origins to make a request
@return {minima.Handler}
*/
func (c *corsWrapper) AllowAll() minima.Handler {
	crs := &corsWrapper{Cors: cors.AllowAll()}
	return crs.Build()
}

/**
@info Sets headers based on default cors setup
@return {minima.Handler}
*/
func (c *corsWrapper) Default() minima.Handler {
	crs := &corsWrapper{Cors: cors.Default()}
	return crs.Build()
}

/**
@info Sets headers based on the options given
@return {minima.Handler}
*/
func (c *corsWrapper) NewCors(options Options) minima.Handler {
	crs := &corsWrapper{Cors: cors.New(options), optionPassthrough: options.OptionsPassthrough}
	return crs.Build()
}
