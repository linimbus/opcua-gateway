package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func VersionGet() string {
	return "v1.0.0"
}

func AppNameGet() string {
	return "OPCUA Gateway Windows " + VersionGet()
}

func CompanyGet() string {
	return "linimbus@126.com"
}

func SwitchName(flag bool) string {
	if flag {
		return "Yes"
	}
	return "No"
}

func TimeStampGet() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func DatetimeToTime(opcuaTime uint64) time.Time {
	ts := (opcuaTime - uint64(116444736000000000))
	return time.Unix(int64(ts)/10000000, int64(ts%10000000)*100).Local()
}

func DatetimeToString(opcuaTime uint64) string {
	return DatetimeToTime(opcuaTime).Format("2006-01-02T15:04:05.000000")
}

func DefaultFont() Font {
	return Font{Family: "Segoe UI", PointSize: 9}
}

func SaveToFile(name string, body []byte) error {
	err := os.WriteFile(name, body, 0664)
	if err != nil {
		logs.Error("save file %s failed, %s", name, err.Error())
	}
	return err
}

func InterfaceGet(iface *net.Interface) ([]net.IP, error) {
	addrs, err := iface.Addrs()
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0)
	for _, v := range addrs {
		ipone, _, err := net.ParseCIDR(v.String())
		if err != nil {
			continue
		}
		if len(ipone) > 0 {
			ips = append(ips, ipone)
		}
	}
	if len(ips) == 0 {
		return nil, fmt.Errorf("interface not any address")
	}
	return ips, nil
}

func InterfaceOptions() []string {
	output := []string{"0.0.0.0", "::"}
	ifaces, err := net.Interfaces()
	if err != nil {
		logs.Error("interface query failed, %s", err.Error())
		return output
	}
	for _, v := range ifaces {
		if v.Flags&net.FlagUp == 0 {
			continue
		}
		address, err := InterfaceGet(&v)
		if err != nil {
			logs.Warning("interface get failed, %s", err.Error())
			continue
		}
		for _, addr := range address {
			output = append(output, addr.String())
		}
	}
	return output
}

func ListenTest(addr string, port int) error {
	address := fmt.Sprintf("%s:%d", addr, port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		logs.Error("listen [%s] failed: %s", address, err.Error())
		return err
	}
	listen.Close()
	return nil
}

func CapSignal(proc func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		proc()
		logs.Error("recv signcal %s, ready to exit", sig.String())
		os.Exit(-1)
	}()
}

func CopyClipboard() (string, error) {
	text, err := walk.Clipboard().Text()
	if err != nil {
		logs.Error(err.Error())
		return "", fmt.Errorf("can not find the any clipboard")
	}
	return text, nil
}

func PasteClipboard(input string) error {
	err := walk.Clipboard().SetText(input)
	if err != nil {
		logs.Error(err.Error())
	}
	return err
}

var constEscapeMaps = map[rune]rune{
	'-': '_', '?': '_', '!': '_', ':': '_', ';': '_', '&': '_', '^': '_',
	'(': '_', ')': '_', '#': '_', '@': '_', '/': '_', '\\': '_', '"': '_',
	'\'': '_', '<': '_', '>': '_', '{': '_', '}': '_', '*': '_', ' ': '_',
	'\r': '_', '\n': '_', '\t': '_', '|': '_', '=': '_', '+': '_', '.': '$',
}

func EscapeString(str string) string {
	output := make([]rune, 0)
	for _, c := range str {
		value, ok := constEscapeMaps[c]
		if ok {
			output = append(output, rune(value))
		} else {
			output = append(output, c)
		}
	}
	return string(output)
}

func ColumnName(str string) string {
	name := EscapeString(str)
	if len(name) < 60 {
		return name
	}
	return "_" + name[len(name)-60:]
}

func SmallLength(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BoolListToString(values []bool) string {
	var buffer bytes.Buffer
	for i, v := range values {
		if v {
			buffer.WriteString("1")
		} else {
			buffer.WriteString("0")
		}
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func BoolToString(value bool) string {
	if value {
		return "1"
	}
	return "0"
}

func Int8ListToString(values []int8) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", int16(v)))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Uint8ListToString(values []uint8) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", uint16(v)))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Int16ListToString(values []int16) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Uint16ListToString(values []uint16) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Int32ListToString(values []int32) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Uint32ListToString(values []uint32) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Int64ListToString(values []int64) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func Uint64ListToString(values []uint64) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%d", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func DateTimeListToString(values []uint64) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(DatetimeToString(v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func FloatListToString(values []float32) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%0.5f", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func DoubleListToString(values []float64) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(fmt.Sprintf("%0.5f", v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func StringListToString(values []string) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(v)
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func ByteListToString(values [][]byte) string {
	var buffer bytes.Buffer
	for i, v := range values {
		buffer.WriteString(base64.StdEncoding.EncodeToString(v))
		if i+1 != len(values) {
			buffer.WriteString(",")
		}
	}
	return buffer.String()
}

func StringToBool(value string) bool {
	if value == "1" || value == "true" {
		return true
	}
	return false
}

func BoolListCompare(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Int8ListCompare(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Uint8ListCompare(a, b []uint8) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Int16ListCompare(a, b []int16) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Uint16ListCompare(a, b []uint16) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Int32ListCompare(a, b []int32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Uint32ListCompare(a, b []uint32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Int64ListCompare(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Uint64ListCompare(a, b []uint64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func StringListCompare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func FloatListCompare(a, b []float32) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func DoubleListCompare(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ByteListCompare(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !bytes.Equal(a[i], b[i]) {
			return false
		}
	}
	return true
}

func BoolListClone(a []bool) []bool {
	b := make([]bool, len(a))
	copy(b, a)
	return b
}

func Int8ListClone(a []int8) []int8 {
	b := make([]int8, len(a))
	copy(b, a)
	return b
}

func Uint8ListClone(a []uint8) []uint8 {
	b := make([]uint8, len(a))
	copy(b, a)
	return b
}

func Int16ListClone(a []int16) []int16 {
	b := make([]int16, len(a))
	copy(b, a)
	return b
}

func Uint16ListClone(a []uint16) []uint16 {
	b := make([]uint16, len(a))
	copy(b, a)
	return b
}

func Int32ListClone(a []int32) []int32 {
	b := make([]int32, len(a))
	copy(b, a)
	return b
}

func Uint32ListClone(a []uint32) []uint32 {
	b := make([]uint32, len(a))
	copy(b, a)
	return b
}

func Int64ListClone(a []int64) []int64 {
	b := make([]int64, len(a))
	copy(b, a)
	return b
}

func Uint64ListClone(a []uint64) []uint64 {
	b := make([]uint64, len(a))
	copy(b, a)
	return b
}

func FloatListClone(a []float32) []float32 {
	b := make([]float32, len(a))
	copy(b, a)
	return b
}

func DoubleListClone(a []float64) []float64 {
	b := make([]float64, len(a))
	copy(b, a)
	return b
}

func StringListClone(a []string) []string {
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func ByteListClone(a [][]byte) [][]byte {
	b := make([][]byte, len(a))
	for i := range a {
		b[i] = make([]byte, len(a[i]))
		copy(b[i], a[i])
	}
	return b
}

func ByteClone(a []byte) []byte {
	b := make([]byte, len(a))
	copy(b, a)
	return b
}
