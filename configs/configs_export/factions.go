package configs_export

import (
	"github.com/darklab8/fl-configs/configs/settings/logus"
)

type Reputation struct {
	Name     string
	Rep      float64
	Empathy  float64
	Nickname string
}

type Faction struct {
	Name        string
	Infocard    Infocard
	ShortName   string
	Nickname    string
	Reputations []Reputation
}

func (e *Exporter) GetFactions() []Faction {
	var factions []Faction = make([]Faction, 0, 100)

	for _, group := range e.configs.InitialWorld.Groups {
		faction := Faction{
			Nickname: group.Nickname.Get(),
		}

		if faction_name, ok := e.configs.Infocards.Infonames[group.IdsName.Get()]; ok {
			faction.Name = string(faction_name)
		}

		if infocard, ok := e.configs.Infocards.Infocards[group.IdsInfo.Get()]; ok {
			infocard_parts, err := infocard.XmlToText()
			logus.Log.CheckError(err, "failed to xml infocard")
			faction.Infocard.Lines = append(faction.Infocard.Lines, infocard_parts...)
		}

		if short_name, ok := e.configs.Infocards.Infonames[group.IdsShortName.Get()]; ok {
			faction.ShortName = string(short_name)
		}

		empathy_rates, empathy_exists := e.configs.Empathy.RepoChangeMap.MapGetValue(faction.Nickname)

		for _, reputation := range group.Relationships {
			rep_to_add := &Reputation{}
			rep_to_add.Nickname = reputation.TargetNickname.Get()
			rep_to_add.Rep = reputation.Rep.Get()

			target_faction := e.configs.InitialWorld.GroupsMap.MapGet(rep_to_add.Nickname)

			if name, ok := e.configs.Infocards.Infonames[target_faction.IdsName.Get()]; ok {
				rep_to_add.Name = string(name)
			}

			if empathy_exists {
				if empathy_rate, ok := empathy_rates.EmpathyRatesMap.MapGetValue(rep_to_add.Nickname); ok {
					rep_to_add.Empathy = empathy_rate.RepoChange.Get()
				}
			}

			faction.Reputations = append(faction.Reputations, *rep_to_add)
		}

		factions = append(factions, faction)

	}

	return factions
}
