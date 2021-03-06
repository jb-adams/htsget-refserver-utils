// Package htsformats contains objects modeling entities in genomic file formats
// and associated behaviors
//
// Module samrecord_test tests samrecord
package htsformats

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// newSamRecordTC test cases for NewSamRecord
var newSamRecordTC = []struct {
	raw, expQname string
}{
	{"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100", "A00111:67:H3M5YDMXX:2:1182:16125:23813"},
	{"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100", "A00111:67:H3M5YDMXX:1:2407:21558:16094"},
	{"A00111:67:H3M5YDMXX:2:2377:18322:22200\t163\tchr1\t24613365\t3\t100M\t=\t24613584\t296\tATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGTTTCAAAGTATTCTGAAGCTTGGAGGATGGTGAAGTAAAGTCC\tFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFF--FFFFFFFFFFFFF-FFFFFFFFFFFFFFFFFFFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100", "A00111:67:H3M5YDMXX:2:2377:18322:22200"},
	{"A00111:67:H3M5YDMXX:1:2344:29939:2018\t99\tchr1\t24613587\t255\t100M\t=\t24613883\t385\tCTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC\t8--FFFFFFFFFFFF--FFFFFFFFFFFFFFFFFF-FFFF-FFFFFFFFFFFFFFFFFFF-FFFFFFFFF-FFFFFFF-FFFFFFFFFFFFFFFF-F-FF\tNH:i:1\tHI:i:1\tNM:i:1\tMD:Z:80T19", "A00111:67:H3M5YDMXX:1:2344:29939:2018"},
	{"A00111:67:H3M5YDMXX:1:1263:33003:30342\t99\tchr1\t24613673\t3\t100M\t=\t24613757\t183\tGCTCAGAAGAATCCTGCAAAGAAAAATACTTCCGAGACGATGAATAGAATTATACCATATCGTAGTCCTTTNTGTACAATAGGAGTGTGGTGGCCTTGGT    F8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF\tNH:i:2\tHI:i:1\tNM:i:1\tMD:Z:71T28", "A00111:67:H3M5YDMXX:1:1263:33003:30342"},
}

// samRecordEmitFieldsTC test cases for emitFields
var samRecordEmitFieldsTC = []struct {
	raw string
}{
	{"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100"},
	{"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100"},
	{"A00111:67:H3M5YDMXX:2:2377:18322:22200\t163\tchr1\t24613365\t3\t100M\t=\t24613584\t296\tATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGTTTCAAAGTATTCTGAAGCTTGGAGGATGGTGAAGTAAAGTCC\tFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFF--FFFFFFFFFFFFF-FFFFFFFFFFFFFFFFFFFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100"},
	{"A00111:67:H3M5YDMXX:1:2344:29939:2018\t99\tchr1\t24613587\t255\t100M\t=\t24613883\t385\tCTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC\t8--FFFFFFFFFFFF--FFFFFFFFFFFFFFFFFF-FFFF-FFFFFFFFFFFFFFFFFFF-FFFFFFFFF-FFFFFFF-FFFFFFFFFFFFFFFF-F-FF\tNH:i:1\tHI:i:1\tNM:i:1\tMD:Z:80T19"},
	{"A00111:67:H3M5YDMXX:1:1263:33003:30342\t99\tchr1\t24613673\t3\t100M\t=\t24613757\t183\tGCTCAGAAGAATCCTGCAAAGAAAAATACTTCCGAGACGATGAATAGAATTATACCATATCGTAGTCCTTTNTGTACAATAGGAGTGTGGTGGCCTTGGT    F8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF\tNH:i:2\tHI:i:1\tNM:i:1\tMD:Z:71T28"},
}

// samRecordGetFieldTC test cases for getField
var samRecordGetFieldTC = []struct {
	raw     string
	cols    []int
	expVals []string
}{
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		[]int{0, 1, 2, 3, 4, 5},
		[]string{"A00111:67:H3M5YDMXX:2:1182:16125:23813", "147", "ERCC-00171", "280", "255", "100M"},
	},
	{
		"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100",
		[]int{6, 7, 8, 9, 10},
		[]string{"=", "24613553", "330", "CAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT", "FFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF"},
	},
	{
		"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100",
		[]int{2, 4, 6, 8, 10},
		[]string{"chr1", "3", "=", "330", "FFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF"},
	},
	{
		"A00111:67:H3M5YDMXX:1:2344:29939:2018\t99\tchr1\t24613587\t255\t100M\t=\t24613883\t385\tCTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC\t8--FFFFFFFFFFFF--FFFFFFFFFFFFFFFFFF-FFFF-FFFFFFFFFFFFFFFFFFF-FFFFFFFFF-FFFFFFF-FFFFFFFFFFFFFFFF-F-FF\tNH:i:1\tHI:i:1\tNM:i:1\tMD:Z:80T19",
		[]int{1, 3, 5, 7, 9},
		[]string{"99", "24613587", "100M", "24613883", "CTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC"},
	},
	{
		"A00111:67:H3M5YDMXX:1:1263:33003:30342\t99\tchr1\t24613673\t3\t100M\t=\t24613757\t183\tGCTCAGAAGAATCCTGCAAAGAAAAATACTTCCGAGACGATGAATAGAATTATACCATATCGTAGTCCTTTNTGTACAATAGGAGTGTGGTGGCCTTGGT\tF8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF\tNH:i:2\tHI:i:1\tNM:i:1\tMD:Z:71T28",
		[]int{10, 7, 4, 1, 0},
		[]string{"F8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF", "24613757", "3", "99", "A00111:67:H3M5YDMXX:1:1263:33003:30342"},
	},
}

// samRecordEmitTagsTC test cases for emitTags
var samRecordEmitTagsTC = []struct {
	raw string
}{
	{"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100"},
	{"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100"},
	{"A00111:67:H3M5YDMXX:2:2377:18322:22200\t163\tchr1\t24613365\t3\t100M\t=\t24613584\t296\tATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGTTTCAAAGTATTCTGAAGCTTGGAGGATGGTGAAGTAAAGTCC\tFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFF--FFFFFFFFFFFFF-FFFFFFFFFFFFFFFFFFFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100"},
	{"A00111:67:H3M5YDMXX:1:2344:29939:2018\t99\tchr1\t24613587\t255\t100M\t=\t24613883\t385\tCTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC\t8--FFFFFFFFFFFF--FFFFFFFFFFFFFFFFFF-FFFF-FFFFFFFFFFFFFFFFFFF-FFFFFFFFF-FFFFFFF-FFFFFFFFFFFFFFFF-F-FF\tNH:i:1\tHI:i:1\tNM:i:1\tMD:Z:80T19"},
	{"A00111:67:H3M5YDMXX:1:1263:33003:30342\t99\tchr1\t24613673\t3\t100M\t=\t24613757\t183\tGCTCAGAAGAATCCTGCAAAGAAAAATACTTCCGAGACGATGAATAGAATTATACCATATCGTAGTCCTTTNTGTACAATAGGAGTGTGGTGGCCTTGGT    F8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF\tNH:i:2\tHI:i:1\tNM:i:1\tMD:Z:71T28"},
}

// samRecordGetTagTC test cases for getTag
var samRecordGetTagTC = []struct {
	raw           string
	keys, expVals []string
}{
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		[]string{"MD"},
		[]string{"MD:Z:100"},
	},
	{
		"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100",
		[]string{"MD", "HI"},
		[]string{"MD:Z:100", "HI:i:1"},
	},
	{
		"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100",
		[]string{"NH", "NM"},
		[]string{"NH:i:2", "NM:i:0"},
	},
	{
		"A00111:67:H3M5YDMXX:1:2344:29939:2018\t99\tchr1\t24613587\t255\t100M\t=\t24613883\t385\tCTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC\t8--FFFFFFFFFFFF--FFFFFFFFFFFFFFFFFF-FFFF-FFFFFFFFFFFFFFFFFFF-FFFFFFFFF-FFFFFFF-FFFFFFFFFFFFFFFF-F-FF\tNH:i:1\tHI:i:1\tNM:i:1\tMD:Z:80T19",
		[]string{"NM", "HI", "NH"},
		[]string{"NM:i:1", "HI:i:1", "NH:i:1"},
	},
	{
		"A00111:67:H3M5YDMXX:1:1263:33003:30342\t99\tchr1\t24613673\t3\t100M\t=\t24613757\t183\tGCTCAGAAGAATCCTGCAAAGAAAAATACTTCCGAGACGATGAATAGAATTATACCATATCGTAGTCCTTTNTGTACAATAGGAGTGTGGTGGCCTTGGT\tF8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF\tNH:i:2\tHI:i:1\tNM:i:1\tMD:Z:71T28",
		[]string{"HI", "MD", "NM", "NH"},
		[]string{"HI:i:1", "MD:Z:71T28", "NM:i:1", "NH:i:2"},
	},
}

// samRecordStringTC test cases for String
var samRecordStringTC = []struct {
	raw, exp string
}{
	{
		"A00111:67:H3M5YDMXX:2:1182:16125:23813\t147\tERCC-00171\t280\t255\t100M\t=\t1\t-379\tAACCAAACATCCGTGCGATTCGTGCCACTCGTAGACGGCATCTCACAGTCACTGAAGGCTATTAAAGAGTTAGCACCCACCATTGGATGAAGCCCAGGAT\tFFFFFFFFFF-FFFFFFFF-FFFF-F-F-FFFFFFFFFFFFFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF\tNH:i:1\tHI:i:1\tNM:i:0\tMD:Z:100",
		"[SamRecord qname=A00111:67:H3M5YDMXX:2:1182:16125:23813]",
	},
	{
		"A00111:67:H3M5YDMXX:1:2407:21558:16094\t99\tchr1\t24613323\t3\t100M\t=\t24613553\t330\tCAATAAGGAATGTTGATCCAATAATTACATGGAGTCCATGGAATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGT\tFFFFFFFFFFFFFFFFFF8FFFFF8FFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFFF-FFFFFFFFF--F-FFFFFFFFF-FFFFF-FFF-F-FFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100",
		"[SamRecord qname=A00111:67:H3M5YDMXX:1:2407:21558:16094]",
	},
	{
		"A00111:67:H3M5YDMXX:2:2377:18322:22200\t163\tchr1\t24613365\t3\t100M\t=\t24613584\t296\tATCCAGTAGCCATGAAGAATGTAGAACCATAGATACCATCTGAAATGGAGAATGATGTTTCAAAGTATTCTGAAGCTTGGAGGATGGTGAAGTAAAGTCC\tFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF-FFFFFF--FFFFFFFFFFFFF-FFFFFFFFFFFFFFFFFFFFFFF\tNH:i:2\tHI:i:1\tNM:i:0\tMD:Z:100",
		"[SamRecord qname=A00111:67:H3M5YDMXX:2:2377:18322:22200]",
	},
	{
		"A00111:67:H3M5YDMXX:1:2344:29939:2018\t99\tchr1\t24613587\t255\t100M\t=\t24613883\t385\tCTTCTAGAGGGTTAAGTGGTGAAATTCCTGTTGGAGGTCAGCAGCCTCCTAGATCATGTGTTGGTACGAGGCTAGAATGACAGAACGCTCAGAAGAATCC\t8--FFFFFFFFFFFF--FFFFFFFFFFFFFFFFFF-FFFF-FFFFFFFFFFFFFFFFFFF-FFFFFFFFF-FFFFFFF-FFFFFFFFFFFFFFFF-F-FF\tNH:i:1\tHI:i:1\tNM:i:1\tMD:Z:80T19",
		"[SamRecord qname=A00111:67:H3M5YDMXX:1:2344:29939:2018]",
	},
	{
		"A00111:67:H3M5YDMXX:1:1263:33003:30342\t99\tchr1\t24613673\t3\t100M\t=\t24613757\t183\tGCTCAGAAGAATCCTGCAAAGAAAAATACTTCCGAGACGATGAATAGAATTATACCATATCGTAGTCCTTTNTGTACAATAGGAGTGTGGTGGCCTTGGT    F8FFFFFFFFFFFFFFFFFFF8FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF#FFFFFFFFFFFFFFF-FFFFFFF-FFFF\tNH:i:2\tHI:i:1\tNM:i:1\tMD:Z:71T28",
		"[SamRecord qname=A00111:67:H3M5YDMXX:1:1263:33003:30342]",
	},
}

// TestNewSamRecord tests NewSamRecord function
func TestNewSamRecord(t *testing.T) {
	for _, tc := range newSamRecordTC {
		samRecord := NewSamRecord(tc.raw)
		assert.Equal(t, tc.expQname, samRecord.qname)
	}
}

// TestEmitFields tests emitFields function
func TestEmitFields(t *testing.T) {
	for _, tc := range samRecordEmitFieldsTC {
		samRecord := NewSamRecord(tc.raw)
		expected := strings.Split(tc.raw, "\t")[:11]
		actual := samRecord.emitFields()
		assert.Equal(t, expected, actual)
	}
}

// TestGetField tests getField function
func TestGetField(t *testing.T) {
	for _, tc := range samRecordGetFieldTC {
		samRecord := NewSamRecord(tc.raw)
		for i := range tc.cols {
			expVal := tc.expVals[i]
			actualVal := samRecord.getField(tc.cols[i])
			assert.Equal(t, expVal, actualVal)
		}
	}
}

// TestEmitTags tests emitTags function
func TestEmitTags(t *testing.T) {
	for _, tc := range samRecordEmitTagsTC {
		samRecord := NewSamRecord(tc.raw)
		expected := strings.Split(tc.raw, "\t")[11:]
		actual := samRecord.emitTags()
		assert.Equal(t, expected, actual)
	}
}

// TestGetTag tests getTag function
func TestGetTag(t *testing.T) {
	for _, tc := range samRecordGetTagTC {
		samRecord := NewSamRecord(tc.raw)
		for i := range tc.keys {
			expected := tc.expVals[i]
			actual := samRecord.getTag(tc.keys[i])
			assert.Equal(t, expected, actual)
		}
	}
}

// TestSamRecordString tests String function
func TestSamRecordString(t *testing.T) {
	for _, tc := range samRecordStringTC {
		samRecord := NewSamRecord(tc.raw)
		assert.Equal(t, tc.exp, samRecord.String())
	}
}
