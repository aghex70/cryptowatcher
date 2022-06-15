package handlers

import (
	"cwapi/internal/core/ports"
	"cwapi/internal/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/gddo/httputil/header"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"strings"
)

type RestHandler struct {
	tradeUseCase ports.TradeUseCase
	userUseCase  ports.UserUseCase
	logger       *zap.Logger
}

func NewRestHandler(tradeUseCase ports.TradeUseCase, userUseCase ports.UserUseCase) RestHandler {
	return RestHandler{
		tradeUseCase: tradeUseCase,
		userUseCase:  userUseCase,
	}
}

func (h RestHandler) FetchOrders(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	// Get source param from request body
	var fr ports.FetchRequest
	err := decodeJSONBody(w, r, &fr)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	res, err := h.tradeUseCase.Fetch(fr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if res != nil {
		_, err := w.Write([]byte(res.String()))
		if err != nil {
			return
		}
	}
}

func (h RestHandler) StopFetchOrders(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	err := r.ParseForm()
	if err != nil {
		return
	}
	source := r.Form.Get("source")
	sfr := ports.StopFetchRequest{Source: source}
	err = h.tradeUseCase.StopFetch(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetTrades(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	trades, err := h.tradeUseCase.Get()
	b, err := json.Marshal(trades)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetSales(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	sales, err := h.tradeUseCase.GetSales()
	b, err := json.Marshal(sales)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetPurchases(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	purchases, err := h.tradeUseCase.GetPurchases()
	b, err := json.Marshal(purchases)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	userID := 1
	user, err := h.userUseCase.Get(userID)
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h RestHandler) GetUserTrades(w http.ResponseWriter, r *http.Request) {
	logger.ZapLogger.Info("FetchOrders")
	userID := 1
	trades, err := h.userUseCase.GetUserTrades(userID)
	b, err := json.Marshal(trades)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type malformedRequest struct {
	status int
	msg    string
}

func (mr *malformedRequest) Error() string {
	return mr.msg
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequest{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := fmt.Sprintf("Request body contains badly-formed JSON")
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequest{status: http.StatusBadRequest, msg: msg}

		case err.Error() == "http: request body too large":
			msg := "Request body must not be larger than 1MB"
			return &malformedRequest{status: http.StatusRequestEntityTooLarge, msg: msg}

		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequest{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}
