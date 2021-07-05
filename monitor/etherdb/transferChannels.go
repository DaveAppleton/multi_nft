package etherdb

import "fmt"

// TransferChannel is a channel returning all the latest transfers
type MultiTransferChannel struct {
	ID    uint64
	Ch    chan MultiTokenTransfer
	Error error
}

// MakeTokenTransferChannel returns a channel to token transfers
func MakeTokenTransferChannel(ID uint64) (ch MultiTransferChannel) {
	ch = MultiTransferChannel{ID: ID, Ch: make(chan MultiTokenTransfer, 50)}
	return
}

// ReadTokenTransfers send db to channel
func (ch *MultiTransferChannel) ReadTokenTransfers() (err error) {
	defer close(ch.Ch)
	tt := MultiTokenTransfer{TokenID: ch.ID}
	ttRes, err := tt.Find()
	if err != nil {
		ch.Error = err
		return
	}
	for _, transfer := range ttRes {
		fmt.Println(transfer.BlockNumber)
		ch.Ch <- transfer
	}
	return
}

// NamedTransferChannel is a channel returning all the latest transfers
type UniqueTransferChannel struct {
	ID    uint64
	Ch    chan UniqueTokenTransfer
	Error error
}

// MakeTokenTransferChannel returns a channel to token transfers
func MakeUniqueTokenTransferChannel(ID uint64) (ch UniqueTransferChannel) {
	ch = UniqueTransferChannel{ID: ID, Ch: make(chan UniqueTokenTransfer, 50)}
	return
}

// ReadTokenTransfers send db to channel
func (ch *UniqueTransferChannel) ReadUniqueTokenTransfers(addr string) (err error) {
	defer close(ch.Ch)
	tt := UniqueTokenTransfer{TokenID: ch.ID}
	ttRes, err := tt.FindAllByAddress(addr)
	if err != nil {
		ch.Error = err
		return
	}
	fmt.Println(len(ttRes), " results found for ", addr)
	for _, transfer := range ttRes {
		ch.Ch <- transfer
	}
	return
}
