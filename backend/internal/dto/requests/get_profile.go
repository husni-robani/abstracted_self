package requests

type GetProfileRequest struct {
	Name bool `form:"name"`
	Summary bool `form:"summary"`
	Bio bool `form:"bio"`
	Taglines bool `form:"taglines"`
}