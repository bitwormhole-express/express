package entity

import "github.com/bitwomrhole-express/express/community-server/app/data/dxo"

// Credential 表示一个凭证
type Credential struct {
	Base

	ID dxo.CredentialID
}
