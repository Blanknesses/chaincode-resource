package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type ABACRequest struct {
	AS AS
	AO AO
}

func (r *ABACRequest) ToBytes() []byte {
	b, err := json.Marshal(*r)
	if err != nil {
		return nil
	}
	return b
}

func (r *ABACRequest) GetId() string {
	return fmt.Sprintf("%x", sha256.Sum256(r.ToBytes()))
}
