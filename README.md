# Турнирный менеджер по итеративной дилемме заключённого на Go

Проект позволяет проводить турниры между стратегиями в повторяющейся дилемме заключённого и проверять тезисы Аксельрода на популяциях.

## Задача

Два игрока могут сотрудничать или предавать. Предательство даёт максимум в одном раунде, но при повторных встречах взаимное сотрудничество выгоднее обоим. Стратегия ищет баланс в повторяющихся взаимодействиях &mdash; с «тенью будущего».

## Быстрый старт

```
just play          # классический круговой турнир -> tournament.csv
just exp0          # тот же турнир в формате эксперимента
just exp1          # один TFT против предателей
just exp2          # кластер TFT
just exp2s         # итерация по размеру кластера
just exp3          # "не завидуй": сравнение побед в матче с общим зачётом
just exp4          # ясность против сложности
just exp5          # шум
just exp7          # разная длина матча
just all           # все эксперименты подряд
```

Для не-just-enjoyers:

```
go run ./cmd/play/
go run ./cmd/experiments/ -e 1 -defectors 20 -rounds 200
go run ./cmd/experiments/ -e 2 -cluster 5 -defectors 20
```

Флаги CLI экспериментов: `-e`, `-rounds`, `-defectors`, `-cluster`.

## Эксперименты

| Код        | Суть                                          |
| ---------- | --------------------------------------------- |
| `0`        | Классический турнир Аксельрода                |
| `1`        | Один Tit-for-Tat в море Always-Defect         |
| `2` / `2s` | Кластер TFT; sweep по K                       |
| `3`        | Победы в отдельных матчах vs суммарный score  |
| `4`        | Простые clear-стратегии vs непрозрачные       |
| `5`        | Шум: случайная инверсия хода с вероятностью ε |
| `7`        | Разная длина матча (тень будущего)            |

Сценарии живут в `cmd/experiments/`.

## Создание стратегии

Реализуйте интерфейс:

```go
type Strategy interface {
	Name() string
	Move(opponentHistory History) Action
	Reset()
}
```

Доступна только история ходов соперника. Внутреннее состояние сбрасывайте в `Reset()`.

Параметрические стратегии &mdash; публичные поля структуры. Чтобы стратегия участвовала в экспериментах через фабрику, добавьте ветку в `strategies.New` и имя в `SpecName`.

## Классический турнир

```go
config := internal.Config{
	Rounds:       200,
	PayoffMatrix: internal.DefaultPayoffMatrix(),
}

tm := internal.NewTournamentManager(config)
for _, spec := range strategies.AllClassicSpecs() {
	tm.AddStrategy(strategies.New(spec))
}

tm.PlayTournament()
tm.PrintResults()
_ = tm.ExportResultsToCSV("results.csv")
```

Или точечно:

```go
tm.AddStrategy(strategies.TitForTat{})
tm.AddStrategy(&strategies.Kamikaze{BetrayalStart: 120})
```

## Структура

```
cmd/play/           — классический турнир
cmd/experiments/    — сценарии экспериментов
internal/           — Config, PlayMatch, Population, TournamentManager
strategies/         — реализации + factory
```

## Источник

Основано на работах Роберта Аксельрода по эволюции сотрудничества:  
https://cs.stanford.edu/people/eroberts/courses/soco/projects/1998-99/game-theory/axelrod.html
