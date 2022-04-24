package order

import (
	"errors"

	"github.com/pMertDogan/picusGoBackend--Patika/picusBootCampFinalProject/domain/basket"
	"go.uber.org/zap"
)

/*
//Create handler DTO is strong data about what is needed to create order
	We need basketID to understand which basket items are used in order
*/
type CompleteOrderDTO struct {
	BasketIDs []int `json:"basketIDs,required"` //	
	Comment string `json:"comment,required"` //
	ShipingAddress string `json:"shipingAddress,required"` //
	BillinAddress string `json:"billingAddress,required"` //
}


//check provided BasketIDs are valid or not
func (dto CompleteOrderDTO) ValideteBasketIDs(userID int) error {
	
	//Get all basket items for user without pagination
	//We dont added basket limit so if user has many basket items we will get all of them and its add extra load to server
	verifiedBaskets, err := basket.Repo().GetBasketsByUserID(userID)
	zap.L().Debug("verifiedBaskets", zap.Any("verifiedBaskets are ", verifiedBaskets))
	if err != nil {
		return  err
	}


	//check is dto.BasketIDs are unique if not return error
	if !UniqueIntSlice(dto.BasketIDs) {
		return errors.New("BasketIDs must be unique")
	}

	//Check each of  dto.BasketIDs in baskets.BasketIDs
	for _, basketID := range dto.BasketIDs {
		//Check if basketID is in baskets.BasketIDs
		if !BasketIDsInBaskets(basketID, verifiedBaskets) {
			//If not return error. 
			return  errors.New("BasketIDs is not valid for customer")
		}
	}
	//return true to indicate that basketIDs are valid
	return nil
}

//check is user send correct basketIDs
func BasketIDsInBaskets(basketIDtoCheck int, verifiedBaskets basket.BasketSToResponseDTO) bool {
	for _, basket := range verifiedBaskets {
		if int(basket.ID) == basketIDtoCheck {
			return true
		}
	}
	return false
}

//check is int slice has unique items
func UniqueIntSlice(slice []int) bool {
	//make map to store unique items
	m := make(map[int]bool)
	for _, item := range slice {
		if _, ok := m[item]; ok {
			return false
		}
		m[item] = true
	}
	return true
}