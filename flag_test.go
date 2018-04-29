package flaggy

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

// debugOff makes defers easier
func debugOff() {
	DebugMode = false
}

func TestGlobs(t *testing.T) {
	t.Skip("This is only used to test os arg globbing")
	for _, a := range os.Args {
		fmt.Println(a)
	}
}

func TestParseArgWithValue(t *testing.T) {

	testCases := make(map[string][]string)
	testCases["-f=test"] = []string{"f", "test"}
	testCases["--f=test"] = []string{"f", "test"}
	testCases["--flag=test"] = []string{"flag", "test"}
	testCases["-flag=test"] = []string{"flag", "test"}
	testCases["----flag=--test"] = []string{"--flag", "--test"}
	testCases["-b"] = []string{"b", ""}
	testCases["--bool"] = []string{"bool", ""}

	for arg, correctValues := range testCases {
		key, value := parseArgWithValue(arg)
		if key != correctValues[0] {
			t.Fatalf("Flag %s parsed key as %s but expected key %s", arg, key, correctValues[0])
		}
		if value != correctValues[1] {
			t.Fatalf("Flag %s parsed value as %s but expected value %s", arg, value, correctValues[1])
		}
		t.Logf("Flag %s parsed key as %s and value as %s correctly", arg, key, value)
	}
}

func TestDetermineArgType(t *testing.T) {

	testCases := make(map[string]string)
	testCases["-f"] = argIsFlagWithSpace
	testCases["--f"] = argIsFlagWithSpace
	testCases["-flag"] = argIsFlagWithSpace
	testCases["--flag"] = argIsFlagWithSpace
	testCases["positionalArg"] = argIsPositional
	testCases["subcommand"] = argIsPositional
	testCases["sub--+/\\324command"] = argIsPositional
	testCases["--flag=CONTENT"] = argIsFlagWithValue
	testCases["-flag=CONTENT"] = argIsFlagWithValue
	testCases["-anotherfl-ag=CONTENT"] = argIsFlagWithValue
	testCases["--anotherfl-ag=CONTENT"] = argIsFlagWithValue
	testCases["1--anotherfl-ag=CONTENT"] = argIsPositional

	for arg, correctArgType := range testCases {
		argType := determineArgType(arg)
		if argType != correctArgType {
			t.Fatalf("Flag %s determined to be type %s but expected type %s", arg, argType, correctArgType)
		} else {
			t.Logf("Flag %s correctly determined to be type %s", arg, argType)
		}
	}
}

// TestInputParsing tests all flag types.
func TestInputParsing(t *testing.T) {

	var err error

	ResetParser()

	inputArgs := []string{}

	var stringFlag string
	err = AddStringFlag(&stringFlag, "s", "string", "string flag")
	if err != nil {
		t.Fatal(err)
	}
	inputArgs = append(inputArgs, "-s", "flaggy")
	// TODO - input args for every flag
	// TODO - desired output for every flag

	var stringSliceFlag []string
	err = AddStringSliceFlag(&stringSliceFlag, "ssf", "stringSlice", "string slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var boolFlag bool
	err = AddBoolFlag(&boolFlag, "bf", "bool", "bool flag")
	if err != nil {
		t.Fatal(err)
	}

	var boolSliceFlag []bool
	err = AddBoolSliceFlag(&boolSliceFlag, "bsf", "boolSlice", "bool slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var byteSliceFlag []byte
	err = AddByteSliceFlag(&byteSliceFlag, "bysf", "byteSlice", "byte slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var durationFlag time.Duration
	err = AddDurationFlag(&durationFlag, "df", "duration", "duration flag")
	if err != nil {
		t.Fatal(err)
	}

	var durationSliceFlag []time.Duration
	err = AddDurationSliceFlag(&durationSliceFlag, "dsf", "durationSlice", "duration slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var float32Flag float32
	err = AddFloat32Flag(&float32Flag, "f32", "float32", "float32 flag")
	if err != nil {
		t.Fatal(err)
	}

	var float32SliceFlag []float32
	err = AddFloat32SliceFlag(&float32SliceFlag, "f32s", "float32Slice", "float32 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var float64Flag float64
	err = AddFloat64Flag(&float64Flag, "f64", "float64", "float64 flag")
	if err != nil {
		t.Fatal(err)
	}

	var float64SliceFlag []float64
	err = AddFloat64SliceFlag(&float64SliceFlag, "f64s", "float64Slice", "float64 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var intFlag int
	err = AddIntFlag(&intFlag, "i", "int", "int flag")
	if err != nil {
		t.Fatal(err)
	}

	var intSliceFlag []int
	err = AddIntSliceFlag(&intSliceFlag, "is", "intSlice", "int slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var uintFlag uint
	err = AddUIntFlag(&uintFlag, "ui", "uint", "uint flag")
	if err != nil {
		t.Fatal(err)
	}

	var uintSliceFlag []uint
	err = AddUIntSliceFlag(&uintSliceFlag, "uis", "uintSlice", "uint slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint64Flag uint64
	err = AddUInt64Flag(&uint64Flag, "ui64", "uint64", "uint64 flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint64SliceFlag []uint64
	err = AddUInt64SliceFlag(&uint64SliceFlag, "ui64s", "uint64Slice", "uint64 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint32Flag uint32
	err = AddUInt32Flag(&uint32Flag, "ui32", "uint32", "uint32 flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint32SliceFlag []uint32
	err = AddUInt32SliceFlag(&uint32SliceFlag, "ui32s", "uint32Slice", "uint32 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint16Flag uint16
	err = AddUInt16Flag(&uint16Flag, "ui16", "uint16", "uint16 flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint16SliceFlag []uint16
	err = AddUInt16SliceFlag(&uint16SliceFlag, "ui16s", "uint16Slice", "uint16 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint8Flag uint8
	err = AddUInt8Flag(&uint8Flag, "ui8", "uint8", "uint8 flag")
	if err != nil {
		t.Fatal(err)
	}

	var uint8SliceFlag []uint8
	err = AddUInt8SliceFlag(&uint8SliceFlag, "ui8s", "uint8Slice", "uint8 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var int64Flag int64
	err = AddInt64Flag(&int64Flag, "i64", "int64", "int64 flag")
	if err != nil {
		t.Fatal(err)
	}

	var int64SliceFlag []int64
	err = AddInt64SliceFlag(&int64SliceFlag, "ui64s", "uint64Slice", "uint64 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var int32Flag int32
	err = AddInt32Flag(&int32Flag, "i32", "int32", "int32 flag")
	if err != nil {
		t.Fatal(err)
	}

	var int32SliceFlag []int32
	err = AddInt32SliceFlag(&int32SliceFlag, "ui32s", "uint32Slice", "uint32 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var int16Flag int16
	err = AddInt16Flag(&int16Flag, "i16", "int16", "int16 flag")
	if err != nil {
		t.Fatal(err)
	}

	var int16SliceFlag []int16
	err = AddInt16SliceFlag(&int16SliceFlag, "ui16s", "uint16Slice", "uint16 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var int8Flag int8
	err = AddInt8Flag(&int8Flag, "i8", "int8", "int8 flag")
	if err != nil {
		t.Fatal(err)
	}

	var int8SliceFlag []int8
	err = AddInt8SliceFlag(&int8SliceFlag, "ui8s", "uint8Slice", "uint8 slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var ipFlag net.IP
	err = AddIPFlag(&ipFlag, "ip", "ipFlag", "ip flag")
	if err != nil {
		t.Fatal(err)
	}

	var ipSliceFlag []net.IP
	err = AddIPSliceFlag(&ipSliceFlag, "ips", "ipFlagSlice", "ip slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var hwFlag net.HardwareAddr
	err = AddHardwareAddrFlag(&hwFlag, "hw", "hwFlag", "hw flag")
	if err != nil {
		t.Fatal(err)
	}

	var hwFlagSlice []net.HardwareAddr
	err = AddHardwareAddrSliceFlag(&hwFlagSlice, "hws", "hwFlagSlice", "hw slice flag")
	if err != nil {
		t.Fatal(err)
	}

	var maskFlag net.IPMask
	err = AddIPMaskFlag(&maskFlag, "m", "mFlag", "mask flag")
	if err != nil {
		t.Fatal(err)
	}

	var maskSliceFlag []net.IPMask
	err = AddIPMaskSliceFlag(&maskSliceFlag, "ms", "mFlagSlice", "mask slice flag")
	if err != nil {
		t.Fatal(err)
	}

	err = ParseArgs(inputArgs)
	if err != nil {
		t.Fatal(err)
	}

}
