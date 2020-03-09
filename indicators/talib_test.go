/*
Copyright 2016 Mark Chenoweth
Copyright 2018 Alessandro Sanino
Licensed under terms of MIT license (see LICENSE)
*/

package indicators

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ok(t *testing.T, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d: unexpected error: %s\n", filepath.Base(file), line, err.Error())
		t.FailNow()
	}
}

func equals(t *testing.T, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("%s:%d:\n\tgo: %#v\n\tpy: %#v\n", filepath.Base(file), line, exp, act)
		t.FailNow()
	}
}

var (
	testOpen   = []float64{7560.2, 7411.75, 7317.94, 7299.2, 7199.58, 7397.35, 7535.46, 7505.86, 7516.8, 7332.05, 7224.5, 7208.95, 7185.15, 7255.94, 7066.35, 7115.08, 6882.08, 6619.53, 7286.91, 7157.53, 7194.13, 7147.4, 7521.88, 7317.47, 7247.45, 7195.8, 7195.17, 7247.7, 7297.43, 7372.79, 7222.74, 7160.69, 7174.7, 6945.7, 7332.58, 7356.05, 7352.12, 7768.17, 8151.79, 8048.94, 7812.7, 8200, 8008.1, 8180.75, 8105.01, 8813.2, 8810.62, 8715.39, 8898.03, 8901.02, 8698.97, 8629.66, 8728.47, 8661.63, 8382.08, 8435.83, 8320, 8600, 8894.57, 9386.57, 9280.83, 9510.84, 9339.26, 9379.6, 9309.69, 9281, 9168.49, 9617.84, 9754.63, 9803.71, 9907.7, 10173.51, 9851.99, 10270.98, 10348.78, 10229.19, 10364.04, 9909.71, 9925.96, 9725.57, 10190.3, 9589.1, 9610.4, 9696.13, 9664.08, 9969.79, 9664.87, 9305.4, 8778.37, 8816.5, 8710, 8525, 8516, 8917.34, 8755.45, 8757.79, 9059, 9154.1, 8901.12, 8035.79}
	testHigh   = []float64{7560.2, 7431.07, 7415, 7772.71, 7500, 7618.99, 7638.88, 7580, 7666, 7400, 7271, 7296.19, 7302.35, 7269, 7225.23, 7147.91, 6938.77, 7449.68, 7372.12, 7218.11, 7194.13, 7530, 7692.98, 7431.1, 7266.81, 7432, 7255.37, 7349.65, 7524.46, 7384.9, 7302.35, 7237.35, 7184.94, 7402.31, 7396.1, 7495, 7817, 8220, 8463.57, 8048.94, 8200, 8286, 8190, 8196.81, 8895, 8903.2, 8852.35, 9015.22, 9000.1, 9188.1, 8740.54, 8778.66, 8792.98, 8665.95, 8530.7, 8437.47, 8600, 9004.35, 9413.24, 9443.96, 9570, 9529.35, 9464.16, 9474.28, 9615, 9348.6, 9775, 9859.57, 9878, 9948.97, 10178.54, 10199.85, 10383.9, 10495, 10500.5, 10398, 10408.04, 10051.24, 9973.45, 10288, 10300, 9706.27, 9773.2, 9722.39, 10024.08, 10027.66, 9682.73, 9369.99, 8974.75, 8900.94, 8805.06, 8756.11, 8980.34, 8921.8, 8850, 9169, 9187.85, 9219.13, 8901.12, 8191.13}
	testLow    = []float64{7233.87, 7140.08, 7238.04, 7087.09, 7150.06, 7305.56, 7486.7, 7382.69, 7268.03, 7150.23, 7122.28, 7072.2, 7179.65, 7007.48, 7007, 6820, 6550, 6425, 7000, 7072.61, 7109.73, 7124.52, 7247.86, 7156, 7110.73, 7150, 7052, 7231, 7274.43, 7199, 7112.55, 7150, 6900, 6853.53, 7256.03, 7310, 7342.46, 7697.03, 7872.09, 7737.97, 7667, 8000, 7960, 8039, 8105.01, 8555, 8573.91, 8661.52, 8798.9, 8461.38, 8507.93, 8480, 8567.68, 8280, 8212.9, 8252.72, 8276.22, 8546.55, 8876, 9215.5, 9166.07, 9195.93, 9280.33, 9135, 9211.07, 9075, 9142.52, 9521, 9706.94, 9658.58, 9885.04, 9731.2, 9706.94, 10247.35, 10068, 10093.26, 9739, 9598.49, 9467.57, 9602.6, 9312, 9396.91, 9562.6, 9568.51, 9664.08, 9480, 9234.21, 8627.78, 8520, 8421.49, 8525, 8410, 8486.72, 8660.04, 8663.76, 8757.79, 8990.01, 8859.49, 7995.48, 8000}
	testClose  = []float64{7412.66, 7309.64, 7300.43, 7197.78, 7390.42, 7541.79, 7502.65, 7524.26, 7340.52, 7216.07, 7207.42, 7188.42, 7257.47, 7059.03, 7115.08, 6882.19, 6612.3, 7286.91, 7149.57, 7194.11, 7147.4, 7509.7, 7316.17, 7251.52, 7195.79, 7188.3, 7246, 7296.24, 7385.54, 7220.24, 7168.36, 7178.68, 6950.56, 7338.91, 7344.48, 7356.7, 7762.74, 8159.01, 8044.44, 7806.78, 8200, 8016.22, 8180.76, 8105.01, 8813.04, 8809.17, 8710.15, 8892.63, 8908.53, 8696.6, 8625.17, 8717.89, 8655.93, 8378.44, 8422.13, 8329.5, 8590.48, 8894.54, 9400, 9289.18, 9500, 9327.85, 9377.17, 9329.39, 9288.09, 9159.37, 9618.42, 9754.63, 9803.42, 9902, 10173.97, 9850.01, 10268.98, 10348.78, 10228.67, 10364.04, 9899.78, 9912.89, 9697.15, 10185.17, 9595.72, 9612.76, 9696.13, 9668.13, 9965.21, 9652.58, 9305.4, 8779.36, 8816.5, 8703.84, 8527.74, 8528.95, 8917.34, 8755.45, 8753.28, 9066.65, 9153.79, 8893.93, 8033.7, 8069.77}
	testVolume = []float64{5477.97290774, 4324.30016255, 2552.19525842, 7737.39153487, 4641.75678879, 4167.2401632, 2034.573719, 1798.77041268, 6582.47031165, 4090.62092926, 3419.93896517, 4376.88899274, 3178.07849595, 2012.34890643, 1848.34572822, 5923.05547919, 7107.25638535, 12873.63439214, 5936.55634536, 3201.66392251, 1689.0365357, 3796.22608082, 6710.7673604, 4194.58980335, 1504.52751296, 3116.74715162, 4024.57086599, 1579.70272647, 2583.85134193, 3722.91334933, 2638.69130899, 1119.10969318, 3972.70795427, 8072.72908749, 3256.73735349, 2707.27385549, 6728.28277831, 12158.32281238, 11913.23248192, 5957.24423978, 8999.98825687, 3788.3273136, 2056.08675369, 4011.43909644, 17009.43193382, 8606.34943936, 6713.2584764, 9363.15892644, 3678.00926615, 8785.449927, 4420.19825219, 3394.68916005, 3143.81212331, 8611.0107876, 6430.43023326, 2812.78995216, 6340.58577579, 9619.7747426, 9565.55935515, 6481.91594603, 7945.15877894, 5009.62230695, 1669.00246242, 3015.79158859, 4758.13350441, 4867.22545881, 8063.96793828, 7554.48867469, 4013.49096731, 3268.73554695, 4830.33787739, 7560.66789084, 8127.59636283, 6553.40942709, 9347.45128374, 6190.37749036, 5762.30634481, 4133.47675014, 5913.20492551, 6754.35394103, 9936.8206814, 6608.4910792, 4271.47074293, 2096.56353593, 3329.22433648, 5877.00770771, 6280.00295543, 11564.19794381, 9273.9036636, 6546.98122765, 2338.06785643, 4156.57173933, 6186.86088929, 4629.97228611, 3353.2325223, 6448.68567336, 4136.42339429, 2805.5915714, 9068.89040196, 1975.36714117}

	testRand = []float64{0.42422904963267427, 0.16755615298728432, 0.5946077386900349, 0.17611040890583352, 0.29152918200482136, 0.27807733751955355, 0.7177400699036796, 0.5036012923358724, 0.1629504791237938, 0.6483065114032258, 0.5703588423748475, 0.7161845737507714, 0.6942714038794598, 0.42176699339445745, 0.7884431075157385, 0.24584359985404292, 0.7480158197252457, 0.2651217282085182, 0.4437589032368914, 0.9845738324910773, 0.5590040804528499, 0.25521017265864154, 0.1372114571360159, 0.1218701299153161, 0.25511876291008395, 0.7483943425884052, 0.076845841747889, 0.5389677976892574, 0.9015900382854415, 0.13503746751073498, 0.17237105554803778, 0.022111455150970016, 0.4735780024560894, 0.694458845807901, 0.5530772348613145, 0.3444350790493579, 0.6468662907768967, 0.6359557337589957, 0.5650572127602662, 0.621587087190788, 0.5634446451263618, 0.6967583014608363, 0.3366771423506647, 0.8920892600559512, 0.00029418556385873984, 0.1664001753124047, 0.2032534540019577, 0.30597531513267284, 0.4581883332445693, 0.4877258346021447}

	// testCrossunder1 = []float64{1, 2, 3, 4, 8, 6, 7}
	// testCrossunder2 = []float64{1, 1, 10, 9, 5, 3, 7}

	// testNothingCrossed1 = []float64{1, 2, 3, 4, 8, 6, 7}
	// testNothingCrossed2 = []float64{1, 4, 5, 9, 5, 3, 7}
	// testCrossover1 = []float64{1, 3, 2, 4, 8, 6, 7}
	// testCrossover2 = []float64{1, 5, 1, 4, 5, 6, 7}
)

func a2s(a []float64) string { // go float64 array to python list initializer string
	return strings.ReplaceAll(fmt.Sprintf("%f", a), " ", ",")
}

func round(input float64) float64 {
	if input < 0 {
		return math.Ceil(input - 0.5)
	}
	return math.Floor(input + 0.5)
}

func compare(t *testing.T, goResult []float64, taCall string) {
	pyprog := fmt.Sprintf(`import talib,numpy
testOpen = numpy.array(%s)
testHigh = numpy.array(%s)
testLow = numpy.array(%s)
testClose = numpy.array(%s)
testVolume = numpy.array(%s)
testRand = numpy.array(%s)
%s
print(' '.join([str(p) for p in result]).replace('nan','0.0'))`,
		a2s(testOpen), a2s(testHigh), a2s(testLow), a2s(testClose), a2s(testVolume), a2s(testRand), taCall)

	pyOut, err := exec.Command("python", "-c", pyprog).Output()
	ok(t, err)

	var pyResult []float64
	strResult := strings.Fields(string(pyOut))
	for _, arg := range strResult {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			pyResult = append(pyResult, n)
		}
	}

	equals(t, len(goResult), len(pyResult))

	for i := 0; i < len(goResult); i++ {
		if (goResult[i] < -0.00000000000001) || (goResult[i] < 0.00000000000001) {
			goResult[i] = 0.0
		}
		if (pyResult[i] < -0.00000000000001) || (pyResult[i] < 0.00000000000001) {
			pyResult[i] = 0.0
		}

		var s1, s2 string
		if (goResult[i] > -1000000) && (goResult[i] < 1000000) {
			s1 = fmt.Sprintf("%.6f", goResult[i])
		} else {
			s1 = fmt.Sprintf("%.1f", round(goResult[i])) // reduce precision for very large numbers
		}

		if (pyResult[i] > -1000000) && (pyResult[i] < 1000000) {
			s2 = fmt.Sprintf("%.6f", pyResult[i])
		} else {
			s2 = fmt.Sprintf("%.1f", round(pyResult[i])) // reduce precision for very large numbers
		}
		if s1[:len(s1)-2] != s2[:len(s2)-2] {
			_, file, line, _ := runtime.Caller(1)
			fmt.Printf("%s:%d:\n\tgo!: %#v\n\tpy!: %#v\n", filepath.Base(file), line, s1, s2)
			t.FailNow()
		}
	}
}

// Ensure that python and talib are installed and in the PATH
func TestMain(m *testing.M) {
	pyout, err := exec.Command("python", "-c", "import talib; print('success')").Output()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if string(pyout[0:7]) != "success" {
		fmt.Println("python and talib must be installed to run tests")
		os.Exit(-1)
	}
	os.Exit(m.Run())
}

// Test all the functions

func TestSma(t *testing.T) {
	result := Sma(testClose, 20)
	compare(t, result, "result = talib.SMA(testClose,20)")
}

func TestEma(t *testing.T) {
	result := Ema(testClose, 5)
	compare(t, result, "result = talib.EMA(testClose,5)")
	result = Ema(testClose, 20)
	compare(t, result, "result = talib.EMA(testClose,20)")
	result = Ema(testClose, 50)
	compare(t, result, "result = talib.EMA(testClose,50)")
	result = Ema(testClose, 100)
	compare(t, result, "result = talib.EMA(testClose,100)")
}

func TestRsi(t *testing.T) {
	result := Rsi(testClose, 10)
	compare(t, result, "result = talib.RSI(testClose,10)")
}

func TestAdd(t *testing.T) {
	result := Add(testHigh, testLow)
	compare(t, result, "result = talib.ADD(testHigh,testLow)")
}

func TestDiv(t *testing.T) {
	result := Div(testHigh, testLow)
	compare(t, result, "result = talib.DIV(testHigh,testLow)")
}

func TestMax(t *testing.T) {
	result := Max(testClose, 10)
	compare(t, result, "result = talib.MAX(testClose,10)")
}

func TestMaxIndex(t *testing.T) {
	result := MaxIndex(testClose, 10)
	compare(t, result, "result = talib.MAXINDEX(testClose,10)")
}

func TestMin(t *testing.T) {
	result := Min(testClose, 10)
	compare(t, result, "result = talib.MIN(testClose,10)")
}

func TestMinIndex(t *testing.T) {
	result := MinIndex(testClose, 10)
	compare(t, result, "result = talib.MININDEX(testClose,10)")
}

func TestMult(t *testing.T) {
	result := Mult(testHigh, testLow)
	compare(t, result, "result = talib.MULT(testHigh,testLow)")
}

func TestSub(t *testing.T) {
	result := Sub(testHigh, testLow)
	compare(t, result, "result = talib.SUB(testHigh,testLow)")
}

func TestRocp(t *testing.T) {
	result := Rocp(testClose, 10)
	compare(t, result, "result = talib.ROCP(testClose,10)")
}

func TestObv(t *testing.T) {
	result := Obv(testClose, testVolume)
	compare(t, result, "result = talib.OBV(testClose,testVolume)")
}

func TestAtr(t *testing.T) {
	result := Atr(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.ATR(testHigh,testLow,testClose,14)")
}

func TestNatr(t *testing.T) {
	result := Natr(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.NATR(testHigh,testLow,testClose,14)")
}

func TestTRange(t *testing.T) {
	result := TRange(testHigh, testLow, testClose)
	compare(t, result, "result = talib.TRANGE(testHigh,testLow,testClose)")
}

func TestAvgPrice(t *testing.T) {
	result := AvgPrice(testOpen, testHigh, testLow, testClose)
	compare(t, result, "result = talib.AVGPRICE(testOpen,testHigh,testLow,testClose)")
}

func TestMedPrice(t *testing.T) {
	result := MedPrice(testHigh, testLow)
	compare(t, result, "result = talib.MEDPRICE(testHigh,testLow)")
}

func TestTypPrice(t *testing.T) {
	result := TypPrice(testHigh, testLow, testClose)
	compare(t, result, "result = talib.TYPPRICE(testHigh,testLow,testClose)")
}

func TestWclPrice(t *testing.T) {
	result := WclPrice(testHigh, testLow, testClose)
	compare(t, result, "result = talib.WCLPRICE(testHigh,testLow,testClose)")
}

func TestAcos(t *testing.T) {
	result := Acos(testRand)
	compare(t, result, "result = talib.ACOS(testRand)")
}

func TestAsin(t *testing.T) {
	result := Asin(testRand)
	compare(t, result, "result = talib.ASIN(testRand)")
}

func TestAtan(t *testing.T) {
	result := Atan(testRand)
	compare(t, result, "result = talib.ATAN(testRand)")
}

func TestCeil(t *testing.T) {
	result := Ceil(testClose)
	compare(t, result, "result = talib.CEIL(testClose)")
}

func TestCos(t *testing.T) {
	result := Cos(testRand)
	compare(t, result, "result = talib.COS(testRand)")
}

func TestCosh(t *testing.T) {
	result := Cosh(testRand)
	compare(t, result, "result = talib.COSH(testRand)")
}

func TestExp(t *testing.T) {
	result := Exp(testRand)
	compare(t, result, "result = talib.EXP(testRand)")
}

func TestFloor(t *testing.T) {
	result := Floor(testClose)
	compare(t, result, "result = talib.FLOOR(testClose)")
}

func TestLn(t *testing.T) {
	result := Ln(testClose)
	compare(t, result, "result = talib.LN(testClose)")
}

func TestLog10(t *testing.T) {
	result := Log10(testClose)
	compare(t, result, "result = talib.LOG10(testClose)")
}

func TestSin(t *testing.T) {
	result := Sin(testRand)
	compare(t, result, "result = talib.SIN(testRand)")
}

func TestSinh(t *testing.T) {
	result := Sinh(testRand)
	compare(t, result, "result = talib.SINH(testRand)")
}

func TestSqrt(t *testing.T) {
	result := Sqrt(testClose)
	compare(t, result, "result = talib.SQRT(testClose)")
}

func TestTan(t *testing.T) {
	result := Tan(testRand)
	compare(t, result, "result = talib.TAN(testRand)")
}

func TestTanh(t *testing.T) {
	result := Tanh(testRand)
	compare(t, result, "result = talib.TANH(testRand)")
}

func TestSum(t *testing.T) {
	result := Sum(testClose, 10)
	compare(t, result, "result = talib.SUM(testClose,10)")
}

func TestVar(t *testing.T) {
	result := Var(testClose, 10)
	compare(t, result, "result = talib.VAR(testClose,10)")
}

func TestTsf(t *testing.T) {
	result := Tsf(testClose, 10)
	compare(t, result, "result = talib.TSF(testClose,10)")
}

func TestStdDev(t *testing.T) {
	result := StdDev(testRand, 10, 1.0)
	compare(t, result, "result = talib.STDDEV(testRand,10,1.0)")
}

func TestLinearRegSlope(t *testing.T) {
	result := LinearRegSlope(testClose, 10)
	compare(t, result, "result = talib.LINEARREG_SLOPE(testClose,10)")
}

func TestLinearRegIntercept(t *testing.T) {
	result := LinearRegIntercept(testClose, 10)
	compare(t, result, "result = talib.LINEARREG_INTERCEPT(testClose,10)")
}

func TestLinearRegAngle(t *testing.T) {
	result := LinearRegAngle(testClose, 10)
	compare(t, result, "result = talib.LINEARREG_ANGLE(testClose,10)")
}

func TestLinearReg(t *testing.T) {
	result := LinearReg(testClose, 10)
	compare(t, result, "result = talib.LINEARREG(testClose,10)")
}

func TestCorrel(t *testing.T) {
	result := Correl(testHigh, testLow, 10)
	compare(t, result, "result = talib.CORREL(testHigh,testLow,10)")
}

func TestBeta(t *testing.T) {
	result := Beta(testHigh, testLow, 5)
	compare(t, result, "result = talib.BETA(testHigh,testLow,5)")
}

func TestHtDcPeriod(t *testing.T) {
	result := HtDcPeriod(testClose)
	compare(t, result, "result = talib.HT_DCPERIOD(testClose)")
}

func TestHtPhasor(t *testing.T) {
	result1, result2 := HtPhasor(testClose)
	compare(t, result1, "result,_ = talib.HT_PHASOR(testClose)")
	compare(t, result2, "_,result = talib.HT_PHASOR(testClose)")
}

func TestHtSine(t *testing.T) {
	result1, result2 := HtSine(testClose)
	compare(t, result1, "result,_ = talib.HT_SINE(testClose)")
	compare(t, result2, "_,result = talib.HT_SINE(testClose)")
}

func TestHtTrendline(t *testing.T) {
	result := HtTrendline(testClose)
	compare(t, result, "result = talib.HT_TRENDLINE(testClose)")
}

func TestHtTrendMode(t *testing.T) {
	result := HtTrendMode(testClose)
	compare(t, result, "result = talib.HT_TRENDMODE(testClose)")
}

func TestWillR(t *testing.T) {
	result := WillR(testHigh, testLow, testClose, 9)
	compare(t, result, "result = talib.WILLR(testHigh,testLow,testClose,9)")
}

func TestAdx(t *testing.T) {
	result := Adx(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.ADX(testHigh,testLow,testClose,14)")
}

func TestAdxR(t *testing.T) {
	result := AdxR(testHigh, testLow, testClose, 5)
	compare(t, result, "result = talib.ADXR(testHigh,testLow,testClose,5)")
}

func TestCci(t *testing.T) {
	result := Cci(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.CCI(testHigh,testLow,testClose,14)")
}

func TestRoc(t *testing.T) {
	result := Roc(testClose, 10)
	compare(t, result, "result = talib.ROC(testClose,10)")
}

func TestRocr(t *testing.T) {
	result := Rocr(testClose, 10)
	compare(t, result, "result = talib.ROCR(testClose,10)")
}

func TestRocr100(t *testing.T) {
	result := Rocr100(testClose, 10)
	compare(t, result, "result = talib.ROCR100(testClose,10)")
}

func TestMom(t *testing.T) {
	result := Mom(testClose, 10)
	compare(t, result, "result = talib.MOM(testClose,10)")
}

func TestBBands(t *testing.T) {
	upper, middle, lower := BBands(testClose, 5, 2.0, 2.0, SMA)
	compare(t, upper, "result,upper,lower = talib.BBANDS(testClose,5,2.0,2.0)")
	compare(t, middle, "upper,result,lower = talib.BBANDS(testClose,5,2.0,2.0)")
	compare(t, lower, "upper,middle,result = talib.BBANDS(testClose,5,2.0,2.0)")
}

func TestDema(t *testing.T) {
	result := Dema(testClose, 10)
	compare(t, result, "result = talib.DEMA(testClose,10)")
}

func TestTema(t *testing.T) {
	result := Tema(testClose, 10)
	compare(t, result, "result = talib.TEMA(testClose,10)")
}

func TestWma(t *testing.T) {
	result := Wma(testClose, 10)
	compare(t, result, "result = talib.WMA(testClose,10)")
}

func TestMa(t *testing.T) {
	result := Ma(testClose, 10, DEMA)
	compare(t, result, "result = talib.MA(testClose,10,talib.MA_Type.DEMA)")
}

func TestTrima(t *testing.T) {
	result := Trima(testClose, 10)
	compare(t, result, "result = talib.TRIMA(testClose,10)")
	result = Trima(testClose, 11)
	compare(t, result, "result = talib.TRIMA(testClose,11)")
}

func TestMidPoint(t *testing.T) {
	result := MidPoint(testClose, 10)
	compare(t, result, "result = talib.MIDPOINT(testClose,10)")
}

func TestMidPrice(t *testing.T) {
	result := MidPrice(testHigh, testLow, 10)
	compare(t, result, "result = talib.MIDPRICE(testHigh,testLow,10)")
}

func TestT3(t *testing.T) {
	result := T3(testClose, 5, 0.7)
	compare(t, result, "result = talib.T3(testClose,5,0.7)")
}

func TestKama(t *testing.T) {
	result := Kama(testClose, 10)
	compare(t, result, "result = talib.KAMA(testClose,10)")
}

func TestMaVp(t *testing.T) {
	periods := make([]float64, len(testClose))
	for i := range testClose {
		periods[i] = 5.0
	}
	result := MaVp(testClose, periods, 2, 10, SMA)
	compare(t, result, "result = talib.MAVP(testClose,numpy.full(len(testClose),5.0),2,10,talib.MA_Type.SMA)")
}

func TestMinusDM(t *testing.T) {
	result := MinusDM(testHigh, testLow, 14)
	compare(t, result, "result = talib.MINUS_DM(testHigh,testLow,14)")
}

func TestPlusDM(t *testing.T) {
	result := PlusDM(testHigh, testLow, 14)
	compare(t, result, "result = talib.PLUS_DM(testHigh,testLow,14)")
}

func TestSar(t *testing.T) {
	result := Sar(testHigh, testLow, 0.0, 0.0)
	compare(t, result, "result = talib.SAR(testHigh,testLow,0.0,0.0)")
}

func TestSarExt(t *testing.T) {
	result := SarExt(testHigh, testLow, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	compare(t, result, "result = talib.SAREXT(testHigh,testLow,0.0,0.0,0.0,0.0,0.0,0.0,0.0,0.0)")
}

func TestMama(t *testing.T) {
	mama, fama := Mama(testClose, 0.5, 0.05)
	compare(t, mama, "result,fama = talib.MAMA(testClose,0.5,0.05)")
	compare(t, fama, "mama,result = talib.MAMA(testClose,0.5,0.05)")
}

func TestMinMax(t *testing.T) {
	min, max := MinMax(testClose, 10)
	compare(t, min, "result,max = talib.MINMAX(testClose,10)")
	compare(t, max, "min,result = talib.MINMAX(testClose,10)")
}

func TestMinMaxIndex(t *testing.T) {
	minidx, maxidx := MinMaxIndex(testClose, 10)
	compare(t, minidx, "result,maxidx = talib.MINMAXINDEX(testClose,10)")
	compare(t, maxidx, "minidx,result = talib.MINMAXINDEX(testClose,10)")
}

func TestApo(t *testing.T) {
	result := Apo(testClose, 12, 26, SMA)
	compare(t, result, "result = talib.APO(testClose,12,26,talib.MA_Type.SMA)")
	result = Apo(testClose, 26, 12, SMA)
	compare(t, result, "result = talib.APO(testClose,26,12,talib.MA_Type.SMA)")
}

func TestPpo(t *testing.T) {
	result := Ppo(testClose, 12, 26, SMA)
	compare(t, result, "result = talib.PPO(testClose,12,26,talib.MA_Type.SMA)")
	result = Ppo(testClose, 26, 12, SMA)
	compare(t, result, "result = talib.PPO(testClose,26,12,talib.MA_Type.SMA)")
}

func TestAroon(t *testing.T) {
	dn, up := Aroon(testHigh, testLow, 14)
	compare(t, dn, "result,up = talib.AROON(testHigh,testLow,14)")
	compare(t, up, "dn,result = talib.AROON(testHigh,testLow,14)")
}

func TestAroonOsc(t *testing.T) {
	result := AroonOsc(testHigh, testLow, 14)
	compare(t, result, "result = talib.AROONOSC(testHigh,testLow,14)")
}

func TestBop(t *testing.T) {
	result := Bop(testOpen, testHigh, testLow, testClose)
	compare(t, result, "result = talib.BOP(testOpen,testHigh,testLow,testClose)")
}

func TestCmo(t *testing.T) {
	result := Cmo(testClose, 14)
	compare(t, result, "result = talib.CMO(testClose,14)")
}

func TestDx(t *testing.T) {
	result := Dx(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.DX(testHigh,testLow,testClose,14)")
}

func TestMinusDI(t *testing.T) {
	result := MinusDI(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.MINUS_DI(testHigh,testLow,testClose,14)")
}

func TestPlusDI(t *testing.T) {
	result := PlusDI(testHigh, testLow, testClose, 14)
	compare(t, result, "result = talib.PLUS_DI(testHigh,testLow,testClose,14)")
}

//
func TestMfi(t *testing.T) {
	result := Mfi(testHigh, testLow, testClose, testVolume, 14)
	compare(t, result, "result = talib.MFI(testHigh,testLow,testClose,testVolume,14)")
}

func TestUltOsc(t *testing.T) {
	result := UltOsc(testHigh, testLow, testClose, 7, 14, 28)
	compare(t, result, "result = talib.ULTOSC(testHigh,testLow,testClose,7,14,28)")
}

func TestStoch(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 5, 3, SMA, 3, SMA)
	compare(t, slowk, "result,slowd = talib.STOCH(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
	compare(t, slowd, "slowk,result = talib.STOCH(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
}

func TestStoch2(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 12, 3, SMA, 3, SMA)
	compare(t, slowk, "result,slowd = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
	compare(t, slowd, "slowk,result = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,3,talib.MA_Type.SMA)")
}

func TestStoch3(t *testing.T) {
	slowk, slowd := Stoch(testHigh, testLow, testClose, 12, 3, SMA, 15, SMA)
	compare(t, slowk, "result,slowd = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,15,talib.MA_Type.SMA)")
	compare(t, slowd, "slowk,result = talib.STOCH(testHigh,testLow,testClose,12,3,talib.MA_Type.SMA,15,talib.MA_Type.SMA)")
}

func TestStochF(t *testing.T) {
	fastk, fastd := StochF(testHigh, testLow, testClose, 5, 3, SMA)
	compare(t, fastk, "result,fastd = talib.STOCHF(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA)")
	compare(t, fastd, "fastk,result = talib.STOCHF(testHigh,testLow,testClose,5,3,talib.MA_Type.SMA)")
}

func TestStochRsi(t *testing.T) {
	fastk, fastd := StochRsi(testClose, 14, 5, 2, SMA)
	compare(t, fastk, "result,fastd = talib.STOCHRSI(testClose,14,5,2,talib.MA_Type.SMA)")
	compare(t, fastd, "fastk,result = talib.STOCHRSI(testClose,14,5,2,talib.MA_Type.SMA)")
}

func TestMacdExt(t *testing.T) {
	macd, macdsignal, macdhist := MacdExt(testClose, 12, SMA, 26, SMA, 9, SMA)
	compare(t, macd, "result, macdsignal, macdhist = talib.MACDEXT(testClose,12,talib.MA_Type.SMA,26,talib.MA_Type.SMA,9,talib.MA_Type.SMA)")
	compare(t, macdsignal, "macd, result, macdhist = talib.MACDEXT(testClose,12,talib.MA_Type.SMA,26,talib.MA_Type.SMA,9,talib.MA_Type.SMA)")
	compare(t, macdhist, "macd, macdsignal, result = talib.MACDEXT(testClose,12,talib.MA_Type.SMA,26,talib.MA_Type.SMA,9,talib.MA_Type.SMA)")
}

func TestTrix(t *testing.T) {
	result := Trix(testClose, 5)
	compare(t, result, "result = talib.TRIX(testClose,5)")
	result = Trix(testClose, 30)
	compare(t, result, "result = talib.TRIX(testClose,30)")
}

func TestMacd(t *testing.T) {
	macd, macdsignal, macdhist := Macd(testClose, 12, 26, 9)
	unstable := 100
	compare(t, macd[unstable:], fmt.Sprintf("result, macdsignal, macdhist = talib.MACD(testClose,12,26,9); result = result[%d:]", unstable))
	compare(t, macdsignal[unstable:], fmt.Sprintf("macd, result, macdhist = talib.MACD(testClose,12,26,9); result = result[%d:]", unstable))
	compare(t, macdhist[unstable:], fmt.Sprintf("macd, macdsignal, result = talib.MACD(testClose,12,26,9); result = result[%d:]", unstable))
}

func TestMacdFix(t *testing.T) {
	macd, macdsignal, macdhist := MacdFix(testClose, 9)
	unstable := 100
	compare(t, macd[unstable:], fmt.Sprintf("result, macdsignal, macdhist = talib.MACDFIX(testClose,9); result = result[%d:]", unstable))
	compare(t, macdsignal[unstable:], fmt.Sprintf("macd, result, macdhist = talib.MACDFIX(testClose,9); result = result[%d:]", unstable))
	compare(t, macdhist[unstable:], fmt.Sprintf("macd, macdsignal, result = talib.MACDFIX(testClose,9); result = result[%d:]", unstable))
}

func TestAd(t *testing.T) {
	result := Ad(testHigh, testLow, testClose, testVolume)
	compare(t, result, "result = talib.AD(testHigh,testLow,testClose,testVolume)")
}

func TestAdOsc(t *testing.T) {
	result := AdOsc(testHigh, testLow, testClose, testVolume, 3, 10)
	compare(t, result, "result = talib.ADOSC(testHigh,testLow,testClose,testVolume,3,10)")
}

// func TestHeikinashiCandles(t *testing.T) {
// 	expectedHighs := []float64{200.24, 198.62, 198.62, 201.99, 202.25, 200.46, 201.33, 197.03, 197.93, 197.74, 198.62, 199.54, 202.09, 201.93, 201.4, 199.99, 200.16, 198.21, 198.08, 197.95, 200.71, 201.23, 202.13, 203.05, 201.48, 202.93, 203.26, 204.76, 205.6, 206.07, 205.97, 206.17, 207.06, 206.94, 207.76, 207.95, 207.43, 207.3, 207.77, 207.76, 206.23, 206.54, 205.7, 204.57, 202.63, 201.35, 202.99, 203.73, 204.47, 204.21, 207, 206.21, 207.68, 207.77, 207.07, 206.03, 203.1, 202.69, 205.3, 204.8, 203.15, 203.7, 205.15, 205.45, 205.21, 205.87, 206.76, 207.29, 206.39, 207.7, 207.64, 205.91, 206.92, 207.52, 207.51, 208.58, 208.61, 209.11, 208.15, 207.94, 207.02, 207.43, 208.66, 208.11, 206.6, 206.06, 208.5, 208.53, 207.29, 207.87, 208.96, 209.24, 210.02, 210.19, 210.39, 210.36, 210.16, 209.54, 209.61, 209.22, 209.06, 208.98, 208.83, 209.3, 208.5, 207.24, 206.5, 205.79, 208.06, 208.73, 208.13, 206.13, 207.02, 207.97, 209.96, 209.21, 210.24, 210.09, 209.82, 208.91, 208.25, 207.51, 205.03, 205.73, 205.97, 205.35, 205.87, 204.47, 205.06, 205.68, 207.58, 208.72, 208.94, 209.95, 210.2, 210.82, 210.39, 209.43, 209.31, 208.04, 205.25, 207.18, 208.71, 208.69, 209.11, 208.2, 207.93, 208.97, 208.09, 206.04, 208.34, 207.15, 206.83, 207.23, 207.19, 208.26, 208.35, 207.69, 205.99, 201.68, 195.3, 193.29, 192.64, 197.21, 197.63, 196.93, 192.62, 193.3, 195.86, 191.72, 195.42, 197.26, 195.04, 194.64, 194.83, 196.79, 198.19, 200.65, 197.5, 196.51, 193.31, 193.52, 192.31, 193.85, 190.77, 188.62, 190.7, 191.35, 193.88, 197.56, 197.8, 198.65, 200.36, 200.71, 200.57, 200.96, 199.68, 201.16, 202.09, 202.17, 202.63, 202.58, 204.29, 206.72, 206.14, 205.78, 207.74, 208.03, 208.2, 209.37, 210.41, 210.25, 209.73, 209.08, 208.25, 207.37, 207.7, 205.83, 203.46, 204.47, 205.82, 207.66, 207.81, 208.88, 208.74, 208.59, 208.5, 208.56, 208.65, 209.57, 209.75, 207.91, 208.73, 208.49, 207.06, 207.45, 206.2, 202.93, 201.85, 204.89, 207.16, 207.25, 202.93, 201.88, 203.85, 206.07, 206.33, 205.26, 207.79, 207.21, 205.89}
// 	expectedOpens := []float64{201.745, 198.83999999999997, 196.89, 197.785, 200.815, 201.175, 199.41500000000002, 198.99, 196.205, 196.275, 196.15, 198.14, 198.2, 200.9, 201.23000000000002, 200.935, 198.755, 198.065, 197.12, 195.96499999999997, 196.925, 199.8, 199.865, 201.33499999999998, 201.885, 200.56, 202.175, 202.59, 204.195, 205.19, 205.515, 205.64999999999998, 205.45499999999998, 206.10500000000002, 206.81, 207.19, 207.365, 207.175, 206.695, 207.11, 207.01999999999998, 206.065, 206.28, 204.245, 203.845, 201.685, 200.755, 202.01, 202.13, 203.445, 203.625, 204.7, 205.485, 206.735, 206.88, 206.015, 204.13, 201.865, 202.18, 204.325, 203.57, 202.78, 202.765, 203.32999999999998, 204.285, 204.47, 205.04000000000002, 206.3, 206.16, 205.855, 206.91, 206.87, 205.14499999999998, 206.135, 206.805, 206.985, 207.315, 208.305, 208.2, 207.745, 207.135, 205.85500000000002, 206.73000000000002, 207.925, 206.64, 205.49, 205.095, 207.905, 207.745, 206.47, 206.915, 208.37, 209.07, 209.3, 209.755, 209.64, 209.73000000000002, 209.64, 208.195, 208.615, 209.03, 208.39999999999998, 208.4, 207.845, 208.6, 207.265, 206.535, 205.75, 205.15, 206.835, 208.20499999999998, 206.99, 205.56, 206.26999999999998, 207.25, 208.685, 208.8, 209.56, 209.75, 208.78, 208.155, 207.73000000000002, 204.45, 204.26999999999998, 205.32, 205.4, 203.96, 205.19, 203.20499999999998, 203.69, 205.095, 207.06, 207.875, 208.34, 209.74, 210.03, 210.32, 209.75, 208.815, 208.525, 206.83499999999998, 204.575, 206.255, 207.8, 208.16500000000002, 208.625, 207.8, 207.22, 207.935, 207.005, 205.755, 207.60500000000002, 206.505, 205.715, 206.385, 206.615, 207.32999999999998, 207.8, 206.4, 202.97, 197.57, 186.41, 189.235, 191.135, 195.95499999999998, 196.695, 195.7, 190.315, 192.86, 193.74, 190.59, 194.51, 194.88, 193.04500000000002, 193.89, 194.305, 195.35, 197.29500000000002, 197.66500000000002, 194.42000000000002, 195.29000000000002, 192.74, 192.70499999999998, 191.385, 192.61, 188.775, 187.08499999999998, 189.87, 190.965, 191.25, 196.3, 196.88, 197.975, 198.895, 200.165, 200.28, 199.265, 198.555, 200.025, 201.85, 201.735, 201.78, 201.535, 202.91500000000002, 206.15, 205.925, 205.18, 206.745, 207.35500000000002, 207.26, 208.12, 209.24, 209.61, 209.05, 208.65, 207.45999999999998, 206.805, 207.075, 204.45499999999998, 202.24, 202.76, 204.51, 206.16, 207.34, 208.14, 207.985, 207.375, 208.17000000000002, 208.255, 207.985, 208.815, 208.335, 205.99, 206.385, 207.555, 205.5, 204.55, 204.425, 201.42000000000002, 201.285, 203.655, 205.97500000000002, 205.41, 201.395, 201.54, 203.11, 205.35500000000002, 205.7, 205.03500000000003, 206.95499999999998, 206.52}
// 	expectedCloses := []float64{198.79999999999998, 196.81, 197.7525, 200.8725, 201, 199.2825, 198.94250000000002, 196, 196.335, 196.14499999999998, 197.755, 198.205, 200.53250000000003, 201.265, 200.75, 198.79, 198.04, 196.7775, 196.2775, 196.415, 199.69, 200.09, 201.3575, 201.89999999999998, 200.6525, 201.95499999999998, 202.5275, 203.98499999999999, 205.13, 205.4925, 205.595, 205.5225, 205.945, 206.695, 207.16, 207.40749999999997, 207.0425, 206.75750000000002, 207.1125, 206.9075, 205.7975, 206.17749999999998, 204.275, 203.9025, 201.6975, 200.7825, 202.015, 202.10750000000002, 203.265, 203.565, 204.70999999999998, 205.495, 206.82999999999998, 207.04999999999998, 206.1325, 204.185, 201.93, 202.17499999999998, 204.40750000000003, 203.7575, 202.495, 202.845, 203.4425, 204.495, 204.4875, 204.965, 206.2525, 206.3325, 205.72499999999997, 207.035, 206.9625, 204.9825, 206.20999999999998, 206.7625, 206.7675, 207.47250000000003, 208.2475, 208.1775, 207.4125, 207.1225, 205.76500000000001, 206.7125, 208.0675, 206.7025, 205.26500000000001, 205.12, 207.9375, 207.8, 206.385, 207.03, 208.3175, 208.97, 209.35500000000002, 209.755, 209.7, 209.74, 209.745, 208.20000000000002, 208.565, 208.89, 208.335, 208.265, 207.865, 208.62, 207.365, 206.49499999999998, 205.7725, 205.1225, 206.9275, 208.2475, 207.1175, 205.4375, 206.24249999999998, 207.1275, 208.655, 208.71, 209.66500000000002, 209.705, 208.88, 208.16750000000002, 207.64000000000001, 204.8675, 204.14499999999998, 205.1625, 205.3225, 204.1325, 204.525, 203.2175, 203.7375, 204.6375, 207.0825, 207.95000000000002, 208.335, 209.66750000000002, 209.93, 210.33, 209.735, 208.90499999999997, 208.4475, 206.7525, 204.595, 206.05, 207.8275, 208.03, 208.55, 207.535, 207.215, 208.0625, 206.8625, 205.5325, 207.63, 206.405, 205.3375, 206.4275, 206.595, 207.195, 207.83249999999998, 206.3875, 203.395, 198.04, 187.125, 189.1525, 190.3, 195.54250000000002, 196.6875, 195.79000000000002, 190.4675, 192.3275, 194.03500000000003, 190.5975, 194.3625, 194.805, 193.3075, 193.70000000000002, 194.1775, 195.32, 197.25, 198.26500000000001, 195.03750000000002, 195.28749999999997, 192.55249999999998, 192.675, 191.1275, 192.4375, 188.7125, 187.15249999999997, 189.69, 190.495, 191.095, 196.3325, 196.8475, 197.72750000000002, 198.89249999999998, 200.1075, 200.2125, 199.59, 198.6375, 199.91750000000002, 201.63, 201.64249999999998, 201.885, 201.5275, 202.695, 206.025, 205.8325, 205.1775, 206.555, 207.43, 207.3075, 208.1375, 209.3375, 209.4875, 208.92, 208.40250000000003, 207.225, 206.735, 207.07, 204.5875, 202.29500000000002, 202.7425, 204.6275, 206.1875, 207.365, 208.19500000000002, 208, 207.38, 208.1525, 208.17249999999999, 207.9875, 208.7675, 208.35500000000002, 205.8575, 206.4725, 207.3925, 205.65500000000003, 204.88, 204.745, 201.52249999999998, 200.7975, 203.4675, 205.675, 205.42499999999998, 201.38750000000002, 201.2625, 202.90499999999997, 205.34, 205.7875, 204.8175, 207.0425, 206.5025, 204.69}
// 	expectedLows := []float64{197.64, 195.78, 197.35, 199.89, 200.12, 198.55, 197.99, 195.61, 195, 194.75, 197.97, 197.43, 199.87, 200.83, 200.57, 198.64, 196.09, 196.33, 195.42, 196.01, 198.9, 199.8, 200.72, 201.39, 200.49, 201.72, 202.43, 203.69, 204.84, 205.17, 205.42, 205.18, 205.24, 206.68, 206.85, 207.35, 207.11, 206.4, 206.52, 206.85, 205.98, 206.2, 203.3, 203.54, 200.84, 200.37, 201.11, 201.67, 202.53, 203.49, 203.2, 205.26, 206.39, 206.67, 205.51, 202.5, 201.71, 201.88, 203.7, 203.16, 202.44, 202.36, 202.12, 204, 204.26, 204.49, 205.89, 205.78, 205.54, 206.72, 206.7, 204.66, 205.75, 206.28, 206.68, 206.82, 208.3, 207.43, 207.4, 207.04, 205.16, 206.08, 207.88, 205.59, 204.74, 204.63, 207.54, 207.27, 206.29, 206.69, 207.89, 209.07, 208.88, 209.65, 209.51, 209.34, 209.62, 207.36, 207.9, 208.97, 207.79, 208.22, 207.68, 208.56, 206.8, 206.45, 205.18, 205.15, 206.05, 208.13, 206.68, 205.33, 205.62, 207.25, 207.96, 208.48, 209.55, 209.71, 208.18, 207.54, 207.5, 203.15, 203.57, 205.21, 205.03, 203.49, 204.67, 202.27, 202.63, 205, 206.68, 207.4, 208.28, 209.53, 209.94, 210.24, 209.42, 208.6, 207.86, 205.7, 204.5, 205.49, 207.16, 207.84, 208.17, 207.47, 207.06, 207.75, 206.05, 205.65, 206.97, 206.35, 204.82, 206.35, 206.13, 206.4, 207.66, 206.02, 201.71, 195.64, 185.42, 185.2, 189.96, 194.84, 196.31, 195.48, 189.65, 192.47, 193.39, 190.46, 193.77, 192.64, 192.41, 193.22, 193.84, 194.44, 196.62, 197.52, 194.29, 195.28, 192.73, 192.45, 191.01, 191.73, 186.9, 187.01, 189.24, 190.94, 188.65, 195.3, 196.62, 197.72, 197.77, 200.14, 200.23, 199.07, 198.11, 198.9, 201.63, 201.3, 201.65, 200.66, 201.78, 206.02, 205.78, 204.98, 205.78, 207.12, 206.7, 207.09, 208.73, 209.12, 208.91, 208.5, 206.85, 206.28, 206.51, 203.63, 201.34, 201.12, 204.25, 204.82, 207.32, 208.07, 207.83, 206.64, 208.08, 208.19, 207.46, 208.2, 207.3, 204.39, 204.39, 207.12, 205.27, 204.13, 204.2, 200.69, 200.87, 203.49, 205.15, 203.65, 200.02, 201.41, 202.72, 204.69, 205.68, 204.86, 206.51, 205.93, 203.87}
//
// 	resultHigh, resultOpen, resultClose, resultLow := HeikinashiCandles(testHigh, testOpen, testClose, testLow)
// 	for i := range expectedHighs {
// 		if resultHigh[i] != expectedHighs[i] {
// 			t.Errorf("Highs error: Expected %f at cell %d got %f ", expectedHighs[i], i, resultHigh[i])
// 		}
// 	}
//
// 	for i, expected := range expectedOpens {
// 		if resultOpen[i] != expected {
// 			t.Errorf("Opens error: Expected %f at cell %d got %f ", expected, i, resultOpen[i])
// 		}
// 	}
//
// 	for i, expected := range expectedCloses {
// 		if resultClose[i] != expected {
// 			t.Errorf("Closes error: Expected %f at cell %d got %f ", expected, i, resultClose[i])
// 		}
// 	}
//
// 	for i, expected := range expectedLows {
// 		if resultLow[i] != expected {
// 			t.Errorf("Lows error: Expected %f at cell %d got %f ", expected, i, resultLow[i])
// 		}
// 	}
// }

func TestHlc3(t *testing.T) {
	hlc3 := Hlc3(testHigh, testLow, testClose)
	expectedHlc3 := []float64{7402.243333333333, 7293.596666666667, 7317.823333333334, 7352.526666666666, 7346.826666666668, 7488.78, 7542.743333333333, 7495.649999999999, 7424.849999999999, 7255.433333333333, 7200.233333333333, 7185.603333333333, 7246.490000000001, 7111.836666666666, 7115.7699999999995, 6950.033333333333, 6700.356666666667, 7053.863333333334, 7173.8966666666665, 7161.61, 7150.420000000001, 7388.073333333334, 7419.003333333334, 7279.540000000001, 7191.110000000001, 7256.766666666666, 7184.456666666666, 7292.296666666666, 7394.81, 7268.046666666666, 7194.420000000001, 7188.676666666666, 7011.833333333333, 7198.25, 7332.203333333334, 7387.233333333334, 7640.733333333333, 8025.346666666667, 8126.7, 7864.563333333333, 8022.333333333333, 8100.740000000001, 8110.253333333334, 8113.606666666667, 8604.35, 8755.79, 8712.136666666667, 8856.456666666665, 8902.51, 8782.026666666667, 8624.546666666667, 8658.85, 8672.196666666667, 8441.463333333333, 8388.576666666666, 8339.896666666666, 8488.9, 8815.146666666667, 9229.746666666666, 9316.213333333333, 9412.023333333333, 9351.043333333333, 9373.886666666665, 9312.89, 9371.386666666667, 9194.323333333334, 9511.980000000001, 9711.733333333332, 9796.12, 9836.516666666666, 10079.183333333334, 9927.020000000002, 10119.94, 10363.71, 10265.723333333333, 10285.1, 10015.606666666667, 9854.206666666667, 9712.723333333333, 10025.256666666666, 9735.906666666668, 9571.980000000001, 9677.31, 9653.01, 9884.456666666667, 9720.08, 9407.446666666665, 8925.710000000001, 8770.416666666666, 8675.423333333334, 8619.266666666665, 8565.02, 8794.8, 8779.096666666666, 8755.68, 8997.813333333334, 9110.550000000001, 8990.85, 8310.1, 8086.966666666667}
	for i := range expectedHlc3 {
		if hlc3[i] != expectedHlc3[i] {
			t.Fatalf("Assertion error: expected %f, got %f", expectedHlc3[i], hlc3[i])
		}
	}
}

// func TestCrossover(t *testing.T) {
// 	if Crossover(testCrossunder1, testCrossunder2) == true {
// 		t.Error("Crossover: Not expected and found")
// 	}
//
// 	if Crossover(testNothingCrossed1, testNothingCrossed2) == true {
// 		t.Error("Crossover: Not expected and found")
// 	}
//
// 	if Crossover(testCrossover1, testCrossover2) == false {
// 		t.Error("Crossover: Expected and not found")
// 	}
// }

// func TestCrossunder(t *testing.T) {
// 	// Crossunder
// 	series1 := []float64{1, 2, 3, 4, 8, 6, 7}
// 	series2 := []float64{1, 1, 10, 9, 5, 3, 7}

// 	if Crossunder(series1, series2) == false {
// 		t.Error("Crossunder: Expected and not found")
// 	}

// 	// Nothing
// 	series1 = []float64{1, 2, 3, 4, 8, 6, 7}
// 	series2 = []float64{1, 4, 5, 9, 5, 3, 7}

// 	if Crossunder(series1, series2) == true {
// 		t.Error("Crossunder: Not expected and found")
// 	}

// 	// Crossover
// 	series1 = []float64{1, 3, 2, 4, 8, 6, 7}
// 	series2 = []float64{1, 5, 1, 4, 5, 6, 7}

// 	if Crossunder(series1, series2) == true {
// 		t.Error("Crossunder: Not expected and found")
// 	}
// }

func TestGroupCandles(t *testing.T) {
	testHigh := []float64{1, 2, 3, 4, 4.1, 4.5, 4, 2, 5, 3.2}
	testOpens := []float64{0.5, 0.3, 1, 3.1, 3.9, 2.1, 3, 0.9, 1.4, 3.2}
	testCloses := []float64{0.3, 1, 3, 4, 2.1, 3, 1, 1.5, 3.2, 1}
	testLows := []float64{0.1, 0.3, 2, 4, 1, 2, 0.5, 0.9, 1, 1}

	expGroupedHighs := []float64{2, 4, 4.5, 4, 5}
	expGroupedOpens := []float64{0.5, 1, 3.9, 3, 1.4}
	expGroupedCloses := []float64{1, 4, 3, 1.5, 1}
	expGroupedLows := []float64{0.1, 2, 1, 0.5, 1}

	testGroupedHighs, testGroupedOpens, testGroupedCloses, testGroupedLows, err := GroupCandles(testHigh, testOpens, testCloses, testLows, 2)
	require.NoError(t, err, "Unexpected error")

	assert.EqualValues(t, expGroupedHighs, testGroupedHighs, "Highs not expected")
	assert.EqualValues(t, expGroupedOpens, testGroupedOpens, "Opens not expected")
	assert.EqualValues(t, expGroupedCloses, testGroupedCloses, "Closes not expected")
	assert.EqualValues(t, expGroupedLows, testGroupedLows, "Lows not expected")
}
