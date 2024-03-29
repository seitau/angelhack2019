package main

func detectArea(lat float64, lng float64) [4]float64 {
	if 42.35 <= lat && lat <= 43.20 && 140.45 <= lng && lng <= 141.525 {
		return [4]float64{42.35, 43.20, 140.45, 141.525}
	} else if 37.58 <= lat && lat <= 38.4175 && 138.7515 <= lng && lng <= 139.62625 {
		return [4]float64{37.58, 38.4175, 138.7515, 139.62625}
	} else if 37.9185 <= lat && lat <= 38.835 && 140.835 <= lng && lng <= 141.2505 {
		return [4]float64{37.9185, 38.835, 140.835, 141.2505}
	} else if 35.20 <= lat && lat <= 36.12525 && 139.075 <= lng && lng <= 140.2505 {
		return [4]float64{35.20, 36.12525, 139.075, 140.2505}
	} else if 36.12525 <= lat && lat <= 36.835 && 136.075 <= lng && lng <= 137.2505 {
		return [4]float64{36.12525, 36.835, 136.075, 137.2505}
	} else if 34.835 <= lat && lat <= 35.10 && 136.075 <= lng && lng <= 137.2505 {
		return [4]float64{34.835, 35.10, 136.075, 137.2505}
	} else if 34.9185 <= lat && lat <= 35.668 && 136.225 <= lng && lng <= 137.225 {
		return [4]float64{34.9185, 35.668, 136.225, 137.225}
	} else if 34.2505 <= lat && lat <= 34.9185 && 134.525 <= lng && lng <= 136.0 {
		return [4]float64{34.2505, 34.9185, 134.525, 136.0}
	} else if 33.9185 <= lat && lat <= 34.35 && 134.835 <= lng && lng <= 133.835 {
		return [4]float64{33.9185, 34.35, 134.835, 133.835}
	} else if 34.12525 <= lat && lat <= 34.835 && 131.525 <= lng && lng <= 133.00 {
		return [4]float64{34.12525, 34.835, 131.525, 133.00}
	} else if 33.20 <= lat && lat <= 34.12525 && 130.0 <= lng && lng <= 131.075 {
		return [4]float64{33.20, 34.12525, 130.0, 131.075}
	} else if 31.10 <= lat && lat <= 31.835 && 130.0 <= lng && lng <= 131.0 {
		return [4]float64{31.10, 31.835, 130.0, 131.0}
	}
	return [4]float64{0, 0, 0, 0}
}
