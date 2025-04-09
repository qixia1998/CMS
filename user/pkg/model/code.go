package model

import errs "common/errrs"

var (
	NoLegalMobile = errs.NewError(2001, "手机号不合法")
)
