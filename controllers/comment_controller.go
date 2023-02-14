package controllers

import (
	"encoding/json"
	"go-way-to-intermediate/apperrors"
	"go-way-to-intermediate/controllers/services"
	"go-way-to-intermediate/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment のハンドラ
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {

	var reqComment models.Comment

	// jsonデコードを行なっている場所
	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		// http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		apperrors.ErrorHandler(w, req, err)
	}

	// サービス層から処理結果を受け取る場所
	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		// http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
