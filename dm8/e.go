/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm8

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"math"
)

type dm_build_997 struct{}

var Dm_build_998 = &dm_build_997{}

func (Dm_build_1000 *dm_build_997) Dm_build_999(dm_build_1001 []byte, dm_build_1002 int, dm_build_1003 byte) int {
	dm_build_1001[dm_build_1002] = dm_build_1003
	return 1
}

func (Dm_build_1005 *dm_build_997) Dm_build_1004(dm_build_1006 []byte, dm_build_1007 int, dm_build_1008 int8) int {
	dm_build_1006[dm_build_1007] = byte(dm_build_1008)
	return 1
}

func (Dm_build_1010 *dm_build_997) Dm_build_1009(dm_build_1011 []byte, dm_build_1012 int, dm_build_1013 int16) int {
	dm_build_1011[dm_build_1012] = byte(dm_build_1013)
	dm_build_1012++
	dm_build_1011[dm_build_1012] = byte(dm_build_1013 >> 8)
	return 2
}

func (Dm_build_1015 *dm_build_997) Dm_build_1014(dm_build_1016 []byte, dm_build_1017 int, dm_build_1018 int32) int {
	dm_build_1016[dm_build_1017] = byte(dm_build_1018)
	dm_build_1017++
	dm_build_1016[dm_build_1017] = byte(dm_build_1018 >> 8)
	dm_build_1017++
	dm_build_1016[dm_build_1017] = byte(dm_build_1018 >> 16)
	dm_build_1017++
	dm_build_1016[dm_build_1017] = byte(dm_build_1018 >> 24)
	dm_build_1017++
	return 4
}

func (Dm_build_1020 *dm_build_997) Dm_build_1019(dm_build_1021 []byte, dm_build_1022 int, dm_build_1023 int64) int {
	dm_build_1021[dm_build_1022] = byte(dm_build_1023)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 8)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 16)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 24)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 32)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 40)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 48)
	dm_build_1022++
	dm_build_1021[dm_build_1022] = byte(dm_build_1023 >> 56)
	return 8
}

func (Dm_build_1025 *dm_build_997) Dm_build_1024(dm_build_1026 []byte, dm_build_1027 int, dm_build_1028 float32) int {
	return Dm_build_1025.Dm_build_1044(dm_build_1026, dm_build_1027, math.Float32bits(dm_build_1028))
}

func (Dm_build_1030 *dm_build_997) Dm_build_1029(dm_build_1031 []byte, dm_build_1032 int, dm_build_1033 float64) int {
	return Dm_build_1030.Dm_build_1049(dm_build_1031, dm_build_1032, math.Float64bits(dm_build_1033))
}

func (Dm_build_1035 *dm_build_997) Dm_build_1034(dm_build_1036 []byte, dm_build_1037 int, dm_build_1038 uint8) int {
	dm_build_1036[dm_build_1037] = byte(dm_build_1038)
	return 1
}

func (Dm_build_1040 *dm_build_997) Dm_build_1039(dm_build_1041 []byte, dm_build_1042 int, dm_build_1043 uint16) int {
	dm_build_1041[dm_build_1042] = byte(dm_build_1043)
	dm_build_1042++
	dm_build_1041[dm_build_1042] = byte(dm_build_1043 >> 8)
	return 2
}

func (Dm_build_1045 *dm_build_997) Dm_build_1044(dm_build_1046 []byte, dm_build_1047 int, dm_build_1048 uint32) int {
	dm_build_1046[dm_build_1047] = byte(dm_build_1048)
	dm_build_1047++
	dm_build_1046[dm_build_1047] = byte(dm_build_1048 >> 8)
	dm_build_1047++
	dm_build_1046[dm_build_1047] = byte(dm_build_1048 >> 16)
	dm_build_1047++
	dm_build_1046[dm_build_1047] = byte(dm_build_1048 >> 24)
	return 3
}

func (Dm_build_1050 *dm_build_997) Dm_build_1049(dm_build_1051 []byte, dm_build_1052 int, dm_build_1053 uint64) int {
	dm_build_1051[dm_build_1052] = byte(dm_build_1053)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 8)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 16)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 24)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 32)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 40)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 48)
	dm_build_1052++
	dm_build_1051[dm_build_1052] = byte(dm_build_1053 >> 56)
	return 3
}

func (Dm_build_1055 *dm_build_997) Dm_build_1054(dm_build_1056 []byte, dm_build_1057 int, dm_build_1058 []byte, dm_build_1059 int, dm_build_1060 int) int {
	copy(dm_build_1056[dm_build_1057:dm_build_1057+dm_build_1060], dm_build_1058[dm_build_1059:dm_build_1059+dm_build_1060])
	return dm_build_1060
}

func (Dm_build_1062 *dm_build_997) Dm_build_1061(dm_build_1063 []byte, dm_build_1064 int, dm_build_1065 []byte, dm_build_1066 int, dm_build_1067 int) int {
	dm_build_1064 += Dm_build_1062.Dm_build_1044(dm_build_1063, dm_build_1064, uint32(dm_build_1067))
	return 4 + Dm_build_1062.Dm_build_1054(dm_build_1063, dm_build_1064, dm_build_1065, dm_build_1066, dm_build_1067)
}

func (Dm_build_1069 *dm_build_997) Dm_build_1068(dm_build_1070 []byte, dm_build_1071 int, dm_build_1072 []byte, dm_build_1073 int, dm_build_1074 int) int {
	dm_build_1071 += Dm_build_1069.Dm_build_1039(dm_build_1070, dm_build_1071, uint16(dm_build_1074))
	return 2 + Dm_build_1069.Dm_build_1054(dm_build_1070, dm_build_1071, dm_build_1072, dm_build_1073, dm_build_1074)
}

func (Dm_build_1076 *dm_build_997) Dm_build_1075(dm_build_1077 []byte, dm_build_1078 int, dm_build_1079 string, dm_build_1080 string, dm_build_1081 *DmConnection) int {
	dm_build_1082 := Dm_build_1076.Dm_build_1214(dm_build_1079, dm_build_1080, dm_build_1081)
	dm_build_1078 += Dm_build_1076.Dm_build_1044(dm_build_1077, dm_build_1078, uint32(len(dm_build_1082)))
	return 4 + Dm_build_1076.Dm_build_1054(dm_build_1077, dm_build_1078, dm_build_1082, 0, len(dm_build_1082))
}

func (Dm_build_1084 *dm_build_997) Dm_build_1083(dm_build_1085 []byte, dm_build_1086 int, dm_build_1087 string, dm_build_1088 string, dm_build_1089 *DmConnection) int {
	dm_build_1090 := Dm_build_1084.Dm_build_1214(dm_build_1087, dm_build_1088, dm_build_1089)

	dm_build_1086 += Dm_build_1084.Dm_build_1039(dm_build_1085, dm_build_1086, uint16(len(dm_build_1090)))
	return 2 + Dm_build_1084.Dm_build_1054(dm_build_1085, dm_build_1086, dm_build_1090, 0, len(dm_build_1090))
}

func (Dm_build_1092 *dm_build_997) Dm_build_1091(dm_build_1093 []byte, dm_build_1094 int) byte {
	return dm_build_1093[dm_build_1094]
}

func (Dm_build_1096 *dm_build_997) Dm_build_1095(dm_build_1097 []byte, dm_build_1098 int) int16 {
	var dm_build_1099 int16
	dm_build_1099 = int16(dm_build_1097[dm_build_1098] & 0xff)
	dm_build_1098++
	dm_build_1099 |= int16(dm_build_1097[dm_build_1098]&0xff) << 8
	return dm_build_1099
}

func (Dm_build_1101 *dm_build_997) Dm_build_1100(dm_build_1102 []byte, dm_build_1103 int) int32 {
	var dm_build_1104 int32
	dm_build_1104 = int32(dm_build_1102[dm_build_1103] & 0xff)
	dm_build_1103++
	dm_build_1104 |= int32(dm_build_1102[dm_build_1103]&0xff) << 8
	dm_build_1103++
	dm_build_1104 |= int32(dm_build_1102[dm_build_1103]&0xff) << 16
	dm_build_1103++
	dm_build_1104 |= int32(dm_build_1102[dm_build_1103]&0xff) << 24
	return dm_build_1104
}

func (Dm_build_1106 *dm_build_997) Dm_build_1105(dm_build_1107 []byte, dm_build_1108 int) int64 {
	var dm_build_1109 int64
	dm_build_1109 = int64(dm_build_1107[dm_build_1108] & 0xff)
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 8
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 16
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 24
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 32
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 40
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 48
	dm_build_1108++
	dm_build_1109 |= int64(dm_build_1107[dm_build_1108]&0xff) << 56
	return dm_build_1109
}

func (Dm_build_1111 *dm_build_997) Dm_build_1110(dm_build_1112 []byte, dm_build_1113 int) float32 {
	return math.Float32frombits(Dm_build_1111.Dm_build_1127(dm_build_1112, dm_build_1113))
}

func (Dm_build_1115 *dm_build_997) Dm_build_1114(dm_build_1116 []byte, dm_build_1117 int) float64 {
	return math.Float64frombits(Dm_build_1115.Dm_build_1132(dm_build_1116, dm_build_1117))
}

func (Dm_build_1119 *dm_build_997) Dm_build_1118(dm_build_1120 []byte, dm_build_1121 int) uint8 {
	return uint8(dm_build_1120[dm_build_1121] & 0xff)
}

func (Dm_build_1123 *dm_build_997) Dm_build_1122(dm_build_1124 []byte, dm_build_1125 int) uint16 {
	var dm_build_1126 uint16
	dm_build_1126 = uint16(dm_build_1124[dm_build_1125] & 0xff)
	dm_build_1125++
	dm_build_1126 |= uint16(dm_build_1124[dm_build_1125]&0xff) << 8
	return dm_build_1126
}

func (Dm_build_1128 *dm_build_997) Dm_build_1127(dm_build_1129 []byte, dm_build_1130 int) uint32 {
	var dm_build_1131 uint32
	dm_build_1131 = uint32(dm_build_1129[dm_build_1130] & 0xff)
	dm_build_1130++
	dm_build_1131 |= uint32(dm_build_1129[dm_build_1130]&0xff) << 8
	dm_build_1130++
	dm_build_1131 |= uint32(dm_build_1129[dm_build_1130]&0xff) << 16
	dm_build_1130++
	dm_build_1131 |= uint32(dm_build_1129[dm_build_1130]&0xff) << 24
	return dm_build_1131
}

func (Dm_build_1133 *dm_build_997) Dm_build_1132(dm_build_1134 []byte, dm_build_1135 int) uint64 {
	var dm_build_1136 uint64
	dm_build_1136 = uint64(dm_build_1134[dm_build_1135] & 0xff)
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 8
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 16
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 24
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 32
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 40
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 48
	dm_build_1135++
	dm_build_1136 |= uint64(dm_build_1134[dm_build_1135]&0xff) << 56
	return dm_build_1136
}

func (Dm_build_1138 *dm_build_997) Dm_build_1137(dm_build_1139 []byte, dm_build_1140 int) []byte {
	dm_build_1141 := Dm_build_1138.Dm_build_1127(dm_build_1139, dm_build_1140)

	dm_build_1142 := make([]byte, dm_build_1141)
	copy(dm_build_1142[:int(dm_build_1141)], dm_build_1139[dm_build_1140+4:dm_build_1140+4+int(dm_build_1141)])
	return dm_build_1142
}

func (Dm_build_1144 *dm_build_997) Dm_build_1143(dm_build_1145 []byte, dm_build_1146 int) []byte {
	dm_build_1147 := Dm_build_1144.Dm_build_1122(dm_build_1145, dm_build_1146)

	dm_build_1148 := make([]byte, dm_build_1147)
	copy(dm_build_1148[:int(dm_build_1147)], dm_build_1145[dm_build_1146+2:dm_build_1146+2+int(dm_build_1147)])
	return dm_build_1148
}

func (Dm_build_1150 *dm_build_997) Dm_build_1149(dm_build_1151 []byte, dm_build_1152 int, dm_build_1153 int) []byte {

	dm_build_1154 := make([]byte, dm_build_1153)
	copy(dm_build_1154[:dm_build_1153], dm_build_1151[dm_build_1152:dm_build_1152+dm_build_1153])
	return dm_build_1154
}

func (Dm_build_1156 *dm_build_997) Dm_build_1155(dm_build_1157 []byte, dm_build_1158 int, dm_build_1159 int, dm_build_1160 string, dm_build_1161 *DmConnection) string {
	return Dm_build_1156.Dm_build_1250(dm_build_1157[dm_build_1158:dm_build_1158+dm_build_1159], dm_build_1160, dm_build_1161)
}

func (Dm_build_1163 *dm_build_997) Dm_build_1162(dm_build_1164 []byte, dm_build_1165 int, dm_build_1166 string, dm_build_1167 *DmConnection) string {
	dm_build_1168 := Dm_build_1163.Dm_build_1127(dm_build_1164, dm_build_1165)
	dm_build_1165 += 4
	return Dm_build_1163.Dm_build_1155(dm_build_1164, dm_build_1165, int(dm_build_1168), dm_build_1166, dm_build_1167)
}

func (Dm_build_1170 *dm_build_997) Dm_build_1169(dm_build_1171 []byte, dm_build_1172 int, dm_build_1173 string, dm_build_1174 *DmConnection) string {
	dm_build_1175 := Dm_build_1170.Dm_build_1122(dm_build_1171, dm_build_1172)
	dm_build_1172 += 2
	return Dm_build_1170.Dm_build_1155(dm_build_1171, dm_build_1172, int(dm_build_1175), dm_build_1173, dm_build_1174)
}

func (Dm_build_1177 *dm_build_997) Dm_build_1176(dm_build_1178 byte) []byte {
	return []byte{dm_build_1178}
}

func (Dm_build_1180 *dm_build_997) Dm_build_1179(dm_build_1181 int8) []byte {
	return []byte{byte(dm_build_1181)}
}

func (Dm_build_1183 *dm_build_997) Dm_build_1182(dm_build_1184 int16) []byte {
	return []byte{byte(dm_build_1184), byte(dm_build_1184 >> 8)}
}

func (Dm_build_1186 *dm_build_997) Dm_build_1185(dm_build_1187 int32) []byte {
	return []byte{byte(dm_build_1187), byte(dm_build_1187 >> 8), byte(dm_build_1187 >> 16), byte(dm_build_1187 >> 24)}
}

func (Dm_build_1189 *dm_build_997) Dm_build_1188(dm_build_1190 int64) []byte {
	return []byte{byte(dm_build_1190), byte(dm_build_1190 >> 8), byte(dm_build_1190 >> 16), byte(dm_build_1190 >> 24), byte(dm_build_1190 >> 32),
		byte(dm_build_1190 >> 40), byte(dm_build_1190 >> 48), byte(dm_build_1190 >> 56)}
}

func (Dm_build_1192 *dm_build_997) Dm_build_1191(dm_build_1193 float32) []byte {
	return Dm_build_1192.Dm_build_1203(math.Float32bits(dm_build_1193))
}

func (Dm_build_1195 *dm_build_997) Dm_build_1194(dm_build_1196 float64) []byte {
	return Dm_build_1195.Dm_build_1206(math.Float64bits(dm_build_1196))
}

func (Dm_build_1198 *dm_build_997) Dm_build_1197(dm_build_1199 uint8) []byte {
	return []byte{byte(dm_build_1199)}
}

func (Dm_build_1201 *dm_build_997) Dm_build_1200(dm_build_1202 uint16) []byte {
	return []byte{byte(dm_build_1202), byte(dm_build_1202 >> 8)}
}

func (Dm_build_1204 *dm_build_997) Dm_build_1203(dm_build_1205 uint32) []byte {
	return []byte{byte(dm_build_1205), byte(dm_build_1205 >> 8), byte(dm_build_1205 >> 16), byte(dm_build_1205 >> 24)}
}

func (Dm_build_1207 *dm_build_997) Dm_build_1206(dm_build_1208 uint64) []byte {
	return []byte{byte(dm_build_1208), byte(dm_build_1208 >> 8), byte(dm_build_1208 >> 16), byte(dm_build_1208 >> 24), byte(dm_build_1208 >> 32), byte(dm_build_1208 >> 40), byte(dm_build_1208 >> 48), byte(dm_build_1208 >> 56)}
}

func (Dm_build_1210 *dm_build_997) Dm_build_1209(dm_build_1211 []byte, dm_build_1212 string, dm_build_1213 *DmConnection) []byte {
	if dm_build_1212 == "UTF-8" {
		return dm_build_1211
	}

	if dm_build_1213 == nil {
		if e := dm_build_1255(dm_build_1212); e != nil {
			tmp, err := ioutil.ReadAll(
				transform.NewReader(bytes.NewReader(dm_build_1211), e.NewEncoder()),
			)
			if err != nil {
				panic("UTF8 To Charset error!")
			}

			return tmp
		}

		panic("Unsupported Charset!")
	}

	if dm_build_1213.encodeBuffer == nil {
		dm_build_1213.encodeBuffer = bytes.NewBuffer(nil)
		dm_build_1213.encode = dm_build_1255(dm_build_1213.getServerEncoding())
		dm_build_1213.transformReaderDst = make([]byte, 4096)
		dm_build_1213.transformReaderSrc = make([]byte, 4096)
	}

	if e := dm_build_1213.encode; e != nil {

		dm_build_1213.encodeBuffer.Reset()

		n, err := dm_build_1213.encodeBuffer.ReadFrom(
			Dm_build_1269(bytes.NewReader(dm_build_1211), e.NewEncoder(), dm_build_1213.transformReaderDst, dm_build_1213.transformReaderSrc),
		)
		if err != nil {
			panic("UTF8 To Charset error!")
		}
		var tmp = make([]byte, n)
		if _, err = dm_build_1213.encodeBuffer.Read(tmp); err != nil {
			panic("UTF8 To Charset error!")
		}
		return tmp
	}

	panic("Unsupported Charset!")
}

func (Dm_build_1215 *dm_build_997) Dm_build_1214(dm_build_1216 string, dm_build_1217 string, dm_build_1218 *DmConnection) []byte {
	return Dm_build_1215.Dm_build_1209([]byte(dm_build_1216), dm_build_1217, dm_build_1218)
}

func (Dm_build_1220 *dm_build_997) Dm_build_1219(dm_build_1221 []byte) byte {
	return Dm_build_1220.Dm_build_1091(dm_build_1221, 0)
}

func (Dm_build_1223 *dm_build_997) Dm_build_1222(dm_build_1224 []byte) int16 {
	return Dm_build_1223.Dm_build_1095(dm_build_1224, 0)
}

func (Dm_build_1226 *dm_build_997) Dm_build_1225(dm_build_1227 []byte) int32 {
	return Dm_build_1226.Dm_build_1100(dm_build_1227, 0)
}

func (Dm_build_1229 *dm_build_997) Dm_build_1228(dm_build_1230 []byte) int64 {
	return Dm_build_1229.Dm_build_1105(dm_build_1230, 0)
}

func (Dm_build_1232 *dm_build_997) Dm_build_1231(dm_build_1233 []byte) float32 {
	return Dm_build_1232.Dm_build_1110(dm_build_1233, 0)
}

func (Dm_build_1235 *dm_build_997) Dm_build_1234(dm_build_1236 []byte) float64 {
	return Dm_build_1235.Dm_build_1114(dm_build_1236, 0)
}

func (Dm_build_1238 *dm_build_997) Dm_build_1237(dm_build_1239 []byte) uint8 {
	return Dm_build_1238.Dm_build_1118(dm_build_1239, 0)
}

func (Dm_build_1241 *dm_build_997) Dm_build_1240(dm_build_1242 []byte) uint16 {
	return Dm_build_1241.Dm_build_1122(dm_build_1242, 0)
}

func (Dm_build_1244 *dm_build_997) Dm_build_1243(dm_build_1245 []byte) uint32 {
	return Dm_build_1244.Dm_build_1127(dm_build_1245, 0)
}

func (Dm_build_1247 *dm_build_997) Dm_build_1246(dm_build_1248 []byte, dm_build_1249 string) []byte {
	if dm_build_1249 == "UTF-8" {
		return dm_build_1248
	}

	if e := dm_build_1255(dm_build_1249); e != nil {

		tmp, err := ioutil.ReadAll(
			transform.NewReader(bytes.NewReader(dm_build_1248), e.NewDecoder()),
		)
		if err != nil {

			panic("Charset To UTF8 error!")
		}

		return tmp
	}

	panic("Unsupported Charset!")

}

func (Dm_build_1251 *dm_build_997) Dm_build_1250(dm_build_1252 []byte, dm_build_1253 string, dm_build_1254 *DmConnection) string {
	return string(Dm_build_1251.Dm_build_1246(dm_build_1252, dm_build_1253))
}

func dm_build_1255(dm_build_1256 string) encoding.Encoding {
	if e, err := ianaindex.MIB.Encoding(dm_build_1256); err == nil && e != nil {
		return e
	}
	return nil
}

type Dm_build_1257 struct {
	dm_build_1258 io.Reader
	dm_build_1259 transform.Transformer
	dm_build_1260 error

	dm_build_1261                []byte
	dm_build_1262, dm_build_1263 int

	dm_build_1264                []byte
	dm_build_1265, dm_build_1266 int

	dm_build_1267 bool
}

const dm_build_1268 = 4096

func Dm_build_1269(dm_build_1270 io.Reader, dm_build_1271 transform.Transformer, dm_build_1272 []byte, dm_build_1273 []byte) *Dm_build_1257 {
	dm_build_1271.Reset()
	return &Dm_build_1257{
		dm_build_1258: dm_build_1270,
		dm_build_1259: dm_build_1271,
		dm_build_1261: dm_build_1272,
		dm_build_1264: dm_build_1273,
	}
}

func (dm_build_1275 *Dm_build_1257) Read(dm_build_1276 []byte) (int, error) {
	dm_build_1277, dm_build_1278 := 0, error(nil)
	for {

		if dm_build_1275.dm_build_1262 != dm_build_1275.dm_build_1263 {
			dm_build_1277 = copy(dm_build_1276, dm_build_1275.dm_build_1261[dm_build_1275.dm_build_1262:dm_build_1275.dm_build_1263])
			dm_build_1275.dm_build_1262 += dm_build_1277
			if dm_build_1275.dm_build_1262 == dm_build_1275.dm_build_1263 && dm_build_1275.dm_build_1267 {
				return dm_build_1277, dm_build_1275.dm_build_1260
			}
			return dm_build_1277, nil
		} else if dm_build_1275.dm_build_1267 {
			return 0, dm_build_1275.dm_build_1260
		}

		if dm_build_1275.dm_build_1265 != dm_build_1275.dm_build_1266 || dm_build_1275.dm_build_1260 != nil {
			dm_build_1275.dm_build_1262 = 0
			dm_build_1275.dm_build_1263, dm_build_1277, dm_build_1278 = dm_build_1275.dm_build_1259.Transform(dm_build_1275.dm_build_1261, dm_build_1275.dm_build_1264[dm_build_1275.dm_build_1265:dm_build_1275.dm_build_1266], dm_build_1275.dm_build_1260 == io.EOF)
			dm_build_1275.dm_build_1265 += dm_build_1277

			switch {
			case dm_build_1278 == nil:
				if dm_build_1275.dm_build_1265 != dm_build_1275.dm_build_1266 {
					dm_build_1275.dm_build_1260 = nil
				}

				dm_build_1275.dm_build_1267 = dm_build_1275.dm_build_1260 != nil
				continue
			case dm_build_1278 == transform.ErrShortDst && (dm_build_1275.dm_build_1263 != 0 || dm_build_1277 != 0):

				continue
			case dm_build_1278 == transform.ErrShortSrc && dm_build_1275.dm_build_1266-dm_build_1275.dm_build_1265 != len(dm_build_1275.dm_build_1264) && dm_build_1275.dm_build_1260 == nil:

			default:
				dm_build_1275.dm_build_1267 = true

				if dm_build_1275.dm_build_1260 == nil || dm_build_1275.dm_build_1260 == io.EOF {
					dm_build_1275.dm_build_1260 = dm_build_1278
				}
				continue
			}
		}

		if dm_build_1275.dm_build_1265 != 0 {
			dm_build_1275.dm_build_1265, dm_build_1275.dm_build_1266 = 0, copy(dm_build_1275.dm_build_1264, dm_build_1275.dm_build_1264[dm_build_1275.dm_build_1265:dm_build_1275.dm_build_1266])
		}
		dm_build_1277, dm_build_1275.dm_build_1260 = dm_build_1275.dm_build_1258.Read(dm_build_1275.dm_build_1264[dm_build_1275.dm_build_1266:])
		dm_build_1275.dm_build_1266 += dm_build_1277
	}
}
