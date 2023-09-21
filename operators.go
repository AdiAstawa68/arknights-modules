package arknightsmod

type Operator struct {
	Name,
	Class,
	Archetype,
	Region string
	HealthPoint,
	DeploymentPoint,
	AttackPower,
	Defense,
	Resistance,
	RedeploymentTime int
}

func (op Operator) introduction(username string) string {
	return "Hello Dr." + username + " my name is " + op.Name
}
func (op Operator) Onboard(username string) string {
	return op.introduction(username) + ", from " + op.Region + ", it's a pleasure to meet you!."
}
func GetAllOperator() []Operator {
	return []Operator{
		{Name: "Surtr", DeploymentPoint: 21},
		{Name: "Saria", DeploymentPoint: 22},
		{Name: "Phantom", DeploymentPoint: 10},
		{Name: "Yato", DeploymentPoint: 5},
		{Name: "Leizi", DeploymentPoint: 31},
	}
}
