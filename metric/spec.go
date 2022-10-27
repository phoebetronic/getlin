package metric

type Interface interface {
	Get() Getter
	Set() Setter
}

type Getter interface {
	Mat() Matrix
	Sta() States
}

type Setter interface {
	Mat(int, int)
	Sta(float32)
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
	// Ppv returns the precision currently recorded by this confusion matrix.
	//
	//     TP / (TP + FP)
	//
	Ppv() float32
}

type States interface {
	// Nrm returns the normalized probability distribution across the currently
	// recorded state buckets.
	Nrm() [41]float32
	// Raw returns the currently recorded list of raw state buckets.
	Raw() [41]float32
}
