package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/tolstovrob/axelrod-tournament/internal"
	"github.com/tolstovrob/axelrod-tournament/strategies"
)

func newPop(specs []internal.Spec, cfg internal.Config, noise float64) *internal.Population {
	return internal.NewPopulation(specs, cfg, noise, strategies.New, strategies.SpecName)
}

func printHeader(w *tabwriter.Writer, title string) {
	fmt.Fprintf(w, "\n=== %s ===\n", title)
}

func printRanking(w *tabwriter.Writer, pop *internal.Population) {
	sorted := pop.SortedByScore()
	fmt.Fprintf(w, "Место\tСтратегия\tСр. счёт\tПобед\tИгр\n")
	for i, ind := range sorted {
		fmt.Fprintf(w, "%d\t%s\t%.1f\t%d\t%d\n",
			i+1, strategies.SpecName(ind.Spec), ind.Score, ind.Wins, ind.Games)
	}
}

func exp0Classic(cfg internal.Config) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, "Эксперимент 0: Классический турнир Аксельрода")

	pop := newPop(strategies.AllClassicSpecs(), cfg, 0)
	pop.RunRoundRobin()
	printRanking(w, pop)
	fmt.Fprintf(w, "\nСредний счёт популяции: %.1f\n", pop.AverageScore())
	w.Flush()
}

func exp1LoneNice(cfg internal.Config, nDefectors int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, fmt.Sprintf("Эксперимент 1: Один TFT vs %d предателей", nDefectors))

	specs := []internal.Spec{{Kind: "tit_for_tat"}}
	for i := 0; i < nDefectors; i++ {
		specs = append(specs, internal.Spec{Kind: "always_defect"})
	}

	pop := newPop(specs, cfg, 0)
	pop.RunRoundRobin()

	tftAvg, _ := pop.ScoreByKind("tit_for_tat")
	defAvg, defCount := pop.ScoreByKind("always_defect")

	fmt.Fprintf(w, "TFT avg score:\t%.1f\n", tftAvg)
	fmt.Fprintf(w, "Defect avg score:\t%.1f  (n=%d)\n", defAvg, defCount)
	fmt.Fprintf(w, "Разница (TFT - Defect):\t%.1f\n", tftAvg-defAvg)
	if tftAvg < defAvg {
		fmt.Fprintf(w, "\nВывод: одинокий TFT проигрывает по среднему score.\n")
	}
	w.Flush()
}

func exp2Cluster(cfg internal.Config, clusterSize, nDefectors int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, fmt.Sprintf(
		"Эксперимент 2: Кластер TFT (K=%d) vs %d предателей",
		clusterSize, nDefectors,
	))

	specs := make([]internal.Spec, 0, clusterSize+nDefectors)
	for i := 0; i < clusterSize; i++ {
		specs = append(specs, internal.Spec{Kind: "tit_for_tat"})
	}
	for i := 0; i < nDefectors; i++ {
		specs = append(specs, internal.Spec{Kind: "always_defect"})
	}

	pop := newPop(specs, cfg, 0)
	pop.RunRoundRobin()

	tftAvg, tftCount := pop.ScoreByKind("tit_for_tat")
	defAvg, defCount := pop.ScoreByKind("always_defect")

	fmt.Fprintf(w, "TFT avg score:\t%.1f  (n=%d)\n", tftAvg, tftCount)
	fmt.Fprintf(w, "Defect avg score:\t%.1f  (n=%d)\n", defAvg, defCount)
	fmt.Fprintf(w, "Разница (TFT - Defect):\t%.1f\n", tftAvg-defAvg)
	if tftAvg > defAvg {
		fmt.Fprintf(w, "\nВывод: при K=%d кластер TFT уже обгоняет дефекторов по среднему score.\n", clusterSize)
	} else {
		fmt.Fprintf(w, "\nВывод: при K=%d кластера ещё недостаточно.\n", clusterSize)
	}
	w.Flush()
}

func exp2ClusterSweep(cfg internal.Config, maxK, nDefectors int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, fmt.Sprintf(
		"Эксперимент 2 (sweep): K = 1..%d при %d предателях",
		maxK, nDefectors,
	))

	fmt.Fprintf(w, "K\tTFT ср.\tСр. пред.\tРазн.\tTFT выигрывает?\n")
	for k := 1; k <= maxK; k++ {
		specs := make([]internal.Spec, 0, k+nDefectors)
		for i := 0; i < k; i++ {
			specs = append(specs, internal.Spec{Kind: "tit_for_tat"})
		}
		for i := 0; i < nDefectors; i++ {
			specs = append(specs, internal.Spec{Kind: "always_defect"})
		}

		pop := newPop(specs, cfg, 0)
		pop.RunRoundRobin()

		tftAvg, _ := pop.ScoreByKind("tit_for_tat")
		defAvg, _ := pop.ScoreByKind("always_defect")
		delta := tftAvg - defAvg
		win := "нет"
		if delta > 0 {
			win = "да"
		}
		fmt.Fprintf(w, "%d\t%.1f\t%.1f\t%+.1f\t%s\n", k, tftAvg, defAvg, delta, win)
	}
	w.Flush()
}

func exp3Envy(cfg internal.Config) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, "Эксперимент 3: Не завидуй")

	specs := []internal.Spec{
		{Kind: "tit_for_tat"},
		{Kind: "always_defect"},
		{Kind: "always_cooperate"},
		{Kind: "grudger"},
		{Kind: "pavlov"},
		{Kind: "suspicious_tit_for_tat"},
		{Kind: "hard_joss"},
		{Kind: "soft_joss"},
		{Kind: "generous_tit_for_tat"},
		{Kind: "prober"},
		{Kind: "random"},
		{Kind: "alternator"},
	}

	pop := newPop(specs, cfg, 0)
	pop.RunRoundRobin()

	fmt.Fprintf(w, "\nРейтинг по общему счёту:\n")
	printRanking(w, pop)

	fmt.Fprintf(w, "\nСравнение побед в матчах и общему счёту:\n")
	fmt.Fprintf(w, "Стратегия\tПобед\tСр. счёт\tПобед/Игры\n")
	for _, ind := range pop.SortedByScore() {
		ratio := 0.0
		if ind.Games > 0 {
			ratio = float64(ind.Wins) / float64(ind.Games)
		}
		fmt.Fprintf(w, "%s\t%d\t%.1f\t%.2f\n",
			strategies.SpecName(ind.Spec), ind.Wins, ind.Score, ratio)
	}

	fmt.Fprintf(w, "\nВывод: стратегии с высоким Побед/Игры не всегда лидируют по среднему счёту.\n")
	fmt.Fprintf(w, "Tit-for-Tat часто проигрывает отдельные матчи, но набирает много очков в общем зачёте.\n")
	w.Flush()
}

func exp4Clarity(cfg internal.Config) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, "Эксперимент 4: Ясность vs сложность")

	clear := []internal.Spec{
		{Kind: "tit_for_tat"},
		{Kind: "tit_for_two_tats"},
		{Kind: "grudger"},
		{Kind: "pavlov"},
		{Kind: "always_cooperate"},
		{Kind: "always_defect"},
	}
	opaque := []internal.Spec{
		{Kind: "random"},
		{Kind: "detective"},
		{Kind: "prober"},
		{Kind: "adaptive", Params: map[string]float64{"threshold": 0.5}},
		{Kind: "kamikaze", Params: map[string]float64{"betrayal_start": 100}},
		{Kind: "alternator"},
	}

	pop := newPop(append(clear, opaque...), cfg, 0)
	pop.RunRoundRobin()

	fmt.Fprintf(w, "\nПолный рейтинг:\n")
	printRanking(w, pop)

	clearSum, clearN := 0.0, 0
	opaqueSum, opaqueN := 0.0, 0
	clearKinds := map[string]bool{
		"tit_for_tat": true, "tit_for_two_tats": true, "grudger": true,
		"pavlov": true, "always_cooperate": true, "always_defect": true,
	}
	for _, ind := range pop.Individuals {
		if clearKinds[ind.Spec.Kind] {
			clearSum += ind.Score
			clearN++
		} else {
			opaqueSum += ind.Score
			opaqueN++
		}
	}

	fmt.Fprintf(w, "\nСредний score простых стратегий:  %.1f\n", clearSum/float64(clearN))
	fmt.Fprintf(w, "Средний score сложных стратегий: %.1f\n", opaqueSum/float64(opaqueN))
	w.Flush()
}

func exp5Noise(cfg internal.Config, epsilons []float64) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, "Эксперимент 5: Влияние шума")

	specs := []internal.Spec{
		{Kind: "tit_for_tat"},
		{Kind: "generous_tit_for_tat"},
		{Kind: "forgiving_tit_for_tat"},
		{Kind: "soft_joss"},
		{Kind: "hard_joss"},
		{Kind: "grudger"},
		{Kind: "pavlov"},
		{Kind: "always_defect"},
		{Kind: "always_cooperate"},
	}

	fmt.Fprintf(w, "eps\tСтратегия\tСр. счёт\tКоопер.\n")
	for _, eps := range epsilons {
		pop := newPop(specs, cfg, eps)
		pop.RunRoundRobin()
		coop := pop.CooperationRate()

		for i, ind := range pop.SortedByScore() {
			marker := ""
			if i == 0 {
				marker = " ←"
			}
			fmt.Fprintf(w, "%.2f\t%s\t%.1f\t%.3f%s\n",
				eps, strategies.SpecName(ind.Spec), ind.Score, coop, marker)
		}
		fmt.Fprintf(w, "\n")
	}

	fmt.Fprintf(w, "Вывод: при росте eps чистый TFT теряет позиции относительно прощающих стратегий.\n")
	w.Flush()
}

func exp7Shadow(cfg internal.Config, lengths []int) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	printHeader(w, "Эксперимент 7: Длина тени будущего")

	baseSpecs := []internal.Spec{
		{Kind: "tit_for_tat"},
		{Kind: "always_defect"},
		{Kind: "always_cooperate"},
		{Kind: "grudger"},
		{Kind: "pavlov"},
		{Kind: "generous_tit_for_tat"},
		{Kind: "suspicious_tit_for_tat"},
		{Kind: "random"},
	}

	fmt.Fprintf(w, "Раунды\tTFT\tDefect\tCooperate\tGrudger\tPavlov\tGenerous\n")
	for _, rounds := range lengths {
		c := cfg
		c.Rounds = rounds
		pop := newPop(baseSpecs, c, 0)
		pop.RunRoundRobin()

		get := func(kind string) float64 {
			avg, _ := pop.ScoreByKind(kind)
			return avg
		}

		fmt.Fprintf(w, "%d\t%.1f\t%.1f\t%.1f\t%.1f\t%.1f\t%.1f\n",
			rounds,
			get("tit_for_tat"),
			get("always_defect"),
			get("always_cooperate"),
			get("grudger"),
			get("pavlov"),
			get("generous_tit_for_tat"),
		)
	}

	fmt.Fprintf(w, "\nВывод: при коротких играх преимущество смещается к более жадным стратегиям.\n")
	w.Flush()
}

func runAll(cfg internal.Config) {
	sep := strings.Repeat("─", 60)

	fmt.Println(sep)
	exp0Classic(cfg)

	fmt.Println(sep)
	exp1LoneNice(cfg, 20)

	fmt.Println(sep)
	exp2Cluster(cfg, 1, 20)
	exp2Cluster(cfg, 5, 20)
	exp2Cluster(cfg, 10, 20)

	fmt.Println(sep)
	exp2ClusterSweep(cfg, 12, 20)

	fmt.Println(sep)
	exp3Envy(cfg)

	fmt.Println(sep)
	exp4Clarity(cfg)

	fmt.Println(sep)
	exp5Noise(cfg, []float64{0.0, 0.01, 0.05, 0.10})

	fmt.Println(sep)
	exp7Shadow(cfg, []int{5, 10, 20, 50, 100, 200})

	fmt.Println(sep)
	fmt.Println("Все эксперименты завершены.")
}
