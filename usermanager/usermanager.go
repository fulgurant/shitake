package usermanager

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	datastore "github.com/fulgurant/datastore"

	"go.uber.org/zap"
)

// Constants
var (
	userBucket = []byte("user")
)

// Errors
var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrNoDatastore   = errors.New("no datastore specified")
	ErrNoHasher      = errors.New("no hasher specified")
)

type UserManager struct {
	options *Options
}

func New(options *Options) (*UserManager, error) {
	if options.ds == nil {
		return nil, ErrNoDatastore
	}

	if options.hasher == nil {
		return nil, ErrNoHasher
	}

	ds := &UserManager{
		options: options,
	}

	return ds, nil
}

func (um *UserManager) RegisterEndpoints(r gin.IRouter) {
	r.POST("/user/signup", um.postSignup)

	if um.options.logger != nil {
		um.options.logger.Info("endpoints registered")
	}

	if um.options.debug {
		r.GET("/users", um.getUsers)

		if um.options.logger != nil {
			um.options.logger.Info("debug endpoints registered")
		}
	}
}

func (um *UserManager) postSignup(ctx *gin.Context) {
	u := User{}

	if err := ctx.ShouldBind(&u); err != nil {
		if um.options.logger != nil {
			um.options.logger.Info("could not bind user body")
		}
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	if err := um.Signup(&u); err != nil {
		if um.options.logger != nil {
			um.options.logger.Info("could not signup", zap.Error(err))
		}
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}

	ctx.String(http.StatusOK, "OK")
}

// getUsers is a debug endpoint
func (um *UserManager) getUsers(ctx *gin.Context) {
	builder := strings.Builder{}

	err := um.options.ds.List(userBucket, nil, func(key []byte, value []byte) error {
		builder.Write(key)
		builder.WriteRune('\n')
		return nil
	})

	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, builder.String())
}

func (um *UserManager) Signup(u *User) error {
	if _, err := um.options.ds.Get(userBucket, []byte(u.Name)); err != datastore.ErrNotFound {
		return ErrAlreadyExists
	}

	if err := u.Check(); err != nil {
		return err
	}

	if err := u.Hash(um.options.hasher); err != nil {
		return err
	}

	u.Approved = um.options.autoApprove

	b, err := u.ToBytes()
	if err != nil {
		return err
	}

	if err := um.options.ds.Set(userBucket, []byte(u.Name), b); err != nil {
		if um.options.logger != nil {
			um.options.logger.Error("could not signup", zap.Error(err))
		}
		return err
	}

	if um.options.logger != nil {
		um.options.logger.Info("signup", zap.String("username", u.Name))
	}

	return nil
}
