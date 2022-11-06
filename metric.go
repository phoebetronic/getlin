package getlin

type Metric interface {
	Get() Getter
	Set() Setter
}

type Getter interface {
	Mat() Matrix
	Sta() States
}

type Setter interface {
	Mat(int, int)
	Sta(int, int)
}

type Matrix interface {
	// Acc returns the accuracy currently recorded by this confusion matrix.
	//
	//     (TP + TN) / (TP + TN + FP + FN)
	//
	Acc() float32
	// Mae returns the mean absolute error of the currently recorded predicted
	// and true labels.
	//
	//     (FP + FN) / (TP + TN + FP + FN)
	//
	Mae() float32
	// Raw returns the currently recorded list of raw confusion matrix states.
	Raw() [4]int
}

type States interface {
	// Ind takes the ratio to be mapped to its raw interal state bucket. A
	// Clause ratio can be used to update the states distribution like shown
	// below.
	//
	//     met.Set().Sta(met.Get().Sta().Ind(rat), 1)
	//
	Ind(float32) int
	// Nrm returns the normalized probability distribution across the currently
	// recorded state buckets.
	Nrm() [41]float32
	// Raw returns the currently recorded list of raw state buckets.
	Raw() [41]int
}
