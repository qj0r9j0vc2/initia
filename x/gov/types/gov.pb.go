// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: initia/gov/v1/gov.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params defines the parameters for the x/gov module.
type Params struct {
	// Minimum deposit for a proposal to enter voting period.
	MinDeposit []types.Coin `protobuf:"bytes,1,rep,name=min_deposit,json=minDeposit,proto3" json:"min_deposit"`
	// Maximum period for Atom holders to deposit on a proposal. Initial value: 2
	// months.
	MaxDepositPeriod time.Duration `protobuf:"bytes,2,opt,name=max_deposit_period,json=maxDepositPeriod,proto3,stdduration" json:"max_deposit_period"`
	// Duration of the voting period.
	VotingPeriod time.Duration `protobuf:"bytes,3,opt,name=voting_period,json=votingPeriod,proto3,stdduration" json:"voting_period"`
	//  Minimum percentage of total stake needed to vote for a result to be
	//  considered valid.
	Quorum string `protobuf:"bytes,4,opt,name=quorum,proto3" json:"quorum,omitempty"`
	//  Minimum proportion of Yes votes for proposal to pass. Default value: 0.5.
	Threshold string `protobuf:"bytes,5,opt,name=threshold,proto3" json:"threshold,omitempty"`
	//  Minimum value of Veto votes to Total votes ratio for proposal to be
	//  vetoed. Default value: 1/3.
	VetoThreshold string `protobuf:"bytes,6,opt,name=veto_threshold,json=vetoThreshold,proto3" json:"veto_threshold,omitempty"`
	//  The ratio representing the proportion of the deposit value that must be paid at proposal submission.
	MinInitialDepositRatio string `protobuf:"bytes,7,opt,name=min_initial_deposit_ratio,json=minInitialDepositRatio,proto3" json:"min_initial_deposit_ratio,omitempty"`
	// burn deposits if a proposal does not meet quorum
	BurnVoteQuorum bool `protobuf:"varint,13,opt,name=burn_vote_quorum,json=burnVoteQuorum,proto3" json:"burn_vote_quorum,omitempty"`
	// burn deposits if the proposal does not enter voting period
	BurnProposalDepositPrevote bool `protobuf:"varint,14,opt,name=burn_proposal_deposit_prevote,json=burnProposalDepositPrevote,proto3" json:"burn_proposal_deposit_prevote,omitempty"`
	// burn deposits if quorum with vote type no_veto is met
	BurnVoteVeto bool `protobuf:"varint,15,opt,name=burn_vote_veto,json=burnVoteVeto,proto3" json:"burn_vote_veto,omitempty"`
	// Minimum deposit for a emergency proposal to enter voting period.
	EmergencyMinDeposit []types.Coin `protobuf:"bytes,16,rep,name=emergency_min_deposit,json=emergencyMinDeposit,proto3" json:"emergency_min_deposit"`
	// Tally interval for emergency proposal.
	EmergencyTallyInterval time.Duration `protobuf:"bytes,17,opt,name=emergency_tally_interval,json=emergencyTallyInterval,proto3,stdduration" json:"emergency_tally_interval"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6adfe7a550f5e4ec, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMinDeposit() []types.Coin {
	if m != nil {
		return m.MinDeposit
	}
	return nil
}

func (m *Params) GetMaxDepositPeriod() time.Duration {
	if m != nil {
		return m.MaxDepositPeriod
	}
	return 0
}

func (m *Params) GetVotingPeriod() time.Duration {
	if m != nil {
		return m.VotingPeriod
	}
	return 0
}

func (m *Params) GetQuorum() string {
	if m != nil {
		return m.Quorum
	}
	return ""
}

func (m *Params) GetThreshold() string {
	if m != nil {
		return m.Threshold
	}
	return ""
}

func (m *Params) GetVetoThreshold() string {
	if m != nil {
		return m.VetoThreshold
	}
	return ""
}

func (m *Params) GetMinInitialDepositRatio() string {
	if m != nil {
		return m.MinInitialDepositRatio
	}
	return ""
}

func (m *Params) GetBurnVoteQuorum() bool {
	if m != nil {
		return m.BurnVoteQuorum
	}
	return false
}

func (m *Params) GetBurnProposalDepositPrevote() bool {
	if m != nil {
		return m.BurnProposalDepositPrevote
	}
	return false
}

func (m *Params) GetBurnVoteVeto() bool {
	if m != nil {
		return m.BurnVoteVeto
	}
	return false
}

func (m *Params) GetEmergencyMinDeposit() []types.Coin {
	if m != nil {
		return m.EmergencyMinDeposit
	}
	return nil
}

func (m *Params) GetEmergencyTallyInterval() time.Duration {
	if m != nil {
		return m.EmergencyTallyInterval
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "initia.gov.v1.Params")
}

func init() { proto.RegisterFile("initia/gov/v1/gov.proto", fileDescriptor_6adfe7a550f5e4ec) }

var fileDescriptor_6adfe7a550f5e4ec = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xb1, 0x6f, 0xd3, 0x4e,
	0x14, 0xc7, 0xe3, 0x5f, 0x7f, 0x04, 0x7a, 0x6d, 0x42, 0x6b, 0xa0, 0x75, 0x22, 0xe1, 0x44, 0x08,
	0xa1, 0x50, 0x51, 0x5b, 0x01, 0xb1, 0xb0, 0x11, 0x8a, 0x50, 0x87, 0x4a, 0x25, 0xaa, 0x22, 0xc4,
	0x62, 0x9d, 0x93, 0xc3, 0x39, 0xc9, 0x77, 0xcf, 0xd8, 0x67, 0x2b, 0xf9, 0x17, 0x98, 0x18, 0x3b,
	0x76, 0xec, 0xd8, 0x81, 0x81, 0x3f, 0xa1, 0x63, 0xc5, 0xc4, 0x04, 0x28, 0x19, 0xca, 0x9f, 0x81,
	0xce, 0x77, 0x4e, 0xaa, 0x28, 0x4b, 0x96, 0xe4, 0xfc, 0xde, 0xf7, 0x7d, 0xde, 0x7d, 0xef, 0xde,
	0xa1, 0x5d, 0xca, 0xa9, 0xa0, 0xd8, 0x0d, 0x20, 0x73, 0xb3, 0xb6, 0xfc, 0x73, 0xa2, 0x18, 0x04,
	0x98, 0x15, 0x95, 0x70, 0x64, 0x24, 0x6b, 0xd7, 0xed, 0x3e, 0x24, 0x0c, 0x12, 0xd7, 0xc7, 0x09,
	0x71, 0xb3, 0xb6, 0x4f, 0x04, 0x6e, 0xbb, 0x7d, 0xa0, 0x5c, 0xc9, 0xeb, 0xf7, 0x03, 0x08, 0x20,
	0x5f, 0xba, 0x72, 0xa5, 0xa3, 0xb5, 0x00, 0x20, 0x08, 0x89, 0x9b, 0x7f, 0xf9, 0xe9, 0x27, 0x17,
	0xf3, 0xb1, 0x4e, 0x35, 0x16, 0x53, 0x82, 0x32, 0x92, 0x08, 0xcc, 0x22, 0x2d, 0xb0, 0x17, 0x05,
	0x83, 0x34, 0xc6, 0x82, 0x42, 0xd1, 0xb1, 0xa6, 0x76, 0xe4, 0xa9, 0xa6, 0xea, 0x43, 0xa7, 0x76,
	0xf5, 0x66, 0x17, 0x4d, 0xd5, 0xb7, 0x31, 0xa3, 0x1c, 0xdc, 0xfc, 0x57, 0x85, 0x1e, 0x7d, 0x2f,
	0xa3, 0xf2, 0x31, 0x8e, 0x31, 0x4b, 0xcc, 0xb7, 0x68, 0x83, 0x51, 0xee, 0x0d, 0x48, 0x04, 0x09,
	0x15, 0x96, 0xd1, 0x5c, 0x6b, 0x6d, 0x3c, 0xaf, 0x39, 0x1a, 0x2d, 0x9d, 0x3b, 0xda, 0xb9, 0xf3,
	0x06, 0x28, 0xef, 0xac, 0x5f, 0xfe, 0x6a, 0x94, 0xce, 0xaf, 0x2f, 0xf6, 0x8c, 0x2e, 0x62, 0x94,
	0x1f, 0xa8, 0x3a, 0xb3, 0x87, 0x4c, 0x86, 0x47, 0x05, 0xc6, 0x8b, 0x48, 0x4c, 0x61, 0x60, 0xfd,
	0xd7, 0x34, 0x72, 0x9a, 0x72, 0xe5, 0x14, 0xae, 0x9c, 0x03, 0xed, 0xaa, 0x53, 0x91, 0xb4, 0xd3,
	0xdf, 0x0d, 0x43, 0x11, 0xb7, 0x18, 0x1e, 0x69, 0xe2, 0x71, 0x4e, 0x30, 0x8f, 0x50, 0x25, 0x03,
	0x41, 0x79, 0x50, 0x20, 0xd7, 0x56, 0x44, 0x6e, 0xaa, 0x72, 0x8d, 0x7b, 0x82, 0xca, 0x9f, 0x53,
	0x88, 0x53, 0x66, 0xfd, 0xdf, 0x34, 0x5a, 0xeb, 0x9d, 0xea, 0x8f, 0x6f, 0xfb, 0x48, 0x7b, 0x3d,
	0x20, 0xfd, 0xae, 0xce, 0x9a, 0xcf, 0xd0, 0xba, 0x18, 0xc6, 0x24, 0x19, 0x42, 0x38, 0xb0, 0x6e,
	0x2d, 0x95, 0xce, 0x05, 0xe6, 0x4b, 0x54, 0xcd, 0x88, 0x00, 0x6f, 0x5e, 0x52, 0x5e, 0x5a, 0x52,
	0x91, 0xaa, 0x93, 0x59, 0xd9, 0x21, 0xaa, 0xc9, 0xa3, 0x57, 0x33, 0x17, 0xce, 0xce, 0x2e, 0xf7,
	0x61, 0xdd, 0x5e, 0x4a, 0xd8, 0x61, 0x94, 0x1f, 0x2a, 0xbd, 0x3e, 0xa7, 0xae, 0x54, 0x9b, 0x2d,
	0xb4, 0xe5, 0xa7, 0x31, 0xf7, 0x32, 0x10, 0xc4, 0xd3, 0x0e, 0x2b, 0x4d, 0xa3, 0x75, 0xa7, 0x5b,
	0x95, 0xf1, 0x1e, 0x08, 0xf2, 0x5e, 0x39, 0x7b, 0x8d, 0x1e, 0xe6, 0xca, 0x28, 0x86, 0x08, 0x92,
	0x1b, 0x6d, 0xa3, 0x98, 0xc8, 0x6a, 0xab, 0x9a, 0x97, 0xd5, 0xa5, 0xe8, 0x58, 0x6b, 0x8a, 0x2b,
	0x51, 0x0a, 0xf3, 0x31, 0xaa, 0xce, 0x9b, 0x49, 0x4b, 0xd6, 0xdd, 0xbc, 0x66, 0xb3, 0x68, 0xd5,
	0x23, 0x02, 0xcc, 0x0f, 0xe8, 0x01, 0x61, 0x24, 0x0e, 0x08, 0xef, 0x8f, 0xbd, 0x9b, 0x23, 0xb6,
	0xb5, 0xc2, 0x88, 0xdd, 0x9b, 0x21, 0x8e, 0xe6, 0xb3, 0xe6, 0x23, 0x6b, 0x4e, 0x16, 0x38, 0x0c,
	0xc7, 0x1e, 0xe5, 0x82, 0xc4, 0x19, 0x0e, 0xad, 0xed, 0x15, 0xc7, 0x63, 0x67, 0x46, 0x3a, 0x91,
	0xa0, 0x43, 0xcd, 0x79, 0xb5, 0x7b, 0x7a, 0xd6, 0x28, 0xfd, 0x3d, 0x6b, 0x18, 0x5f, 0xae, 0x2f,
	0xf6, 0x90, 0x7c, 0x53, 0xea, 0xbd, 0x74, 0xde, 0x9d, 0x4f, 0x6c, 0xe3, 0x72, 0x62, 0x1b, 0x57,
	0x13, 0xdb, 0xf8, 0x33, 0xb1, 0x8d, 0xaf, 0x53, 0xbb, 0x74, 0x35, 0xb5, 0x4b, 0x3f, 0xa7, 0x76,
	0xe9, 0xe3, 0xd3, 0x80, 0x8a, 0x61, 0xea, 0x3b, 0x7d, 0x60, 0xae, 0xba, 0xd7, 0xfd, 0x10, 0xfb,
	0x89, 0x5e, 0xbb, 0xa3, 0xfc, 0x75, 0x8a, 0x71, 0x44, 0x12, 0xbf, 0x9c, 0xef, 0xed, 0xc5, 0xbf,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xdc, 0x9c, 0x0d, 0x8d, 0x04, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.MinDeposit) != len(that1.MinDeposit) {
		return false
	}
	for i := range this.MinDeposit {
		if !this.MinDeposit[i].Equal(&that1.MinDeposit[i]) {
			return false
		}
	}
	if this.MaxDepositPeriod != that1.MaxDepositPeriod {
		return false
	}
	if this.VotingPeriod != that1.VotingPeriod {
		return false
	}
	if this.Quorum != that1.Quorum {
		return false
	}
	if this.Threshold != that1.Threshold {
		return false
	}
	if this.VetoThreshold != that1.VetoThreshold {
		return false
	}
	if this.MinInitialDepositRatio != that1.MinInitialDepositRatio {
		return false
	}
	if this.BurnVoteQuorum != that1.BurnVoteQuorum {
		return false
	}
	if this.BurnProposalDepositPrevote != that1.BurnProposalDepositPrevote {
		return false
	}
	if this.BurnVoteVeto != that1.BurnVoteVeto {
		return false
	}
	if len(this.EmergencyMinDeposit) != len(that1.EmergencyMinDeposit) {
		return false
	}
	for i := range this.EmergencyMinDeposit {
		if !this.EmergencyMinDeposit[i].Equal(&that1.EmergencyMinDeposit[i]) {
			return false
		}
	}
	if this.EmergencyTallyInterval != that1.EmergencyTallyInterval {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.EmergencyTallyInterval, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.EmergencyTallyInterval):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGov(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x1
	i--
	dAtA[i] = 0x8a
	if len(m.EmergencyMinDeposit) > 0 {
		for iNdEx := len(m.EmergencyMinDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.EmergencyMinDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if m.BurnVoteVeto {
		i--
		if m.BurnVoteVeto {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if m.BurnProposalDepositPrevote {
		i--
		if m.BurnProposalDepositPrevote {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x70
	}
	if m.BurnVoteQuorum {
		i--
		if m.BurnVoteQuorum {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x68
	}
	if len(m.MinInitialDepositRatio) > 0 {
		i -= len(m.MinInitialDepositRatio)
		copy(dAtA[i:], m.MinInitialDepositRatio)
		i = encodeVarintGov(dAtA, i, uint64(len(m.MinInitialDepositRatio)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.VetoThreshold) > 0 {
		i -= len(m.VetoThreshold)
		copy(dAtA[i:], m.VetoThreshold)
		i = encodeVarintGov(dAtA, i, uint64(len(m.VetoThreshold)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Threshold) > 0 {
		i -= len(m.Threshold)
		copy(dAtA[i:], m.Threshold)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Threshold)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Quorum) > 0 {
		i -= len(m.Quorum)
		copy(dAtA[i:], m.Quorum)
		i = encodeVarintGov(dAtA, i, uint64(len(m.Quorum)))
		i--
		dAtA[i] = 0x22
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.VotingPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.VotingPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGov(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.MaxDepositPeriod, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxDepositPeriod):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGov(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.MinDeposit) > 0 {
		for iNdEx := len(m.MinDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGov(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGov(dAtA []byte, offset int, v uint64) int {
	offset -= sovGov(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.MinDeposit) > 0 {
		for _, e := range m.MinDeposit {
			l = e.Size()
			n += 1 + l + sovGov(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.MaxDepositPeriod)
	n += 1 + l + sovGov(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.VotingPeriod)
	n += 1 + l + sovGov(uint64(l))
	l = len(m.Quorum)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.Threshold)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.VetoThreshold)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	l = len(m.MinInitialDepositRatio)
	if l > 0 {
		n += 1 + l + sovGov(uint64(l))
	}
	if m.BurnVoteQuorum {
		n += 2
	}
	if m.BurnProposalDepositPrevote {
		n += 2
	}
	if m.BurnVoteVeto {
		n += 2
	}
	if len(m.EmergencyMinDeposit) > 0 {
		for _, e := range m.EmergencyMinDeposit {
			l = e.Size()
			n += 2 + l + sovGov(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.EmergencyTallyInterval)
	n += 2 + l + sovGov(uint64(l))
	return n
}

func sovGov(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGov(x uint64) (n int) {
	return sovGov(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGov
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinDeposit = append(m.MinDeposit, types.Coin{})
			if err := m.MinDeposit[len(m.MinDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDepositPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.MaxDepositPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPeriod", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.VotingPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quorum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Quorum = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Threshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Threshold = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VetoThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VetoThreshold = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinInitialDepositRatio", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinInitialDepositRatio = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnVoteQuorum", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BurnVoteQuorum = bool(v != 0)
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnProposalDepositPrevote", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BurnProposalDepositPrevote = bool(v != 0)
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnVoteVeto", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.BurnVoteVeto = bool(v != 0)
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyMinDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EmergencyMinDeposit = append(m.EmergencyMinDeposit, types.Coin{})
			if err := m.EmergencyMinDeposit[len(m.EmergencyMinDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmergencyTallyInterval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGov
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGov
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGov
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.EmergencyTallyInterval, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGov(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGov
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipGov(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGov
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowGov
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthGov
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGov
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGov
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGov        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGov          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGov = fmt.Errorf("proto: unexpected end of group")
)
