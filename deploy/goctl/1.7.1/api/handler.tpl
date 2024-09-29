package {{.PkgName}}

import (
    {{if .HasRequest}}"errors"{{end}}
	"net/http"
	{{if .HasRequest}}"reflect"{{end}}
    {{if .HasRequest}}"go_zero_pgsql/common/errorx"{{end}}
    {{if .HasRequest}}translations "github.com/go-playground/validator/v10/translations/zh"{{end}}

	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
	{{.ImportPackages}}
	xhttp "github.com/zeromicro/x/http"
    {{if .HasRequest}}"github.com/go-playground/locales/zh"{{end}}
    {{if .HasRequest}}ut "github.com/go-playground/universal-translator"{{end}}
    {{if .HasRequest}}"github.com/go-playground/validator/v10"{{end}}
)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
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

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		 if err != nil {
            // code-data 响应格式
            xhttp.JsonBaseResponseCtx(r.Context(), w, err)
        } else {
            // code-data 响应格式
            xhttp.JsonBaseResponseCtx(r.Context(), w, {{if .HasResp}}resp{{else}}nil{{end}})
        }
	}
}
