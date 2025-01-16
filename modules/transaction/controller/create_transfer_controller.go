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

		business := service.GetTransferService(store)

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

		// Update from account balance
		if err := UpdateAccountBalance(db, data.FromAccountId, -data.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update to account balance
		if err := UpdateAccountBalance(db, data.ToAccountId, data.Amount); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func CreateEntries(db *gorm.DB, data *model.EntriesCreation) error {
	store := database.NewSQLStore(db)

	business := service.GetEntriesService(store)

	if err := business.CreateEntries(data); err != nil {
		return err
	}

	return nil
}

func UpdateAccountBalance(db *gorm.DB, accountId int, amount int64) error {
	store := database.NewSQLStore(db)

	business := service.GetAccountService(store)

	data, err := business.GetAccountById(accountId)
	if err != nil {
		return err
	}

	updateData := model.AccountUpdate{Balance: data.Balance + amount}

	if err := business.UpdateAccount(&updateData, accountId); err != nil {
		return err
	}

	return nil
}
