package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"project2/internal/api/middleware"
	service_interfaces "project2/internal/domain/interfaces/service"
	"project2/internal/models"
	"project2/pkg/errs"
	"project2/pkg/logger"
	"time"
)

type LeaderboardHandler struct {
	leaderBoardService service_interfaces.LeaderboardService
}

func NewLeaderboardHandler(leaderBoardService service_interfaces.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{
		leaderBoardService: leaderBoardService,
	}
}

func (l *LeaderboardHandler) GetGameLeaderboardHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameIdStr := vars["gameID"]

	// Start log
	logger.Logger.Infow("Processing request to get game leaderboard", "gameID", gameIdStr, "method", r.Method, "time", time.Now())

	gameId, err := uuid.Parse(gameIdStr)
	if err != nil {
		logger.Logger.Errorw("Error parsing game ID", "gameID", gameIdStr, "error", err, "time", time.Now())
		errs.NewInternalServerError("Couldn't parse game id").ToJSON(w)
		return
	}

	leaderboard, err := l.leaderBoardService.GetGameLeaderboard(r.Context(), gameId)
	if err != nil {
		logger.Logger.Errorw("Error fetching leaderboard", "gameID", gameIdStr, "error", err, "time", time.Now())
		errs.NewInternalServerError("Couldn't get leaderboard").ToJSON(w)
		return
	}

	// Respond with leaderboard
	w.Header().Set("Content-Type", "application/json")
	jsonResponse := map[string]any{
		"code":    http.StatusOK,
		"message": "Success",
		"leaderboard": func() []models.Leaderboard {
			if leaderboard == nil {
				return []models.Leaderboard{}
			}
			return leaderboard
		}(),
	}
	json.NewEncoder(w).Encode(jsonResponse)
	logger.Logger.Infow("Successfully fetched game leaderboard", "gameID", gameIdStr, "method", r.Method, "time", time.Now())
}

//func (l *LeaderboardHandler) AddLossToUserHandler(w http.ResponseWriter, r *http.Request) {
//	userIdStr, ok := r.Context().Value(middleware.UserIdKey).(string)
//	if !ok {
//		logger.Logger.Errorw("Error finding userId in context", "method", r.Method, "time", time.Now())
//		errs.NewUnexpectedError("Could not find the userId").ToJSON(w)
//		return
//	}
//
//	logger.Logger.Infow("Processing request to add loss", "userID", userIdStr, "method", r.Method, "time", time.Now())
//	userId, err := uuid.Parse(userIdStr)
//	if err != nil {
//		logger.Logger.Errorw("Error parsing user ID", "userID", userIdStr, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't parse user id").ToJSON(w)
//		return
//	}
//
//	var requestBody struct {
//		GameId    string `json:"game_id" validate:"required"`
//		BookingId string `json:"booking_id" validate:"required"`
//	}
//	err = json.NewDecoder(r.Body).Decode(&requestBody)
//	if err != nil {
//		logger.Logger.Errorw("Error decoding request body", "method", r.Method, "error", err, "time", time.Now())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err = validate.Struct(requestBody)
//	if err != nil {
//		logger.Logger.Errorw("Validation error", "method", r.Method, "error", err, "requestBody", requestBody, "time", time.Now())
//		errs.NewInvalidBodyError("Invalid request body").ToJSON(w)
//		return
//	}
//
//	gameId, err := uuid.Parse(requestBody.GameId)
//	if err != nil {
//		logger.Logger.Errorw("Error parsing game ID", "gameID", requestBody.GameId, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't parse game id").ToJSON(w)
//		return
//	}
//	bookingId, err := uuid.Parse(requestBody.BookingId)
//	if err != nil {
//		logger.Logger.Errorw("Error parsing booking ID", "bookingID", requestBody.BookingId, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't parse booking id").ToJSON(w)
//		return
//	}
//
//	err = l.leaderBoardService.AddLossToUser(r.Context(), userId, gameId, bookingId)
//	if err != nil {
//		logger.Logger.Errorw("Error adding loss to user", "userID", userIdStr, "gameID", requestBody.GameId, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't add loss to user").ToJSON(w)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	jsonResponse := map[string]any{
//		"code":    http.StatusOK,
//		"message": "Success",
//	}
//	json.NewEncoder(w).Encode(jsonResponse)
//	logger.Logger.Infow("Loss added successfully", "userID", userIdStr, "gameID", requestBody.GameId, "method", r.Method, "time", time.Now())
//}
//
//func (l *LeaderboardHandler) AddWinToUserHandler(w http.ResponseWriter, r *http.Request) {
//	userIdStr, ok := r.Context().Value(middleware.UserIdKey).(string)
//	if !ok {
//		logger.Logger.Errorw("Error finding userId in context", "method", r.Method, "time", time.Now())
//		errs.NewUnexpectedError("Could not find the userId").ToJSON(w)
//		return
//	}
//
//	logger.Logger.Infow("Processing request to add win", "userID", userIdStr, "method", r.Method, "time", time.Now())
//	userId, err := uuid.Parse(userIdStr)
//	if err != nil {
//		logger.Logger.Errorw("Error parsing user ID", "userID", userIdStr, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't parse user id").ToJSON(w)
//		return
//	}
//
//	var requestBody struct {
//		GameId    string `json:"game_id" validate:"required"`
//		BookingId string `json:"booking_id" validate:"required"`
//	}
//	err = json.NewDecoder(r.Body).Decode(&requestBody)
//	if err != nil {
//		logger.Logger.Errorw("Error decoding request body", "method", r.Method, "error", err, "time", time.Now())
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	err = validate.Struct(requestBody)
//	if err != nil {
//		logger.Logger.Errorw("Validation error", "method", r.Method, "error", err, "requestBody", requestBody, "time", time.Now())
//		errs.NewInvalidBodyError("Invalid request body").ToJSON(w)
//		return
//	}
//
//	gameId, err := uuid.Parse(requestBody.GameId)
//	if err != nil {
//		logger.Logger.Errorw("Error parsing game ID", "gameID", requestBody.GameId, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't parse game id").ToJSON(w)
//		return
//	}
//	bookingId, err := uuid.Parse(requestBody.BookingId)
//	if err != nil {
//		logger.Logger.Errorw("Error parsing booking ID", "bookingID", requestBody.BookingId, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't parse booking id").ToJSON(w)
//		return
//	}
//
//	err = l.leaderBoardService.AddWinToUser(r.Context(), userId, gameId, bookingId)
//	if err != nil {
//		logger.Logger.Errorw("Error adding win to user", "userID", userIdStr, "gameID", requestBody.GameId, "error", err, "time", time.Now())
//		errs.NewInternalServerError("Couldn't add win to user").ToJSON(w)
//		return
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	jsonResponse := map[string]any{
//		"code":    http.StatusOK,
//		"message": "Success",
//	}
//	json.NewEncoder(w).Encode(jsonResponse)
//	logger.Logger.Infow("Win added successfully", "userID", userIdStr, "gameID", requestBody.GameId, "method", r.Method, "time", time.Now())
//}

func (l *LeaderboardHandler) RecordUserResultHandler(w http.ResponseWriter, r *http.Request) {
	userIdStr, ok := r.Context().Value(middleware.UserIdKey).(string)
	if !ok {
		logger.Logger.Errorw("Error finding userId in context", "method", r.Method, "time", time.Now())
		errs.NewUnexpectedError("Could not find the userId").ToJSON(w)
		return
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		logger.Logger.Errorw("Error parsing user ID", "userID", userIdStr, "error", err, "time", time.Now())
		errs.NewInternalServerError("Couldn't parse user id").ToJSON(w)
		return
	}

	var requestBody struct {
		GameId    string `json:"game_id" validate:"required"`
		BookingId string `json:"booking_id" validate:"required"`
		Result    string `json:"result" validate:"required"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		logger.Logger.Errorw("Error decoding request body", "method", r.Method, "error", err, "time", time.Now())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = validate.Struct(requestBody)
	if err != nil {
		logger.Logger.Errorw("Validation error", "method", r.Method, "error", err, "requestBody", requestBody, "time", time.Now())
		errs.NewBadRequestError("Invalid request body").ToJSON(w)
		return
	}

	gameId, err := uuid.Parse(requestBody.GameId)
	if err != nil {
		logger.Logger.Errorw("Error parsing game ID", "gameID", requestBody.GameId, "error", err, "time", time.Now())
		errs.NewInternalServerError("Couldn't parse game id").ToJSON(w)
		return
	}
	bookingId, err := uuid.Parse(requestBody.BookingId)
	if err != nil {
		logger.Logger.Errorw("Error parsing booking ID", "bookingID", requestBody.BookingId, "error", err, "time", time.Now())
		errs.NewInternalServerError("Couldn't parse booking id").ToJSON(w)
		return
	}

	switch requestBody.Result {
	case "win":
		err = l.leaderBoardService.AddWinToUser(r.Context(), userId, gameId, bookingId)
		if err != nil {
			logger.Logger.Errorw("Error adding win to user", "userID", userIdStr, "gameID", requestBody.GameId, "error", err, "time", time.Now())
			errs.NewInternalServerError("Couldn't add win to user").ToJSON(w)
			return
		}

	case "loss":
		err = l.leaderBoardService.AddLossToUser(r.Context(), userId, gameId, bookingId)
		if err != nil {
			logger.Logger.Errorw("Error adding loss to user", "userID", userIdStr, "gameID", requestBody.GameId, "error", err, "time", time.Now())
			errs.NewInternalServerError("Couldn't add win to user").ToJSON(w)
			return
		}
	default:
		errs.NewBadRequestError("Invalid or missing result parameter in the URL").ToJSON(w)
		logger.Logger.Errorw("Invalid or no parameter sent in the URl", "userID", userIdStr, "gameID", requestBody.GameId, "parameter", requestBody.Result, "time", time.Now())
		return
	}

}
