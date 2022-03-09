package entity

import "bitwomrhole.com/djaf/express-go-server/server/data/dxo"

// Credential 表示一个凭证
type Credential struct {
	Base

	ID dxo.CredentialID
}
