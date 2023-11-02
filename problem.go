package problem

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	URLAboutBlank = "about:blank"

	fieldDetail    = "detail"
	fieldError     = "error"
	fieldInstance  = "instance"
	fieldPath      = "path"
	fieldStatus    = "status"
	fieldTimestamp = "timestamp"
	fieldTitle     = "title"
	fieldType      = "type"
)

var ErrJSON = errors.New("JSON error")

// Details is the problem details object as defined by the RFC 7807.
type Details struct { //nolint:musttag
	Type     *url.URL
	Title    string
	Status   int
	Detail   string
	Instance *url.URL

	extensions map[string]any
}

// New return a new problem details object with the given HTTP status code and detail.
func New(status int, detail string) *Details {
	return &Details{ //nolint:exhaustruct
		Detail: detail,
		Status: status,
		extensions: map[string]any{
			fieldTimestamp: time.Now(),
		},
	}
}

// Error sets the given error in the problem details with the key "error".
func (pb *Details) Error(err error) *Details {
	if err != nil {
		pb.Put(fieldError, err.Error())
	}

	return pb
}

// Request sets the URI from the given request in the problem details with the key "path".
func (pb *Details) Request(req *http.Request) *Details {
	pb.Put(fieldPath, req.RequestURI)

	return pb
}

// Put adds the given value with the given key in the problem details.
func (pb *Details) Put(key string, value any) *Details {
	pb.extensions[key] = value

	return pb
}

// Get returns the value associated with the given key.
func (pb *Details) Get(key string) (any, bool) {
	value, found := pb.extensions[key]

	return value, found
}

func (pb *Details) MarshalJSON() ([]byte, error) {
	data := pb.asMap()

	res, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSON, err.Error())
	}

	return res, nil
}

func (pb *Details) asMap() map[string]any {
	result := pb.extensions

	if pb.Type == nil {
		result[fieldType] = URLAboutBlank

		if status, found := StatusByCode[pb.Status]; found {
			result[fieldTitle] = status
		}
	} else {
		result[fieldType] = pb.Type.String()

		if pb.Title != "" {
			result[fieldTitle] = pb.Title
		}
	}

	if pb.Status > 0 {
		result[fieldStatus] = pb.Status
	}

	if pb.Detail != "" {
		result[fieldDetail] = pb.Detail
	}

	if pb.Instance != nil {
		result[fieldInstance] = pb.Instance.String()
	}

	return result
}

func (pb *Details) UnmarshalJSON(bytes []byte) error {
	var err error

	data := make(map[string]any)

	if err = json.Unmarshal(bytes, &data); err != nil {
		return fmt.Errorf("%w: %s", ErrJSON, err.Error())
	}

	pb.Title = pop[string](data, fieldTitle)
	pb.Status = int(pop[float64](data, fieldStatus))
	pb.Detail = pop[string](data, fieldDetail)

	if pb.Type, err = popURL(data, fieldType, URLAboutBlank); err != nil {
		return err
	}

	if pb.Instance, err = popURL(data, fieldInstance, ""); err != nil {
		return err
	}

	pb.extensions = data

	return nil
}

func pop[T any](data map[string]any, key string) T { //nolint:ireturn
	var dft T

	value, found := data[key]
	if !found {
		return dft
	}

	res, ok := value.(T)
	if !ok {
		return dft
	}

	delete(data, key)

	return res
}

func popURL(data map[string]any, key, dft string) (*url.URL, error) {
	s := pop[string](data, key)
	if s == "" {
		s = dft
	}

	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrJSON, err.Error())
	}

	return u, nil
}
