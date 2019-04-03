package imagefilter

import "image/draw"
import "./edgedetection"

func CannyEdgeDetection(img draw.Image) draw.Image {
	result := Grayscale(img)
	result = GaussianBlur(result)

	var gradientAngles []float64
	result, gradientAngles = edgedetection.SobelGradientCalculation(result)
	result = edgedetection.NonMaxSuppression(result, gradientAngles)
	result, weakColor, strongColor := edgedetection.DoubleThreshold(result, 0.05, 0.09)
	result = edgedetection.Hysterisis(result, weakColor, strongColor)

	return result
}
