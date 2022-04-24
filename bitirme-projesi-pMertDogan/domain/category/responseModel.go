package category

// import "github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain"

//we can use domain.ResponseModel or this one instead.
type ResponseModel struct{
	// Response    domain.ResponseModel
	ErrMsg string `json:"errMsg,omitempty"`
	ResponseCode int `json:"responseCode"`
	Data Categorys `json:"data,omitempty"`
}