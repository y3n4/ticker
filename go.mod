module github.com/systemli/ticker

require (
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/appleboy/gofight/v2 v2.1.2
	github.com/asdine/storm v2.1.2+incompatible
	github.com/dghubble/go-twitter v0.0.0-20200725221434-4bc8ad7ad1b4
	github.com/dghubble/oauth1 v0.6.0
	github.com/disintegration/imaging v1.6.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/size v0.0.0-20200815104238-dc717522c4e2
	github.com/gin-gonic/gin v1.7.4
	github.com/google/uuid v1.1.1
	github.com/paulmach/go.geojson v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.7.1
	github.com/sethvargo/go-password v0.2.0
	github.com/sirupsen/logrus v1.6.0
	github.com/spf13/afero v1.3.4
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/toorop/gin-logrus v0.0.0-20190701131413-6c374ad36b67
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
)

require (
	github.com/DataDog/zstd v1.4.0 // indirect
	github.com/Sereal/Sereal v0.0.0-20190606082811-cf1bab6c7a3a // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dghubble/sling v1.3.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/google/go-querystring v1.0.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.2 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.8.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.13.0 // indirect
	github.com/prometheus/procfs v0.1.3 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tidwall/gjson v1.11.0 // indirect
	github.com/ugorji/go/codec v1.2.6 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/image v0.0.0-20200801110659-972c09e46d76 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/ini.v1 v1.60.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

go 1.17

replace github.com/dghubble/go-twitter => ./forks/go-twitter
