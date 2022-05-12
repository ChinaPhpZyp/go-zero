package {{.PkgName}}

import (
	"net/http"

	"github.com/xiaoshouchen/go-zero/rest/httpx"
	{{.ImportPackages}}
	"hbb_micro/common/response"
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
			response.Fail(w, err)
			return
		}

		{{end}}l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
		if err != nil {
			response.Fail(w, err)
		} else {
			{{if .HasResp}}response.Success(w, resp){{else}}response.Success(w, nil){{end}}
		}
	}
}
