package characterLevel

var expTable map[int]int = map[int]int{
	1:  0,
	2:  1000,
	3:  2325,
	4:  4025,
	5:  6175,
	6:  8800,
	7:  11950,
	8:  15675,
	9:  20025,
	10: 25025,
	11: 30725,
	12: 37175,
	13: 44400,
	14: 52450,
	15: 61375,
	16: 71200,
	17: 81950,
	18: 93675,
	19: 106400,
	20: 120175,
	21: 135050,
	22: 151850,
	23: 169850,
	24: 189100,
	25: 209650,
	26: 231525,
	27: 254775,
	28: 279425,
	29: 305525,
	30: 333100,
	31: 362200,
	32: 392850,
	33: 425100,
	34: 458975,
	35: 494525,
	36: 531775,
	37: 570750,
	38: 611500,
	39: 654075,
	40: 698500,
	41: 744800,
	42: 795425,
	43: 848125,
	44: 902900,
	45: 959800,
	46: 1018875,
	47: 1080150,
	48: 1143675,
	49: 1209475,
	50: 1277600,
	51: 1348075,
	52: 1424575,
	53: 1503625,
	54: 1585275,
	55: 1669550,
	56: 1756500,
	57: 1846150,
	58: 1938550,
	59: 2033725,
	60: 2131725,
	61: 2232600,
	62: 2341550,
	63: 2453600,
	64: 2568775,
	65: 2687100,
	66: 2808625,
	67: 2933400,
	68: 3061475,
	69: 3192875,
	70: 3327650,
	71: 3465825,
	72: 3614525,
	73: 3766900,
	74: 3922975,
	75: 4082800,
	76: 4246400,
	77: 4413825,
	78: 4585125,
	79: 4760350,
	80: 4939525,
	81: 5122700,
	82: 5338925,
	83: 5581950,
	84: 5855050,
	85: 6161850,
	86: 6506450,
	87: 6893400,
	88: 7327825,
	89: 7815450,
	90: 8362650,
}

func GetRequiredEXP(currentLv, targetLv int) int {
	return expTable[targetLv] - expTable[currentLv]
}
