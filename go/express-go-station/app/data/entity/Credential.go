package entity

import "github.com/bitwomrhole-express/express/station/app/data/dxo"

// Credential 表示一个凭证
type Credential struct {
	Base

	ID dxo.CredentialID
}
