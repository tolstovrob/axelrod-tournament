package strategies

import (
	"fmt"

	"github.com/tolstovrob/axelrod-tournament/internal"
)

var names = map[string]string{
	"always_defect":          "Предатель",
	"always_cooperate":       "Добряк",
	"tit_for_tat":            "Око за око",
	"tit_for_two_tats":       "Око за два ока",
	"two_tits_for_tat":       "Два ока за око",
	"pavlov":                 "Павлов",
	"grudger":                "Злопамятный",
	"detective":              "Детектив",
	"soft_joss":              "Мягкий Джосс",
	"hard_joss":              "Жесткий Джосс",
	"forgiving_tit_for_tat":  "Прощающее око за око",
	"suspicious_tit_for_tat": "Подозрительное око за око",
	"spiteful":               "Мстительный",
	"alternator":             "Переключатель",
	"majority":               "Большинство",
	"prober":                 "Исследователь",
	"gradual":                "Постепенный",
	"generous_tit_for_tat":   "Великодушное око за око",
	"random":                 "Рандом",
}

func SpecName(s internal.Spec) string {
	switch s.Kind {
	case "adaptive":
		return fmt.Sprintf("Адаптивная (%.2f)", s.Params["threshold"])
	case "kamikaze":
		return fmt.Sprintf("Камикадзе (%d)", int(s.Params["betrayal_start"]))
	default:
		return defaultName(s.Kind)
	}
}

func defaultName(kind string) string {
	if n, ok := names[kind]; ok {
		return n
	}
	return kind
}

func New(spec internal.Spec) internal.Strategy {
	switch spec.Kind {
	case "always_defect":
		return AlwaysDefect{}
	case "always_cooperate":
		return AlwaysCooperate{}
	case "tit_for_tat":
		return TitForTat{}
	case "tit_for_two_tats":
		return TitForTwoTats{}
	case "two_tits_for_tat":
		return &TwoTitsForTat{}
	case "pavlov":
		return &Pavlov{}
	case "grudger":
		return &Grudger{}
	case "detective":
		return &Detective{}
	case "soft_joss":
		return SoftJoss{}
	case "hard_joss":
		return HardJoss{}
	case "forgiving_tit_for_tat":
		return &ForgivingTitForTat{}
	case "suspicious_tit_for_tat":
		return SuspiciousTitForTat{}
	case "spiteful":
		return &Spiteful{}
	case "alternator":
		return Alternator{}
	case "majority":
		return Majority{}
	case "prober":
		return &Prober{}
	case "gradual":
		return &Gradual{}
	case "generous_tit_for_tat":
		return GenerousTitForTat{}
	case "random":
		return Random{}
	case "adaptive":
		th := 0.5
		if v, ok := spec.Params["threshold"]; ok {
			th = v
		}
		return AdaptiveStrategy{Threshold: th}
	case "kamikaze":
		start := 120
		if v, ok := spec.Params["betrayal_start"]; ok {
			start = int(v)
		}
		return &Kamikaze{BetrayalStart: start}
	default:
		return AlwaysDefect{}
	}
}

func AllClassicSpecs() []internal.Spec {
	return []internal.Spec{
		{Kind: "always_defect"},
		{Kind: "always_cooperate"},
		{Kind: "tit_for_tat"},
		{Kind: "tit_for_two_tats"},
		{Kind: "two_tits_for_tat"},
		{Kind: "pavlov"},
		{Kind: "grudger"},
		{Kind: "detective"},
		{Kind: "soft_joss"},
		{Kind: "hard_joss"},
		{Kind: "forgiving_tit_for_tat"},
		{Kind: "suspicious_tit_for_tat"},
		{Kind: "spiteful"},
		{Kind: "alternator"},
		{Kind: "majority"},
		{Kind: "prober"},
		{Kind: "gradual"},
		{Kind: "generous_tit_for_tat"},
		{Kind: "adaptive", Params: map[string]float64{"threshold": 0.3}},
		{Kind: "adaptive", Params: map[string]float64{"threshold": 0.5}},
		{Kind: "adaptive", Params: map[string]float64{"threshold": 0.7}},
		{Kind: "kamikaze", Params: map[string]float64{"betrayal_start": 160}},
		{Kind: "kamikaze", Params: map[string]float64{"betrayal_start": 120}},
		{Kind: "random"},
	}
}
