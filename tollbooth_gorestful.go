package tollbooth_gorestful

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/emicklei/go-restful"
)

func LimitHandler(handler restful.RouteFunction, lmt *limiter.Limiter) restful.RouteFunction {
	return func(request *restful.Request, response *restful.Response) {
		httpError := tollbooth.LimitByRequest(lmt, request.Request)
		if httpError != nil {
			response.AddHeader("Content-Type", lmt.GetMessageContentType())
			response.WriteErrorString(httpError.StatusCode, httpError.Message)
			return
		}

		handler(request, response)
	}
}
