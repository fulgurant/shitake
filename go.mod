module github.com/fulgurant/shitake

go 1.16

replace (
	github.com/fulgurant/datastore => ../datastore
	github.com/fulgurant/health => ../health
	github.com/fulgurant/server => ../server
	github.com/fulgurant/simplehash => ../simplehash
)

require (
	github.com/alecthomas/kong v0.2.16
	github.com/fulgurant/datastore v0.0.4
	github.com/fulgurant/health v0.0.2
	github.com/fulgurant/server v0.0.0-20210417220906-cfd6b150159d
	github.com/fulgurant/simplehash v0.0.3
	github.com/gin-gonic/gin v1.7.1
	go.uber.org/zap v1.16.0
)
