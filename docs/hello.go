package docs

import "github.com/vaclav-dvorak/go-api-hello-world/api"

// swagger:operation GET /hello/{name} hello Hello
//
// Returns a simple Hello message
// ---
// consumes:
// - text/plain
// produces:
// - application/json
// parameters:
// - name: name
//   in: path
//   description: Name to be returned.
//   required: true
//   type: string
// responses:
//   200:
//     $ref: '#/responses/helloResponse'

// The hello message.
// swagger:response helloResponse
type helloResponseWrapper struct {
  // in:body
  Body api.HelloResponse
}
