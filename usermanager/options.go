package usermanager

import (
	datastore "github.com/fulgurant/datastore"
	simplehash "github.com/fulgurant/simplehash"
	"go.uber.org/zap"
)

type Options struct {
	logger      *zap.Logger
	ds          datastore.GetSetter
	hasher      simplehash.Hasher
	autoApprove bool
	debug       bool
}

func DefaultOptions() *Options {
	return &Options{}
}

func (o *Options) WithAutoApprove(value bool) *Options {
	o.autoApprove = value
	return o
}

func (o *Options) WithDebug(value bool) *Options {
	o.debug = value
	return o
}

func (o *Options) WithGetSetter(value datastore.GetSetter) *Options {
	o.ds = value
	return o
}

func (o *Options) WithLogger(value *zap.Logger) *Options {
	o.logger = value.With(zap.String("module", "usermanager"))
	return o
}

func (o *Options) WithHasher(value simplehash.Hasher) *Options {
	o.hasher = value
	return o
}
