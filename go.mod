module github.com/systemli/ticker

require (
	github.com/appleboy/gin-jwt/v2 v2.8.0
	github.com/appleboy/gofight/v2 v2.1.2
	github.com/asdine/storm v2.1.2+incompatible
	github.com/dghubble/go-twitter v0.0.0-20211115160449-93a8679adecb
	github.com/dghubble/oauth1 v0.7.1
	github.com/disintegration/imaging v1.6.2
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/size v0.0.0-20220102055520-f75bacbc2df3
	github.com/gin-gonic/gin v1.7.7
	github.com/go-telegram-bot-api/telegram-bot-api/v5 v5.5.1
	github.com/google/uuid v1.3.0
	github.com/paulmach/go.geojson v1.4.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.0
	github.com/sethvargo/go-password v0.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/afero v1.8.0
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.2
	github.com/swaggo/swag v1.7.8
	github.com/toorop/gin-logrus v0.0.0-20210225092905-2c785434f26f
	golang.org/x/crypto v0.0.0-20220112180741-5e0467b6c7ce
)

require (
	github.com/DataDog/zstd v1.4.0 // indirect
	github.com/Sereal/Sereal v0.0.0-20190606082811-cf1bab6c7a3a // indirect
	github.com/go-playground/validator/v10 v10.10.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/image v0.0.0-20211028202545-6944b10bf410 // indirect
	gopkg.in/ini.v1 v1.66.3 // indirect
)

go 1.16

replace github.com/dghubble/go-twitter => github.com/0x46616c6b/go-twitter v0.0.0-media
