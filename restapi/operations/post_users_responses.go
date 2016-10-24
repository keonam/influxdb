package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/chronograf/models"
)

/*PostUsersCreated Successfully created new user

swagger:response postUsersCreated
*/
type PostUsersCreated struct {
	/*Location of the newly created user resource.
	  Required: true
	*/
	Location string `json:"Location"`

	// In: body
	Payload *models.User `json:"body,omitempty"`
}

// NewPostUsersCreated creates PostUsersCreated with default headers values
func NewPostUsersCreated() *PostUsersCreated {
	return &PostUsersCreated{}
}

// WithLocation adds the location to the post users created response
func (o *PostUsersCreated) WithLocation(location string) *PostUsersCreated {
	o.Location = location
	return o
}

// SetLocation sets the location to the post users created response
func (o *PostUsersCreated) SetLocation(location string) {
	o.Location = location
}

// WithPayload adds the payload to the post users created response
func (o *PostUsersCreated) WithPayload(payload *models.User) *PostUsersCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users created response
func (o *PostUsersCreated) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Location
	rw.Header().Add("Location", fmt.Sprintf("%v", o.Location))

	rw.WriteHeader(201)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PostUsersDefault A processing or an unexpected error.

swagger:response postUsersDefault
*/
type PostUsersDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostUsersDefault creates PostUsersDefault with default headers values
func NewPostUsersDefault(code int) *PostUsersDefault {
	if code <= 0 {
		code = 500
	}

	return &PostUsersDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post users default response
func (o *PostUsersDefault) WithStatusCode(code int) *PostUsersDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post users default response
func (o *PostUsersDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post users default response
func (o *PostUsersDefault) WithPayload(payload *models.Error) *PostUsersDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post users default response
func (o *PostUsersDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostUsersDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
