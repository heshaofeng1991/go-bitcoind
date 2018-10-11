package bitcoind

// A ScriptSig represents a scriptsyg
type ScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

// Vin represent an IN value
type Vin struct {
	Coinbase  string    `json:"coinbase"`
	Txid      string    `json:"txid"`
	Vout      int       `json:"vout"`
	ScriptSig ScriptSig `json:"scriptSig"`
	Sequence  uint32    `json:"sequence"`
}

type ScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs,omitempty"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses,omitempty"`
}

// Vout represent an OUT value
type Vout struct {
	Value        float64      `json:"value"`
	N            int          `json:"n"`
	ScriptPubKey ScriptPubKey `json:"scriptPubKey"`
}

type RawTransaction struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		Sequence int64 `json:"sequence"`
	} `json:"vin"`
	Vout []struct {
		Value        float64 `json:"value"`
		N            int     `json:"n"`
		ScriptPubKey struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
	} `json:"vout"`
	Hex           string `json:"hex"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
}

// TransactionDetails represents details about a transaction
type TransactionDetails struct {
	Account  string  `json:"account"`
	Address  string  `json:"address,omitempty"`
	Category string  `json:"category"`
	Amount   float64 `json:"amount"`
	Fee      float64 `json:"fee,omitempty"`
}

// Transaction represents a transaction
type Transaction struct {
	Amount            float64       `json:"amount"`
	Fee               float64       `json:"fee"`
	Confirmations     int64         `json:"confirmations"`
	Blockhash         string        `json:"blockhash"`
	Blockindex        int64         `json:"blockindex"`
	Blocktime         int64         `json:"blocktime"`
	Txid              string        `json:"txid"`
	Walletconflicts   []interface{} `json:"walletconflicts"`
	Time              int64         `json:"time"`
	Timereceived      int64         `json:"timereceived"`
	Bip125Replaceable string        `json:"bip125-replaceable"`
	Details           []struct {
		Account   string  `json:"account"`
		Address   string  `json:"address"`
		Category  string  `json:"category"`
		Amount    float64 `json:"amount"`
		Label     string  `json:"label"`
		Vout      int64   `json:"vout"`
		Fee       float64 `json:"fee,omitempty"`
		Abandoned bool    `json:"abandoned,omitempty"`
	} `json:"details"`
	Hex string `json:"hex"`
}

//type Transaction struct {
//	Amount          float64              `json:"amount"`
//	Account         string               `json:"account,omitempty"`
//	Address         string               `json:"address,omitempty"`
//	Category        string               `json:"category,omitempty"`
//	Fee             float64              `json:"fee,omitempty"`
//	Confirmations   int64                `json:"confirmations"`
//	BlockHash       string               `json:"blockhash"`
//	BlockIndex      int64                `json:"blockindex"`
//	BlockTime       int64                `json:"blocktime"`
//	TxID            string               `json:"txid"`
//	WalletConflicts []string             `json:"walletconflicts"`
//	Time            int64                `json:"time"`
//	TimeReceived    int64                `json:"timereceived"`
//	Details         []TransactionDetails `json:"details,omitempty"`
//	Hex             string               `json:"hex,omitempty"`
//}

// UTransactionOut represents a unspent transaction out (UTXO)
type UTransactionOut struct {
	Bestblock     string       `json:"bestblock"`
	Confirmations uint32       `json:"confirmations"`
	Value         float64      `json:"value"`
	ScriptPubKey  ScriptPubKey `json:"scriptPubKey"`
	Version       uint32       `json:"version"`
	Coinbase      bool         `json:"coinbase"`
}

// TransactionOutSet represents statistics about the unspent transaction output database
type TransactionOutSet struct {
	Height          uint32  `json:"height"`
	Bestblock       string  `json:"bestblock"`
	Transactions    float64 `json:"transactions"`
	TxOuts          float64 `json:"txouts"`
	BytesSerialized float64 `json:"bytes_serialized"`
	HashSerialized  string  `json:"hash_serialized"`
	TotalAmount     float64 `json:"total_amount"`
}
