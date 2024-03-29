package configs_export

type Reputation struct {
	Name     string
	Rep      float64
	Empathy  float64
	Nickname string
}

type Faction struct {
	Name      string
	ShortName string
	Nickname  string

	ObjectDestruction float64
	MissionSuccess    float64
	MissionFailure    float64
	MissionAbort      float64

	InfonameID  int
	InfocardID  int
	Infocard    InfocardKey
	Reputations []Reputation
}

func (e *Exporter) GetFactions() []Faction {
	var factions []Faction = make([]Faction, 0, 100)

	for _, group := range e.configs.InitialWorld.Groups {
		var nickname string = group.Nickname.Get()
		faction := Faction{
			Nickname:   nickname,
			InfonameID: group.IdsName.Get(),
			InfocardID: group.IdsInfo.Get(),
			Infocard:   InfocardKey(nickname),
		}

		if faction_name, ok := e.configs.Infocards.Infonames[group.IdsName.Get()]; ok {
			faction.Name = string(faction_name)
		}

		e.infocards_parser.Set(InfocardKey(nickname), group.IdsInfo.Get())

		if short_name, ok := e.configs.Infocards.Infonames[group.IdsShortName.Get()]; ok {
			faction.ShortName = string(short_name)
		}

		empathy_rates, empathy_exists := e.configs.Empathy.RepoChangeMap.MapGetValue(faction.Nickname)

		if empathy_exists {
			faction.ObjectDestruction = empathy_rates.ObjectDestruction.Get()
			faction.MissionSuccess = empathy_rates.MissionSuccess.Get()
			faction.MissionFailure = empathy_rates.MissionFailure.Get()
			faction.MissionAbort = empathy_rates.MissionAbort.Get()
		}

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
