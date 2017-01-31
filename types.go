package nicehash

type AlgoType int

const (
	AlgoTypeScrypt AlgoType = iota
	AlgoTypeSHA256
	AlgoTypeScryptNf
	AlgoTypeX11
	AlgoTypeX13
	AlgoTypeKeccak
	AlgoTypeX15
	AlgoTypeNist5
	AlgoTypeNeoScrypt
	AlgoTypeLyra2RE
	AlgoTypeWhirlpoolX
	AlgoTypeQubit
	AlgoTypeQuark
	AlgoTypeAxiom
	AlgoTypeLyra2REv2
	AlgoTypeScryptJaneNf16
	AlgoTypeBlake256r8
	AlgoTypeBlake256r14
	AlgoTypeBlake256r8vnl
	AlgoTypeHodl
	AlgoTypeDaggerHashimoto
	AlgoTypeDecred
	AlgoTypeCryptoNight
	AlgoTypeLbry
	AlgoTypeEquihash
)

func (t AlgoType) ToString() string {
	switch t {
	case AlgoTypeScrypt:
		return "Scrypt";
	case AlgoTypeSHA256:
		return "SHA256";
	case AlgoTypeScryptNf:
		return "ScryptNf";
	case AlgoTypeX11:
		return "X11";
	case AlgoTypeX13:
		return "X13";
	case AlgoTypeKeccak:
		return "Keccak";
	case AlgoTypeX15:
		return "X15";
	case AlgoTypeNist5:
		return "Nist5";
	case AlgoTypeNeoScrypt:
		return "NeoScrypt";
	case AlgoTypeLyra2RE:
		return "Lyra2RE";
	case AlgoTypeWhirlpoolX:
		return "WhirlpoolX";
	case AlgoTypeQubit:
		return "Qubit";
	case AlgoTypeQuark:
		return "Quark";
	case AlgoTypeAxiom:
		return "Axiom";
	case AlgoTypeLyra2REv2:
		return "Lyra2REv2";
	case AlgoTypeScryptJaneNf16:
		return "ScryptJaneNf16";
	case AlgoTypeBlake256r8:
		return "Blake256r8";
	case AlgoTypeBlake256r14:
		return "Blake256r14";
	case AlgoTypeBlake256r8vnl:
		return "Blake256r8vnl";
	case AlgoTypeHodl:
		return "Hodl";
	case AlgoTypeDaggerHashimoto:
		return "DaggerHashimoto";
	case AlgoTypeDecred:
		return "Decred";
	case AlgoTypeCryptoNight:
		return "CryptoNight";
	case AlgoTypeLbry:
		return "Lbry";
	case AlgoTypeEquihash:
		return "Equihash";
	}
	return "NA"
}

type Location int

const (
	LocationNiceHash Location = iota
	LocationWestHash
)

func (t Location) ToString() string {
	switch t {
	case LocationNiceHash:
		return "NiceHash";
	case LocationWestHash:
		return "WestHash";
	}
	return "NA"
}

type OrderType int

const (
	OrderTypeStandard OrderType = iota
	OrderTypeFixed
)

func (t OrderType) ToString() string {
	switch t {
	case OrderTypeStandard:
		return "Standard"
	case OrderTypeFixed:
		return "Fixed"
	}
	return "NA"
}
