/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package controllers

import "Syns/servers/syns-tokens-ms/dao"

// @notice Root struct for other methods in controller
type SynsTokenController struct {
	SynsTokenDao dao.SynsTokenDao
}

// @dev Constructor
func SynsTokenControllerConstructor(synsTokenDao dao.SynsTokenDao) *SynsTokenController{
	return &SynsTokenController {
		SynsTokenDao: synsTokenDao,
	}
}
