package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Mo3g4u/blog-myapi/common"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	// 変換先であるMyAppError型の変数を先に用意
	var appErr *MyAppError
	// errors.As関数で引数のerrをMyAppError型のappErrに変換する
	if !errors.As(err, &appErr) {
		// もし変換に失敗したらUnknowsエラーを変数appErrに手動で格納
		appErr = &MyAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("[%d]error: %s\n", traceID, appErr)

	var statusCode int
	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorizated:
		statusCode = http.StatusUnauthorized
	case NotMatchUser:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
