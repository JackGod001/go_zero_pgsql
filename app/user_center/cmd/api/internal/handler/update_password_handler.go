package handler

import (
	"errors"
	translations "github.com/go-playground/validator/v10/translations/zh"
	"go_zero_pgsql/common/errorx"
	"net/http"
	"reflect"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_pgsql/app/user_center/cmd/api/internal/logic"
	"go_zero_pgsql/app/user_center/cmd/api/internal/svc"
	"go_zero_pgsql/app/user_center/cmd/api/internal/types"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	xhttp "github.com/zeromicro/x/http"
)

// 修改密码
func UpdatePasswordHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePasswordRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, err.Error()))
			return
		}

		validate := validator.New()
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

		trans, _ := ut.New(zh.New()).GetTranslator("zh")
		validateErr := translations.RegisterDefaultTranslations(validate, trans)
		if validateErr = validate.StructCtx(r.Context(), req); validateErr != nil {
			for _, err := range validateErr.(validator.ValidationErrors) {
				httpx.Error(w, errorx.NewHandlerError(errorx.ParamErrorCode, errors.New(err.Translate(trans)).Error()))
				return
			}
		}

		l := logic.NewUpdatePasswordLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePassword(&req)
		if err != nil {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			// code-data 响应格式
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
