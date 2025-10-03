module cuhara.qua.go

go 1.24.0

toolchain go1.24.1

require gorm.io/gorm v1.31.0

require (
	github.com/aarondl/null/v8 v8.1.3
	github.com/aarondl/sqlboiler/v4 v4.19.5
	github.com/aarondl/strmangle v0.0.9
	github.com/friendsofgo/errors v0.9.2
	github.com/getkin/kin-openapi v0.132.0
	github.com/labstack/echo/v4 v4.13.4
	github.com/oapi-codegen/runtime v1.1.2
	github.com/rs/zerolog v1.34.0
)

require (
	github.com/aarondl/inflect v0.0.2 // indirect
	github.com/aarondl/randomize v0.0.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-openapi/jsonpointer v0.22.0 // indirect
	github.com/go-openapi/swag/jsonname v0.24.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oasdiff/yaml v0.0.0-20250309154309-f31be36b4037 // indirect
	github.com/oasdiff/yaml3 v0.0.0-20250309153720-d2182401db90 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/net v0.44.0 // indirect
	golang.org/x/sys v0.36.0 // indirect
	golang.org/x/time v0.11.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)

require (
	github.com/go-playground/validator/v10 v10.27.0
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/lib/pq v1.10.9
	github.com/oapi-codegen/echo-middleware v1.0.2
	github.com/subosito/gotenv v1.6.0
	github.com/timewasted/go-accept-headers v0.0.0-20130320203746-c78f304b1b09
	golang.org/x/crypto v0.42.0
	golang.org/x/text v0.29.0 // indirect
)

replace github.com/volatiletech/sqlboiler/v4 => github.com/aarondl/sqlboiler/v4 v4.19.5
