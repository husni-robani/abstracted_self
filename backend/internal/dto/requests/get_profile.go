package requests

type GetProfileRequest struct {
	Name bool `form:"name"`
	Summary bool `form:"summary"`
	Bio bool `form:"bio"`
	Taglines bool `form:"taglines"`
	Resume bool `form:"resume"`
	SkillSet bool `form:"skill_set"`
	Cache bool `form:"cache"`
}