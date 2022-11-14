package main

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

// /v1/:key
func getPutHandler(s Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Param("key")

		value, err := io.ReadAll(c.Request().Body)
		if err != nil {
			http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
			return nil
		}

		err = s.Put(key, string(value))
		if err != nil {
			if err != nil {
				http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
				return nil
			}
		}

		c.Response().WriteHeader(http.StatusCreated)

		return nil
	}
}

// /v1/:key
func getGetHandler(s Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Param("key")

		value, err := s.Get(key)
		if err != nil {
			var message string
			var code int
			switch errors.Cause(err) {
			case ErrNoSuchKey:
				message = err.Error()
				code = http.StatusNotFound
			default:
				message = err.Error()
				code = http.StatusInternalServerError
			}
			http.Error(c.Response().Writer, message, code)
			return nil
		}

		return c.String(http.StatusOK, value)
	}
}

// /v1/:key
func getDeleteHandler(s Store) echo.HandlerFunc {
	return func(c echo.Context) error {
		key := c.Param("key")

		err := s.Delete(key)
		if err != nil {
			http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
			return nil
		}

		c.Response().WriteHeader(http.StatusCreated)

		return nil
	}
}
