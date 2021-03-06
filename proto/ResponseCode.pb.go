// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ResponseCode.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResponseCodeEnum int32

const (
	ResponseCodeEnum_OK                                       ResponseCodeEnum = 0
	ResponseCodeEnum_INVALID_TRANSACTION                      ResponseCodeEnum = 1
	ResponseCodeEnum_PAYER_ACCOUNT_NOT_FOUND                  ResponseCodeEnum = 2
	ResponseCodeEnum_INVALID_NODE_ACCOUNT                     ResponseCodeEnum = 3
	ResponseCodeEnum_TRANSACTION_EXPIRED                      ResponseCodeEnum = 4
	ResponseCodeEnum_INVALID_TRANSACTION_START                ResponseCodeEnum = 5
	ResponseCodeEnum_INVALID_TRANSACTION_DURATION             ResponseCodeEnum = 6
	ResponseCodeEnum_INVALID_SIGNATURE                        ResponseCodeEnum = 7
	ResponseCodeEnum_MEMO_TOO_LONG                            ResponseCodeEnum = 8
	ResponseCodeEnum_INSUFFICIENT_TX_FEE                      ResponseCodeEnum = 9
	ResponseCodeEnum_INSUFFICIENT_PAYER_BALANCE               ResponseCodeEnum = 10
	ResponseCodeEnum_DUPLICATE_TRANSACTION                    ResponseCodeEnum = 11
	ResponseCodeEnum_BUSY                                     ResponseCodeEnum = 12
	ResponseCodeEnum_NOT_SUPPORTED                            ResponseCodeEnum = 13
	ResponseCodeEnum_INVALID_FILE_ID                          ResponseCodeEnum = 14
	ResponseCodeEnum_INVALID_ACCOUNT_ID                       ResponseCodeEnum = 15
	ResponseCodeEnum_INVALID_CONTRACT_ID                      ResponseCodeEnum = 16
	ResponseCodeEnum_INVALID_TRANSACTION_ID                   ResponseCodeEnum = 17
	ResponseCodeEnum_RECEIPT_NOT_FOUND                        ResponseCodeEnum = 18
	ResponseCodeEnum_RECORD_NOT_FOUND                         ResponseCodeEnum = 19
	ResponseCodeEnum_INVALID_SOLIDITY_ID                      ResponseCodeEnum = 20
	ResponseCodeEnum_UNKNOWN                                  ResponseCodeEnum = 21
	ResponseCodeEnum_SUCCESS                                  ResponseCodeEnum = 22
	ResponseCodeEnum_FAIL_INVALID                             ResponseCodeEnum = 23
	ResponseCodeEnum_FAIL_FEE                                 ResponseCodeEnum = 24
	ResponseCodeEnum_FAIL_BALANCE                             ResponseCodeEnum = 25
	ResponseCodeEnum_KEY_REQUIRED                             ResponseCodeEnum = 26
	ResponseCodeEnum_BAD_ENCODING                             ResponseCodeEnum = 27
	ResponseCodeEnum_INSUFFICIENT_ACCOUNT_BALANCE             ResponseCodeEnum = 28
	ResponseCodeEnum_INVALID_SOLIDITY_ADDRESS                 ResponseCodeEnum = 29
	ResponseCodeEnum_INSUFFICIENT_GAS                         ResponseCodeEnum = 30
	ResponseCodeEnum_CONTRACT_SIZE_LIMIT_EXCEEDED             ResponseCodeEnum = 31
	ResponseCodeEnum_LOCAL_CALL_MODIFICATION_EXCEPTION        ResponseCodeEnum = 32
	ResponseCodeEnum_CONTRACT_REVERT_EXECUTED                 ResponseCodeEnum = 33
	ResponseCodeEnum_CONTRACT_EXECUTION_EXCEPTION             ResponseCodeEnum = 34
	ResponseCodeEnum_INVALID_RECEIVING_NODE_ACCOUNT           ResponseCodeEnum = 35
	ResponseCodeEnum_MISSING_QUERY_HEADER                     ResponseCodeEnum = 36
	ResponseCodeEnum_ACCOUNT_UPDATE_FAILED                    ResponseCodeEnum = 37
	ResponseCodeEnum_INVALID_KEY_ENCODING                     ResponseCodeEnum = 38
	ResponseCodeEnum_NULL_SOLIDITY_ADDRESS                    ResponseCodeEnum = 39
	ResponseCodeEnum_CONTRACT_UPDATE_FAILED                   ResponseCodeEnum = 40
	ResponseCodeEnum_INVALID_QUERY_HEADER                     ResponseCodeEnum = 41
	ResponseCodeEnum_INVALID_FEE_SUBMITTED                    ResponseCodeEnum = 42
	ResponseCodeEnum_INVALID_PAYER_SIGNATURE                  ResponseCodeEnum = 43
	ResponseCodeEnum_KEY_NOT_PROVIDED                         ResponseCodeEnum = 44
	ResponseCodeEnum_INVALID_EXPIRATION_TIME                  ResponseCodeEnum = 45
	ResponseCodeEnum_NO_WACL_KEY                              ResponseCodeEnum = 46
	ResponseCodeEnum_FILE_CONTENT_EMPTY                       ResponseCodeEnum = 47
	ResponseCodeEnum_INVALID_ACCOUNT_AMOUNTS                  ResponseCodeEnum = 48
	ResponseCodeEnum_EMPTY_TRANSACTION_BODY                   ResponseCodeEnum = 49
	ResponseCodeEnum_INVALID_TRANSACTION_BODY                 ResponseCodeEnum = 50
	ResponseCodeEnum_INVALID_SIGNATURE_TYPE_MISMATCHING_KEY   ResponseCodeEnum = 51
	ResponseCodeEnum_INVALID_SIGNATURE_COUNT_MISMATCHING_KEY  ResponseCodeEnum = 52
	ResponseCodeEnum_EMPTY_CLAIM_BODY                         ResponseCodeEnum = 53
	ResponseCodeEnum_EMPTY_CLAIM_HASH                         ResponseCodeEnum = 54
	ResponseCodeEnum_EMPTY_CLAIM_KEYS                         ResponseCodeEnum = 55
	ResponseCodeEnum_INVALID_CLAIM_HASH_SIZE                  ResponseCodeEnum = 56
	ResponseCodeEnum_EMPTY_QUERY_BODY                         ResponseCodeEnum = 57
	ResponseCodeEnum_EMPTY_CLAIM_QUERY                        ResponseCodeEnum = 58
	ResponseCodeEnum_CLAIM_NOT_FOUND                          ResponseCodeEnum = 59
	ResponseCodeEnum_ACCOUNT_ID_DOES_NOT_EXIST                ResponseCodeEnum = 60
	ResponseCodeEnum_CLAIM_ALREADY_EXISTS                     ResponseCodeEnum = 61
	ResponseCodeEnum_INVALID_FILE_WACL                        ResponseCodeEnum = 62
	ResponseCodeEnum_SERIALIZATION_FAILED                     ResponseCodeEnum = 63
	ResponseCodeEnum_TRANSACTION_OVERSIZE                     ResponseCodeEnum = 64
	ResponseCodeEnum_TRANSACTION_TOO_MANY_LAYERS              ResponseCodeEnum = 65
	ResponseCodeEnum_CONTRACT_DELETED                         ResponseCodeEnum = 66
	ResponseCodeEnum_PLATFORM_NOT_ACTIVE                      ResponseCodeEnum = 67
	ResponseCodeEnum_KEY_PREFIX_MISMATCH                      ResponseCodeEnum = 68
	ResponseCodeEnum_PLATFORM_TRANSACTION_NOT_CREATED         ResponseCodeEnum = 69
	ResponseCodeEnum_INVALID_RENEWAL_PERIOD                   ResponseCodeEnum = 70
	ResponseCodeEnum_INVALID_PAYER_ACCOUNT_ID                 ResponseCodeEnum = 71
	ResponseCodeEnum_ACCOUNT_DELETED                          ResponseCodeEnum = 72
	ResponseCodeEnum_FILE_DELETED                             ResponseCodeEnum = 73
	ResponseCodeEnum_ACCOUNT_REPEATED_IN_ACCOUNT_AMOUNTS      ResponseCodeEnum = 74
	ResponseCodeEnum_SETTING_NEGATIVE_ACCOUNT_BALANCE         ResponseCodeEnum = 75
	ResponseCodeEnum_OBTAINER_REQUIRED                        ResponseCodeEnum = 76
	ResponseCodeEnum_OBTAINER_SAME_CONTRACT_ID                ResponseCodeEnum = 77
	ResponseCodeEnum_OBTAINER_DOES_NOT_EXIST                  ResponseCodeEnum = 78
	ResponseCodeEnum_MODIFYING_IMMUTABLE_CONTRACT             ResponseCodeEnum = 79
	ResponseCodeEnum_FILE_SYSTEM_EXCEPTION                    ResponseCodeEnum = 80
	ResponseCodeEnum_AUTORENEW_DURATION_NOT_IN_RANGE          ResponseCodeEnum = 81
	ResponseCodeEnum_ERROR_DECODING_BYTESTRING                ResponseCodeEnum = 82
	ResponseCodeEnum_CONTRACT_FILE_EMPTY                      ResponseCodeEnum = 83
	ResponseCodeEnum_CONTRACT_BYTECODE_EMPTY                  ResponseCodeEnum = 84
	ResponseCodeEnum_INVALID_INITIAL_BALANCE                  ResponseCodeEnum = 85
	ResponseCodeEnum_INVALID_RECEIVE_RECORD_THRESHOLD         ResponseCodeEnum = 86
	ResponseCodeEnum_INVALID_SEND_RECORD_THRESHOLD            ResponseCodeEnum = 87
	ResponseCodeEnum_ACCOUNT_IS_NOT_GENESIS_ACCOUNT           ResponseCodeEnum = 88
	ResponseCodeEnum_PAYER_ACCOUNT_UNAUTHORIZED               ResponseCodeEnum = 89
	ResponseCodeEnum_INVALID_FREEZE_TRANSACTION_BODY          ResponseCodeEnum = 90
	ResponseCodeEnum_FREEZE_TRANSACTION_BODY_NOT_FOUND        ResponseCodeEnum = 91
	ResponseCodeEnum_TRANSFER_LIST_SIZE_LIMIT_EXCEEDED        ResponseCodeEnum = 92
	ResponseCodeEnum_RESULT_SIZE_LIMIT_EXCEEDED               ResponseCodeEnum = 93
	ResponseCodeEnum_NOT_SPECIAL_ACCOUNT                      ResponseCodeEnum = 94
	ResponseCodeEnum_CONTRACT_NEGATIVE_GAS                    ResponseCodeEnum = 95
	ResponseCodeEnum_CONTRACT_NEGATIVE_VALUE                  ResponseCodeEnum = 96
	ResponseCodeEnum_INVALID_FEE_FILE                         ResponseCodeEnum = 97
	ResponseCodeEnum_INVALID_EXCHANGE_RATE_FILE               ResponseCodeEnum = 98
	ResponseCodeEnum_INSUFFICIENT_LOCAL_CALL_GAS              ResponseCodeEnum = 99
	ResponseCodeEnum_ENTITY_NOT_ALLOWED_TO_DELETE             ResponseCodeEnum = 100
	ResponseCodeEnum_AUTHORIZATION_FAILED                     ResponseCodeEnum = 101
	ResponseCodeEnum_FILE_UPLOADED_PROTO_INVALID              ResponseCodeEnum = 102
	ResponseCodeEnum_FILE_UPLOADED_PROTO_NOT_SAVED_TO_DISK    ResponseCodeEnum = 103
	ResponseCodeEnum_FEE_SCHEDULE_FILE_PART_UPLOADED          ResponseCodeEnum = 104
	ResponseCodeEnum_EXCHANGE_RATE_CHANGE_LIMIT_EXCEEDED      ResponseCodeEnum = 105
	ResponseCodeEnum_MAX_CONTRACT_STORAGE_EXCEEDED            ResponseCodeEnum = 106
	ResponseCodeEnum_TRANSAFER_ACCOUNT_SAME_AS_DELETE_ACCOUNT ResponseCodeEnum = 107
	ResponseCodeEnum_TOTAL_LEDGER_BALANCE_INVALID             ResponseCodeEnum = 108
	ResponseCodeEnum_EXPIRATION_REDUCTION_NOT_ALLOWED         ResponseCodeEnum = 110
	ResponseCodeEnum_MAX_GAS_LIMIT_EXCEEDED                   ResponseCodeEnum = 111
	ResponseCodeEnum_MAX_FILE_SIZE_EXCEEDED                   ResponseCodeEnum = 112
)

var ResponseCodeEnum_name = map[int32]string{
	0:   "OK",
	1:   "INVALID_TRANSACTION",
	2:   "PAYER_ACCOUNT_NOT_FOUND",
	3:   "INVALID_NODE_ACCOUNT",
	4:   "TRANSACTION_EXPIRED",
	5:   "INVALID_TRANSACTION_START",
	6:   "INVALID_TRANSACTION_DURATION",
	7:   "INVALID_SIGNATURE",
	8:   "MEMO_TOO_LONG",
	9:   "INSUFFICIENT_TX_FEE",
	10:  "INSUFFICIENT_PAYER_BALANCE",
	11:  "DUPLICATE_TRANSACTION",
	12:  "BUSY",
	13:  "NOT_SUPPORTED",
	14:  "INVALID_FILE_ID",
	15:  "INVALID_ACCOUNT_ID",
	16:  "INVALID_CONTRACT_ID",
	17:  "INVALID_TRANSACTION_ID",
	18:  "RECEIPT_NOT_FOUND",
	19:  "RECORD_NOT_FOUND",
	20:  "INVALID_SOLIDITY_ID",
	21:  "UNKNOWN",
	22:  "SUCCESS",
	23:  "FAIL_INVALID",
	24:  "FAIL_FEE",
	25:  "FAIL_BALANCE",
	26:  "KEY_REQUIRED",
	27:  "BAD_ENCODING",
	28:  "INSUFFICIENT_ACCOUNT_BALANCE",
	29:  "INVALID_SOLIDITY_ADDRESS",
	30:  "INSUFFICIENT_GAS",
	31:  "CONTRACT_SIZE_LIMIT_EXCEEDED",
	32:  "LOCAL_CALL_MODIFICATION_EXCEPTION",
	33:  "CONTRACT_REVERT_EXECUTED",
	34:  "CONTRACT_EXECUTION_EXCEPTION",
	35:  "INVALID_RECEIVING_NODE_ACCOUNT",
	36:  "MISSING_QUERY_HEADER",
	37:  "ACCOUNT_UPDATE_FAILED",
	38:  "INVALID_KEY_ENCODING",
	39:  "NULL_SOLIDITY_ADDRESS",
	40:  "CONTRACT_UPDATE_FAILED",
	41:  "INVALID_QUERY_HEADER",
	42:  "INVALID_FEE_SUBMITTED",
	43:  "INVALID_PAYER_SIGNATURE",
	44:  "KEY_NOT_PROVIDED",
	45:  "INVALID_EXPIRATION_TIME",
	46:  "NO_WACL_KEY",
	47:  "FILE_CONTENT_EMPTY",
	48:  "INVALID_ACCOUNT_AMOUNTS",
	49:  "EMPTY_TRANSACTION_BODY",
	50:  "INVALID_TRANSACTION_BODY",
	51:  "INVALID_SIGNATURE_TYPE_MISMATCHING_KEY",
	52:  "INVALID_SIGNATURE_COUNT_MISMATCHING_KEY",
	53:  "EMPTY_CLAIM_BODY",
	54:  "EMPTY_CLAIM_HASH",
	55:  "EMPTY_CLAIM_KEYS",
	56:  "INVALID_CLAIM_HASH_SIZE",
	57:  "EMPTY_QUERY_BODY",
	58:  "EMPTY_CLAIM_QUERY",
	59:  "CLAIM_NOT_FOUND",
	60:  "ACCOUNT_ID_DOES_NOT_EXIST",
	61:  "CLAIM_ALREADY_EXISTS",
	62:  "INVALID_FILE_WACL",
	63:  "SERIALIZATION_FAILED",
	64:  "TRANSACTION_OVERSIZE",
	65:  "TRANSACTION_TOO_MANY_LAYERS",
	66:  "CONTRACT_DELETED",
	67:  "PLATFORM_NOT_ACTIVE",
	68:  "KEY_PREFIX_MISMATCH",
	69:  "PLATFORM_TRANSACTION_NOT_CREATED",
	70:  "INVALID_RENEWAL_PERIOD",
	71:  "INVALID_PAYER_ACCOUNT_ID",
	72:  "ACCOUNT_DELETED",
	73:  "FILE_DELETED",
	74:  "ACCOUNT_REPEATED_IN_ACCOUNT_AMOUNTS",
	75:  "SETTING_NEGATIVE_ACCOUNT_BALANCE",
	76:  "OBTAINER_REQUIRED",
	77:  "OBTAINER_SAME_CONTRACT_ID",
	78:  "OBTAINER_DOES_NOT_EXIST",
	79:  "MODIFYING_IMMUTABLE_CONTRACT",
	80:  "FILE_SYSTEM_EXCEPTION",
	81:  "AUTORENEW_DURATION_NOT_IN_RANGE",
	82:  "ERROR_DECODING_BYTESTRING",
	83:  "CONTRACT_FILE_EMPTY",
	84:  "CONTRACT_BYTECODE_EMPTY",
	85:  "INVALID_INITIAL_BALANCE",
	86:  "INVALID_RECEIVE_RECORD_THRESHOLD",
	87:  "INVALID_SEND_RECORD_THRESHOLD",
	88:  "ACCOUNT_IS_NOT_GENESIS_ACCOUNT",
	89:  "PAYER_ACCOUNT_UNAUTHORIZED",
	90:  "INVALID_FREEZE_TRANSACTION_BODY",
	91:  "FREEZE_TRANSACTION_BODY_NOT_FOUND",
	92:  "TRANSFER_LIST_SIZE_LIMIT_EXCEEDED",
	93:  "RESULT_SIZE_LIMIT_EXCEEDED",
	94:  "NOT_SPECIAL_ACCOUNT",
	95:  "CONTRACT_NEGATIVE_GAS",
	96:  "CONTRACT_NEGATIVE_VALUE",
	97:  "INVALID_FEE_FILE",
	98:  "INVALID_EXCHANGE_RATE_FILE",
	99:  "INSUFFICIENT_LOCAL_CALL_GAS",
	100: "ENTITY_NOT_ALLOWED_TO_DELETE",
	101: "AUTHORIZATION_FAILED",
	102: "FILE_UPLOADED_PROTO_INVALID",
	103: "FILE_UPLOADED_PROTO_NOT_SAVED_TO_DISK",
	104: "FEE_SCHEDULE_FILE_PART_UPLOADED",
	105: "EXCHANGE_RATE_CHANGE_LIMIT_EXCEEDED",
	106: "MAX_CONTRACT_STORAGE_EXCEEDED",
	107: "TRANSAFER_ACCOUNT_SAME_AS_DELETE_ACCOUNT",
	108: "TOTAL_LEDGER_BALANCE_INVALID",
	110: "EXPIRATION_REDUCTION_NOT_ALLOWED",
	111: "MAX_GAS_LIMIT_EXCEEDED",
	112: "MAX_FILE_SIZE_EXCEEDED",
}

var ResponseCodeEnum_value = map[string]int32{
	"OK":                                       0,
	"INVALID_TRANSACTION":                      1,
	"PAYER_ACCOUNT_NOT_FOUND":                  2,
	"INVALID_NODE_ACCOUNT":                     3,
	"TRANSACTION_EXPIRED":                      4,
	"INVALID_TRANSACTION_START":                5,
	"INVALID_TRANSACTION_DURATION":             6,
	"INVALID_SIGNATURE":                        7,
	"MEMO_TOO_LONG":                            8,
	"INSUFFICIENT_TX_FEE":                      9,
	"INSUFFICIENT_PAYER_BALANCE":               10,
	"DUPLICATE_TRANSACTION":                    11,
	"BUSY":                                     12,
	"NOT_SUPPORTED":                            13,
	"INVALID_FILE_ID":                          14,
	"INVALID_ACCOUNT_ID":                       15,
	"INVALID_CONTRACT_ID":                      16,
	"INVALID_TRANSACTION_ID":                   17,
	"RECEIPT_NOT_FOUND":                        18,
	"RECORD_NOT_FOUND":                         19,
	"INVALID_SOLIDITY_ID":                      20,
	"UNKNOWN":                                  21,
	"SUCCESS":                                  22,
	"FAIL_INVALID":                             23,
	"FAIL_FEE":                                 24,
	"FAIL_BALANCE":                             25,
	"KEY_REQUIRED":                             26,
	"BAD_ENCODING":                             27,
	"INSUFFICIENT_ACCOUNT_BALANCE":             28,
	"INVALID_SOLIDITY_ADDRESS":                 29,
	"INSUFFICIENT_GAS":                         30,
	"CONTRACT_SIZE_LIMIT_EXCEEDED":             31,
	"LOCAL_CALL_MODIFICATION_EXCEPTION":        32,
	"CONTRACT_REVERT_EXECUTED":                 33,
	"CONTRACT_EXECUTION_EXCEPTION":             34,
	"INVALID_RECEIVING_NODE_ACCOUNT":           35,
	"MISSING_QUERY_HEADER":                     36,
	"ACCOUNT_UPDATE_FAILED":                    37,
	"INVALID_KEY_ENCODING":                     38,
	"NULL_SOLIDITY_ADDRESS":                    39,
	"CONTRACT_UPDATE_FAILED":                   40,
	"INVALID_QUERY_HEADER":                     41,
	"INVALID_FEE_SUBMITTED":                    42,
	"INVALID_PAYER_SIGNATURE":                  43,
	"KEY_NOT_PROVIDED":                         44,
	"INVALID_EXPIRATION_TIME":                  45,
	"NO_WACL_KEY":                              46,
	"FILE_CONTENT_EMPTY":                       47,
	"INVALID_ACCOUNT_AMOUNTS":                  48,
	"EMPTY_TRANSACTION_BODY":                   49,
	"INVALID_TRANSACTION_BODY":                 50,
	"INVALID_SIGNATURE_TYPE_MISMATCHING_KEY":   51,
	"INVALID_SIGNATURE_COUNT_MISMATCHING_KEY":  52,
	"EMPTY_CLAIM_BODY":                         53,
	"EMPTY_CLAIM_HASH":                         54,
	"EMPTY_CLAIM_KEYS":                         55,
	"INVALID_CLAIM_HASH_SIZE":                  56,
	"EMPTY_QUERY_BODY":                         57,
	"EMPTY_CLAIM_QUERY":                        58,
	"CLAIM_NOT_FOUND":                          59,
	"ACCOUNT_ID_DOES_NOT_EXIST":                60,
	"CLAIM_ALREADY_EXISTS":                     61,
	"INVALID_FILE_WACL":                        62,
	"SERIALIZATION_FAILED":                     63,
	"TRANSACTION_OVERSIZE":                     64,
	"TRANSACTION_TOO_MANY_LAYERS":              65,
	"CONTRACT_DELETED":                         66,
	"PLATFORM_NOT_ACTIVE":                      67,
	"KEY_PREFIX_MISMATCH":                      68,
	"PLATFORM_TRANSACTION_NOT_CREATED":         69,
	"INVALID_RENEWAL_PERIOD":                   70,
	"INVALID_PAYER_ACCOUNT_ID":                 71,
	"ACCOUNT_DELETED":                          72,
	"FILE_DELETED":                             73,
	"ACCOUNT_REPEATED_IN_ACCOUNT_AMOUNTS":      74,
	"SETTING_NEGATIVE_ACCOUNT_BALANCE":         75,
	"OBTAINER_REQUIRED":                        76,
	"OBTAINER_SAME_CONTRACT_ID":                77,
	"OBTAINER_DOES_NOT_EXIST":                  78,
	"MODIFYING_IMMUTABLE_CONTRACT":             79,
	"FILE_SYSTEM_EXCEPTION":                    80,
	"AUTORENEW_DURATION_NOT_IN_RANGE":          81,
	"ERROR_DECODING_BYTESTRING":                82,
	"CONTRACT_FILE_EMPTY":                      83,
	"CONTRACT_BYTECODE_EMPTY":                  84,
	"INVALID_INITIAL_BALANCE":                  85,
	"INVALID_RECEIVE_RECORD_THRESHOLD":         86,
	"INVALID_SEND_RECORD_THRESHOLD":            87,
	"ACCOUNT_IS_NOT_GENESIS_ACCOUNT":           88,
	"PAYER_ACCOUNT_UNAUTHORIZED":               89,
	"INVALID_FREEZE_TRANSACTION_BODY":          90,
	"FREEZE_TRANSACTION_BODY_NOT_FOUND":        91,
	"TRANSFER_LIST_SIZE_LIMIT_EXCEEDED":        92,
	"RESULT_SIZE_LIMIT_EXCEEDED":               93,
	"NOT_SPECIAL_ACCOUNT":                      94,
	"CONTRACT_NEGATIVE_GAS":                    95,
	"CONTRACT_NEGATIVE_VALUE":                  96,
	"INVALID_FEE_FILE":                         97,
	"INVALID_EXCHANGE_RATE_FILE":               98,
	"INSUFFICIENT_LOCAL_CALL_GAS":              99,
	"ENTITY_NOT_ALLOWED_TO_DELETE":             100,
	"AUTHORIZATION_FAILED":                     101,
	"FILE_UPLOADED_PROTO_INVALID":              102,
	"FILE_UPLOADED_PROTO_NOT_SAVED_TO_DISK":    103,
	"FEE_SCHEDULE_FILE_PART_UPLOADED":          104,
	"EXCHANGE_RATE_CHANGE_LIMIT_EXCEEDED":      105,
	"MAX_CONTRACT_STORAGE_EXCEEDED":            106,
	"TRANSAFER_ACCOUNT_SAME_AS_DELETE_ACCOUNT": 107,
	"TOTAL_LEDGER_BALANCE_INVALID":             108,
	"EXPIRATION_REDUCTION_NOT_ALLOWED":         110,
	"MAX_GAS_LIMIT_EXCEEDED":                   111,
	"MAX_FILE_SIZE_EXCEEDED":                   112,
}

func (x ResponseCodeEnum) String() string {
	return proto.EnumName(ResponseCodeEnum_name, int32(x))
}

func (ResponseCodeEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1af33712ce1d0093, []int{0}
}

func init() {
	proto.RegisterEnum("proto.ResponseCodeEnum", ResponseCodeEnum_name, ResponseCodeEnum_value)
}

func init() { proto.RegisterFile("ResponseCode.proto", fileDescriptor_1af33712ce1d0093) }

var fileDescriptor_1af33712ce1d0093 = []byte{
	// 1401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x57, 0x79, 0x7f, 0xd4, 0x36,
	0x10, 0x2d, 0x94, 0xab, 0x02, 0xca, 0x20, 0xae, 0x10, 0x20, 0xdc, 0x37, 0x4d, 0x0f, 0x7a, 0xdf,
	0x5a, 0x6b, 0x76, 0x57, 0x8d, 0x2c, 0x19, 0x49, 0xde, 0xac, 0xd3, 0x43, 0x0d, 0x90, 0x92, 0x1e,
	0x90, 0xfc, 0xa0, 0xfd, 0xae, 0xfd, 0x38, 0xfd, 0x8d, 0xbc, 0xf6, 0xda, 0xd9, 0xf4, 0xaf, 0x24,
	0x1a, 0xcd, 0x68, 0xe6, 0xcd, 0x9b, 0x37, 0x0e, 0xe3, 0x6e, 0xeb, 0xcd, 0xee, 0xce, 0xab, 0x37,
	0x5b, 0xd9, 0xce, 0xf3, 0xad, 0xd5, 0xdd, 0xd7, 0x3b, 0x7f, 0xef, 0xf0, 0xc3, 0xe9, 0xc7, 0x83,
	0x7f, 0x97, 0x18, 0x74, 0xad, 0xf8, 0xea, 0x9f, 0x97, 0xfc, 0x08, 0x3b, 0x68, 0xd7, 0xe0, 0x2d,
	0x7e, 0x81, 0x9d, 0x51, 0x66, 0x22, 0xb4, 0x92, 0x31, 0x38, 0x61, 0xbc, 0xc8, 0x82, 0xb2, 0x06,
	0x0e, 0xf0, 0x4b, 0xec, 0x42, 0x21, 0x2a, 0x74, 0x51, 0x64, 0x99, 0x2d, 0x4d, 0x88, 0xc6, 0x86,
	0x38, 0xb4, 0xa5, 0x91, 0x70, 0x90, 0x2f, 0xb1, 0xb3, 0x8d, 0x97, 0xb1, 0x12, 0x9b, 0x3b, 0xf0,
	0x36, 0xc5, 0xeb, 0xc4, 0x89, 0x38, 0x2d, 0x94, 0x43, 0x09, 0x87, 0xf8, 0x15, 0x76, 0x71, 0x9f,
	0x87, 0xa2, 0x0f, 0xc2, 0x05, 0x38, 0xcc, 0xaf, 0xb1, 0xcb, 0xfb, 0x99, 0x65, 0xe9, 0x44, 0x4a,
	0xe8, 0x08, 0x3f, 0xc7, 0x4e, 0x37, 0x37, 0xbc, 0x1a, 0x19, 0x11, 0x4a, 0x87, 0x70, 0x94, 0x9f,
	0x66, 0x27, 0x73, 0xcc, 0x6d, 0x0c, 0xd6, 0x46, 0x6d, 0xcd, 0x08, 0x8e, 0xd5, 0x35, 0xf9, 0x72,
	0x38, 0x54, 0x99, 0x42, 0x13, 0x62, 0x98, 0xc6, 0x21, 0x22, 0xbc, 0xc3, 0x57, 0xd8, 0x72, 0xcf,
	0x50, 0x17, 0x38, 0x10, 0x5a, 0x98, 0x0c, 0x81, 0xf1, 0x8b, 0xec, 0x9c, 0x2c, 0x0b, 0xad, 0x32,
	0x11, 0xb0, 0x07, 0xc7, 0x71, 0x7e, 0x8c, 0x1d, 0x1a, 0x94, 0xbe, 0x82, 0x13, 0xf4, 0x20, 0x41,
	0xe1, 0xcb, 0xa2, 0xb0, 0x2e, 0xa0, 0x84, 0x93, 0xfc, 0x0c, 0x3b, 0xd5, 0xa4, 0x36, 0x54, 0x1a,
	0xa3, 0x92, 0xf0, 0x2e, 0x3f, 0xcf, 0x78, 0x73, 0xd8, 0x40, 0xa8, 0x24, 0x9c, 0xea, 0x22, 0x9e,
	0x59, 0x13, 0x9c, 0xc8, 0x92, 0x01, 0xf8, 0x32, 0x3b, 0xbf, 0x1f, 0x04, 0x4a, 0xc2, 0x69, 0x2a,
	0xde, 0x61, 0x86, 0xaa, 0xe8, 0xf6, 0x81, 0xf3, 0xb3, 0x0c, 0x1c, 0x66, 0xd6, 0xc9, 0xce, 0xe9,
	0x99, 0xee, 0x0b, 0xde, 0x6a, 0x25, 0x55, 0xa8, 0x28, 0xca, 0x59, 0x7e, 0x9c, 0x1d, 0x2d, 0xcd,
	0x9a, 0xb1, 0xeb, 0x06, 0xce, 0xd1, 0x1f, 0xbe, 0xcc, 0x32, 0xf4, 0x1e, 0xce, 0x73, 0x60, 0x27,
	0x86, 0x42, 0xe9, 0x38, 0xf3, 0x83, 0x0b, 0xfc, 0x04, 0x3b, 0x96, 0x4e, 0x08, 0xb9, 0xa5, 0xd6,
	0xde, 0x60, 0x75, 0x91, 0x4e, 0xd6, 0xb0, 0x8a, 0x0e, 0x9f, 0x94, 0xa9, 0xc3, 0xcb, 0x74, 0x32,
	0x10, 0x32, 0xa2, 0xc9, 0xac, 0x54, 0x66, 0x04, 0x97, 0xea, 0xa6, 0x76, 0xf0, 0x6e, 0x70, 0x68,
	0xa2, 0x5c, 0xe6, 0x97, 0xd9, 0xd2, 0x42, 0xaa, 0x42, 0x4a, 0x47, 0x59, 0x5d, 0xa1, 0xf2, 0x7a,
	0xfe, 0x23, 0xe1, 0x61, 0x85, 0xa2, 0xb6, 0xc0, 0x79, 0xb5, 0x81, 0x51, 0xab, 0x5c, 0x85, 0x88,
	0xd3, 0x0c, 0x51, 0xa2, 0x84, 0xab, 0xfc, 0x36, 0xbb, 0xae, 0x6d, 0x26, 0x74, 0xcc, 0x84, 0xd6,
	0x31, 0xb7, 0x52, 0x0d, 0xa9, 0xa7, 0x35, 0x21, 0x33, 0x2c, 0x52, 0x4f, 0xaf, 0xd1, 0xe3, 0x6d,
	0x20, 0x87, 0x13, 0x74, 0x14, 0x04, 0xb3, 0x92, 0x9a, 0x7a, 0xbd, 0xf7, 0x4c, 0x7d, 0xdc, 0xf7,
	0xbf, 0xc1, 0x6f, 0xb0, 0x95, 0x26, 0xf9, 0xd4, 0x9c, 0x89, 0x32, 0xa3, 0xfe, 0x3c, 0xdc, 0xa4,
	0x49, 0xc9, 0x95, 0xf7, 0x64, 0x79, 0x52, 0xa2, 0xab, 0xe2, 0x18, 0x85, 0x44, 0x07, 0xb7, 0x88,
	0x6c, 0x0d, 0x1e, 0x65, 0x21, 0x89, 0x71, 0x84, 0x30, 0x4a, 0xb8, 0xdd, 0x1d, 0x2f, 0xc2, 0xb8,
	0x45, 0xf4, 0x0e, 0x39, 0x99, 0x52, 0xeb, 0x45, 0xb0, 0xee, 0x12, 0x7d, 0xda, 0x7c, 0xfb, 0x01,
	0xef, 0x75, 0x03, 0xf6, 0xb2, 0xb8, 0x4f, 0x01, 0x5b, 0xea, 0x22, 0x46, 0x5f, 0x0e, 0x72, 0x15,
	0x08, 0x80, 0x07, 0xa4, 0x00, 0x8d, 0xa9, 0x1e, 0x94, 0xf9, 0xd8, 0x3d, 0xa4, 0xd6, 0x50, 0x6a,
	0x44, 0xbb, 0xc2, 0xd9, 0x89, 0x22, 0xe0, 0x1f, 0x75, 0x5d, 0xd2, 0xe4, 0xd7, 0x98, 0x07, 0x95,
	0x23, 0xbc, 0xc7, 0x4f, 0xb1, 0xe3, 0xc6, 0xc6, 0x75, 0x91, 0x69, 0xaa, 0x0a, 0x56, 0x69, 0x42,
	0xd2, 0xb8, 0x50, 0xda, 0xd4, 0x5e, 0xcc, 0x8b, 0x50, 0xc1, 0xfb, 0xdd, 0x28, 0x0d, 0x42, 0x22,
	0xa7, 0x1f, 0x1e, 0x3e, 0xa0, 0x32, 0xd3, 0xbd, 0xde, 0x8c, 0x0c, 0xac, 0xac, 0xe0, 0xc3, 0x2e,
	0x9b, 0x16, 0xac, 0x1f, 0xf1, 0x07, 0xec, 0xce, 0x82, 0x80, 0xc4, 0x50, 0x15, 0x18, 0x73, 0xe5,
	0x73, 0x11, 0xb2, 0x31, 0x75, 0x89, 0x52, 0x7b, 0xcc, 0x1f, 0xb2, 0xbb, 0x8b, 0x77, 0xeb, 0x54,
	0xf6, 0x5e, 0xfe, 0x98, 0xb0, 0xa8, 0x53, 0xca, 0xb4, 0x50, 0x79, 0xfd, 0xdc, 0x27, 0x7b, 0x4f,
	0xc7, 0xc2, 0x8f, 0xe1, 0xd3, 0xbd, 0xa7, 0x6b, 0x58, 0x79, 0xf8, 0xac, 0x5b, 0xf1, 0xfc, 0x76,
	0x22, 0x37, 0x7c, 0x3e, 0x77, 0xa9, 0x5b, 0x97, 0xc2, 0x7f, 0x41, 0x8a, 0xd0, 0x0d, 0x94, 0x6c,
	0xf0, 0x25, 0x49, 0x51, 0x7d, 0x30, 0x17, 0x84, 0xaf, 0x48, 0x7b, 0xe7, 0x12, 0x14, 0xa5, 0x45,
	0x9f, 0xcc, 0x38, 0x55, 0x3e, 0xc0, 0xd7, 0xc4, 0x8e, 0xda, 0x47, 0x68, 0x87, 0x42, 0x56, 0xb5,
	0xc1, 0xc3, 0x37, 0x5d, 0xcd, 0x4d, 0x9d, 0xa2, 0xe6, 0xc1, 0xb7, 0xe4, 0xe0, 0xd1, 0x29, 0xa1,
	0xd5, 0x46, 0xdd, 0xe1, 0x19, 0xd1, 0xbe, 0x23, 0x4b, 0x17, 0x79, 0x3b, 0x41, 0x97, 0xaa, 0xf8,
	0x9e, 0x5f, 0x65, 0x97, 0xba, 0x16, 0x92, 0xeb, 0x5c, 0x98, 0x2a, 0x6a, 0xa2, 0x96, 0x07, 0x41,
	0x65, 0xb6, 0xfc, 0x95, 0xa8, 0x91, 0x48, 0x38, 0x20, 0x2d, 0x2b, 0xb4, 0x08, 0x43, 0xeb, 0xea,
	0x92, 0xc8, 0x7d, 0x82, 0x90, 0x91, 0x81, 0x08, 0x58, 0x38, 0x1c, 0xaa, 0x69, 0xdb, 0x14, 0x90,
	0xfc, 0x16, 0xbb, 0xd6, 0x7a, 0x74, 0x5f, 0x24, 0xef, 0xcc, 0xa1, 0xa0, 0xb8, 0xd8, 0x15, 0x5b,
	0x87, 0x06, 0xd7, 0x85, 0x8e, 0x05, 0x3a, 0x65, 0x25, 0x0c, 0xbb, 0x34, 0xea, 0xaf, 0x40, 0x25,
	0x61, 0x44, 0x08, 0x37, 0x7f, 0x37, 0x69, 0x8e, 0x93, 0x3e, 0x12, 0x40, 0xcd, 0x89, 0xe2, 0x77,
	0xd9, 0xcd, 0xe6, 0x9a, 0xc3, 0x22, 0x3d, 0x1b, 0x95, 0x59, 0x20, 0xf4, 0x0f, 0x94, 0xaf, 0xc7,
	0x10, 0x92, 0x76, 0xe0, 0x48, 0x50, 0x79, 0x0b, 0x42, 0xb9, 0x46, 0x9d, 0xb0, 0x83, 0x20, 0x94,
	0x41, 0x37, 0xd7, 0x5c, 0x4d, 0x9d, 0x6d, 0x8f, 0xbd, 0xc8, 0xb1, 0xb7, 0x52, 0x72, 0xe2, 0x55,
	0x6b, 0xde, 0xd3, 0x76, 0x43, 0x02, 0x97, 0xa4, 0xb1, 0xa2, 0xa7, 0x55, 0x9e, 0x97, 0x41, 0x0c,
	0xf4, 0x3c, 0x02, 0x58, 0x12, 0x87, 0x54, 0x95, 0xaf, 0x7c, 0xc0, 0xbc, 0xa3, 0x7d, 0x05, 0xbf,
	0xc9, 0xae, 0x8a, 0x32, 0xd8, 0x84, 0x5d, 0xbb, 0xa5, 0x53, 0x78, 0x65, 0xa2, 0x13, 0x66, 0x84,
	0xf0, 0x84, 0xb2, 0x43, 0xe7, 0xac, 0x8b, 0x12, 0x6b, 0x05, 0x8b, 0x83, 0x2a, 0xa0, 0x0f, 0x8e,
	0xc4, 0xcc, 0x51, 0x0b, 0xdb, 0x74, 0xd3, 0x3b, 0xb5, 0x00, 0x78, 0x4a, 0xbb, 0x35, 0x90, 0x47,
	0x46, 0x9a, 0x5a, 0x1b, 0x43, 0x77, 0x56, 0x94, 0x51, 0x41, 0x89, 0xf9, 0x56, 0x2a, 0x09, 0xcc,
	0xbe, 0x24, 0x63, 0x9c, 0x2d, 0xc8, 0x30, 0x76, 0xe8, 0xc7, 0x56, 0x4b, 0x98, 0xf0, 0xeb, 0xec,
	0x4a, 0x3b, 0xdd, 0x68, 0xe4, 0xe2, 0x95, 0x75, 0xd2, 0xf6, 0xb6, 0xeb, 0x35, 0x6c, 0x23, 0x34,
	0xe8, 0x95, 0x6f, 0xb5, 0x7d, 0x4a, 0x9f, 0x13, 0x7d, 0x7e, 0x94, 0x46, 0x94, 0x61, 0x6c, 0x9d,
	0xda, 0x40, 0x09, 0x15, 0x61, 0xd4, 0x4e, 0x8f, 0x43, 0xdc, 0xc0, 0x45, 0x55, 0xda, 0xa0, 0x5d,
	0xf5, 0x3f, 0xc6, 0xce, 0x08, 0xff, 0x48, 0xd7, 0x92, 0x7d, 0x88, 0x2e, 0x6a, 0xe5, 0xf7, 0xdf,
	0x7c, 0x3f, 0x51, 0x4a, 0x0e, 0x7d, 0xa9, 0xf7, 0xb7, 0xff, 0x4c, 0x90, 0xa7, 0x8f, 0x97, 0x02,
	0x33, 0x02, 0xae, 0xa9, 0xe5, 0x17, 0x6a, 0x75, 0x0b, 0x79, 0x4b, 0x43, 0xda, 0xb7, 0xb1, 0xd7,
	0x8d, 0xd6, 0x34, 0x11, 0xba, 0x44, 0xf8, 0xb5, 0x5e, 0xd1, 0xf3, 0xfd, 0x41, 0x6d, 0x84, 0xcd,
	0xfa, 0x43, 0xab, 0xd9, 0x03, 0xd9, 0x98, 0xe8, 0x10, 0x5d, 0xda, 0x48, 0x64, 0x7f, 0x4a, 0x62,
	0xd0, 0x5b, 0xec, 0x9d, 0x6d, 0x4d, 0x6f, 0x3e, 0x23, 0x6e, 0xa2, 0x09, 0xb4, 0xe0, 0xd2, 0xd0,
	0x6b, 0x6d, 0xd7, 0x51, 0xc6, 0x60, 0x67, 0x03, 0x06, 0xcf, 0x49, 0x69, 0x1a, 0xb0, 0x7b, 0x1a,
	0xb4, 0x45, 0xc1, 0x13, 0x9b, 0xca, 0x42, 0x5b, 0x21, 0x51, 0xd2, 0x82, 0x0a, 0xb6, 0xfd, 0xb4,
	0xf9, 0x8d, 0xdf, 0x67, 0xb7, 0xf7, 0xbb, 0x90, 0x80, 0x11, 0x93, 0xd9, 0x3b, 0xca, 0xaf, 0xc1,
	0x0b, 0x6a, 0x61, 0x5a, 0x8b, 0xd9, 0x18, 0x65, 0xa9, 0xeb, 0xfc, 0x63, 0x21, 0x5c, 0x68, 0x9d,
	0x61, 0x9b, 0x46, 0xbd, 0x5f, 0xe5, 0xec, 0xf7, 0x3d, 0xe8, 0xff, 0x4e, 0xbc, 0xcb, 0xc5, 0x74,
	0x3e, 0xa3, 0x3e, 0x58, 0x27, 0x46, 0x38, 0xbf, 0xf2, 0x07, 0x7f, 0xc4, 0xee, 0xd5, 0x3c, 0x18,
	0x76, 0x78, 0x95, 0x26, 0x5b, 0xf8, 0x59, 0xf1, 0x6d, 0xd7, 0xfe, 0x24, 0x98, 0x82, 0x0d, 0x42,
	0x47, 0x8d, 0x72, 0x34, 0xff, 0x94, 0x6d, 0x6b, 0xfd, 0x8b, 0x06, 0xa2, 0xb3, 0x89, 0x1d, 0xca,
	0x72, 0xae, 0x86, 0x33, 0x58, 0xe1, 0x15, 0xa9, 0x21, 0x25, 0x36, 0x12, 0x7e, 0x6f, 0xd2, 0x3b,
	0x8d, 0xad, 0x16, 0x02, 0x22, 0x55, 0x6b, 0xdb, 0x1d, 0xac, 0xb0, 0xe5, 0x67, 0x3b, 0x2f, 0x57,
	0xb7, 0xb7, 0x9e, 0x6f, 0xbd, 0xde, 0x5c, 0xdd, 0xde, 0x7c, 0xb3, 0xfd, 0xe2, 0xf5, 0xe6, 0xee,
	0x76, 0xfd, 0xff, 0x47, 0x71, 0xe0, 0xe9, 0x91, 0xf4, 0xcb, 0xe3, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x81, 0x7f, 0xf7, 0x80, 0x9e, 0x0c, 0x00, 0x00,
}
