package transaction

import (
	"gogo/common"
	"gogo/modules/transaction/database"
	"gogo/modules/transaction/model"
	"gogo/modules/transaction/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransfer(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.TransfersCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := database.NewSQLStore(db)

		business := service.GetCreateTransferService(store)

		//Create transfer
		if err := business.CreateTransfer(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Create entries from account
		entriesFromData := model.EntriesCreation{
			AccountId: data.FromAccountId,
			Amount:    -data.Amount,
		}

		if err := CreateEntries(db, &entriesFromData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//Create entries to account
		entriesToData := model.EntriesCreation{
			AccountId: data.ToAccountId,
			Amount:    data.Amount,
		}

		if err := CreateEntries(db, &entriesToData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func CreateEntries(db *gorm.DB, data *model.EntriesCreation) error {
	store := database.NewSQLStore(db)

	business := service.GetCreateEntriesService(store)

	if err := business.CreateEntries(data); err != nil {
		return err
	}

	return nil
}
