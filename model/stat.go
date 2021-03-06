package model

type Stat struct {
	Hp             int32 `json:"hp" bson:"hp"`
	Attack         int32 `json:"attack" bson:"attack"`
	Defense        int32 `json:"defense" bson:"defense"`
	SpecialAttack  int32 `json:"specialAttack" bson:"special_attack"`
	SpecialDefense int32 `json:"specialDefense" bson:"special_defense"`
	Speed          int32 `json:"speed" bson:"speed"`
}

func (s *Stat) valid() bool {
	return (s.Hp > 0 &&
		s.Attack > 0 &&
		s.Defense > 0 &&
		s.SpecialAttack > 0 &&
		s.SpecialDefense > 0 &&
		s.Speed > 0)
}
