package intel_gpu

type Output struct {
	Period struct {
		Duration float64 `json:"duration"`
		Unit     string  `json:"unit"`
	} `json:"period"`
	Frequency struct {
		Requested float64 `json:"requested"`
		Actual    float64 `json:"actual"`
		Unit      string  `json:"unit"`
	} `json:"frequency"`
	Interrupts struct {
		Count float64 `json:"count"`
		Unit  string  `json:"unit"`
	} `json:"interrupts"`
	Rc6 struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"rc6"`
	Power struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"power"`
	ImcBandwidth struct {
		Reads  float64 `json:"reads"`
		Writes float64 `json:"writes"`
		Unit   string  `json:"unit"`
	} `json:"imc-bandwidth"`
	Engines struct {
		Render3D0 struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Render/3D/0"`
		Blitter0 struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Blitter/0"`
		Video0 struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Video/0"`
		Video1 struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"Video/1"`
		Videoenhance0 struct {
			Busy float64 `json:"busy"`
			Sema float64 `json:"sema"`
			Wait float64 `json:"wait"`
			Unit string  `json:"unit"`
		} `json:"VideoEnhance/0"`
	} `json:"engines"`
}
