// Code generated by protoc-gen-go. DO NOT EDIT.
// source: infra/appengine/test-results/model/gen/testresults.proto

/*
Package gen is a generated protocol buffer package.

It is generated from these files:
	infra/appengine/test-results/model/gen/testresults.proto

It has these top-level messages:
	TestResults
	TestRun
	TestResultEvent
	WriteLatency
*/
package gen

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ResultType represents the result of a test, either actual or expected.
type ResultType int32

const (
	// SKIP means the test was not run.
	ResultType_SKIP ResultType = 0
	// PASS means the test ran as expected.
	ResultType_PASS ResultType = 1
	// FAIL means the test did not run as expected.
	ResultType_FAIL ResultType = 2
	// CRASH means the test runner crashed during the test.
	ResultType_CRASH ResultType = 3
	// TIMEOUT means the test hung (did not complete) and was aborted.
	ResultType_TIMEOUT ResultType = 4
	// MISSING is layout test specific. The test completed but we could not find
	// an expected baseline to compare against.
	ResultType_MISSING ResultType = 5
	// LEAK is layout test specific. Memory leaks were detected during the test execution.
	ResultType_LEAK ResultType = 6
	// SLOW is layout test specific. The test is expected to take longer than normal to run.
	ResultType_SLOW ResultType = 7
	// TEXT is layout test specific, deprecated. The test is expected to produce
	// a text-only failure (the image, if present, will match). Normally you will
	// see "FAIL" instead.
	ResultType_TEXT ResultType = 8
	// AUDIO is layout test specific, deprecated. The test is expected to produce
	// audio output that doesn't match the expected result. Normally you will see
	// "FAIL" instead.
	ResultType_AUDIO ResultType = 9
	// IMAGE is layout test specific. The test produces image (and possibly text
	// output). The image output doesn't match what we'd expect, but the text output,
	// if present, does.
	ResultType_IMAGE ResultType = 10
	// IMAGE_TEXT is layout test specific, deprecated. The test produces image
	// and text output, both of which fail to match what we expect. Normally you
	// will see "FAIL" instead.
	ResultType_IMAGE_TEXT ResultType = 11
	// REBASELINE is layout test specific. The expected test result is out of date
	// and will be ignored (any result other than a crash or timeout will be
	// considered as passing). This test result should only ever show up on local
	// test runs, not on bots (it is forbidden to check in a TestExpectations file
	// with this expectation). This should never show up as an "actual" result.
	ResultType_REBASELINE ResultType = 12
	// NEEDS_REBASELINE is layout test specific. The expected test result is out
	// of date and will be ignored (as above); the auto-rebaseline-bot will look
	// for tests of this type and automatically update them. This should never
	// show up as an "actual" result.
	ResultType_NEEDS_REBASELINE ResultType = 13
	// NEEDS_MANUAL_REBASELINE is layout test specific. The expected test result
	// is out of date and will be ignored (as above). This result may be checked
	// in to the TestExpectations file, but the auto-rebasline-bot will ignore
	// these entries. This should never show up as an "actual" result.
	ResultType_NEEDS_MANUAL_REBASELINE ResultType = 14
	// UNKNOWN is an unrecognized or unknown test result type.
	ResultType_UNKNOWN ResultType = 15
)

var ResultType_name = map[int32]string{
	0:  "SKIP",
	1:  "PASS",
	2:  "FAIL",
	3:  "CRASH",
	4:  "TIMEOUT",
	5:  "MISSING",
	6:  "LEAK",
	7:  "SLOW",
	8:  "TEXT",
	9:  "AUDIO",
	10: "IMAGE",
	11: "IMAGE_TEXT",
	12: "REBASELINE",
	13: "NEEDS_REBASELINE",
	14: "NEEDS_MANUAL_REBASELINE",
	15: "UNKNOWN",
}
var ResultType_value = map[string]int32{
	"SKIP":                    0,
	"PASS":                    1,
	"FAIL":                    2,
	"CRASH":                   3,
	"TIMEOUT":                 4,
	"MISSING":                 5,
	"LEAK":                    6,
	"SLOW":                    7,
	"TEXT":                    8,
	"AUDIO":                   9,
	"IMAGE":                   10,
	"IMAGE_TEXT":              11,
	"REBASELINE":              12,
	"NEEDS_REBASELINE":        13,
	"NEEDS_MANUAL_REBASELINE": 14,
	"UNKNOWN":                 15,
}

func (x ResultType) String() string {
	return proto.EnumName(ResultType_name, int32(x))
}
func (ResultType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// TestResults encapsulates the results of a run of a set of tests.
type TestResults struct {
	// Interrupted is true if the test run was interrupted and terminated early
	// (either via the runner bailing out or the user hitting ctrl-C, etc.)
	// If true, this indicates that not all of the tests in the suite were run
	// and the results are at best incomplete and possibly totally invalid.
	Interrupted bool `protobuf:"varint,1,opt,name=interrupted" json:"interrupted,omitempty"`
	// NumFailuresByType is a summary of the totals of each result type.
	// If a test was run more than once, only the first invocation's result is
	// included in the totals. Each key is one of the result types listed below.
	// A missing result type is the same as being present and set to zero (0).
	NumFailuresByType []*TestResults_FailuresByType `protobuf:"bytes,2,rep,name=num_failures_by_type,json=numFailuresByType" json:"num_failures_by_type,omitempty"`
	// PathDelimiter is the separator string to use in between components of a
	// tests name; normally "." for GTest- and Python-based tests and "/" for
	// layout tests; if not present, you should default to "/" for backwards-compatibility.
	PathDelimiter string `protobuf:"bytes,3,opt,name=path_delimiter,json=pathDelimiter" json:"path_delimiter,omitempty"`
	// SecondsSinceEpoch is the start time of the test run expressed as a
	// floating-point offset in seconds from the UNIX epoch.
	SecondsSinceEpoch float32 `protobuf:"fixed32,4,opt,name=seconds_since_epoch,json=secondsSinceEpoch" json:"seconds_since_epoch,omitempty"`
	// TestResults is the set of actual test results. Each directory or module
	// component in the test name is a node in the trie, and the leaf contains
	// the dict of per-test fields as described below.
	//
	// In the original JSON, this is a trie. Here, just flatten out as path strings.
	// They *should* compress just fine in transit/at rest, but we should verify
	// with measurements in the wild.
	Tests map[string]*TestRun `protobuf:"bytes,5,rep,name=Tests" json:"Tests,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// BuildId is the build_id in chrome-infra-events:raw_events.completed_builds_legacy
	BuildId string `protobuf:"bytes,6,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
}

func (m *TestResults) Reset()                    { *m = TestResults{} }
func (m *TestResults) String() string            { return proto.CompactTextString(m) }
func (*TestResults) ProtoMessage()               {}
func (*TestResults) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestResults) GetInterrupted() bool {
	if m != nil {
		return m.Interrupted
	}
	return false
}

func (m *TestResults) GetNumFailuresByType() []*TestResults_FailuresByType {
	if m != nil {
		return m.NumFailuresByType
	}
	return nil
}

func (m *TestResults) GetPathDelimiter() string {
	if m != nil {
		return m.PathDelimiter
	}
	return ""
}

func (m *TestResults) GetSecondsSinceEpoch() float32 {
	if m != nil {
		return m.SecondsSinceEpoch
	}
	return 0
}

func (m *TestResults) GetTests() map[string]*TestRun {
	if m != nil {
		return m.Tests
	}
	return nil
}

func (m *TestResults) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

// FailuresByType is a workaround for the lack of support in proto3 for
// enum types as keys in maps.
type TestResults_FailuresByType struct {
	Type  ResultType `protobuf:"varint,1,opt,name=type,enum=testresults.events.ResultType" json:"type,omitempty"`
	Count int64      `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *TestResults_FailuresByType) Reset()                    { *m = TestResults_FailuresByType{} }
func (m *TestResults_FailuresByType) String() string            { return proto.CompactTextString(m) }
func (*TestResults_FailuresByType) ProtoMessage()               {}
func (*TestResults_FailuresByType) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *TestResults_FailuresByType) GetType() ResultType {
	if m != nil {
		return m.Type
	}
	return ResultType_SKIP
}

func (m *TestResults_FailuresByType) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// TestRun represents the output results of one test run.
type TestRun struct {
	// Actual is an ordered list of the results the test actually produced.
	// {"FAIL, "PASS"} means that a test was run twice, failed the first time,
	// and then passed when it was retried. If a test produces multiple
	// different results, then it was actually flaky during the run.
	Actual []ResultType `protobuf:"varint,1,rep,packed,name=actual,enum=testresults.events.ResultType" json:"actual,omitempty"`
	// Expected is an unordered list of the result types expected for the test,
	// e.g. {"FAIL", "PASS"} means that a test is expected to either pass or fail.
	// A test that contains multiple values is expected to be flaky.
	Expected []ResultType `protobuf:"varint,2,rep,packed,name=expected,enum=testresults.events.ResultType" json:"expected,omitempty"`
	// Bugs is a list of URLs to bug database entries associated with each test.
	Bugs []string `protobuf:"bytes,3,rep,name=bugs" json:"bugs,omitempty"`
	// IsUnexpected indicates that the failure was unexpected (a regression).
	// If false, the failure was expected and will be ignored.
	IsUnexpected bool `protobuf:"varint,4,opt,name=is_unexpected,json=isUnexpected" json:"is_unexpected,omitempty"`
	// Time is the time it took in seconds to execute the first invocation of the test.
	Time float32 `protobuf:"fixed32,5,opt,name=time" json:"time,omitempty"`
	// Times are the times in seconds of each invocation of the test.
	Times []float32 `protobuf:"fixed32,6,rep,packed,name=times" json:"times,omitempty"`
	// HasRepaintOverlay indicates that the test output contains the data needed
	// to draw repaint overlays to help explain the results (only used in layout tests).
	HasRepaintOverlay bool `protobuf:"varint,7,opt,name=has_repaint_overlay,json=hasRepaintOverlay" json:"has_repaint_overlay,omitempty"`
	// IsMissingAudio indicates taht the test was supposed to have an audio
	// baseline to compare against, and we didn't find one.
	IsMissingAudio bool `protobuf:"varint,8,opt,name=is_missing_audio,json=isMissingAudio" json:"is_missing_audio,omitempty"`
	// IsMissingTest indicates that the test was supposed to have a text baseline
	// to compare against, and we didn't find one.
	IsMissingText bool `protobuf:"varint,9,opt,name=is_missing_text,json=isMissingText" json:"is_missing_text,omitempty"`
	// IsMissingVideo indicates that the test was supposed to have an image
	// baseline to compare against and we didn't find one.
	IsMissingVideo bool `protobuf:"varint,10,opt,name=is_missing_video,json=isMissingVideo" json:"is_missing_video,omitempty"`
	// IsTestHarnessTest indicates that the layout test was written using the
	// w3c's test harness and we don't necessarily have any baselines to compare against.
	IsTestharnessTest bool `protobuf:"varint,11,opt,name=is_testharness_test,json=isTestharnessTest" json:"is_testharness_test,omitempty"`
	// ReftestType may be one of "==" or "!=" to indicate that the test is a
	// "reference test" and the results were expected to match the reference or
	// not match the reference, respectively (only used in layout tests).
	ReftestType string `protobuf:"bytes,12,opt,name=reftest_type,json=reftestType" json:"reftest_type,omitempty"`
	// Name is the name of the test or test suite.
	Name string `protobuf:"bytes,13,opt,name=name" json:"name,omitempty"`
}

func (m *TestRun) Reset()                    { *m = TestRun{} }
func (m *TestRun) String() string            { return proto.CompactTextString(m) }
func (*TestRun) ProtoMessage()               {}
func (*TestRun) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestRun) GetActual() []ResultType {
	if m != nil {
		return m.Actual
	}
	return nil
}

func (m *TestRun) GetExpected() []ResultType {
	if m != nil {
		return m.Expected
	}
	return nil
}

func (m *TestRun) GetBugs() []string {
	if m != nil {
		return m.Bugs
	}
	return nil
}

func (m *TestRun) GetIsUnexpected() bool {
	if m != nil {
		return m.IsUnexpected
	}
	return false
}

func (m *TestRun) GetTime() float32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *TestRun) GetTimes() []float32 {
	if m != nil {
		return m.Times
	}
	return nil
}

func (m *TestRun) GetHasRepaintOverlay() bool {
	if m != nil {
		return m.HasRepaintOverlay
	}
	return false
}

func (m *TestRun) GetIsMissingAudio() bool {
	if m != nil {
		return m.IsMissingAudio
	}
	return false
}

func (m *TestRun) GetIsMissingText() bool {
	if m != nil {
		return m.IsMissingText
	}
	return false
}

func (m *TestRun) GetIsMissingVideo() bool {
	if m != nil {
		return m.IsMissingVideo
	}
	return false
}

func (m *TestRun) GetIsTestharnessTest() bool {
	if m != nil {
		return m.IsTestharnessTest
	}
	return false
}

func (m *TestRun) GetReftestType() string {
	if m != nil {
		return m.ReftestType
	}
	return ""
}

func (m *TestRun) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// TestResultEvent is suitable for representing a row in a BigQuery table. Note
// that while TestResults looks like a more obvious choice, the repeated
// TestRun messages in .tests will often exceed the row size limit for BQ.
// This message flattens TestRun.tests out into one row per test, which should
// more easily fit within the BQ row size limit.
type TestResultEvent struct {
	// Path is the full joined path of the test.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// TestType is derived from step name, by stripping anything but the first
	// word (before the first space), but preserving '(with patch)' suffix if it
	// was present in the original step name.
	TestType string `protobuf:"bytes,2,opt,name=test_type,json=testType" json:"test_type,omitempty"`
	// StepName is the name of the step that was running the tests. Test type
	// above is a normalized version of this name.
	StepName string `protobuf:"bytes,3,opt,name=step_name,json=stepName" json:"step_name,omitempty"`
	// Interrupted is true if the test run was interrupted and terminated early
	// (either via the runner bailing out or the user hitting ctrl-C, etc.) If
	// true, this indicates that not all of the tests in the suite were run and
	// the results are at best incomplete and possibly totally invalid.
	Interrupted bool `protobuf:"varint,4,opt,name=interrupted" json:"interrupted,omitempty"`
	// StartTime is The start time of the test run expressed as a number of
	// microseconds from the UNIX epoch.
	StartTime *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	// Run is the result of the test run.
	Run  *TestRun   `protobuf:"bytes,6,opt,name=run" json:"run,omitempty"`
	Runs []*TestRun `protobuf:"bytes,7,rep,name=runs" json:"runs,omitempty"`
	// BuildId is the build_id in chrome-infra-events:raw_events.completed_builds_legacy
	BuildId      string                        `protobuf:"bytes,8,opt,name=build_id,json=buildId" json:"build_id,omitempty"`
	BuildbotInfo *TestResultEvent_BuildbotInfo `protobuf:"bytes,9,opt,name=buildbot_info,json=buildbotInfo" json:"buildbot_info,omitempty"`
	WriteTime    *google_protobuf.Timestamp    `protobuf:"bytes,10,opt,name=write_time,json=writeTime" json:"write_time,omitempty"`
}

func (m *TestResultEvent) Reset()                    { *m = TestResultEvent{} }
func (m *TestResultEvent) String() string            { return proto.CompactTextString(m) }
func (*TestResultEvent) ProtoMessage()               {}
func (*TestResultEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TestResultEvent) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *TestResultEvent) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *TestResultEvent) GetStepName() string {
	if m != nil {
		return m.StepName
	}
	return ""
}

func (m *TestResultEvent) GetInterrupted() bool {
	if m != nil {
		return m.Interrupted
	}
	return false
}

func (m *TestResultEvent) GetStartTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TestResultEvent) GetRun() *TestRun {
	if m != nil {
		return m.Run
	}
	return nil
}

func (m *TestResultEvent) GetRuns() []*TestRun {
	if m != nil {
		return m.Runs
	}
	return nil
}

func (m *TestResultEvent) GetBuildId() string {
	if m != nil {
		return m.BuildId
	}
	return ""
}

func (m *TestResultEvent) GetBuildbotInfo() *TestResultEvent_BuildbotInfo {
	if m != nil {
		return m.BuildbotInfo
	}
	return nil
}

func (m *TestResultEvent) GetWriteTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.WriteTime
	}
	return nil
}

type TestResultEvent_BuildbotInfo struct {
	MasterName  string `protobuf:"bytes,1,opt,name=master_name,json=masterName" json:"master_name,omitempty"`
	BuilderName string `protobuf:"bytes,2,opt,name=builder_name,json=builderName" json:"builder_name,omitempty"`
	BuildNumber int64  `protobuf:"varint,3,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
}

func (m *TestResultEvent_BuildbotInfo) Reset()                    { *m = TestResultEvent_BuildbotInfo{} }
func (m *TestResultEvent_BuildbotInfo) String() string            { return proto.CompactTextString(m) }
func (*TestResultEvent_BuildbotInfo) ProtoMessage()               {}
func (*TestResultEvent_BuildbotInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

func (m *TestResultEvent_BuildbotInfo) GetMasterName() string {
	if m != nil {
		return m.MasterName
	}
	return ""
}

func (m *TestResultEvent_BuildbotInfo) GetBuilderName() string {
	if m != nil {
		return m.BuilderName
	}
	return ""
}

func (m *TestResultEvent_BuildbotInfo) GetBuildNumber() int64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

type WriteLatency struct {
	Master       string `protobuf:"bytes,1,opt,name=master" json:"master,omitempty"`
	Builder      string `protobuf:"bytes,2,opt,name=builder" json:"builder,omitempty"`
	BuildNumber  int64  `protobuf:"varint,3,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	WriteLatency int64  `protobuf:"varint,4,opt,name=write_latency,json=writeLatency" json:"write_latency,omitempty"`
	TableName    string `protobuf:"bytes,5,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
}

func (m *WriteLatency) Reset()                    { *m = WriteLatency{} }
func (m *WriteLatency) String() string            { return proto.CompactTextString(m) }
func (*WriteLatency) ProtoMessage()               {}
func (*WriteLatency) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *WriteLatency) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *WriteLatency) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *WriteLatency) GetBuildNumber() int64 {
	if m != nil {
		return m.BuildNumber
	}
	return 0
}

func (m *WriteLatency) GetWriteLatency() int64 {
	if m != nil {
		return m.WriteLatency
	}
	return 0
}

func (m *WriteLatency) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func init() {
	proto.RegisterType((*TestResults)(nil), "testresults.events.TestResults")
	proto.RegisterType((*TestResults_FailuresByType)(nil), "testresults.events.TestResults.FailuresByType")
	proto.RegisterType((*TestRun)(nil), "testresults.events.TestRun")
	proto.RegisterType((*TestResultEvent)(nil), "testresults.events.TestResultEvent")
	proto.RegisterType((*TestResultEvent_BuildbotInfo)(nil), "testresults.events.TestResultEvent.BuildbotInfo")
	proto.RegisterType((*WriteLatency)(nil), "testresults.events.WriteLatency")
	proto.RegisterEnum("testresults.events.ResultType", ResultType_name, ResultType_value)
}

func init() {
	proto.RegisterFile("infra/appengine/test-results/model/gen/testresults.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1018 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5f, 0x6f, 0x1b, 0x45,
	0x10, 0xc7, 0x3e, 0xff, 0x9d, 0xb3, 0x9d, 0xcb, 0x12, 0xc1, 0x91, 0x08, 0xea, 0x06, 0x81, 0xac,
	0x4a, 0xb5, 0xc1, 0x48, 0xa8, 0xf4, 0x09, 0x87, 0x5c, 0x8b, 0x15, 0xc7, 0xa9, 0xce, 0x36, 0x41,
	0x7d, 0x39, 0xad, 0xed, 0xb5, 0xb3, 0xe2, 0x6e, 0xef, 0x74, 0xbb, 0x97, 0xc6, 0x9f, 0x88, 0xaf,
	0xc1, 0x77, 0xe1, 0x23, 0xf0, 0xc0, 0x2b, 0xda, 0xd9, 0x4b, 0xe2, 0xa4, 0xb4, 0xc9, 0xdb, 0xcc,
	0x6f, 0x7e, 0xf3, 0x67, 0x67, 0x66, 0x07, 0x5e, 0x70, 0xb1, 0x4a, 0x69, 0x8f, 0x26, 0x09, 0x13,
	0x6b, 0x2e, 0x58, 0x4f, 0x31, 0xa9, 0x9e, 0xa7, 0x4c, 0x66, 0xa1, 0x92, 0xbd, 0x28, 0x5e, 0xb2,
	0xb0, 0xb7, 0x66, 0x02, 0xe1, 0x1c, 0xed, 0x26, 0x69, 0xac, 0x62, 0x42, 0xb6, 0x21, 0x76, 0xc9,
	0x84, 0x92, 0xfb, 0x4f, 0xd6, 0x71, 0xbc, 0x0e, 0x59, 0x0f, 0x19, 0xf3, 0x6c, 0xd5, 0x53, 0x3c,
	0x62, 0x52, 0xd1, 0x28, 0x31, 0x4e, 0x87, 0xff, 0x5a, 0x60, 0x4f, 0x99, 0x54, 0xbe, 0xf1, 0x23,
	0x6d, 0xb0, 0xb9, 0x50, 0x2c, 0x4d, 0xb3, 0x44, 0xb1, 0xa5, 0x5b, 0x68, 0x17, 0x3a, 0x35, 0x7f,
	0x1b, 0x22, 0x01, 0xec, 0x89, 0x2c, 0x0a, 0x56, 0x94, 0x87, 0x59, 0xca, 0x64, 0x30, 0xdf, 0x04,
	0x6a, 0x93, 0x30, 0xb7, 0xd8, 0xb6, 0x3a, 0x76, 0xbf, 0xdb, 0x7d, 0xbf, 0x8a, 0xee, 0x56, 0x82,
	0xee, 0xab, 0xdc, 0xef, 0x68, 0x33, 0xdd, 0x24, 0xcc, 0xdf, 0x15, 0x59, 0x74, 0x17, 0x22, 0xdf,
	0x40, 0x2b, 0xa1, 0xea, 0x22, 0x58, 0xb2, 0x90, 0x47, 0x5c, 0xb1, 0xd4, 0xb5, 0xda, 0x85, 0x4e,
	0xdd, 0x6f, 0x6a, 0xf4, 0xf8, 0x1a, 0x24, 0x5d, 0xf8, 0x54, 0xb2, 0x45, 0x2c, 0x96, 0x32, 0x90,
	0x5c, 0x2c, 0x58, 0xc0, 0x92, 0x78, 0x71, 0xe1, 0x96, 0xda, 0x85, 0x4e, 0xd1, 0xdf, 0xcd, 0x4d,
	0x13, 0x6d, 0xf1, 0xb4, 0x81, 0xfc, 0x0c, 0x65, 0x5d, 0x87, 0x74, 0xcb, 0x58, 0xe8, 0xb3, 0x87,
	0x0a, 0x45, 0xb2, 0x27, 0x54, 0xba, 0xf1, 0x8d, 0x23, 0xf9, 0x02, 0x6a, 0xf3, 0x8c, 0x87, 0xcb,
	0x80, 0x2f, 0xdd, 0x0a, 0x96, 0x54, 0x45, 0x7d, 0xb8, 0xdc, 0x7f, 0x0b, 0xad, 0x7b, 0xaf, 0xe8,
	0x43, 0x09, 0xdb, 0xa2, 0x3b, 0xd8, 0xea, 0x7f, 0xf5, 0x7f, 0xd9, 0x4c, 0x26, 0x6c, 0x03, 0x72,
	0xc9, 0x1e, 0x94, 0x17, 0x71, 0x26, 0x94, 0x5b, 0x6c, 0x17, 0x3a, 0x96, 0x6f, 0x94, 0xfd, 0x19,
	0xc0, 0x6d, 0x2d, 0xc4, 0x01, 0xeb, 0x0f, 0xb6, 0xc1, 0xb0, 0x75, 0x5f, 0x8b, 0xe4, 0x7b, 0x28,
	0x5f, 0xd2, 0x30, 0x63, 0xe8, 0x65, 0xf7, 0x0f, 0x3e, 0xf8, 0xb0, 0x4c, 0xf8, 0x86, 0xf9, 0xb2,
	0xf8, 0xa2, 0x70, 0xf8, 0x8f, 0x05, 0xd5, 0x1c, 0x26, 0x3f, 0x42, 0x85, 0x2e, 0x54, 0x46, 0x43,
	0xb7, 0xd0, 0xb6, 0x1e, 0x51, 0x6e, 0xce, 0x26, 0x2f, 0xa1, 0xc6, 0xae, 0x12, 0xb6, 0xd0, 0xab,
	0x52, 0x7c, 0x94, 0xe7, 0x0d, 0x9f, 0x10, 0x28, 0xcd, 0xb3, 0xb5, 0x74, 0xad, 0xb6, 0xd5, 0xa9,
	0xfb, 0x28, 0x93, 0xaf, 0xa1, 0xc9, 0x65, 0x90, 0x89, 0x9b, 0xa0, 0x25, 0xdc, 0xbf, 0x06, 0x97,
	0x33, 0xb1, 0xed, 0xa8, 0xb7, 0xd8, 0x2d, 0xe3, 0xa4, 0x51, 0xd6, 0x9d, 0xc3, 0xcd, 0x76, 0x2b,
	0x6d, 0xab, 0x53, 0xf4, 0x8d, 0xa2, 0x57, 0xe4, 0x82, 0xca, 0x20, 0x65, 0x09, 0xe5, 0x42, 0x05,
	0xf1, 0x25, 0x4b, 0x43, 0xba, 0x71, 0xab, 0x18, 0x74, 0xf7, 0x82, 0x4a, 0xdf, 0x58, 0xce, 0x8c,
	0x81, 0x74, 0xc0, 0xe1, 0x32, 0x88, 0xb8, 0x94, 0x5c, 0xac, 0x03, 0x9a, 0x2d, 0x79, 0xec, 0xd6,
	0x90, 0xdc, 0xe2, 0xf2, 0xd4, 0xc0, 0x03, 0x8d, 0x92, 0x6f, 0x61, 0x67, 0x8b, 0xa9, 0xd8, 0x95,
	0x72, 0xeb, 0x48, 0x6c, 0xde, 0x10, 0xa7, 0xec, 0x4a, 0xdd, 0x8b, 0x78, 0xc9, 0x97, 0x2c, 0x76,
	0xe1, 0x5e, 0xc4, 0xdf, 0x34, 0xaa, 0x6b, 0xe5, 0x32, 0xd0, 0xcd, 0xbb, 0xa0, 0xa9, 0x60, 0xd2,
	0xc8, 0xae, 0x6d, 0x6a, 0xe5, 0x72, 0x7a, 0x6b, 0xd1, 0x22, 0x79, 0x0a, 0x8d, 0x94, 0xad, 0x34,
	0xc7, 0x7c, 0xbf, 0x06, 0x2e, 0x84, 0x9d, 0x63, 0xb8, 0x82, 0x04, 0x4a, 0x82, 0x46, 0xcc, 0x6d,
	0xa2, 0x09, 0xe5, 0xc3, 0xbf, 0x4a, 0xb0, 0x73, 0xbb, 0xe5, 0x9e, 0x9e, 0x8f, 0xe6, 0xe9, 0xaf,
	0x95, 0xef, 0x14, 0xca, 0xe4, 0x00, 0xea, 0xb7, 0xb1, 0x8b, 0x68, 0xa8, 0xdd, 0x04, 0x3e, 0x80,
	0xba, 0x54, 0x2c, 0x09, 0x30, 0xba, 0xf9, 0x9c, 0x35, 0x0d, 0x8c, 0x69, 0xc4, 0xee, 0x5f, 0x90,
	0xd2, 0xfb, 0x17, 0xe4, 0x27, 0x00, 0xa9, 0x68, 0xaa, 0x82, 0x9b, 0x31, 0xda, 0xfd, 0xfd, 0xae,
	0xb9, 0x54, 0xdd, 0xeb, 0x4b, 0xd5, 0x9d, 0x5e, 0x5f, 0x2a, 0xbf, 0x8e, 0x6c, 0xad, 0x93, 0xe7,
	0x60, 0xa5, 0x99, 0xc0, 0xdf, 0xf7, 0xc0, 0xa6, 0x6b, 0x1e, 0xe9, 0x41, 0x29, 0xcd, 0x84, 0x74,
	0xab, 0xf8, 0xe5, 0x3f, 0xca, 0x47, 0xe2, 0x9d, 0x2f, 0x5e, 0xbb, 0xf3, 0xc5, 0xc9, 0x0c, 0x9a,
	0x28, 0xce, 0x63, 0x15, 0x70, 0xb1, 0x8a, 0x71, 0xe0, 0x76, 0xff, 0xbb, 0x8f, 0xdf, 0x11, 0xec,
	0x70, 0xf7, 0x28, 0x77, 0x1c, 0x8a, 0x55, 0xec, 0x37, 0xe6, 0x5b, 0x9a, 0x6e, 0xc6, 0xbb, 0x94,
	0x2b, 0x66, 0x9a, 0x01, 0x0f, 0x37, 0x03, 0xd9, 0x5a, 0xdf, 0xcf, 0xa0, 0xb1, 0x1d, 0x98, 0x3c,
	0x01, 0x3b, 0xa2, 0x52, 0xb1, 0xd4, 0x0c, 0xc6, 0x8c, 0x13, 0x0c, 0x84, 0xa3, 0x79, 0x0a, 0x26,
	0xf7, 0x35, 0xc3, 0xcc, 0xd5, 0xce, 0xb1, 0x3b, 0x94, 0x40, 0x64, 0xd1, 0x3c, 0x3f, 0xbd, 0x56,
	0x4e, 0x19, 0x23, 0x74, 0xf8, 0x67, 0x01, 0x1a, 0xe7, 0xba, 0x88, 0x11, 0x55, 0x4c, 0x2c, 0x36,
	0xe4, 0x33, 0xa8, 0x98, 0x24, 0x79, 0xca, 0x5c, 0x23, 0x2e, 0x54, 0xf3, 0xd0, 0x79, 0xa6, 0x6b,
	0xf5, 0x11, 0x59, 0xf4, 0x29, 0x30, 0x7d, 0x09, 0x4d, 0x16, 0x5c, 0x24, 0xcb, 0x6f, 0xbc, 0xdb,
	0xce, 0xfc, 0x25, 0x80, 0xa2, 0xf3, 0x90, 0x99, 0xe7, 0x94, 0x31, 0x49, 0x1d, 0x11, 0xfd, 0x98,
	0x67, 0x7f, 0x17, 0x00, 0x6e, 0x6f, 0x0f, 0xa9, 0x41, 0x69, 0x72, 0x32, 0x7c, 0xe3, 0x7c, 0xa2,
	0xa5, 0x37, 0x83, 0xc9, 0xc4, 0x29, 0x68, 0xe9, 0xd5, 0x60, 0x38, 0x72, 0x8a, 0xa4, 0x0e, 0xe5,
	0x5f, 0xfc, 0xc1, 0xe4, 0x57, 0xc7, 0x22, 0x36, 0x54, 0xa7, 0xc3, 0x53, 0xef, 0x6c, 0x36, 0x75,
	0x4a, 0x5a, 0x39, 0x1d, 0x4e, 0x26, 0xc3, 0xf1, 0x6b, 0xa7, 0xac, 0xe9, 0x23, 0x6f, 0x70, 0xe2,
	0x54, 0x30, 0xd8, 0xe8, 0xec, 0xdc, 0xa9, 0x6a, 0x69, 0xea, 0xfd, 0x3e, 0x75, 0x6a, 0x3a, 0xc4,
	0x60, 0x76, 0x3c, 0x3c, 0x73, 0xea, 0x5a, 0x1c, 0x9e, 0x0e, 0x5e, 0x7b, 0x0e, 0x90, 0x16, 0x00,
	0x8a, 0x01, 0xb2, 0x6c, 0xad, 0xfb, 0xde, 0xd1, 0x60, 0xe2, 0x8d, 0x86, 0x63, 0xcf, 0x69, 0x90,
	0x3d, 0x70, 0xc6, 0x9e, 0x77, 0x3c, 0x09, 0xb6, 0xd0, 0x26, 0x39, 0x80, 0xcf, 0x0d, 0x7a, 0x3a,
	0x18, 0xcf, 0x06, 0xa3, 0x6d, 0x63, 0x4b, 0xd7, 0x34, 0x1b, 0x9f, 0x8c, 0xcf, 0xce, 0xc7, 0xce,
	0xce, 0x51, 0xf9, 0xad, 0xb5, 0x66, 0x62, 0x5e, 0xc1, 0x65, 0xf9, 0xe1, 0xbf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x64, 0x5f, 0x46, 0x55, 0x41, 0x08, 0x00, 0x00,
}
