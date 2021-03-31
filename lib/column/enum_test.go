package column

import (
	"fmt"
	"testing"
)

func TestParseE(t *testing.T) {
	var s = `'' = -128, 'Web' = -127, 'Panasonic' = -126, 'samsung' = -125, 'LG' = -124, 'Linux' = -123, 'N/A' = -122, 'iOS' = -121, 'Android OS' = -120, 'Samsung' = -119, 'Windows 10' = -118, 'Mac OS' = -117, 'Sony' = -116, 'OnePlus' = -115, 'Philips' = -114, 'Windows 8.1' = -113, 'LENOVO' = -112, 'lg' = -111, 'motorola' = -110, 'HUAWEI' = -109, 'PHILIPS' = -108, 'Windows 7' = -107, 'HONOR' = -106, 'Lenovo' = -105, 'NVIDIA' = -104, 'Technicolor' = -103, 'NULL' = -102, 'HTC' = -101, 'htc' = -100, 'Google' = -99, 'Sagemcom' = -98, 'asus' = -97, 'Fairphone' = -96, 'google' = -95, 'LGE' = -94, 'Xiaomi' = -93, 'BullittGroupLimited' = -92, 'OUKITEL' = -91, 'ZTE' = -90, 'Huawei' = -89, 'Windows 8' = -88, 'Ulefone' = -87, 'null' = -86, 'lge' = -85, 'HMDGlobal' = -84, 'Nokia' = -83, 'fengmi' = -82, 'minix' = -81, 'TCL' = -80, 'Redmi' = -79, 'HMD Global' = -78, 'Vestel' = -77, 'ZUK' = -76, 'XGIMI' = -75, 'LandRover' = -74, 'rockchip' = -73, 'DOOGEE' = -72, 'AGM' = -71, 'Elephone' = -70, 'MINIX' = -69, 'MI' = -68, 'Doro' = -67, 'OPPO' = -66, 'Windows Vista' = -65, 'Hewlett-Packard' = -64, 'Acer' = -63, 'Cat' = -62, 'MBX' = -61, 'CHG_Tv_Hub' = -60, 'Teclast' = -59, 'UMIDIGI' = -58, 'xiaomi' = -57, 'XGODY' = -56, 'Apple' = -55, 'Harman TV' = -54, 'LeMobile' = -53, 'Coolpad' = -52, 'LeEco' = -51, 'JTY' = -50, 'ZoundIndustriesSmartphonesAB' = -49, 'hena' = -48, 'POCO' = -47, 'Canal' = -46, 'atvXperience' = -45, 'Amlogic' = -44, 'alps' = -43, 'lephone' = -42, 'Fake' = -41, 'EasyPlayTV' = -40, 'My' = -39, 'BlackBerry' = -38, 'Netxeon' = -37, 'Valencia2_Y100pro' = -36, 'Hisilicon' = -35, 'Asus' = -34, 'realme' = -33, 'Blackview' = -32, 'YotaDevicesLimited' = -31, 'QUALCOMMYotaDevices' = -30, 'Razer' = -29, 'AZW' = -28, 'Inter Sales A/S' = -27, 'DNA' = -26, 'Realtek' = -25, 'WIKO' = -24, 'VS' = -23, 'SHARP' = -22, 'unknown' = -21, 'Andersson' = -20, 'zte' = -19, 'vivo' = -18, 'Windows XP' = -17, 'Get' = -16, 'PHONEMAX' = -15, 'Wileyfox' = -14, 'BNO' = -13, 'Proscan' = -12, 'UHANS' = -11, 'amlogic' = -10, 'GAEMODEL' = -9, 'blackberry' = -8, 'A-gold' = -7, 'Allwinner' = -6, 'SDMC' = -5, 'KT107' = -4, 'surface' = -3, 'HOMTOM' = -2, 'acer' = -1, 'FORMULER' = 0, 'wheatek' = 1, 'razer' = 2, 'PROBOX2' = 3, 'blackshark' = 4, 'PIPO' = 5, 'CHUWI' = 6, 'softwinner' = 7, 'VIM' = 8, 'ulefone' = 9, 'vernee' = 10, 'Gigaset' = 11, 'Unblocktech' = 12, 'Amazon' = 13, 'IBDL' = 14, 'MD501' = 15, 'coosea' = 16, 'BLU' = 17, 'honor' = 18, 'Unihertz' = 19, 'Chengdu XGimi Technology Co.,Ltd' = 20, 'CUBOT' = 21, 'InfomirSA' = 22, 'Htc' = 23, 'Vernee' = 24, 'ALLDOCUBE' = 25, 'RED' = 26, 'BMXC' = 27, 'Trekstor' = 28, 'techain' = 29, 'LEAGOO' = 30, 'lenovo' = 31, 'Kurio' = 32, 'CROSSCALL' = 33, 'Ematic' = 34, 'CARBAYTA' = 35, 'nvidia' = 36, 'Amino' = 37, 'TrekStor' = 38, 'MStarSemiconductor,Inc.' = 39, 'None' = 40, 'ChengduXGimiTechnologyCo.,Ltd' = 41, 'MStar Semiconductor, Inc.' = 42, 'DENVER' = 43, 'Infomir SA' = 44, 'Droi' = 45, 'dev' = 46, 'Mara' = 47, 'InterSalesA/S' = 48, 'nubia' = 49, 'MTK' = 50, 'KaonMedia' = 51, 'Fengmi' = 52, 'Open BSD' = 53, 'Planet' = 54, 'HarmanTV' = 55, 'MAG' = 56, 'Essential Products' = 57, 'Stofa' = 58, 'Alcatel' = 59, 'UnknownManufacturer' = 60, 'bq' = 61, 'Simbans' = 62, 'essential' = 63, 'RiksTV' = 64, 'Element' = 65, 'Jlinksz' = 66, 'NOMU_S30' = 67, 'VKworld' = 68, 'Windows Mobile' = 69, 'Quidbox' = 70, 'UGOOS' = 71, 'EssentialProducts' = 72, 'GlobalDistributionFZE' = 73, 'WeTek' = 74, 'F3_Pro' = 75, 'iBallSlide' = 76, 'Atari' = 77, 'Marshall' = 78, 'Dr.Ing.h.c.F.PorscheAG' = 79, 'AllCall' = 80, 'GBR' = 81, 'X18' = 82, 'ID2ME' = 83, 'Meizu' = 84, 'NewBund' = 85, 'X23' = 86, 'Alldocube' = 87, 'HX' = 88, 'RKM' = 89, 'SHIFT' = 90, 'Arcadyan' = 91`
	var res = enumRegExp.FindAllStringSubmatch(s, -1)

	var refNumbers = map[string]struct{}{}
	for i := -128; i <= 91; i++ {
		refNumbers[fmt.Sprintf("%d", i)] = struct{}{}
	}

	var resNumbers = map[string]struct{}{}

	for _, v := range res {
		resNumbers[v[2]] = struct{}{}
	}
	if len(resNumbers) != len(refNumbers) {
		t.Fatal("length mismatch")
	}
	for k := range refNumbers {
		if _, exist := resNumbers[k]; !exist {
			t.Fatalf(`%s not exist in resNumbers`, k)
		}
	}
}
