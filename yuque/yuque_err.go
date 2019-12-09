package yuque

import "github.com/pubgo/g/xerror"

type YQErr xerror.Err

func (t *YQErr) Error() error {
	return nil
}

func (t *YQErr) IsNotFoundErr() bool {
	return true
}
