package tx

import "strconv"

// Transaction types
type TransactionType uint8

const (
	Miner_Transaction      TransactionType = 0x00
	Issue_Transaction      TransactionType = 0x01
	Claim_Transaction      TransactionType = 0x02
	Enrollment_Transaction TransactionType = 0x20
	Register_Transaction   TransactionType = 0x40
	Contract_Transaction   TransactionType = 0x80
	State_Transaction      TransactionType = 0x90
	/// <summary>
	/// Publish scripts to the blockchain for being invoked later.
	/// </summary>
	Publish_Transaction    TransactionType = 0xd0
	Invocation_Transaction TransactionType = 0xd1
)

func (t TransactionType) String() string {
	switch t {
	case 0:
		return "MinerTransaction"
	case 1:
		return "IssueTransaction"
	case 2:
		return "ClaimTransaction"
	case 32:
		return "EnrollmentTransaction"
	case 64:
		return "RegisterTransaction"
	case 128:
		return "ContractTransaction"
	case 144:
		return "StateTransaction"
	case 208:
		return "PublishTransaction"
	case 209:
		return "InvocationTransaction"
	default:
		return "TransactionType=" + strconv.FormatUint(uint64(t), 10)
	}
}