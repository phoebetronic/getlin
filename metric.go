package getlin

type Metric interface {
	Get() Getter
	Set() Setter
}

type Getter interface {
	Err() Errors
	Sta() States
}

type Setter interface {
	Err(float32, float32)
	Sta(int, int)
}

type Errors interface {
	// Mae returns the mean absolute error of the currently recorded true
	// weights and weight predictions.
	Mae() float32
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
