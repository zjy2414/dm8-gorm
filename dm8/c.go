/*
 * Copyright (c) 2000-2018, 达梦数据库有限公司.
 * All rights reserved.
 */
package dm8

import (
	"io"
	"math"
)

type Dm_build_1357 struct {
	dm_build_1358 []byte
	dm_build_1359 int
}

func Dm_build_1360(dm_build_1361 int) *Dm_build_1357 {
	return &Dm_build_1357{make([]byte, 0, dm_build_1361), 0}
}

func Dm_build_1362(dm_build_1363 []byte) *Dm_build_1357 {
	return &Dm_build_1357{dm_build_1363, 0}
}

func (dm_build_1365 *Dm_build_1357) dm_build_1364(dm_build_1366 int) *Dm_build_1357 {

	dm_build_1367 := len(dm_build_1365.dm_build_1358)
	dm_build_1368 := cap(dm_build_1365.dm_build_1358)

	if dm_build_1367+dm_build_1366 <= dm_build_1368 {
		dm_build_1365.dm_build_1358 = dm_build_1365.dm_build_1358[:dm_build_1367+dm_build_1366]
	} else {

		var calCap = int64(math.Max(float64(2*dm_build_1368), float64(dm_build_1366+dm_build_1367)))

		nbuf := make([]byte, dm_build_1366+dm_build_1367, calCap)
		copy(nbuf, dm_build_1365.dm_build_1358)
		dm_build_1365.dm_build_1358 = nbuf
	}

	return dm_build_1365
}

func (dm_build_1370 *Dm_build_1357) Dm_build_1369() int {
	return len(dm_build_1370.dm_build_1358)
}

func (dm_build_1372 *Dm_build_1357) Dm_build_1371(dm_build_1373 int) *Dm_build_1357 {
	for i := dm_build_1373; i < len(dm_build_1372.dm_build_1358); i++ {
		dm_build_1372.dm_build_1358[i] = 0
	}
	dm_build_1372.dm_build_1358 = dm_build_1372.dm_build_1358[:dm_build_1373]
	return dm_build_1372
}

func (dm_build_1375 *Dm_build_1357) Dm_build_1374(dm_build_1376 int) *Dm_build_1357 {
	dm_build_1375.dm_build_1359 = dm_build_1376
	return dm_build_1375
}

func (dm_build_1378 *Dm_build_1357) Dm_build_1377() int {
	return dm_build_1378.dm_build_1359
}

func (dm_build_1380 *Dm_build_1357) Dm_build_1379(dm_build_1381 bool) int {
	return len(dm_build_1380.dm_build_1358) - dm_build_1380.dm_build_1359
}

func (dm_build_1383 *Dm_build_1357) Dm_build_1382(dm_build_1384 int, dm_build_1385 bool, dm_build_1386 bool) *Dm_build_1357 {

	if dm_build_1385 {
		if dm_build_1386 {
			dm_build_1383.dm_build_1364(dm_build_1384)
		} else {
			dm_build_1383.dm_build_1358 = dm_build_1383.dm_build_1358[:len(dm_build_1383.dm_build_1358)-dm_build_1384]
		}
	} else {
		if dm_build_1386 {
			dm_build_1383.dm_build_1359 += dm_build_1384
		} else {
			dm_build_1383.dm_build_1359 -= dm_build_1384
		}
	}

	return dm_build_1383
}

func (dm_build_1388 *Dm_build_1357) Dm_build_1387(dm_build_1389 io.Reader, dm_build_1390 int) (int, error) {
	dm_build_1391 := len(dm_build_1388.dm_build_1358)
	dm_build_1388.dm_build_1364(dm_build_1390)
	dm_build_1392 := 0
	for dm_build_1390 > 0 {
		n, err := dm_build_1389.Read(dm_build_1388.dm_build_1358[dm_build_1391+dm_build_1392:])
		if n > 0 && err == io.EOF {
			dm_build_1392 += n
			dm_build_1388.dm_build_1358 = dm_build_1388.dm_build_1358[:dm_build_1391+dm_build_1392]
			return dm_build_1392, nil
		} else if n > 0 && err == nil {
			dm_build_1390 -= n
			dm_build_1392 += n
		} else if n == 0 && err != nil {
			return -1, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
		}
	}

	return dm_build_1392, nil
}

func (dm_build_1394 *Dm_build_1357) Dm_build_1393(dm_build_1395 io.Writer) (*Dm_build_1357, error) {
	if _, err := dm_build_1395.Write(dm_build_1394.dm_build_1358); err != nil {
		return nil, ECGO_COMMUNITION_ERROR.addDetailln(err.Error()).throw()
	}
	return dm_build_1394, nil
}

func (dm_build_1397 *Dm_build_1357) Dm_build_1396(dm_build_1398 bool) int {
	dm_build_1399 := len(dm_build_1397.dm_build_1358)
	dm_build_1397.dm_build_1364(1)

	if dm_build_1398 {
		return copy(dm_build_1397.dm_build_1358[dm_build_1399:], []byte{1})
	} else {
		return copy(dm_build_1397.dm_build_1358[dm_build_1399:], []byte{0})
	}
}

func (dm_build_1401 *Dm_build_1357) Dm_build_1400(dm_build_1402 byte) int {
	dm_build_1403 := len(dm_build_1401.dm_build_1358)
	dm_build_1401.dm_build_1364(1)

	return copy(dm_build_1401.dm_build_1358[dm_build_1403:], Dm_build_998.Dm_build_1176(dm_build_1402))
}

func (dm_build_1405 *Dm_build_1357) Dm_build_1404(dm_build_1406 int8) int {
	dm_build_1407 := len(dm_build_1405.dm_build_1358)
	dm_build_1405.dm_build_1364(1)

	return copy(dm_build_1405.dm_build_1358[dm_build_1407:], Dm_build_998.Dm_build_1179(dm_build_1406))
}

func (dm_build_1409 *Dm_build_1357) Dm_build_1408(dm_build_1410 int16) int {
	dm_build_1411 := len(dm_build_1409.dm_build_1358)
	dm_build_1409.dm_build_1364(2)

	return copy(dm_build_1409.dm_build_1358[dm_build_1411:], Dm_build_998.Dm_build_1182(dm_build_1410))
}

func (dm_build_1413 *Dm_build_1357) Dm_build_1412(dm_build_1414 int32) int {
	dm_build_1415 := len(dm_build_1413.dm_build_1358)
	dm_build_1413.dm_build_1364(4)

	return copy(dm_build_1413.dm_build_1358[dm_build_1415:], Dm_build_998.Dm_build_1185(dm_build_1414))
}

func (dm_build_1417 *Dm_build_1357) Dm_build_1416(dm_build_1418 uint8) int {
	dm_build_1419 := len(dm_build_1417.dm_build_1358)
	dm_build_1417.dm_build_1364(1)

	return copy(dm_build_1417.dm_build_1358[dm_build_1419:], Dm_build_998.Dm_build_1197(dm_build_1418))
}

func (dm_build_1421 *Dm_build_1357) Dm_build_1420(dm_build_1422 uint16) int {
	dm_build_1423 := len(dm_build_1421.dm_build_1358)
	dm_build_1421.dm_build_1364(2)

	return copy(dm_build_1421.dm_build_1358[dm_build_1423:], Dm_build_998.Dm_build_1200(dm_build_1422))
}

func (dm_build_1425 *Dm_build_1357) Dm_build_1424(dm_build_1426 uint32) int {
	dm_build_1427 := len(dm_build_1425.dm_build_1358)
	dm_build_1425.dm_build_1364(4)

	return copy(dm_build_1425.dm_build_1358[dm_build_1427:], Dm_build_998.Dm_build_1203(dm_build_1426))
}

func (dm_build_1429 *Dm_build_1357) Dm_build_1428(dm_build_1430 uint64) int {
	dm_build_1431 := len(dm_build_1429.dm_build_1358)
	dm_build_1429.dm_build_1364(8)

	return copy(dm_build_1429.dm_build_1358[dm_build_1431:], Dm_build_998.Dm_build_1206(dm_build_1430))
}

func (dm_build_1433 *Dm_build_1357) Dm_build_1432(dm_build_1434 float32) int {
	dm_build_1435 := len(dm_build_1433.dm_build_1358)
	dm_build_1433.dm_build_1364(4)

	return copy(dm_build_1433.dm_build_1358[dm_build_1435:], Dm_build_998.Dm_build_1203(math.Float32bits(dm_build_1434)))
}

func (dm_build_1437 *Dm_build_1357) Dm_build_1436(dm_build_1438 float64) int {
	dm_build_1439 := len(dm_build_1437.dm_build_1358)
	dm_build_1437.dm_build_1364(8)

	return copy(dm_build_1437.dm_build_1358[dm_build_1439:], Dm_build_998.Dm_build_1206(math.Float64bits(dm_build_1438)))
}

func (dm_build_1441 *Dm_build_1357) Dm_build_1440(dm_build_1442 []byte) int {
	dm_build_1443 := len(dm_build_1441.dm_build_1358)
	dm_build_1441.dm_build_1364(len(dm_build_1442))
	return copy(dm_build_1441.dm_build_1358[dm_build_1443:], dm_build_1442)
}

func (dm_build_1445 *Dm_build_1357) Dm_build_1444(dm_build_1446 []byte) int {
	return dm_build_1445.Dm_build_1412(int32(len(dm_build_1446))) + dm_build_1445.Dm_build_1440(dm_build_1446)
}

func (dm_build_1448 *Dm_build_1357) Dm_build_1447(dm_build_1449 []byte) int {
	return dm_build_1448.Dm_build_1416(uint8(len(dm_build_1449))) + dm_build_1448.Dm_build_1440(dm_build_1449)
}

func (dm_build_1451 *Dm_build_1357) Dm_build_1450(dm_build_1452 []byte) int {
	return dm_build_1451.Dm_build_1420(uint16(len(dm_build_1452))) + dm_build_1451.Dm_build_1440(dm_build_1452)
}

func (dm_build_1454 *Dm_build_1357) Dm_build_1453(dm_build_1455 []byte) int {
	return dm_build_1454.Dm_build_1440(dm_build_1455) + dm_build_1454.Dm_build_1400(0)
}

func (dm_build_1457 *Dm_build_1357) Dm_build_1456(dm_build_1458 string, dm_build_1459 string, dm_build_1460 *DmConnection) int {
	dm_build_1461 := Dm_build_998.Dm_build_1214(dm_build_1458, dm_build_1459, dm_build_1460)
	return dm_build_1457.Dm_build_1444(dm_build_1461)
}

func (dm_build_1463 *Dm_build_1357) Dm_build_1462(dm_build_1464 string, dm_build_1465 string, dm_build_1466 *DmConnection) int {
	dm_build_1467 := Dm_build_998.Dm_build_1214(dm_build_1464, dm_build_1465, dm_build_1466)
	return dm_build_1463.Dm_build_1447(dm_build_1467)
}

func (dm_build_1469 *Dm_build_1357) Dm_build_1468(dm_build_1470 string, dm_build_1471 string, dm_build_1472 *DmConnection) int {
	dm_build_1473 := Dm_build_998.Dm_build_1214(dm_build_1470, dm_build_1471, dm_build_1472)
	return dm_build_1469.Dm_build_1450(dm_build_1473)
}

func (dm_build_1475 *Dm_build_1357) Dm_build_1474(dm_build_1476 string, dm_build_1477 string, dm_build_1478 *DmConnection) int {
	dm_build_1479 := Dm_build_998.Dm_build_1214(dm_build_1476, dm_build_1477, dm_build_1478)
	return dm_build_1475.Dm_build_1453(dm_build_1479)
}

func (dm_build_1481 *Dm_build_1357) Dm_build_1480() byte {
	dm_build_1482 := Dm_build_998.Dm_build_1091(dm_build_1481.dm_build_1358, dm_build_1481.dm_build_1359)
	dm_build_1481.dm_build_1359++
	return dm_build_1482
}

func (dm_build_1484 *Dm_build_1357) Dm_build_1483() int16 {
	dm_build_1485 := Dm_build_998.Dm_build_1095(dm_build_1484.dm_build_1358, dm_build_1484.dm_build_1359)
	dm_build_1484.dm_build_1359 += 2
	return dm_build_1485
}

func (dm_build_1487 *Dm_build_1357) Dm_build_1486() int32 {
	dm_build_1488 := Dm_build_998.Dm_build_1100(dm_build_1487.dm_build_1358, dm_build_1487.dm_build_1359)
	dm_build_1487.dm_build_1359 += 4
	return dm_build_1488
}

func (dm_build_1490 *Dm_build_1357) Dm_build_1489() int64 {
	dm_build_1491 := Dm_build_998.Dm_build_1105(dm_build_1490.dm_build_1358, dm_build_1490.dm_build_1359)
	dm_build_1490.dm_build_1359 += 8
	return dm_build_1491
}

func (dm_build_1493 *Dm_build_1357) Dm_build_1492() float32 {
	dm_build_1494 := Dm_build_998.Dm_build_1110(dm_build_1493.dm_build_1358, dm_build_1493.dm_build_1359)
	dm_build_1493.dm_build_1359 += 4
	return dm_build_1494
}

func (dm_build_1496 *Dm_build_1357) Dm_build_1495() float64 {
	dm_build_1497 := Dm_build_998.Dm_build_1114(dm_build_1496.dm_build_1358, dm_build_1496.dm_build_1359)
	dm_build_1496.dm_build_1359 += 8
	return dm_build_1497
}

func (dm_build_1499 *Dm_build_1357) Dm_build_1498() uint8 {
	dm_build_1500 := Dm_build_998.Dm_build_1118(dm_build_1499.dm_build_1358, dm_build_1499.dm_build_1359)
	dm_build_1499.dm_build_1359 += 1
	return dm_build_1500
}

func (dm_build_1502 *Dm_build_1357) Dm_build_1501() uint16 {
	dm_build_1503 := Dm_build_998.Dm_build_1122(dm_build_1502.dm_build_1358, dm_build_1502.dm_build_1359)
	dm_build_1502.dm_build_1359 += 2
	return dm_build_1503
}

func (dm_build_1505 *Dm_build_1357) Dm_build_1504() uint32 {
	dm_build_1506 := Dm_build_998.Dm_build_1127(dm_build_1505.dm_build_1358, dm_build_1505.dm_build_1359)
	dm_build_1505.dm_build_1359 += 4
	return dm_build_1506
}

func (dm_build_1508 *Dm_build_1357) Dm_build_1507(dm_build_1509 int) []byte {
	dm_build_1510 := Dm_build_998.Dm_build_1149(dm_build_1508.dm_build_1358, dm_build_1508.dm_build_1359, dm_build_1509)
	dm_build_1508.dm_build_1359 += dm_build_1509
	return dm_build_1510
}

func (dm_build_1512 *Dm_build_1357) Dm_build_1511() []byte {
	return dm_build_1512.Dm_build_1507(int(dm_build_1512.Dm_build_1486()))
}

func (dm_build_1514 *Dm_build_1357) Dm_build_1513() []byte {
	return dm_build_1514.Dm_build_1507(int(dm_build_1514.Dm_build_1480()))
}

func (dm_build_1516 *Dm_build_1357) Dm_build_1515() []byte {
	return dm_build_1516.Dm_build_1507(int(dm_build_1516.Dm_build_1483()))
}

func (dm_build_1518 *Dm_build_1357) Dm_build_1517(dm_build_1519 int) []byte {
	return dm_build_1518.Dm_build_1507(dm_build_1519)
}

func (dm_build_1521 *Dm_build_1357) Dm_build_1520() []byte {
	dm_build_1522 := 0
	for dm_build_1521.Dm_build_1480() != 0 {
		dm_build_1522++
	}
	dm_build_1521.Dm_build_1382(dm_build_1522, false, false)
	return dm_build_1521.Dm_build_1507(dm_build_1522)
}

func (dm_build_1524 *Dm_build_1357) Dm_build_1523(dm_build_1525 int, dm_build_1526 string, dm_build_1527 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1524.Dm_build_1507(dm_build_1525), dm_build_1526, dm_build_1527)
}

func (dm_build_1529 *Dm_build_1357) Dm_build_1528(dm_build_1530 string, dm_build_1531 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1529.Dm_build_1511(), dm_build_1530, dm_build_1531)
}

func (dm_build_1533 *Dm_build_1357) Dm_build_1532(dm_build_1534 string, dm_build_1535 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1533.Dm_build_1513(), dm_build_1534, dm_build_1535)
}

func (dm_build_1537 *Dm_build_1357) Dm_build_1536(dm_build_1538 string, dm_build_1539 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1537.Dm_build_1515(), dm_build_1538, dm_build_1539)
}

func (dm_build_1541 *Dm_build_1357) Dm_build_1540(dm_build_1542 string, dm_build_1543 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1541.Dm_build_1520(), dm_build_1542, dm_build_1543)
}

func (dm_build_1545 *Dm_build_1357) Dm_build_1544(dm_build_1546 int, dm_build_1547 byte) int {
	return dm_build_1545.Dm_build_1580(dm_build_1546, Dm_build_998.Dm_build_1176(dm_build_1547))
}

func (dm_build_1549 *Dm_build_1357) Dm_build_1548(dm_build_1550 int, dm_build_1551 int16) int {
	return dm_build_1549.Dm_build_1580(dm_build_1550, Dm_build_998.Dm_build_1182(dm_build_1551))
}

func (dm_build_1553 *Dm_build_1357) Dm_build_1552(dm_build_1554 int, dm_build_1555 int32) int {
	return dm_build_1553.Dm_build_1580(dm_build_1554, Dm_build_998.Dm_build_1185(dm_build_1555))
}

func (dm_build_1557 *Dm_build_1357) Dm_build_1556(dm_build_1558 int, dm_build_1559 int64) int {
	return dm_build_1557.Dm_build_1580(dm_build_1558, Dm_build_998.Dm_build_1188(dm_build_1559))
}

func (dm_build_1561 *Dm_build_1357) Dm_build_1560(dm_build_1562 int, dm_build_1563 float32) int {
	return dm_build_1561.Dm_build_1580(dm_build_1562, Dm_build_998.Dm_build_1191(dm_build_1563))
}

func (dm_build_1565 *Dm_build_1357) Dm_build_1564(dm_build_1566 int, dm_build_1567 float64) int {
	return dm_build_1565.Dm_build_1580(dm_build_1566, Dm_build_998.Dm_build_1194(dm_build_1567))
}

func (dm_build_1569 *Dm_build_1357) Dm_build_1568(dm_build_1570 int, dm_build_1571 uint8) int {
	return dm_build_1569.Dm_build_1580(dm_build_1570, Dm_build_998.Dm_build_1197(dm_build_1571))
}

func (dm_build_1573 *Dm_build_1357) Dm_build_1572(dm_build_1574 int, dm_build_1575 uint16) int {
	return dm_build_1573.Dm_build_1580(dm_build_1574, Dm_build_998.Dm_build_1200(dm_build_1575))
}

func (dm_build_1577 *Dm_build_1357) Dm_build_1576(dm_build_1578 int, dm_build_1579 uint32) int {
	return dm_build_1577.Dm_build_1580(dm_build_1578, Dm_build_998.Dm_build_1203(dm_build_1579))
}

func (dm_build_1581 *Dm_build_1357) Dm_build_1580(dm_build_1582 int, dm_build_1583 []byte) int {
	return copy(dm_build_1581.dm_build_1358[dm_build_1582:], dm_build_1583)
}

func (dm_build_1585 *Dm_build_1357) Dm_build_1584(dm_build_1586 int, dm_build_1587 []byte) int {
	return dm_build_1585.Dm_build_1552(dm_build_1586, int32(len(dm_build_1587))) + dm_build_1585.Dm_build_1580(dm_build_1586+4, dm_build_1587)
}

func (dm_build_1589 *Dm_build_1357) Dm_build_1588(dm_build_1590 int, dm_build_1591 []byte) int {
	return dm_build_1589.Dm_build_1544(dm_build_1590, byte(len(dm_build_1591))) + dm_build_1589.Dm_build_1580(dm_build_1590+1, dm_build_1591)
}

func (dm_build_1593 *Dm_build_1357) Dm_build_1592(dm_build_1594 int, dm_build_1595 []byte) int {
	return dm_build_1593.Dm_build_1548(dm_build_1594, int16(len(dm_build_1595))) + dm_build_1593.Dm_build_1580(dm_build_1594+2, dm_build_1595)
}

func (dm_build_1597 *Dm_build_1357) Dm_build_1596(dm_build_1598 int, dm_build_1599 []byte) int {
	return dm_build_1597.Dm_build_1580(dm_build_1598, dm_build_1599) + dm_build_1597.Dm_build_1544(dm_build_1598+len(dm_build_1599), 0)
}

func (dm_build_1601 *Dm_build_1357) Dm_build_1600(dm_build_1602 int, dm_build_1603 string, dm_build_1604 string, dm_build_1605 *DmConnection) int {
	return dm_build_1601.Dm_build_1584(dm_build_1602, Dm_build_998.Dm_build_1214(dm_build_1603, dm_build_1604, dm_build_1605))
}

func (dm_build_1607 *Dm_build_1357) Dm_build_1606(dm_build_1608 int, dm_build_1609 string, dm_build_1610 string, dm_build_1611 *DmConnection) int {
	return dm_build_1607.Dm_build_1588(dm_build_1608, Dm_build_998.Dm_build_1214(dm_build_1609, dm_build_1610, dm_build_1611))
}

func (dm_build_1613 *Dm_build_1357) Dm_build_1612(dm_build_1614 int, dm_build_1615 string, dm_build_1616 string, dm_build_1617 *DmConnection) int {
	return dm_build_1613.Dm_build_1592(dm_build_1614, Dm_build_998.Dm_build_1214(dm_build_1615, dm_build_1616, dm_build_1617))
}

func (dm_build_1619 *Dm_build_1357) Dm_build_1618(dm_build_1620 int, dm_build_1621 string, dm_build_1622 string, dm_build_1623 *DmConnection) int {
	return dm_build_1619.Dm_build_1596(dm_build_1620, Dm_build_998.Dm_build_1214(dm_build_1621, dm_build_1622, dm_build_1623))
}

func (dm_build_1625 *Dm_build_1357) Dm_build_1624(dm_build_1626 int) byte {
	return Dm_build_998.Dm_build_1219(dm_build_1625.Dm_build_1651(dm_build_1626, 1))
}

func (dm_build_1628 *Dm_build_1357) Dm_build_1627(dm_build_1629 int) int16 {
	return Dm_build_998.Dm_build_1222(dm_build_1628.Dm_build_1651(dm_build_1629, 2))
}

func (dm_build_1631 *Dm_build_1357) Dm_build_1630(dm_build_1632 int) int32 {
	return Dm_build_998.Dm_build_1225(dm_build_1631.Dm_build_1651(dm_build_1632, 4))
}

func (dm_build_1634 *Dm_build_1357) Dm_build_1633(dm_build_1635 int) int64 {
	return Dm_build_998.Dm_build_1228(dm_build_1634.Dm_build_1651(dm_build_1635, 8))
}

func (dm_build_1637 *Dm_build_1357) Dm_build_1636(dm_build_1638 int) float32 {
	return Dm_build_998.Dm_build_1231(dm_build_1637.Dm_build_1651(dm_build_1638, 4))
}

func (dm_build_1640 *Dm_build_1357) Dm_build_1639(dm_build_1641 int) float64 {
	return Dm_build_998.Dm_build_1234(dm_build_1640.Dm_build_1651(dm_build_1641, 8))
}

func (dm_build_1643 *Dm_build_1357) Dm_build_1642(dm_build_1644 int) uint8 {
	return Dm_build_998.Dm_build_1237(dm_build_1643.Dm_build_1651(dm_build_1644, 1))
}

func (dm_build_1646 *Dm_build_1357) Dm_build_1645(dm_build_1647 int) uint16 {
	return Dm_build_998.Dm_build_1240(dm_build_1646.Dm_build_1651(dm_build_1647, 2))
}

func (dm_build_1649 *Dm_build_1357) Dm_build_1648(dm_build_1650 int) uint32 {
	return Dm_build_998.Dm_build_1243(dm_build_1649.Dm_build_1651(dm_build_1650, 4))
}

func (dm_build_1652 *Dm_build_1357) Dm_build_1651(dm_build_1653 int, dm_build_1654 int) []byte {
	return dm_build_1652.dm_build_1358[dm_build_1653 : dm_build_1653+dm_build_1654]
}

func (dm_build_1656 *Dm_build_1357) Dm_build_1655(dm_build_1657 int) []byte {
	dm_build_1658 := dm_build_1656.Dm_build_1630(dm_build_1657)
	return dm_build_1656.Dm_build_1651(dm_build_1657+4, int(dm_build_1658))
}

func (dm_build_1660 *Dm_build_1357) Dm_build_1659(dm_build_1661 int) []byte {
	dm_build_1662 := dm_build_1660.Dm_build_1624(dm_build_1661)
	return dm_build_1660.Dm_build_1651(dm_build_1661+1, int(dm_build_1662))
}

func (dm_build_1664 *Dm_build_1357) Dm_build_1663(dm_build_1665 int) []byte {
	dm_build_1666 := dm_build_1664.Dm_build_1627(dm_build_1665)
	return dm_build_1664.Dm_build_1651(dm_build_1665+2, int(dm_build_1666))
}

func (dm_build_1668 *Dm_build_1357) Dm_build_1667(dm_build_1669 int) []byte {
	dm_build_1670 := 0
	for dm_build_1668.Dm_build_1624(dm_build_1669) != 0 {
		dm_build_1669++
		dm_build_1670++
	}

	return dm_build_1668.Dm_build_1651(dm_build_1669-dm_build_1670, int(dm_build_1670))
}

func (dm_build_1672 *Dm_build_1357) Dm_build_1671(dm_build_1673 int, dm_build_1674 string, dm_build_1675 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1672.Dm_build_1655(dm_build_1673), dm_build_1674, dm_build_1675)
}

func (dm_build_1677 *Dm_build_1357) Dm_build_1676(dm_build_1678 int, dm_build_1679 string, dm_build_1680 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1677.Dm_build_1659(dm_build_1678), dm_build_1679, dm_build_1680)
}

func (dm_build_1682 *Dm_build_1357) Dm_build_1681(dm_build_1683 int, dm_build_1684 string, dm_build_1685 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1682.Dm_build_1663(dm_build_1683), dm_build_1684, dm_build_1685)
}

func (dm_build_1687 *Dm_build_1357) Dm_build_1686(dm_build_1688 int, dm_build_1689 string, dm_build_1690 *DmConnection) string {
	return Dm_build_998.Dm_build_1250(dm_build_1687.Dm_build_1667(dm_build_1688), dm_build_1689, dm_build_1690)
}
