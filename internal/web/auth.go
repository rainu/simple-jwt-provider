package web

import (
	"encoding/json"
	"errors"
	"github.com/leberKleber/simple-jwt-provider/internal"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		EMail    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if requestBody.EMail == "" {
		writeError(w, http.StatusBadRequest, "email must be set")
		return
	}

	if requestBody.Password == "" {
		writeError(w, http.StatusBadRequest, "password must be set")
		return
	}

	jwt, err := s.p.Login(requestBody.EMail, requestBody.Password)
	if err != nil {
		if errors.Is(err, internal.ErrIncorrectPassword) || errors.Is(err, internal.ErrUserNotFound) {
			logrus.WithField("email", requestBody.EMail).Warn("somebody tried to login with invalid credentials")
			writeError(w, http.StatusUnauthorized, "invalid credentials")
			return
		}

		logrus.WithError(err).Error("Failed to login User")
		writeInternalServerError(w)
		return
	}

	err = json.NewEncoder(w).Encode(struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: jwt,
	})
	if err != nil {
		logrus.WithError(err).Error("Failed marshal request response")
		writeInternalServerError(w)
		return
	}
}

func (s *Server) passwordResetRequestHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		EMail string `json:"email"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if requestBody.EMail == "" {
		writeError(w, http.StatusBadRequest, "email must be set")
		return
	}

	err = s.p.CreatePasswordResetRequest(requestBody.EMail)
	if err != nil {
		if errors.Is(err, internal.ErrUserNotFound) {
			logrus.WithField("email", requestBody.EMail).Warn("somebody tried to create a reset-password-request for non existing User")
			w.WriteHeader(http.StatusCreated)
			return
		}

		logrus.WithError(err).Error("Failed to create password-reset-request")
		writeInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	return
}

func (s *Server) passwordResetHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := struct {
		EMail      string `json:"email"`
		ResetToken string `json:"reset_token"`
		Password   string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	if requestBody.EMail == "" {
		writeError(w, http.StatusBadRequest, "email must be set")
		return
	}

	if requestBody.ResetToken == "" {
		writeError(w, http.StatusBadRequest, "reset-token must be set")
		return
	}

	if requestBody.Password == "" {
		writeError(w, http.StatusBadRequest, "password must be set")
		return
	}

	err = s.p.ResetPassword(requestBody.EMail, requestBody.ResetToken, requestBody.Password)
	if err != nil {
		if errors.Is(err, internal.ErrNoValidTokenFound) {
			writeError(w, http.StatusBadRequest, "reset-token is invalid or token email combination is not correct")
			return
		}
		logrus.WithError(err).Error("Failed to create password-reset-request")
		writeInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
