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

var (
	testOpen   = []float64{7535.46, 7505.86, 7516.8, 7332.05, 7224.5, 7208.95, 7185.15, 7255.94, 7066.35, 7115.08, 6882.08, 6619.53, 7286.91, 7157.53, 7194.13, 7147.4, 7521.88, 7317.47, 7247.45, 7195.8, 7195.17, 7247.7, 7297.43, 7372.79, 7222.74, 7160.69, 7174.7, 6945.7, 7332.58, 7356.05, 7352.12, 7768.17, 8151.79, 8048.94, 7812.7, 8200, 8008.1, 8180.75, 8105.01, 8813.2, 8810.62, 8715.39, 8898.03, 8901.02, 8698.97, 8629.66, 8728.47, 8661.63, 8382.08, 8435.83, 8320, 8600, 8894.57, 9386.57, 9280.83, 9510.84, 9339.26, 9379.6, 9309.69, 9281, 9168.49, 9617.84, 9754.63, 9803.71, 9907.7, 10173.51, 9851.99, 10270.98, 10348.78, 10229.19, 10364.04, 9909.71, 9925.96, 9725.57, 10190.3, 9589.1, 9610.4, 9696.13, 9664.08, 9969.79, 9664.87, 9305.4, 8778.37, 8816.5, 8710, 8525, 8516, 8917.34, 8755.45, 8757.79, 9059, 9154.1, 8901.12, 8035.79, 7937.2, 7893.63, 7956.29, 4841.67, 5623.16, 5167.38}
	testHigh   = []float64{7638.88, 7580, 7666, 7400, 7271, 7296.19, 7302.35, 7269, 7225.23, 7147.91, 6938.77, 7449.68, 7372.12, 7218.11, 7194.13, 7530, 7692.98, 7431.1, 7266.81, 7432, 7255.37, 7349.65, 7524.46, 7384.9, 7302.35, 7237.35, 7184.94, 7402.31, 7396.1, 7495, 7817, 8220, 8463.57, 8048.94, 8200, 8286, 8190, 8196.81, 8895, 8903.2, 8852.35, 9015.22, 9000.1, 9188.1, 8740.54, 8778.66, 8792.98, 8665.95, 8530.7, 8437.47, 8600, 9004.35, 9413.24, 9443.96, 9570, 9529.35, 9464.16, 9474.28, 9615, 9348.6, 9775, 9859.57, 9878, 9948.97, 10178.54, 10199.85, 10383.9, 10495, 10500.5, 10398, 10408.04, 10051.24, 9973.45, 10288, 10300, 9706.27, 9773.2, 9722.39, 10024.08, 10027.66, 9682.73, 9369.99, 8974.75, 8900.94, 8805.06, 8756.11, 8980.34, 8921.8, 8850, 9169, 9187.85, 9219.13, 8901.12, 8191.13, 8158.42, 7988.78, 7969.9, 5990.35, 5663.01, 5965.77}
	testLow    = []float64{7486.7, 7382.69, 7268.03, 7150.23, 7122.28, 7072.2, 7179.65, 7007.48, 7007, 6820, 6550, 6425, 7000, 7072.61, 7109.73, 7124.52, 7247.86, 7156, 7110.73, 7150, 7052, 7231, 7274.43, 7199, 7112.55, 7150, 6900, 6853.53, 7256.03, 7310, 7342.46, 7697.03, 7872.09, 7737.97, 7667, 8000, 7960, 8039, 8105.01, 8555, 8573.91, 8661.52, 8798.9, 8461.38, 8507.93, 8480, 8567.68, 8280, 8212.9, 8252.72, 8276.22, 8546.55, 8876, 9215.5, 9166.07, 9195.93, 9280.33, 9135, 9211.07, 9075, 9142.52, 9521, 9706.94, 9658.58, 9885.04, 9731.2, 9706.94, 10247.35, 10068, 10093.26, 9739, 9598.49, 9467.57, 9602.6, 9312, 9396.91, 9562.6, 9568.51, 9664.08, 9480, 9234.21, 8627.78, 8520, 8421.49, 8525, 8410, 8486.72, 8660.04, 8663.76, 8757.79, 8990.01, 8859.49, 7995.48, 7636, 7733.99, 7591.99, 4776.59, 3850, 5051, 5092.34}
	testClose  = []float64{7502.65, 7524.26, 7340.52, 7216.07, 7207.42, 7188.42, 7257.47, 7059.03, 7115.08, 6882.19, 6612.3, 7286.91, 7149.57, 7194.11, 7147.4, 7509.7, 7316.17, 7251.52, 7195.79, 7188.3, 7246, 7296.24, 7385.54, 7220.24, 7168.36, 7178.68, 6950.56, 7338.91, 7344.48, 7356.7, 7762.74, 8159.01, 8044.44, 7806.78, 8200, 8016.22, 8180.76, 8105.01, 8813.04, 8809.17, 8710.15, 8892.63, 8908.53, 8696.6, 8625.17, 8717.89, 8655.93, 8378.44, 8422.13, 8329.5, 8590.48, 8894.54, 9400, 9289.18, 9500, 9327.85, 9377.17, 9329.39, 9288.09, 9159.37, 9618.42, 9754.63, 9803.42, 9902, 10173.97, 9850.01, 10268.98, 10348.78, 10228.67, 10364.04, 9899.78, 9912.89, 9697.15, 10185.17, 9595.72, 9612.76, 9696.13, 9668.13, 9965.21, 9652.58, 9305.4, 8779.36, 8816.5, 8703.84, 8527.74, 8528.95, 8917.34, 8755.45, 8753.28, 9066.65, 9153.79, 8893.93, 8033.7, 7936.25, 7885.92, 7934.57, 4841.67, 5622.74, 5169.37, 5372.6}
	testVolume = []float64{2034.57, 1798.77, 6582.47, 4090.62, 3419.94, 4376.89, 3178.08, 2012.35, 1848.35, 5923.06, 7107.26, 12873.63, 5936.56, 3201.66, 1689.04, 3796.23, 6710.77, 4194.59, 1504.53, 3116.75, 4024.57, 1579.7, 2583.85, 3722.91, 2638.69, 1119.11, 3972.71, 8072.73, 3256.74, 2707.27, 6728.28, 12158.32, 11913.23, 5957.24, 8999.99, 3788.33, 2056.09, 4011.44, 17009.43, 8606.35, 6713.26, 9363.16, 3678.01, 8785.45, 4420.2, 3394.69, 3143.81, 8611.01, 6430.43, 2812.79, 6340.59, 9619.77, 9565.56, 6481.92, 7945.16, 5009.62, 1669.0, 3015.79, 4758.13, 4867.23, 8063.97, 7554.49, 4013.49, 3268.74, 4830.34, 7560.67, 8127.6, 6553.41, 9347.45, 6190.38, 5762.31, 4133.48, 5913.2, 6754.35, 9936.82, 6608.49, 4271.47, 2096.56, 3329.22, 5877.01, 6280.0, 11564.2, 9273.9, 6546.98, 2338.07, 4156.57, 6186.86, 4629.97, 3353.23, 6448.69, 4136.42, 2805.59, 9068.89, 12177.39, 8274.95, 7289.61, 58513.39, 54419.64, 15370.87, 15675.36}

	testRand = []float64{0.42422904963267427, 0.16755615298728432, 0.5946077386900349, 0.17611040890583352, 0.29152918200482136, 0.27807733751955355, 0.7177400699036796, 0.5036012923358724, 0.1629504791237938, 0.6483065114032258, 0.5703588423748475, 0.7161845737507714, 0.6942714038794598, 0.42176699339445745, 0.7884431075157385, 0.24584359985404292, 0.7480158197252457, 0.2651217282085182, 0.4437589032368914, 0.9845738324910773, 0.5590040804528499, 0.25521017265864154, 0.1372114571360159, 0.1218701299153161, 0.25511876291008395, 0.7483943425884052, 0.076845841747889, 0.5389677976892574, 0.9015900382854415, 0.13503746751073498, 0.17237105554803778, 0.022111455150970016, 0.4735780024560894, 0.694458845807901, 0.5530772348613145, 0.3444350790493579, 0.6468662907768967, 0.6359557337589957, 0.5650572127602662, 0.621587087190788, 0.5634446451263618, 0.6967583014608363, 0.3366771423506647, 0.8920892600559512, 0.00029418556385873984, 0.1664001753124047, 0.2032534540019577, 0.30597531513267284, 0.4581883332445693, 0.4877258346021447}

	testCrossunder1 = []float64{1, 2, 3, 4, 8, 6, 7}
	testCrossunder2 = []float64{1, 1, 10, 9, 5, 3, 7}

	testNothingCrossed1 = []float64{1, 2, 3, 4, 8, 6, 7}
	testNothingCrossed2 = []float64{1, 4, 5, 9, 5, 3, 7}
	testCrossover1      = []float64{1, 3, 2, 4, 8, 3, 8}
	testCrossover2      = []float64{1, 5, 1, 4, 5, 6, 7}
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
	result := Rsi(testClose, 14)
	compare(t, result, "result = talib.RSI(testClose,14)")
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
	t.Skip()
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

func TestGenerateExpectedCandlesOutput(t *testing.T) {
	t.Skip()
	resultHigh, resultOpen, resultClose, resultLow := HeikinashiCandles(testHigh, testOpen, testClose, testLow)

	var outResultHigh string
	for x := range resultHigh {
		outResultHigh += fmt.Sprintf("%v, ", resultHigh[x])
	}

	var outResultOpen string
	for x := range resultOpen {
		outResultOpen += fmt.Sprintf("%v, ", resultOpen[x])
	}
	var outResultClose string
	for x := range resultClose {
		outResultClose += fmt.Sprintf("%v, ", resultClose[x])
	}

	var outResultLow string
	for x := range resultLow {
		outResultLow += fmt.Sprintf("%v, ", resultLow[x])
	}

	fmt.Printf("expectedHighs := []float64{%v}\n", outResultHigh[0:len(outResultHigh)-2])
	fmt.Printf("expectedOpens := []float64{%v}\n", outResultOpen[0:len(outResultOpen)-2])
	fmt.Printf("expectedCloses := []float64{%v}\n", outResultClose[0:len(outResultClose)-2])
	fmt.Printf("expectedLows := []float64{%v}\n", outResultLow[0:len(outResultLow)-2])
}

func TestHeikinashiCandles(t *testing.T) {
	expectedHighs := []float64{0, 7580, 7666, 7400, 7271, 7296.19, 7302.35, 7269, 7225.23, 7147.91, 6938.77, 7449.68, 7372.12, 7218.11, 7194.13, 7530, 7692.98, 7431.1, 7266.81, 7432, 7255.37, 7349.65, 7524.46, 7384.9, 7302.35, 7237.35, 7184.94, 7402.31, 7396.1, 7495, 7817, 8220, 8463.57, 8048.94, 8200, 8286, 8190, 8196.81, 8895, 8903.2, 8852.35, 9015.22, 9000.1, 9188.1, 8740.54, 8778.66, 8792.98, 8665.95, 8530.7, 8437.47, 8600, 9004.35, 9413.24, 9443.96, 9570, 9529.35, 9464.16, 9474.28, 9615, 9348.6, 9775, 9859.57, 9878, 9948.97, 10178.54, 10199.85, 10383.9, 10495, 10500.5, 10398, 10408.04, 10051.24, 9973.45, 10288, 10300, 9706.27, 9773.2, 9722.39, 10024.08, 10027.66, 9682.73, 9369.99, 8974.75, 8900.94, 8805.06, 8756.11, 8980.34, 8921.8, 8850, 9169, 9187.85, 9219.13, 8901.12, 8191.13, 8158.42, 7988.78, 7969.9, 5990.35, 5663.01, 5965.77}
	expectedOpens := []float64{0, 7519.055, 7515.0599999999995, 7428.66, 7274.0599999999995, 7215.96, 7198.6849999999995, 7221.3099999999995, 7157.485, 7090.715, 6998.635, 6747.1900000000005, 6953.219999999999, 7218.24, 7175.82, 7170.764999999999, 7328.549999999999, 7419.025, 7284.495000000001, 7221.62, 7192.05, 7220.585, 7271.969999999999, 7341.485000000001, 7296.514999999999, 7195.549999999999, 7169.6849999999995, 7062.63, 7142.305, 7338.53, 7356.375, 7557.43, 7963.59, 8098.115, 7927.86, 8006.35, 8108.110000000001, 8094.43, 8142.88, 8459.025000000001, 8811.185000000001, 8760.385, 8804.009999999998, 8903.28, 8798.810000000001, 8662.07, 8673.775, 8692.2, 8520.035, 8402.105, 8382.665, 8455.24, 8747.27, 9147.285, 9337.875, 9390.415, 9419.345000000001, 9358.215, 9354.494999999999, 9298.89, 9220.185000000001, 9393.455, 9686.235, 9779.025, 9852.855, 10040.835, 10011.76, 10060.485, 10309.880000000001, 10288.725, 10296.615000000002, 10131.91, 9911.3, 9811.555, 9955.369999999999, 9893.009999999998, 9600.93, 9653.265, 9682.13, 9814.645, 9811.185000000001, 9485.135, 9042.380000000001, 8797.435000000001, 8760.17, 8618.869999999999, 8526.975, 8716.67, 8836.395, 8754.365000000002, 8912.220000000001, 9106.395, 9024.015, 8467.41, 7986.02, 7911.5599999999995, 7914.1, 6398.98, 5232.205, 5396.264999999999}
	expectedCloses := []float64{0, 7498.2025, 7447.8375, 7274.5875, 7206.299999999999, 7191.44, 7231.155000000001, 7147.862499999999, 7103.415, 6991.295, 6745.7875, 6945.28, 7202.15, 7160.59, 7161.3475, 7327.905, 7444.7225, 7289.0225, 7205.195, 7241.525, 7187.135, 7281.147499999999, 7370.465, 7294.2325, 7201.5, 7181.68, 7052.55, 7135.112499999999, 7332.2975, 7379.4375, 7568.58, 7961.0525, 8132.9725, 7910.6575, 7969.925, 8125.555, 8084.715, 8130.3925, 8479.515000000001, 8770.1425, 8736.7575, 8821.189999999999, 8901.390000000001, 8811.775, 8643.1525, 8651.5525, 8686.265, 8496.505000000001, 8386.9525, 8363.88, 8446.675, 8761.36, 9145.9525, 9333.8025, 9379.225, 9390.9925, 9365.23, 9329.567500000001, 9355.962500000001, 9215.9925, 9426.107499999998, 9688.26, 9785.7475, 9828.315, 10036.3125, 9988.642500000002, 10052.9525, 10340.5275, 10286.4875, 10271.122500000001, 10102.715, 9868.082499999999, 9766.0325, 9950.335, 9849.505, 9576.260000000002, 9660.582499999999, 9663.789999999999, 9829.3625, 9782.5075, 9471.8025, 9020.6325, 8772.405, 8710.692500000001, 8641.949999999999, 8555.015, 8725.1, 8813.657500000001, 8755.622500000001, 8937.8075, 9097.6625, 9031.6625, 8457.855, 7949.7925, 7928.8825, 7852.2425, 6386.1125, 5076.1900000000005, 5376.635, 5399.5225}
	expectedLows := []float64{0, 7505.86, 7340.52, 7216.07, 7207.42, 7188.42, 7185.15, 7059.03, 7066.35, 6882.19, 6612.3, 6619.53, 7149.57, 7157.53, 7147.4, 7147.4, 7316.17, 7251.52, 7195.79, 7188.3, 7195.17, 7247.7, 7297.43, 7220.24, 7168.36, 7160.69, 6950.56, 6945.7, 7332.58, 7356.05, 7352.12, 7768.17, 8044.44, 7806.78, 7812.7, 8016.22, 8008.1, 8105.01, 8105.01, 8809.17, 8710.15, 8715.39, 8898.03, 8696.6, 8625.17, 8629.66, 8655.93, 8378.44, 8382.08, 8329.5, 8320, 8600, 8894.57, 9289.18, 9280.83, 9327.85, 9339.26, 9329.39, 9288.09, 9159.37, 9168.49, 9617.84, 9754.63, 9803.71, 9907.7, 9850.01, 9851.99, 10270.98, 10228.67, 10229.19, 9899.78, 9909.71, 9697.15, 9725.57, 9595.72, 9589.1, 9610.4, 9668.13, 9664.08, 9652.58, 9305.4, 8779.36, 8778.37, 8703.84, 8527.74, 8525, 8516, 8755.45, 8753.28, 8757.79, 9059, 8893.93, 8033.7, 7936.25, 7885.92, 7893.63, 4841.67, 4841.67, 5169.37, 5167.38}

	resultHigh, resultOpen, resultClose, resultLow := HeikinashiCandles(testHigh, testOpen, testClose, testLow)
	for i := range expectedHighs {
		if resultHigh[i] != expectedHighs[i] {
			t.Errorf("Highs error: Expected %f at cell %d got %f ", expectedHighs[i], i, resultHigh[i])
		}
	}

	for i, expected := range expectedOpens {
		if resultOpen[i] != expected {
			t.Errorf("Opens error: Expected %f at cell %d got %f ", expected, i, resultOpen[i])
		}
	}

	for i, expected := range expectedCloses {
		if resultClose[i] != expected {
			t.Errorf("Closes error: Expected %f at cell %d got %f ", expected, i, resultClose[i])
		}
	}

	for i, expected := range expectedLows {
		if resultLow[i] != expected {
			t.Errorf("Lows error: Expected %f at cell %d got %f ", expected, i, resultLow[i])
		}
	}
}

func TestHlc3(t *testing.T) {
	hlc3 := Hlc3(testHigh, testLow, testClose)
	expectedHlc3 := []float64{7542.743333333333, 7495.649999999999, 7424.849999999999, 7255.433333333333, 7200.233333333333, 7185.603333333333, 7246.490000000001, 7111.836666666666, 7115.7699999999995, 6950.033333333333, 6700.356666666667, 7053.863333333334, 7173.8966666666665, 7161.61, 7150.420000000001, 7388.073333333334, 7419.003333333334, 7279.540000000001, 7191.110000000001, 7256.766666666666, 7184.456666666666, 7292.296666666666, 7394.81, 7268.046666666666, 7194.420000000001, 7188.676666666666, 7011.833333333333, 7198.25, 7332.203333333334, 7387.233333333334, 7640.733333333333, 8025.346666666667, 8126.7, 7864.563333333333, 8022.333333333333, 8100.740000000001, 8110.253333333334, 8113.606666666667, 8604.35, 8755.79, 8712.136666666667, 8856.456666666665, 8902.51, 8782.026666666667, 8624.546666666667, 8658.85, 8672.196666666667, 8441.463333333333, 8388.576666666666, 8339.896666666666, 8488.9, 8815.146666666667, 9229.746666666666, 9316.213333333333, 9412.023333333333, 9351.043333333333, 9373.886666666665, 9312.89, 9371.386666666667, 9194.323333333334, 9511.980000000001, 9711.733333333332, 9796.12, 9836.516666666666, 10079.183333333334, 9927.020000000002, 10119.94, 10363.71, 10265.723333333333, 10285.1, 10015.606666666667, 9854.206666666667, 9712.723333333333, 10025.256666666666, 9735.906666666668, 9571.980000000001, 9677.31, 9653.01, 9884.456666666667, 9720.08, 9407.446666666665, 8925.710000000001, 8770.416666666666, 8675.423333333334, 8619.266666666665, 8565.02, 8794.8, 8779.096666666666, 8755.68, 8997.813333333334, 9110.550000000001, 8990.85, 8310.1, 7921.126666666667, 7926.110000000001, 7838.446666666667, 5862.72, 5154.363333333334, 5294.46, 5476.903333333333}

	for i := range expectedHlc3 {
		if hlc3[i] != expectedHlc3[i] {
			t.Fatalf("Assertion error: expected %f, got %f", expectedHlc3[i], hlc3[i])
		}
	}
}

func TestCrossover(t *testing.T) {
	if Crossover(testCrossunder1, testCrossunder2) == true {
		t.Error("Crossover: Not expected and found")
	}

	if Crossover(testNothingCrossed1, testNothingCrossed2) == true {
		t.Error("Crossover: Not expected and found")
	}

	if Crossover(testCrossover1, testCrossover2) == false {
		t.Error("Crossover: Expected and not found")
	}
}

func TestCrossunder(t *testing.T) {
	// Crossunder
	series1 := []float64{1, 2, 3, 4, 8, 6, 7}
	series2 := []float64{1, 1, 10, 9, 5, 3, 7}

	if Crossunder(series1, series2) == false {
		t.Error("Crossunder: Expected and not found")
	}

	// Nothing
	series1 = []float64{1, 3, 2, 4, 8, 3, 8}
	series2 = []float64{1, 4, 5, 9, 5, 3, 7}

	if Crossunder(series1, series2) == true {
		t.Error("Crossunder: Not expected and found")
	}

	// Crossover
	series1 = []float64{1, 3, 2, 4, 8, 6, 7}
	series2 = []float64{1, 5, 1, 4, 5, 6, 7}

	if Crossunder(series1, series2) == true {
		t.Error("Crossunder: Not expected and found")
	}
}

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
