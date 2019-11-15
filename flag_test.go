package flaggy

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

// debugOff makes defers easier and turns off debug mode
func debugOff() {
	DebugMode = false
}

// debugOn turns on debug mode
func debugOn() {
	DebugMode = true
}

func TestGlobs(t *testing.T) {
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
	defer debugOff()
	DebugMode = true

	ResetParser()
	var err error
	inputArgs := []string{}

	// Setup input arguments for every input type

	var stringFlag = "defaultVar"
	String(&stringFlag, "s", "string", "string flag")
	inputArgs = append(inputArgs, "-s", "flaggy")
	var stringFlagExpected = "flaggy"

	var stringSliceFlag []string
	StringSlice(&stringSliceFlag, "ssf", "stringSlice", "string slice flag")
	inputArgs = append(inputArgs, "-ssf", "one", "-ssf", "two")
	var stringSliceFlagExpected = []string{"one", "two"}

	var boolFlag bool
	Bool(&boolFlag, "bf", "bool", "bool flag")
	inputArgs = append(inputArgs, "-bf")
	var boolFlagExpected = true

	var boolSliceFlag []bool
	BoolSlice(&boolSliceFlag, "bsf", "boolSlice", "bool slice flag")
	inputArgs = append(inputArgs, "-bsf", "-bsf")
	var boolSliceFlagExpected = []bool{true, true}

	var byteSliceFlag []byte
	ByteSlice(&byteSliceFlag, "bysf", "byteSlice", "byte slice flag")
	inputArgs = append(inputArgs, "-bysf", "17", "-bysf", "18")
	var byteSliceFlagExpected = []uint8{17, 18}

	var durationFlag time.Duration
	Duration(&durationFlag, "df", "duration", "duration flag")
	inputArgs = append(inputArgs, "-df", "33s")
	var durationFlagExpected = time.Second * 33

	var durationSliceFlag []time.Duration
	DurationSlice(&durationSliceFlag, "dsf", "durationSlice", "duration slice flag")
	inputArgs = append(inputArgs, "-dsf", "33s", "-dsf", "1h")
	var durationSliceFlagExpected = []time.Duration{time.Second * 33, time.Hour}

	var float32Flag float32
	Float32(&float32Flag, "f32", "float32", "float32 flag")
	inputArgs = append(inputArgs, "-f32", "33.343")
	var float32FlagExpected float32 = 33.343

	var float32SliceFlag []float32
	Float32Slice(&float32SliceFlag, "f32s", "float32Slice", "float32 slice flag")
	inputArgs = append(inputArgs, "-f32s", "33.343", "-f32s", "33.222")
	var float32SliceFlagExpected = []float32{33.343, 33.222}

	var float64Flag float64
	Float64(&float64Flag, "f64", "float64", "float64 flag")
	inputArgs = append(inputArgs, "-f64", "33.222343")
	var float64FlagExpected = 33.222343

	var float64SliceFlag []float64
	Float64Slice(&float64SliceFlag, "f64s", "float64Slice", "float64 slice flag")
	inputArgs = append(inputArgs, "-f64s", "64.343", "-f64s", "64.222")
	var float64SliceFlagExpected = []float64{64.343, 64.222}

	var intFlag int
	Int(&intFlag, "i", "int", "int flag")
	inputArgs = append(inputArgs, "-i", "3553")
	var intFlagExpected = 3553

	var intSliceFlag []int
	IntSlice(&intSliceFlag, "is", "intSlice", "int slice flag")
	inputArgs = append(inputArgs, "-is", "6446", "-is", "64")
	var intSliceFlagExpected = []int{6446, 64}

	var uintFlag uint
	UInt(&uintFlag, "ui", "uint", "uint flag")
	inputArgs = append(inputArgs, "-ui", "3553")
	var uintFlagExpected uint = 3553

	var uintSliceFlag []uint
	UIntSlice(&uintSliceFlag, "uis", "uintSlice", "uint slice flag")
	inputArgs = append(inputArgs, "-uis", "6446", "-uis", "64")
	var uintSliceFlagExpected = []uint{6446, 64}

	var uint64Flag uint64
	UInt64(&uint64Flag, "ui64", "uint64", "uint64 flag")
	inputArgs = append(inputArgs, "-ui64", "3553")
	var uint64FlagExpected uint64 = 3553

	var uint64SliceFlag []uint64
	UInt64Slice(&uint64SliceFlag, "ui64s", "uint64Slice", "uint64 slice flag")
	inputArgs = append(inputArgs, "-ui64s", "6446", "-ui64s", "64")
	var uint64SliceFlagExpected = []uint64{6446, 64}

	var uint32Flag uint32
	UInt32(&uint32Flag, "ui32", "uint32", "uint32 flag")
	inputArgs = append(inputArgs, "-ui32", "6446")
	var uint32FlagExpected uint32 = 6446

	var uint32SliceFlag []uint32
	UInt32Slice(&uint32SliceFlag, "ui32s", "uint32Slice", "uint32 slice flag")
	inputArgs = append(inputArgs, "-ui32s", "6446", "-ui32s", "64")
	var uint32SliceFlagExpected = []uint32{6446, 64}

	var uint16Flag uint16
	UInt16(&uint16Flag, "ui16", "uint16", "uint16 flag")
	inputArgs = append(inputArgs, "-ui16", "6446")
	var uint16FlagExpected uint16 = 6446

	var uint16SliceFlag []uint16
	UInt16Slice(&uint16SliceFlag, "ui16s", "uint16Slice", "uint16 slice flag")
	inputArgs = append(inputArgs, "-ui16s", "6446", "-ui16s", "64")
	var uint16SliceFlagExpected = []uint16{6446, 64}

	var uint8Flag uint8
	UInt8(&uint8Flag, "ui8", "uint8", "uint8 flag")
	inputArgs = append(inputArgs, "-ui8", "50")
	var uint8FlagExpected uint8 = 50

	var uint8SliceFlag []uint8
	UInt8Slice(&uint8SliceFlag, "ui8s", "uint8Slice", "uint8 slice flag")
	inputArgs = append(inputArgs, "-ui8s", "3", "-ui8s", "2")
	var uint8SliceFlagExpected = []uint8{uint8(3), uint8(2)}

	var int64Flag int64
	Int64(&int64Flag, "i64", "i64", "int64 flag")
	inputArgs = append(inputArgs, "-i64", "33445566")
	var int64FlagExpected int64 = 33445566

	var int64SliceFlag []int64
	Int64Slice(&int64SliceFlag, "i64s", "int64Slice", "int64 slice flag")
	if err != nil {
		t.Fatal(err)
	}
	inputArgs = append(inputArgs, "-i64s", "40", "-i64s", "50")
	var int64SliceFlagExpected = []int64{40, 50}

	var int32Flag int32
	Int32(&int32Flag, "i32", "int32", "int32 flag")
	inputArgs = append(inputArgs, "-i32", "445566")
	var int32FlagExpected int32 = 445566

	var int32SliceFlag []int32
	Int32Slice(&int32SliceFlag, "i32s", "int32Slice", "uint32 slice flag")
	inputArgs = append(inputArgs, "-i32s", "40", "-i32s", "50")
	var int32SliceFlagExpected = []int32{40, 50}

	var int16Flag int16
	Int16(&int16Flag, "i16", "int16", "int16 flag")
	if err != nil {
		t.Fatal(err)
	}
	inputArgs = append(inputArgs, "-i16", "5566")
	var int16FlagExpected int16 = 5566

	var int16SliceFlag []int16
	Int16Slice(&int16SliceFlag, "i16s", "int16Slice", "int16 slice flag")
	inputArgs = append(inputArgs, "-i16s", "40", "-i16s", "50")
	var int16SliceFlagExpected = []int16{40, 50}

	var int8Flag int8
	Int8(&int8Flag, "i8", "int8", "int8 flag")
	inputArgs = append(inputArgs, "-i8", "32")
	var int8FlagExpected int8 = 32

	var int8SliceFlag []int8
	Int8Slice(&int8SliceFlag, "i8s", "int8Slice", "uint8 slice flag")
	inputArgs = append(inputArgs, "-i8s", "4", "-i8s", "2")
	var int8SliceFlagExpected = []int8{4, 2}

	var ipFlag net.IP
	IP(&ipFlag, "ip", "ipFlag", "ip flag")
	inputArgs = append(inputArgs, "-ip", "1.1.1.1")
	var ipFlagExpected = net.IPv4(1, 1, 1, 1)

	var ipSliceFlag []net.IP
	IPSlice(&ipSliceFlag, "ips", "ipFlagSlice", "ip slice flag")
	inputArgs = append(inputArgs, "-ips", "1.1.1.1", "-ips", "4.4.4.4")
	var ipSliceFlagExpected = []net.IP{net.IPv4(1, 1, 1, 1), net.IPv4(4, 4, 4, 4)}

	var hwFlag net.HardwareAddr
	HardwareAddr(&hwFlag, "hw", "hwFlag", "hw flag")
	inputArgs = append(inputArgs, "-hw", "32:00:16:46:20:00")
	hwFlagExpected, err := net.ParseMAC("32:00:16:46:20:00")
	if err != nil {
		t.Fatal(err)
	}

	var hwFlagSlice []net.HardwareAddr
	HardwareAddrSlice(&hwFlagSlice, "hws", "hwFlagSlice", "hw slice flag")
	inputArgs = append(inputArgs, "-hws", "32:00:16:46:20:00", "-hws", "32:00:16:46:20:01")
	macA, err := net.ParseMAC("32:00:16:46:20:00")
	if err != nil {
		t.Fatal(err)
	}
	macB, err := net.ParseMAC("32:00:16:46:20:01")
	if err != nil {
		t.Fatal(err)
	}
	var hwFlagSliceExpected = []net.HardwareAddr{macA, macB}

	var maskFlag net.IPMask
	IPMask(&maskFlag, "m", "mFlag", "mask flag")
	inputArgs = append(inputArgs, "-m", "255.255.255.255")
	var maskFlagExpected = net.IPMask([]byte{255, 255, 255, 255})

	var maskSliceFlag []net.IPMask
	IPMaskSlice(&maskSliceFlag, "ms", "mFlagSlice", "mask slice flag")
	if err != nil {
		t.Fatal(err)
	}
	inputArgs = append(inputArgs, "-ms", "255.255.255.255", "-ms", "255.255.255.0")
	var maskSliceFlagExpected = []net.IPMask{net.IPMask([]byte{255, 255, 255, 255}), net.IPMask([]byte{255, 255, 255, 0})}

	// display help with all flags used
	ShowHelp("Showing help for test: " + t.Name())

	// Parse arguments
	ParseArgs(inputArgs)

	// validate parsed values
	if stringFlag != stringFlagExpected {
		t.Fatal("string flag incorrect", stringFlag, stringFlagExpected)
	}

	for i, f := range stringSliceFlagExpected {
		if stringSliceFlag[i] != f {
			t.Fatal("stringSlice value incorrect", stringSliceFlag[i], f)
		}
	}

	if boolFlag != boolFlagExpected {
		t.Fatal("bool flag incorrect", boolFlag, boolFlagExpected)
	}

	for i, f := range boolSliceFlagExpected {
		if boolSliceFlag[i] != f {
			t.Fatal("boolSlice value incorrect", boolSliceFlag[i], f)
		}
	}

	for i, f := range byteSliceFlagExpected {
		if byteSliceFlag[i] != f {
			t.Fatal("byteSlice value incorrect", boolSliceFlag[i], f)
		}
	}

	if durationFlag != durationFlagExpected {
		t.Fatal("duration flag incorrect", durationFlag, durationFlagExpected)
	}

	for i, f := range durationSliceFlagExpected {
		if durationSliceFlag[i] != f {
			t.Fatal("durationSlice value incorrect", durationSliceFlag[i], f)
		}
	}

	if float32Flag != float32FlagExpected {
		t.Fatal("float32 flag incorrect", float32Flag, float32FlagExpected)
	}

	for i, f := range float32SliceFlagExpected {
		if float32SliceFlag[i] != f {
			t.Fatal("float32Slice value incorrect", float32SliceFlag[i], f)
		}
	}

	if float64Flag != float64FlagExpected {
		t.Fatal("float64 flag incorrect", float64Flag, float64FlagExpected)
	}

	for i, f := range float64SliceFlagExpected {
		if float64SliceFlag[i] != f {
			t.Fatal("float64Slice value incorrect", float64SliceFlag[i], f)
		}
	}

	if intFlag != intFlagExpected {
		t.Fatal("int flag incorrect", intFlag, intFlagExpected)
	}

	for i, f := range intSliceFlagExpected {
		if intSliceFlag[i] != f {
			t.Fatal("intSlice value incorrect", intSliceFlag[i], f)
		}
	}

	if uintFlag != uintFlagExpected {
		t.Fatal("uint flag incorrect", uintFlag, uintFlagExpected)
	}

	for i, f := range uintSliceFlagExpected {
		if uintSliceFlag[i] != f {
			t.Fatal("uintSlice value incorrect", uintSliceFlag[i], f)
		}
	}

	if uint64Flag != uint64FlagExpected {
		t.Fatal("uint64 flag incorrect", uint64Flag, uint64FlagExpected)
	}

	for i, f := range uint64SliceFlagExpected {
		if uint64SliceFlag[i] != f {
			t.Fatal("uint64Slice value incorrect", uint64SliceFlag[i], f)
		}
	}

	if uint32Flag != uint32FlagExpected {
		t.Fatal("uint32 flag incorrect", uint32Flag, uint32FlagExpected)
	}

	for i, f := range uint32SliceFlagExpected {
		if uint32SliceFlag[i] != f {
			t.Fatal("uint32Slice value incorrect", uint32SliceFlag[i], f)
		}
	}

	if uint16Flag != uint16FlagExpected {
		t.Fatal("uint16 flag incorrect", uint16Flag, uint16FlagExpected)
	}

	for i, f := range uint16SliceFlagExpected {
		if uint16SliceFlag[i] != f {
			t.Fatal("uint16Slice value incorrect", uint16SliceFlag[i], f)
		}
	}

	if uint8Flag != uint8FlagExpected {
		t.Fatal("uint8 flag incorrect", uint8Flag, uint8FlagExpected)
	}

	for i, f := range uint8SliceFlagExpected {
		if uint8SliceFlag[i] != f {
			t.Fatal("uint8Slice value", i, "incorrect", uint8SliceFlag[i], f)
		}
	}

	if int64Flag != int64FlagExpected {
		t.Fatal("int64 flag incorrect", int64Flag, int64FlagExpected)
	}

	for i, f := range int64SliceFlagExpected {
		if int64SliceFlag[i] != f {
			t.Fatal("int64Slice value incorrect", int64SliceFlag[i], f)
		}
	}

	if int32Flag != int32FlagExpected {
		t.Fatal("int32 flag incorrect", int32Flag, int32FlagExpected)
	}

	for i, f := range int32SliceFlagExpected {
		if int32SliceFlag[i] != f {
			t.Fatal("int32Slice value incorrect", int32SliceFlag[i], f)
		}
	}

	if int16Flag != int16FlagExpected {
		t.Fatal("int16 flag incorrect", int16Flag, int16FlagExpected)
	}

	for i, f := range int16SliceFlagExpected {
		if int16SliceFlag[i] != f {
			t.Fatal("int16Slice value incorrect", int16SliceFlag[i], f)
		}
	}

	if int8Flag != int8FlagExpected {
		t.Fatal("int8 flag incorrect", int8Flag, int8FlagExpected)
	}

	for i, f := range int8SliceFlagExpected {
		if int8SliceFlag[i] != f {
			t.Fatal("int8Slice value incorrect", int8SliceFlag[i], f)
		}
	}

	if !ipFlag.Equal(ipFlagExpected) {
		t.Fatal("ip flag incorrect", ipFlag, ipFlagExpected)
	}

	for i, f := range ipSliceFlagExpected {
		if !f.Equal(ipSliceFlag[i]) {
			t.Fatal("ipSlice value incorrect", ipSliceFlag[i], f)
		}
	}

	if hwFlag.String() != hwFlagExpected.String() {
		t.Fatal("hw flag incorrect", hwFlag, hwFlagExpected)
	}

	for i, f := range hwFlagSliceExpected {
		if f.String() != hwFlagSlice[i].String() {
			t.Fatal("hw flag slice value incorrect", hwFlagSlice[i].String(), f.String())
		}
	}

	if maskFlag.String() != maskFlagExpected.String() {
		t.Fatal("mask flag incorrect", maskFlag, maskFlagExpected)
	}

	for i, f := range maskSliceFlagExpected {
		if f.String() != maskSliceFlag[i].String() {
			t.Fatal("mask flag slice value incorrect", maskSliceFlag[i].String(), f.String())
		}
	}
}
