package dxo

import "strconv"

////////////////////////////////////////////////////////////////////////////////

// AccountID 表示一个账号的ID
type AccountID int64

// BucketID 表示一个存储桶的ID
type BucketID int64

// CredentialID 表示凭证的ID
type CredentialID int64

// PackageID 表示一个包裹的ID
type PackageID int64

////////////////////////////////////////////////////////////////////////////////

// Int64 把 AccountID 转换成 int64
func (id AccountID) Int64() int64 {
	return int64(id)
}

////////////////////////////////////////////////////////////////////////////////

// String 把 AccountID 转换成 string
func (id AccountID) String() string {
	n := int64(id)
	return strconv.FormatInt(n, 10)
}

func (v BucketID) String() string {
	n := int64(v)
	return strconv.FormatInt(n, 10)
}

func (v PackageID) String() string {
	n := int64(v)
	return strconv.FormatInt(n, 10)
}

////////////////////////////////////////////////////////////////////////////////
