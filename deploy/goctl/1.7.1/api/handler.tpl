package {{.PkgName}}

import (
	{{.ImportPackages}}
	"net/http"
    xhttp "github.com/zeromicro/x/http"
	"github.com/zeromicro/go-zero/rest/httpx"

)

{{if .HasDoc}}{{.Doc}}{{end}}
func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
            return
        }


		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		 if err != nil {
            err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
        } else {
            xhttp.JsonBaseResponseCtx(r.Context(), w, {{if .HasResp}}resp{{else}}nil{{end}})
        }
	}
}
