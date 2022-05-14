package m

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

// Define the Policy structure
type Policy struct {
	AS AS
	AO AO
	AP int //1 is allow , 0 is deney
	AE AE
}

type AS struct {
	UserAccount string `json:"userAccount"`
	UserRole   string `json:"userRole"`
	UserOrg  string `json:"userOrg"`
}

type AO struct {
	IIoTID string `json:"IIoTID"`
}

type AE struct {
	CreatedTime string  `json:"createdTime"`
	EndTime     string  `json:"endTime"`
	AllowedIP   string `json:"allowedIP"`
	RateLimit string `json:"rateLimit"`
	FlowLimit string `json:"flowLimit"`
}

func (p *Policy) ToBytes() []byte {
	b, err := json.Marshal(*p)
	if err != nil {
		return nil
	}
	return b
}

func (p *Policy) GetID() string {
	as, _ := json.Marshal(p.AS)
	ao, _ := json.Marshal(p.AO)
	result := append(as, ao...)
	return fmt.Sprintf("%x", sha256.Sum256(result))
}
