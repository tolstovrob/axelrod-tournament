package internal

import (
	"sort"
)

type StrategyFactory func(Spec) Strategy

type Individual struct {
	Spec  Spec
	Score float64
	Wins  int
	Games int
}

type Population struct {
	Individuals []Individual
	Config      Config
	Noise       float64
	Factory     StrategyFactory
	NameOf      func(Spec) string
}

func NewPopulation(specs []Spec, cfg Config, noise float64, factory StrategyFactory, nameOf func(Spec) string) *Population {
	inds := make([]Individual, len(specs))
	for i, sp := range specs {
		inds[i] = Individual{Spec: sp}
	}
	return &Population{
		Individuals: inds,
		Config:      cfg,
		Noise:       noise,
		Factory:     factory,
		NameOf:      nameOf,
	}
}

func (p *Population) RunRoundRobin() {
	n := len(p.Individuals)
	for i := range p.Individuals {
		p.Individuals[i].Score = 0
		p.Individuals[i].Wins = 0
		p.Individuals[i].Games = 0
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s1 := p.Factory(p.Individuals[i].Spec)
			s2 := p.Factory(p.Individuals[j].Spec)

			res := PlayMatch(s1, s2, p.Config, p.Noise)

			p.Individuals[i].Score += float64(res.Score1)
			p.Individuals[j].Score += float64(res.Score2)
			p.Individuals[i].Games++
			p.Individuals[j].Games++

			if res.Score1 > res.Score2 {
				p.Individuals[i].Wins++
			} else if res.Score2 > res.Score1 {
				p.Individuals[j].Wins++
			}
		}
	}

	matches := float64(n - 1)
	if matches > 0 {
		for i := range p.Individuals {
			p.Individuals[i].Score /= matches
		}
	}
}

func (p *Population) SortedByScore() []Individual {
	sorted := make([]Individual, len(p.Individuals))
	copy(sorted, p.Individuals)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score > sorted[j].Score
	})
	return sorted
}

func (p *Population) AverageScore() float64 {
	if len(p.Individuals) == 0 {
		return 0
	}
	sum := 0.0
	for _, ind := range p.Individuals {
		sum += ind.Score
	}
	return sum / float64(len(p.Individuals))
}

func (p *Population) ScoreByKind(kind string) (avg float64, count int) {
	sum := 0.0
	for _, ind := range p.Individuals {
		if ind.Spec.Kind == kind {
			sum += ind.Score
			count++
		}
	}
	if count == 0 {
		return 0, 0
	}
	return sum / float64(count), count
}

func (p *Population) CooperationRate() float64 {
	n := len(p.Individuals)
	if n < 2 {
		return 0
	}

	totalCoop := 0
	totalRounds := 0

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			s1 := p.Factory(p.Individuals[i].Spec)
			s2 := p.Factory(p.Individuals[j].Spec)
			res := PlayMatch(s1, s2, p.Config, p.Noise)
			totalCoop += res.CooperationRounds
			totalRounds += res.TotalRounds
		}
	}

	if totalRounds == 0 {
		return 0
	}
	return float64(totalCoop) / float64(totalRounds)
}
