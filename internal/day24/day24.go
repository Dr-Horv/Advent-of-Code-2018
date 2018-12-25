package day24

import (
	"errors"
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"regexp"
	"sort"
	"strings"
)

type groupType int

const (
	ImmuneSystem groupType = iota
	Infection
)

func (g groupType) opposite() groupType {
	if g == ImmuneSystem {
		return Infection
	} else {
		return ImmuneSystem
	}
}

func (g groupType) String() string {
	if g == ImmuneSystem {
		return "Immune"
	} else {
		return "Infection"
	}
}

type group struct {
	id           int
	groupType    groupType
	units        int
	health       int
	immunities   []string
	weaknesses   []string
	attackDamage int
	attackType   string
	initiative   int
}

func (g group) String() string {
	return fmt.Sprintf("%v %v u: %v h: %v i: %v w: %v dmg: %v, t: %v ini %v",
		g.id,
		g.groupType,
		g.units,
		g.health,
		g.immunities,
		g.weaknesses,
		g.attackDamage,
		g.attackType,
		g.initiative)
}

func (g group) idString() string {
	return fmt.Sprintf("%v %v", g.id, g.groupType)
}

type targeting struct {
	attacker *group
	defender *group
}

func (g group) isImmune(attacker *group) bool {
	for _, immunity := range g.immunities {
		if attacker.attackType == immunity {
			return true
		}
	}

	return false
}

func (g group) isWeak(attacker *group) bool {
	for _, weakness := range g.weaknesses {
		if attacker.attackType == weakness {
			return true
		}
	}

	return false
}

func (g group) effectivePower() int {
	return g.units * g.attackDamage
}

func Solve(lines []string, partOne bool) string {

	var re = regexp.MustCompile(`(?m)(\d+) units each with (\d+) hit points (.*)with an attack that does (\d+) (\w+) damage at initiative (\d+)`)

	addingToImmuneSystem := true
	immuneSystem := make([]*group, 0)
	infection := make([]*group, 0)
	infectionId := 1
	immuneId := 1
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			continue
		}

		if strings.Contains(l, "Immune System:") {
			addingToImmuneSystem = true
			continue
		} else if strings.Contains(l, "Infection:") {
			addingToImmuneSystem = false
			continue
		}

		matches := re.FindStringSubmatch(strings.TrimSpace(l))
		immunities, weaknesses := createImmunitiesAndWeaknesses(matches[3])
		groupType := ImmuneSystem
		id := immuneId
		if !addingToImmuneSystem {
			groupType = Infection
			id = infectionId
		}

		g := &group{
			id,
			groupType,
			StrConv(matches[1]),
			StrConv(matches[2]),
			immunities,
			weaknesses,
			StrConv(matches[4]),
			matches[5],
			StrConv(matches[6])}

		if addingToImmuneSystem {
			immuneSystem = append(immuneSystem, g)
			immuneId++
		} else {
			infection = append(infection, g)
			infectionId++
		}
	}

	if partOne {
		everyone, _ := battle(immuneSystem, infection)

		sum := 0
		for _, g := range everyone {
			sum += g.units
		}

		return fmt.Sprint(sum)
	}

	boost := 1
	var everyone []*group
	var err error
	for {
		boostedImmuneSystem := make([]*group, 0)
		for _, g := range immuneSystem {
			bg := &group{
				g.id,
				g.groupType,
				g.units,
				g.health,
				g.immunities,
				g.weaknesses,
				g.attackDamage + boost,
				g.attackType,
				g.initiative,
			}
			boostedImmuneSystem = append(boostedImmuneSystem, bg)
		}

		everyone, err = battle(boostedImmuneSystem, infection)
		if err == nil && everyone[0].groupType == ImmuneSystem {
			break
		}

		boost++
	}

	sum := 0
	for _, g := range everyone {
		sum += g.units
	}

	return fmt.Sprint(sum)

}

func copy(groups []*group) []*group {
	ng := make([]*group, 0)
	for _, g := range groups {
		bg := &group{
			g.id,
			g.groupType,
			g.units,
			g.health,
			g.immunities,
			g.weaknesses,
			g.attackDamage,
			g.attackType,
			g.initiative,
		}
		ng = append(ng, bg)
	}

	return ng
}

func battle(immuneSystem []*group, infection []*group) ([]*group, error) {
	immuneSystem = copy(immuneSystem)
	infection = copy(infection)
	everyone := make([]*group, 0)
	for {
		everyone = make([]*group, 0)
		immunities := 0
		for _, g := range immuneSystem {
			if g.units <= 0 {
				continue
			}
			immunities++
			everyone = append(everyone, g)
		}

		infections := 0
		for _, g := range infection {
			if g.units <= 0 {
				continue
			}
			infections++
			everyone = append(everyone, g)
		}

		// fmt.Printf("Infection: %v Immune: %v\n", infections, immunities)
		if infections == 0 || immunities == 0 {
			break
		}

		sort.Slice(everyone, func(i, j int) bool {
			g1 := everyone[i]
			g2 := everyone[j]

			if g1.effectivePower() > g2.effectivePower() {
				return true
			} else if g1.effectivePower() == g2.effectivePower() {
				return g1.initiative > g2.initiative
			}

			return false
		})

		targets := make(map[string]*targeting)

		mostDamage := -1
		var targetedGroup *group
		for _, ag := range everyone {
			opposite := ag.groupType.opposite()
			mostDamage = -1
			for _, dg := range everyone {

				_, found := targets[dg.idString()]
				if found {
					continue
				}

				if dg.groupType != opposite {
					continue
				}

				dmg := calculateDamage(ag, dg)
				if dmg == mostDamage {
					if dg.effectivePower() > targetedGroup.effectivePower() {
						targetedGroup = dg
					}
				} else if dmg > mostDamage {
					targetedGroup = dg
					mostDamage = dmg
				}
			}

			if mostDamage > 0 {
				targets[targetedGroup.idString()] = &targeting{ag, targetedGroup}
			}
		}

		attacking := make(map[string]*targeting)
		for _, t := range targets {
			attacking[t.attacker.idString()] = t
		}

		sort.Slice(everyone, func(i, j int) bool {
			g1 := everyone[i]
			g2 := everyone[j]
			return g1.initiative > g2.initiative
		})

		totalDeaths := 0
		for _, a := range everyone {
			if a.units == 0 {
				continue
			}

			targeting, found := attacking[a.idString()]
			if !found {
				continue
			}

			if targeting.attacker.units <= 0 {
				continue
			}

			dmg := calculateDamage(targeting.attacker, targeting.defender)
			deaths := int(float64(dmg) / float64(targeting.defender.health))
			totalDeaths += deaths

			targeting.defender.units -= deaths
		}

		//fmt.Printf("Deaths: %v\n", totalDeaths)

		if totalDeaths == 0 {
			return nil, errors.New("no deaths")
		}

	}
	return everyone, nil
}

func calculateDamage(attacker *group, defender *group) int {
	if defender.isImmune(attacker) {
		return 0
	}

	if defender.isWeak(attacker) {
		return attacker.effectivePower() * 2
	}

	return attacker.effectivePower()
}

func createImmunitiesAndWeaknesses(s string) ([]string, []string) {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return []string{}, []string{}
	}

	addingWeaknesses := false
	addingImmunities := false

	cleaned := strings.Replace(trimmed, "(", "", -1)
	cleaned = strings.Replace(cleaned, ")", "", -1)
	cleaned = strings.Replace(cleaned, ";", "", -1)

	weaknesses := make([]string, 0)
	immunities := make([]string, 0)

	for _, op := range strings.Split(cleaned, " ") {
		p := strings.TrimSpace(op)
		p = strings.Replace(p, ",", "", -1)

		if p == "to" {
			continue
		}

		if p == "weak" {
			addingWeaknesses = true
			addingImmunities = false
			continue
		}

		if p == "immune" {
			addingWeaknesses = false
			addingImmunities = true
			continue
		}

		if addingImmunities {
			immunities = append(immunities, p)
		}

		if addingWeaknesses {
			weaknesses = append(weaknesses, p)
		}
	}

	return immunities, weaknesses
}
