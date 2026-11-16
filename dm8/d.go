/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm8

import (
	"container/list"
	"io"
)

type Dm_build_1279 struct {
	dm_build_1280 *list.List
	dm_build_1281 *dm_build_1333
	dm_build_1282 int
}

func Dm_build_1283() *Dm_build_1279 {
	return &Dm_build_1279{
		dm_build_1280: list.New(),
		dm_build_1282: 0,
	}
}

func (dm_build_1285 *Dm_build_1279) Dm_build_1284() int {
	return dm_build_1285.dm_build_1282
}

func (dm_build_1287 *Dm_build_1279) Dm_build_1286(dm_build_1288 *Dm_build_1357, dm_build_1289 int) int {
	var dm_build_1290 = 0
	var dm_build_1291 = 0
	for dm_build_1290 < dm_build_1289 && dm_build_1287.dm_build_1281 != nil {
		dm_build_1291 = dm_build_1287.dm_build_1281.dm_build_1341(dm_build_1288, dm_build_1289-dm_build_1290)
		if dm_build_1287.dm_build_1281.dm_build_1336 == 0 {
			dm_build_1287.dm_build_1323()
		}
		dm_build_1290 += dm_build_1291
		dm_build_1287.dm_build_1282 -= dm_build_1291
	}
	return dm_build_1290
}

func (dm_build_1293 *Dm_build_1279) Dm_build_1292(dm_build_1294 []byte, dm_build_1295 int, dm_build_1296 int) int {
	var dm_build_1297 = 0
	var dm_build_1298 = 0
	for dm_build_1297 < dm_build_1296 && dm_build_1293.dm_build_1281 != nil {
		dm_build_1298 = dm_build_1293.dm_build_1281.dm_build_1345(dm_build_1294, dm_build_1295, dm_build_1296-dm_build_1297)
		if dm_build_1293.dm_build_1281.dm_build_1336 == 0 {
			dm_build_1293.dm_build_1323()
		}
		dm_build_1297 += dm_build_1298
		dm_build_1293.dm_build_1282 -= dm_build_1298
		dm_build_1295 += dm_build_1298
	}
	return dm_build_1297
}

func (dm_build_1300 *Dm_build_1279) Dm_build_1299(dm_build_1301 io.Writer, dm_build_1302 int) int {
	var dm_build_1303 = 0
	var dm_build_1304 = 0
	for dm_build_1303 < dm_build_1302 && dm_build_1300.dm_build_1281 != nil {
		dm_build_1304 = dm_build_1300.dm_build_1281.dm_build_1350(dm_build_1301, dm_build_1302-dm_build_1303)
		if dm_build_1300.dm_build_1281.dm_build_1336 == 0 {
			dm_build_1300.dm_build_1323()
		}
		dm_build_1303 += dm_build_1304
		dm_build_1300.dm_build_1282 -= dm_build_1304
	}
	return dm_build_1303
}

func (dm_build_1306 *Dm_build_1279) Dm_build_1305(dm_build_1307 []byte, dm_build_1308 int, dm_build_1309 int) {
	if dm_build_1309 == 0 {
		return
	}
	var dm_build_1310 = dm_build_1337(dm_build_1307, dm_build_1308, dm_build_1309)
	if dm_build_1306.dm_build_1281 == nil {
		dm_build_1306.dm_build_1281 = dm_build_1310
	} else {
		dm_build_1306.dm_build_1280.PushBack(dm_build_1310)
	}
	dm_build_1306.dm_build_1282 += dm_build_1309
}

func (dm_build_1312 *Dm_build_1279) dm_build_1311(dm_build_1313 int) byte {
	var dm_build_1314 = dm_build_1313
	var dm_build_1315 = dm_build_1312.dm_build_1281
	for dm_build_1314 > 0 && dm_build_1315 != nil {
		if dm_build_1315.dm_build_1336 == 0 {
			continue
		}
		if dm_build_1314 > dm_build_1315.dm_build_1336-1 {
			dm_build_1314 -= dm_build_1315.dm_build_1336
			dm_build_1315 = dm_build_1312.dm_build_1280.Front().Value.(*dm_build_1333)
		} else {
			break
		}
	}
	return dm_build_1315.dm_build_1354(dm_build_1314)
}
func (dm_build_1317 *Dm_build_1279) Dm_build_1316(dm_build_1318 *Dm_build_1279) {
	if dm_build_1318.dm_build_1282 == 0 {
		return
	}
	var dm_build_1319 = dm_build_1318.dm_build_1281
	for dm_build_1319 != nil {
		dm_build_1317.dm_build_1320(dm_build_1319)
		dm_build_1318.dm_build_1323()
		dm_build_1319 = dm_build_1318.dm_build_1281
	}
	dm_build_1318.dm_build_1282 = 0
}
func (dm_build_1321 *Dm_build_1279) dm_build_1320(dm_build_1322 *dm_build_1333) {
	if dm_build_1322.dm_build_1336 == 0 {
		return
	}
	if dm_build_1321.dm_build_1281 == nil {
		dm_build_1321.dm_build_1281 = dm_build_1322
	} else {
		dm_build_1321.dm_build_1280.PushBack(dm_build_1322)
	}
	dm_build_1321.dm_build_1282 += dm_build_1322.dm_build_1336
}

func (dm_build_1324 *Dm_build_1279) dm_build_1323() {
	var dm_build_1325 = dm_build_1324.dm_build_1280.Front()
	if dm_build_1325 == nil {
		dm_build_1324.dm_build_1281 = nil
	} else {
		dm_build_1324.dm_build_1281 = dm_build_1325.Value.(*dm_build_1333)
		dm_build_1324.dm_build_1280.Remove(dm_build_1325)
	}
}

func (dm_build_1327 *Dm_build_1279) Dm_build_1326() []byte {
	var dm_build_1328 = make([]byte, dm_build_1327.dm_build_1282)
	var dm_build_1329 = dm_build_1327.dm_build_1281
	var dm_build_1330 = 0
	var dm_build_1331 = len(dm_build_1328)
	var dm_build_1332 = 0
	for dm_build_1329 != nil {
		if dm_build_1329.dm_build_1336 > 0 {
			if dm_build_1331 > dm_build_1329.dm_build_1336 {
				dm_build_1332 = dm_build_1329.dm_build_1336
			} else {
				dm_build_1332 = dm_build_1331
			}
			copy(dm_build_1328[dm_build_1330:dm_build_1330+dm_build_1332], dm_build_1329.dm_build_1334[dm_build_1329.dm_build_1335:dm_build_1329.dm_build_1335+dm_build_1332])
			dm_build_1330 += dm_build_1332
			dm_build_1331 -= dm_build_1332
		}
		if dm_build_1327.dm_build_1280.Front() == nil {
			dm_build_1329 = nil
		} else {
			dm_build_1329 = dm_build_1327.dm_build_1280.Front().Value.(*dm_build_1333)
		}
	}
	return dm_build_1328
}

type dm_build_1333 struct {
	dm_build_1334 []byte
	dm_build_1335 int
	dm_build_1336 int
}

func dm_build_1337(dm_build_1338 []byte, dm_build_1339 int, dm_build_1340 int) *dm_build_1333 {
	return &dm_build_1333{
		dm_build_1338,
		dm_build_1339,
		dm_build_1340,
	}
}

func (dm_build_1342 *dm_build_1333) dm_build_1341(dm_build_1343 *Dm_build_1357, dm_build_1344 int) int {
	if dm_build_1342.dm_build_1336 <= dm_build_1344 {
		dm_build_1344 = dm_build_1342.dm_build_1336
	}
	dm_build_1343.Dm_build_1440(dm_build_1342.dm_build_1334[dm_build_1342.dm_build_1335 : dm_build_1342.dm_build_1335+dm_build_1344])
	dm_build_1342.dm_build_1335 += dm_build_1344
	dm_build_1342.dm_build_1336 -= dm_build_1344
	return dm_build_1344
}

func (dm_build_1346 *dm_build_1333) dm_build_1345(dm_build_1347 []byte, dm_build_1348 int, dm_build_1349 int) int {
	if dm_build_1346.dm_build_1336 <= dm_build_1349 {
		dm_build_1349 = dm_build_1346.dm_build_1336
	}
	copy(dm_build_1347[dm_build_1348:dm_build_1348+dm_build_1349], dm_build_1346.dm_build_1334[dm_build_1346.dm_build_1335:dm_build_1346.dm_build_1335+dm_build_1349])
	dm_build_1346.dm_build_1335 += dm_build_1349
	dm_build_1346.dm_build_1336 -= dm_build_1349
	return dm_build_1349
}

func (dm_build_1351 *dm_build_1333) dm_build_1350(dm_build_1352 io.Writer, dm_build_1353 int) int {
	if dm_build_1351.dm_build_1336 <= dm_build_1353 {
		dm_build_1353 = dm_build_1351.dm_build_1336
	}
	dm_build_1352.Write(dm_build_1351.dm_build_1334[dm_build_1351.dm_build_1335 : dm_build_1351.dm_build_1335+dm_build_1353])
	dm_build_1351.dm_build_1335 += dm_build_1353
	dm_build_1351.dm_build_1336 -= dm_build_1353
	return dm_build_1353
}
func (dm_build_1355 *dm_build_1333) dm_build_1354(dm_build_1356 int) byte {
	return dm_build_1355.dm_build_1334[dm_build_1355.dm_build_1335+dm_build_1356]
}
