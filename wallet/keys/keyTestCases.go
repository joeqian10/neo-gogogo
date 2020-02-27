package keys

// Ktype represents key test case values (different encodings of the key).
type Ktype struct {
	Address,
	PrivateKey,
	PublicKey,
	Wif,
	Passphrase,
	Nep2key,
	ScriptHash string
}

// Arr contains a set of known keys in Ktype format.
var KeyCases = []Ktype{
	{
		Address:    "AJntkozhVgbc6irY9hRFtNUvuPZS4YcUyD",
		PrivateKey: "831cb932167332a768f1c898d2cf4586a14aa606b7f078eba028c849c306cce6",
		PublicKey:  "03b7a7f933199f28cc1c48d22a21c78ac3992cf7fceb038a9c670fe55444426619",
		Wif:        "L1caMUAsHr2dKwhqbMpYRcCzmzvZTfYZSCBefgARhz9iimAFRn1z",
		Passphrase: "neo-gogogo",
		Nep2key:    "6PYRGTFTzMq5BTcZphryV8KEtyhGJCK7fRZr6etdAKrb2gowRvSjsSA5XZ",
		ScriptHash: "62d6035f671f46b1ab5715eef3a903910bb81921",
	},
	{
		Address:    "AHbwJGdhUy3d1BwhXQKc1VrNojjBTX6g87",
		PrivateKey: "82a4ff38f5de304c2fae62d1a736c343816412f7e4fe3badaf5e1e940c8f07c3",
		PublicKey:  "027d73c8b02e446340caceee7a517cddff72440e60c28cbb84884f307760ecad5b",
		Wif:        "L1bfdDaFQErh7gGMz32zgBXZCN65AKeexKzxeEvS7d4Cq6zf2Rpf",
		Passphrase: "",
		Nep2key:    "6PYP4RrgZTU8zcb5RN2xzeV2ufddJRe2DKom3kW8AhxuLCtFYKAt3dGZ4G",
		ScriptHash: "80176d20bd72e6f8b16332fdfa75bd63cd230f14",
	},
	{
		Address:    "AKeMGQcrN4Su5mEK7T1NgCCyBtY16vD16Y",
		PrivateKey: "31ab808b68c25377b2068500e264f164d1d75eda748a8e0a98db4c74db181b66",
		PublicKey:  "038a2151948a908cdf2d680eead6512217769e34b9db196574572cb98e273516a1",
		Wif:        "KxtGC6QHFKSiHVLY1ANkwS78ebfhworv6LnkJH2MUxE8AbbgAHVW",
		Passphrase: "密码@👍",
		Nep2key:    "6PYUKxjoTp8CmbR4ZDLLdugviWzf7QuxikGir7mudeAcSo9dviT6x6MtPB",
		ScriptHash: "6ea9062064e43b49df964d13da0a413c113b742a",
	},
	{
		Address:    "AQzRMe3zyGS8W177xLJfewRRQZY2kddMun",
		PrivateKey: "b503c7727756aaed15e2cfc1d571e00109c92b075d080038d208f8b9cab917af",
		PublicKey:  "03d08d6f766b54e35745bc99d643c939ec6f3d37004f2a59006be0e53610f0be25",
		Wif:        "L3Hab7wL43SbWLnkfnVCp6dT99xzfB4qLZxeR9dFgcWWPirwKyXp",
		Passphrase: "1",
		Nep2key:    "6PYLjXkQzADs7r36XQjByJXoggm3beRh6UzxuN59NZiBxFBnm1HPvv3ytM",
		ScriptHash: "d2601d3651a41a7faf6e02a280ab26d9dc971865",
	},
}
