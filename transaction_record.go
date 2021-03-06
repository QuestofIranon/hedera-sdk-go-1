package hedera

import (
	"fmt"
	"time"

	"github.com/hashgraph/hedera-sdk-go/proto"
)

type TransactionRecord struct {
	Receipt            TransactionReceipt
	Hash               []byte
	ConsensusTimestamp time.Time
	TransactionID      TransactionID
	TransactionMemo    string
	TransactionFee     uint64
	Transfers          []Transfer
	callResult         *ContractFunctionResult
	callResultIsCreate bool
}

func (record TransactionRecord) ContractExecuteResult() (ContractFunctionResult, error) {
	if record.callResult == nil || record.callResultIsCreate {
		return ContractFunctionResult{}, fmt.Errorf("record does not contain a contract execute result")
	}

	return *record.callResult, nil
}

func (record TransactionRecord) ContractCreateResult() (ContractFunctionResult, error) {
	if record.callResult == nil || !record.callResultIsCreate {
		return ContractFunctionResult{}, fmt.Errorf("record does not contain a contract create result")
	}

	return *record.callResult, nil
}

func transactionRecordFromProto(pb *proto.TransactionRecord) TransactionRecord {
	var transferList = []Transfer{}

	for _, element := range pb.TransferList.AccountAmounts {
		transferList = append(transferList, transferFromProto(element))
	}

	callResultIsCreate := false
	var callResult *ContractFunctionResult = nil
	if pb.GetContractCreateResult() != nil {
		callResultIsCreate = false
		result := contractFunctionResultFromProto(pb.GetContractCreateResult())
		callResult = &result
	} else {
		result := contractFunctionResultFromProto(pb.GetContractCallResult())
		callResult = &result
	}

	return TransactionRecord{
		Receipt:            transactionReceiptFromProto(pb.Receipt),
		Hash:               pb.TransactionHash,
		ConsensusTimestamp: timeFromProto(pb.ConsensusTimestamp),
		TransactionID:      transactionIDFromProto(pb.TransactionID),
		TransactionMemo:    pb.Memo,
		TransactionFee:     pb.TransactionFee,
		Transfers:          transferList,
		callResultIsCreate: callResultIsCreate,
		callResult:         callResult,
	}
}
